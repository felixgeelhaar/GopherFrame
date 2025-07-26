package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinalVictory(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("DataFrameErrorScenarios", func(t *testing.T) {
		// Test various error conditions and edge cases to boost coverage
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "test_data", Type: arrow.PrimitiveTypes.Int64},
				{Name: "empty_column", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		intBuilder := array.NewInt64Builder(pool)
		strBuilder := array.NewStringBuilder(pool)

		// Add some test data
		intBuilder.AppendValues([]int64{1, 2, 3}, nil)
		strBuilder.AppendValues([]string{"a", "b", "c"}, nil)

		intArray := intBuilder.NewArray()
		strArray := strBuilder.NewArray()
		defer intArray.Release()
		defer strArray.Release()

		record := array.NewRecord(schema, []arrow.Array{intArray, strArray}, 3)
		df := NewDataFrame(record)
		defer df.Release()

		// Test Column access with various invalid names
		invalidNames := []string{
			"",
			"nonexistent",
			"NONEXISTENT",
			"test_data_wrong",
			"empty_column_wrong",
			"123invalid",
			"invalid!@#",
			"test data", // with space
		}

		for _, name := range invalidNames {
			_, err := df.Column(name)
			assert.Error(t, err, "Expected error for column name: %s", name)
		}

		// Test valid column access multiple times to ensure consistency
		for i := 0; i < 10; i++ {
			series, err := df.Column("test_data")
			require.NoError(t, err)
			assert.Equal(t, "test_data", series.Field().Name)
			assert.Equal(t, 3, series.Len())

			series2, err := df.Column("empty_column")
			require.NoError(t, err)
			assert.Equal(t, "empty_column", series2.Field().Name)
			assert.Equal(t, 3, series2.Len())
		}
	})

	t.Run("DataFrameMemoryScenarios", func(t *testing.T) {
		// Test DataFrame creation and release patterns
		for iteration := 0; iteration < 20; iteration++ {
			schema := arrow.NewSchema(
				[]arrow.Field{
					{Name: "iteration_data", Type: arrow.PrimitiveTypes.Int32},
					{Name: "float_data", Type: arrow.PrimitiveTypes.Float64},
					{Name: "string_data", Type: arrow.BinaryTypes.String},
				},
				nil,
			)

			intBuilder := array.NewInt32Builder(pool)
			floatBuilder := array.NewFloat64Builder(pool)
			strBuilder := array.NewStringBuilder(pool)

			// Vary the data size
			dataSize := iteration + 1
			for i := 0; i < dataSize; i++ {
				intBuilder.Append(int32(i + iteration))
				floatBuilder.Append(float64(i+iteration) * 1.5)
				strBuilder.AppendString(string(rune('A' + (i % 26))))
			}

			intArray := intBuilder.NewArray()
			floatArray := floatBuilder.NewArray()
			strArray := strBuilder.NewArray()

			record := array.NewRecord(schema, []arrow.Array{intArray, floatArray, strArray}, int64(dataSize))
			df := NewDataFrame(record)

			// Test all DataFrame methods
			assert.Equal(t, int64(dataSize), df.NumRows())
			assert.Equal(t, int64(3), df.NumCols())

			// Test Schema
			dfSchema := df.Schema()
			assert.Equal(t, 3, len(dfSchema.Fields()))

			// Test Columns
			columns := df.Columns()
			assert.Len(t, columns, 3)

			// Test Record
			dfRecord := df.Record()
			assert.NotNil(t, dfRecord)

			// Test String representation
			str := df.String()
			assert.NotEmpty(t, str)

			// Release everything
			intArray.Release()
			floatArray.Release()
			strArray.Release()
			df.Release()
		}
	})

	t.Run("SeriesEdgeCases", func(t *testing.T) {
		// Test Series with edge case data
		testCases := []struct {
			name     string
			dataType arrow.DataType
			dataSize int
			hasNulls bool
		}{
			{"ZeroLengthSeries", arrow.PrimitiveTypes.Int64, 0, false},
			{"SingleElementSeries", arrow.PrimitiveTypes.Float64, 1, false},
			{"LargeSeriesNoNulls", arrow.BinaryTypes.String, 100, false},
			{"SmallSeriesWithNulls", arrow.PrimitiveTypes.Int32, 5, true},
			{"BooleanSeriesWithNulls", arrow.FixedWidthTypes.Boolean, 10, true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var arr arrow.Array

				switch tc.dataType {
				case arrow.PrimitiveTypes.Int64:
					builder := array.NewInt64Builder(pool)
					for i := 0; i < tc.dataSize; i++ {
						if tc.hasNulls && i%2 == 0 {
							builder.AppendNull()
						} else {
							builder.Append(int64(i))
						}
					}
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Int32:
					builder := array.NewInt32Builder(pool)
					for i := 0; i < tc.dataSize; i++ {
						if tc.hasNulls && i%2 == 0 {
							builder.AppendNull()
						} else {
							builder.Append(int32(i))
						}
					}
					arr = builder.NewArray()
				case arrow.PrimitiveTypes.Float64:
					builder := array.NewFloat64Builder(pool)
					for i := 0; i < tc.dataSize; i++ {
						if tc.hasNulls && i%2 == 0 {
							builder.AppendNull()
						} else {
							builder.Append(float64(i) * 1.5)
						}
					}
					arr = builder.NewArray()
				case arrow.BinaryTypes.String:
					builder := array.NewStringBuilder(pool)
					for i := 0; i < tc.dataSize; i++ {
						if tc.hasNulls && i%2 == 0 {
							builder.AppendNull()
						} else {
							builder.AppendString(string(rune('A' + (i % 26))))
						}
					}
					arr = builder.NewArray()
				case arrow.FixedWidthTypes.Boolean:
					builder := array.NewBooleanBuilder(pool)
					for i := 0; i < tc.dataSize; i++ {
						if tc.hasNulls && i%2 == 0 {
							builder.AppendNull()
						} else {
							builder.Append(i%2 == 1)
						}
					}
					arr = builder.NewArray()
				}

				defer arr.Release()

				field := arrow.Field{Name: tc.name, Type: tc.dataType, Nullable: tc.hasNulls}
				series := &Series{array: arr, field: field}

				// Test all Series methods
				assert.Equal(t, tc.dataSize, series.Len())
				assert.Equal(t, tc.name, series.Field().Name)
				assert.Equal(t, tc.dataType, series.Field().Type)
				assert.Equal(t, tc.hasNulls, series.Field().Nullable)

				// Test Array method
				retrievedArray := series.Array()
				assert.NotNil(t, retrievedArray)
				assert.Equal(t, tc.dataSize, retrievedArray.Len())
				assert.Equal(t, tc.dataType, retrievedArray.DataType())

				if tc.hasNulls && tc.dataSize > 0 {
					assert.True(t, retrievedArray.NullN() > 0)
				}
			})
		}
	})
}
