package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorizedExprMethods(t *testing.T) {
	// Create a vectorized expression for testing
	operands := []Expr{Col("x"), Col("y")}
	vExpr := NewVectorizedExpr(operands, "add", nil)

	t.Run("VectorizedExprArithmetic", func(t *testing.T) {
		// Test Sub
		subExpr := vExpr.Sub(Lit(5))
		assert.NotNil(t, subExpr)
		assert.IsType(t, &BinaryExpr{}, subExpr)

		// Test Mul
		mulExpr := vExpr.Mul(Lit(2))
		assert.NotNil(t, mulExpr)
		assert.IsType(t, &BinaryExpr{}, mulExpr)

		// Test Div
		divExpr := vExpr.Div(Lit(2))
		assert.NotNil(t, divExpr)
		assert.IsType(t, &BinaryExpr{}, divExpr)
	})

	t.Run("VectorizedExprComparison", func(t *testing.T) {
		// Test Gt
		gtExpr := vExpr.Gt(Lit(10))
		assert.NotNil(t, gtExpr)
		assert.IsType(t, &BinaryExpr{}, gtExpr)

		// Test Lt
		ltExpr := vExpr.Lt(Lit(20))
		assert.NotNil(t, ltExpr)
		assert.IsType(t, &BinaryExpr{}, ltExpr)

		// Test Eq
		eqExpr := vExpr.Eq(Lit(15))
		assert.NotNil(t, eqExpr)
		assert.IsType(t, &BinaryExpr{}, eqExpr)
	})

	t.Run("VectorizedExprStringOperations", func(t *testing.T) {
		// Test Contains
		containsExpr := vExpr.Contains(Lit("test"))
		assert.NotNil(t, containsExpr)
		assert.IsType(t, &BinaryExpr{}, containsExpr)

		// Test StartsWith
		startsExpr := vExpr.StartsWith(Lit("start"))
		assert.NotNil(t, startsExpr)
		assert.IsType(t, &BinaryExpr{}, startsExpr)

		// Test EndsWith
		endsExpr := vExpr.EndsWith(Lit("end"))
		assert.NotNil(t, endsExpr)
		assert.IsType(t, &BinaryExpr{}, endsExpr)
	})

	t.Run("VectorizedExprDateOperations", func(t *testing.T) {
		// Test AddDays
		addDaysExpr := vExpr.AddDays(Lit(10))
		assert.NotNil(t, addDaysExpr)
		assert.IsType(t, &BinaryExpr{}, addDaysExpr)

		// Test AddMonths
		addMonthsExpr := vExpr.AddMonths(Lit(2))
		assert.NotNil(t, addMonthsExpr)
		assert.IsType(t, &BinaryExpr{}, addMonthsExpr)

		// Test AddYears
		addYearsExpr := vExpr.AddYears(Lit(1))
		assert.NotNil(t, addYearsExpr)
		assert.IsType(t, &BinaryExpr{}, addYearsExpr)

		// Test Year
		yearExpr := vExpr.Year()
		assert.NotNil(t, yearExpr)
		assert.IsType(t, &UnaryExpr{}, yearExpr)

		// Test Month
		monthExpr := vExpr.Month()
		assert.NotNil(t, monthExpr)
		assert.IsType(t, &UnaryExpr{}, monthExpr)

		// Test Day
		dayExpr := vExpr.Day()
		assert.NotNil(t, dayExpr)
		assert.IsType(t, &UnaryExpr{}, dayExpr)

		// Test DayOfWeek
		dayOfWeekExpr := vExpr.DayOfWeek()
		assert.NotNil(t, dayOfWeekExpr)
		assert.IsType(t, &UnaryExpr{}, dayOfWeekExpr)

		// Test DateTrunc
		dateTruncExpr := vExpr.DateTrunc("month")
		assert.NotNil(t, dateTruncExpr)
		assert.IsType(t, &UnaryExpr{}, dateTruncExpr)
	})
}

func TestVectorizedFunctions(t *testing.T) {
	t.Run("ColumnVectorizedFunctions", func(t *testing.T) {
		col := Col("data")

		// Test VectorizedSub
		subExpr := col.VectorizedSub(Col("other"))
		assert.NotNil(t, subExpr)
		assert.IsType(t, &VectorizedExpr{}, subExpr)

		// Test VectorizedDiv
		divExpr := col.VectorizedDiv(Col("divisor"))
		assert.NotNil(t, divExpr)
		assert.IsType(t, &VectorizedExpr{}, divExpr)

		// Test VectorizedLt
		ltExpr := col.VectorizedLt(Col("threshold"))
		assert.NotNil(t, ltExpr)
		assert.IsType(t, &VectorizedExpr{}, ltExpr)

		// Test VectorizedEq
		eqExpr := col.VectorizedEq(Col("target"))
		assert.NotNil(t, eqExpr)
		assert.IsType(t, &VectorizedExpr{}, eqExpr)

		// Test VectorizedMean
		meanExpr := col.VectorizedMean()
		assert.NotNil(t, meanExpr)
		assert.IsType(t, &VectorizedExpr{}, meanExpr)

		// Test VectorizedMin
		minExpr := col.VectorizedMin()
		assert.NotNil(t, minExpr)
		assert.IsType(t, &VectorizedExpr{}, minExpr)

		// Test VectorizedMax
		maxExpr := col.VectorizedMax()
		assert.NotNil(t, maxExpr)
		assert.IsType(t, &VectorizedExpr{}, maxExpr)

		// Test VectorizedStdDev
		stdDevExpr := col.VectorizedStdDev()
		assert.NotNil(t, stdDevExpr)
		assert.IsType(t, &VectorizedExpr{}, stdDevExpr)

		// Test VectorizedVariance
		varExpr := col.VectorizedVariance()
		assert.NotNil(t, varExpr)
		assert.IsType(t, &VectorizedExpr{}, varExpr)
	})

	t.Run("LiteralVectorizedFunctions", func(t *testing.T) {
		lit := Lit(10)

		// Test all vectorized functions on literals
		assert.NotNil(t, lit.VectorizedAdd(Lit(5)))
		assert.NotNil(t, lit.VectorizedSub(Lit(5)))
		assert.NotNil(t, lit.VectorizedMul(Lit(5)))
		assert.NotNil(t, lit.VectorizedDiv(Lit(5)))
		assert.NotNil(t, lit.VectorizedGt(Lit(5)))
		assert.NotNil(t, lit.VectorizedLt(Lit(5)))
		assert.NotNil(t, lit.VectorizedEq(Lit(5)))
		assert.NotNil(t, lit.VectorizedSum())
		assert.NotNil(t, lit.VectorizedMean())
		assert.NotNil(t, lit.VectorizedMin())
		assert.NotNil(t, lit.VectorizedMax())
		assert.NotNil(t, lit.VectorizedStdDev())
		assert.NotNil(t, lit.VectorizedVariance())
	})

	t.Run("UnaryVectorizedFunctions", func(t *testing.T) {
		unary := Col("data").Year()

		// Test all vectorized functions on unary expressions
		assert.NotNil(t, unary.VectorizedAdd(Lit(5)))
		assert.NotNil(t, unary.VectorizedSub(Lit(5)))
		assert.NotNil(t, unary.VectorizedMul(Lit(5)))
		assert.NotNil(t, unary.VectorizedDiv(Lit(5)))
		assert.NotNil(t, unary.VectorizedGt(Lit(5)))
		assert.NotNil(t, unary.VectorizedLt(Lit(5)))
		assert.NotNil(t, unary.VectorizedEq(Lit(5)))
		assert.NotNil(t, unary.VectorizedSum())
		assert.NotNil(t, unary.VectorizedMean())
		assert.NotNil(t, unary.VectorizedMin())
		assert.NotNil(t, unary.VectorizedMax())
		assert.NotNil(t, unary.VectorizedStdDev())
		assert.NotNil(t, unary.VectorizedVariance())
	})

	t.Run("BinaryVectorizedFunctions", func(t *testing.T) {
		binary := Col("a").Add(Col("b"))

		// Test all vectorized functions on binary expressions
		assert.NotNil(t, binary.VectorizedAdd(Lit(5)))
		assert.NotNil(t, binary.VectorizedSub(Lit(5)))
		assert.NotNil(t, binary.VectorizedMul(Lit(5)))
		assert.NotNil(t, binary.VectorizedDiv(Lit(5)))
		assert.NotNil(t, binary.VectorizedGt(Lit(5)))
		assert.NotNil(t, binary.VectorizedLt(Lit(5)))
		assert.NotNil(t, binary.VectorizedEq(Lit(5)))
		assert.NotNil(t, binary.VectorizedSum())
		assert.NotNil(t, binary.VectorizedMean())
		assert.NotNil(t, binary.VectorizedMin())
		assert.NotNil(t, binary.VectorizedMax())
		assert.NotNil(t, binary.VectorizedStdDev())
		assert.NotNil(t, binary.VectorizedVariance())
	})

	t.Run("VectorizedExprVectorizedFunctions", func(t *testing.T) {
		vexpr := NewVectorizedExpr([]Expr{Col("x"), Col("y")}, "add", nil)

		// Test all vectorized functions on vectorized expressions
		assert.NotNil(t, vexpr.VectorizedAdd(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedSub(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedMul(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedDiv(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedGt(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedLt(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedEq(Lit(5)))
		assert.NotNil(t, vexpr.VectorizedSum())
		assert.NotNil(t, vexpr.VectorizedMean())
		assert.NotNil(t, vexpr.VectorizedMin())
		assert.NotNil(t, vexpr.VectorizedMax())
		assert.NotNil(t, vexpr.VectorizedStdDev())
		assert.NotNil(t, vexpr.VectorizedVariance())
	})
}
