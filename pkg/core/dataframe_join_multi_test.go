package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMultiKeyTestDataFrames(t *testing.T) (*DataFrame, *DataFrame) {
	t.Helper()
	pool := memory.NewGoAllocator()

	// Left: department + role -> name
	leftSchema := arrow.NewSchema([]arrow.Field{
		{Name: "dept", Type: arrow.BinaryTypes.String},
		{Name: "role", Type: arrow.BinaryTypes.String},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}, nil)

	d1 := array.NewStringBuilder(pool)
	d1.AppendValues([]string{"eng", "eng", "sales", "sales"}, nil)
	d1Arr := d1.NewArray()
	d1.Release()

	r1 := array.NewStringBuilder(pool)
	r1.AppendValues([]string{"lead", "dev", "lead", "rep"}, nil)
	r1Arr := r1.NewArray()
	r1.Release()

	n1 := array.NewStringBuilder(pool)
	n1.AppendValues([]string{"Alice", "Bob", "Carol", "Dave"}, nil)
	n1Arr := n1.NewArray()
	n1.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{d1Arr, r1Arr, n1Arr}, 4)

	// Right: department + role -> salary
	rightSchema := arrow.NewSchema([]arrow.Field{
		{Name: "department", Type: arrow.BinaryTypes.String},
		{Name: "position", Type: arrow.BinaryTypes.String},
		{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	d2 := array.NewStringBuilder(pool)
	d2.AppendValues([]string{"eng", "eng", "sales"}, nil)
	d2Arr := d2.NewArray()
	d2.Release()

	r2 := array.NewStringBuilder(pool)
	r2.AppendValues([]string{"lead", "dev", "lead"}, nil)
	r2Arr := r2.NewArray()
	r2.Release()

	s2 := array.NewFloat64Builder(pool)
	s2.AppendValues([]float64{150000, 120000, 130000}, nil)
	s2Arr := s2.NewArray()
	s2.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{d2Arr, r2Arr, s2Arr}, 3)

	return NewDataFrame(leftRecord), NewDataFrame(rightRecord)
}

func TestInnerJoinMulti_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	result, err := left.InnerJoinMulti(right, []string{"dept", "role"}, []string{"department", "position"})
	require.NoError(t, err)
	defer result.Release()

	// eng/lead, eng/dev, sales/lead match; sales/rep does not
	assert.Equal(t, int64(3), result.NumRows())
	assert.True(t, result.HasColumn("salary"))
}

func TestLeftJoinMulti_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	result, err := left.LeftJoinMulti(right, []string{"dept", "role"}, []string{"department", "position"})
	require.NoError(t, err)
	defer result.Release()

	// All 4 left rows preserved; sales/rep has null salary
	assert.Equal(t, int64(4), result.NumRows())
}

func TestRightJoinMulti_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	result, err := left.RightJoinMulti(right, []string{"dept", "role"}, []string{"department", "position"})
	require.NoError(t, err)
	defer result.Release()

	// All 3 right rows preserved
	assert.Equal(t, int64(3), result.NumRows())
}

func TestFullOuterJoinMulti_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	result, err := left.FullOuterJoinMulti(right, []string{"dept", "role"}, []string{"department", "position"})
	require.NoError(t, err)
	defer result.Release()

	// 3 matched + 1 left-only (sales/rep) = 4
	assert.Equal(t, int64(4), result.NumRows())
}

func TestJoinMulti_KeyMismatch_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	_, err := left.JoinMulti(right, []string{"dept"}, []string{"department", "position"}, InnerJoin)
	assert.Error(t, err)
}

func TestJoinMulti_EmptyKeys_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	_, err := left.JoinMulti(right, []string{}, []string{}, InnerJoin)
	assert.Error(t, err)
}

func TestJoinMulti_InvalidColumn_Core(t *testing.T) {
	left, right := createMultiKeyTestDataFrames(t)
	defer left.Release()
	defer right.Release()

	_, err := left.JoinMulti(right, []string{"nonexistent"}, []string{"department"}, InnerJoin)
	assert.Error(t, err)
}
