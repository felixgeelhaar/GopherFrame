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

func TestExpressionErrors(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create minimal test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "num", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 2)
	df := core.NewDataFrame(record)
	defer df.Release()

	t.Run("ColumnNotFound", func(t *testing.T) {
		// Test accessing non-existent column
		expr := Col("nonexistent")
		_, err := expr.Evaluate(df)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "column not found")
	})

	t.Run("TypeInferenceEdgeCases", func(t *testing.T) {
		// Test type inference with various Go types
		types := []interface{}{
			int8(1), int16(2), int32(3), int64(4),
			uint8(5), uint16(6), uint32(7), uint64(8),
			float32(9.0), float64(10.0),
			"string", true, false,
		}

		for _, val := range types {
			dtype := inferDataType(val)
			assert.NotNil(t, dtype, "Type inference failed for %T", val)
		}
	})

	t.Run("ArrayHelperFunctions", func(t *testing.T) {
		// Test helper functions for coverage
		intBuilder := array.NewInt64Builder(pool)
		intBuilder.AppendValues([]int64{100, 200}, nil)
		intArray := intBuilder.NewArray()
		defer intArray.Release()

		// Test asInt64Array
		if arr, ok := asInt64Array(intArray); ok {
			assert.Equal(t, int64(100), arr.Value(0))
			assert.Equal(t, int64(200), arr.Value(1))
		}

		// Test with non-int64 array should return false
		strBuilder := array.NewStringBuilder(pool)
		strBuilder.AppendValues([]string{"a", "b"}, nil)
		strArray := strBuilder.NewArray()
		defer strArray.Release()

		if _, ok := asInt64Array(strArray); ok {
			t.Error("asInt64Array should return false for string array")
		}

		// Test asStringArray
		if arr, ok := asStringArray(strArray); ok {
			assert.Equal(t, "a", arr.Value(0))
			assert.Equal(t, "b", arr.Value(1))
		}

		// Test with non-string array should return false
		if _, ok := asStringArray(intArray); ok {
			t.Error("asStringArray should return false for int array")
		}
	})
}

func TestExpressionConstructors(t *testing.T) {
	t.Run("ExpressionCreation", func(t *testing.T) {
		// Test various expression constructors
		col := Col("test")
		assert.NotNil(t, col)
		assert.IsType(t, &ColumnExpr{}, col)

		lit := Lit(42)
		assert.NotNil(t, lit)
		assert.IsType(t, &LiteralExpr{}, lit)

		// Test binary expression creation
		binExpr := col.Add(lit)
		assert.NotNil(t, binExpr)
		assert.IsType(t, &BinaryExpr{}, binExpr)
	})

	t.Run("ExpressionStringOutput", func(t *testing.T) {
		// Test string representations for coverage
		expressions := []Expr{
			Col("column_name"),
			Lit(123),
			Lit("text"),
			Lit(true),
			Lit(3.14),
			Col("a").Add(Lit(1)),
			Col("b").Lt(Lit(10)),
			Col("c").Eq(Lit("value")),
		}

		for _, expr := range expressions {
			str := expr.String()
			assert.NotEmpty(t, str, "String() should not be empty for %T", expr)
		}
	})
}

func TestMemoryPoolExtended(t *testing.T) {
	t.Run("MemoryPoolOperations", func(t *testing.T) {
		pool := NewMemoryPool()
		assert.NotNil(t, pool)

		// Test getting allocator
		allocator := pool.GetAllocator()
		assert.NotNil(t, allocator)

		// Test multiple allocator calls return same instance
		allocator2 := pool.GetAllocator()
		assert.Equal(t, allocator, allocator2)
	})

	t.Run("PooledExpressionEvaluation", func(t *testing.T) {
		// Create test data for pooled evaluation
		memPool := memory.NewGoAllocator()
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "data", Type: arrow.PrimitiveTypes.Int64},
			},
			nil,
		)

		builder := array.NewInt64Builder(memPool)
		builder.AppendValues([]int64{1, 2, 3}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 3)
		df := core.NewDataFrame(record)
		defer df.Release()

		// Test expression evaluation with memory pool
		expr := Col("data").Add(Lit(10))
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})
}

func TestVectorizedExpressions(t *testing.T) {
	t.Run("VectorizedExpressionProperties", func(t *testing.T) {
		operands := []Expr{Col("x"), Col("y")}
		vexpr := NewVectorizedExpr(operands, "add", nil)

		// Test that vectorized expression implements Expr interface
		assert.NotNil(t, vexpr)

		// Test string representation
		str := vexpr.String()
		assert.Contains(t, str, "vectorized")
		assert.Contains(t, str, "add")

		// Test that it can be used in further operations
		combined := vexpr.Add(Lit(5))
		assert.NotNil(t, combined)
	})
}

func TestExpressionInterfaces(t *testing.T) {
	t.Run("ExprInterface", func(t *testing.T) {
		// Ensure all expression types implement Expr interface
		var exprs []Expr = []Expr{
			Col("test"),
			Lit(42),
			Col("a").Add(Lit(1)),
			NewVectorizedExpr([]Expr{Col("x")}, "test", nil),
		}

		for _, expr := range exprs {
			// Test that all implement basic Expr methods
			assert.NotEmpty(t, expr.String())

			// Test that they can be combined
			combined := expr.Add(Lit(1))
			assert.NotNil(t, combined)
		}
	})
}
