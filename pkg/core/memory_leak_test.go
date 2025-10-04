package core

import (
	"runtime"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// TestMemoryLeak_DataFrameRelease verifies that Release() properly frees memory
func TestMemoryLeak_DataFrameRelease(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Get baseline memory
	runtime.GC()
	var baselineStats runtime.MemStats
	runtime.ReadMemStats(&baselineStats)
	baselineAlloc := baselineStats.Alloc

	// Create and release many DataFrames
	iterations := 1000
	for i := 0; i < iterations; i++ {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		// Create arrays
		idBuilder := array.NewInt64Builder(pool)
		valueBuilder := array.NewFloat64Builder(pool)

		for j := 0; j < 100; j++ {
			idBuilder.Append(int64(j))
			valueBuilder.Append(float64(j) * 1.5)
		}

		idArray := idBuilder.NewArray()
		valueArray := valueBuilder.NewArray()

		record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 100)
		df := NewDataFrame(record)

		// Release everything
		df.Release()
		idArray.Release()
		valueArray.Release()
		record.Release()
	}

	// Force GC and check memory
	runtime.GC()
	var afterStats runtime.MemStats
	runtime.ReadMemStats(&afterStats)
	afterAlloc := afterStats.Alloc

	// Memory should not grow significantly
	// Allow for some overhead (10KB per iteration is too much)
	maxGrowth := uint64(iterations * 10 * 1024)

	// Calculate growth, handling the case where memory decreased
	var actualGrowth uint64
	var decreased bool
	if afterAlloc >= baselineAlloc {
		actualGrowth = afterAlloc - baselineAlloc
	} else {
		// Memory decreased - this is good! No leak.
		actualGrowth = 0
		decreased = true
	}

	if actualGrowth > maxGrowth {
		t.Errorf("Potential memory leak detected: grew by %d bytes (max allowed: %d)",
			actualGrowth, maxGrowth)
	}

	if decreased {
		t.Logf("✅ Memory decreased: baseline %d bytes, after %d bytes (freed %d bytes)",
			baselineAlloc, afterAlloc, baselineAlloc-afterAlloc)
	} else {
		t.Logf("Memory baseline: %d bytes, after: %d bytes, growth: %d bytes",
			baselineAlloc, afterAlloc, actualGrowth)
	}
}

// TestMemoryLeak_SeriesRelease verifies Series Release() properly frees memory
func TestMemoryLeak_SeriesRelease(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Get baseline memory
	runtime.GC()
	var baselineStats runtime.MemStats
	runtime.ReadMemStats(&baselineStats)
	baselineAlloc := baselineStats.Alloc

	// Create and release many Series
	iterations := 1000
	for i := 0; i < iterations; i++ {
		builder := array.NewInt64Builder(pool)
		for j := 0; j < 100; j++ {
			builder.Append(int64(j))
		}

		arr := builder.NewArray()
		field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(arr, field)

		// Release everything
		series.Release()
		arr.Release()
	}

	// Force GC and check memory
	runtime.GC()
	var afterStats runtime.MemStats
	runtime.ReadMemStats(&afterStats)
	afterAlloc := afterStats.Alloc

	// Memory should not grow significantly
	maxGrowth := uint64(iterations * 10 * 1024)

	// Calculate growth, handling the case where memory decreased
	var actualGrowth uint64
	var decreased bool
	if afterAlloc >= baselineAlloc {
		actualGrowth = afterAlloc - baselineAlloc
	} else {
		// Memory decreased - this is good! No leak.
		actualGrowth = 0
		decreased = true
	}

	if actualGrowth > maxGrowth {
		t.Errorf("Potential memory leak detected: grew by %d bytes (max allowed: %d)",
			actualGrowth, maxGrowth)
	}

	if decreased {
		t.Logf("✅ Memory decreased: baseline %d bytes, after %d bytes (freed %d bytes)",
			baselineAlloc, afterAlloc, baselineAlloc-afterAlloc)
	} else {
		t.Logf("Memory baseline: %d bytes, after: %d bytes, growth: %d bytes",
			baselineAlloc, afterAlloc, actualGrowth)
	}
}

// TestMemoryLeak_Operations verifies operations don't leak memory
func TestMemoryLeak_Operations(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	for j := 0; j < 1000; j++ {
		idBuilder.Append(int64(j))
		valueBuilder.Append(float64(j) * 1.5)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 1000)
	df := NewDataFrame(record)
	defer df.Release()

	// Get baseline memory
	runtime.GC()
	var baselineStats runtime.MemStats
	runtime.ReadMemStats(&baselineStats)
	baselineAlloc := baselineStats.Alloc

	// Perform operations and release results
	iterations := 500
	for i := 0; i < iterations; i++ {
		// Select
		selected, _ := df.Select([]string{"id"})
		selected.Release()

		// Column access
		col, _ := df.Column("value")
		col.Release()

		// ColumnAt
		colAt, _ := df.ColumnAt(0)
		colAt.Release()
	}

	// Force GC and check memory
	runtime.GC()
	var afterStats runtime.MemStats
	runtime.ReadMemStats(&afterStats)
	afterAlloc := afterStats.Alloc

	// Memory should not grow significantly
	maxGrowth := uint64(iterations * 5 * 1024)

	// Calculate growth, handling the case where memory decreased
	var actualGrowth uint64
	var decreased bool
	if afterAlloc >= baselineAlloc {
		actualGrowth = afterAlloc - baselineAlloc
	} else {
		// Memory decreased - this is good! No leak.
		actualGrowth = 0
		decreased = true
	}

	if actualGrowth > maxGrowth {
		t.Errorf("Potential memory leak in operations: grew by %d bytes (max allowed: %d)",
			actualGrowth, maxGrowth)
	}

	if decreased {
		t.Logf("✅ Memory decreased: baseline %d bytes, after %d bytes (freed %d bytes)",
			baselineAlloc, afterAlloc, baselineAlloc-afterAlloc)
	} else {
		t.Logf("Memory baseline: %d bytes, after: %d bytes, growth: %d bytes",
			baselineAlloc, afterAlloc, actualGrowth)
	}
}
