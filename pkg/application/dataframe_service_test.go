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
		t.Fatal("NewDataFrameService should not return nil")
	}

	if service.groupByService == nil {
		t.Error("DataFrameService should have groupBy service initialized")
	}
}

func TestDataFrameService_LoadFromParquet(t *testing.T) {
	service := NewDataFrameService()

	// Test with non-existent file - should return error
	df, err := service.LoadFromParquet("nonexistent.parquet")
	if err == nil {
		t.Error("LoadFromParquet with non-existent file should return error")
	}

	if df != nil {
		t.Error("DataFrame should be nil when load fails")
	}
}

func TestDataFrameService_SaveToParquet(t *testing.T) {
	service := NewDataFrameService()

	// Create a test DataFrame
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Build test data
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 3)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test saving to invalid path - should return error
	err := service.SaveToParquet(df, "/invalid/path/test.parquet")
	if err == nil {
		t.Error("SaveToParquet with invalid path should return error")
	}
}

func TestDataFrameService_GroupBy(t *testing.T) {
	service := NewDataFrameService()

	// Create test DataFrame with data suitable for grouping
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
			{Name: "bonus", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build test data
	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"Engineering", "Engineering", "Sales", "Sales"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryBuilder.AppendValues([]float64{100000, 120000, 80000, 85000}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	bonusBuilder := array.NewFloat64Builder(pool)
	bonusBuilder.AppendValues([]float64{10000, 15000, 8000, 9000}, nil)
	bonusArray := bonusBuilder.NewArray()
	defer bonusArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, salaryArray, bonusArray}, 4)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test GroupBy
	request := aggregation.GroupByRequest{
		GroupColumns: []string{"department"},
		Aggregations: []aggregation.AggregationSpec{
			{Column: "salary", Type: aggregation.Sum, Alias: "total_salary"},
			{Column: "bonus", Type: aggregation.Mean, Alias: "avg_bonus"},
		},
	}

	result := service.GroupBy(df, request)
	if result.Error != nil {
		t.Fatalf("GroupBy failed: %v", result.Error)
	}

	if result.DataFrame == nil {
		t.Fatal("GroupBy result DataFrame should not be nil")
	}
	defer result.DataFrame.Release()

	// Verify results
	if result.DataFrame.NumRows() != 2 { // 2 departments
		t.Errorf("Expected 2 rows, got %d", result.DataFrame.NumRows())
	}

	if result.DataFrame.NumCols() != 3 { // department + 2 aggregations
		t.Errorf("Expected 3 columns, got %d", result.DataFrame.NumCols())
	}

	// Test with invalid request - no group columns
	invalidRequest := aggregation.GroupByRequest{
		GroupColumns: []string{},
		Aggregations: []aggregation.AggregationSpec{
			{Column: "salary", Type: aggregation.Sum, Alias: "total_salary"},
		},
	}

	invalidResult := service.GroupBy(df, invalidRequest)
	if invalidResult.Error == nil {
		t.Error("GroupBy with no group columns should return error")
	}
}

func TestGroupByBuilder_NewGroupByBuilder(t *testing.T) {
	builder := NewGroupByBuilder("category")
	if builder == nil {
		t.Fatal("NewGroupByBuilder should not return nil")
	}

	if len(builder.groupColumns) != 1 {
		t.Errorf("Expected 1 group column, got %d", len(builder.groupColumns))
	}

	if builder.groupColumns[0] != "category" {
		t.Errorf("Expected group column 'category', got %s", builder.groupColumns[0])
	}

	if len(builder.aggregations) != 0 {
		t.Error("New GroupByBuilder should have empty aggregations")
	}
}

func TestGroupByBuilder_Sum(t *testing.T) {
	builder := NewGroupByBuilder("category")

	// Test Sum aggregation
	result := builder.Sum("value")
	if result != builder {
		t.Error("Sum should return the same builder for chaining")
	}

	if len(builder.aggregations) != 1 {
		t.Errorf("Expected 1 aggregation, got %d", len(builder.aggregations))
	}

	agg := builder.aggregations[0]
	if agg.Column != "value" {
		t.Errorf("Expected column 'value', got %s", agg.Column)
	}

	if agg.Type != aggregation.Sum {
		t.Errorf("Expected Sum aggregation type, got %v", agg.Type)
	}

	if agg.Alias != "value_sum" {
		t.Errorf("Expected alias 'value_sum', got %s", agg.Alias)
	}
}

func TestGroupByBuilder_Mean(t *testing.T) {
	builder := NewGroupByBuilder("category")

	// Test Mean aggregation
	result := builder.Mean("value")
	if result != builder {
		t.Error("Mean should return the same builder for chaining")
	}

	if len(builder.aggregations) != 1 {
		t.Errorf("Expected 1 aggregation, got %d", len(builder.aggregations))
	}

	agg := builder.aggregations[0]
	if agg.Column != "value" {
		t.Errorf("Expected column 'value', got %s", agg.Column)
	}

	if agg.Type != aggregation.Mean {
		t.Errorf("Expected Mean aggregation type, got %v", agg.Type)
	}

	if agg.Alias != "value_mean" {
		t.Errorf("Expected alias 'value_mean', got %s", agg.Alias)
	}
}

func TestGroupByBuilder_Count(t *testing.T) {
	builder := NewGroupByBuilder("category")

	// Test Count aggregation
	result := builder.Count("value")
	if result != builder {
		t.Error("Count should return the same builder for chaining")
	}

	if len(builder.aggregations) != 1 {
		t.Errorf("Expected 1 aggregation, got %d", len(builder.aggregations))
	}

	agg := builder.aggregations[0]
	if agg.Column != "value" {
		t.Errorf("Expected column 'value', got %s", agg.Column)
	}

	if agg.Type != aggregation.Count {
		t.Errorf("Expected Count aggregation type, got %v", agg.Type)
	}

	if agg.Alias != "value_count" {
		t.Errorf("Expected alias 'value_count', got %s", agg.Alias)
	}
}

func TestGroupByBuilder_Min(t *testing.T) {
	builder := NewGroupByBuilder("category")

	// Test Min aggregation
	result := builder.Min("value")
	if result != builder {
		t.Error("Min should return the same builder for chaining")
	}

	if len(builder.aggregations) != 1 {
		t.Errorf("Expected 1 aggregation, got %d", len(builder.aggregations))
	}

	agg := builder.aggregations[0]
	if agg.Column != "value" {
		t.Errorf("Expected column 'value', got %s", agg.Column)
	}

	if agg.Type != aggregation.Min {
		t.Errorf("Expected Min aggregation type, got %v", agg.Type)
	}

	if agg.Alias != "value_min" {
		t.Errorf("Expected alias 'value_min', got %s", agg.Alias)
	}
}

func TestGroupByBuilder_Max(t *testing.T) {
	builder := NewGroupByBuilder("category")

	// Test Max aggregation
	result := builder.Max("value")
	if result != builder {
		t.Error("Max should return the same builder for chaining")
	}

	if len(builder.aggregations) != 1 {
		t.Errorf("Expected 1 aggregation, got %d", len(builder.aggregations))
	}

	agg := builder.aggregations[0]
	if agg.Column != "value" {
		t.Errorf("Expected column 'value', got %s", agg.Column)
	}

	if agg.Type != aggregation.Max {
		t.Errorf("Expected Max aggregation type, got %v", agg.Type)
	}

	if agg.Alias != "value_max" {
		t.Errorf("Expected alias 'value_max', got %s", agg.Alias)
	}
}

func TestGroupByBuilder_Build(t *testing.T) {
	// Test Build functionality
	builder := NewGroupByBuilder("category")
	request := builder.Sum("value").As("total_value").Build()

	if len(request.GroupColumns) != 1 {
		t.Errorf("Expected 1 group column, got %d", len(request.GroupColumns))
	}

	if request.GroupColumns[0] != "category" {
		t.Errorf("Expected group column 'category', got %s", request.GroupColumns[0])
	}

	if len(request.Aggregations) != 1 {
		t.Errorf("Expected 1 aggregation, got %d", len(request.Aggregations))
	}

	agg := request.Aggregations[0]
	if agg.Column != "value" {
		t.Errorf("Expected column 'value', got %s", agg.Column)
	}

	if agg.Type != aggregation.Sum {
		t.Errorf("Expected Sum aggregation type, got %v", agg.Type)
	}

	if agg.Alias != "total_value" {
		t.Errorf("Expected alias 'total_value', got %s", agg.Alias)
	}
}
