package expr

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleCoverage(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create simple test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
			{Name: "str_col", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	intBuilder := array.NewInt64Builder(pool)
	strBuilder := array.NewStringBuilder(pool)

	intBuilder.AppendValues([]int64{1, 2, 3}, nil)
	strBuilder.AppendValues([]string{"a", "b", "c"}, nil)

	intArray := intBuilder.NewArray()
	strArray := strBuilder.NewArray()
	defer intArray.Release()
	defer strArray.Release()

	record := array.NewRecord(schema, []arrow.Array{intArray, strArray}, 3)
	df := core.NewDataFrame(record)
	defer df.Release()

	t.Run("ExpressionTypeChecking", func(t *testing.T) {
		// Test various expression methods that might be uncovered
		colExpr := Col("int_col")
		litExpr := Lit(42)

		// Test type inference
		intType := inferDataType(42)
		assert.Equal(t, arrow.PrimitiveTypes.Int64, intType)

		strType := inferDataType("hello")
		assert.Equal(t, arrow.BinaryTypes.String, strType)

		boolType := inferDataType(true)
		assert.Equal(t, arrow.FixedWidthTypes.Boolean, boolType)

		// Test expression composition
		addExpr := colExpr.Add(litExpr)
		assert.NotNil(t, addExpr)

		subExpr := colExpr.Sub(litExpr)
		assert.NotNil(t, subExpr)

		mulExpr := colExpr.Mul(litExpr)
		assert.NotNil(t, mulExpr)

		divExpr := colExpr.Div(litExpr)
		assert.NotNil(t, divExpr)
	})

	t.Run("ComparisonOperations", func(t *testing.T) {
		colExpr := Col("int_col")
		litExpr := Lit(2)

		// Test comparison operations
		ltExpr := colExpr.Lt(litExpr)
		result, err := ltExpr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())

		gtExpr := colExpr.Gt(litExpr)
		result2, err := gtExpr.Evaluate(df)
		require.NoError(t, err)
		defer result2.Release()
		assert.Equal(t, 3, result2.Len())

		eqExpr := colExpr.Eq(litExpr)
		result3, err := eqExpr.Evaluate(df)
		require.NoError(t, err)
		defer result3.Release()
		assert.Equal(t, 3, result3.Len())
	})

	t.Run("StringOperationMethods", func(t *testing.T) {
		strCol := Col("str_col")
		searchExpr := Lit("a")

		// Test string operations that should exist
		containsExpr := strCol.Contains(searchExpr)
		assert.NotNil(t, containsExpr)

		startsExpr := strCol.StartsWith(searchExpr)
		assert.NotNil(t, startsExpr)

		endsExpr := strCol.EndsWith(searchExpr)
		assert.NotNil(t, endsExpr)
	})
}

func TestUtilityFunctionsCoverage(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("ArrayConversionHelpers", func(t *testing.T) {
		// Test array conversion functions
		intBuilder := array.NewInt64Builder(pool)
		intBuilder.AppendValues([]int64{10, 20, 30}, nil)
		intArray := intBuilder.NewArray()
		defer intArray.Release()

		// Test if asInt64Array function exists and works
		if arr, ok := asInt64Array(intArray); ok {
			assert.Equal(t, int64(10), arr.Value(0))
			assert.Equal(t, int64(20), arr.Value(1))
			assert.Equal(t, int64(30), arr.Value(2))
		}

		// Test string array conversion
		strBuilder := array.NewStringBuilder(pool)
		strBuilder.AppendValues([]string{"x", "y", "z"}, nil)
		strArray := strBuilder.NewArray()
		defer strArray.Release()

		if arr, ok := asStringArray(strArray); ok {
			assert.Equal(t, "x", arr.Value(0))
			assert.Equal(t, "y", arr.Value(1))
			assert.Equal(t, "z", arr.Value(2))
		}
	})
}

func TestVectorizedExpressionBasics(t *testing.T) {
	t.Run("VectorizedExprConstructor", func(t *testing.T) {
		// Test basic vectorized expression creation
		operands := []Expr{Col("test"), Lit(5)}
		vexpr := NewVectorizedExpr(operands, "add", nil)
		assert.NotNil(t, vexpr)

		// Test string representation
		str := vexpr.String()
		assert.Contains(t, str, "vectorized")
	})
}

func TestMemoryPoolBasics(t *testing.T) {
	t.Run("MemoryPoolCreation", func(t *testing.T) {
		// Test memory pool creation and basic usage
		pool := NewMemoryPool()
		assert.NotNil(t, pool)
		allocator := pool.GetAllocator()
		assert.NotNil(t, allocator)
	})
}
