package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSeriesGetterMethods tests Series getter methods
func TestSeriesGetterMethods(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Test GetInt64
	t.Run("GetInt64", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{100, 200, 300}, []bool{true, false, true})
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "int_col", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(int64Array, field)
		defer series.Release()

		// Test valid value
		val, err := series.GetInt64(0)
		assert.NoError(t, err)
		assert.Equal(t, int64(100), val)

		// Test null value - should return error
		_, err = series.GetInt64(1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "null value")

		// Test out of bounds
		_, err = series.GetInt64(10)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "index out of range")
	})

	// Test GetFloat64
	t.Run("GetFloat64", func(t *testing.T) {
		float64Builder := array.NewFloat64Builder(pool)
		float64Builder.AppendValues([]float64{1.5, 2.5, 3.5}, []bool{true, false, true})
		float64Array := float64Builder.NewArray()
		defer float64Array.Release()

		field := arrow.Field{Name: "float_col", Type: arrow.PrimitiveTypes.Float64}
		series := NewSeries(float64Array, field)
		defer series.Release()

		// Test valid value
		val, err := series.GetFloat64(0)
		assert.NoError(t, err)
		assert.Equal(t, 1.5, val)

		// Test null value - should return error
		_, err = series.GetFloat64(1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "null value")

		// Test out of bounds
		_, err = series.GetFloat64(10)
		assert.Error(t, err)
	})

	// Test GetString
	t.Run("GetString", func(t *testing.T) {
		stringBuilder := array.NewStringBuilder(pool)
		stringBuilder.AppendValues([]string{"hello", "world", "test"}, []bool{true, false, true})
		stringArray := stringBuilder.NewArray()
		defer stringArray.Release()

		field := arrow.Field{Name: "string_col", Type: arrow.BinaryTypes.String}
		series := NewSeries(stringArray, field)
		defer series.Release()

		// Test valid value
		val, err := series.GetString(0)
		assert.NoError(t, err)
		assert.Equal(t, "hello", val)

		// Test null value
		val, err = series.GetString(1)
		assert.NoError(t, err)
		assert.Equal(t, "", val) // null returns empty string

		// Test out of bounds
		_, err = series.GetString(10)
		assert.Error(t, err)
	})

	// Test GetBool
	t.Run("GetBool", func(t *testing.T) {
		boolBuilder := array.NewBooleanBuilder(pool)
		boolBuilder.AppendValues([]bool{true, false, true}, []bool{true, true, false})
		boolArray := boolBuilder.NewArray()
		defer boolArray.Release()

		field := arrow.Field{Name: "bool_col", Type: arrow.FixedWidthTypes.Boolean}
		series := NewSeries(boolArray, field)
		defer series.Release()

		// Test valid value
		val, err := series.GetBool(0)
		assert.NoError(t, err)
		assert.Equal(t, true, val)

		// Test valid false value
		val, err = series.GetBool(1)
		assert.NoError(t, err)
		assert.Equal(t, false, val)

		// Test null value - should return error
		_, err = series.GetBool(2)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "null value")

		// Test out of bounds
		_, err = series.GetBool(10)
		assert.Error(t, err)
	})

	// Test GetValue with different types
	t.Run("GetValue", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{42}, nil)
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "value_col", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(int64Array, field)
		defer series.Release()

		// Test valid index
		val := series.GetValue(0)
		assert.Equal(t, int64(42), val)

		// Test out of bounds (returns nil)
		val = series.GetValue(10)
		assert.Nil(t, val)
	})
}

// TestSeriesUtilityMethods tests Series utility methods
func TestSeriesUtilityMethods(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Test Equal
	t.Run("Equal", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{1, 2, 3}, nil)
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "col1", Type: arrow.PrimitiveTypes.Int64}
		series1 := NewSeries(int64Array, field)
		defer series1.Release()

		// Test equal series
		series2 := NewSeries(int64Array, field)
		defer series2.Release()
		assert.True(t, series1.Equal(series2), "Equal series should return true")

		// Test different names
		field3 := arrow.Field{Name: "col2", Type: arrow.PrimitiveTypes.Int64}
		series3 := NewSeries(int64Array, field3)
		defer series3.Release()
		assert.False(t, series1.Equal(series3), "Different names should return false")

		// Test different lengths
		int64Builder2 := array.NewInt64Builder(pool)
		int64Builder2.AppendValues([]int64{1, 2}, nil)
		int64Array2 := int64Builder2.NewArray()
		defer int64Array2.Release()

		series4 := NewSeries(int64Array2, field)
		defer series4.Release()
		assert.False(t, series1.Equal(series4), "Different lengths should return false")
	})

	// Test Validate
	t.Run("Validate", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{1, 2}, nil)
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "col1", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(int64Array, field)
		defer series.Release()

		// Test validation passes
		err := series.Validate()
		assert.NoError(t, err)

		// Test with type mismatch (field says float64 but array is int64)
		mismatchField := arrow.Field{Name: "col1", Type: arrow.PrimitiveTypes.Float64}
		series2 := &Series{
			array: int64Array,
			field: mismatchField,
		}
		int64Array.Retain() // Increment ref count since we're reusing the array
		err = series2.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "does not match")
		series2.Release()
	})

	// Test Slice
	t.Run("Slice", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "col1", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(int64Array, field)
		defer series.Release()

		// Test valid slice (offset=1, length=3)
		sliced, err := series.Slice(1, 3)
		require.NoError(t, err)
		assert.Equal(t, 3, sliced.Len())
		sliced.Release()

		// Test negative length
		_, err = series.Slice(1, -1)
		assert.Error(t, err)

		// Test out of bounds
		_, err = series.Slice(0, 10)
		assert.Error(t, err)
	})

	// Test Head
	t.Run("Head", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "col1", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(int64Array, field)
		defer series.Release()

		// Test head with n < length
		head, err := series.Head(3)
		require.NoError(t, err)
		assert.Equal(t, 3, head.Len())
		head.Release()

		// Test head with n > length (should return all)
		head2, err := series.Head(10)
		require.NoError(t, err)
		assert.Equal(t, 5, head2.Len())
		head2.Release()

		// Test head with n = 0 (should return empty, no error)
		head3, err := series.Head(0)
		require.NoError(t, err)
		assert.Equal(t, 0, head3.Len())
		head3.Release()

		// Test head with negative n (should error)
		_, err = series.Head(-1)
		assert.Error(t, err)
	})

	// Test Tail
	t.Run("Tail", func(t *testing.T) {
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
		int64Array := int64Builder.NewArray()
		defer int64Array.Release()

		field := arrow.Field{Name: "col1", Type: arrow.PrimitiveTypes.Int64}
		series := NewSeries(int64Array, field)
		defer series.Release()

		// Test tail with n < length
		tail, err := series.Tail(3)
		require.NoError(t, err)
		assert.Equal(t, 3, tail.Len())
		tail.Release()

		// Test tail with n > length (should return all)
		tail2, err := series.Tail(10)
		require.NoError(t, err)
		assert.Equal(t, 5, tail2.Len())
		tail2.Release()

		// Test tail with n = 0 (should return empty, no error)
		tail3, err := series.Tail(0)
		require.NoError(t, err)
		assert.Equal(t, 0, tail3.Len())
		tail3.Release()

		// Test tail with negative n (should error)
		_, err = series.Tail(-1)
		assert.Error(t, err)
	})
}

// TestSortMultipleEdgeCases tests SortMultiple with various scenarios
func TestSortMultipleEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	categoryBuilder.AppendValues([]string{"A", "B", "A", "B"}, nil)
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewInt64Builder(pool)
	valueBuilder.AppendValues([]int64{30, 10, 20, 40}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 4)
	defer record.Release()

	df := &DataFrame{
		record:    record,
		allocator: pool,
	}

	// Test sort by multiple columns
	sortKeys := []SortKey{
		{Column: "category", Ascending: true},
		{Column: "value", Ascending: false},
	}
	sorted, err := df.SortMultiple(sortKeys)
	require.NoError(t, err)
	defer sorted.Release()

	// Verify result has correct number of rows
	assert.Equal(t, int64(4), sorted.NumRows())

	// Test with empty sort keys
	_, err = df.SortMultiple([]SortKey{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no sort keys provided")

	// Test with invalid column
	invalidKeys := []SortKey{{Column: "nonexistent", Ascending: true}}
	_, err = df.SortMultiple(invalidKeys)
	assert.Error(t, err)
}
