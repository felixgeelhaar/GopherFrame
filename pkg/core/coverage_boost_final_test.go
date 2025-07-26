package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataFrameExtensiveCoverage(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("DataFrameWithManyColumns", func(t *testing.T) {
		// Create DataFrame with 20+ columns of different types
		fields := []arrow.Field{
			{Name: "col1", Type: arrow.PrimitiveTypes.Int8},
			{Name: "col2", Type: arrow.PrimitiveTypes.Int16},
			{Name: "col3", Type: arrow.PrimitiveTypes.Int32},
			{Name: "col4", Type: arrow.PrimitiveTypes.Int64},
			{Name: "col5", Type: arrow.PrimitiveTypes.Uint8},
			{Name: "col6", Type: arrow.PrimitiveTypes.Uint16},
			{Name: "col7", Type: arrow.PrimitiveTypes.Uint32},
			{Name: "col8", Type: arrow.PrimitiveTypes.Uint64},
			{Name: "col9", Type: arrow.PrimitiveTypes.Float32},
			{Name: "col10", Type: arrow.PrimitiveTypes.Float64},
			{Name: "col11", Type: arrow.FixedWidthTypes.Boolean},
			{Name: "col12", Type: arrow.BinaryTypes.String},
			{Name: "col13", Type: arrow.FixedWidthTypes.Date32},
			{Name: "col14", Type: arrow.FixedWidthTypes.Date64},
			{Name: "col15", Type: &arrow.TimestampType{Unit: arrow.Second}},
			{Name: "col16", Type: &arrow.TimestampType{Unit: arrow.Millisecond}},
			{Name: "col17", Type: &arrow.TimestampType{Unit: arrow.Microsecond}},
			{Name: "col18", Type: &arrow.TimestampType{Unit: arrow.Nanosecond}},
			{Name: "col19", Type: &arrow.Time32Type{Unit: arrow.Second}},
			{Name: "col20", Type: &arrow.Time64Type{Unit: arrow.Microsecond}},
		}

		schema := arrow.NewSchema(fields, nil)

		// Create arrays for all columns
		arrays := make([]arrow.Array, len(fields))

		// Col1: Int8
		int8Builder := array.NewInt8Builder(pool)
		int8Builder.AppendValues([]int8{1, 2, 3, 4, 5}, nil)
		arrays[0] = int8Builder.NewArray()

		// Col2: Int16
		int16Builder := array.NewInt16Builder(pool)
		int16Builder.AppendValues([]int16{10, 20, 30, 40, 50}, nil)
		arrays[1] = int16Builder.NewArray()

		// Col3: Int32
		int32Builder := array.NewInt32Builder(pool)
		int32Builder.AppendValues([]int32{100, 200, 300, 400, 500}, nil)
		arrays[2] = int32Builder.NewArray()

		// Col4: Int64
		int64Builder := array.NewInt64Builder(pool)
		int64Builder.AppendValues([]int64{1000, 2000, 3000, 4000, 5000}, nil)
		arrays[3] = int64Builder.NewArray()

		// Col5: Uint8
		uint8Builder := array.NewUint8Builder(pool)
		uint8Builder.AppendValues([]uint8{1, 2, 3, 4, 5}, nil)
		arrays[4] = uint8Builder.NewArray()

		// Col6: Uint16
		uint16Builder := array.NewUint16Builder(pool)
		uint16Builder.AppendValues([]uint16{10, 20, 30, 40, 50}, nil)
		arrays[5] = uint16Builder.NewArray()

		// Col7: Uint32
		uint32Builder := array.NewUint32Builder(pool)
		uint32Builder.AppendValues([]uint32{100, 200, 300, 400, 500}, nil)
		arrays[6] = uint32Builder.NewArray()

		// Col8: Uint64
		uint64Builder := array.NewUint64Builder(pool)
		uint64Builder.AppendValues([]uint64{1000, 2000, 3000, 4000, 5000}, nil)
		arrays[7] = uint64Builder.NewArray()

		// Col9: Float32
		float32Builder := array.NewFloat32Builder(pool)
		float32Builder.AppendValues([]float32{1.1, 2.2, 3.3, 4.4, 5.5}, nil)
		arrays[8] = float32Builder.NewArray()

		// Col10: Float64
		float64Builder := array.NewFloat64Builder(pool)
		float64Builder.AppendValues([]float64{10.1, 20.2, 30.3, 40.4, 50.5}, nil)
		arrays[9] = float64Builder.NewArray()

		// Col11: Boolean
		boolBuilder := array.NewBooleanBuilder(pool)
		boolBuilder.AppendValues([]bool{true, false, true, false, true}, nil)
		arrays[10] = boolBuilder.NewArray()

		// Col12: String
		stringBuilder := array.NewStringBuilder(pool)
		stringBuilder.AppendValues([]string{"a", "b", "c", "d", "e"}, nil)
		arrays[11] = stringBuilder.NewArray()

		// Col13: Date32
		date32Builder := array.NewDate32Builder(pool)
		date32Builder.AppendValues([]arrow.Date32{19523, 19524, 19525, 19526, 19527}, nil)
		arrays[12] = date32Builder.NewArray()

		// Col14: Date64
		date64Builder := array.NewDate64Builder(pool)
		date64Builder.AppendValues([]arrow.Date64{1686787200000, 1686873600000, 1686960000000, 1687046400000, 1687132800000}, nil)
		arrays[13] = date64Builder.NewArray()

		// Col15: Timestamp(Second)
		timestampSecBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Second})
		timestampSecBuilder.AppendValues([]arrow.Timestamp{1686787200, 1686873600, 1686960000, 1687046400, 1687132800}, nil)
		arrays[14] = timestampSecBuilder.NewArray()

		// Col16: Timestamp(Millisecond)
		timestampMsBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Millisecond})
		timestampMsBuilder.AppendValues([]arrow.Timestamp{1686787200000, 1686873600000, 1686960000000, 1687046400000, 1687132800000}, nil)
		arrays[15] = timestampMsBuilder.NewArray()

		// Col17: Timestamp(Microsecond)
		timestampUsBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Microsecond})
		timestampUsBuilder.AppendValues([]arrow.Timestamp{1686787200000000, 1686873600000000, 1686960000000000, 1687046400000000, 1687132800000000}, nil)
		arrays[16] = timestampUsBuilder.NewArray()

		// Col18: Timestamp(Nanosecond)
		timestampNsBuilder := array.NewTimestampBuilder(pool, &arrow.TimestampType{Unit: arrow.Nanosecond})
		timestampNsBuilder.AppendValues([]arrow.Timestamp{1686787200000000000, 1686873600000000000, 1686960000000000000, 1687046400000000000, 1687132800000000000}, nil)
		arrays[17] = timestampNsBuilder.NewArray()

		// Col19: Time32(Second)
		time32Builder := array.NewTime32Builder(pool, &arrow.Time32Type{Unit: arrow.Second})
		time32Builder.AppendValues([]arrow.Time32{3600, 7200, 10800, 14400, 18000}, nil)
		arrays[18] = time32Builder.NewArray()

		// Col20: Time64(Microsecond)
		time64Builder := array.NewTime64Builder(pool, &arrow.Time64Type{Unit: arrow.Microsecond})
		time64Builder.AppendValues([]arrow.Time64{3600000000, 7200000000, 10800000000, 14400000000, 18000000000}, nil)
		arrays[19] = time64Builder.NewArray()

		// Release all arrays when done
		defer func() {
			for _, arr := range arrays {
				arr.Release()
			}
		}()

		// Create DataFrame
		record := array.NewRecord(schema, arrays, 5)
		df := NewDataFrame(record)
		defer df.Release()

		// Test basic properties
		assert.Equal(t, int64(5), df.NumRows())
		assert.Equal(t, int64(20), df.NumCols())

		// Test Schema
		dfSchema := df.Schema()
		assert.Equal(t, 20, len(dfSchema.Fields()))

		// Test Columns method
		columns := df.Columns()
		assert.Len(t, columns, 20)

		// Test accessing all columns by name
		for i, field := range fields {
			series, err := df.Column(field.Name)
			require.NoError(t, err)
			assert.Equal(t, field.Name, series.Field().Name)
			assert.Equal(t, 5, series.Len())

			// Verify column order
			assert.Equal(t, field.Name, columns[i].Field().Name)
		}

		// Test Record method
		dfRecord := df.Record()
		assert.NotNil(t, dfRecord)
		assert.Equal(t, int64(5), dfRecord.NumRows())
		assert.Equal(t, int64(20), dfRecord.NumCols())

		// Test String method with large DataFrame
		str := df.String()
		assert.NotEmpty(t, str)
		assert.Contains(t, str, "DataFrame")
	})

	t.Run("DataFrameWithNullValues", func(t *testing.T) {
		// Create DataFrame with null values in different types
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "nullable_int", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
				{Name: "nullable_float", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
				{Name: "nullable_string", Type: arrow.BinaryTypes.String, Nullable: true},
				{Name: "nullable_bool", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
			},
			nil,
		)

		// Create arrays with null values
		intBuilder := array.NewInt64Builder(pool)
		intBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, []bool{true, false, true, false, true})

		floatBuilder := array.NewFloat64Builder(pool)
		floatBuilder.AppendValues([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, []bool{false, true, false, true, false})

		stringBuilder := array.NewStringBuilder(pool)
		stringBuilder.AppendValues([]string{"a", "b", "c", "d", "e"}, []bool{true, true, false, false, true})

		boolBuilder := array.NewBooleanBuilder(pool)
		boolBuilder.AppendValues([]bool{true, false, true, false, true}, []bool{false, false, true, true, false})

		intArray := intBuilder.NewArray()
		floatArray := floatBuilder.NewArray()
		stringArray := stringBuilder.NewArray()
		boolArray := boolBuilder.NewArray()

		defer intArray.Release()
		defer floatArray.Release()
		defer stringArray.Release()
		defer boolArray.Release()

		record := array.NewRecord(schema, []arrow.Array{intArray, floatArray, stringArray, boolArray}, 5)
		df := NewDataFrame(record)
		defer df.Release()

		// Test DataFrame with null values
		assert.Equal(t, int64(5), df.NumRows())
		assert.Equal(t, int64(4), df.NumCols())

		// Test accessing columns with nulls
		intSeries, err := df.Column("nullable_int")
		require.NoError(t, err)
		assert.Equal(t, 2, intSeries.Array().NullN()) // 2 null values

		floatSeries, err := df.Column("nullable_float")
		require.NoError(t, err)
		assert.Equal(t, 3, floatSeries.Array().NullN()) // 3 null values

		stringSeries, err := df.Column("nullable_string")
		require.NoError(t, err)
		assert.Equal(t, 2, stringSeries.Array().NullN()) // 2 null values

		boolSeries, err := df.Column("nullable_bool")
		require.NoError(t, err)
		assert.Equal(t, 3, boolSeries.Array().NullN()) // 3 null values
	})

	t.Run("DataFrameLargeRowCount", func(t *testing.T) {
		// Test DataFrame with larger number of rows
		numRows := 1000
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		intBuilder := array.NewInt64Builder(pool)
		floatBuilder := array.NewFloat64Builder(pool)

		// Generate large dataset
		for i := 0; i < numRows; i++ {
			intBuilder.Append(int64(i))
			floatBuilder.Append(float64(i) * 1.5)
		}

		intArray := intBuilder.NewArray()
		floatArray := floatBuilder.NewArray()
		defer intArray.Release()
		defer floatArray.Release()

		record := array.NewRecord(schema, []arrow.Array{intArray, floatArray}, int64(numRows))
		df := NewDataFrame(record)
		defer df.Release()

		// Test properties with large DataFrame
		assert.Equal(t, int64(numRows), df.NumRows())
		assert.Equal(t, int64(2), df.NumCols())

		// Test column access
		idSeries, err := df.Column("id")
		require.NoError(t, err)
		assert.Equal(t, numRows, idSeries.Len())

		valueSeries, err := df.Column("value")
		require.NoError(t, err)
		assert.Equal(t, numRows, valueSeries.Len())

		// Test that string representation handles large DataFrame
		str := df.String()
		assert.NotEmpty(t, str)
	})
}

func TestSeriesExtensiveCoverage(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("SeriesWithAllArrowTypes", func(t *testing.T) {
		// Test Series creation with various Arrow types
		testCases := []struct {
			name     string
			dataType arrow.DataType
			values   interface{}
		}{
			{"Int8", arrow.PrimitiveTypes.Int8, []int8{1, 2, 3}},
			{"Int16", arrow.PrimitiveTypes.Int16, []int16{10, 20, 30}},
			{"Int32", arrow.PrimitiveTypes.Int32, []int32{100, 200, 300}},
			{"Int64", arrow.PrimitiveTypes.Int64, []int64{1000, 2000, 3000}},
			{"Uint8", arrow.PrimitiveTypes.Uint8, []uint8{1, 2, 3}},
			{"Uint16", arrow.PrimitiveTypes.Uint16, []uint16{10, 20, 30}},
			{"Uint32", arrow.PrimitiveTypes.Uint32, []uint32{100, 200, 300}},
			{"Uint64", arrow.PrimitiveTypes.Uint64, []uint64{1000, 2000, 3000}},
			{"Float32", arrow.PrimitiveTypes.Float32, []float32{1.1, 2.2, 3.3}},
			{"Float64", arrow.PrimitiveTypes.Float64, []float64{10.1, 20.2, 30.3}},
			{"Boolean", arrow.FixedWidthTypes.Boolean, []bool{true, false, true}},
			{"String", arrow.BinaryTypes.String, []string{"a", "b", "c"}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var arr arrow.Array

				// Create appropriate builder and array
				switch tc.dataType {
				case arrow.PrimitiveTypes.Int8:
					builder := array.NewInt8Builder(pool)
					builder.AppendValues(tc.values.([]int8), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Int16:
					builder := array.NewInt16Builder(pool)
					builder.AppendValues(tc.values.([]int16), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Int32:
					builder := array.NewInt32Builder(pool)
					builder.AppendValues(tc.values.([]int32), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Int64:
					builder := array.NewInt64Builder(pool)
					builder.AppendValues(tc.values.([]int64), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Uint8:
					builder := array.NewUint8Builder(pool)
					builder.AppendValues(tc.values.([]uint8), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Uint16:
					builder := array.NewUint16Builder(pool)
					builder.AppendValues(tc.values.([]uint16), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Uint32:
					builder := array.NewUint32Builder(pool)
					builder.AppendValues(tc.values.([]uint32), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Uint64:
					builder := array.NewUint64Builder(pool)
					builder.AppendValues(tc.values.([]uint64), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Float32:
					builder := array.NewFloat32Builder(pool)
					builder.AppendValues(tc.values.([]float32), nil)
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Float64:
					builder := array.NewFloat64Builder(pool)
					builder.AppendValues(tc.values.([]float64), nil)
					arr = builder.NewArray()
				case arrow.FixedWidthTypes.Boolean:
					builder := array.NewBooleanBuilder(pool)
					builder.AppendValues(tc.values.([]bool), nil)
					arr = builder.NewArray()
				case arrow.BinaryTypes.String:
					builder := array.NewStringBuilder(pool)
					builder.AppendValues(tc.values.([]string), nil)
					arr = builder.NewArray()
				}

				defer arr.Release()

				field := arrow.Field{Name: tc.name + "_series", Type: tc.dataType}
				series := &Series{array: arr, field: field}

				// Test all Series methods
				assert.Equal(t, 3, series.Len())
				assert.Equal(t, tc.name+"_series", series.Field().Name)
				assert.Equal(t, tc.dataType, series.Field().Type)

				// Test Array method
				retrievedArray := series.Array()
				assert.NotNil(t, retrievedArray)
				assert.Equal(t, 3, retrievedArray.Len())
				assert.Equal(t, tc.dataType, retrievedArray.DataType())
			})
		}
	})
}
