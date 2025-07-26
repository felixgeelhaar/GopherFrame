package expr

import (
	"runtime"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMemoryPoolUsage tests that expressions use shared memory pools efficiently
func TestMemoryPoolUsage(t *testing.T) {
	t.Run("PooledVsNonPooled", func(t *testing.T) {
		df := createTestDataFrameWithPool(memory.NewGoAllocator())
		defer df.Release()

		// Test with shared pool
		sharedPool := memory.NewGoAllocator()
		expr := Col("value").Add(Lit(10.0))

		result1, err := expr.EvaluateWithPool(df, sharedPool)
		require.NoError(t, err)
		defer result1.Release()

		result2, err := expr.EvaluateWithPool(df, sharedPool)
		require.NoError(t, err)
		defer result2.Release()

		// Verify results are correct
		assert.Equal(t, df.NumRows(), int64(result1.Len()))
		assert.Equal(t, df.NumRows(), int64(result2.Len()))

		// Results should be equal
		assert.Equal(t, result1.DataType(), result2.DataType())
	})

	t.Run("ChainedOperationsMemoryEfficiency", func(t *testing.T) {
		df := createTestDataFrameWithPool(memory.NewGoAllocator())
		defer df.Release()

		pool := memory.NewGoAllocator()

		// Create a complex chained expression
		expr := Col("value").
			Add(Lit(10.0)).
			Mul(Lit(2.0)).
			Sub(Lit(5.0))

		result, err := expr.EvaluateWithPool(df, pool)
		require.NoError(t, err)
		defer result.Release()

		// Verify result correctness
		assert.Equal(t, df.NumRows(), int64(result.Len()))
		assert.Equal(t, arrow.FLOAT64, result.DataType().ID())
	})
}

// BenchmarkMemoryPoolPerformance compares pooled vs non-pooled allocations
func BenchmarkMemoryPoolPerformance(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("WithPool_Size_"+string(rune(size/1000))+"K", func(b *testing.B) {
			pool := memory.NewGoAllocator()
			df := createBenchmarkDataFrameWithPool(size, pool)
			defer df.Release()

			expr := Col("value").Add(Lit(10.0)).Mul(Lit(2.0))

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := expr.EvaluateWithPool(df, pool)
				if err != nil {
					b.Fatalf("Evaluation failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("WithoutPool_Size_"+string(rune(size/1000))+"K", func(b *testing.B) {
			df := createBenchmarkDataFrameWithPool(size, memory.NewGoAllocator())
			defer df.Release()

			expr := Col("value").Add(Lit(10.0)).Mul(Lit(2.0))

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := expr.Evaluate(df)
				if err != nil {
					b.Fatalf("Evaluation failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

// TestMemoryLeaks ensures expressions don't leak memory
func TestMemoryLeaks(t *testing.T) {
	// Get initial memory stats
	var m1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// Perform many operations
	for i := 0; i < 100; i++ {
		df := createTestDataFrameForMemory(1000)
		expr := Col("value").Add(Lit(float64(i))).Mul(Lit(2.0))

		result, err := expr.Evaluate(df)
		require.NoError(t, err)

		// Immediately release
		result.Release()
		df.Release()
	}

	// Force GC and get final memory stats
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Memory increase should be minimal (less than 10MB)
	memIncrease := int64(m2.Alloc) - int64(m1.Alloc)
	assert.Less(t, memIncrease, int64(10*1024*1024), "Memory leak detected: %d bytes increase", memIncrease)
}

// AllocationTracker wraps an allocator to track memory usage
type AllocationTracker struct {
	memory.Allocator
	allocations int
	totalBytes  int64
}

func (a *AllocationTracker) Allocate(size int) []byte {
	a.allocations++
	a.totalBytes += int64(size)
	return a.Allocator.Allocate(size)
}

func (a *AllocationTracker) Reset() {
	a.allocations = 0
	a.totalBytes = 0
}

func (a *AllocationTracker) AllocationCount() int {
	return a.allocations
}

func (a *AllocationTracker) TotalAllocated() int64 {
	return a.totalBytes
}

// Helper functions
func createTestDataFrameWithPool(pool memory.Allocator) *core.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	for i := 0; i < 1000; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i) * 1.5)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 1000)
	defer record.Release()

	return core.NewDataFrame(record)
}

func createBenchmarkDataFrameWithPool(size int, pool memory.Allocator) *core.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i%100) + float64(i%10)*0.1)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, int64(size))
	defer record.Release()

	return core.NewDataFrame(record)
}

func createTestDataFrameForMemory(size int) *core.DataFrame {
	return createBenchmarkDataFrameWithPool(size, memory.NewGoAllocator())
}
