package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// BenchmarkCoreOperations benchmarks the hot path operations
func BenchmarkCoreOperations(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		// Filter benchmarks
		b.Run("Filter_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			// Create boolean mask - select ~50% of rows
			mask := createBooleanMask(size, 0.5)
			defer mask.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Filter(mask)
				if err != nil {
					b.Fatalf("Filter failed: %v", err)
				}
				result.Release()
			}
		})

		// Select benchmarks
		b.Run("Select_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			columns := []string{"id", "value"}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Select(columns)
				if err != nil {
					b.Fatalf("Select failed: %v", err)
				}
				result.Release()
			}
		})

		// WithColumn benchmarks
		b.Run("WithColumn_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			// Create new column with computed values
			newColumn := createComputedColumn(size)
			defer newColumn.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.WithColumn("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumn failed: %v", err)
				}
				result.Release()
			}
		})

		// Combined operations benchmark
		b.Run("CombinedOps_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			mask := createBooleanMask(size, 0.3)
			defer mask.Release()

			newColumn := createComputedColumn(size)
			defer newColumn.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Chain operations: WithColumn -> Filter -> Select
				withCol, err := df.WithColumn("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumn failed: %v", err)
				}

				filtered, err := withCol.Filter(mask)
				withCol.Release()
				if err != nil {
					b.Fatalf("Filter failed: %v", err)
				}

				selected, err := filtered.Select([]string{"id", "computed"})
				filtered.Release()
				if err != nil {
					b.Fatalf("Select failed: %v", err)
				}

				selected.Release()
			}
		})
	}
}

// BenchmarkFilterTypes benchmarks filtering on different data types
func BenchmarkFilterTypes(b *testing.B) {
	size := 50000

	b.Run("FilterInt64", func(b *testing.B) {
		df := createInt64DataFrame(size)
		defer df.Release()

		mask := createBooleanMask(size, 0.3)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.Filter(mask)
			if err != nil {
				b.Fatalf("Filter failed: %v", err)
			}
			result.Release()
		}
	})

	b.Run("FilterFloat64", func(b *testing.B) {
		df := createFloat64DataFrame(size)
		defer df.Release()

		mask := createBooleanMask(size, 0.3)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.Filter(mask)
			if err != nil {
				b.Fatalf("Filter failed: %v", err)
			}
			result.Release()
		}
	})

	b.Run("FilterString", func(b *testing.B) {
		df := createStringDataFrame(size)
		defer df.Release()

		mask := createBooleanMask(size, 0.3)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.Filter(mask)
			if err != nil {
				b.Fatalf("Filter failed: %v", err)
			}
			result.Release()
		}
	})

	b.Run("FilterMixed", func(b *testing.B) {
		df := createMixedTypeDataFrameNoBoolean(size)
		defer df.Release()

		mask := createBooleanMask(size, 0.3)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.Filter(mask)
			if err != nil {
				b.Fatalf("Filter failed: %v", err)
			}
			result.Release()
		}
	})
}

// BenchmarkSelectColumns benchmarks selecting different numbers of columns
func BenchmarkSelectColumns(b *testing.B) {
	size := 50000
	df := createWideDataFrame(size, 20) // 20 columns
	defer df.Release()

	testCases := []struct {
		name    string
		columns []string
		numCols int
	}{
		{"Select1Col", []string{"col_0"}, 1},
		{"Select5Cols", []string{"col_0", "col_1", "col_2", "col_3", "col_4"}, 5},
		{"Select10Cols", []string{"col_0", "col_1", "col_2", "col_3", "col_4", "col_5", "col_6", "col_7", "col_8", "col_9"}, 10},
		{"Select15Cols", []string{"col_0", "col_1", "col_2", "col_3", "col_4", "col_5", "col_6", "col_7", "col_8", "col_9", "col_10", "col_11", "col_12", "col_13", "col_14"}, 15},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Select(tc.columns)
				if err != nil {
					b.Fatalf("Select failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

// BenchmarkMemoryAllocation tests memory allocation patterns
func BenchmarkMemoryAllocation(b *testing.B) {
	size := 10000

	b.Run("FilterWithGoAllocator", func(b *testing.B) {
		df := createBenchmarkDataFrameWithPool(size, memory.NewGoAllocator())
		defer df.Release()

		mask := createBooleanMask(size, 0.5)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.Filter(mask)
			if err != nil {
				b.Fatalf("Filter failed: %v", err)
			}
			result.Release()
		}
	})

	b.Run("FilterWithPooling", func(b *testing.B) {
		pool := memory.NewGoAllocator()
		df := createBenchmarkDataFrameWithPool(size, pool)
		defer df.Release()

		mask := createBooleanMaskWithPool(size, 0.5, pool)
		defer mask.Release()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := df.Filter(mask)
			if err != nil {
				b.Fatalf("Filter failed: %v", err)
			}
			result.Release()
		}
	})
}

// Helper functions for benchmark data creation
func createBenchmarkDataFrame(size int) *DataFrame {
	return createBenchmarkDataFrameWithPool(size, memory.NewGoAllocator())
}

func createBenchmarkDataFrameWithPool(size int, pool memory.Allocator) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	categoryBuilder := array.NewStringBuilder(pool)

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i) * 1.5)
		categoryBuilder.Append("category_" + string(rune('A'+i%5)))
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray, categoryArray}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createBooleanMask(size int, selectivity float64) arrow.Array {
	return createBooleanMaskWithPool(size, selectivity, memory.NewGoAllocator())
}

func createBooleanMaskWithPool(size int, selectivity float64, pool memory.Allocator) arrow.Array {
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	threshold := int(float64(size) * selectivity)
	for i := 0; i < size; i++ {
		builder.Append(i < threshold)
	}

	return builder.NewArray()
}

func createComputedColumn(size int) arrow.Array {
	pool := memory.NewGoAllocator()
	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	for i := 0; i < size; i++ {
		builder.Append(float64(i)*2.0 + 10.0)
	}

	return builder.NewArray()
}

func createInt64DataFrame(size int) *DataFrame {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	for i := 0; i < size; i++ {
		builder.Append(int64(i))
	}

	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createFloat64DataFrame(size int) *DataFrame {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "float_col", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	builder := array.NewFloat64Builder(pool)
	for i := 0; i < size; i++ {
		builder.Append(float64(i) * 1.5)
	}

	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createStringDataFrame(size int) *DataFrame {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "string_col", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	builder := array.NewStringBuilder(pool)
	for i := 0; i < size; i++ {
		builder.Append("string_" + string(rune('A'+i%26)))
	}

	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createMixedTypeDataFrameNoBoolean(size int) *DataFrame {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
			{Name: "float_col", Type: arrow.PrimitiveTypes.Float64},
			{Name: "string_col", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	intBuilder := array.NewInt64Builder(pool)
	floatBuilder := array.NewFloat64Builder(pool)
	stringBuilder := array.NewStringBuilder(pool)

	for i := 0; i < size; i++ {
		intBuilder.Append(int64(i))
		floatBuilder.Append(float64(i) * 1.5)
		stringBuilder.Append("item_" + string(rune('A'+i%26)))
	}

	intArray := intBuilder.NewArray()
	defer intArray.Release()

	floatArray := floatBuilder.NewArray()
	defer floatArray.Release()

	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	record := array.NewRecord(schema, []arrow.Array{intArray, floatArray, stringArray}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createWideDataFrame(size, numCols int) *DataFrame {
	pool := memory.NewGoAllocator()

	fields := make([]arrow.Field, numCols)
	builders := make([]*array.Float64Builder, numCols)
	arrays := make([]arrow.Array, numCols)

	for i := 0; i < numCols; i++ {
		fields[i] = arrow.Field{Name: "col_" + string(rune('0'+i/10)) + string(rune('0'+i%10)), Type: arrow.PrimitiveTypes.Float64}
		builders[i] = array.NewFloat64Builder(pool)

		for j := 0; j < size; j++ {
			builders[i].Append(float64(j) * float64(i+1))
		}

		arrays[i] = builders[i].NewArray()
	}

	// Cleanup builders
	for _, builder := range builders {
		builder.Release()
	}

	// Cleanup arrays (deferred)
	defer func() {
		for _, arr := range arrays {
			arr.Release()
		}
	}()

	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, arrays, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func formatSize(size int) string {
	if size >= 1000 {
		return string(rune('0'+size/1000)) + "K"
	}
	return string(rune('0'+size/100)) + "H"
}
