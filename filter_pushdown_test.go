package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryPlan_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "a", Type: arrow.PrimitiveTypes.Float64},
		{Name: "b", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	ab := array.NewFloat64Builder(pool)
	ab.AppendValues([]float64{1, 2, 3, 4, 5}, nil)
	aa := ab.NewArray()
	defer aa.Release()
	ab.Release()
	bb := array.NewFloat64Builder(pool)
	bb.AppendValues([]float64{10, 20, 30, 40, 50}, nil)
	ba := bb.NewArray()
	defer ba.Release()
	bb.Release()

	rec := array.NewRecord(schema, []arrow.Array{aa, ba}, 5)
	defer rec.Release()
	df := NewDataFrame(rec)

	// Without optimization, this would: select -> filter -> with_column
	// With pushdown: filter -> select -> with_column
	result := NewQueryPlan(df).
		Filter(df.Col("a").Gt(Lit(2.0))).
		Select("a", "b").
		Execute()

	require.NoError(t, result.Err())
	assert.Equal(t, int64(3), result.NumRows()) // 3, 4, 5
}

func TestQueryPlan_FilterPushdown(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
		{Name: "y", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.AppendValues([]float64{1, 2, 3, 4}, nil)
	xa := xb.NewArray()
	defer xa.Release()
	xb.Release()
	yb := array.NewFloat64Builder(pool)
	yb.AppendValues([]float64{10, 20, 30, 40}, nil)
	ya := yb.NewArray()
	defer ya.Release()
	yb.Release()

	rec := array.NewRecord(schema, []arrow.Array{xa, ya}, 4)
	defer rec.Release()
	df := NewDataFrame(rec)

	// Filter on "x" should be pushed before WithColumn("z")
	result := NewQueryPlan(df).
		WithColumn("z", df.Col("x").Add(df.Col("y"))).
		Filter(df.Col("x").Gt(Lit(2.0))).
		Execute()

	require.NoError(t, result.Err())
	assert.Equal(t, int64(2), result.NumRows()) // x=3,4
	assert.True(t, result.HasColumn("z"))
}

func TestPushFiltersDown_Logic(t *testing.T) {
	ops := []queryOp{
		{opType: "with_column", colName: "z"},
		{opType: "filter", predicate: Col("x").Gt(Lit(0.0))},
	}

	optimized := pushFiltersDown(ops)
	// Filter on "x" should be pushed before with_column "z"
	assert.Equal(t, "filter", optimized[0].opType)
	assert.Equal(t, "with_column", optimized[1].opType)
}

func TestPushFiltersDown_NoPushIfDependsOnComputed(t *testing.T) {
	ops := []queryOp{
		{opType: "with_column", colName: "computed"},
		{opType: "filter", predicate: Col("computed").Gt(Lit(0.0))},
	}

	optimized := pushFiltersDown(ops)
	// Filter references "computed", so it cannot be pushed before with_column
	assert.Equal(t, "with_column", optimized[0].opType)
	assert.Equal(t, "filter", optimized[1].opType)
}
