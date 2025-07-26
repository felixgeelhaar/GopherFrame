package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataFrameJoins(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create first DataFrame
	schema1 := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder1 := array.NewInt64Builder(pool)
	nameBuilder1 := array.NewStringBuilder(pool)
	idBuilder1.AppendValues([]int64{1, 2, 3}, nil)
	nameBuilder1.AppendValues([]string{"Alice", "Bob", "Charlie"}, nil)

	idArray1 := idBuilder1.NewArray()
	nameArray1 := nameBuilder1.NewArray()
	defer idArray1.Release()
	defer nameArray1.Release()

	record1 := array.NewRecord(schema1, []arrow.Array{idArray1, nameArray1}, 3)
	df1 := NewDataFrame(record1)
	defer df1.Release()

	// Create second DataFrame
	schema2 := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "age", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder2 := array.NewInt64Builder(pool)
	ageBuilder2 := array.NewInt64Builder(pool)
	idBuilder2.AppendValues([]int64{1, 2, 4}, nil) // Note: 4 doesn't exist in df1, 3 doesn't exist in df2
	ageBuilder2.AppendValues([]int64{25, 30, 35}, nil)

	idArray2 := idBuilder2.NewArray()
	ageArray2 := ageBuilder2.NewArray()
	defer idArray2.Release()
	defer ageArray2.Release()

	record2 := array.NewRecord(schema2, []arrow.Array{idArray2, ageArray2}, 3)
	df2 := NewDataFrame(record2)
	defer df2.Release()

	t.Run("InnerJoin", func(t *testing.T) {
		result, err := df1.InnerJoin(df2, "id", "id")
		require.NoError(t, err)
		defer result.Release()

		// Should have 2 rows (ids 1 and 2 match)
		assert.Equal(t, int64(2), result.NumRows())
		assert.Equal(t, int64(3), result.NumCols()) // id, name, age

		// Check schema
		schema := result.Schema()
		assert.Equal(t, "id", schema.Field(0).Name)
		assert.Equal(t, "name", schema.Field(1).Name)
		assert.Equal(t, "age", schema.Field(2).Name)

		// Check data
		record := result.Record()
		idCol := record.Column(0).(*array.Int64)
		nameCol := record.Column(1).(*array.String)
		ageCol := record.Column(2).(*array.Int64)

		assert.Equal(t, int64(1), idCol.Value(0))
		assert.Equal(t, "Alice", nameCol.Value(0))
		assert.Equal(t, int64(25), ageCol.Value(0))

		assert.Equal(t, int64(2), idCol.Value(1))
		assert.Equal(t, "Bob", nameCol.Value(1))
		assert.Equal(t, int64(30), ageCol.Value(1))
	})

	t.Run("LeftJoin", func(t *testing.T) {
		result, err := df1.LeftJoin(df2, "id", "id")
		require.NoError(t, err)
		defer result.Release()

		// Should have 3 rows (all from left df)
		assert.Equal(t, int64(3), result.NumRows())
		assert.Equal(t, int64(3), result.NumCols()) // id, name, age

		// Check data
		record := result.Record()
		idCol := record.Column(0).(*array.Int64)
		nameCol := record.Column(1).(*array.String)
		ageCol := record.Column(2).(*array.Int64)

		assert.Equal(t, int64(1), idCol.Value(0))
		assert.Equal(t, "Alice", nameCol.Value(0))
		assert.Equal(t, int64(25), ageCol.Value(0))

		assert.Equal(t, int64(2), idCol.Value(1))
		assert.Equal(t, "Bob", nameCol.Value(1))
		assert.Equal(t, int64(30), ageCol.Value(1))

		assert.Equal(t, int64(3), idCol.Value(2))
		assert.Equal(t, "Charlie", nameCol.Value(2))
		assert.True(t, ageCol.IsNull(2)) // No matching age for Charlie
	})

	t.Run("JoinWithCustomType", func(t *testing.T) {
		// Test join with explicit join type
		result, err := df1.Join(df2, "id", "id", InnerJoin)
		require.NoError(t, err)
		defer result.Release()

		// Should be same as InnerJoin
		assert.Equal(t, int64(2), result.NumRows())
		assert.Equal(t, int64(3), result.NumCols())
	})

	t.Run("JoinErrors", func(t *testing.T) {
		// Test join with non-existent column
		_, err := df1.InnerJoin(df2, "nonexistent", "id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "column not found")

		// Test with unsupported join type
		_, err = df1.Join(df2, "id", "id", JoinType(999))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported join type")
	})
}

func TestDataFrameIterator(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	valueBuilder.AppendValues([]float64{1.1, 2.2, 3.3}, nil)

	idArray := idBuilder.NewArray()
	valueArray := valueBuilder.NewArray()
	defer idArray.Release()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 3)
	df := NewDataFrame(record)
	defer df.Release()

	t.Run("DataFrameAccess", func(t *testing.T) {
		// Test basic access methods
		assert.Equal(t, int64(3), df.NumRows())
		assert.Equal(t, int64(2), df.NumCols())

		schema := df.Schema()
		assert.NotNil(t, schema)
		assert.Equal(t, 2, len(schema.Fields()))
	})
}

func TestUtilityFunctions(t *testing.T) {
	t.Run("GetColumnIndex", func(t *testing.T) {
		pool := memory.NewGoAllocator()

		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "name", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		idBuilder := array.NewInt64Builder(pool)
		nameBuilder := array.NewStringBuilder(pool)
		idBuilder.Append(1)
		nameBuilder.Append("test")

		idArray := idBuilder.NewArray()
		nameArray := nameBuilder.NewArray()
		defer idArray.Release()
		defer nameArray.Release()

		record := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 1)
		df := NewDataFrame(record)
		defer df.Release()

		// Test existing column
		idx := df.getColumnIndex("id")
		assert.Equal(t, 0, idx)

		idx = df.getColumnIndex("name")
		assert.Equal(t, 1, idx)

		// Test non-existent column
		idx = df.getColumnIndex("nonexistent")
		assert.Equal(t, -1, idx)
	})

	t.Run("ExtractValue", func(t *testing.T) {
		pool := memory.NewGoAllocator()

		// Test with int64 array
		intBuilder := array.NewInt64Builder(pool)
		intBuilder.AppendValues([]int64{1, 2, 3}, nil)
		intArray := intBuilder.NewArray()
		defer intArray.Release()

		val := extractValue(intArray, 1)
		assert.Equal(t, int64(2), val)

		// Test with string array
		strBuilder := array.NewStringBuilder(pool)
		strBuilder.AppendValues([]string{"a", "b", "c"}, nil)
		strArray := strBuilder.NewArray()
		defer strArray.Release()

		val = extractValue(strArray, 0)
		assert.Equal(t, "a", val)

		// Test with float64 array
		floatBuilder := array.NewFloat64Builder(pool)
		floatBuilder.AppendValues([]float64{1.1, 2.2, 3.3}, nil)
		floatArray := floatBuilder.NewArray()
		defer floatArray.Release()

		val = extractValue(floatArray, 2)
		assert.Equal(t, float64(3.3), val)

		// Test with boolean array
		boolBuilder := array.NewBooleanBuilder(pool)
		boolBuilder.AppendValues([]bool{true, false, true}, nil)
		boolArray := boolBuilder.NewArray()
		defer boolArray.Release()

		val = extractValue(boolArray, 1)
		assert.Equal(t, false, val)

		// Test with date type (should extract as string representation)
		dateBuilder := array.NewDate32Builder(pool)
		dateBuilder.Append(arrow.Date32(19000))
		dateArray := dateBuilder.NewArray()
		defer dateArray.Release()

		val = extractValue(dateArray, 0)
		// Date extraction returns the date value, not nil
		assert.NotNil(t, val)
	})
}

func TestParallelCombineResults(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test schema
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Create first chunk
	idBuilder1 := array.NewInt64Builder(pool)
	valueBuilder1 := array.NewFloat64Builder(pool)
	idBuilder1.AppendValues([]int64{1, 2}, nil)
	valueBuilder1.AppendValues([]float64{1.1, 2.2}, nil)

	idArray1 := idBuilder1.NewArray()
	valueArray1 := valueBuilder1.NewArray()
	defer idArray1.Release()
	defer valueArray1.Release()

	record1 := array.NewRecord(schema, []arrow.Array{idArray1, valueArray1}, 2)
	df1 := NewDataFrame(record1)
	defer df1.Release()

	// Create second chunk
	idBuilder2 := array.NewInt64Builder(pool)
	valueBuilder2 := array.NewFloat64Builder(pool)
	idBuilder2.AppendValues([]int64{3, 4}, nil)
	valueBuilder2.AppendValues([]float64{3.3, 4.4}, nil)

	idArray2 := idBuilder2.NewArray()
	valueArray2 := valueBuilder2.NewArray()
	defer idArray2.Release()
	defer valueArray2.Release()

	record2 := array.NewRecord(schema, []arrow.Array{idArray2, valueArray2}, 2)
	df2 := NewDataFrame(record2)
	defer df2.Release()

	t.Run("CombineResults", func(t *testing.T) {
		processor := NewParallelProcessor(DefaultParallelOptions())

		// Create arrays from the dataframes
		arrays := []arrow.Array{df1.Record().Column(0), df2.Record().Column(0)}

		result, err := processor.CombineResults(arrays)
		require.NoError(t, err)
		defer result.Release()

		// Should have 4 values total
		assert.Equal(t, 4, result.Len())

		// Check combined data
		idCol := result.(*array.Int64)
		assert.Equal(t, int64(1), idCol.Value(0))
		assert.Equal(t, int64(3), idCol.Value(2))
	})
}
