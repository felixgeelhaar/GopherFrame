package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// Phase 2 Feature Benchmarks
// Benchmarks for joins, window functions, temporal operations, string operations, and statistical aggregations

const (
	smallSize  = 1000
	mediumSize = 10000
	largeSize  = 100000
	partitions = 10 // Number of partitions for window functions
	windowSize = 7  // Window size for rolling aggregations
)

// ========== Join Operation Benchmarks ==========

func BenchmarkInnerJoin_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	left := createJoinDataFrame(pool, smallSize, "left")
	defer left.Release()
	right := createJoinDataFrame(pool, smallSize/2, "right")
	defer right.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := left.InnerJoin(right, "key", "key")
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkInnerJoin_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	left := createJoinDataFrame(pool, mediumSize, "left")
	defer left.Release()
	right := createJoinDataFrame(pool, mediumSize/2, "right")
	defer right.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := left.InnerJoin(right, "key", "key")
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkLeftJoin_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	left := createJoinDataFrame(pool, smallSize, "left")
	defer left.Release()
	right := createJoinDataFrame(pool, smallSize/2, "right")
	defer right.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := left.LeftJoin(right, "key", "key")
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkLeftJoin_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	left := createJoinDataFrame(pool, mediumSize, "left")
	defer left.Release()
	right := createJoinDataFrame(pool, mediumSize/2, "right")
	defer right.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := left.LeftJoin(right, "key", "key")
		if result != nil {
			result.Release()
		}
	}
}

// ========== Window Function Benchmarks ==========

func BenchmarkRowNumber_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createPartitionedDataFrame(pool, smallSize, partitions)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().PartitionBy("partition").OrderBy("value").Over(RowNumber().As("row_num"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkRowNumber_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createPartitionedDataFrame(pool, mediumSize, partitions)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().PartitionBy("partition").OrderBy("value").Over(RowNumber().As("row_num"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkRank_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createPartitionedDataFrame(pool, smallSize, partitions)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().PartitionBy("partition").OrderBy("value").Over(Rank().As("rank"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkRank_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createPartitionedDataFrame(pool, mediumSize, partitions)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().PartitionBy("partition").OrderBy("value").Over(Rank().As("rank"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkLag_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createPartitionedDataFrame(pool, smallSize, partitions)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().PartitionBy("partition").OrderBy("value").Over(Lag("value", 1).As("prev_value"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkLag_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createPartitionedDataFrame(pool, mediumSize, partitions)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().PartitionBy("partition").OrderBy("value").Over(Lag("value", 1).As("prev_value"))
		if result != nil {
			result.Release()
		}
	}
}

// ========== Rolling Aggregation Benchmarks ==========

func BenchmarkRollingSum_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, smallSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Rows(windowSize).Over(RollingSum("value").As("rolling_sum"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkRollingSum_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, mediumSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Rows(windowSize).Over(RollingSum("value").As("rolling_sum"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkRollingMean_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, smallSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Rows(windowSize).Over(RollingMean("value").As("rolling_mean"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkRollingMean_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, mediumSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Rows(windowSize).Over(RollingMean("value").As("rolling_mean"))
		if result != nil {
			result.Release()
		}
	}
}

// ========== Cumulative Aggregation Benchmarks ==========

func BenchmarkCumSum_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, smallSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Over(CumSum("value").As("cumsum"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkCumSum_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, mediumSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Over(CumSum("value").As("cumsum"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkCumMax_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, smallSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Over(CumMax("value").As("cummax"))
		if result != nil {
			result.Release()
		}
	}
}

func BenchmarkCumMax_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createTimeSeriesDataFrame(pool, mediumSize)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := df.Window().Over(CumMax("value").As("cummax"))
		if result != nil {
			result.Release()
		}
	}
}

// ========== Helper Functions ==========

func createJoinDataFrame(pool memory.Allocator, size int, prefix string) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "key", Type: arrow.PrimitiveTypes.Int64},
			{Name: prefix + "_value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	keyBuilder := array.NewInt64Builder(pool)
	defer keyBuilder.Release()
	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()

	for i := 0; i < size; i++ {
		keyBuilder.Append(int64(i % (size / 2))) // Create some duplicates
		valueBuilder.Append(float64(i * 10))
	}

	keyArray := keyBuilder.NewArray()
	defer keyArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(
		schema,
		[]arrow.Array{keyArray, valueArray},
		int64(size),
	)

	return NewDataFrame(record)
}

func createPartitionedDataFrame(pool memory.Allocator, size int, numPartitions int) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "partition", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	defer idBuilder.Release()
	partitionBuilder := array.NewStringBuilder(pool)
	defer partitionBuilder.Release()
	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		partitionBuilder.Append(string(rune('A' + (i % numPartitions))))
		valueBuilder.Append(float64(i % 100))
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()
	partitionArray := partitionBuilder.NewArray()
	defer partitionArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, partitionArray, valueArray},
		int64(size),
	)

	return NewDataFrame(record)
}

func createTimeSeriesDataFrame(pool memory.Allocator, size int) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "date", Type: arrow.PrimitiveTypes.Int64}, // Simplified as Int64
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	dateBuilder := array.NewInt64Builder(pool)
	defer dateBuilder.Release()
	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()

	for i := 0; i < size; i++ {
		dateBuilder.Append(int64(i))
		valueBuilder.Append(float64((i % 50) + 10)) // Values between 10-60
	}

	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(
		schema,
		[]arrow.Array{dateArray, valueArray},
		int64(size),
	)

	return NewDataFrame(record)
}

// ========== Memory Allocation Benchmarks ==========

func BenchmarkMemoryAllocation_CreateRelease_1K(b *testing.B) {
	pool := memory.NewGoAllocator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df := createBenchDataFrameGopherFrame(pool, smallSize)
		df.Release()
	}
}

func BenchmarkMemoryAllocation_CreateRelease_10K(b *testing.B) {
	pool := memory.NewGoAllocator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df := createBenchDataFrameGopherFrame(pool, mediumSize)
		df.Release()
	}
}

func BenchmarkMemoryAllocation_Join_1K(b *testing.B) {
	pool := memory.NewGoAllocator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		left := createJoinDataFrame(pool, smallSize, "left")
		right := createJoinDataFrame(pool, smallSize/2, "right")
		result, _ := left.InnerJoin(right, "key", "key")
		if result != nil {
			result.Release()
		}
		left.Release()
		right.Release()
	}
}
