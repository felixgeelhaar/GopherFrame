package application

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/aggregation"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

func TestNewDataFrameService(t *testing.T) {
	service := NewDataFrameService()
	if service == nil {
		t.Error("NewDataFrameService should not return nil")
	}
}

func TestDataFrameService_GroupBy(t *testing.T) {
	service := NewDataFrameService()

	// Create test dataframe
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

	categoryArr := categoryBuilder.NewArray()
	valueArr := valueBuilder.NewArray()
	defer categoryArr.Release()
	defer valueArr.Release()

	record := array.NewRecord(schema, []arrow.Array{categoryArr, valueArr}, 4)
	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Create aggregation request
	request := aggregation.GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []aggregation.AggregationSpec{
			{Column: "value", Type: aggregation.Sum, Alias: "value_sum"},
		},
	}

	result := service.GroupBy(df, request)
	if result.Error != nil {
		t.Fatalf("GroupBy failed: %v", result.Error)
	}
	defer result.DataFrame.Release()

	// Verify results
	if result.DataFrame.NumRows() != 2 {
		t.Errorf("Expected 2 rows, got %d", result.DataFrame.NumRows())
	}

	if result.DataFrame.NumCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", result.DataFrame.NumCols())
	}
}

func TestDataFrameService_GroupBy_ErrorCases(t *testing.T) {
	service := NewDataFrameService()

	// Create empty dataframe
	schema := arrow.NewSchema([]arrow.Field{}, nil)
	emptyRecord := array.NewRecord(schema, []arrow.Array{}, 0)
	emptyDF := dataframe.NewDataFrame(emptyRecord)
	defer emptyDF.Release()

	// Test with no group columns
	request := aggregation.GroupByRequest{
		GroupColumns: []string{},
		Aggregations: []aggregation.AggregationSpec{
			{Column: "value", Type: aggregation.Sum, Alias: "sum"},
		},
	}

	result := service.GroupBy(emptyDF, request)
	if result.Error == nil {
		t.Error("Expected error for empty group columns")
	}

	// Test with no aggregations
	request = aggregation.GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []aggregation.AggregationSpec{},
	}

	result = service.GroupBy(emptyDF, request)
	if result.Error == nil {
		t.Error("Expected error for empty aggregations")
	}
}

func TestDataFrameService_MultipleAggregations(t *testing.T) {
	service := NewDataFrameService()

	// Create test dataframe
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

	categoryBuilder.AppendValues([]string{"A", "A", "B", "B"}, nil)
	valueBuilder.AppendValues([]float64{10.0, 20.0, 30.0, 40.0}, nil)

	categoryArr := categoryBuilder.NewArray()
	valueArr := valueBuilder.NewArray()
	defer categoryArr.Release()
	defer valueArr.Release()

	record := array.NewRecord(schema, []arrow.Array{categoryArr, valueArr}, 4)
	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test multiple aggregations
	request := aggregation.GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []aggregation.AggregationSpec{
			{Column: "value", Type: aggregation.Sum, Alias: "sum"},
			{Column: "value", Type: aggregation.Mean, Alias: "avg"},
			{Column: "value", Type: aggregation.Count, Alias: "count"},
		},
	}

	result := service.GroupBy(df, request)
	if result.Error != nil {
		t.Fatalf("Multiple aggregations failed: %v", result.Error)
	}
	defer result.DataFrame.Release()

	// Should have category + 3 aggregation columns = 4 total
	expectedCols := int64(4)
	if result.DataFrame.NumCols() != expectedCols {
		t.Errorf("Expected %d columns, got %d", expectedCols, result.DataFrame.NumCols())
	}

	// Check column names
	names := result.DataFrame.ColumnNames()
	expectedNames := []string{"category", "sum", "avg", "count"}

	if len(names) != len(expectedNames) {
		t.Fatalf("Expected %d column names, got %d", len(expectedNames), len(names))
	}

	for i, expected := range expectedNames {
		if names[i] != expected {
			t.Errorf("Column %d: expected %s, got %s", i, expected, names[i])
		}
	}
}
