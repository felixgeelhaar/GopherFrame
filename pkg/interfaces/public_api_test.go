package interfaces

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestNewDataFrame_DDD(t *testing.T) {
	// Create test data using Arrow
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	categoryBuilder := array.NewStringBuilder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	idBuilder.AppendValues([]int64{1, 2, 3, 4}, nil)
	categoryBuilder.AppendValues([]string{"A", "B", "A", "B"}, nil)
	valueBuilder.AppendValues([]float64{10.0, 20.0, 30.0, 40.0}, nil)

	idArray := idBuilder.NewArray()
	categoryArray := categoryBuilder.NewArray()
	valueArray := valueBuilder.NewArray()
	defer idArray.Release()
	defer categoryArray.Release()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, categoryArray, valueArray}, 4)

	// Test the DDD-structured public API
	df := NewDataFrame(record)
	defer df.Release()

	// Test basic properties
	if df.NumRows() != 4 {
		t.Errorf("Expected 4 rows, got %d", df.NumRows())
	}

	if df.NumCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NumCols())
	}

	// Test column names
	colNames := df.ColumnNames()
	expectedNames := []string{"id", "category", "value"}
	for i, expected := range expectedNames {
		if colNames[i] != expected {
			t.Errorf("Column %d: expected %s, got %s", i, expected, colNames[i])
		}
	}

	// Test HasColumn
	if !df.HasColumn("category") {
		t.Error("Expected HasColumn('category') to return true")
	}

	if df.HasColumn("nonexistent") {
		t.Error("Expected HasColumn('nonexistent') to return false")
	}
}

func TestGroupBy_DDD(t *testing.T) {
	// Create test data
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	categoryBuilder.AppendValues([]string{"A", "B", "A", "B"}, nil)
	valueBuilder.AppendValues([]float64{10.0, 20.0, 30.0, 40.0}, nil)

	categoryArray := categoryBuilder.NewArray()
	valueArray := valueBuilder.NewArray()
	defer categoryArray.Release()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 4)

	df := NewDataFrame(record)
	defer df.Release()

	// Test group-by with sum aggregation using the new DDD structure
	result := df.GroupBy("category").Sum("value")
	if result.Err() != nil {
		t.Fatalf("GroupBy Sum failed: %v", result.Err())
	}
	defer result.Release()

	// Should have 2 groups (A and B)
	if result.NumRows() != 2 {
		t.Errorf("Expected 2 groups, got %d", result.NumRows())
	}

	// Should have 2 columns (category and value_sum)
	if result.NumCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", result.NumCols())
	}
}

func TestAggregationSpecs_DDD(t *testing.T) {
	// Test aggregation specification builders
	sumSpec := Sum("value")
	if sumSpec.Column != "value" {
		t.Errorf("Expected column 'value', got '%s'", sumSpec.Column)
	}
	if sumSpec.Type != SumAgg {
		t.Errorf("Expected SumAgg type, got %d", sumSpec.Type)
	}
	if sumSpec.Alias != "value_sum" {
		t.Errorf("Expected alias 'value_sum', got '%s'", sumSpec.Alias)
	}

	// Test custom alias
	customSum := Sum("value").As("total")
	if customSum.Alias != "total" {
		t.Errorf("Expected alias 'total', got '%s'", customSum.Alias)
	}

	// Test other aggregation types
	meanSpec := Mean("value")
	if meanSpec.Type != MeanAgg {
		t.Errorf("Expected MeanAgg type, got %d", meanSpec.Type)
	}

	countSpec := Count("value")
	if countSpec.Type != CountAgg {
		t.Errorf("Expected CountAgg type, got %d", countSpec.Type)
	}
}
