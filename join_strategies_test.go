package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createSortedJoinDFs(t *testing.T) (*DataFrame, *DataFrame) {
	t.Helper()
	pool := memory.NewGoAllocator()

	// Left (sorted by id)
	ls := arrow.NewSchema([]arrow.Field{
		{Name: "id", Type: arrow.BinaryTypes.String},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}, nil)
	lb1 := array.NewStringBuilder(pool)
	lb1.AppendValues([]string{"1", "2", "3", "4"}, nil)
	la1 := lb1.NewArray()
	lb1.Release()
	lb2 := array.NewStringBuilder(pool)
	lb2.AppendValues([]string{"Alice", "Bob", "Carol", "Dave"}, nil)
	la2 := lb2.NewArray()
	lb2.Release()
	lr := array.NewRecord(ls, []arrow.Array{la1, la2}, 4)

	// Right (sorted by uid)
	rs := arrow.NewSchema([]arrow.Field{
		{Name: "uid", Type: arrow.BinaryTypes.String},
		{Name: "score", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	rb1 := array.NewStringBuilder(pool)
	rb1.AppendValues([]string{"1", "2", "4"}, nil)
	ra1 := rb1.NewArray()
	rb1.Release()
	rb2 := array.NewFloat64Builder(pool)
	rb2.AppendValues([]float64{95, 87, 72}, nil)
	ra2 := rb2.NewArray()
	rb2.Release()
	rr := array.NewRecord(rs, []arrow.Array{ra1, ra2}, 3)

	return NewDataFrame(lr), NewDataFrame(rr)
}

func TestMergeJoin_Basic(t *testing.T) {
	left, right := createSortedJoinDFs(t)
	result := left.MergeJoin(right, "id", "uid")
	require.NoError(t, result.Err())
	assert.Equal(t, int64(3), result.NumRows()) // 1,2,4 match; 3 doesn't
	assert.True(t, result.HasColumn("name"))
	assert.True(t, result.HasColumn("score"))
}

func TestBroadcastJoin_Basic(t *testing.T) {
	left, right := createSortedJoinDFs(t)
	result := left.BroadcastJoin(right, "id", "uid")
	require.NoError(t, result.Err())
	assert.Equal(t, int64(3), result.NumRows())
}

func TestAutoJoin_SmallRight(t *testing.T) {
	left, right := createSortedJoinDFs(t)
	result := left.AutoJoin(right, "id", "uid")
	require.NoError(t, result.Err())
	assert.Equal(t, int64(3), result.NumRows())
}

func TestCrossTab_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "dept", Type: arrow.BinaryTypes.String},
		{Name: "status", Type: arrow.BinaryTypes.String},
	}, nil)
	db := array.NewStringBuilder(pool)
	db.AppendValues([]string{"eng", "eng", "sales", "sales", "eng"}, nil)
	da := db.NewArray()
	db.Release()
	sb := array.NewStringBuilder(pool)
	sb.AppendValues([]string{"active", "inactive", "active", "active", "active"}, nil)
	sa := sb.NewArray()
	sb.Release()

	rec := array.NewRecord(schema, []arrow.Array{da, sa}, 5)
	df := NewDataFrame(rec)

	result := df.CrossTab("dept", "status")
	require.NoError(t, result.Err())
	assert.Equal(t, int64(2), result.NumRows()) // eng, sales
	assert.True(t, result.HasColumn("active"))
	assert.True(t, result.HasColumn("inactive"))
}

func TestDetectOutliersIQR_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	vb := array.NewFloat64Builder(pool)
	vb.AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, nil)
	va := vb.NewArray()
	vb.Release()

	rec := array.NewRecord(schema, []arrow.Array{va}, 10)
	df := NewDataFrame(rec)

	result, err := df.DetectOutliersIQR("val", 1.5)
	require.NoError(t, err)
	assert.Greater(t, result.Count, 0) // 100 should be an outlier
	assert.Equal(t, "IQR", result.Method)
}

func TestDetectOutliersZScore_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	vb := array.NewFloat64Builder(pool)
	vb.AppendValues([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, nil)
	va := vb.NewArray()
	vb.Release()

	rec := array.NewRecord(schema, []arrow.Array{va}, 10)
	df := NewDataFrame(rec)

	result, err := df.DetectOutliersZScore("val", 2.0)
	require.NoError(t, err)
	assert.Greater(t, result.Count, 0)
	assert.Equal(t, "z-score", result.Method)
}

func TestDetectOutliers_NonNumeric(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
	}, nil)
	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"a", "b"}, nil)
	na := nb.NewArray()
	nb.Release()

	rec := array.NewRecord(schema, []arrow.Array{na}, 2)
	df := NewDataFrame(rec)

	_, err := df.DetectOutliersIQR("name", 1.5)
	assert.Error(t, err)
}

func TestMergeJoin_InvalidColumn(t *testing.T) {
	left, right := createSortedJoinDFs(t)
	result := left.MergeJoin(right, "nonexistent", "uid")
	assert.Error(t, result.Err())
}

func TestBroadcastJoin_NilOther(t *testing.T) {
	left, _ := createSortedJoinDFs(t)
	result := left.BroadcastJoin(nil, "id", "uid")
	assert.Error(t, result.Err())
}
