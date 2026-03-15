package gopherframe

import (
	"math"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestVarianceAggregation(t *testing.T) {
	// Known values: [2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0]
	// Mean = 5.0, Sample Variance = 4.571428... (N-1 = 7)
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	catBuilder := array.NewStringBuilder(pool)
	catBuilder.AppendValues([]string{"A", "A", "A", "A", "A", "A", "A", "A"}, nil)
	catArray := catBuilder.NewArray()
	defer catArray.Release()

	valBuilder := array.NewFloat64Builder(pool)
	valBuilder.AppendValues([]float64{2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0}, nil)
	valArray := valBuilder.NewArray()
	defer valArray.Release()

	record := array.NewRecord(schema, []arrow.Array{catArray, valArray}, 8)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	result := df.GroupBy("category").Agg(Variance("value"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("Variance aggregation failed: %v", result.Err())
	}

	if result.NumRows() != 1 {
		t.Fatalf("Expected 1 row, got %d", result.NumRows())
	}

	names := result.ColumnNames()
	if names[1] != "value_variance" {
		t.Errorf("Expected column name 'value_variance', got '%s'", names[1])
	}

	// Sample variance of [2,4,4,4,5,5,7,9]: mean=5.0
	// Sum of squared deviations = 9+1+1+1+0+0+4+16 = 32
	// Sample variance = 32 / 7 = 4.571428571428571
	expected := 32.0 / 7.0
	actual := getFloat64Value(t, result, 1, 0)
	if math.Abs(actual-expected) > 1e-10 {
		t.Errorf("Expected variance %.10f, got %.10f", expected, actual)
	}
}

func TestStdDevAggregation(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	catBuilder := array.NewStringBuilder(pool)
	catBuilder.AppendValues([]string{"A", "A", "A", "A", "A", "A", "A", "A"}, nil)
	catArray := catBuilder.NewArray()
	defer catArray.Release()

	valBuilder := array.NewFloat64Builder(pool)
	valBuilder.AppendValues([]float64{2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0}, nil)
	valArray := valBuilder.NewArray()
	defer valArray.Release()

	record := array.NewRecord(schema, []arrow.Array{catArray, valArray}, 8)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	result := df.GroupBy("category").Agg(StdDev("value"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("StdDev aggregation failed: %v", result.Err())
	}

	if result.NumRows() != 1 {
		t.Fatalf("Expected 1 row, got %d", result.NumRows())
	}

	names := result.ColumnNames()
	if names[1] != "value_stddev" {
		t.Errorf("Expected column name 'value_stddev', got '%s'", names[1])
	}

	// StdDev = sqrt(32/7)
	expected := math.Sqrt(32.0 / 7.0)
	actual := getFloat64Value(t, result, 1, 0)
	if math.Abs(actual-expected) > 1e-10 {
		t.Errorf("Expected stddev %.10f, got %.10f", expected, actual)
	}
}

func TestVarianceWithNulls(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	catBuilder := array.NewStringBuilder(pool)
	catBuilder.AppendValues([]string{"A", "A", "A", "A", "A"}, nil)
	catArray := catBuilder.NewArray()
	defer catArray.Release()

	// Values: [10.0, null, 20.0, null, 30.0] -> non-null: [10, 20, 30]
	valBuilder := array.NewFloat64Builder(pool)
	valBuilder.Append(10.0)
	valBuilder.AppendNull()
	valBuilder.Append(20.0)
	valBuilder.AppendNull()
	valBuilder.Append(30.0)
	valArray := valBuilder.NewArray()
	defer valArray.Release()

	record := array.NewRecord(schema, []arrow.Array{catArray, valArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	result := df.GroupBy("category").Agg(Variance("value"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("Variance with nulls failed: %v", result.Err())
	}

	// Non-null values: [10, 20, 30], mean=20
	// Sum of squared deviations = 100 + 0 + 100 = 200
	// Sample variance = 200 / 2 = 100.0
	expected := 100.0
	actual := getFloat64Value(t, result, 1, 0)
	if math.Abs(actual-expected) > 1e-10 {
		t.Errorf("Expected variance %.10f, got %.10f", expected, actual)
	}
}

func TestVarianceSingleValue(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Group "A" has 2 values, group "B" has 1 value, group "C" has 0 non-null values
	catBuilder := array.NewStringBuilder(pool)
	catBuilder.AppendValues([]string{"A", "A", "B", "C"}, nil)
	catArray := catBuilder.NewArray()
	defer catArray.Release()

	valBuilder := array.NewFloat64Builder(pool)
	valBuilder.Append(10.0)
	valBuilder.Append(20.0)
	valBuilder.Append(5.0)
	valBuilder.AppendNull()
	valArray := valBuilder.NewArray()
	defer valArray.Release()

	record := array.NewRecord(schema, []arrow.Array{catArray, valArray}, 4)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	result := df.GroupBy("category").Agg(Variance("value"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("Variance single value failed: %v", result.Err())
	}

	if result.NumRows() != 3 {
		t.Fatalf("Expected 3 groups, got %d", result.NumRows())
	}

	// Group A: [10, 20], mean=15, variance = (25+25)/1 = 50
	actualA := getFloat64Value(t, result, 1, 0)
	if math.Abs(actualA-50.0) > 1e-10 {
		t.Errorf("Expected group A variance 50.0, got %.10f", actualA)
	}

	// Group B: single value -> null
	if !isNullValue(t, result, 1, 1) {
		t.Error("Expected group B variance to be null (single value)")
	}

	// Group C: no non-null values -> null
	if !isNullValue(t, result, 1, 2) {
		t.Error("Expected group C variance to be null (no non-null values)")
	}
}

func TestVarianceGroupBy(t *testing.T) {
	// Full GroupBy().Agg(Variance()) test with multiple groups
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"Eng", "Eng", "Eng", "Sales", "Sales", "Sales"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryBuilder.AppendValues([]float64{100.0, 110.0, 120.0, 50.0, 70.0, 80.0}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, salaryArray}, 6)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	result := df.GroupBy("department").Agg(
		Mean("salary"),
		Variance("salary"),
	)
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("Variance GroupBy failed: %v", result.Err())
	}

	if result.NumRows() != 2 {
		t.Fatalf("Expected 2 groups, got %d", result.NumRows())
	}

	names := result.ColumnNames()
	expectedNames := []string{"department", "salary_mean", "salary_variance"}
	for i, expected := range expectedNames {
		if names[i] != expected {
			t.Errorf("Expected column %d to be '%s', got '%s'", i, expected, names[i])
		}
	}

	// Eng: [100, 110, 120], mean=110, variance = (100+0+100)/2 = 100
	engVariance := getFloat64Value(t, result, 2, 0)
	if math.Abs(engVariance-100.0) > 1e-10 {
		t.Errorf("Expected Eng variance 100.0, got %.10f", engVariance)
	}

	// Sales: [50, 70, 80], mean=200/3=66.666...,
	// deviations: (50-66.667)^2 + (70-66.667)^2 + (80-66.667)^2
	// = 277.778 + 11.111 + 177.778 = 466.667 (approx)
	// variance = 466.667 / 2 = 233.333...
	salesMean := (50.0 + 70.0 + 80.0) / 3.0
	salesSumSqDev := (50.0-salesMean)*(50.0-salesMean) + (70.0-salesMean)*(70.0-salesMean) + (80.0-salesMean)*(80.0-salesMean)
	expectedSalesVariance := salesSumSqDev / 2.0

	salesVariance := getFloat64Value(t, result, 2, 1)
	if math.Abs(salesVariance-expectedSalesVariance) > 1e-10 {
		t.Errorf("Expected Sales variance %.10f, got %.10f", expectedSalesVariance, salesVariance)
	}
}

func TestStdDevGroupBy(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"Eng", "Eng", "Eng", "Sales", "Sales", "Sales"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryBuilder.AppendValues([]float64{100.0, 110.0, 120.0, 50.0, 70.0, 80.0}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, salaryArray}, 6)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	result := df.GroupBy("department").Agg(StdDev("salary"))
	defer result.Release()

	if result.Err() != nil {
		t.Fatalf("StdDev GroupBy failed: %v", result.Err())
	}

	if result.NumRows() != 2 {
		t.Fatalf("Expected 2 groups, got %d", result.NumRows())
	}

	names := result.ColumnNames()
	if names[1] != "salary_stddev" {
		t.Errorf("Expected column name 'salary_stddev', got '%s'", names[1])
	}

	// Eng: variance=100, stddev=10
	engStdDev := getFloat64Value(t, result, 1, 0)
	if math.Abs(engStdDev-10.0) > 1e-10 {
		t.Errorf("Expected Eng stddev 10.0, got %.10f", engStdDev)
	}

	// Sales: stddev = sqrt(variance)
	salesMean := (50.0 + 70.0 + 80.0) / 3.0
	salesSumSqDev := (50.0-salesMean)*(50.0-salesMean) + (70.0-salesMean)*(70.0-salesMean) + (80.0-salesMean)*(80.0-salesMean)
	expectedSalesStdDev := math.Sqrt(salesSumSqDev / 2.0)

	salesStdDev := getFloat64Value(t, result, 1, 1)
	if math.Abs(salesStdDev-expectedSalesStdDev) > 1e-10 {
		t.Errorf("Expected Sales stddev %.10f, got %.10f", expectedSalesStdDev, salesStdDev)
	}
}

// getFloat64Value extracts a float64 value from a DataFrame result at the given column and row.
func getFloat64Value(t *testing.T, df *DataFrame, colIdx int, rowIdx int) float64 {
	t.Helper()
	rec := df.Record()
	col := rec.Column(colIdx)
	f64arr, ok := col.(*array.Float64)
	if !ok {
		t.Fatalf("Column %d is not float64", colIdx)
	}
	if f64arr.IsNull(rowIdx) {
		t.Fatalf("Value at column %d, row %d is null", colIdx, rowIdx)
	}
	return f64arr.Value(rowIdx)
}

// isNullValue checks if a value at the given column and row is null.
func isNullValue(t *testing.T, df *DataFrame, colIdx int, rowIdx int) bool {
	t.Helper()
	rec := df.Record()
	col := rec.Column(colIdx)
	return col.IsNull(rowIdx)
}
