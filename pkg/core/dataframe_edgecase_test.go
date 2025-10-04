package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDataFrame_EqualEdgeCases tests Equal method with various edge cases
func TestDataFrame_EqualEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create base DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{10.5, 20.5, 30.5}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 3)
	defer record.Release()

	df1 := NewDataFrame(record)
	defer df1.Release()

	// Test: Different number of columns
	schema2 := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder2 := array.NewInt64Builder(pool)
	idBuilder2.AppendValues([]int64{1, 2, 3}, nil)
	idArray2 := idBuilder2.NewArray()
	defer idArray2.Release()

	record2 := array.NewRecord(schema2, []arrow.Array{idArray2}, 3)
	defer record2.Release()

	df2 := NewDataFrame(record2)
	defer df2.Release()

	assert.False(t, df1.Equal(df2), "DataFrames with different column counts should not be equal")

	// Test: Same structure, different data
	idBuilder3 := array.NewInt64Builder(pool)
	idBuilder3.AppendValues([]int64{4, 5, 6}, nil)
	idArray3 := idBuilder3.NewArray()
	defer idArray3.Release()

	valueBuilder3 := array.NewFloat64Builder(pool)
	valueBuilder3.AppendValues([]float64{40.5, 50.5, 60.5}, nil)
	valueArray3 := valueBuilder3.NewArray()
	defer valueArray3.Release()

	record3 := array.NewRecord(schema, []arrow.Array{idArray3, valueArray3}, 3)
	defer record3.Release()

	df3 := NewDataFrame(record3)
	defer df3.Release()

	assert.False(t, df1.Equal(df3), "DataFrames with different data should not be equal")
}

// TestDataFrame_ValidateEdgeCases tests Validate method edge cases
func TestDataFrame_ValidateEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Valid DataFrame
	err := df.Validate()
	assert.NoError(t, err, "Valid DataFrame should pass validation")

	// Test: DataFrame with nil allocator (should still validate)
	df.allocator = nil
	err = df.Validate()
	assert.NoError(t, err, "DataFrame with nil allocator should still validate")
}

// TestDataFrame_StringEdgeCases tests String method edge cases
func TestDataFrame_StringEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	str := df.String()
	assert.Contains(t, str, "DataFrame", "String should contain 'DataFrame'")
	assert.Contains(t, str, "rows: 2", "String should contain row count")
	assert.Contains(t, str, "cols: 2", "String should contain column count")
	assert.Contains(t, str, "id: type=int64", "String should contain schema info")
}

// TestDataFrame_WithColumnEdgeCases tests WithColumn edge cases
func TestDataFrame_WithColumnEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{10.5, 20.5, 30.5}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Replace existing column
	newValueBuilder := array.NewFloat64Builder(pool)
	newValueBuilder.AppendValues([]float64{100.5, 200.5, 300.5}, nil)
	newValueArray := newValueBuilder.NewArray()
	defer newValueArray.Release()

	newValueSeries := NewSeries(newValueArray, arrow.Field{Name: "value", Type: arrow.PrimitiveTypes.Float64})
	defer newValueSeries.Release()

	dfReplaced, err := df.WithColumn("value", newValueSeries.Array())
	require.NoError(t, err)
	defer dfReplaced.Release()

	replacedSeries, err := dfReplaced.Column("value")
	require.NoError(t, err)
	replacedCol := replacedSeries.Array().(*array.Float64)
	assert.Equal(t, 100.5, replacedCol.Value(0), "Replaced column should have new values")

	// Test: Add column with wrong length
	shortBuilder := array.NewInt64Builder(pool)
	shortBuilder.AppendValues([]int64{1, 2}, nil) // Only 2 values instead of 3
	shortArray := shortBuilder.NewArray()
	defer shortArray.Release()

	shortSeries := NewSeries(shortArray, arrow.Field{Name: "short", Type: arrow.PrimitiveTypes.Int64})
	defer shortSeries.Release()

	_, err = df.WithColumn("short", shortSeries.Array())
	assert.Error(t, err, "Adding column with wrong length should error")
	assert.Contains(t, err.Error(), "does not match DataFrame rows", "Error should mention row mismatch")
}

// TestDataFrame_FilterEdgeCases tests Filter edge cases
func TestDataFrame_FilterEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{10.5, 20.5, 30.5, 40.5, 50.5}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Filter with all true (all rows pass)
	trueBuilder := array.NewBooleanBuilder(pool)
	trueBuilder.AppendValues([]bool{true, true, true, true, true}, nil)
	trueMask := trueBuilder.NewArray()
	defer trueMask.Release()

	trueSeries := NewSeries(trueMask, arrow.Field{Name: "mask", Type: arrow.FixedWidthTypes.Boolean})
	defer trueSeries.Release()

	allTrue, err := df.Filter(trueSeries.Array())
	require.NoError(t, err)
	defer allTrue.Release()
	assert.Equal(t, int64(5), allTrue.NumRows(), "All true mask should keep all rows")

	// Test: Filter with all false (no rows pass)
	falseBuilder := array.NewBooleanBuilder(pool)
	falseBuilder.AppendValues([]bool{false, false, false, false, false}, nil)
	falseMask := falseBuilder.NewArray()
	defer falseMask.Release()

	falseSeries := NewSeries(falseMask, arrow.Field{Name: "mask", Type: arrow.FixedWidthTypes.Boolean})
	defer falseSeries.Release()

	allFalse, err := df.Filter(falseSeries.Array())
	require.NoError(t, err)
	defer allFalse.Release()
	assert.Equal(t, int64(0), allFalse.NumRows(), "All false mask should remove all rows")

	// Test: Filter with nulls in mask (nulls preserved)
	nullBuilder := array.NewBooleanBuilder(pool)
	nullBuilder.AppendValues([]bool{true, false, true, false, true}, []bool{true, true, false, true, true})
	nullMask := nullBuilder.NewArray()
	defer nullMask.Release()

	nullSeries := NewSeries(nullMask, arrow.Field{Name: "mask", Type: arrow.FixedWidthTypes.Boolean})
	defer nullSeries.Release()

	withNulls, err := df.Filter(nullSeries.Array())
	require.NoError(t, err)
	defer withNulls.Release()
	assert.Equal(t, int64(2), withNulls.NumRows(), "Filter with nulls should only keep true values")
}

// TestCompareValuesEdgeCases tests compareValues helper with edge cases
func TestCompareValuesEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	df := &DataFrame{
		allocator: pool,
	}

	// Test: Float64 comparison
	float64Builder := array.NewFloat64Builder(pool)
	float64Builder.AppendValues([]float64{1.5, 2.5, 1.5}, nil)
	float64Array := float64Builder.NewArray()
	defer float64Array.Release()

	assert.Equal(t, -1, df.compareValues(float64Array, arrow.PrimitiveTypes.Float64, 0, 1), "1.5 < 2.5")
	assert.Equal(t, 1, df.compareValues(float64Array, arrow.PrimitiveTypes.Float64, 1, 0), "2.5 > 1.5")
	assert.Equal(t, 0, df.compareValues(float64Array, arrow.PrimitiveTypes.Float64, 0, 2), "1.5 == 1.5")

	// Test: String comparison
	stringBuilder := array.NewStringBuilder(pool)
	stringBuilder.AppendValues([]string{"apple", "banana", "apple"}, nil)
	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	assert.Equal(t, -1, df.compareValues(stringArray, arrow.BinaryTypes.String, 0, 1), "apple < banana")
	assert.Equal(t, 1, df.compareValues(stringArray, arrow.BinaryTypes.String, 1, 0), "banana > apple")
	assert.Equal(t, 0, df.compareValues(stringArray, arrow.BinaryTypes.String, 0, 2), "apple == apple")

	// Test: Boolean comparison
	boolBuilder := array.NewBooleanBuilder(pool)
	boolBuilder.AppendValues([]bool{false, true, false}, nil)
	boolArray := boolBuilder.NewArray()
	defer boolArray.Release()

	assert.Equal(t, -1, df.compareValues(boolArray, arrow.FixedWidthTypes.Boolean, 0, 1), "false < true")
	assert.Equal(t, 1, df.compareValues(boolArray, arrow.FixedWidthTypes.Boolean, 1, 0), "true > false")
	assert.Equal(t, 0, df.compareValues(boolArray, arrow.FixedWidthTypes.Boolean, 0, 2), "false == false")

	// Test: Int64 comparison
	int64Builder := array.NewInt64Builder(pool)
	int64Builder.AppendValues([]int64{10, 20, 10}, nil)
	int64Array := int64Builder.NewArray()
	defer int64Array.Release()

	assert.Equal(t, -1, df.compareValues(int64Array, arrow.PrimitiveTypes.Int64, 0, 1), "10 < 20")
	assert.Equal(t, 1, df.compareValues(int64Array, arrow.PrimitiveTypes.Int64, 1, 0), "20 > 10")
	assert.Equal(t, 0, df.compareValues(int64Array, arrow.PrimitiveTypes.Int64, 0, 2), "10 == 10")
}

// TestSortMultipleMoreEdgeCases tests additional SortMultiple edge cases
func TestSortMultipleMoreEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	categoryBuilder.AppendValues([]string{"A", "B", "A", "B", "A"}, nil)
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewInt64Builder(pool)
	valueBuilder.AppendValues([]int64{30, 10, 20, 40, 20}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{9.5, 8.5, 9.0, 8.0, 9.2}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray, scoreArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Sort by three columns
	sortKeys := []SortKey{
		{Column: "category", Ascending: true},
		{Column: "value", Ascending: true},
		{Column: "score", Ascending: false},
	}

	sorted, err := df.SortMultiple(sortKeys)
	require.NoError(t, err)
	defer sorted.Release()

	assert.Equal(t, int64(5), sorted.NumRows(), "Sorted DataFrame should have same row count")

	// Verify category column is sorted
	categorySeries, err := sorted.Column("category")
	require.NoError(t, err)
	categoryCol := categorySeries.Array().(*array.String)
	assert.Equal(t, "A", categoryCol.Value(0))
	assert.Equal(t, "A", categoryCol.Value(1))
	assert.Equal(t, "A", categoryCol.Value(2))
}
