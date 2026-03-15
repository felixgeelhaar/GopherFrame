package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScalarUDF_Multiply(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "a", Type: arrow.PrimitiveTypes.Float64},
		{Name: "b", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	ab := array.NewFloat64Builder(pool)
	defer ab.Release()
	ab.AppendValues([]float64{1, 2, 3}, nil)
	aArr := ab.NewArray()
	defer aArr.Release()

	bb := array.NewFloat64Builder(pool)
	defer bb.Release()
	bb.AppendValues([]float64{10, 20, 30}, nil)
	bArr := bb.NewArray()
	defer bArr.Release()

	record := array.NewRecord(schema, []arrow.Array{aArr, bArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)

	udf := ScalarUDF([]string{"a", "b"}, arrow.PrimitiveTypes.Float64, func(row map[string]interface{}) (interface{}, error) {
		a := row["a"].(float64)
		b := row["b"].(float64)
		return a * b, nil
	})

	result := df.WithColumn("product", udf)
	require.NoError(t, result.Err())
	assert.Equal(t, int64(3), result.NumCols())
	assert.True(t, result.HasColumn("product"))
}

func TestScalarUDF_StringConcat(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "first", Type: arrow.BinaryTypes.String},
		{Name: "last", Type: arrow.BinaryTypes.String},
	}, nil)

	fb := array.NewStringBuilder(pool)
	defer fb.Release()
	fb.AppendValues([]string{"Alice", "Bob"}, nil)
	fArr := fb.NewArray()
	defer fArr.Release()

	lb := array.NewStringBuilder(pool)
	defer lb.Release()
	lb.AppendValues([]string{"Smith", "Jones"}, nil)
	lArr := lb.NewArray()
	defer lArr.Release()

	record := array.NewRecord(schema, []arrow.Array{fArr, lArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)

	udf := ScalarUDF([]string{"first", "last"}, arrow.BinaryTypes.String, func(row map[string]interface{}) (interface{}, error) {
		return row["first"].(string) + " " + row["last"].(string), nil
	})

	result := df.WithColumn("full_name", udf)
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("full_name"))
}

func TestVectorUDF_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	xb := array.NewFloat64Builder(pool)
	defer xb.Release()
	xb.AppendValues([]float64{1, 4, 9, 16}, nil)
	xArr := xb.NewArray()
	defer xArr.Release()

	record := array.NewRecord(schema, []arrow.Array{xArr}, 4)
	defer record.Release()

	df := NewDataFrame(record)

	udf := VectorUDF([]string{"x"}, arrow.PrimitiveTypes.Float64, func(cols map[string]arrow.Array) (arrow.Array, error) {
		xArr := cols["x"].(*array.Float64)
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()
		for i := 0; i < xArr.Len(); i++ {
			builder.Append(xArr.Value(i) * 2)
		}
		return builder.NewArray(), nil
	})

	result := df.WithColumn("x_doubled", udf)
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("x_doubled"))
}

func TestScalarUDF_WithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	vb := array.NewFloat64Builder(pool)
	defer vb.Release()
	vb.Append(1.0)
	vb.AppendNull()
	vb.Append(3.0)
	vArr := vb.NewArray()
	defer vArr.Release()

	record := array.NewRecord(schema, []arrow.Array{vArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)

	udf := ScalarUDF([]string{"val"}, arrow.PrimitiveTypes.Float64, func(row map[string]interface{}) (interface{}, error) {
		if row["val"] == nil {
			return nil, nil
		}
		return row["val"].(float64) * 10, nil
	})

	result := df.WithColumn("scaled", udf)
	require.NoError(t, result.Err())
}

func TestScalarUDF_BoolOutput(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "age", Type: arrow.PrimitiveTypes.Int64},
	}, nil)

	ab := array.NewInt64Builder(pool)
	defer ab.Release()
	ab.AppendValues([]int64{15, 21, 17, 30}, nil)
	aArr := ab.NewArray()
	defer aArr.Release()

	record := array.NewRecord(schema, []arrow.Array{aArr}, 4)
	defer record.Release()

	df := NewDataFrame(record)

	udf := ScalarUDF([]string{"age"}, arrow.FixedWidthTypes.Boolean, func(row map[string]interface{}) (interface{}, error) {
		return row["age"].(int64) >= 18, nil
	})

	result := df.WithColumn("is_adult", udf)
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("is_adult"))
}
