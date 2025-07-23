package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestDataFrameCreation(t *testing.T) {
	// Test creating a DataFrame from an Arrow Record
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Build test data
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

	// Create record
	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 3)
	defer record.Release()

	// Create DataFrame
	df := NewDataFrame(record)
	defer df.Release()

	// Test basic properties
	if df.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", df.NumRows())
	}

	if df.NumCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", df.NumCols())
	}

	// Test column names
	names := df.ColumnNames()
	expectedNames := []string{"id", "name"}
	if len(names) != len(expectedNames) {
		t.Errorf("Expected %d column names, got %d", len(expectedNames), len(names))
	}
	for i, expected := range expectedNames {
		if names[i] != expected {
			t.Errorf("Expected column name %s, got %s", expected, names[i])
		}
	}

	// Test schema
	if df.Schema() == nil {
		t.Error("Expected non-nil schema")
	}

	if !df.Schema().Equal(schema) {
		t.Error("Schema should match the input schema")
	}
}

func TestDataFrameColumnAccess(t *testing.T) {
	// Create test DataFrame
	df := createTestDataFrame(t)
	defer df.Release()

	// Test column access by name
	idSeries, err := df.Column("id")
	if err != nil {
		t.Errorf("Failed to get id column: %v", err)
	}
	defer idSeries.Release()

	if idSeries.Name() != "id" {
		t.Errorf("Expected column name 'id', got '%s'", idSeries.Name())
	}

	if idSeries.Len() != 3 {
		t.Errorf("Expected column length 3, got %d", idSeries.Len())
	}

	// Test non-existent column
	_, err = df.Column("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent column")
	}

	// Test HasColumn
	if !df.HasColumn("id") {
		t.Error("Expected HasColumn('id') to be true")
	}

	if df.HasColumn("nonexistent") {
		t.Error("Expected HasColumn('nonexistent') to be false")
	}
}

func TestDataFrameValidation(t *testing.T) {
	// Create test DataFrame
	df := createTestDataFrame(t)
	defer df.Release()

	// Test validation
	if err := df.Validate(); err != nil {
		t.Errorf("Expected valid DataFrame, got error: %v", err)
	}
}

func TestNewDataFrameWithAllocator(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray}, 3)
	defer record.Release()

	df := NewDataFrameWithAllocator(record, pool)
	defer df.Release()

	if df.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", df.NumRows())
	}
}

func TestNewDataFrameFromStorage(t *testing.T) {
	// This test would require a proper storage backend implementation
	// For now, we'll skip this test as it requires complex setup
	t.Skip("NewDataFrameFromStorage requires storage backend implementation")
}

func TestDataFrame_Record(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	record := df.Record()
	if record == nil {
		t.Error("Expected non-nil record")
	}

	if record.NumRows() != df.NumRows() {
		t.Errorf("Expected record rows %d, got %d", df.NumRows(), record.NumRows())
	}
}

func TestDataFrame_ColumnAt(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	series, err := df.ColumnAt(0)
	if err != nil {
		t.Errorf("Failed to get column at index 0: %v", err)
	}
	defer series.Release()

	if series.Name() != "id" {
		t.Errorf("Expected column name 'id', got '%s'", series.Name())
	}

	// Test invalid index
	_, err = df.ColumnAt(10)
	if err == nil {
		t.Error("Expected error for invalid column index")
	}
}

func TestDataFrame_Columns(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	columns := df.Columns()
	if len(columns) != 3 {
		t.Errorf("Expected 3 columns, got %d", len(columns))
	}

	for _, series := range columns {
		defer series.Release()
	}

	if columns[0].Name() != "id" {
		t.Errorf("Expected first column name 'id', got '%s'", columns[0].Name())
	}
}

func TestDataFrame_Equal(t *testing.T) {
	df1 := createTestDataFrame(t)
	defer df1.Release()

	df2 := createTestDataFrame(t)
	defer df2.Release()

	if !df1.Equal(df2) {
		t.Error("Expected DataFrames to be equal")
	}

	// Test with different data
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{4, 5, 6}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray}, 3)
	defer record.Release()

	df3 := NewDataFrame(record)
	defer df3.Release()

	if df1.Equal(df3) {
		t.Error("Expected DataFrames to be different")
	}
}

func TestDataFrame_String(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	str := df.String()
	if str == "" {
		t.Error("Expected non-empty string representation")
	}

	// Check that string contains column names
	if !contains(str, "id") {
		t.Error("Expected string to contain 'id' column name")
	}
}

func TestDataFrame_Clone(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	cloned := df.Clone()
	defer cloned.Release()

	if !df.Equal(cloned) {
		t.Error("Expected cloned DataFrame to be equal to original")
	}

	if df.NumRows() != cloned.NumRows() {
		t.Errorf("Expected cloned rows %d, got %d", df.NumRows(), cloned.NumRows())
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestDataFrame_Select(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	// Test selecting specific columns
	selected, err := df.Select([]string{"id", "name"})
	if err != nil {
		t.Fatalf("Failed to select columns: %v", err)
	}
	defer selected.Release()

	if selected.NumCols() != 2 {
		t.Errorf("Expected 2 columns after select, got %d", selected.NumCols())
	}

	if selected.NumRows() != df.NumRows() {
		t.Errorf("Expected same number of rows after select, got %d", selected.NumRows())
	}

	// Test selecting non-existent column
	_, err = df.Select([]string{"nonexistent"})
	if err == nil {
		t.Error("Expected error when selecting non-existent column")
	}
}

func TestDataFrame_WithColumn(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	// Create a new column
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{10, 20, 30}, nil)
	newColumnArray := builder.NewArray()
	defer newColumnArray.Release()

	// Add the column
	newDF, err := df.WithColumn("new_col", newColumnArray)
	if err != nil {
		t.Fatalf("Failed to add column: %v", err)
	}
	defer newDF.Release()

	if newDF.NumCols() != df.NumCols()+1 {
		t.Errorf("Expected %d columns after adding column, got %d", df.NumCols()+1, newDF.NumCols())
	}

	if !newDF.HasColumn("new_col") {
		t.Error("Expected new DataFrame to have 'new_col' column")
	}

	// Test adding column with wrong length
	wrongLengthBuilder := array.NewInt64Builder(pool)
	wrongLengthBuilder.AppendValues([]int64{1, 2}, nil) // Wrong length (2 instead of 3)
	wrongLengthArray := wrongLengthBuilder.NewArray()
	defer wrongLengthArray.Release()

	_, err = df.WithColumn("wrong_len", wrongLengthArray)
	if err == nil {
		t.Error("Expected error when adding column with wrong length")
	}
}

func TestDataFrame_Filter(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	// Create a filter mask - select rows where id > 1
	pool := memory.NewGoAllocator()
	maskBuilder := array.NewBooleanBuilder(pool)
	maskBuilder.AppendValues([]bool{false, true, true}, nil) // Skip first row
	mask := maskBuilder.NewArray()
	defer mask.Release()

	filtered, err := df.Filter(mask)
	if err != nil {
		t.Fatalf("Failed to filter DataFrame: %v", err)
	}
	defer filtered.Release()

	if filtered.NumRows() != 2 {
		t.Errorf("Expected 2 rows after filtering, got %d", filtered.NumRows())
	}

	if filtered.NumCols() != df.NumCols() {
		t.Errorf("Expected same number of columns after filtering, got %d", filtered.NumCols())
	}

	// Test with wrong mask length
	wrongMaskBuilder := array.NewBooleanBuilder(pool)
	wrongMaskBuilder.AppendValues([]bool{true, false}, nil) // Wrong length
	wrongMask := wrongMaskBuilder.NewArray()
	defer wrongMask.Release()

	_, err = df.Filter(wrongMask)
	if err == nil {
		t.Error("Expected error when filtering with wrong mask length")
	}
}

// Helper function to create a test DataFrame
func createTestDataFrame(_ *testing.T) *DataFrame {
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
