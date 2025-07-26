package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteralExprMethods(t *testing.T) {
	t.Run("LiteralStringOperations", func(t *testing.T) {
		lit := Lit("test")

		// Test Contains
		containsExpr := lit.Contains(Lit("es"))
		assert.NotNil(t, containsExpr)
		assert.IsType(t, &BinaryExpr{}, containsExpr)

		// Test StartsWith
		startsExpr := lit.StartsWith(Lit("te"))
		assert.NotNil(t, startsExpr)
		assert.IsType(t, &BinaryExpr{}, startsExpr)

		// Test EndsWith
		endsExpr := lit.EndsWith(Lit("st"))
		assert.NotNil(t, endsExpr)
		assert.IsType(t, &BinaryExpr{}, endsExpr)
	})

	t.Run("LiteralDateOperations", func(t *testing.T) {
		dateLit := Lit(19523) // Represents a date value

		// Test AddDays
		addDaysExpr := dateLit.AddDays(Lit(10))
		assert.NotNil(t, addDaysExpr)
		assert.IsType(t, &BinaryExpr{}, addDaysExpr)

		// Test AddMonths
		addMonthsExpr := dateLit.AddMonths(Lit(2))
		assert.NotNil(t, addMonthsExpr)
		assert.IsType(t, &BinaryExpr{}, addMonthsExpr)

		// Test AddYears
		addYearsExpr := dateLit.AddYears(Lit(1))
		assert.NotNil(t, addYearsExpr)
		assert.IsType(t, &BinaryExpr{}, addYearsExpr)

		// Test Year
		yearExpr := dateLit.Year()
		assert.NotNil(t, yearExpr)
		assert.IsType(t, &UnaryExpr{}, yearExpr)

		// Test Month
		monthExpr := dateLit.Month()
		assert.NotNil(t, monthExpr)
		assert.IsType(t, &UnaryExpr{}, monthExpr)

		// Test Day
		dayExpr := dateLit.Day()
		assert.NotNil(t, dayExpr)
		assert.IsType(t, &UnaryExpr{}, dayExpr)

		// Test DayOfWeek
		dayOfWeekExpr := dateLit.DayOfWeek()
		assert.NotNil(t, dayOfWeekExpr)
		assert.IsType(t, &UnaryExpr{}, dayOfWeekExpr)

		// Test DateTrunc
		dateTruncExpr := dateLit.DateTrunc("month")
		assert.NotNil(t, dateTruncExpr)
		assert.IsType(t, &UnaryExpr{}, dateTruncExpr)
	})
}
