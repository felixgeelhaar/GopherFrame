package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestDataFramePublicAPI(t *testing.T) {
	// Test creating a DataFrame using the public API
	df := createSampleDataFrame()
	defer df.Release()

	// Test basic properties
	if df.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", df.NumRows())
	}

	if df.NumCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NumCols())
	}

	// Test column names
	names := df.ColumnNames()
	expected := []string{"id", "name", "score"}
	for i, exp := range expected {
		if names[i] != exp {
			t.Errorf("Expected column name %s, got %s", exp, names[i])
		}
	}
}

func TestDataFrameFilter(t *testing.T) {
	df := createSampleDataFrame()
	defer df.Release()

	// Test filter with simple expression: id > 1
	filtered := df.Filter(Col("score").Gt(Lit(90.0)))
	defer filtered.Release()

	// Should return 2 rows (Alice: 95.5, Charlie: 92.1)
	if filtered.NumRows() != 2 {
		t.Errorf("Expected 2 rows after filter, got %d", filtered.NumRows())
	}
}

func TestDataFrameSelect(t *testing.T) {
	df := createSampleDataFrame()
	defer df.Release()

	// Test selecting specific columns
	selected := df.Select("name", "score")
	defer selected.Release()

	// Should have 2 columns
	if selected.NumCols() != 2 {
		t.Errorf("Expected 2 columns after select, got %d", selected.NumCols())
	}

	// Should still have all rows
	if selected.NumRows() != 3 {
		t.Errorf("Expected 3 rows after select, got %d", selected.NumRows())
	}

	// Check column names
	names := selected.ColumnNames()
	expected := []string{"name", "score"}
	for i, exp := range expected {
		if names[i] != exp {
			t.Errorf("Expected column name %s, got %s", exp, names[i])
		}
	}
}

func TestDataFrameWithColumn(t *testing.T) {
	df := createSampleDataFrame()
	defer df.Release()

	// Add a new column: bonus = score * 0.1
	withBonus := df.WithColumn("bonus", Col("score").Mul(Lit(0.1)))
	defer withBonus.Release()

	// Should have 4 columns now
	if withBonus.NumCols() != 4 {
		t.Errorf("Expected 4 columns after WithColumn, got %d", withBonus.NumCols())
	}

	// Check that new column exists
	if !withBonus.HasColumn("bonus") {
		t.Error("Expected new column 'bonus' to exist")
	}
}

func TestDataFrameChaining(t *testing.T) {
	df := createSampleDataFrame()
	defer df.Release()

	// Test chaining operations: filter, select, add column
	result := df.
		Filter(Col("score").Gt(Lit(90.0))).
		Select("name", "score").
		WithColumn("grade", Lit("A"))
	defer result.Release()

	// Should have high-scoring students with grade A
	if result.NumRows() != 2 {
		t.Errorf("Expected 2 rows after chaining, got %d", result.NumRows())
	}

	if result.NumCols() != 3 {
		t.Errorf("Expected 3 columns after chaining, got %d", result.NumCols())
	}

	// Check that grade column exists and has correct values
	if !result.HasColumn("grade") {
		t.Error("Expected 'grade' column to exist")
	}
}

// Helper function to create sample data for testing
func createSampleDataFrame() *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Create ID column
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	// Create name column
	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	// Create score column
	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{95.5, 87.2, 92.1}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	// Create record
	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray}, 3)
	defer record.Release()

	return NewDataFrame(record)
}
