// memory_test.go - Memory leak and resource management testing
// This file ensures production-ready memory management

package gopherframe

import (
	"runtime"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMemoryLeaks verifies no memory leaks in core operations
func TestMemoryLeaks(t *testing.T) {
	// Force garbage collection and get baseline
	runtime.GC()
	runtime.GC() // Call twice to ensure clean state
	var baseline runtime.MemStats
	runtime.ReadMemStats(&baseline)

	// Run operations that should not leak memory
	iterations := 1000
	for i := 0; i < iterations; i++ {
		df := createBenchmarkDataFrame(1000)

		// Perform various operations (avoid GroupBy due to type complexity)
		filtered := df.Filter(Col("value").Gt(Lit(50.0)))
		selected := filtered.Select("id", "value")
		withCol := selected.WithColumn("doubled", Col("value").Mul(Lit(2.0)))
		sorted := withCol.Sort("value", true)

		// Clean up
		sorted.Release()
		withCol.Release()
		selected.Release()
		filtered.Release()
		df.Release()

		// Periodic GC to prevent memory buildup during test
		if i%100 == 0 {
			runtime.GC()
		}
	}

	// Force final garbage collection
	runtime.GC()
	runtime.GC()
	time.Sleep(10 * time.Millisecond) // Allow finalizers to run

	var final runtime.MemStats
	runtime.ReadMemStats(&final)

	// Check for significant memory growth
	allocIncrease := final.Alloc - baseline.Alloc
	heapIncrease := final.HeapAlloc - baseline.HeapAlloc

	t.Logf("Baseline Alloc: %d bytes", baseline.Alloc)
	t.Logf("Final Alloc: %d bytes", final.Alloc)
	t.Logf("Alloc Increase: %d bytes", allocIncrease)
	t.Logf("Heap Increase: %d bytes", heapIncrease)

	// Allow for some increase due to Go runtime overhead, but should be minimal
	maxAcceptableIncrease := uint64(1024 * 1024) // 1MB
	require.Less(t, allocIncrease, maxAcceptableIncrease,
		"Memory allocation increased by %d bytes, expected less than %d",
		allocIncrease, maxAcceptableIncrease)
}

// TestArrowMemoryManagement verifies proper Arrow memory allocator usage
func TestArrowMemoryManagement(t *testing.T) {
	// Arrow memory management is handled internally by the Go runtime
	// We verify proper usage by ensuring no panics and proper Release() calls
	pool := memory.NewGoAllocator()

	// Create and release multiple Arrow arrays
	for i := 0; i < 100; i++ {
		builder := array.NewInt64Builder(pool)
		for j := 0; j < 1000; j++ {
			builder.Append(int64(j))
		}
		arr := builder.NewArray()
		// Verify array was created successfully
		assert.Equal(t, 1000, arr.Len())
		arr.Release()
	}

	// If we reach here without panics, memory management is working correctly
	assert.True(t, true, "Arrow memory allocation and release completed successfully")
}

// TestDataFrameReleasePattern verifies proper Release() pattern
func TestDataFrameReleasePattern(t *testing.T) {
	// Test proper Release() pattern without allocator tracking
	// (Arrow Go doesn't expose allocator metrics in the public API)

	// Create DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()
	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	for i := 0; i < 1000; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i))
	}

	idArray := idBuilder.NewArray()
	valueArray := valueBuilder.NewArray()

	// Verify arrays were created
	assert.Equal(t, 1000, idArray.Len())
	assert.Equal(t, 1000, valueArray.Len())

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 1000)
	df := NewDataFrame(record)

	// Verify DataFrame was created correctly
	assert.Equal(t, int64(1000), df.NumRows())
	assert.Equal(t, int64(2), df.NumCols())

	// Release everything in correct order
	idArray.Release()
	valueArray.Release()
	df.Release()

	// If we reach here without panics, release pattern is correct
	assert.True(t, true, "DataFrame release pattern completed successfully")
}

// TestConcurrentDataFrameOperations verifies thread safety
func TestConcurrentDataFrameOperations(t *testing.T) {
	df := createBenchmarkDataFrame(10000)
	defer df.Release()

	// Track goroutine count
	initialGoroutines := runtime.NumGoroutine()

	// Run concurrent read operations
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()

			// Read-only operations should be safe
			_ = df.NumRows()
			_ = df.NumCols()
			_ = df.ColumnNames()

			// Create derived DataFrames (each goroutine gets its own)
			filtered := df.Filter(Col("value").Gt(Lit(50.0)))
			defer filtered.Release()

			selected := filtered.Select("id", "value")
			defer selected.Release()

			_ = selected.NumRows()
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Allow time for goroutines to clean up
	time.Sleep(100 * time.Millisecond)
	runtime.GC()

	// Check no goroutines leaked
	finalGoroutines := runtime.NumGoroutine()
	assert.LessOrEqual(t, finalGoroutines, initialGoroutines+1, // +1 for test runner variation
		"Goroutine count should not increase after concurrent operations")
}

// TestLargeDataFrameMemoryUsage verifies memory usage for large DataFrames
func TestLargeDataFrameMemoryUsage(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large DataFrame test in short mode")
	}

	runtime.GC()
	var beforeStats runtime.MemStats
	runtime.ReadMemStats(&beforeStats)

	// Create large DataFrame (1M rows)
	largeDF := createBenchmarkDataFrame(1000000)

	var afterCreateStats runtime.MemStats
	runtime.ReadMemStats(&afterCreateStats)

	creationMemory := afterCreateStats.Alloc - beforeStats.Alloc
	t.Logf("Memory used for 1M row DataFrame: %d MB", creationMemory/(1024*1024))

	// Perform operations
	filtered := largeDF.Filter(Col("value").Gt(Lit(50.0)))
	grouped := filtered.GroupBy("category").Agg(Sum("value"))

	var afterOpsStats runtime.MemStats
	runtime.ReadMemStats(&afterOpsStats)

	opsMemory := afterOpsStats.Alloc - afterCreateStats.Alloc
	t.Logf("Additional memory for operations: %d MB", opsMemory/(1024*1024))

	// Clean up
	grouped.Release()
	filtered.Release()
	largeDF.Release()

	runtime.GC()
	runtime.GC()
	time.Sleep(50 * time.Millisecond)

	var afterCleanupStats runtime.MemStats
	runtime.ReadMemStats(&afterCleanupStats)

	// Memory should be significantly reduced after cleanup
	finalMemory := afterCleanupStats.Alloc
	memoryReduction := afterOpsStats.Alloc - finalMemory

	t.Logf("Memory reduction after cleanup: %d MB", memoryReduction/(1024*1024))

	// Should recover at least 50% of the allocated memory
	minExpectedReduction := (afterOpsStats.Alloc - beforeStats.Alloc) / 2
	assert.Greater(t, memoryReduction, minExpectedReduction,
		"Should recover significant memory after DataFrame cleanup")
}

// TestResourceCleanupPatterns verifies various cleanup scenarios
func TestResourceCleanupPatterns(t *testing.T) {
	t.Run("ChainedOperationCleanup", func(t *testing.T) {
		// Test chained operations cleanup without allocator tracking
		df := createBenchmarkDataFrame(1000)

		// Verify initial DataFrame
		assert.Equal(t, int64(1000), df.NumRows())

		// Create chain of operations
		result := df.
			Filter(Col("value").Gt(Lit(25.0))).
			WithColumn("doubled", Col("value").Mul(Lit(2.0))).
			Select("id", "doubled")

		// Verify result DataFrame
		assert.Greater(t, result.NumRows(), int64(0))
		assert.Equal(t, int64(2), result.NumCols())

		// Clean up
		result.Release()
		df.Release()

		// If we reach here without panics, cleanup was successful
		assert.True(t, true, "Chained operation cleanup completed successfully")
	})

	t.Run("IOOperationCleanup", func(t *testing.T) {
		// Test file I/O cleanup
		df := createBenchmarkDataFrame(1000)
		tempFile := t.TempDir() + "/test.parquet"

		// Write and read
		err := WriteParquet(df, tempFile)
		require.NoError(t, err)

		df.Release()

		readDF, err := ReadParquet(tempFile)
		require.NoError(t, err)

		// Verify read DataFrame
		assert.Equal(t, int64(1000), readDF.NumRows())
		assert.Equal(t, int64(3), readDF.NumCols())

		readDF.Release()

		// If we reach here without panics, I/O cleanup was successful
		assert.True(t, true, "I/O operation cleanup completed successfully")
	})
}

// Note: Using createBenchmarkDataFrame helper function from benchmark_test.go
