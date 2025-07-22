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

// Helper function to create a test DataFrame
func createTestDataFrame(t *testing.T) *DataFrame {
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