package main

import (
	"fmt"
	"log"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gopherframe "github.com/felixgeelhaar/GopherFrame"
)

func main() {
	pool := memory.NewGoAllocator()

	// Create sample data
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Int64},
			{Name: "experience", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	// Build columns
	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"Engineering", "Sales", "Engineering", "Marketing", "Sales"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	salaryBuilder := array.NewInt64Builder(pool)
	salaryBuilder.AppendValues([]int64{95000, 65000, 105000, 70000, 72000}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	expBuilder := array.NewInt64Builder(pool)
	expBuilder.AppendValues([]int64{5, 3, 8, 4, 6}, nil)
	expArray := expBuilder.NewArray()
	defer expArray.Release()

	// Create record and DataFrame
	record := array.NewRecord(schema, []arrow.Array{nameArray, deptArray, salaryArray, expArray}, 5)
	defer record.Release()

	df := gopherframe.NewDataFrame(record)
	defer df.Release()

	fmt.Println("Original DataFrame:")
	printDataFrame(df)

	// Sort by salary descending
	sortedBySalary := df.Sort("salary", false)
	if sortedBySalary.Err() != nil {
		log.Fatalf("Sort by salary failed: %v", sortedBySalary.Err())
	}
	defer sortedBySalary.Release()

	fmt.Println("\nSorted by salary (descending):")
	printDataFrame(sortedBySalary)

	// Multi-column sort: department ascending, then salary descending
	multiSorted := df.SortMultiple([]gopherframe.SortKey{
		{Column: "department", Ascending: true},
		{Column: "salary", Ascending: false},
	})
	if multiSorted.Err() != nil {
		log.Fatalf("Multi-column sort failed: %v", multiSorted.Err())
	}
	defer multiSorted.Release()

	fmt.Println("\nSorted by department (asc), then salary (desc):")
	printDataFrame(multiSorted)

	// Chain sorting with filtering
	highEarners := df.
		Filter(gopherframe.Col("salary").Gt(gopherframe.Lit(int64(70000)))).
		Sort("experience", false)

	if highEarners.Err() != nil {
		log.Fatalf("Chained operations failed: %v", highEarners.Err())
	}
	defer highEarners.Release()

	fmt.Println("\nHigh earners (>70k), sorted by experience (desc):")
	printDataFrame(highEarners)
}

// Helper function to print DataFrame contents
func printDataFrame(df *gopherframe.DataFrame) {
	if df.Err() != nil {
		fmt.Printf("DataFrame error: %v\n", df.Err())
		return
	}

	fmt.Printf("Rows: %d, Columns: %d\n", df.NumRows(), df.NumCols())
	fmt.Printf("Columns: %v\n", df.ColumnNames())
	fmt.Println("Schema:", df.Schema())
}
