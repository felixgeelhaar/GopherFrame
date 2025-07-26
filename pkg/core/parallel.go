package core

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// ParallelStrategy defines different approaches to parallelization
type ParallelStrategy int

const (
	StrategyAuto ParallelStrategy = iota
	StrategyFixedChunks
	StrategyCacheFriendly
	StrategyMemoryBound
)

// ParallelOptions configures parallel processing behavior
type ParallelOptions struct {
	NumWorkers      int              // Number of parallel workers (0 = auto)
	ChunkSize       int64            // Rows per chunk (0 = auto)
	MinParallelRows int64            // Minimum rows to enable parallelization
	MemoryPool      memory.Allocator // Shared memory pool
	Strategy        ParallelStrategy // Chunking strategy
	Context         context.Context  // Cancellation context
}

// DefaultParallelOptions returns sensible defaults for parallel processing
func DefaultParallelOptions() *ParallelOptions {
	return &ParallelOptions{
		NumWorkers:      runtime.NumCPU(),
		ChunkSize:       0, // Auto-calculate
		MinParallelRows: 10000,
		MemoryPool:      memory.NewGoAllocator(),
		Strategy:        StrategyAuto,
		Context:         context.Background(),
	}
}

// ChunkSpec defines a chunk of work for parallel processing
type ChunkSpec struct {
	StartRow int64
	EndRow   int64
	ChunkID  int
}

// ChunkProcessFunc defines the signature for chunk processing functions
type ChunkProcessFunc func(chunk ChunkSpec, data *DataFrame, pool memory.Allocator) (arrow.Array, error)

// ChunkResult holds the result of processing a chunk
type ChunkResult struct {
	ChunkID int
	Result  arrow.Array
	Error   error
}

// ParallelProcessor manages parallel execution of DataFrame operations
type ParallelProcessor struct {
	options *ParallelOptions
	workers int
}

// NewParallelProcessor creates a new parallel processor with the given options
func NewParallelProcessor(opts *ParallelOptions) *ParallelProcessor {
	if opts == nil {
		opts = DefaultParallelOptions()
	}

	// Ensure context is never nil
	if opts.Context == nil {
		opts.Context = context.Background()
	}

	// Ensure memory pool is never nil
	if opts.MemoryPool == nil {
		opts.MemoryPool = memory.NewGoAllocator()
	}

	workers := opts.NumWorkers
	if workers <= 0 {
		workers = runtime.NumCPU()
	}

	return &ParallelProcessor{
		options: opts,
		workers: workers,
	}
}

// ShouldParallelize determines if parallelization is beneficial for the given DataFrame
func (p *ParallelProcessor) ShouldParallelize(df *DataFrame) bool {
	return df.NumRows() >= p.options.MinParallelRows && p.workers > 1
}

// CalculateChunks divides the DataFrame into optimal chunks for parallel processing
func (p *ParallelProcessor) CalculateChunks(df *DataFrame) []ChunkSpec {
	totalRows := df.NumRows()

	if !p.ShouldParallelize(df) {
		return []ChunkSpec{{StartRow: 0, EndRow: totalRows, ChunkID: 0}}
	}

	var chunkSize int64
	if p.options.ChunkSize > 0 {
		chunkSize = p.options.ChunkSize
	} else {
		chunkSize = p.calculateOptimalChunkSize(totalRows)
	}

	chunks := make([]ChunkSpec, 0)
	chunkID := 0

	for start := int64(0); start < totalRows; start += chunkSize {
		end := start + chunkSize
		if end > totalRows {
			end = totalRows
		}

		chunks = append(chunks, ChunkSpec{
			StartRow: start,
			EndRow:   end,
			ChunkID:  chunkID,
		})
		chunkID++
	}

	return chunks
}

// calculateOptimalChunkSize determines the best chunk size based on strategy and system characteristics
func (p *ParallelProcessor) calculateOptimalChunkSize(totalRows int64) int64 {
	switch p.options.Strategy {
	case StrategyFixedChunks:
		return totalRows / int64(p.workers)
	case StrategyCacheFriendly:
		// Aim for chunks that fit in L3 cache (assume 8MB)
		return min(totalRows/int64(p.workers), 100000)
	case StrategyMemoryBound:
		// Conservative chunks to limit memory usage
		return min(totalRows/int64(p.workers*2), 50000)
	default: // StrategyAuto
		// Balance between parallelism and overhead
		optimalSize := totalRows / int64(p.workers*2)
		return max(min(optimalSize, 100000), 10000)
	}
}

// ProcessChunks executes the given processing function on all chunks in parallel
func (p *ParallelProcessor) ProcessChunks(df *DataFrame, chunks []ChunkSpec, processFunc ChunkProcessFunc) ([]arrow.Array, error) {
	if len(chunks) == 1 {
		// Single chunk - no need for parallelization
		result, err := processFunc(chunks[0], df, p.options.MemoryPool)
		if err != nil {
			return nil, err
		}
		return []arrow.Array{result}, nil
	}

	// Parallel processing
	resultChan := make(chan ChunkResult, len(chunks))
	semaphore := make(chan struct{}, p.workers)

	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go func(c ChunkSpec) {
			defer wg.Done()

			// Acquire worker slot
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Check for cancellation
			select {
			case <-p.options.Context.Done():
				resultChan <- ChunkResult{
					ChunkID: c.ChunkID,
					Error:   p.options.Context.Err(),
				}
				return
			default:
			}

			// Create chunk-specific DataFrame
			chunkDF, err := p.createChunkDataFrame(df, c)
			if err != nil {
				resultChan <- ChunkResult{
					ChunkID: c.ChunkID,
					Error:   fmt.Errorf("failed to create chunk DataFrame: %w", err),
				}
				return
			}
			defer chunkDF.Release()

			// Process the chunk
			result, err := processFunc(c, chunkDF, p.options.MemoryPool)
			resultChan <- ChunkResult{
				ChunkID: c.ChunkID,
				Result:  result,
				Error:   err,
			}
		}(chunk)
	}

	// Wait for all workers to complete
	wg.Wait()
	close(resultChan)

	// Collect and order results
	results := make([]arrow.Array, len(chunks))
	var firstError error

	for result := range resultChan {
		if result.Error != nil && firstError == nil {
			firstError = result.Error
		}
		if result.Result != nil {
			results[result.ChunkID] = result.Result
		}
	}

	if firstError != nil {
		// Clean up any successful results
		for _, result := range results {
			if result != nil {
				result.Release()
			}
		}
		return nil, firstError
	}

	return results, nil
}

// createChunkDataFrame creates a new DataFrame containing only the specified row range
func (p *ParallelProcessor) createChunkDataFrame(df *DataFrame, chunk ChunkSpec) (*DataFrame, error) {
	schema := df.record.Schema()
	chunkSize := chunk.EndRow - chunk.StartRow

	if chunkSize <= 0 {
		return nil, fmt.Errorf("invalid chunk size: %d", chunkSize)
	}

	// Create new arrays for the chunk
	chunkColumns := make([]arrow.Array, len(schema.Fields()))

	for colIdx := range schema.Fields() {
		column := df.record.Column(colIdx)
		chunkArray, err := p.sliceArray(column, chunk.StartRow, chunk.EndRow)
		if err != nil {
			// Clean up any arrays created so far
			for i := 0; i < colIdx; i++ {
				if chunkColumns[i] != nil {
					chunkColumns[i].Release()
				}
			}
			return nil, fmt.Errorf("failed to slice column %d: %w", colIdx, err)
		}
		chunkColumns[colIdx] = chunkArray
	}

	// Create new record for the chunk
	chunkRecord := array.NewRecord(schema, chunkColumns, chunkSize)

	// Clean up chunk columns (record retains them)
	for _, col := range chunkColumns {
		col.Release()
	}

	return NewDataFrame(chunkRecord), nil
}

// sliceArray creates a new array containing elements from start to end (exclusive)
func (p *ParallelProcessor) sliceArray(arr arrow.Array, start, end int64) (arrow.Array, error) {
	if start < 0 || end > int64(arr.Len()) || start >= end {
		return nil, fmt.Errorf("invalid slice range [%d:%d) for array of length %d", start, end, arr.Len())
	}

	// Use Arrow's slicing capability
	sliced := array.NewSlice(arr, start, end)
	sliced.Retain() // Ensure the slice is retained for the caller
	return sliced, nil
}

// CombineResults concatenates multiple arrays into a single array of the same type
func (p *ParallelProcessor) CombineResults(results []arrow.Array) (arrow.Array, error) {
	if len(results) == 0 {
		return nil, fmt.Errorf("no results to combine")
	}

	if len(results) == 1 {
		results[0].Retain()
		return results[0], nil
	}

	// Verify all arrays have the same type
	expectedType := results[0].DataType()
	for i, result := range results {
		if !arrow.TypeEqual(result.DataType(), expectedType) {
			return nil, fmt.Errorf("type mismatch at index %d: expected %s, got %s",
				i, expectedType, result.DataType())
		}
	}

	// Calculate total length
	totalLen := 0
	for _, result := range results {
		totalLen += result.Len()
	}

	// Create builder for the combined result
	builder, err := p.createBuilderForType(expectedType, totalLen)
	if err != nil {
		return nil, fmt.Errorf("failed to create builder: %w", err)
	}
	defer builder.Release()

	// Append all results
	for _, result := range results {
		if err := p.appendArrayToBuilder(builder, result); err != nil {
			return nil, fmt.Errorf("failed to append array: %w", err)
		}
	}

	return builder.NewArray(), nil
}

// createBuilderForType creates an appropriate builder for the given Arrow type
func (p *ParallelProcessor) createBuilderForType(dataType arrow.DataType, capacity int) (array.Builder, error) {
	pool := p.options.MemoryPool

	switch dataType.ID() {
	case arrow.BOOL:
		builder := array.NewBooleanBuilder(pool)
		builder.Reserve(capacity)
		return builder, nil
	case arrow.INT64:
		builder := array.NewInt64Builder(pool)
		builder.Reserve(capacity)
		return builder, nil
	case arrow.FLOAT64:
		builder := array.NewFloat64Builder(pool)
		builder.Reserve(capacity)
		return builder, nil
	case arrow.STRING:
		builder := array.NewStringBuilder(pool)
		builder.Reserve(capacity)
		return builder, nil
	default:
		return nil, fmt.Errorf("unsupported data type for parallel processing: %s", dataType)
	}
}

// appendArrayToBuilder appends all elements from the array to the builder
func (p *ParallelProcessor) appendArrayToBuilder(builder array.Builder, arr arrow.Array) error {
	switch b := builder.(type) {
	case *array.BooleanBuilder:
		boolArr := arr.(*array.Boolean)
		for i := 0; i < arr.Len(); i++ {
			if boolArr.IsNull(i) {
				b.AppendNull()
			} else {
				b.Append(boolArr.Value(i))
			}
		}
	case *array.Int64Builder:
		intArr := arr.(*array.Int64)
		for i := 0; i < arr.Len(); i++ {
			if intArr.IsNull(i) {
				b.AppendNull()
			} else {
				b.Append(intArr.Value(i))
			}
		}
	case *array.Float64Builder:
		floatArr := arr.(*array.Float64)
		for i := 0; i < arr.Len(); i++ {
			if floatArr.IsNull(i) {
				b.AppendNull()
			} else {
				b.Append(floatArr.Value(i))
			}
		}
	case *array.StringBuilder:
		strArr := arr.(*array.String)
		for i := 0; i < arr.Len(); i++ {
			if strArr.IsNull(i) {
				b.AppendNull()
			} else {
				b.Append(strArr.Value(i))
			}
		}
	default:
		return fmt.Errorf("unsupported builder type: %T", builder)
	}

	return nil
}

// Utility functions
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
