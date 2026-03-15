package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDateColumn_ISO(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "date_str", Type: arrow.BinaryTypes.String},
	}, nil)
	db := array.NewStringBuilder(pool)
	db.AppendValues([]string{"2024-01-15", "2024-06-30", "2024-12-25"}, nil)
	da := db.NewArray()
	defer da.Release()
	db.Release()

	rec := array.NewRecord(schema, []arrow.Array{da}, 3)
	defer rec.Release()

	df := NewDataFrame(rec)
	result := df.ParseDateColumn("date_str", "date")
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("date"))
	assert.Equal(t, int64(2), result.NumCols()) // original + new timestamp
}

func TestParseDateColumn_USFormat(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "date_str", Type: arrow.BinaryTypes.String},
	}, nil)
	db := array.NewStringBuilder(pool)
	db.AppendValues([]string{"01/15/2024", "06/30/2024"}, nil)
	da := db.NewArray()
	defer da.Release()
	db.Release()

	rec := array.NewRecord(schema, []arrow.Array{da}, 2)
	defer rec.Release()

	df := NewDataFrame(rec)
	result := df.ParseDateColumn("date_str", "date")
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("date"))
}

func TestParseDateWithFormat(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "ts", Type: arrow.BinaryTypes.String},
	}, nil)
	tb := array.NewStringBuilder(pool)
	tb.AppendValues([]string{"2024-01-15 10:30:00", "2024-06-30 14:00:00"}, nil)
	ta := tb.NewArray()
	defer ta.Release()
	tb.Release()

	rec := array.NewRecord(schema, []arrow.Array{ta}, 2)
	defer rec.Release()

	df := NewDataFrame(rec)
	result := df.ParseDateWithFormat("ts", "timestamp", "2006-01-02 15:04:05")
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("timestamp"))
}

func TestParseDateColumn_WithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "date_str", Type: arrow.BinaryTypes.String},
	}, nil)
	db := array.NewStringBuilder(pool)
	db.Append("2024-01-15")
	db.AppendNull()
	db.Append("2024-12-25")
	da := db.NewArray()
	defer da.Release()
	db.Release()

	rec := array.NewRecord(schema, []arrow.Array{da}, 3)
	defer rec.Release()

	df := NewDataFrame(rec)
	result := df.ParseDateColumn("date_str", "date")
	require.NoError(t, result.Err())
}

func TestParseDateColumn_InvalidColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.Append(1.0)
	xa := xb.NewArray()
	defer xa.Release()
	xb.Release()

	rec := array.NewRecord(schema, []arrow.Array{xa}, 1)
	defer rec.Release()

	df := NewDataFrame(rec)
	result := df.ParseDateColumn("x", "date")
	assert.Error(t, result.Err()) // Not a string column
}

func TestInferDateFormat(t *testing.T) {
	assert.Equal(t, "2006-01-02", inferDateFormat("2024-01-15"))
	assert.Equal(t, "01/02/2006", inferDateFormat("01/15/2024"))
	assert.Equal(t, "2006-01-02 15:04:05", inferDateFormat("2024-01-15 10:30:00"))
	assert.Equal(t, "", inferDateFormat("not a date"))
}
