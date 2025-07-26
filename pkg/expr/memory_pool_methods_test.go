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

func TestMemoryPoolMethods(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
			{Name: "date_col", Type: arrow.FixedWidthTypes.Date32},
		},
		nil,
	)

	intBuilder := array.NewInt64Builder(pool)
	dateBuilder := array.NewDate32Builder(pool)

	intBuilder.AppendValues([]int64{1, 2, 3}, nil)
	dateBuilder.AppendValues([]arrow.Date32{19523, 19524, 19525}, nil) // Test dates

	intArray := intBuilder.NewArray()
	dateArray := dateBuilder.NewArray()
	defer intArray.Release()
	defer dateArray.Release()

	record := array.NewRecord(schema, []arrow.Array{intArray, dateArray}, 3)
	df := core.NewDataFrame(record)
	defer df.Release()

	// Create allocator for pooled operations
	memAllocator := memory.NewGoAllocator()

	t.Run("ColumnExprEvaluateWithPool", func(t *testing.T) {
		col := Col("int_col")

		// Test EvaluateWithPool method (public API)
		result, err := col.EvaluateWithPool(df, memAllocator)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("LiteralExprEvaluateWithPool", func(t *testing.T) {
		lit := Lit(42)

		// Test EvaluateWithPool method (public API)
		result, err := lit.EvaluateWithPool(df, memAllocator)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("UnaryExprEvaluateWithPool", func(t *testing.T) {
		unary := Col("date_col").Year()

		// Test EvaluateWithPool method (public API)
		result, err := unary.EvaluateWithPool(df, memAllocator)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("BinaryExprEvaluateWithPool", func(t *testing.T) {
		binary := Col("int_col").Add(Lit(10))

		// Test EvaluateWithPool method (public API)
		result, err := binary.EvaluateWithPool(df, memAllocator)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("VectorizedExprEvaluateWithPool", func(t *testing.T) {
		vexpr := NewVectorizedExpr([]Expr{Col("int_col"), Lit(5)}, "add", nil)

		// Test EvaluateWithPool method (public API)
		result, err := vexpr.EvaluateWithPool(df, memAllocator)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})

	t.Run("GenericEvaluateWithPool", func(t *testing.T) {
		expr := Col("int_col")

		// Test the helper function evaluateWithPool
		result, err := evaluateWithPool(expr, df, memAllocator)
		require.NoError(t, err)
		defer result.Release()
		assert.Equal(t, 3, result.Len())
	})
}

func TestExprContext(t *testing.T) {
	t.Run("NewExprContext", func(t *testing.T) {
		ctx := NewExprContext()

		assert.NotNil(t, ctx)
		assert.NotNil(t, ctx.Pool)
	})

	t.Run("DefaultContext", func(t *testing.T) {
		// Test that DefaultContext is available
		assert.NotNil(t, DefaultContext)
		assert.NotNil(t, DefaultContext.Pool)
	})
}
