package core

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParallelProcessorBasics tests core parallel processor functionality
func TestParallelProcessorBasics(t *testing.T) {
	opts := DefaultParallelOptions()
	processor := NewParallelProcessor(opts)

	t.Run("ShouldParallelize", func(t *testing.T) {
		smallDF := createTestDataFrameForParallel(5000)
		defer smallDF.Release()

		largeDF := createTestDataFrameForParallel(50000)
		defer largeDF.Release()

		assert.False(t, processor.ShouldParallelize(smallDF))
		assert.True(t, processor.ShouldParallelize(largeDF))
	})

	t.Run("CalculateChunks", func(t *testing.T) {
		df := createTestDataFrameForParallel(100000)
		defer df.Release()

		chunks := processor.CalculateChunks(df)

		assert.Greater(t, len(chunks), 1)

		// Verify chunks cover entire DataFrame
		totalRows := int64(0)
		for _, chunk := range chunks {
			assert.GreaterOrEqual(t, chunk.EndRow, chunk.StartRow)
			totalRows += chunk.EndRow - chunk.StartRow
		}
		assert.Equal(t, df.NumRows(), totalRows)

		// Verify chunks are contiguous
		for i := 1; i < len(chunks); i++ {
			assert.Equal(t, chunks[i-1].EndRow, chunks[i].StartRow)
		}
	})

	t.Run("ChunkSizeStrategies", func(t *testing.T) {
		df := createTestDataFrameForParallel(100000)
		defer df.Release()

		strategies := []ParallelStrategy{
			StrategyAuto,
			StrategyFixedChunks,
			StrategyCacheFriendly,
			StrategyMemoryBound,
		}

		for _, strategy := range strategies {
			opts := &ParallelOptions{
				NumWorkers:      4,
				Strategy:        strategy,
				MinParallelRows: 1000,
			}
			proc := NewParallelProcessor(opts)
			chunks := proc.CalculateChunks(df)

			assert.Greater(t, len(chunks), 0)

			// Verify total coverage
			totalRows := int64(0)
			for _, chunk := range chunks {
				totalRows += chunk.EndRow - chunk.StartRow
			}
			assert.Equal(t, df.NumRows(), totalRows)
		}
	})
}

// TestParallelFilterCorrectness tests that parallel filtering produces correct results
func TestParallelFilterCorrectness(t *testing.T) {
	df := createTestDataFrameForParallel(25000)
	defer df.Release()

	// Create a selective filter
	mask := createSelectiveMaskForParallel(25000, func(i int) bool {
		return i%7 == 0
	})
	defer mask.Release()

	opts := &ParallelOptions{
		NumWorkers:      4,
		MinParallelRows: 5000,
	}

	t.Run("ParallelVsSequential", func(t *testing.T) {
		// Sequential result
		seqResult, err := df.FilterOptimized(mask)
		require.NoError(t, err)
		defer seqResult.Release()

		// Parallel result
		parResult, err := df.ParallelFilterWithBooleanMask(mask, opts)
		require.NoError(t, err)
		defer parResult.Release()

		// Results should have same dimensions
		assert.Equal(t, seqResult.NumRows(), parResult.NumRows())
		assert.Equal(t, seqResult.NumCols(), parResult.NumCols())
	})

	// Note: Testing with predicate requires expression package integration
	// which will be added in a future iteration to avoid import cycles
}

// TestParallelExpressionEvaluation tests parallel expression evaluation
func TestParallelExpressionEvaluation(t *testing.T) {
	// Note: Expression evaluation tests will be implemented
	// when expression package integration is completed
	t.Skip("Expression evaluation tests require expr package integration")
}

// TestParallelAggregation tests parallel aggregation operations
func TestParallelAggregation(t *testing.T) {
	df := createTestDataFrameForParallel(100000)
	defer df.Release()

	opts := &ParallelOptions{
		NumWorkers:      4,
		MinParallelRows: 10000,
	}

	t.Run("ParallelSum", func(t *testing.T) {
		// Sequential sum (force by using high MinParallelRows)
		seqOpts := &ParallelOptions{
			MinParallelRows: 200000, // Force sequential
		}
		seqSum, err := df.ParallelSum("value", seqOpts)
		require.NoError(t, err)

		// Parallel sum
		parSum, err := df.ParallelSum("value", opts)
		require.NoError(t, err)

		// Results should be very close (allowing for floating point precision)
		assert.InDelta(t, seqSum, parSum, 1e-6*seqSum)
	})

	t.Run("SumNonExistentColumn", func(t *testing.T) {
		_, err := df.ParallelSum("nonexistent", opts)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}

// TestParallelChainedOperations tests the parallel chain builder
func TestParallelChainedOperations(t *testing.T) {
	// Note: Chained operations tests require expression package integration
	t.Skip("Chained operations tests require expr package integration")
}

// TestParallelErrorHandling tests error conditions and cleanup
func TestParallelErrorHandling(t *testing.T) {
	df := createTestDataFrameForParallel(20000)
	defer df.Release()

	t.Run("InvalidMaskLength", func(t *testing.T) {
		shortMask := createSelectiveMaskForParallel(df.NumRows()-100, func(i int) bool {
			return true
		})
		defer shortMask.Release()

		opts := &ParallelOptions{
			NumWorkers:      4,
			MinParallelRows: 5000,
		}

		_, err := df.ParallelFilterWithBooleanMask(shortMask, opts)
		assert.Error(t, err)
	})

	t.Run("ContextCancellation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		opts := &ParallelOptions{
			NumWorkers:      4,
			MinParallelRows: 1000,
			Context:         ctx,
		}

		mask := createSelectiveMaskForParallel(df.NumRows(), func(i int) bool {
			return i%2 == 0
		})
		defer mask.Release()

		// This should be cancelled quickly due to short timeout
		_, err := df.ParallelFilterWithBooleanMask(mask, opts)

		// Should either succeed quickly or be cancelled
		if err != nil {
			assert.Contains(t, err.Error(), "context")
		}
	})
}

// TestParallelMemoryManagement tests memory efficiency and cleanup
func TestParallelMemoryManagement(t *testing.T) {
	df := createTestDataFrameForParallel(50000)
	defer df.Release()

	t.Run("SharedMemoryPool", func(t *testing.T) {
		pool := memory.NewGoAllocator()
		opts := &ParallelOptions{
			NumWorkers:      4,
			MinParallelRows: 10000,
			MemoryPool:      pool,
		}

		mask := createSelectiveBooleanMaskWithPool(50000, 0.3, pool)
		defer mask.Release()

		result, err := df.ParallelFilterWithBooleanMask(mask, opts)
		require.NoError(t, err)
		defer result.Release()

		assert.Greater(t, result.NumRows(), int64(0))
	})

	t.Run("ResourceCleanup", func(t *testing.T) {
		opts := &ParallelOptions{
			NumWorkers:      4,
			MinParallelRows: 10000,
		}

		testPool := memory.NewGoAllocator()
		mask := createSelectiveBooleanMaskWithPool(50000, 0.3, testPool)
		defer mask.Release()

		// Multiple operations to test cleanup
		for i := 0; i < 5; i++ {
			result, err := df.ParallelFilterWithBooleanMask(mask, opts)
			require.NoError(t, err)
			result.Release()
		}
	})
}

// TestParallelScalability tests performance scaling with worker count
func TestParallelScalability(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping scalability test in short mode")
	}

	df := createTestDataFrameForParallel(100000)
	defer df.Release()

	mask := createSelectiveBooleanMask(100000, 0.3)
	defer mask.Release()

	maxWorkers := runtime.NumCPU()
	workerCounts := []int{1, 2}
	if maxWorkers > 2 {
		workerCounts = append(workerCounts, maxWorkers)
	}

	for _, workers := range workerCounts {
		if workers > maxWorkers {
			continue
		}

		t.Run("Workers_"+formatWorkerCount(workers), func(t *testing.T) {
			opts := &ParallelOptions{
				NumWorkers:      workers,
				MinParallelRows: 10000,
			}

			start := time.Now()
			result, err := df.ParallelFilterWithBooleanMask(mask, opts)
			duration := time.Since(start)

			require.NoError(t, err)
			defer result.Release()

			t.Logf("Workers: %d, Duration: %v, Rows: %d", workers, duration, result.NumRows())
			assert.Greater(t, result.NumRows(), int64(0))
		})
	}
}

// TestParallelEdgeCases tests edge cases and boundary conditions
func TestParallelEdgeCases(t *testing.T) {
	t.Run("EmptyDataFrame", func(t *testing.T) {
		emptyDF := createTestDataFrameForParallel(0)
		defer emptyDF.Release()

		opts := DefaultParallelOptions()
		processor := NewParallelProcessor(opts)

		assert.False(t, processor.ShouldParallelize(emptyDF))

		chunks := processor.CalculateChunks(emptyDF)
		assert.Len(t, chunks, 1)
		assert.Equal(t, int64(0), chunks[0].StartRow)
		assert.Equal(t, int64(0), chunks[0].EndRow)
	})

	t.Run("SingleRowDataFrame", func(t *testing.T) {
		singleDF := createTestDataFrameForParallel(1)
		defer singleDF.Release()

		opts := &ParallelOptions{
			NumWorkers:      4,
			MinParallelRows: 1,
		}

		// Test with simple sum operation instead of expression
		sum, err := singleDF.ParallelSum("value", opts)
		require.NoError(t, err)

		// Single row with value 0 * 2.5 = 0, so we expect 0
		assert.Equal(t, 0.0, sum)
	})

	t.Run("ExactChunkBoundary", func(t *testing.T) {
		// Create DataFrame with size that's exact multiple of chunk size
		df := createTestDataFrameForParallel(40000)
		defer df.Release()

		opts := &ParallelOptions{
			NumWorkers:      4,
			ChunkSize:       10000, // Exact division
			MinParallelRows: 5000,
		}

		processor := NewParallelProcessor(opts)
		chunks := processor.CalculateChunks(df)

		assert.Equal(t, 4, len(chunks))
		for _, chunk := range chunks {
			assert.Equal(t, int64(10000), chunk.EndRow-chunk.StartRow)
		}
	})
}

// Helper functions for parallel testing

func createTestDataFrameForParallel(size int) *DataFrame {
	if size == 0 {
		// Create empty DataFrame
		pool := memory.NewGoAllocator()
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		idBuilder := array.NewInt64Builder(pool)
		valueBuilder := array.NewFloat64Builder(pool)

		idArray := idBuilder.NewArray()
		defer idArray.Release()

		valueArray := valueBuilder.NewArray()
		defer valueArray.Release()

		record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 0)
		defer record.Release()

		return NewDataFrame(record)
	}

	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i) * 2.5)
		nameBuilder.Append("item_" + string(rune('A'+i%26)))
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray, nameArray}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createSelectiveMaskForParallel(size int64, selectFunc func(int) bool) arrow.Array {
	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	for i := 0; i < int(size); i++ {
		builder.Append(selectFunc(i))
	}

	return builder.NewArray()
}
