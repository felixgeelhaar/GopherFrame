package core

import (
	"runtime"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// BenchmarkParallelVsSequential compares parallel and sequential processing performance
func BenchmarkParallelVsSequential(b *testing.B) {
	sizes := []int{50000, 200000, 1000000}

	for _, size := range sizes {
		df := createLargeDataFrame(size)
		defer df.Release()

		mask := createSelectiveBooleanMask(size, 0.3)
		defer mask.Release()

		sizeStr := formatLargeSize(size)

		// Filter benchmarks
		b.Run("Filter_Sequential_"+sizeStr, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterOptimized(mask)
				if err != nil {
					b.Fatalf("Sequential filter failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Filter_Parallel_"+sizeStr, func(b *testing.B) {
			opts := &ParallelOptions{
				NumWorkers:      runtime.NumCPU(),
				MinParallelRows: 10000,
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.ParallelFilterWithBooleanMask(mask, opts)
				if err != nil {
					b.Fatalf("Parallel filter failed: %v", err)
				}
				result.Release()
			}
		})

		// Note: Vectorized filtering with expressions requires expr package integration
		// This benchmark will be added when expression integration is completed
	}
}

// BenchmarkParallelExpressionEvaluation tests parallel expression performance
func BenchmarkParallelExpressionEvaluation(b *testing.B) {
	// Note: Expression evaluation benchmarks require expr package integration
	// These will be implemented when expression integration is completed
	b.Skip("Expression evaluation benchmarks require expr package integration")
}

// BenchmarkParallelAggregation tests parallel aggregation performance
func BenchmarkParallelAggregation(b *testing.B) {
	sizes := []int{100000, 1000000, 5000000}

	for _, size := range sizes {
		df := createLargeDataFrame(size)
		defer df.Release()

		sizeStr := formatLargeSize(size)

		b.Run("Sum_Sequential_"+sizeStr, func(b *testing.B) {
			seqOpts := &ParallelOptions{
				MinParallelRows: int64(size + 1), // Force sequential
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := df.ParallelSum("value", seqOpts)
				if err != nil {
					b.Fatalf("Sequential sum failed: %v", err)
				}
			}
		})

		b.Run("Sum_Parallel_"+sizeStr, func(b *testing.B) {
			opts := &ParallelOptions{
				NumWorkers:      runtime.NumCPU(),
				MinParallelRows: 25000,
				Strategy:        StrategyAuto,
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := df.ParallelSum("value", opts)
				if err != nil {
					b.Fatalf("Parallel sum failed: %v", err)
				}
			}
		})
	}
}

// BenchmarkParallelChunkingStrategies tests different chunking strategies
func BenchmarkParallelChunkingStrategies(b *testing.B) {
	size := 500000
	df := createLargeDataFrame(size)
	defer df.Release()

	mask := createSelectiveBooleanMask(size, 0.5)
	defer mask.Release()

	strategies := []struct {
		name     string
		strategy ParallelStrategy
	}{
		{"Auto", StrategyAuto},
		{"FixedChunks", StrategyFixedChunks},
		{"CacheFriendly", StrategyCacheFriendly},
		{"MemoryBound", StrategyMemoryBound},
	}

	for _, strat := range strategies {
		b.Run("Strategy_"+strat.name, func(b *testing.B) {
			opts := &ParallelOptions{
				NumWorkers:      runtime.NumCPU(),
				MinParallelRows: 10000,
				Strategy:        strat.strategy,
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.ParallelFilterWithBooleanMask(mask, opts)
				if err != nil {
					b.Fatalf("Filter with strategy %s failed: %v", strat.name, err)
				}
				result.Release()
			}
		})
	}
}

// BenchmarkParallelWorkerCounts tests performance with different worker counts
func BenchmarkParallelWorkerCounts(b *testing.B) {
	size := 1000000
	df := createLargeDataFrame(size)
	defer df.Release()

	mask := createSelectiveBooleanMask(size, 0.3)
	defer mask.Release()

	maxWorkers := runtime.NumCPU()
	workerCounts := []int{1, 2, 4}
	if maxWorkers > 4 {
		workerCounts = append(workerCounts, maxWorkers/2, maxWorkers)
	}

	for _, workers := range workerCounts {
		if workers > maxWorkers {
			continue
		}

		b.Run("Workers_"+formatWorkerCount(workers), func(b *testing.B) {
			opts := &ParallelOptions{
				NumWorkers:      workers,
				MinParallelRows: 50000,
				Strategy:        StrategyAuto,
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.ParallelFilterWithBooleanMask(mask, opts)
				if err != nil {
					b.Fatalf("Filter with %d workers failed: %v", workers, err)
				}
				result.Release()
			}
		})
	}
}

// BenchmarkParallelChainedOperations tests chained parallel operations
func BenchmarkParallelChainedOperations(b *testing.B) {
	// Note: Chained operations benchmarks require expr package integration
	b.Skip("Chained operations benchmarks require expr package integration")
}

// BenchmarkParallelMemoryUsage tests memory efficiency of parallel operations
func BenchmarkParallelMemoryUsage(b *testing.B) {
	size := 500000
	df := createLargeDataFrame(size)
	defer df.Release()

	// Test with shared memory pool
	b.Run("SharedMemoryPool", func(b *testing.B) {
		pool := memory.NewGoAllocator()
		opts := &ParallelOptions{
			NumWorkers:      runtime.NumCPU(),
			MinParallelRows: 50000,
			MemoryPool:      pool,
		}

		mask := createSelectiveBooleanMaskWithPool(size, 0.4, pool)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.ParallelFilterWithBooleanMask(mask, opts)
			if err != nil {
				b.Fatalf("Parallel filter with shared pool failed: %v", err)
			}
			result.Release()
		}
	})

	// Test with separate pools
	b.Run("SeparatePools", func(b *testing.B) {
		opts := &ParallelOptions{
			NumWorkers:      runtime.NumCPU(),
			MinParallelRows: 50000,
			MemoryPool:      nil, // Will create separate pools
		}

		mask := createSelectiveBooleanMask(size, 0.4)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.ParallelFilterWithBooleanMask(mask, opts)
			if err != nil {
				b.Fatalf("Parallel filter with separate pools failed: %v", err)
			}
			result.Release()
		}
	})
}

// Helper functions for large dataset creation and formatting

func createLargeDataFrame(size int) *DataFrame {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "flag", Type: arrow.FixedWidthTypes.Boolean},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	categoryBuilder := array.NewStringBuilder(pool)
	flagBuilder := array.NewBooleanBuilder(pool)

	// Reserve capacity for better performance
	idBuilder.Reserve(size)
	valueBuilder.Reserve(size)
	categoryBuilder.Reserve(size)
	flagBuilder.Reserve(size)

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i%10000) * 1.337)
		categoryBuilder.Append("cat_" + string(rune('A'+i%26)))
		flagBuilder.Append(i%3 == 0)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	flagArray := flagBuilder.NewArray()
	defer flagArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray, categoryArray, flagArray}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createSelectiveBooleanMask(size int, selectivity float64) arrow.Array {
	return createSelectiveBooleanMaskWithPool(size, selectivity, memory.NewGoAllocator())
}

func createSelectiveBooleanMaskWithPool(size int, selectivity float64, pool memory.Allocator) arrow.Array {
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	builder.Reserve(size)
	threshold := int(float64(size) * selectivity)

	for i := 0; i < size; i++ {
		// Create a pattern that gives approximately the desired selectivity
		builder.Append(i%int(1.0/selectivity) < threshold/int(1.0/selectivity))
	}

	return builder.NewArray()
}

func formatLargeSize(size int) string {
	if size >= 1000000 {
		return string(rune('0'+size/1000000)) + "M"
	} else if size >= 1000 {
		return string(rune('0'+size/1000)) + "K"
	}
	return string(rune('0'+size/100)) + "H"
}

func formatWorkerCount(workers int) string {
	if workers < 10 {
		return string(rune('0' + workers))
	}
	return string(rune('0'+workers/10)) + string(rune('0'+workers%10))
}
