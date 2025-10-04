package core

import (
	"strconv"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// Benchmark Comparison: GopherFrame vs Gota
//
// These benchmarks compare GopherFrame performance against Gota for common operations.
// Tests use identical datasets and operations to ensure fair comparison.

const (
	benchSize1K   = 1000
	benchSize10K  = 10000
	benchSize100K = 100000
)

// ========== DataFrame Creation ==========

func BenchmarkCreate_GopherFrame_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df := createBenchDataFrameGopherFrame(pool, benchSize1K)
		df.Release()
	}
}

func BenchmarkCreate_Gota_1K(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = createBenchDataFrameGota(benchSize1K)
	}
}

func BenchmarkCreate_GopherFrame_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		df := createBenchDataFrameGopherFrame(pool, benchSize10K)
		df.Release()
	}
}

func BenchmarkCreate_Gota_10K(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = createBenchDataFrameGota(benchSize10K)
	}
}

// ========== Filter Operations ==========

func BenchmarkFilter_GopherFrame_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createBenchDataFrameGopherFrame(pool, benchSize1K)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Filter: value > 500
		predicateBuilder := array.NewBooleanBuilder(pool)
		valueCol := df.record.Column(1).(*array.Int64)
		for j := 0; j < int(df.record.NumRows()); j++ {
			predicateBuilder.Append(valueCol.Value(j) > 500)
		}
		predicateArray := predicateBuilder.NewArray()

		filtered, _ := df.Filter(predicateArray)
		if filtered != nil {
			filtered.Release()
		}
		predicateArray.Release()
		predicateBuilder.Release()
	}
}

func BenchmarkFilter_Gota_1K(b *testing.B) {
	df := createBenchDataFrameGota(benchSize1K)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Filter: value > 500
		_ = df.Filter(
			dataframe.F{
				Colname:    "value",
				Comparator: series.Greater,
				Comparando: 500,
			},
		)
	}
}

func BenchmarkFilter_GopherFrame_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createBenchDataFrameGopherFrame(pool, benchSize10K)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		predicateBuilder := array.NewBooleanBuilder(pool)
		valueCol := df.record.Column(1).(*array.Int64)
		for j := 0; j < int(df.record.NumRows()); j++ {
			predicateBuilder.Append(valueCol.Value(j) > 500)
		}
		predicateArray := predicateBuilder.NewArray()

		filtered, _ := df.Filter(predicateArray)
		if filtered != nil {
			filtered.Release()
		}
		predicateArray.Release()
		predicateBuilder.Release()
	}
}

func BenchmarkFilter_Gota_10K(b *testing.B) {
	df := createBenchDataFrameGota(benchSize10K)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = df.Filter(
			dataframe.F{
				Colname:    "value",
				Comparator: series.Greater,
				Comparando: 500,
			},
		)
	}
}

// ========== Select Operations ==========

func BenchmarkSelect_GopherFrame_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createBenchDataFrameGopherFrame(pool, benchSize1K)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selected, _ := df.Select([]string{"id", "value"})
		if selected != nil {
			selected.Release()
		}
	}
}

func BenchmarkSelect_Gota_1K(b *testing.B) {
	df := createBenchDataFrameGota(benchSize1K)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = df.Select([]string{"id", "value"})
	}
}

func BenchmarkSelect_GopherFrame_10K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createBenchDataFrameGopherFrame(pool, benchSize10K)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selected, _ := df.Select([]string{"id", "value"})
		if selected != nil {
			selected.Release()
		}
	}
}

func BenchmarkSelect_Gota_10K(b *testing.B) {
	df := createBenchDataFrameGota(benchSize10K)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = df.Select([]string{"id", "value"})
	}
}

// ========== Helper Functions ==========

func createBenchDataFrameGopherFrame(pool memory.Allocator, size int) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	defer idBuilder.Release()
	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()

	categories := []string{"A", "B", "C", "D"}
	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(int64(i * 10))
		categoryBuilder.Append(categories[i%4])
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, valueArray, categoryArray},
		int64(size),
	)

	return NewDataFrame(record)
}

func createBenchDataFrameGota(size int) dataframe.DataFrame {
	ids := make([]int, size)
	values := make([]int, size)
	categories := make([]string, size)

	cats := []string{"A", "B", "C", "D"}
	for i := 0; i < size; i++ {
		ids[i] = i
		values[i] = i * 10
		categories[i] = cats[i%4]
	}

	return dataframe.New(
		series.New(ids, series.Int, "id"),
		series.New(values, series.Int, "value"),
		series.New(categories, series.String, "category"),
	)
}

// ========== Column Access ==========

func BenchmarkColumnAccess_GopherFrame_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createBenchDataFrameGopherFrame(pool, benchSize1K)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		col, _ := df.Column("value")
		if col != nil {
			col.Release()
		}
	}
}

func BenchmarkColumnAccess_Gota_1K(b *testing.B) {
	df := createBenchDataFrameGota(benchSize1K)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = df.Col("value")
	}
}

// ========== Iteration ==========

func BenchmarkIteration_GopherFrame_1K(b *testing.B) {
	pool := memory.NewGoAllocator()
	df := createBenchDataFrameGopherFrame(pool, benchSize1K)
	defer df.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		valueCol := df.record.Column(1).(*array.Int64)
		sum := int64(0)
		for j := 0; j < valueCol.Len(); j++ {
			sum += valueCol.Value(j)
		}
		_ = sum
	}
}

func BenchmarkIteration_Gota_1K(b *testing.B) {
	df := createBenchDataFrameGota(benchSize1K)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		records := df.Records()
		sum := 0
		for j := 1; j < len(records); j++ { // Skip header
			val, _ := strconv.Atoi(records[j][1])
			sum += val
		}
		_ = sum
	}
}
