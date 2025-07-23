package main

import (
	"fmt"
	"log"
	"os"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
	fmt.Println("🐹 GopherFrame Complete Demonstration")
	fmt.Println("====================================")

	// 1. Create a DataFrame with sample data
	fmt.Println("\n1. Creating DataFrame...")
	df := createSampleDataFrame()
	defer df.Release()

	fmt.Printf("Created DataFrame with %d rows and %d columns\n", df.NumRows(), df.NumCols())
	fmt.Printf("Columns: %v\n", df.ColumnNames())

	// 2. Basic operations
	fmt.Println("\n2. Basic Operations...")

	// Filter operation
	highScores := df.Filter(gf.Col("score").Gt(gf.Lit(85.0)))
	defer highScores.Release()
	fmt.Printf("High scores (>85): %d rows\n", highScores.NumRows())

	// Select operation
	summary := df.Select("name", "score", "department")
	defer summary.Release()
	fmt.Printf("Selected columns: %v\n", summary.ColumnNames())

	// WithColumn operation
	withBonus := df.WithColumn("bonus", gf.Col("score").Mul(gf.Lit(0.1)))
	defer withBonus.Release()
	fmt.Printf("Added bonus column: %d columns total\n", withBonus.NumCols())

	// 3. Chained operations
	fmt.Println("\n3. Chained Operations...")
	result := df.
		Filter(gf.Col("score").Gt(gf.Lit(75.0))).
		WithColumn("grade", gf.Col("score").Div(gf.Lit(10.0))).
		Select("name", "department", "score", "grade")
	defer result.Release()

	fmt.Printf("Chained result: %d rows, %d columns\n", result.NumRows(), result.NumCols())

	// 4. GroupBy and Aggregation
	fmt.Println("\n4. GroupBy Operations...")
	groupedResults := df.GroupBy("department").Agg(
		gf.Sum("score").As("total_score"),
		gf.Mean("score").As("avg_score"),
		gf.Count("score").As("count"),
		gf.Max("score").As("max_score"),
		gf.Min("score").As("min_score"),
	)
	defer groupedResults.Release()

	fmt.Printf("Department stats: %d departments\n", groupedResults.NumRows())
	fmt.Printf("Aggregation columns: %v\n", groupedResults.ColumnNames())

	// 5. I/O Operations
	fmt.Println("\n5. I/O Operations...")

	// Write to Parquet
	parquetFile := "demo_output.parquet"
	err := gf.WriteParquet(df, parquetFile)
	if err != nil {
		log.Printf("Parquet write error: %v", err)
	} else {
		fmt.Printf("✓ Written to %s\n", parquetFile)

		// Read back from Parquet
		readDF, err := gf.ReadParquet(parquetFile)
		if err != nil {
			log.Printf("Parquet read error: %v", err)
		} else {
			defer readDF.Release()
			fmt.Printf("✓ Read back %d rows from Parquet\n", readDF.NumRows())
		}

		// Clean up
		if err := os.Remove(parquetFile); err != nil {
			log.Printf("Warning: failed to remove temporary file %s: %v", parquetFile, err)
		}
	}

	// Write to CSV
	csvFile := "demo_output.csv"
	err = gf.WriteCSV(df, csvFile)
	if err != nil {
		log.Printf("CSV write error: %v", err)
	} else {
		fmt.Printf("✓ Written to %s\n", csvFile)

		// Read back from CSV
		readCSV, err := gf.ReadCSV(csvFile)
		if err != nil {
			log.Printf("CSV read error: %v", err)
		} else {
			defer readCSV.Release()
			fmt.Printf("✓ Read back %d rows from CSV\n", readCSV.NumRows())
		}

		// Clean up
		if err := os.Remove(csvFile); err != nil {
			log.Printf("Warning: failed to remove temporary file %s: %v", csvFile, err)
		}
	}

	// 6. Performance demonstration
	fmt.Println("\n6. Performance Test...")
	largeDF := createLargeDataFrame(50000)
	defer largeDF.Release()

	fmt.Printf("Created large DataFrame: %d rows\n", largeDF.NumRows())

	// Complex chain of operations
	complexResult := largeDF.
		Filter(gf.Col("value").Gt(gf.Lit(500.0))).
		WithColumn("category_upper", gf.Col("category")). // String ops not implemented yet
		GroupBy("category").
		Agg(gf.Mean("value").As("avg_value"))
	defer complexResult.Release()

	fmt.Printf("Complex operation result: %d groups\n", complexResult.NumRows())

	fmt.Println("\n🎉 GopherFrame demonstration completed successfully!")
	fmt.Println("\nKey Features Demonstrated:")
	fmt.Println("• DataFrame creation and manipulation")
	fmt.Println("• Filter, Select, WithColumn operations")
	fmt.Println("• Method chaining for fluent API")
	fmt.Println("• GroupBy with multiple aggregations")
	fmt.Println("• Parquet and CSV I/O")
	fmt.Println("• High-performance operations on large datasets")
	fmt.Println("• Memory-safe operations with proper cleanup")
}

func createSampleDataFrame() *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Sample employee data
	ids := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	names := []string{
		"Alice", "Bob", "Charlie", "Diana", "Eve",
		"Frank", "Grace", "Henry", "Ivy", "Jack",
	}
	departments := []string{
		"Engineering", "Engineering", "Sales", "Marketing", "Engineering",
		"Sales", "Marketing", "Engineering", "Sales", "Marketing",
	}
	scores := []float64{92.5, 87.2, 78.9, 94.1, 89.7, 82.3, 91.8, 85.6, 88.4, 79.2}

	// Build arrays
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues(ids, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues(names, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues(departments, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues(scores, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, nameArray, deptArray, scoreArray},
		int64(len(ids)),
	)

	return gf.NewDataFrame(record)
}

func createLargeDataFrame(size int) *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Create large dataset
	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	categoryBuilder := array.NewStringBuilder(pool)

	categories := []string{"A", "B", "C", "D", "E"}

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i%1000) + float64(i%100)*0.01)
		categoryBuilder.Append(categories[i%len(categories)])
	}

	idArray := idBuilder.NewArray()
	valueArray := valueBuilder.NewArray()
	categoryArray := categoryBuilder.NewArray()

	defer idArray.Release()
	defer valueArray.Release()
	defer categoryArray.Release()

	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, valueArray, categoryArray},
		int64(size),
	)

	return gf.NewDataFrame(record)
}
