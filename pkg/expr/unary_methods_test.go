package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnaryExprMethods(t *testing.T) {
	// Create a unary expression for testing
	baseExpr := Col("test")
	unaryExpr := baseExpr.Year() // This creates a UnaryExpr

	t.Run("UnaryExprName", func(t *testing.T) {
		name := unaryExpr.Name()
		assert.NotEmpty(t, name)
	})

	t.Run("UnaryExprString", func(t *testing.T) {
		str := unaryExpr.String()
		assert.NotEmpty(t, str)
		assert.Contains(t, str, "year")
	})

	t.Run("UnaryExprArithmetic", func(t *testing.T) {
		// Test Add
		addExpr := unaryExpr.Add(Lit(10))
		assert.NotNil(t, addExpr)
		assert.IsType(t, &BinaryExpr{}, addExpr)

		// Test Sub
		subExpr := unaryExpr.Sub(Lit(5))
		assert.NotNil(t, subExpr)
		assert.IsType(t, &BinaryExpr{}, subExpr)

		// Test Mul
		mulExpr := unaryExpr.Mul(Lit(2))
		assert.NotNil(t, mulExpr)
		assert.IsType(t, &BinaryExpr{}, mulExpr)

		// Test Div
		divExpr := unaryExpr.Div(Lit(2))
		assert.NotNil(t, divExpr)
		assert.IsType(t, &BinaryExpr{}, divExpr)
	})

	t.Run("UnaryExprComparison", func(t *testing.T) {
		// Test Gt
		gtExpr := unaryExpr.Gt(Lit(2020))
		assert.NotNil(t, gtExpr)
		assert.IsType(t, &BinaryExpr{}, gtExpr)

		// Test Lt
		ltExpr := unaryExpr.Lt(Lit(2025))
		assert.NotNil(t, ltExpr)
		assert.IsType(t, &BinaryExpr{}, ltExpr)

		// Test Eq
		eqExpr := unaryExpr.Eq(Lit(2023))
		assert.NotNil(t, eqExpr)
		assert.IsType(t, &BinaryExpr{}, eqExpr)
	})

	t.Run("UnaryExprStringOperations", func(t *testing.T) {
		// Create a unary expression (e.g., from a date operation)
		dateUnary := Col("date").Year()

		// Test Contains (even though it might not make semantic sense, it should still create the expression)
		containsExpr := dateUnary.Contains(Lit("TEST"))
		assert.NotNil(t, containsExpr)
		assert.IsType(t, &BinaryExpr{}, containsExpr)

		// Test StartsWith
		startsExpr := dateUnary.StartsWith(Lit("T"))
		assert.NotNil(t, startsExpr)
		assert.IsType(t, &BinaryExpr{}, startsExpr)

		// Test EndsWith
		endsExpr := dateUnary.EndsWith(Lit("T"))
		assert.NotNil(t, endsExpr)
		assert.IsType(t, &BinaryExpr{}, endsExpr)
	})

	t.Run("UnaryExprDateOperations", func(t *testing.T) {
		// Test date operations on unary expressions
		yearUnary := Col("date").Year()

		// Test AddDays
		addDaysExpr := yearUnary.AddDays(Lit(10))
		assert.NotNil(t, addDaysExpr)
		assert.IsType(t, &BinaryExpr{}, addDaysExpr)

		// Test AddMonths
		addMonthsExpr := yearUnary.AddMonths(Lit(2))
		assert.NotNil(t, addMonthsExpr)
		assert.IsType(t, &BinaryExpr{}, addMonthsExpr)

		// Test AddYears
		addYearsExpr := yearUnary.AddYears(Lit(1))
		assert.NotNil(t, addYearsExpr)
		assert.IsType(t, &BinaryExpr{}, addYearsExpr)

		// Test Year on a unary expression
		yearExpr := yearUnary.Year()
		assert.NotNil(t, yearExpr)
		assert.IsType(t, &UnaryExpr{}, yearExpr)

		// Test Month
		monthExpr := yearUnary.Month()
		assert.NotNil(t, monthExpr)
		assert.IsType(t, &UnaryExpr{}, monthExpr)

		// Test Day
		dayExpr := yearUnary.Day()
		assert.NotNil(t, dayExpr)
		assert.IsType(t, &UnaryExpr{}, dayExpr)

		// Test DayOfWeek
		dayOfWeekExpr := yearUnary.DayOfWeek()
		assert.NotNil(t, dayOfWeekExpr)
		assert.IsType(t, &UnaryExpr{}, dayOfWeekExpr)

		// Test DateTrunc
		dateTruncExpr := yearUnary.DateTrunc("month")
		assert.NotNil(t, dateTruncExpr)
		assert.IsType(t, &UnaryExpr{}, dateTruncExpr)
	})
}
