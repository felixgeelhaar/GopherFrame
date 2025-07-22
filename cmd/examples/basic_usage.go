// Package main demonstrates basic usage of GopherFrame DataFrame operations.
package main

import (
	"fmt"
	"log"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/gopherFrame"
)

func main() {
	fmt.Println("🐹 GopherFrame - Production-First DataFrame for Go")
	fmt.Println("==================================================")
	
	// Create sample data - this would typically come from files or databases
	df := createSampleData()
	defer df.Release()
	
	fmt.Printf("Original DataFrame: %d rows, %d columns\n", df.NumRows(), df.NumCols())
	fmt.Printf("Columns: %v\n\n", df.ColumnNames())
	
	// Example 1: Filter high-performing students (score > 90)
	fmt.Println("📊 Filter: Students with score > 90")
	highPerformers := df.Filter(gf.Col("score").Gt(gf.Lit(90.0)))
	defer highPerformers.Release()
	
	if highPerformers.Err() != nil {
		log.Fatalf("Filter error: %v", highPerformers.Err())
	}
	
	fmt.Printf("High performers: %d students\n\n", highPerformers.NumRows())
	
	// Example 2: Select specific columns
	fmt.Println("🎯 Select: Name and score columns only")
	nameScore := df.Select("name", "score")
	defer nameScore.Release()
	
	if nameScore.Err() != nil {
		log.Fatalf("Select error: %v", nameScore.Err())
	}
	
	fmt.Printf("Selected columns: %v\n\n", nameScore.ColumnNames())
	
	// Example 3: Add calculated column (bonus = score * 0.1)
	fmt.Println("➕ WithColumn: Adding bonus calculation")
	withBonus := df.WithColumn("bonus", gf.Col("score").Mul(gf.Lit(0.1)))
	defer withBonus.Release()
	
	if withBonus.Err() != nil {
		log.Fatalf("WithColumn error: %v", withBonus.Err())
	}
	
	fmt.Printf("With bonus column: %v\n\n", withBonus.ColumnNames())
	
	// Example 4: Chain operations
	fmt.Println("🔗 Chaining: Filter → Select → Add Grade")
	result := df.
		Filter(gf.Col("score").Gt(gf.Lit(85.0))).
		Select("name", "score").
		WithColumn("grade", gf.Lit("A"))
	defer result.Release()
	
	if result.Err() != nil {
		log.Fatalf("Chaining error: %v", result.Err())
	}
	
	fmt.Printf("Final result: %d rows, columns: %v\n", result.NumRows(), result.ColumnNames())
	
	fmt.Println("\n✅ All operations completed successfully!")
	fmt.Println("🚀 GopherFrame is ready for production workloads!")
}

func createSampleData() *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
			{Name: "department", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	pool := memory.NewGoAllocator()
	
	// Create ID column
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()
	
	// Create name column
	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()
	
	// Create score column
	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{95.5, 87.2, 92.1, 88.8, 94.3}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()
	
	// Create department column
	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"Engineering", "Design", "Engineering", "Marketing", "Engineering"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()
	
	// Create record
	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray, deptArray}, 5)
	defer record.Release()
	
	return gf.NewDataFrame(record)
}