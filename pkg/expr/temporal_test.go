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

// TestExpr_Year tests extracting year component from timestamps
func TestExpr_Year(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test timestamps
	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	// Test data: 2024-03-15 10:30:45, 2023-12-31 23:59:59, null
	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}
	builder.AppendNull()

	tsArray := builder.NewArray()
	defer tsArray.Release()

	// Create DataFrame with timestamp column
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Year() extraction
	expr := Col("timestamp").Year()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(2024), int64Result.Value(0))
	assert.Equal(t, int64(2023), int64Result.Value(1))
	assert.True(t, int64Result.IsNull(2))
}

// TestExpr_Month tests extracting month component from timestamps
func TestExpr_Month(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").Month()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(3), int64Result.Value(0))
	assert.Equal(t, int64(12), int64Result.Value(1))
	assert.Equal(t, int64(1), int64Result.Value(2))
}

// TestExpr_Day tests extracting day component from timestamps
func TestExpr_Day(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), // leap year
		time.Date(2024, 1, 1, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").Day()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(15), int64Result.Value(0))
	assert.Equal(t, int64(29), int64Result.Value(1))
	assert.Equal(t, int64(1), int64Result.Value(2))
}

// TestExpr_Hour tests extracting hour component from timestamps
func TestExpr_Hour(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 3, 15, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").Hour()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(10), int64Result.Value(0))
	assert.Equal(t, int64(0), int64Result.Value(1))
	assert.Equal(t, int64(23), int64Result.Value(2))
}

// TestExpr_Minute tests extracting minute component from timestamps
func TestExpr_Minute(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 3, 15, 10, 0, 0, 0, time.UTC),
		time.Date(2024, 3, 15, 10, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").Minute()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(30), int64Result.Value(0))
	assert.Equal(t, int64(0), int64Result.Value(1))
	assert.Equal(t, int64(59), int64Result.Value(2))
}

// TestExpr_Second tests extracting second component from timestamps
func TestExpr_Second(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
		time.Date(2024, 3, 15, 10, 30, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").Second()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(45), int64Result.Value(0))
	assert.Equal(t, int64(0), int64Result.Value(1))
	assert.Equal(t, int64(59), int64Result.Value(2))
}

// TestExpr_TruncateToYear tests truncating timestamps to year
func TestExpr_TruncateToYear(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").TruncateToYear()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	// Verify truncation
	t1 := toTime(tsResult.Value(0))
	assert.Equal(t, 2024, t1.Year())
	assert.Equal(t, time.Month(1), t1.Month())
	assert.Equal(t, 1, t1.Day())
	assert.Equal(t, 0, t1.Hour())

	t2 := toTime(tsResult.Value(1))
	assert.Equal(t, 2023, t2.Year())
	assert.Equal(t, time.Month(1), t2.Month())
	assert.Equal(t, 1, t2.Day())
}

// TestExpr_TruncateToMonth tests truncating timestamps to month
func TestExpr_TruncateToMonth(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").TruncateToMonth()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	t1 := toTime(tsResult.Value(0))
	assert.Equal(t, 2024, t1.Year())
	assert.Equal(t, time.Month(3), t1.Month())
	assert.Equal(t, 1, t1.Day())
	assert.Equal(t, 0, t1.Hour())

	t2 := toTime(tsResult.Value(1))
	assert.Equal(t, 2024, t2.Year())
	assert.Equal(t, time.Month(12), t2.Month())
	assert.Equal(t, 1, t2.Day())
}

// TestExpr_TruncateToDay tests truncating timestamps to day
func TestExpr_TruncateToDay(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 3, 15, 23, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").TruncateToDay()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	for i := 0; i < 2; i++ {
		tResult := toTime(tsResult.Value(i))
		assert.Equal(t, 2024, tResult.Year())
		assert.Equal(t, time.Month(3), tResult.Month())
		assert.Equal(t, 15, tResult.Day())
		assert.Equal(t, 0, tResult.Hour())
		assert.Equal(t, 0, tResult.Minute())
		assert.Equal(t, 0, tResult.Second())
	}
}

// TestExpr_TruncateToHour tests truncating timestamps to hour
func TestExpr_TruncateToHour(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	times := []time.Time{
		time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC),
		time.Date(2024, 3, 15, 10, 59, 59, 0, time.UTC),
	}

	for _, tm := range times {
		ts, err := arrow.TimestampFromTime(tm, arrow.Second)
		require.NoError(t, err)
		builder.Append(ts)
	}

	tsArray := builder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").TruncateToHour()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	for i := 0; i < 2; i++ {
		tResult := toTime(tsResult.Value(i))
		assert.Equal(t, 10, tResult.Hour())
		assert.Equal(t, 0, tResult.Minute())
		assert.Equal(t, 0, tResult.Second())
	}
}

// TestExpr_AddDays tests adding days to timestamps
func TestExpr_AddDays(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	tsBuilder := array.NewTimestampBuilder(pool, tsType)
	defer tsBuilder.Release()

	baseTime := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	ts, err := arrow.TimestampFromTime(baseTime, arrow.Second)
	require.NoError(t, err)
	tsBuilder.Append(ts)
	tsBuilder.Append(ts)
	tsBuilder.AppendNull()

	tsArray := tsBuilder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Add 7 days, -2 days, and to null
	expr := Col("timestamp").AddDays(Lit(int64(7)))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	t1 := toTime(tsResult.Value(0))
	expected := baseTime.AddDate(0, 0, 7)
	assert.Equal(t, expected.Year(), t1.Year())
	assert.Equal(t, expected.Month(), t1.Month())
	assert.Equal(t, expected.Day(), t1.Day())
}

// TestExpr_AddHours tests adding hours to timestamps
func TestExpr_AddHours(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	tsBuilder := array.NewTimestampBuilder(pool, tsType)
	defer tsBuilder.Release()

	baseTime := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	ts, err := arrow.TimestampFromTime(baseTime, arrow.Second)
	require.NoError(t, err)
	tsBuilder.Append(ts)

	tsArray := tsBuilder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").AddHours(Lit(int64(25)))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	t1 := toTime(tsResult.Value(0))
	expected := baseTime.Add(25 * time.Hour)
	assert.Equal(t, expected.Day(), t1.Day())
	assert.Equal(t, expected.Hour(), t1.Hour())
}

// TestExpr_AddMinutes tests adding minutes to timestamps
func TestExpr_AddMinutes(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	tsBuilder := array.NewTimestampBuilder(pool, tsType)
	defer tsBuilder.Release()

	baseTime := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	ts, err := arrow.TimestampFromTime(baseTime, arrow.Second)
	require.NoError(t, err)
	tsBuilder.Append(ts)

	tsArray := tsBuilder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").AddMinutes(Lit(int64(90)))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	t1 := toTime(tsResult.Value(0))
	expected := baseTime.Add(90 * time.Minute)
	assert.Equal(t, expected.Hour(), t1.Hour())
	assert.Equal(t, expected.Minute(), t1.Minute())
}

// TestExpr_AddSeconds tests adding seconds to timestamps
func TestExpr_AddSeconds(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	tsBuilder := array.NewTimestampBuilder(pool, tsType)
	defer tsBuilder.Release()

	baseTime := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	ts, err := arrow.TimestampFromTime(baseTime, arrow.Second)
	require.NoError(t, err)
	tsBuilder.Append(ts)

	tsArray := tsBuilder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("timestamp").AddSeconds(Lit(int64(3600)))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	tsResult := result.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	t1 := toTime(tsResult.Value(0))
	expected := baseTime.Add(3600 * time.Second)
	assert.Equal(t, expected.Hour(), t1.Hour())
	assert.Equal(t, expected.Minute(), t1.Minute())
	assert.Equal(t, expected.Second(), t1.Second())
}

// TestExpr_Temporal_MultipleOperations tests chaining temporal operations
func TestExpr_Temporal_MultipleOperations(t *testing.T) {
	pool := memory.NewGoAllocator()

	tsType := &arrow.TimestampType{Unit: arrow.Second}
	tsBuilder := array.NewTimestampBuilder(pool, tsType)
	defer tsBuilder.Release()

	baseTime := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	ts, err := arrow.TimestampFromTime(baseTime, arrow.Second)
	require.NoError(t, err)
	tsBuilder.Append(ts)

	tsArray := tsBuilder.NewArray()
	defer tsArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "timestamp", Type: tsType, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArray}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Chain operations: add 7 days, truncate to day, extract year
	expr1 := Col("timestamp").AddDays(Lit(int64(7)))
	result1, err := expr1.Evaluate(df)
	require.NoError(t, err)
	defer result1.Release()

	// Create new DF with result to test truncation
	schema2 := arrow.NewSchema([]arrow.Field{
		{Name: "ts", Type: tsType, Nullable: true},
	}, nil)
	record2 := array.NewRecord(schema2, []arrow.Array{result1}, 1)
	defer record2.Release()

	df2 := core.NewDataFrame(record2)
	defer df2.Release()

	expr2 := Col("ts").TruncateToDay()
	result2, err := expr2.Evaluate(df2)
	require.NoError(t, err)
	defer result2.Release()

	tsResult := result2.(*array.Timestamp)
	toTime, err := tsType.GetToTimeFunc()
	require.NoError(t, err)

	t1 := toTime(tsResult.Value(0))
	expected := baseTime.AddDate(0, 0, 7)
	assert.Equal(t, expected.Year(), t1.Year())
	assert.Equal(t, expected.Month(), t1.Month())
	assert.Equal(t, expected.Day(), t1.Day())
	assert.Equal(t, 0, t1.Hour())
	assert.Equal(t, 0, t1.Minute())
	assert.Equal(t, 0, t1.Second())
}
