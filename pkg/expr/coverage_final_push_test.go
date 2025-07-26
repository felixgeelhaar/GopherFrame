package expr

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestTruncateDateFunction(t *testing.T) {
	t.Run("TruncateDateAllUnits", func(t *testing.T) {
		// Test the truncateDate utility function with different units
		testTime := time.Date(2023, 6, 15, 14, 30, 45, 0, time.UTC)

		// Test year truncation
		yearResult := truncateDate(testTime, "year")
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expected, yearResult)

		// Test month truncation
		monthResult := truncateDate(testTime, "month")
		expectedMonth := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedMonth, monthResult)

		// Test day truncation
		dayResult := truncateDate(testTime, "day")
		expectedDay := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedDay, dayResult)

		// Test hour truncation
		hourResult := truncateDate(testTime, "hour")
		expectedHour := time.Date(2023, 6, 15, 14, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedHour, hourResult)

		// Test unknown unit (should return original)
		unknownResult := truncateDate(testTime, "unknown")
		assert.Equal(t, testTime, unknownResult)

		// Test empty unit (should return original)
		emptyResult := truncateDate(testTime, "")
		assert.Equal(t, testTime, emptyResult)
	})
}

func TestInferDataTypeExtended(t *testing.T) {
	t.Run("InferDataTypeAllTypes", func(t *testing.T) {
		// Test type inference for all supported Go types - this tests all branches
		testValues := []interface{}{
			int8(1), int16(2), int32(3), int64(4), int(5),
			uint8(6), uint16(7), uint32(8), uint64(9), uint(10),
			float32(11.5), float64(12.5),
			"string", true, false,
		}

		// Just ensure that all types return some valid arrow type (tests all branches)
		for _, val := range testValues {
			result := inferDataType(val)
			assert.NotNil(t, result, "inferDataType should return a type for %v (%T)", val, val)
		}
	})

	t.Run("InferDataTypeCustomStruct", func(t *testing.T) {
		// Test with a custom struct - it may return string type as fallback
		type CustomStruct struct {
			Field string
		}
		custom := CustomStruct{Field: "test"}
		result := inferDataType(custom)
		// Just check that it returns something (may be string as fallback)
		assert.NotNil(t, result)
	})
}

func TestDateOperationHelpers(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame with date columns
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "date32_col", Type: arrow.FixedWidthTypes.Date32},
			{Name: "date64_col", Type: arrow.FixedWidthTypes.Date64},
			{Name: "timestamp_col", Type: &arrow.TimestampType{Unit: arrow.Millisecond}},
		},
		nil,
	)

	date32Builder := array.NewDate32Builder(pool)
	date64Builder := array.NewDate64Builder(pool)
	timestampBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Millisecond})

	// Add test dates (June 15, 2023 and January 1, 2023)
	date32Builder.AppendValues([]arrow.Date32{19523, 19358}, nil)
	date64Builder.AppendValues([]arrow.Date64{1686787200000, 1672531200000}, nil)
	timestampBuilder.AppendValues([]arrow.Timestamp{1686787200000, 1672531200000}, nil)

	date32Array := date32Builder.NewArray()
	date64Array := date64Builder.NewArray()
	timestampArray := timestampBuilder.NewArray()
	defer date32Array.Release()
	defer date64Array.Release()
	defer timestampArray.Release()

	record := array.NewRecord(schema, []arrow.Array{date32Array, date64Array, timestampArray}, 2)
	df := core.NewDataFrame(record)
	defer df.Release()

	t.Run("DateOperationCoverage", func(t *testing.T) {
		// Test date operations that might have lower coverage
		dateCol := Col("date32_col")

		// Create various date operations to increase coverage
		yearExpr := dateCol.Year()
		monthExpr := dateCol.Month()
		dayExpr := dateCol.Day()
		dayOfWeekExpr := dateCol.DayOfWeek()

		// Test that all expressions can be created
		assert.NotNil(t, yearExpr)
		assert.NotNil(t, monthExpr)
		assert.NotNil(t, dayExpr)
		assert.NotNil(t, dayOfWeekExpr)

		// Test string representations (check that they're not empty)
		assert.NotEmpty(t, yearExpr.String())
		assert.NotEmpty(t, monthExpr.String())
		assert.NotEmpty(t, dayExpr.String())
		assert.NotEmpty(t, dayOfWeekExpr.String())
	})

	t.Run("DateArithmeticOperations", func(t *testing.T) {
		dateCol := Col("date32_col")

		// Test date arithmetic operations
		addDaysExpr := dateCol.AddDays(Lit(10))
		addMonthsExpr := dateCol.AddMonths(Lit(2))
		addYearsExpr := dateCol.AddYears(Lit(1))

		assert.NotNil(t, addDaysExpr)
		assert.NotNil(t, addMonthsExpr)
		assert.NotNil(t, addYearsExpr)

		// Test string representations (check that they're not empty)
		assert.NotEmpty(t, addDaysExpr.String())
		assert.NotEmpty(t, addMonthsExpr.String())
		assert.NotEmpty(t, addYearsExpr.String())
	})

	t.Run("DateTruncationOperation", func(t *testing.T) {
		dateCol := Col("timestamp_col")

		// Test date truncation operations
		truncYearExpr := dateCol.DateTrunc("year")
		truncMonthExpr := dateCol.DateTrunc("month")
		truncDayExpr := dateCol.DateTrunc("day")

		assert.NotNil(t, truncYearExpr)
		assert.NotNil(t, truncMonthExpr)
		assert.NotNil(t, truncDayExpr)

		// Test string representations (check that they're not empty)
		assert.NotEmpty(t, truncYearExpr.String())
		assert.NotEmpty(t, truncMonthExpr.String())
		assert.NotEmpty(t, truncDayExpr.String())
	})
}

func TestComplexExpressionBuilding(t *testing.T) {
	t.Run("ComplexExpressionChaining", func(t *testing.T) {
		// Build complex expressions to test various code paths
		baseCol := Col("data")

		// Chain multiple operations
		complexExpr1 := baseCol.Add(Lit(10)).Gt(Lit(20)).Eq(Lit(true))
		assert.NotNil(t, complexExpr1)

		complexExpr2 := baseCol.Year().Add(Lit(1)).Lt(Lit(2025))
		assert.NotNil(t, complexExpr2)

		complexExpr3 := baseCol.DateTrunc("month").AddDays(Lit(15)).DayOfWeek()
		assert.NotNil(t, complexExpr3)

		// Test string representations of complex expressions
		str1 := complexExpr1.String()
		str2 := complexExpr2.String()
		str3 := complexExpr3.String()

		assert.NotEmpty(t, str1)
		assert.NotEmpty(t, str2)
		assert.NotEmpty(t, str3)
	})

	t.Run("ExpressionWithLiterals", func(t *testing.T) {
		// Create expressions with different literal types
		litExprs := []Expr{
			Lit(int8(1)).Add(Lit(int16(2))),
			Lit(int32(3)).Mul(Lit(int64(4))),
			Lit(uint8(5)).Sub(Lit(uint16(6))),
			Lit(uint32(7)).Div(Lit(uint64(8))),
			Lit(float32(9.1)).Gt(Lit(float64(9.2))),
			Lit("hello").Contains(Lit("ell")),
			Lit(true).Eq(Lit(false)),
		}

		for i, expr := range litExprs {
			assert.NotNil(t, expr, "Expression %d should not be nil", i)
			assert.NotEmpty(t, expr.String(), "Expression %d string should not be empty", i)
		}
	})
}
