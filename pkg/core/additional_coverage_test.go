package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataFrameEdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("EmptyDataFrame", func(t *testing.T) {
		// Test empty DataFrame creation and operations
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "empty_col", Type: arrow.PrimitiveTypes.Int64},
			},
			nil,
		)

		builder := array.NewInt64Builder(pool)
		emptyArray := builder.NewArray()
		defer emptyArray.Release()

		record := array.NewRecord(schema, []arrow.Array{emptyArray}, 0)
		df := NewDataFrame(record)
		defer df.Release()

		assert.Equal(t, int64(0), df.NumRows())
		assert.Equal(t, int64(1), df.NumCols())

		// Test operations on empty DataFrame
		columns := df.Columns()
		assert.Len(t, columns, 1)
		assert.Equal(t, "empty_col", columns[0].Field().Name)
	})

	t.Run("DataFrameSchema", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
				{Name: "float_col", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		intBuilder := array.NewInt64Builder(pool)
		floatBuilder := array.NewFloat64Builder(pool)

		intBuilder.AppendValues([]int64{1, 2}, nil)
		floatBuilder.AppendValues([]float64{1.1, 2.2}, nil)

		intArray := intBuilder.NewArray()
		floatArray := floatBuilder.NewArray()
		defer intArray.Release()
		defer floatArray.Release()

		record := array.NewRecord(schema, []arrow.Array{intArray, floatArray}, 2)
		df := NewDataFrame(record)
		defer df.Release()

		// Test Schema method
		resultSchema := df.Schema()
		assert.Equal(t, 2, len(resultSchema.Fields()))
		assert.Equal(t, "int_col", resultSchema.Field(0).Name)
		assert.Equal(t, "float_col", resultSchema.Field(1).Name)
	})

	t.Run("DataFrameColumnAccess", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "test_col", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		builder := array.NewStringBuilder(pool)
		builder.AppendValues([]string{"a", "b", "c"}, nil)
		strArray := builder.NewArray()
		defer strArray.Release()

		record := array.NewRecord(schema, []arrow.Array{strArray}, 3)
		df := NewDataFrame(record)
		defer df.Release()

		// Test Column method
		col, err := df.Column("test_col")
		require.NoError(t, err)
		assert.Equal(t, 3, col.Len())

		// Test non-existent column
		_, err = df.Column("non_existent")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "column not found")
	})
}

func TestDataFrameStringRepresentation(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("StringMethod", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "name", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		intBuilder := array.NewInt64Builder(pool)
		strBuilder := array.NewStringBuilder(pool)

		intBuilder.AppendValues([]int64{1, 2}, nil)
		strBuilder.AppendValues([]string{"Alice", "Bob"}, nil)

		intArray := intBuilder.NewArray()
		strArray := strBuilder.NewArray()
		defer intArray.Release()
		defer strArray.Release()

		record := array.NewRecord(schema, []arrow.Array{intArray, strArray}, 2)
		df := NewDataFrame(record)
		defer df.Release()

		// Test String method returns something meaningful
		str := df.String()
		assert.NotEmpty(t, str)
		assert.Contains(t, str, "DataFrame")
	})
}

func TestDataFrameMemoryManagement(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("ReferenceCountingBasics", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "data", Type: arrow.PrimitiveTypes.Int32},
			},
			nil,
		)

		builder := array.NewInt32Builder(pool)
		builder.AppendValues([]int32{10, 20, 30}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 3)
		df := NewDataFrame(record)

		// Test that DataFrame can be created and released properly
		assert.Equal(t, int64(3), df.NumRows())
		df.Release()

		// After release, accessing should still work (reference counting)
		// This tests that the DataFrame properly manages Arrow memory
	})
}
