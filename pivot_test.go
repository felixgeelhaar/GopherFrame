package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPivot_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "metric", Type: arrow.BinaryTypes.String},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"Alice", "Alice", "Bob", "Bob"}, nil)
	nArr := nb.NewArray()
	defer nArr.Release()
	nb.Release()

	mb := array.NewStringBuilder(pool)
	mb.AppendValues([]string{"height", "weight", "height", "weight"}, nil)
	mArr := mb.NewArray()
	defer mArr.Release()
	mb.Release()

	vb := array.NewFloat64Builder(pool)
	vb.AppendValues([]float64{165, 55, 180, 80}, nil)
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{nArr, mArr, vArr}, 4)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Pivot([]string{"name"}, "metric", "value")

	require.NoError(t, result.Err())
	assert.Equal(t, int64(2), result.NumRows())
	assert.True(t, result.HasColumn("name"))
	assert.True(t, result.HasColumn("height"))
	assert.True(t, result.HasColumn("weight"))
}

func TestPivot_MissingValues(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "metric", Type: arrow.BinaryTypes.String},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"Alice", "Bob"}, nil)
	nArr := nb.NewArray()
	defer nArr.Release()
	nb.Release()

	mb := array.NewStringBuilder(pool)
	mb.AppendValues([]string{"height", "weight"}, nil)
	mArr := mb.NewArray()
	defer mArr.Release()
	mb.Release()

	vb := array.NewFloat64Builder(pool)
	vb.AppendValues([]float64{165, 80}, nil)
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{nArr, mArr, vArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Pivot([]string{"name"}, "metric", "value")

	require.NoError(t, result.Err())
	assert.Equal(t, int64(2), result.NumRows())
	// Alice has height but not weight, Bob has weight but not height
}

func TestPivot_InvalidColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "a", Type: arrow.BinaryTypes.String},
	}, nil)
	ab := array.NewStringBuilder(pool)
	ab.Append("x")
	aArr := ab.NewArray()
	defer aArr.Release()
	ab.Release()

	record := array.NewRecord(schema, []arrow.Array{aArr}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Pivot([]string{"a"}, "nonexistent", "a")
	assert.Error(t, result.Err())
}

func TestUnpivot_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "height", Type: arrow.PrimitiveTypes.Float64},
		{Name: "weight", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"Alice", "Bob"}, nil)
	nArr := nb.NewArray()
	defer nArr.Release()
	nb.Release()

	hb := array.NewFloat64Builder(pool)
	hb.AppendValues([]float64{165, 180}, nil)
	hArr := hb.NewArray()
	defer hArr.Release()
	hb.Release()

	wb := array.NewFloat64Builder(pool)
	wb.AppendValues([]float64{55, 80}, nil)
	wArr := wb.NewArray()
	defer wArr.Release()
	wb.Release()

	record := array.NewRecord(schema, []arrow.Array{nArr, hArr, wArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Unpivot([]string{"name"}, []string{"height", "weight"}, "metric", "value")

	require.NoError(t, result.Err())
	assert.Equal(t, int64(4), result.NumRows()) // 2 rows * 2 value cols
	assert.True(t, result.HasColumn("name"))
	assert.True(t, result.HasColumn("metric"))
	assert.True(t, result.HasColumn("value"))
}

func TestUnpivot_InvalidColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "a", Type: arrow.BinaryTypes.String},
	}, nil)
	ab := array.NewStringBuilder(pool)
	ab.Append("x")
	aArr := ab.NewArray()
	defer aArr.Release()
	ab.Release()

	record := array.NewRecord(schema, []arrow.Array{aArr}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Unpivot([]string{"a"}, []string{"nonexistent"}, "var", "val")
	assert.Error(t, result.Err())
}
