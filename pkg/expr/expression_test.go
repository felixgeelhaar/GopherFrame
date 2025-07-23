package expr

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
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
		{int(42), arrow.INT64}, // int promotes to int64
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

// TestBinaryExpressionAdd tests addition operations with TDD
func TestBinaryExpressionAdd(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []interface{}
	}{
		{
			name:     "Add two float64 columns",
			left:     Col("score"),
			right:    Col("score"),
			expected: []interface{}{191.0, 174.4, 184.2}, // score + score
		},
		{
			name:     "Add float64 column to literal",
			left:     Col("score"),
			right:    Lit(5.0),
			expected: []interface{}{100.5, 92.2, 97.1}, // score + 5.0
		},
		{
			name:     "Add int64 column to literal",
			left:     Col("id"),
			right:    Lit(int64(100)),
			expected: []interface{}{int64(101), int64(102), int64(103)}, // id + 100
		},
		{
			name:     "Add two int64 expressions",
			left:     Col("id"),
			right:    Col("id"),
			expected: []interface{}{int64(2), int64(4), int64(6)}, // id + id
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			addExpr := NewBinaryExpr(tc.left, tc.right, "add")

			result, err := addExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Add operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			if result.Len() != len(tc.expected) {
				t.Fatalf("Expected %d results, got %d", len(tc.expected), result.Len())
			}

			for i, expected := range tc.expected {
				switch expectedVal := expected.(type) {
				case float64:
					floatArray := result.(*array.Float64)
					actual := floatArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %.1f, got %.1f", i, expectedVal, actual)
					}
				case int64:
					intArray := result.(*array.Int64)
					actual := intArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %d, got %d", i, expectedVal, actual)
					}
				}
			}
		})
	}
}

// TestBinaryExpressionSubtract tests subtraction operations with TDD
func TestBinaryExpressionSubtract(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []interface{}
	}{
		{
			name:     "Subtract literal from float64 column",
			left:     Col("score"),
			right:    Lit(5.0),
			expected: []interface{}{90.5, 82.2, 87.1}, // score - 5.0
		},
		{
			name:     "Subtract int64 from int64",
			left:     Col("id"),
			right:    Lit(int64(1)),
			expected: []interface{}{int64(0), int64(1), int64(2)}, // id - 1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			subExpr := NewBinaryExpr(tc.left, tc.right, "subtract")

			result, err := subExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Subtract operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			for i, expected := range tc.expected {
				switch expectedVal := expected.(type) {
				case float64:
					floatArray := result.(*array.Float64)
					actual := floatArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %.1f, got %.1f", i, expectedVal, actual)
					}
				case int64:
					intArray := result.(*array.Int64)
					actual := intArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %d, got %d", i, expectedVal, actual)
					}
				}
			}
		})
	}
}

// TestBinaryExpressionLessThan tests less-than comparison operations with TDD
func TestBinaryExpressionLessThan(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []bool
	}{
		{
			name:     "Float64 column less than literal",
			left:     Col("score"),
			right:    Lit(90.0),
			expected: []bool{false, true, false}, // [95.5, 87.2, 92.1] < 90.0
		},
		{
			name:     "Int64 column less than literal",
			left:     Col("id"),
			right:    Lit(int64(3)),
			expected: []bool{true, true, false}, // [1, 2, 3] < 3
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ltExpr := NewBinaryExpr(tc.left, tc.right, "less")

			result, err := ltExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Less than operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			boolArray := result.(*array.Boolean)
			if boolArray.Len() != len(tc.expected) {
				t.Fatalf("Expected %d results, got %d", len(tc.expected), boolArray.Len())
			}

			for i, expected := range tc.expected {
				actual := boolArray.Value(i)
				if actual != expected {
					t.Errorf("Row %d: expected %v, got %v", i, expected, actual)
				}
			}
		})
	}
}

// TestBinaryExpressionEqual tests equality comparison operations with TDD
func TestBinaryExpressionEqual(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []bool
	}{
		{
			name:     "String column equals literal",
			left:     Col("name"),
			right:    Lit("Bob"),
			expected: []bool{false, true, false}, // ["Alice", "Bob", "Charlie"] == "Bob"
		},
		{
			name:     "Int64 column equals literal",
			left:     Col("id"),
			right:    Lit(int64(2)),
			expected: []bool{false, true, false}, // [1, 2, 3] == 2
		},
		{
			name:     "Float64 column equals literal",
			left:     Col("score"),
			right:    Lit(87.2),
			expected: []bool{false, true, false}, // [95.5, 87.2, 92.1] == 87.2
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eqExpr := NewBinaryExpr(tc.left, tc.right, "equal")

			result, err := eqExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Equal operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			boolArray := result.(*array.Boolean)
			if boolArray.Len() != len(tc.expected) {
				t.Fatalf("Expected %d results, got %d", len(tc.expected), boolArray.Len())
			}

			for i, expected := range tc.expected {
				actual := boolArray.Value(i)
				if actual != expected {
					t.Errorf("Row %d: expected %v, got %v", i, expected, actual)
				}
			}
		})
	}
}

// TestBinaryExpressionDivide tests division operations with TDD
func TestBinaryExpressionDivide(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []interface{}
	}{
		{
			name:     "Divide float64 column by literal",
			left:     Col("score"),
			right:    Lit(2.0),
			expected: []interface{}{47.75, 43.6, 46.05}, // score / 2.0
		},
		{
			name:     "Divide int64 column by literal",
			left:     Col("id"),
			right:    Lit(int64(2)),
			expected: []interface{}{int64(0), int64(1), int64(1)}, // id / 2 (integer division)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			divExpr := NewBinaryExpr(tc.left, tc.right, "divide")

			result, err := divExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Divide operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			for i, expected := range tc.expected {
				switch expectedVal := expected.(type) {
				case float64:
					floatArray := result.(*array.Float64)
					actual := floatArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %.2f, got %.2f", i, expectedVal, actual)
					}
				case int64:
					intArray := result.(*array.Int64)
					actual := intArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %d, got %d", i, expectedVal, actual)
					}
				}
			}
		})
	}
}

// TestBinaryExpressionMultiply tests multiplication operations with TDD
func TestBinaryExpressionMultiply(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []interface{}
	}{
		{
			name:     "Multiply float64 column by literal",
			left:     Col("score"),
			right:    Lit(2.0),
			expected: []interface{}{191.0, 174.4, 184.2}, // score * 2.0
		},
		{
			name:     "Multiply two float64 columns",
			left:     Col("score"),
			right:    Col("score"),
			expected: []interface{}{9120.25, 7603.84, 8482.41}, // score * score
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mulExpr := NewBinaryExpr(tc.left, tc.right, "multiply")

			result, err := mulExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Multiply operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			for i, expected := range tc.expected {
				switch expectedVal := expected.(type) {
				case float64:
					floatArray := result.(*array.Float64)
					actual := floatArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %.1f, got %.1f", i, expectedVal, actual)
					}
				case int64:
					intArray := result.(*array.Int64)
					actual := intArray.Value(i)
					if actual != expectedVal {
						t.Errorf("Row %d: expected %d, got %d", i, expectedVal, actual)
					}
				}
			}
		})
	}
}

// TestBinaryExpressionGreaterThan tests greater-than comparison operations with TDD
func TestBinaryExpressionGreaterThan(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	testCases := []struct {
		name     string
		left     Expr
		right    Expr
		expected []bool
	}{
		{
			name:     "Float64 column greater than literal",
			left:     Col("score"),
			right:    Lit(90.0),
			expected: []bool{true, false, true}, // [95.5, 87.2, 92.1] > 90.0
		},
		{
			name:     "Int64 column greater than literal",
			left:     Col("id"),
			right:    Lit(int64(2)),
			expected: []bool{false, false, true}, // [1, 2, 3] > 2
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gtExpr := NewBinaryExpr(tc.left, tc.right, "greater")

			result, err := gtExpr.Evaluate(df)
			if err != nil {
				t.Fatalf("Greater than operation failed: %v", err)
			}
			defer result.Release()

			// Verify results
			boolArray := result.(*array.Boolean)
			if boolArray.Len() != len(tc.expected) {
				t.Fatalf("Expected %d results, got %d", len(tc.expected), boolArray.Len())
			}

			for i, expected := range tc.expected {
				actual := boolArray.Value(i)
				if actual != expected {
					t.Errorf("Row %d: expected %v, got %v", i, expected, actual)
				}
			}
		})
	}
}

// TestExpressionStringMethods tests String() methods for all expression types
func TestExpressionStringMethods(t *testing.T) {
	testCases := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{
			name:     "Column expression string",
			expr:     Col("test_column"),
			expected: "Col(test_column)",
		},
		{
			name:     "Int64 literal string",
			expr:     Lit(int64(42)),
			expected: "Lit(42)",
		},
		{
			name:     "Float64 literal string",
			expr:     Lit(3.14),
			expected: "Lit(3.14)",
		},
		{
			name:     "String literal string",
			expr:     Lit("hello"),
			expected: "Lit(hello)",
		},
		{
			name:     "Boolean literal string",
			expr:     Lit(true),
			expected: "Lit(true)",
		},
		{
			name:     "Binary add expression string",
			expr:     NewBinaryExpr(Col("a"), Col("b"), "add"),
			expected: "(Col(a) add Col(b))",
		},
		{
			name:     "Binary subtract expression string",
			expr:     NewBinaryExpr(Col("x"), Lit(5), "subtract"),
			expected: "(Col(x) subtract Lit(5))",
		},
		{
			name:     "Binary multiply expression string",
			expr:     NewBinaryExpr(Col("y"), Lit(2.0), "multiply"),
			expected: "(Col(y) multiply Lit(2))",
		},
		{
			name:     "Binary divide expression string",
			expr:     NewBinaryExpr(Col("z"), Lit(3), "divide"),
			expected: "(Col(z) divide Lit(3))",
		},
		{
			name:     "Binary less than expression string",
			expr:     NewBinaryExpr(Col("score"), Lit(90), "less"),
			expected: "(Col(score) less Lit(90))",
		},
		{
			name:     "Binary greater than expression string",
			expr:     NewBinaryExpr(Col("age"), Lit(18), "greater"),
			expected: "(Col(age) greater Lit(18))",
		},
		{
			name:     "Binary equal expression string",
			expr:     NewBinaryExpr(Col("status"), Lit("active"), "equal"),
			expected: "(Col(status) equal Lit(active))",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.expr.String()
			if actual != tc.expected {
				t.Errorf("Expected string '%s', got '%s'", tc.expected, actual)
			}
		})
	}
}

// TestExpressionConvenienceMethods tests the convenience methods (Add, Sub, etc.)
func TestExpressionConvenienceMethods(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	col := Col("score")
	literal := Lit(10.0)

	// Test Add method
	addExpr := col.Add(literal)
	addResult, err := addExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Add method failed: %v", err)
	}
	if addResult != nil {
		defer addResult.Release()
	}

	// Test Sub method
	subExpr := col.Sub(literal)
	subResult, err := subExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Sub method failed: %v", err)
	}
	if subResult != nil {
		defer subResult.Release()
	}

	// Test Mul method
	mulExpr := col.Mul(literal)
	mulResult, err := mulExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Mul method failed: %v", err)
	}
	if mulResult != nil {
		defer mulResult.Release()
	}

	// Test Div method
	divExpr := col.Div(literal)
	divResult, err := divExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Div method failed: %v", err)
	}
	if divResult != nil {
		defer divResult.Release()
	}

	// Test Lt method
	ltExpr := col.Lt(literal)
	ltResult, err := ltExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Lt method failed: %v", err)
	}
	if ltResult != nil {
		defer ltResult.Release()
	}

	// Test Gt method
	gtExpr := col.Gt(literal)
	gtResult, err := gtExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Gt method failed: %v", err)
	}
	if gtResult != nil {
		defer gtResult.Release()
	}

	// Test Eq method
	eqExpr := col.Eq(literal)
	eqResult, err := eqExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Eq method failed: %v", err)
	}
	if eqResult != nil {
		defer eqResult.Release()
	}
}

// TestUnsupportedBinaryOperation tests error handling for unsupported operations
func TestUnsupportedBinaryOperation(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	// Test unsupported operation
	unsupportedExpr := NewBinaryExpr(Col("id"), Lit(5), "unsupported_op")
	_, err := unsupportedExpr.Evaluate(df)

	if err == nil {
		t.Error("Expected error for unsupported binary operation")
	}
}

// TestLiteralExprConvenienceMethods tests the convenience methods on LiteralExpr
func TestLiteralExprConvenienceMethods(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	literal := Lit(10.0)
	col := Col("score")

	// Test all convenience methods on literal expressions
	addExpr := literal.Add(col)
	addResult, err := addExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Add method failed: %v", err)
	}
	if addResult != nil {
		defer addResult.Release()
	}

	subExpr := literal.Sub(col)
	subResult, err := subExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Sub method failed: %v", err)
	}
	if subResult != nil {
		defer subResult.Release()
	}

	mulExpr := literal.Mul(col)
	mulResult, err := mulExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Mul method failed: %v", err)
	}
	if mulResult != nil {
		defer mulResult.Release()
	}

	divExpr := literal.Div(col)
	divResult, err := divExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Div method failed: %v", err)
	}
	if divResult != nil {
		defer divResult.Release()
	}

	ltExpr := literal.Lt(col)
	ltResult, err := ltExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Lt method failed: %v", err)
	}
	if ltResult != nil {
		defer ltResult.Release()
	}

	gtExpr := literal.Gt(col)
	gtResult, err := gtExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Gt method failed: %v", err)
	}
	if gtResult != nil {
		defer gtResult.Release()
	}

	eqExpr := literal.Eq(col)
	eqResult, err := eqExpr.Evaluate(df)
	if err != nil {
		t.Errorf("Literal Eq method failed: %v", err)
	}
	if eqResult != nil {
		defer eqResult.Release()
	}
}

// TestBinaryExprConvenienceMethods tests the convenience methods on BinaryExpr
func TestBinaryExprConvenienceMethods(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	binaryExpr := NewBinaryExpr(Col("score"), Lit(10.0), "add")
	col := Col("score") // Use same type to avoid mismatch errors

	// Test BinaryExpr Name method
	expectedName := "(score add Lit(10))"
	actualName := binaryExpr.Name()
	if actualName != expectedName {
		t.Errorf("Expected binary expr name '%s', got '%s'", expectedName, actualName)
	}

	// Test all convenience methods on binary expressions
	addExpr := binaryExpr.Add(col)
	addResult, err := addExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Add method failed: %v", err)
	}
	if addResult != nil {
		defer addResult.Release()
	}

	subExpr := binaryExpr.Sub(col)
	subResult, err := subExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Sub method failed: %v", err)
	}
	if subResult != nil {
		defer subResult.Release()
	}

	mulExpr := binaryExpr.Mul(col)
	mulResult, err := mulExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Mul method failed: %v", err)
	}
	if mulResult != nil {
		defer mulResult.Release()
	}

	divExpr := binaryExpr.Div(col)
	divResult, err := divExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Div method failed: %v", err)
	}
	if divResult != nil {
		defer divResult.Release()
	}

	ltExpr := binaryExpr.Lt(col)
	ltResult, err := ltExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Lt method failed: %v", err)
	}
	if ltResult != nil {
		defer ltResult.Release()
	}

	gtExpr := binaryExpr.Gt(col)
	gtResult, err := gtExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Gt method failed: %v", err)
	}
	if gtResult != nil {
		defer gtResult.Release()
	}

	eqExpr := binaryExpr.Eq(col)
	eqResult, err := eqExpr.Evaluate(df)
	if err != nil {
		t.Errorf("BinaryExpr Eq method failed: %v", err)
	}
	if eqResult != nil {
		defer eqResult.Release()
	}
}

// TestInferDataTypeEdgeCases tests additional data type inference cases
func TestInferDataTypeEdgeCases(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected arrow.Type
	}{
		{int32(42), arrow.INT64},        // int32 should promote to int64
		{float32(3.14), arrow.FLOAT32},  // float32 stays as float32
		{[]byte("hello"), arrow.STRING}, // byte slice becomes string
	}

	for _, tc := range testCases {
		result := inferDataType(tc.value)
		if result.ID() != tc.expected {
			t.Errorf("For value %v (%T), expected type %s, got %s",
				tc.value, tc.value, tc.expected, result.ID())
		}
	}

	// Test unsupported type (should default to string)
	unsupportedValue := make(chan int)
	result := inferDataType(unsupportedValue)
	if result.ID() != arrow.STRING {
		t.Errorf("Expected unsupported type to default to STRING, got %s", result.ID())
	}
}

// TestTypeMismatchErrors tests error cases for type mismatches
func TestTypeMismatchErrors(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	// Test string operations on incompatible types - should cause error paths in evaluation
	stringExpr := NewBinaryExpr(Col("name"), Col("score"), "add") // string + float64
	_, err := stringExpr.Evaluate(df)
	if err == nil {
		t.Error("Expected error for string + float64 operation")
	}

	// Test unsupported data type combinations
	unsupportedMul := NewBinaryExpr(Col("name"), Col("name"), "multiply") // string * string
	_, err = unsupportedMul.Evaluate(df)
	if err == nil {
		t.Error("Expected error for string * string operation")
	}
}

// TestLiteralExprWithComplexType tests literal expressions with more data types
func TestLiteralExprWithComplexType(t *testing.T) {
	df := createTestDataFrame(t)
	defer df.Release()

	// Test with int type (not int64)
	intLit := Lit(42) // This creates int, not int64
	result, err := intLit.Evaluate(df)
	if err != nil {
		t.Errorf("Failed to evaluate int literal: %v", err)
	}
	if result != nil {
		defer result.Release()
		// Verify it was converted to int64
		if result.DataType().ID() != arrow.INT64 {
			t.Errorf("Expected int to be converted to INT64, got %s", result.DataType().ID())
		}
	}

	// Test with unsupported type - this should cause an error
	// since float32 is not supported in the current implementation
	float32Lit := Lit(float32(3.14))
	_, err = float32Lit.Evaluate(df)
	if err == nil {
		t.Error("Expected error for unsupported float32 literal type")
	}
}

// Helper function to create a test DataFrame
func createTestDataFrame(_ *testing.T) *core.DataFrame {
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
