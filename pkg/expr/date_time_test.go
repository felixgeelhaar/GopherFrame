package expr

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDateTimeExpressions tests all date/time operations
func TestDateTimeExpressions(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{"YearExtraction", testYearExtraction},
		{"MonthExtraction", testMonthExtraction},
		{"DayExtraction", testDayExtraction},
		{"DayOfWeekExtraction", testDayOfWeekExtraction},
		{"DateTruncation", testDateTruncation},
		{"AddDays", testAddDays},
		{"AddMonths", testAddMonths},
		{"AddYears", testAddYears},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func createDateTestDataFrame() *core.DataFrame {
	pool := memory.NewGoAllocator()

	// Create test dates: 2023-06-15, 2023-12-31, 2024-02-29 (leap year)
	testDates := []arrow.Date32{
		arrow.Date32(19523), // 2023-06-15 (days since epoch)
		arrow.Date32(19722), // 2023-12-31
		arrow.Date32(19782), // 2024-02-29 (leap year)
	}

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "date_col", Type: arrow.FixedWidthTypes.Date32},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	dateBuilder := array.NewDate32Builder(pool)
	for _, date := range testDates {
		dateBuilder.Append(date)
	}
	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, dateArray}, 3)
	defer record.Release()

	return core.NewDataFrame(record)
}

func testYearExtraction(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test Year() extraction
	yearExpr := Col("date_col").Year()
	result, err := yearExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.INT32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	yearArray := result.(*array.Int32)
	expected := []int32{2023, 2023, 2024}
	for i, expectedYear := range expected {
		assert.False(t, yearArray.IsNull(i), "Year should not be null at index %d", i)
		assert.Equal(t, expectedYear, yearArray.Value(i), "Year mismatch at index %d", i)
	}
}

func testMonthExtraction(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test Month() extraction
	monthExpr := Col("date_col").Month()
	result, err := monthExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.INT32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	monthArray := result.(*array.Int32)
	expected := []int32{6, 12, 2} // June, December, February
	for i, expectedMonth := range expected {
		assert.False(t, monthArray.IsNull(i), "Month should not be null at index %d", i)
		assert.Equal(t, expectedMonth, monthArray.Value(i), "Month mismatch at index %d", i)
	}
}

func testDayExtraction(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test Day() extraction
	dayExpr := Col("date_col").Day()
	result, err := dayExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.INT32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	dayArray := result.(*array.Int32)
	expected := []int32{15, 31, 29} // 15th, 31st, 29th
	for i, expectedDay := range expected {
		assert.False(t, dayArray.IsNull(i), "Day should not be null at index %d", i)
		assert.Equal(t, expectedDay, dayArray.Value(i), "Day mismatch at index %d", i)
	}
}

func testDayOfWeekExtraction(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test DayOfWeek() extraction (ISO: Monday=1, Sunday=7)
	dayOfWeekExpr := Col("date_col").DayOfWeek()
	result, err := dayOfWeekExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.INT32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	dayOfWeekArray := result.(*array.Int32)
	// 2023-06-15 = Thursday (4), 2023-12-31 = Sunday (7), 2024-02-29 = Thursday (4)
	// Note: Go's Sunday=0, but we convert to ISO Monday=1, so Sunday becomes 7
	expected := []int32{4, 7, 4}
	for i, expectedDayOfWeek := range expected {
		assert.False(t, dayOfWeekArray.IsNull(i), "DayOfWeek should not be null at index %d", i)
		assert.Equal(t, expectedDayOfWeek, dayOfWeekArray.Value(i), "DayOfWeek mismatch at index %d", i)
	}
}

func testDateTruncation(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test DateTrunc("month") - should truncate to first day of month
	truncExpr := Col("date_col").DateTrunc("month")
	result, err := truncExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.DATE32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	truncArray := result.(*array.Date32)
	for i := 0; i < truncArray.Len(); i++ {
		assert.False(t, truncArray.IsNull(i), "Truncated date should not be null at index %d", i)
		// Use time conversion to verify the truncation worked correctly
		actualDays := truncArray.Value(i)
		actualTime := time.Unix(int64(actualDays)*86400, 0).UTC()
		assert.Equal(t, 1, actualTime.Day(), "Truncated date should be 1st of month at index %d", i)
	}
}

func testAddDays(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test AddDays(10) - add 10 days to each date
	addDaysExpr := Col("date_col").AddDays(Lit(10))
	result, err := addDaysExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.DATE32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	addDaysArray := result.(*array.Date32)
	// Expected: original dates + 10 days
	originalDates := []arrow.Date32{19523, 19722, 19782}
	for i, originalDate := range originalDates {
		assert.False(t, addDaysArray.IsNull(i), "AddDays result should not be null at index %d", i)
		expected := originalDate + 10
		assert.Equal(t, expected, addDaysArray.Value(i), "AddDays mismatch at index %d", i)
	}
}

func testAddMonths(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test AddMonths(2) - add 2 months to each date
	addMonthsExpr := Col("date_col").AddMonths(Lit(2))
	result, err := addMonthsExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.DATE32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	addMonthsArray := result.(*array.Date32)
	// Expected: 2023-08-15, 2024-03-02 (Dec 31 + 2 months), 2024-04-29
	expectedTimes := []time.Time{
		time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC), // 2023-06-15 + 2 months
		time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC),  // 2023-12-31 + 2 months
		time.Date(2024, 4, 29, 0, 0, 0, 0, time.UTC), // 2024-02-29 + 2 months
	}

	for i, expectedTime := range expectedTimes {
		assert.False(t, addMonthsArray.IsNull(i), "AddMonths result should not be null at index %d", i)
		actualDays := addMonthsArray.Value(i)
		actualTime := time.Unix(int64(actualDays)*86400, 0).UTC()
		assert.Equal(t, expectedTime.Year(), actualTime.Year(), "AddMonths year mismatch at index %d", i)
		assert.Equal(t, expectedTime.Month(), actualTime.Month(), "AddMonths month mismatch at index %d", i)
		assert.Equal(t, expectedTime.Day(), actualTime.Day(), "AddMonths day mismatch at index %d", i)
	}
}

func testAddYears(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	// Test AddYears(1) - add 1 year to each date
	addYearsExpr := Col("date_col").AddYears(Lit(1))
	result, err := addYearsExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, arrow.DATE32, result.DataType().ID())
	assert.Equal(t, 3, result.Len())

	addYearsArray := result.(*array.Date32)
	// Expected: 2024-06-15, 2024-12-31, 2025-03-01 (Feb 29 doesn't exist in non-leap year)
	expectedTimes := []time.Time{
		time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC),  // 2023-06-15 + 1 year
		time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), // 2023-12-31 + 1 year
		time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),   // 2024-02-29 + 1 year (no leap year, so March 1)
	}

	for i, expectedTime := range expectedTimes {
		assert.False(t, addYearsArray.IsNull(i), "AddYears result should not be null at index %d", i)
		actualDays := addYearsArray.Value(i)
		actualTime := time.Unix(int64(actualDays)*86400, 0).UTC()
		assert.Equal(t, expectedTime.Year(), actualTime.Year(), "AddYears year mismatch at index %d", i)
		assert.Equal(t, expectedTime.Month(), actualTime.Month(), "AddYears month mismatch at index %d", i)
		assert.Equal(t, expectedTime.Day(), actualTime.Day(), "AddYears day mismatch at index %d", i)
	}
}

// Test error conditions
func TestDateTimeErrorConditions(t *testing.T) {
	df := createDateTestDataFrame()
	defer df.Release()

	t.Run("YearOnNonDateColumn", func(t *testing.T) {
		yearExpr := Col("id").Year()
		_, err := yearExpr.Evaluate(df)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type for year extraction")
	})

	t.Run("AddDaysWithWrongType", func(t *testing.T) {
		addDaysExpr := Col("date_col").AddDays(Lit("invalid"))
		_, err := addDaysExpr.Evaluate(df)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "AddDays requires integer days parameter")
	})
}

// Test with null values
func TestDateTimeWithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "date_col", Type: arrow.FixedWidthTypes.Date32},
		},
		nil,
	)

	dateBuilder := array.NewDate32Builder(pool)
	dateBuilder.Append(arrow.Date32(19523)) // 2023-06-15
	dateBuilder.AppendNull()                // null value
	dateBuilder.Append(arrow.Date32(19722)) // 2023-12-31
	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()

	record := array.NewRecord(schema, []arrow.Array{dateArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Year() with nulls
	yearExpr := Col("date_col").Year()
	result, err := yearExpr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	yearArray := result.(*array.Int32)
	assert.Equal(t, 3, yearArray.Len())
	assert.False(t, yearArray.IsNull(0))
	assert.True(t, yearArray.IsNull(1)) // Should preserve null
	assert.False(t, yearArray.IsNull(2))

	assert.Equal(t, int32(2023), yearArray.Value(0))
	assert.Equal(t, int32(2023), yearArray.Value(2))
}
