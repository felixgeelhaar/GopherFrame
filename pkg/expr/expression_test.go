package expr

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/gopherFrame/pkg/core"
)

func TestColumnExpression(t *testing.T) {
	// Create test DataFrame
	df := createTestDataFrame(t)
	defer df.Release()
	
	// Create column expression
	expr := Col("id")
	
	// Test name
	if expr.Name() != "id" {
		t.Errorf("Expected name 'id', got '%s'", expr.Name())
	}
	
	// Test evaluation
	result, err := expr.Evaluate(df)
	if err != nil {
		t.Errorf("Failed to evaluate column expression: %v", err)
	}
	defer result.Release()
	
	// Check result type
	if result.DataType().ID() != arrow.INT64 {
		t.Errorf("Expected INT64 result, got %s", result.DataType())
	}
	
	// Check result length matches DataFrame
	if result.Len() != int(df.NumRows()) {
		t.Errorf("Expected result length %d, got %d", df.NumRows(), result.Len())
	}
	
	// Check first value
	int64Array := result.(*array.Int64)
	if int64Array.Value(0) != 1 {
		t.Errorf("Expected first value 1, got %d", int64Array.Value(0))
	}
}

func TestLiteralExpression(t *testing.T) {
	// Create test DataFrame
	df := createTestDataFrame(t)
	defer df.Release()
	
	// Test int64 literal
	intExpr := Lit(int64(42))
	
	if intExpr.Name() != "Lit(42)" {
		t.Errorf("Expected name 'Lit(42)', got '%s'", intExpr.Name())
	}
	
	result, err := intExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Failed to evaluate int literal: %v", err)
	}
	defer result.Release()
	
	// Check that all values are the literal value
	int64Array := result.(*array.Int64)
	for i := 0; i < int(df.NumRows()); i++ {
		if int64Array.Value(i) != 42 {
			t.Errorf("Expected all values to be 42, got %d at index %d", int64Array.Value(i), i)
		}
	}
	
	// Test string literal
	strExpr := Lit("hello")
	strResult, err := strExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Failed to evaluate string literal: %v", err)
	}
	defer strResult.Release()
	
	stringArray := strResult.(*array.String)
	for i := 0; i < int(df.NumRows()); i++ {
		if stringArray.Value(i) != "hello" {
			t.Errorf("Expected all values to be 'hello', got '%s' at index %d", stringArray.Value(i), i)
		}
	}
	
	// Test boolean literal
	boolExpr := Lit(true)
	boolResult, err := boolExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Failed to evaluate boolean literal: %v", err)
	}
	defer boolResult.Release()
	
	boolArray := boolResult.(*array.Boolean)
	for i := 0; i < int(df.NumRows()); i++ {
		if !boolArray.Value(i) {
			t.Errorf("Expected all values to be true, got false at index %d", i)
		}
	}
}

func TestInferDataType(t *testing.T) {
	// Test type inference
	testCases := []struct {
		value    interface{}
		expected arrow.Type
	}{
		{true, arrow.BOOL},
		{int(42), arrow.INT64},  // int promotes to int64
		{int64(42), arrow.INT64},
		{float64(3.14), arrow.FLOAT64},
		{"hello", arrow.STRING},
	}
	
	for _, tc := range testCases {
		result := inferDataType(tc.value)
		if result.ID() != tc.expected {
			t.Errorf("For value %v (%T), expected type %s, got %s", 
				tc.value, tc.value, tc.expected, result.ID())
		}
	}
}

func TestNonExistentColumn(t *testing.T) {
	// Create test DataFrame
	df := createTestDataFrame(t)
	defer df.Release()
	
	// Try to access non-existent column
	expr := Col("nonexistent")
	_, err := expr.Evaluate(df)
	
	if err == nil {
		t.Error("Expected error for non-existent column")
	}
}

// Helper function to create a test DataFrame
func createTestDataFrame(t *testing.T) *core.DataFrame {
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
	
	return core.NewDataFrame(record)
}