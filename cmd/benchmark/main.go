package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
	fmt.Println("GopherFrame Benchmark Suite")
	fmt.Println("===========================")

	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		fmt.Printf("\nBenchmarking with %d rows:\n", size)
		fmt.Println("-" + fmt.Sprintf("%d", len(fmt.Sprintf("Benchmarking with %d rows:", size))))

		df := createBenchmarkDataFrame(size)
		defer df.Release()

		// Benchmark Filter
		start := time.Now()
		filtered := df.Filter(gf.Col("value").Gt(gf.Lit(50.0)))
		filterTime := time.Since(start)
		fmt.Printf("Filter operation: %v\n", filterTime)
		fmt.Printf("  Filtered rows: %d\n", filtered.NumRows())
		filtered.Release()

		// Benchmark Select
		start = time.Now()
		selected := df.Select("id", "value")
		selectTime := time.Since(start)
		fmt.Printf("Select operation: %v\n", selectTime)
		selected.Release()

		// Benchmark WithColumn
		start = time.Now()
		withCol := df.WithColumn("doubled", gf.Col("value").Mul(gf.Lit(2.0)))
		withColTime := time.Since(start)
		fmt.Printf("WithColumn operation: %v\n", withColTime)
		withCol.Release()

		// Benchmark GroupBy
		start = time.Now()
		grouped := df.GroupBy("category").Agg(gf.Sum("value"))
		groupByTime := time.Since(start)
		fmt.Printf("GroupBy+Sum operation: %v\n", groupByTime)
		fmt.Printf("  Groups: %d\n", grouped.NumRows())
		grouped.Release()

		// Benchmark chained operations
		start = time.Now()
		chained := df.
			Filter(gf.Col("value").Gt(gf.Lit(25.0))).
			WithColumn("normalized", gf.Col("value").Div(gf.Lit(100.0))).
			Select("id", "normalized", "category")
		chainedTime := time.Since(start)
		fmt.Printf("Chained operations: %v\n", chainedTime)
		fmt.Printf("  Result rows: %d\n", chained.NumRows())
		chained.Release()

		// Calculate throughput
		throughput := float64(size) / filterTime.Seconds()
		fmt.Printf("\nThroughput (Filter): %.0f rows/second\n", throughput)
	}

	// Test I/O performance
	fmt.Println("\nI/O Performance:")
	fmt.Println("================")

	df := createBenchmarkDataFrame(10000)
	defer df.Release()

	// Parquet write
	start := time.Now()
	err := gf.WriteParquet(df, "benchmark_test.parquet")
	if err != nil {
		log.Fatalf("Failed to write Parquet: %v", err)
	}
	parquetWriteTime := time.Since(start)
	fmt.Printf("Parquet write (10k rows): %v\n", parquetWriteTime)

	// Parquet read
	start = time.Now()
	readDF, err := gf.ReadParquet("benchmark_test.parquet")
	if err != nil {
		log.Fatalf("Failed to read Parquet: %v", err)
	}
	parquetReadTime := time.Since(start)
	fmt.Printf("Parquet read (10k rows): %v\n", parquetReadTime)
	readDF.Release()

	// CSV write
	start = time.Now()
	err = gf.WriteCSV(df, "benchmark_test.csv")
	if err != nil {
		log.Fatalf("Failed to write CSV: %v", err)
	}
	csvWriteTime := time.Since(start)
	fmt.Printf("CSV write (10k rows): %v\n", csvWriteTime)

	// CSV read
	start = time.Now()
	readCSV, err := gf.ReadCSV("benchmark_test.csv")
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}
	csvReadTime := time.Since(start)
	fmt.Printf("CSV read (10k rows): %v\n", csvReadTime)
	readCSV.Release()

	fmt.Println("\nBenchmark completed!")
}

func createBenchmarkDataFrame(size int) *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Create ID column
	idBuilder := array.NewInt64Builder(pool)
	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
	}
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	// Create value column
	valueBuilder := array.NewFloat64Builder(pool)
	for i := 0; i < size; i++ {
		value := float64(i%100) + float64(i%10)*0.1
		valueBuilder.Append(value)
	}
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	// Create category column
	categoryBuilder := array.NewStringBuilder(pool)
	categories := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < size; i++ {
		categoryBuilder.Append(categories[i%len(categories)])
	}
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray, categoryArray}, int64(size))
	return gf.NewDataFrame(record)
}
