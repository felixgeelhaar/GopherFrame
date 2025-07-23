package gopherframe

import (
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// Benchmark constants
const (
	smallSize  = 1000
	mediumSize = 10000
	largeSize  = 100000
)

// BenchmarkDataFrameCreation measures DataFrame creation performance
func BenchmarkDataFrameCreation(b *testing.B) {
	sizes := []int{smallSize, mediumSize, largeSize}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			schema := arrow.NewSchema(
				[]arrow.Field{
					{Name: "id", Type: arrow.PrimitiveTypes.Int64},
					{Name: "value", Type: arrow.PrimitiveTypes.Float64},
					{Name: "category", Type: arrow.BinaryTypes.String},
				},
				nil,
			)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				record := createBenchmarkRecord(schema, size)
				df := NewDataFrame(record)
				df.Release()
			}
		})
	}
}

// BenchmarkFilter measures filter operation performance
func BenchmarkFilter(b *testing.B) {
	sizes := []int{smallSize, mediumSize, largeSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				filtered := df.Filter(Col("value").Gt(Lit(50.0)))
				filtered.Release()
			}
		})
	}
}

// BenchmarkSelect measures column selection performance
func BenchmarkSelect(b *testing.B) {
	sizes := []int{smallSize, mediumSize, largeSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				selected := df.Select("id", "value")
				selected.Release()
			}
		})
	}
}

// BenchmarkWithColumn measures column computation performance
func BenchmarkWithColumn(b *testing.B) {
	sizes := []int{smallSize, mediumSize, largeSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				newDf := df.WithColumn("doubled", Col("value").Mul(Lit(2.0)))
				newDf.Release()
			}
		})
	}
}

// BenchmarkGroupBySum measures GroupBy with sum aggregation performance
func BenchmarkGroupBySum(b *testing.B) {
	sizes := []int{smallSize, mediumSize, largeSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				grouped := df.GroupBy("category").Agg(Sum("value"))
				grouped.Release()
			}
		})
	}
}

// BenchmarkGroupByMultipleAgg measures GroupBy with multiple aggregations
func BenchmarkGroupByMultipleAgg(b *testing.B) {
	sizes := []int{smallSize, mediumSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				grouped := df.GroupBy("category").Agg(
					Sum("value"),
					Mean("value"),
					Count("value"),
					Min("value"),
					Max("value"),
				)
				grouped.Release()
			}
		})
	}
}

// BenchmarkChainedOperations measures performance of chained operations
func BenchmarkChainedOperations(b *testing.B) {
	sizes := []int{smallSize, mediumSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result := df.
					Filter(Col("value").Gt(Lit(25.0))).
					WithColumn("normalized", Col("value").Div(Lit(100.0))).
					Select("id", "normalized", "category")
				result.Release()
			}
		})
	}
}

// BenchmarkParquetWrite measures Parquet write performance
func BenchmarkParquetWrite(b *testing.B) {
	sizes := []int{smallSize, mediumSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			tempFile := b.TempDir() + "/bench.parquet"
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := WriteParquet(df, tempFile)
				if err != nil {
					b.Fatalf("Failed to write Parquet: %v", err)
				}
			}
		})
	}
}

// BenchmarkParquetRead measures Parquet read performance
func BenchmarkParquetRead(b *testing.B) {
	sizes := []int{smallSize, mediumSize}

	for _, size := range sizes {
		// Create test file
		df := createBenchmarkDataFrame(size)
		tempFile := b.TempDir() + "/bench.parquet"
		if err := WriteParquet(df, tempFile); err != nil {
			b.Fatalf("Failed to create test file: %v", err)
		}
		df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				readDf, err := ReadParquet(tempFile)
				if err != nil {
					b.Fatalf("Failed to read Parquet: %v", err)
				}
				readDf.Release()
			}
		})
	}
}

// BenchmarkCSVWrite measures CSV write performance
func BenchmarkCSVWrite(b *testing.B) {
	sizes := []int{smallSize, mediumSize}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			tempFile := b.TempDir() + "/bench.csv"
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := WriteCSV(df, tempFile)
				if err != nil {
					b.Fatalf("Failed to write CSV: %v", err)
				}
			}
		})
	}
}

// BenchmarkMemoryUsage measures memory allocation patterns
func BenchmarkMemoryUsage(b *testing.B) {
	b.Run("Filter_Memory", func(b *testing.B) {
		df := createBenchmarkDataFrame(mediumSize)
		defer df.Release()

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			filtered := df.Filter(Col("value").Gt(Lit(50.0)))
			filtered.Release()
		}
	})

	b.Run("GroupBy_Memory", func(b *testing.B) {
		df := createBenchmarkDataFrame(mediumSize)
		defer df.Release()

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			grouped := df.GroupBy("category").Agg(Sum("value"))
			grouped.Release()
		}
	})
}

// Helper function to create benchmark data
func createBenchmarkDataFrame(size int) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	record := createBenchmarkRecord(schema, size)
	return NewDataFrame(record)
}

// Helper function to create benchmark record
func createBenchmarkRecord(schema *arrow.Schema, size int) arrow.Record {
	pool := memory.NewGoAllocator()

	// Create ID column
	idBuilder := array.NewInt64Builder(pool)
	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
	}
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	// Create value column with distribution
	valueBuilder := array.NewFloat64Builder(pool)
	for i := 0; i < size; i++ {
		// Create a distribution from 0 to 100
		value := float64(i%100) + float64(i%10)*0.1
		valueBuilder.Append(value)
	}
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	// Create category column with reasonable cardinality
	categoryBuilder := array.NewStringBuilder(pool)
	categories := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < size; i++ {
		categoryBuilder.Append(categories[i%len(categories)])
	}
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	return array.NewRecord(schema, []arrow.Array{idArray, valueArray, categoryArray}, int64(size))
}
