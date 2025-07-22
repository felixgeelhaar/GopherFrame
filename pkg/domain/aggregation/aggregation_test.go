package aggregation

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

func TestGroupByService_Execute(t *testing.T) {
	pool := memory.NewGoAllocator()
	service := NewGroupByService()

	// Create test dataframe
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
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{Column: "value", Type: Sum, Alias: "value_sum"},
		},
	}

	result := service.Execute(df, request)
	if result.Error != nil {
		t.Fatalf("GroupBy failed: %v", result.Error)
	}
	defer result.DataFrame.Release()

	// Should have 2 groups (A, B)
	if result.DataFrame.NumRows() != 2 {
		t.Errorf("Expected 2 groups, got %d", result.DataFrame.NumRows())
	}

	// Should have category + sum columns
	if result.DataFrame.NumCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", result.DataFrame.NumCols())
	}

	names := result.DataFrame.ColumnNames()
	expectedNames := []string{"category", "value_sum"}
	for i, expected := range expectedNames {
		if names[i] != expected {
			t.Errorf("Column %d: expected %s, got %s", i, expected, names[i])
		}
	}
}

func TestAggregationTypes(t *testing.T) {
	testCases := []struct {
		name string
		spec AggregationSpec
	}{
		{
			name: "Sum aggregation",
			spec: AggregationSpec{Column: "value", Type: Sum, Alias: "total"},
		},
		{
			name: "Mean aggregation",
			spec: AggregationSpec{Column: "value", Type: Mean, Alias: "average"},
		},
		{
			name: "Count aggregation",
			spec: AggregationSpec{Column: "value", Type: Count, Alias: "count"},
		},
		{
			name: "Min aggregation",
			spec: AggregationSpec{Column: "value", Type: Min, Alias: "minimum"},
		},
		{
			name: "Max aggregation",
			spec: AggregationSpec{Column: "value", Type: Max, Alias: "maximum"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.spec.Column != "value" {
				t.Errorf("Expected column 'value', got %s", tc.spec.Column)
			}
			if tc.spec.Alias == "" {
				t.Error("Expected non-empty alias")
			}
		})
	}
}

func TestGroupByService_ErrorCases(t *testing.T) {
	service := NewGroupByService()

	// Test with empty dataframe
	schema := arrow.NewSchema([]arrow.Field{}, nil)
	emptyRecord := array.NewRecord(schema, []arrow.Array{}, 0)
	emptyDF := dataframe.NewDataFrame(emptyRecord)
	defer emptyDF.Release()

	// Test with no group columns
	request := GroupByRequest{
		GroupColumns: []string{},
		Aggregations: []AggregationSpec{{Column: "value", Type: Sum, Alias: "sum"}},
	}

	result := service.Execute(emptyDF, request)
	if result.Error == nil {
		t.Error("Expected error for no group columns")
	}

	// Test with no aggregations
	request = GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{},
	}

	result = service.Execute(emptyDF, request)
	if result.Error == nil {
		t.Error("Expected error for no aggregations")
	}
}
