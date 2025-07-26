package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataFrameMethods(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("DataFrameRecord", func(t *testing.T) {
		// Test Record() method
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			},
			nil,
		)

		builder := array.NewInt64Builder(pool)
		builder.AppendValues([]int64{1, 2, 3}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 3)
		df := NewDataFrame(record)
		defer df.Release()

		// Test Record method
		retrievedRecord := df.Record()
		assert.NotNil(t, retrievedRecord)
		assert.Equal(t, int64(3), retrievedRecord.NumRows())
		assert.Equal(t, int64(1), retrievedRecord.NumCols())
	})

	t.Run("DataFrameRelease", func(t *testing.T) {
		// Test proper release behavior
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "data", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		builder := array.NewFloat64Builder(pool)
		builder.AppendValues([]float64{1.1, 2.2}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 2)
		df := NewDataFrame(record)

		// Test that DataFrame can be released without issues
		df.Release()

		// Note: We can't test much after release due to potential memory issues
		// But the important part is that Release() doesn't panic
	})

	t.Run("SeriesOperations", func(t *testing.T) {
		// Test Series-related operations
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "values", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		builder := array.NewStringBuilder(pool)
		builder.AppendValues([]string{"hello", "world"}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 2)
		df := NewDataFrame(record)
		defer df.Release()

		// Test Column method
		series, err := df.Column("values")
		require.NoError(t, err)
		assert.NotNil(t, series)
		assert.Equal(t, 2, series.Len())

		// Test Field method on Series
		field := series.Field()
		assert.Equal(t, "values", field.Name)
		assert.Equal(t, arrow.BinaryTypes.String, field.Type)
	})

	t.Run("DataFrameBasicProperties", func(t *testing.T) {
		// Test basic DataFrame properties with different data types
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "int_col", Type: arrow.PrimitiveTypes.Int32},
				{Name: "str_col", Type: arrow.BinaryTypes.String},
				{Name: "bool_col", Type: arrow.FixedWidthTypes.Boolean},
			},
			nil,
		)

		intBuilder := array.NewInt32Builder(pool)
		strBuilder := array.NewStringBuilder(pool)
		boolBuilder := array.NewBooleanBuilder(pool)

		intBuilder.AppendValues([]int32{10, 20, 30}, nil)
		strBuilder.AppendValues([]string{"a", "b", "c"}, nil)
		boolBuilder.AppendValues([]bool{true, false, true}, nil)

		intArr := intBuilder.NewArray()
		strArr := strBuilder.NewArray()
		boolArr := boolBuilder.NewArray()
		defer intArr.Release()
		defer strArr.Release()
		defer boolArr.Release()

		record := array.NewRecord(schema, []arrow.Array{intArr, strArr, boolArr}, 3)
		df := NewDataFrame(record)
		defer df.Release()

		// Test basic properties
		assert.Equal(t, int64(3), df.NumRows())
		assert.Equal(t, int64(3), df.NumCols())

		// Test Schema
		dfSchema := df.Schema()
		assert.Equal(t, 3, len(dfSchema.Fields()))

		// Test Columns
		columns := df.Columns()
		assert.Len(t, columns, 3)

		// Test individual column access
		for i, expectedName := range []string{"int_col", "str_col", "bool_col"} {
			assert.Equal(t, expectedName, columns[i].Field().Name)
		}
	})
}

func TestSeriesMethods(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("SeriesBasicOperations", func(t *testing.T) {
		// Create a Series directly for testing
		builder := array.NewInt64Builder(pool)
		builder.AppendValues([]int64{100, 200, 300}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		field := arrow.Field{Name: "test_series", Type: arrow.PrimitiveTypes.Int64}
		series := &Series{array: arr, field: field}

		// Test basic Series methods
		assert.Equal(t, 3, series.Len())
		assert.Equal(t, "test_series", series.Field().Name)
		assert.Equal(t, arrow.PrimitiveTypes.Int64, series.Field().Type)

		// Test Array method
		retrievedArray := series.Array()
		assert.NotNil(t, retrievedArray)
		assert.Equal(t, 3, retrievedArray.Len())
	})
}
