package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestGroupByBasicAggregation(t *testing.T) {
	// Create test data with groups
	df := createGroupedDataFrame()
	defer df.Release()

	// Test GroupBy with sum aggregation
	result := df.GroupBy("category").Agg(Sum("amount"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("GroupBy aggregation failed: %v", result.Err())
	}

	// Should have one row per unique category
	expectedRows := int64(3) // A, B, C categories
	if result.NumRows() != expectedRows {
		t.Errorf("Expected %d groups, got %d", expectedRows, result.NumRows())
	}

	// Should have category and sum columns
	expectedCols := int64(2) // category, amount_sum
	if result.NumCols() != expectedCols {
		t.Errorf("Expected %d columns, got %d", expectedCols, result.NumCols())
	}

	// Check column names
	names := result.ColumnNames()
	if names[0] != "category" {
		t.Errorf("Expected first column 'category', got '%s'", names[0])
	}
	if names[1] != "amount_sum" {
		t.Errorf("Expected second column 'amount_sum', got '%s'", names[1])
	}
}

func TestGroupByMultipleAggregations(t *testing.T) {
	df := createGroupedDataFrame()
	defer df.Release()

	// Test multiple aggregations
	result := df.GroupBy("category").Agg(
		Sum("amount"),
		Mean("amount"),
		Count("amount"),
		Min("amount"),
		Max("amount"),
	)
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("Multiple aggregations failed: %v", result.Err())
	}

	// Should have category + 5 aggregation columns
	expectedCols := int64(6)
	if result.NumCols() != expectedCols {
		t.Errorf("Expected %d columns, got %d", expectedCols, result.NumCols())
	}

	// Check aggregation column names
	names := result.ColumnNames()
	expectedNames := []string{"category", "amount_sum", "amount_mean", "amount_count", "amount_min", "amount_max"}

	for i, expected := range expectedNames {
		if names[i] != expected {
			t.Errorf("Expected column %d to be '%s', got '%s'", i, expected, names[i])
		}
	}
}

func TestGroupByMultipleColumns(t *testing.T) {
	df := createComplexGroupedDataFrame()
	defer df.Release()

	// Test grouping by multiple columns
	result := df.GroupBy("department", "level").Agg(Sum("salary"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("Multi-column GroupBy failed: %v", result.Err())
	}

	// Should have department, level, and salary_sum columns
	expectedCols := int64(3)
	if result.NumCols() != expectedCols {
		t.Errorf("Expected %d columns, got %d", expectedCols, result.NumCols())
	}

	names := result.ColumnNames()
	expectedNames := []string{"department", "level", "salary_sum"}

	for i, expected := range expectedNames {
		if names[i] != expected {
			t.Errorf("Expected column %d to be '%s', got '%s'", i, expected, names[i])
		}
	}
}

func TestGroupByChaining(t *testing.T) {
	df := createGroupedDataFrame()
	defer df.Release()

	// Test chaining GroupBy with other operations
	result := df.
		Filter(Col("amount").Gt(Lit(10.0))).
		GroupBy("category").
		Agg(Mean("amount")).
		Select("category", "amount_mean")
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("GroupBy chaining failed: %v", result.Err())
	}

	// Should have filtered and aggregated data
	if result.NumCols() != 2 {
		t.Errorf("Expected 2 columns after chaining, got %d", result.NumCols())
	}
}

func TestAggregationFunctions(t *testing.T) {
	// Test individual aggregation function builders

	// Test Sum
	sumAgg := Sum("amount")
	if sumAgg.Name() != "amount_sum" {
		t.Errorf("Sum aggregation name should be 'amount_sum', got '%s'", sumAgg.Name())
	}

	// Test Mean
	meanAgg := Mean("score")
	if meanAgg.Name() != "score_mean" {
		t.Errorf("Mean aggregation name should be 'score_mean', got '%s'", meanAgg.Name())
	}

	// Test Count
	countAgg := Count("id")
	if countAgg.Name() != "id_count" {
		t.Errorf("Count aggregation name should be 'id_count', got '%s'", countAgg.Name())
	}

	// Test custom named aggregation
	customAgg := Sum("revenue").As("total_revenue")
	if customAgg.Name() != "total_revenue" {
		t.Errorf("Custom named aggregation should be 'total_revenue', got '%s'", customAgg.Name())
	}
}

// Helper function to create grouped test data
func createGroupedDataFrame() *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Create ID column: [1, 2, 3, 4, 5, 6]
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3, 4, 5, 6}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	// Create category column: [A, A, B, B, C, C]
	categoryBuilder := array.NewStringBuilder(pool)
	categoryBuilder.AppendValues([]string{"A", "A", "B", "B", "C", "C"}, nil)
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	// Create amount column: [10.0, 20.0, 15.0, 25.0, 5.0, 35.0]
	amountBuilder := array.NewFloat64Builder(pool)
	amountBuilder.AppendValues([]float64{10.0, 20.0, 15.0, 25.0, 5.0, 35.0}, nil)
	amountArray := amountBuilder.NewArray()
	defer amountArray.Release()

	// Create record
	record := array.NewRecord(schema, []arrow.Array{idArray, categoryArray, amountArray}, 6)
	defer record.Release()

	return NewDataFrame(record)
}

// Helper function to create complex grouped test data
func createComplexGroupedDataFrame() *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "level", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Create columns for multi-level grouping
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3, 4, 5, 6}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"Engineering", "Engineering", "Sales", "Sales", "Engineering", "Sales"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	levelBuilder := array.NewStringBuilder(pool)
	levelBuilder.AppendValues([]string{"Senior", "Junior", "Senior", "Junior", "Senior", "Senior"}, nil)
	levelArray := levelBuilder.NewArray()
	defer levelArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryBuilder.AppendValues([]float64{120000, 80000, 100000, 70000, 130000, 110000}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, deptArray, levelArray, salaryArray}, 6)
	defer record.Release()

	return NewDataFrame(record)
}
