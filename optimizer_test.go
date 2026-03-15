package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFoldedLit_Float64(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.AppendValues([]float64{1, 2, 3}, nil)
	xArr := xb.NewArray()
	defer xArr.Release()
	xb.Release()

	record := array.NewRecord(schema, []arrow.Array{xArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)

	// Use FoldedLit as a pre-computed constant
	result := df.WithColumn("const", FoldedLit(42.0))
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("const"))
	assert.Equal(t, int64(3), result.NumRows())
}

func TestFoldedLit_String(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.AppendValues([]float64{1, 2}, nil)
	xArr := xb.NewArray()
	defer xArr.Release()
	xb.Release()

	record := array.NewRecord(schema, []arrow.Array{xArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.WithColumn("label", FoldedLit("constant"))
	require.NoError(t, result.Err())
	assert.True(t, result.HasColumn("label"))
}

func TestOptimize_PassThrough(t *testing.T) {
	// Optimize should return the expression unchanged if no optimization applies
	e := Col("x").Add(Lit(1.0))
	optimized := Optimize(e)
	assert.NotNil(t, optimized)
}

func TestFoldedLit_Chainable(t *testing.T) {
	// FoldedLit should be composable with other expressions
	e := FoldedLit(10.0).Add(Col("x"))
	assert.NotNil(t, e)
	assert.Contains(t, e.String(), "add")
}

func TestFoldedLit_AllTypes(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.Append(1.0)
	xArr := xb.NewArray()
	defer xArr.Release()
	xb.Release()

	record := array.NewRecord(schema, []arrow.Array{xArr}, 1)
	defer record.Release()
	df := NewDataFrame(record)

	// Test int64
	r1 := df.WithColumn("int_const", FoldedLit(int64(99)))
	require.NoError(t, r1.Err())

	// Test bool
	r2 := df.WithColumn("bool_const", FoldedLit(true))
	require.NoError(t, r2.Err())
}
