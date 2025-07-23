package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestDataFrame_PublicSort_SingleColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with mixed order
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Create arrays with unsorted data: [3,1,2], ["Charlie","Alice","Bob"], [85.5,95.0,90.2]
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{3, 1, 2}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Charlie", "Alice", "Bob"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{85.5, 95.0, 90.2}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test sorting by ID ascending
	sorted := df.Sort("id", true)
	if sorted.Err() != nil {
		t.Fatalf("Sort failed: %v", sorted.Err())
	}
	defer sorted.Release()

	// Verify row count is preserved
	if sorted.NumRows() != df.NumRows() {
		t.Errorf("Expected %d rows after sort, got %d", df.NumRows(), sorted.NumRows())
	}

	// Verify that we can chain operations
	filtered := sorted.Filter(Col("score").Gt(Lit(90.0)))
	if filtered.Err() != nil {
		t.Fatalf("Filter after sort failed: %v", filtered.Err())
	}
	defer filtered.Release()

	// Should have 2 rows: Alice (95.0) and Bob (90.2)
	if filtered.NumRows() != 2 {
		t.Errorf("Expected 2 rows after filter, got %d", filtered.NumRows())
	}
}

func TestDataFrame_PublicSort_MultiColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with duplicate values to test multi-column sorting
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"IT", "HR", "IT", "HR", "IT"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	salaryBuilder := array.NewInt64Builder(pool)
	salaryBuilder.AppendValues([]int64{70000, 60000, 80000, 60000, 70000}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, salaryArray, nameArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test multi-column sort: department ascending, then salary descending
	sorted := df.SortMultiple([]SortKey{
		{Column: "department", Ascending: true},
		{Column: "salary", Ascending: false},
	})
	if sorted.Err() != nil {
		t.Fatalf("Multi-column sort failed: %v", sorted.Err())
	}
	defer sorted.Release()

	// Verify row count is preserved
	if sorted.NumRows() != df.NumRows() {
		t.Errorf("Expected %d rows after sort, got %d", df.NumRows(), sorted.NumRows())
	}

	// Test chaining with Select
	selected := sorted.Select("name", "salary")
	if selected.Err() != nil {
		t.Fatalf("Select after sort failed: %v", selected.Err())
	}
	defer selected.Release()

	if selected.NumCols() != 2 {
		t.Errorf("Expected 2 columns after select, got %d", selected.NumCols())
	}
}

func TestDataFrame_PublicSort_ErrorHandling(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test sort on non-existent column
	sorted := df.Sort("nonexistent", true)
	if sorted.Err() == nil {
		t.Error("Expected error for non-existent column")
	}

	// Test multi-column sort with non-existent column
	multiSorted := df.SortMultiple([]SortKey{
		{Column: "test", Ascending: true},
		{Column: "nonexistent", Ascending: false},
	})
	if multiSorted.Err() == nil {
		t.Error("Expected error for non-existent column in multi-sort")
	}
}

func TestDataFrame_PublicSort_ErrorPropagation(t *testing.T) {
	// Test that errors are properly propagated in chained operations
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Create a DataFrame with an error first
	errorDF := df.Sort("nonexistent", true)
	if errorDF.Err() == nil {
		t.Fatal("Expected error for non-existent column")
	}

	// Try to sort the error DataFrame - should propagate the error
	sortedErrorDF := errorDF.Sort("test", true)
	if sortedErrorDF.Err() == nil {
		t.Error("Expected error to be propagated in chained sort")
	}

	// Try to multi-sort the error DataFrame - should propagate the error
	multiSortedErrorDF := errorDF.SortMultiple([]SortKey{{Column: "test", Ascending: true}})
	if multiSortedErrorDF.Err() == nil {
		t.Error("Expected error to be propagated in chained multi-sort")
	}
}

func TestDataFrame_PublicSort_ImmutabilityPreserved(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "values", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{3, 1, 2}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	defer record.Release()

	original := NewDataFrame(record)
	defer original.Release()

	// Sort the DataFrame
	sorted := original.Sort("values", true)
	if sorted.Err() != nil {
		t.Fatalf("Sort failed: %v", sorted.Err())
	}
	defer sorted.Release()

	// Verify original is unchanged by checking if we can still query it
	if original.NumRows() != 3 {
		t.Error("Original DataFrame was modified by sort operation")
	}

	if original.NumCols() != 1 {
		t.Error("Original DataFrame was modified by sort operation")
	}

	// Both should have same row/column counts
	if sorted.NumRows() != original.NumRows() {
		t.Error("Sorted DataFrame should have same row count as original")
	}

	if sorted.NumCols() != original.NumCols() {
		t.Error("Sorted DataFrame should have same column count as original")
	}
}
