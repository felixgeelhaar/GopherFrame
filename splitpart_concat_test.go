package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitPart_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "path", Type: arrow.BinaryTypes.String},
	}, nil)

	pb := array.NewStringBuilder(pool)
	pb.AppendValues([]string{"a/b/c", "x/y", "single"}, nil)
	pArr := pb.NewArray()
	defer pArr.Release()
	pb.Release()

	record := array.NewRecord(schema, []arrow.Array{pArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)

	// Get second part (index 1) when splitting by "/"
	result := df.WithColumn("part", df.Col("path").SplitPart(Lit("/"), Lit(int64(1))))
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("part"))
}

func TestSplitPart_OutOfBounds(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String},
	}, nil)

	tb := array.NewStringBuilder(pool)
	tb.AppendValues([]string{"a,b"}, nil)
	tArr := tb.NewArray()
	defer tArr.Release()
	tb.Release()

	record := array.NewRecord(schema, []arrow.Array{tArr}, 1)
	defer record.Release()

	df := NewDataFrame(record)

	// Index 5 is out of bounds for "a,b" split by "," (only 2 parts)
	result := df.WithColumn("part", df.Col("text").SplitPart(Lit(","), Lit(int64(5))))
	require.NoError(t, result.Err())
}

func TestConcatAgg_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "group", Type: arrow.BinaryTypes.String},
		{Name: "tag", Type: arrow.BinaryTypes.String},
	}, nil)

	gb := array.NewStringBuilder(pool)
	gb.AppendValues([]string{"A", "A", "B", "B", "B"}, nil)
	gArr := gb.NewArray()
	defer gArr.Release()
	gb.Release()

	tb := array.NewStringBuilder(pool)
	tb.AppendValues([]string{"red", "blue", "green", "yellow", "purple"}, nil)
	tArr := tb.NewArray()
	defer tArr.Release()
	tb.Release()

	record := array.NewRecord(schema, []arrow.Array{gArr, tArr}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.GroupBy("group").Agg(ConcatAgg("tag", ", ").As("tags"))

	require.NoError(t, result.Err())
	assert.Equal(t, int64(2), result.NumRows())
	assert.True(t, result.HasColumn("tags"))
}

func TestCustomAgg_Range(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "group", Type: arrow.BinaryTypes.String},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	gb := array.NewStringBuilder(pool)
	gb.AppendValues([]string{"A", "A", "B", "B"}, nil)
	gArr := gb.NewArray()
	defer gArr.Release()
	gb.Release()

	vb := array.NewFloat64Builder(pool)
	vb.AppendValues([]float64{10, 20, 5, 35}, nil)
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{gArr, vArr}, 4)
	defer record.Release()

	df := NewDataFrame(record)

	// Custom aggregation: range (max - min)
	rangeFn := func(values []float64) float64 {
		if len(values) == 0 {
			return 0
		}
		min, max := values[0], values[0]
		for _, v := range values[1:] {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		return max - min
	}

	result := df.GroupBy("group").Agg(CustomAgg("value", "range", rangeFn))
	require.NoError(t, result.Err())
	assert.Equal(t, int64(2), result.NumRows())
	assert.True(t, result.HasColumn("range"))
}

func TestExprSplitPart_ViaExpr(t *testing.T) {
	// Test that SplitPart works through the expr package directly
	e := expr.NewColumnExpr("path")
	sp := e.SplitPart(expr.NewLiteralExpr("/"), expr.NewLiteralExpr(int64(0)))
	assert.NotNil(t, sp)
	assert.Contains(t, sp.String(), "split_part")
}
