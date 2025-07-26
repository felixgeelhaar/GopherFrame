package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataFrameComprehensive(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("DataFrameWithDifferentTypes", func(t *testing.T) {
		// Create DataFrame with many different arrow types
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "int8_col", Type: arrow.PrimitiveTypes.Int8},
				{Name: "int16_col", Type: arrow.PrimitiveTypes.Int16},
				{Name: "int32_col", Type: arrow.PrimitiveTypes.Int32},
				{Name: "int64_col", Type: arrow.PrimitiveTypes.Int64},
				{Name: "uint8_col", Type: arrow.PrimitiveTypes.Uint8},
				{Name: "uint16_col", Type: arrow.PrimitiveTypes.Uint16},
				{Name: "uint32_col", Type: arrow.PrimitiveTypes.Uint32},
				{Name: "uint64_col", Type: arrow.PrimitiveTypes.Uint64},
				{Name: "float32_col", Type: arrow.PrimitiveTypes.Float32},
				{Name: "float64_col", Type: arrow.PrimitiveTypes.Float64},
				{Name: "bool_col", Type: arrow.FixedWidthTypes.Boolean},
				{Name: "string_col", Type: arrow.BinaryTypes.String},
				{Name: "date32_col", Type: arrow.FixedWidthTypes.Date32},
				{Name: "date64_col", Type: arrow.FixedWidthTypes.Date64},
				{Name: "timestamp_col", Type: &arrow.TimestampType{Unit: arrow.Millisecond}},
			},
			nil,
		)

		// Create builders for all types
		int8Builder := array.NewInt8Builder(pool)
		int16Builder := array.NewInt16Builder(pool)
		int32Builder := array.NewInt32Builder(pool)
		int64Builder := array.NewInt64Builder(pool)
		uint8Builder := array.NewUint8Builder(pool)
		uint16Builder := array.NewUint16Builder(pool)
		uint32Builder := array.NewUint32Builder(pool)
		uint64Builder := array.NewUint64Builder(pool)
		float32Builder := array.NewFloat32Builder(pool)
		float64Builder := array.NewFloat64Builder(pool)
		boolBuilder := array.NewBooleanBuilder(pool)
		stringBuilder := array.NewStringBuilder(pool)
		date32Builder := array.NewDate32Builder(pool)
		date64Builder := array.NewDate64Builder(pool)
		timestampBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Millisecond})

		// Append test data
		numRows := 5
		for i := 0; i < numRows; i++ {
			int8Builder.Append(int8(i))
			int16Builder.Append(int16(i * 10))
			int32Builder.Append(int32(i * 100))
			int64Builder.Append(int64(i * 1000))
			uint8Builder.Append(uint8(i + 1))
			uint16Builder.Append(uint16(i + 10))
			uint32Builder.Append(uint32(i + 100))
			uint64Builder.Append(uint64(i + 1000))
			float32Builder.Append(float32(i) + 0.5)
			float64Builder.Append(float64(i) + 0.25)
			boolBuilder.Append(i%2 == 0)
			stringBuilder.AppendString("row" + string(rune('0'+i)))
			date32Builder.Append(arrow.Date32(19000 + i))
			date64Builder.Append(arrow.Date64(1640995200000 + int64(i)*86400000))
			timestampBuilder.Append(arrow.Timestamp(1640995200000 + int64(i)*3600000))
		}

		// Build arrays
		arrays := []arrow.Array{
			int8Builder.NewArray(),
			int16Builder.NewArray(),
			int32Builder.NewArray(),
			int64Builder.NewArray(),
			uint8Builder.NewArray(),
			uint16Builder.NewArray(),
			uint32Builder.NewArray(),
			uint64Builder.NewArray(),
			float32Builder.NewArray(),
			float64Builder.NewArray(),
			boolBuilder.NewArray(),
			stringBuilder.NewArray(),
			date32Builder.NewArray(),
			date64Builder.NewArray(),
			timestampBuilder.NewArray(),
		}

		// Release arrays when done
		defer func() {
			for _, arr := range arrays {
				arr.Release()
			}
		}()

		// Create DataFrame
		record := array.NewRecord(schema, arrays, int64(numRows))
		df := NewDataFrame(record)
		defer df.Release()

		// Test basic properties
		assert.Equal(t, int64(numRows), df.NumRows())
		assert.Equal(t, int64(15), df.NumCols())

		// Test accessing each column
		for i, expectedName := range []string{
			"int8_col", "int16_col", "int32_col", "int64_col",
			"uint8_col", "uint16_col", "uint32_col", "uint64_col",
			"float32_col", "float64_col", "bool_col", "string_col",
			"date32_col", "date64_col", "timestamp_col",
		} {
			series, err := df.Column(expectedName)
			require.NoError(t, err)
			assert.Equal(t, expectedName, series.Field().Name)
			assert.Equal(t, numRows, series.Len())

			// Also test via Columns() method
			columns := df.Columns()
			assert.Equal(t, expectedName, columns[i].Field().Name)
		}

		// Test Schema method
		dfSchema := df.Schema()
		assert.Equal(t, 15, len(dfSchema.Fields()))

		// Test String method
		str := df.String()
		assert.NotEmpty(t, str)
		assert.Contains(t, str, "DataFrame")

		// Test Record method
		dfRecord := df.Record()
		assert.NotNil(t, dfRecord)
		assert.Equal(t, int64(numRows), dfRecord.NumRows())
		assert.Equal(t, int64(15), dfRecord.NumCols())
	})

	t.Run("DataFrameErrorCases", func(t *testing.T) {
		// Create a simple DataFrame
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "test_col", Type: arrow.PrimitiveTypes.Int64},
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

		// Test error cases
		_, err := df.Column("nonexistent_column")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "column not found")

		_, err = df.Column("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "column not found")
	})

	t.Run("DataFrameMemoryHandling", func(t *testing.T) {
		// Test multiple DataFrame creations and releases
		for i := 0; i < 10; i++ {
			schema := arrow.NewSchema(
				[]arrow.Field{
					{Name: "iter_col", Type: arrow.PrimitiveTypes.Int32},
				},
				nil,
			)

			builder := array.NewInt32Builder(pool)
			builder.AppendValues([]int32{int32(i), int32(i + 1), int32(i + 2)}, nil)
			arr := builder.NewArray()

			record := array.NewRecord(schema, []arrow.Array{arr}, 3)
			df := NewDataFrame(record)

			// Test that DataFrame properties work correctly
			assert.Equal(t, int64(3), df.NumRows())
			assert.Equal(t, int64(1), df.NumCols())

			// Release resources
			arr.Release()
			df.Release()
		}
	})
}

func TestSeriesComprehensive(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("SeriesWithDifferentTypes", func(t *testing.T) {
		// Test Series with different Arrow types
		testCases := []struct {
			name      string
			dataType  arrow.DataType
			buildFunc func() arrow.Array
		}{
			{
				name:     "Int8Series",
				dataType: arrow.PrimitiveTypes.Int8,
				buildFunc: func() arrow.Array {
					builder := array.NewInt8Builder(pool)
					builder.AppendValues([]int8{1, 2, 3}, nil)
					return builder.NewArray()
				},
			},
			{
				name:     "Float32Series",
				dataType: arrow.PrimitiveTypes.Float32,
				buildFunc: func() arrow.Array {
					builder := array.NewFloat32Builder(pool)
					builder.AppendValues([]float32{1.1, 2.2, 3.3}, nil)
					return builder.NewArray()
				},
			},
			{
				name:     "BooleanSeries",
				dataType: arrow.FixedWidthTypes.Boolean,
				buildFunc: func() arrow.Array {
					builder := array.NewBooleanBuilder(pool)
					builder.AppendValues([]bool{true, false, true}, nil)
					return builder.NewArray()
				},
			},
			{
				name:     "StringSeries",
				dataType: arrow.BinaryTypes.String,
				buildFunc: func() arrow.Array {
					builder := array.NewStringBuilder(pool)
					builder.AppendValues([]string{"a", "b", "c"}, nil)
					return builder.NewArray()
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				arr := tc.buildFunc()
				defer arr.Release()

				field := arrow.Field{Name: tc.name, Type: tc.dataType}
				series := &Series{array: arr, field: field}

				// Test basic Series methods
				assert.Equal(t, 3, series.Len())
				assert.Equal(t, tc.name, series.Field().Name)
				assert.Equal(t, tc.dataType, series.Field().Type)

				// Test Array method
				retrievedArray := series.Array()
				assert.NotNil(t, retrievedArray)
				assert.Equal(t, 3, retrievedArray.Len())
				assert.Equal(t, tc.dataType, retrievedArray.DataType())
			})
		}
	})

	t.Run("SeriesWithNullValues", func(t *testing.T) {
		// Test Series with null values
		builder := array.NewInt64Builder(pool)
		builder.AppendValues([]int64{1, 2, 3}, []bool{true, false, true}) // Second value is null
		arr := builder.NewArray()
		defer arr.Release()

		field := arrow.Field{Name: "nullable_series", Type: arrow.PrimitiveTypes.Int64, Nullable: true}
		series := &Series{array: arr, field: field}

		assert.Equal(t, 3, series.Len())
		assert.Equal(t, 1, series.Array().NullN()) // One null value
		assert.True(t, series.Field().Nullable)
	})
}
