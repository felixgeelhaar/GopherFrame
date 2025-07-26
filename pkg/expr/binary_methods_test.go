package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryExprMethods(t *testing.T) {
	// Create a binary expression for testing
	binaryExpr := Col("a").Add(Col("b"))

	t.Run("BinaryExprStringOperations", func(t *testing.T) {
		// Test Contains
		containsExpr := binaryExpr.Contains(Lit("test"))
		assert.NotNil(t, containsExpr)
		assert.IsType(t, &BinaryExpr{}, containsExpr)

		// Test StartsWith
		startsExpr := binaryExpr.StartsWith(Lit("start"))
		assert.NotNil(t, startsExpr)
		assert.IsType(t, &BinaryExpr{}, startsExpr)

		// Test EndsWith
		endsExpr := binaryExpr.EndsWith(Lit("end"))
		assert.NotNil(t, endsExpr)
		assert.IsType(t, &BinaryExpr{}, endsExpr)
	})

	t.Run("BinaryExprDateOperations", func(t *testing.T) {
		// Test AddDays
		addDaysExpr := binaryExpr.AddDays(Lit(10))
		assert.NotNil(t, addDaysExpr)
		assert.IsType(t, &BinaryExpr{}, addDaysExpr)

		// Test AddMonths
		addMonthsExpr := binaryExpr.AddMonths(Lit(2))
		assert.NotNil(t, addMonthsExpr)
		assert.IsType(t, &BinaryExpr{}, addMonthsExpr)

		// Test AddYears
		addYearsExpr := binaryExpr.AddYears(Lit(1))
		assert.NotNil(t, addYearsExpr)
		assert.IsType(t, &BinaryExpr{}, addYearsExpr)

		// Test Year
		yearExpr := binaryExpr.Year()
		assert.NotNil(t, yearExpr)
		assert.IsType(t, &UnaryExpr{}, yearExpr)

		// Test Month
		monthExpr := binaryExpr.Month()
		assert.NotNil(t, monthExpr)
		assert.IsType(t, &UnaryExpr{}, monthExpr)

		// Test Day
		dayExpr := binaryExpr.Day()
		assert.NotNil(t, dayExpr)
		assert.IsType(t, &UnaryExpr{}, dayExpr)

		// Test DayOfWeek
		dayOfWeekExpr := binaryExpr.DayOfWeek()
		assert.NotNil(t, dayOfWeekExpr)
		assert.IsType(t, &UnaryExpr{}, dayOfWeekExpr)

		// Test DateTrunc
		dateTruncExpr := binaryExpr.DateTrunc("month")
		assert.NotNil(t, dateTruncExpr)
		assert.IsType(t, &UnaryExpr{}, dateTruncExpr)
	})
}
