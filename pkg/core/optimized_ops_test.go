package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestOptimizedFilterCorrectness tests that optimized filter produces correct results
func TestOptimizedFilterCorrectness(t *testing.T) {
	df := createTestDataFrameForOptimization()
	defer df.Release()

	// Create filter mask: select rows where id % 2 == 0
	mask := createSelectiveMask(df.NumRows(), func(i int) bool {
		return i%2 == 0
	})
	defer mask.Release()

	// Test all filter implementations
	resultStd, err := df.Filter(mask)
	require.NoError(t, err)
	defer resultStd.Release()

	resultOpt, err := df.FilterOptimized(mask)
	require.NoError(t, err)
	defer resultOpt.Release()

	resultVec, err := df.FilterVectorized(mask)
	require.NoError(t, err)
	defer resultVec.Release()

	// All results should have same shape
	assert.Equal(t, resultStd.NumRows(), resultOpt.NumRows())
	assert.Equal(t, resultStd.NumCols(), resultOpt.NumCols())
	assert.Equal(t, resultStd.NumRows(), resultVec.NumRows())
	assert.Equal(t, resultStd.NumCols(), resultVec.NumCols())

	// Verify data correctness for each column
	for colIdx := 0; colIdx < int(df.NumCols()); colIdx++ {
		colStd := resultStd.record.Column(colIdx)
		colOpt := resultOpt.record.Column(colIdx)
		colVec := resultVec.record.Column(colIdx)

		// Test data type consistency
		assert.Equal(t, colStd.DataType(), colOpt.DataType())
		assert.Equal(t, colStd.DataType(), colVec.DataType())

		// Test array length consistency
		assert.Equal(t, colStd.Len(), colOpt.Len())
		assert.Equal(t, colStd.Len(), colVec.Len())

		// Compare actual values based on type
		switch colStd.DataType().ID() {
		case arrow.INT64:
			assertInt64ArraysEqual(t, colStd.(*array.Int64), colOpt.(*array.Int64))
			assertInt64ArraysEqual(t, colStd.(*array.Int64), colVec.(*array.Int64))
		case arrow.FLOAT64:
			assertFloat64ArraysEqual(t, colStd.(*array.Float64), colOpt.(*array.Float64))
			assertFloat64ArraysEqual(t, colStd.(*array.Float64), colVec.(*array.Float64))
		case arrow.STRING:
			assertStringArraysEqual(t, colStd.(*array.String), colOpt.(*array.String))
			assertStringArraysEqual(t, colStd.(*array.String), colVec.(*array.String))
		case arrow.BOOL:
			assertBooleanArraysEqual(t, colStd.(*array.Boolean), colOpt.(*array.Boolean))
			assertBooleanArraysEqual(t, colStd.(*array.Boolean), colVec.(*array.Boolean))
		}
	}
}

func TestOptimizedSelectCorrectness(t *testing.T) {
	df := createTestDataFrameForOptimization()
	defer df.Release()

	columns := []string{"id", "name"}

	resultStd, err := df.Select(columns)
	require.NoError(t, err)
	defer resultStd.Release()

	resultOpt, err := df.SelectOptimized(columns)
	require.NoError(t, err)
	defer resultOpt.Release()

	// Verify same structure
	assert.Equal(t, resultStd.NumRows(), resultOpt.NumRows())
	assert.Equal(t, resultStd.NumCols(), resultOpt.NumCols())
	assert.Equal(t, int64(2), resultOpt.NumCols()) // Should have 2 selected columns

	// Verify column names and data
	for colIdx := 0; colIdx < int(resultStd.NumCols()); colIdx++ {
		fieldStd := resultStd.record.Schema().Field(colIdx)
		fieldOpt := resultOpt.record.Schema().Field(colIdx)

		assert.Equal(t, fieldStd.Name, fieldOpt.Name)
		assert.Equal(t, fieldStd.Type, fieldOpt.Type)

		colStd := resultStd.record.Column(colIdx)
		colOpt := resultOpt.record.Column(colIdx)

		switch colStd.DataType().ID() {
		case arrow.INT64:
			assertInt64ArraysEqual(t, colStd.(*array.Int64), colOpt.(*array.Int64))
		case arrow.STRING:
			assertStringArraysEqual(t, colStd.(*array.String), colOpt.(*array.String))
		}
	}
}

func TestOptimizedWithColumnCorrectness(t *testing.T) {
	df := createTestDataFrameForOptimization()
	defer df.Release()

	// Create new column
	newCol := createNewTestColumn(df.NumRows())
	defer newCol.Release()

	resultStd, err := df.WithColumn("new_col", newCol)
	require.NoError(t, err)
	defer resultStd.Release()

	resultOpt, err := df.WithColumnOptimized("new_col", newCol)
	require.NoError(t, err)
	defer resultOpt.Release()

	// Verify structure
	assert.Equal(t, resultStd.NumRows(), resultOpt.NumRows())
	assert.Equal(t, resultStd.NumCols(), resultOpt.NumCols())
	assert.Equal(t, df.NumCols()+1, resultOpt.NumCols()) // Should have one more column

	// Verify all column data
	for colIdx := 0; colIdx < int(resultStd.NumCols()); colIdx++ {
		fieldStd := resultStd.record.Schema().Field(colIdx)
		fieldOpt := resultOpt.record.Schema().Field(colIdx)

		assert.Equal(t, fieldStd.Name, fieldOpt.Name)
		assert.Equal(t, fieldStd.Type, fieldOpt.Type)

		colStd := resultStd.record.Column(colIdx)
		colOpt := resultOpt.record.Column(colIdx)

		switch colStd.DataType().ID() {
		case arrow.INT64:
			assertInt64ArraysEqual(t, colStd.(*array.Int64), colOpt.(*array.Int64))
		case arrow.FLOAT64:
			assertFloat64ArraysEqual(t, colStd.(*array.Float64), colOpt.(*array.Float64))
		case arrow.STRING:
			assertStringArraysEqual(t, colStd.(*array.String), colOpt.(*array.String))
		case arrow.BOOL:
			assertBooleanArraysEqual(t, colStd.(*array.Boolean), colOpt.(*array.Boolean))
		}
	}
}

func TestOptimizedEmptyFilter(t *testing.T) {
	df := createTestDataFrameForOptimization()
	defer df.Release()

	// Create mask that selects no rows
	mask := createSelectiveMask(df.NumRows(), func(i int) bool {
		return false
	})
	defer mask.Release()

	resultStd, err := df.Filter(mask)
	require.NoError(t, err)
	defer resultStd.Release()

	resultOpt, err := df.FilterOptimized(mask)
	require.NoError(t, err)
	defer resultOpt.Release()

	resultVec, err := df.FilterVectorized(mask)
	require.NoError(t, err)
	defer resultVec.Release()

	// All should produce empty DataFrames
	assert.Equal(t, int64(0), resultStd.NumRows())
	assert.Equal(t, int64(0), resultOpt.NumRows())
	assert.Equal(t, int64(0), resultVec.NumRows())

	// But should preserve schema
	assert.Equal(t, df.NumCols(), resultStd.NumCols())
	assert.Equal(t, df.NumCols(), resultOpt.NumCols())
	assert.Equal(t, df.NumCols(), resultVec.NumCols())
}

func TestOptimizedFullFilter(t *testing.T) {
	df := createTestDataFrameForOptimization()
	defer df.Release()

	// Create mask that selects all rows
	mask := createSelectiveMask(df.NumRows(), func(i int) bool {
		return true
	})
	defer mask.Release()

	resultStd, err := df.Filter(mask)
	require.NoError(t, err)
	defer resultStd.Release()

	resultOpt, err := df.FilterOptimized(mask)
	require.NoError(t, err)
	defer resultOpt.Release()

	resultVec, err := df.FilterVectorized(mask)
	require.NoError(t, err)
	defer resultVec.Release()

	// All should preserve full DataFrame
	assert.Equal(t, df.NumRows(), resultStd.NumRows())
	assert.Equal(t, df.NumRows(), resultOpt.NumRows())
	assert.Equal(t, df.NumRows(), resultVec.NumRows())

	assert.Equal(t, df.NumCols(), resultStd.NumCols())
	assert.Equal(t, df.NumCols(), resultOpt.NumCols())
	assert.Equal(t, df.NumCols(), resultVec.NumCols())
}

func TestOptimizedWithPoolMemoryUsage(t *testing.T) {
	pool := memory.NewGoAllocator()
	df := createTestDataFrameForOptimization()
	defer df.Release()

	mask := createSelectiveMask(df.NumRows(), func(i int) bool {
		return i%3 == 0
	})
	defer mask.Release()

	newCol := createNewTestColumn(df.NumRows())
	defer newCol.Release()

	// Test FilterOptimizedWithPool
	result1, err := df.FilterOptimizedWithPool(mask, pool)
	require.NoError(t, err)
	defer result1.Release()

	// Test WithColumnOptimizedWithPool
	result2, err := df.WithColumnOptimizedWithPool("test_col", newCol, pool)
	require.NoError(t, err)
	defer result2.Release()

	// Verify results are valid
	assert.Greater(t, result1.NumRows(), int64(0))
	assert.Equal(t, df.NumCols(), result1.NumCols())

	assert.Equal(t, df.NumRows(), result2.NumRows())
	assert.Equal(t, df.NumCols()+1, result2.NumCols())
}

func TestOptimizedErrorConditions(t *testing.T) {
	df := createTestDataFrameForOptimization()
	defer df.Release()

	t.Run("FilterWithWrongMaskLength", func(t *testing.T) {
		shortMask := createSelectiveMask(df.NumRows()-1, func(i int) bool {
			return true
		})
		defer shortMask.Release()

		_, err := df.FilterOptimized(shortMask)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "predicate length")

		_, err = df.FilterVectorized(shortMask)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "predicate length")
	})

	t.Run("SelectWithInvalidColumn", func(t *testing.T) {
		_, err := df.SelectOptimized([]string{"nonexistent_column"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("WithColumnWrongLength", func(t *testing.T) {
		shortCol := createNewTestColumn(df.NumRows() - 1)
		defer shortCol.Release()

		_, err := df.WithColumnOptimized("test", shortCol)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "does not match DataFrame rows")
	})
}

// Helper functions for test data creation and validation

func createTestDataFrameForOptimization() *DataFrame {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "active", Type: arrow.FixedWidthTypes.Boolean},
		},
		nil,
	)

	size := 10
	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)
	activeBuilder := array.NewBooleanBuilder(pool)

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i) * 2.5)
		nameBuilder.Append("item_" + string(rune('A'+i%5)))
		activeBuilder.Append(i%3 == 0)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	activeArray := activeBuilder.NewArray()
	defer activeArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray, nameArray, activeArray}, int64(size))
	defer record.Release()

	return NewDataFrame(record)
}

func createSelectiveMask(size int64, selectFunc func(int) bool) arrow.Array {
	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	for i := 0; i < int(size); i++ {
		builder.Append(selectFunc(i))
	}

	return builder.NewArray()
}

func createNewTestColumn(size int64) arrow.Array {
	pool := memory.NewGoAllocator()
	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	for i := 0; i < int(size); i++ {
		builder.Append(float64(i) * 10.0)
	}

	return builder.NewArray()
}

// Array comparison helpers

func assertInt64ArraysEqual(t *testing.T, expected, actual *array.Int64) {
	assert.Equal(t, expected.Len(), actual.Len())
	for i := 0; i < expected.Len(); i++ {
		assert.Equal(t, expected.IsNull(i), actual.IsNull(i))
		if !expected.IsNull(i) {
			assert.Equal(t, expected.Value(i), actual.Value(i))
		}
	}
}

func assertFloat64ArraysEqual(t *testing.T, expected, actual *array.Float64) {
	assert.Equal(t, expected.Len(), actual.Len())
	for i := 0; i < expected.Len(); i++ {
		assert.Equal(t, expected.IsNull(i), actual.IsNull(i))
		if !expected.IsNull(i) {
			assert.InDelta(t, expected.Value(i), actual.Value(i), 1e-10)
		}
	}
}

func assertStringArraysEqual(t *testing.T, expected, actual *array.String) {
	assert.Equal(t, expected.Len(), actual.Len())
	for i := 0; i < expected.Len(); i++ {
		assert.Equal(t, expected.IsNull(i), actual.IsNull(i))
		if !expected.IsNull(i) {
			assert.Equal(t, expected.Value(i), actual.Value(i))
		}
	}
}

func assertBooleanArraysEqual(t *testing.T, expected, actual *array.Boolean) {
	assert.Equal(t, expected.Len(), actual.Len())
	for i := 0; i < expected.Len(); i++ {
		assert.Equal(t, expected.IsNull(i), actual.IsNull(i))
		if !expected.IsNull(i) {
			assert.Equal(t, expected.Value(i), actual.Value(i))
		}
	}
}
