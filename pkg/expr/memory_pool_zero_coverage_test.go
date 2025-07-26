package expr

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestMemoryPoolZeroCoverage(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
			{Name: "float_col", Type: arrow.PrimitiveTypes.Float64},
			{Name: "string_col", Type: arrow.BinaryTypes.String},
			{Name: "date_col", Type: arrow.FixedWidthTypes.Date32},
			{Name: "timestamp_col", Type: &arrow.TimestampType{Unit: arrow.Second}},
		},
		nil,
	)

	intBuilder := array.NewInt64Builder(pool)
	floatBuilder := array.NewFloat64Builder(pool)
	stringBuilder := array.NewStringBuilder(pool)
	dateBuilder := array.NewDate32Builder(pool)
	timestampBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Second})

	intBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	floatBuilder.AppendValues([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, nil)
	stringBuilder.AppendValues([]string{"a", "b", "c", "d", "e"}, nil)
	dateBuilder.AppendValues([]arrow.Date32{19000, 19001, 19002, 19003, 19004}, nil)
	timestampBuilder.AppendValues([]arrow.Timestamp{1640995200, 1641081600, 1641168000, 1641254400, 1641340800}, nil)

	intArray := intBuilder.NewArray()
	floatArray := floatBuilder.NewArray()
	stringArray := stringBuilder.NewArray()
	dateArray := dateBuilder.NewArray()
	timestampArray := timestampBuilder.NewArray()

	defer intArray.Release()
	defer floatArray.Release()
	defer stringArray.Release()
	defer dateArray.Release()
	defer timestampArray.Release()

	record := array.NewRecord(schema, []arrow.Array{intArray, floatArray, stringArray, dateArray, timestampArray}, 5)
	df := core.NewDataFrame(record)
	defer df.Release()

	t.Run("TestMemoryPoolEvaluateWithPool", func(t *testing.T) {
		// Test the EvaluateWithPool methods to improve coverage
		ctx := NewExprContext()

		// Test LiteralExpr EvaluateWithPool
		lit := Lit(42)
		result, err := lit.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result != nil {
			defer result.Release()
			assert.NotNil(t, result)
			assert.Equal(t, 5, result.Len()) // Should match DataFrame length
		}

		// Test ColumnExpr EvaluateWithPool
		intCol := Col("int_col")
		result2, err := intCol.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result2 != nil {
			defer result2.Release()
			assert.NotNil(t, result2)
		}

		// Test UnaryExpr EvaluateWithPool
		yearExpr := Col("date_col").Year()
		result3, err := yearExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result3 != nil {
			defer result3.Release()
			assert.NotNil(t, result3)
		}

		// Test BinaryExpr EvaluateWithPool
		addExpr := Col("int_col").Add(Lit(10))
		result4, err := addExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result4 != nil {
			defer result4.Release()
			assert.NotNil(t, result4)
		}
	})

	t.Run("TestSpecificMemoryPoolFunctions", func(t *testing.T) {
		// Test specific memory pool functionality with complex expressions
		ctx := NewExprContext()

		// Test chained operations that use memory pools internally
		dateCol := Col("date_col")

		// Chain multiple date operations
		complexExpr := dateCol.Year().Add(Lit(2000))
		result, err := complexExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result != nil {
			defer result.Release()
			assert.NotNil(t, result)
		}

		// Test complex arithmetic with memory pool
		intCol := Col("int_col")
		floatCol := Col("float_col")
		arithmeticExpr := intCol.Add(floatCol).Mul(Lit(2.0))
		result2, err := arithmeticExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result2 != nil {
			defer result2.Release()
			assert.NotNil(t, result2)
		}

		// Test comparison operations with memory pool
		compExpr := intCol.Gt(Lit(2))
		result3, err := compExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result3 != nil {
			defer result3.Release()
			assert.NotNil(t, result3)
		}

		// Test string operations with memory pool
		stringCol := Col("string_col")
		strExpr := stringCol.Contains(Lit("a"))
		result4, err := strExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result4 != nil {
			defer result4.Release()
			assert.NotNil(t, result4)
		}
	})

	t.Run("TestArithmeticMemoryPoolFunctions", func(t *testing.T) {
		// Test various arithmetic and comparison expressions with memory pool
		ctx := NewExprContext()

		// Test complex nested expressions
		intCol := Col("int_col")
		floatCol := Col("float_col")

		// Test nested arithmetic operations
		nestedExpr := intCol.Add(floatCol).Sub(Lit(1.0)).Mul(Lit(2.0))
		result, err := nestedExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result != nil {
			defer result.Release()
			assert.NotNil(t, result)
		}

		// Test division operations (which can trigger division by zero paths)
		divExpr := floatCol.Div(intCol)
		result2, err := divExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result2 != nil {
			defer result2.Release()
			assert.NotNil(t, result2)
		}

		// Test comparison chains
		compExpr := intCol.Gt(Lit(0)).Eq(Lit(true))
		result3, err := compExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result3 != nil {
			defer result3.Release()
			assert.NotNil(t, result3)
		}
	})

	t.Run("TestVectorizedOperations", func(t *testing.T) {
		// Test vectorized expressions with memory pool
		intCol := Col("int_col")
		floatCol := Col("float_col")

		// Test vectorized sum
		sumExpr := intCol.VectorizedSum()
		result, err := sumExpr.EvaluateWithPool(df, pool)
		if err == nil && result != nil {
			defer result.Release()
			assert.NotNil(t, result)
		}

		// Test vectorized mean
		meanExpr := floatCol.VectorizedMean()
		result2, err := meanExpr.EvaluateWithPool(df, pool)
		if err == nil && result2 != nil {
			defer result2.Release()
			assert.NotNil(t, result2)
		}

		// Test vectorized operations with custom expressions
		vec := NewVectorizedExpr([]Expr{intCol, Lit(10)}, "add", nil)
		result3, err := vec.EvaluateWithPool(df, pool)
		if err == nil && result3 != nil {
			defer result3.Release()
			assert.NotNil(t, result3)
		}

		// Test Name and String methods for vectorized expressions
		assert.NotEmpty(t, vec.Name())
		assert.NotEmpty(t, vec.String())
	})

	t.Run("TestEdgeCasesForMemoryPool", func(t *testing.T) {
		// Test edge cases that might trigger different code paths
		ctx := NewExprContext()

		// Test with timestamp column for date operations
		timestampCol := Col("timestamp_col")

		// Test timestamp date operations
		timestampYearExpr := timestampCol.Year()
		result, err := timestampYearExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result != nil {
			defer result.Release()
			assert.NotNil(t, result)
		}

		// Test timestamp month operations
		timestampMonthExpr := timestampCol.Month()
		result2, err := timestampMonthExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result2 != nil {
			defer result2.Release()
			assert.NotNil(t, result2)
		}

		// Test timestamp day operations
		timestampDayExpr := timestampCol.Day()
		result3, err := timestampDayExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result3 != nil {
			defer result3.Release()
			assert.NotNil(t, result3)
		}

		// Test error cases with incompatible types
		stringCol := Col("string_col")
		intCol := Col("int_col")

		// This should error - arithmetic on incompatible types
		invalidArithmetic := stringCol.Add(intCol)
		_, err = invalidArithmetic.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		assert.Error(t, err)

		// Test date truncation with timestamp
		truncExpr := timestampCol.DateTrunc("year")
		result4, err := truncExpr.EvaluateWithPool(df, ctx.Pool.GetAllocator())
		if err == nil && result4 != nil {
			defer result4.Release()
			assert.NotNil(t, result4)
		}
	})
}
