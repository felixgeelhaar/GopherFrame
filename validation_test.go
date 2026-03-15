package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createValidationTestDF(t *testing.T) *DataFrame {
	t.Helper()
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "age", Type: arrow.PrimitiveTypes.Int64},
		{Name: "score", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"Alice", "Bob", "Carol"}, nil)
	nArr := nb.NewArray()
	defer nArr.Release()
	nb.Release()

	ab := array.NewInt64Builder(pool)
	ab.AppendValues([]int64{25, 30, 35}, nil)
	aArr := ab.NewArray()
	defer aArr.Release()
	ab.Release()

	sb := array.NewFloat64Builder(pool)
	sb.AppendValues([]float64{95.5, 87.3, 92.1}, nil)
	sArr := sb.NewArray()
	defer sArr.Release()
	sb.Release()

	record := array.NewRecord(schema, []arrow.Array{nArr, aArr, sArr}, 3)
	defer record.Release()

	return NewDataFrame(record)
}

func TestDescribe_Basic(t *testing.T) {
	df := createValidationTestDF(t)

	stats, err := df.Describe()
	require.NoError(t, err)
	assert.Len(t, stats, 3)

	// Check name column
	assert.Equal(t, "name", stats[0].Name)
	assert.Equal(t, int64(3), stats[0].Count)
	assert.Equal(t, int64(0), stats[0].NullCount)
	assert.Equal(t, int64(3), stats[0].Unique)

	// Check score column has numeric stats
	assert.True(t, stats[2].HasNumeric)
	assert.InDelta(t, 91.63, stats[2].Mean, 0.1)
}

func TestDescribeString(t *testing.T) {
	df := createValidationTestDF(t)
	s := df.DescribeString()
	assert.Contains(t, s, "name")
	assert.Contains(t, s, "age")
	assert.Contains(t, s, "score")
	assert.Contains(t, s, "Column")
}

func TestNullCount_NoNulls(t *testing.T) {
	df := createValidationTestDF(t)
	nc := df.NullCount()
	assert.Equal(t, int64(0), nc["name"])
	assert.Equal(t, int64(0), nc["age"])
}

func TestNullCount_WithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	vb := array.NewFloat64Builder(pool)
	vb.Append(1.0)
	vb.AppendNull()
	vb.Append(3.0)
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{vArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	nc := df.NullCount()
	assert.Equal(t, int64(1), nc["val"])
}

func TestIsComplete_True(t *testing.T) {
	df := createValidationTestDF(t)
	assert.True(t, df.IsComplete())
}

func TestIsComplete_False(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	vb := array.NewFloat64Builder(pool)
	vb.Append(1.0)
	vb.AppendNull()
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{vArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	assert.False(t, df.IsComplete())
}

func TestValidate_NotNull_Pass(t *testing.T) {
	df := createValidationTestDF(t)
	result := df.Validate(NotNull("name"))
	assert.True(t, result.Valid)
	assert.Empty(t, result.Violations)
}

func TestValidate_NotNull_Fail(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	vb := array.NewFloat64Builder(pool)
	vb.Append(1.0)
	vb.AppendNull()
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{vArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Validate(NotNull("val"))
	assert.False(t, result.Valid)
	assert.Len(t, result.Violations, 1)
}

func TestValidate_Positive_Pass(t *testing.T) {
	df := createValidationTestDF(t)
	result := df.Validate(Positive("age"))
	assert.True(t, result.Valid)
}

func TestValidate_Positive_Fail(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	vb := array.NewFloat64Builder(pool)
	vb.AppendValues([]float64{1.0, -5.0, 3.0}, nil)
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{vArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Validate(Positive("val"))
	assert.False(t, result.Valid)
}

func TestValidate_InRange(t *testing.T) {
	df := createValidationTestDF(t)
	result := df.Validate(InRange("score", 0, 100))
	assert.True(t, result.Valid)

	result = df.Validate(InRange("score", 90, 100))
	assert.False(t, result.Valid) // 87.3 is out of [90, 100]
}

func TestValidate_Unique_Pass(t *testing.T) {
	df := createValidationTestDF(t)
	result := df.Validate(UniqueValues("name"))
	assert.True(t, result.Valid)
}

func TestValidate_Unique_Fail(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "val", Type: arrow.BinaryTypes.String},
	}, nil)
	vb := array.NewStringBuilder(pool)
	vb.AppendValues([]string{"a", "b", "a"}, nil)
	vArr := vb.NewArray()
	defer vArr.Release()
	vb.Release()

	record := array.NewRecord(schema, []arrow.Array{vArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	result := df.Validate(UniqueValues("val"))
	assert.False(t, result.Valid)
}

func TestValidate_ColumnNotFound(t *testing.T) {
	df := createValidationTestDF(t)
	result := df.Validate(NotNull("nonexistent"))
	assert.False(t, result.Valid)
}

func TestValidate_MultipleRules(t *testing.T) {
	df := createValidationTestDF(t)
	result := df.Validate(
		NotNull("name"),
		Positive("age"),
		InRange("score", 0, 100),
		UniqueValues("name"),
	)
	assert.True(t, result.Valid)
	assert.Empty(t, result.Violations)
}
