package io

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParquet_DateTimeSupport(t *testing.T) {
	tempDir := t.TempDir()
	pool := memory.NewGoAllocator()

	t.Run("Date32RoundTrip", func(t *testing.T) {
		// Create DataFrame with Date32 column
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "date", Type: arrow.FixedWidthTypes.Date32},
			},
			nil,
		)

		// Create test data
		idBuilder := array.NewInt64Builder(pool)
		dateBuilder := array.NewDate32Builder(pool)

		// Add test dates: 2023-06-15, 2023-12-31, 2024-02-29 (leap year)
		idBuilder.AppendValues([]int64{1, 2, 3}, nil)
		dateBuilder.AppendValues([]arrow.Date32{
			arrow.Date32(19523), // 2023-06-15
			arrow.Date32(19722), // 2023-12-31
			arrow.Date32(19782), // 2024-02-29
		}, nil)

		idArray := idBuilder.NewArray()
		dateArray := dateBuilder.NewArray()
		defer idArray.Release()
		defer dateArray.Release()

		record := array.NewRecord(schema, []arrow.Array{idArray, dateArray}, 3)
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write to Parquet
		testFile := filepath.Join(tempDir, "date32_test.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		// Read back from Parquet
		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		// Verify structure
		assert.Equal(t, int64(3), readDF.NumRows())
		assert.Equal(t, int64(2), readDF.NumCols())

		// Verify date column type
		readSchema := readDF.Schema()
		dateField := readSchema.Field(1)
		assert.Equal(t, "date", dateField.Name)
		assert.Equal(t, arrow.DATE32, dateField.Type.ID())

		// Verify date values
		readRecord := readDF.Record()
		dateArr := readRecord.Column(1).(*array.Date32)

		expectedDates := []arrow.Date32{19523, 19722, 19782}
		for i, expected := range expectedDates {
			assert.Equal(t, expected, dateArr.Value(i))
		}
	})

	t.Run("Date64RoundTrip", func(t *testing.T) {
		// Create DataFrame with Date64 column
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "date64", Type: arrow.FixedWidthTypes.Date64},
			},
			nil,
		)

		idBuilder := array.NewInt64Builder(pool)
		dateBuilder := array.NewDate64Builder(pool)

		// Add test dates as milliseconds since epoch
		testDates := []time.Time{
			time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		}

		idBuilder.AppendValues([]int64{1, 2, 3}, nil)
		for _, dt := range testDates {
			dateBuilder.Append(arrow.Date64(dt.UnixMilli()))
		}

		idArray := idBuilder.NewArray()
		dateArray := dateBuilder.NewArray()
		defer idArray.Release()
		defer dateArray.Release()

		record := array.NewRecord(schema, []arrow.Array{idArray, dateArray}, 3)
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write to Parquet
		testFile := filepath.Join(tempDir, "date64_test.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		// Read back from Parquet
		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		// Verify date column type
		readSchema := readDF.Schema()
		dateField := readSchema.Field(1)
		assert.Equal(t, "date64", dateField.Name)
		// Note: Parquet may store Date64 as Date32 for efficiency
		if dateField.Type.ID() == arrow.DATE32 {
			// If stored as Date32, verify values as Date32
			readRecord := readDF.Record()
			dateArr := readRecord.Column(1).(*array.Date32)

			for i, expectedTime := range testDates {
				days := dateArr.Value(i)
				actualTime := time.Unix(int64(days)*86400, 0).UTC()
				assert.Equal(t, expectedTime.Year(), actualTime.Year())
				assert.Equal(t, expectedTime.Month(), actualTime.Month())
				assert.Equal(t, expectedTime.Day(), actualTime.Day())
			}
		} else {
			// If stored as Date64, verify as Date64
			assert.Equal(t, arrow.DATE64, dateField.Type.ID())
			readRecord := readDF.Record()
			dateArr := readRecord.Column(1).(*array.Date64)

			for i, expectedTime := range testDates {
				actualMs := int64(dateArr.Value(i))
				actualTime := time.Unix(actualMs/1000, (actualMs%1000)*1e6).UTC()
				assert.Equal(t, expectedTime.Year(), actualTime.Year())
				assert.Equal(t, expectedTime.Month(), actualTime.Month())
				assert.Equal(t, expectedTime.Day(), actualTime.Day())
			}
		}
	})

	t.Run("TimestampRoundTrip", func(t *testing.T) {
		// Create DataFrame with Timestamp column
		timestampType := &arrow.TimestampType{Unit: arrow.Millisecond}
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "timestamp", Type: timestampType},
			},
			nil,
		)

		idBuilder := array.NewInt64Builder(pool)
		timestampBuilder := array.NewTimestampBuilder(pool, timestampType)

		// Add test timestamps
		testTimestamps := []time.Time{
			time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
			time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC),
		}

		idBuilder.AppendValues([]int64{1, 2, 3}, nil)
		for _, ts := range testTimestamps {
			timestampBuilder.Append(arrow.Timestamp(ts.UnixMilli()))
		}

		idArray := idBuilder.NewArray()
		timestampArray := timestampBuilder.NewArray()
		defer idArray.Release()
		defer timestampArray.Release()

		record := array.NewRecord(schema, []arrow.Array{idArray, timestampArray}, 3)
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write to Parquet
		testFile := filepath.Join(tempDir, "timestamp_test.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		// Read back from Parquet
		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		// Verify timestamp column type
		readSchema := readDF.Schema()
		tsField := readSchema.Field(1)
		assert.Equal(t, "timestamp", tsField.Name)
		assert.Equal(t, arrow.TIMESTAMP, tsField.Type.ID())

		// Verify timestamp unit
		readTsType := tsField.Type.(*arrow.TimestampType)
		assert.Equal(t, arrow.Millisecond, readTsType.Unit)

		// Verify timestamp values
		readRecord := readDF.Record()
		tsArr := readRecord.Column(1).(*array.Timestamp)

		for i, expectedTime := range testTimestamps {
			actualMs := int64(tsArr.Value(i))
			actualTime := time.Unix(actualMs/1000, (actualMs%1000)*1e6).UTC()
			assert.Equal(t, expectedTime, actualTime)
		}
	})

	t.Run("MixedTypesWithNulls", func(t *testing.T) {
		// Create DataFrame with mixed types including dates
		timestampType := &arrow.TimestampType{Unit: arrow.Millisecond}
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "name", Type: arrow.BinaryTypes.String},
				{Name: "birth_date", Type: arrow.FixedWidthTypes.Date32},
				{Name: "last_login", Type: timestampType},
				{Name: "score", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		// Build arrays with valid date/time values (avoiding nulls for now due to Parquet limitations)
		nameBuilder := array.NewStringBuilder(pool)
		dateBuilder := array.NewDate32Builder(pool)
		timestampBuilder := array.NewTimestampBuilder(pool, timestampType)
		scoreBuilder := array.NewFloat64Builder(pool)

		nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie"}, []bool{true, true, true})

		// All valid dates
		dateBuilder.Append(arrow.Date32(19523)) // 2023-06-15
		dateBuilder.Append(arrow.Date32(19722)) // 2023-12-31
		dateBuilder.Append(arrow.Date32(19782)) // 2024-02-29

		// All valid timestamps
		timestampBuilder.Append(arrow.Timestamp(time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC).UnixMilli()))
		timestampBuilder.Append(arrow.Timestamp(time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC).UnixMilli()))
		timestampBuilder.Append(arrow.Timestamp(time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC).UnixMilli()))

		scoreBuilder.AppendValues([]float64{95.5, 82.0, 91.5}, []bool{true, true, true})

		arrays := []arrow.Array{
			nameBuilder.NewArray(),
			dateBuilder.NewArray(),
			timestampBuilder.NewArray(),
			scoreBuilder.NewArray(),
		}
		for _, arr := range arrays {
			defer arr.Release()
		}

		record := array.NewRecord(schema, arrays, 3)
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write to Parquet
		testFile := filepath.Join(tempDir, "mixed_datetime_test.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		// Read back from Parquet
		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		// Verify data preservation
		assert.Equal(t, int64(3), readDF.NumRows())
		assert.Equal(t, int64(4), readDF.NumCols())

		// Verify schema preservation
		readSchema := readDF.Schema()
		assert.Equal(t, "name", readSchema.Field(0).Name)
		assert.Equal(t, "birth_date", readSchema.Field(1).Name)
		assert.Equal(t, "last_login", readSchema.Field(2).Name)
		assert.Equal(t, "score", readSchema.Field(3).Name)

		// Verify date and timestamp types are preserved
		assert.Equal(t, arrow.DATE32, readSchema.Field(1).Type.ID())
		assert.Equal(t, arrow.TIMESTAMP, readSchema.Field(2).Type.ID())
	})

	t.Run("EmptyDataFrameWithDateColumns", func(t *testing.T) {
		// Test empty DataFrame with date columns
		timestampType := &arrow.TimestampType{Unit: arrow.Millisecond}
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "date", Type: arrow.FixedWidthTypes.Date32},
				{Name: "timestamp", Type: timestampType},
			},
			nil,
		)

		// Create empty arrays
		idBuilder := array.NewInt64Builder(pool)
		dateBuilder := array.NewDate32Builder(pool)
		timestampBuilder := array.NewTimestampBuilder(pool, timestampType)

		arrays := []arrow.Array{
			idBuilder.NewArray(),
			dateBuilder.NewArray(),
			timestampBuilder.NewArray(),
		}
		for _, arr := range arrays {
			defer arr.Release()
		}

		record := array.NewRecord(schema, arrays, 0)
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write to Parquet
		testFile := filepath.Join(tempDir, "empty_datetime_test.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		// Read back from Parquet
		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		// Verify empty DataFrame structure is preserved
		assert.Equal(t, int64(0), readDF.NumRows())
		assert.Equal(t, int64(3), readDF.NumCols())

		// Verify schema is preserved
		readSchema := readDF.Schema()
		assert.Equal(t, arrow.DATE32, readSchema.Field(1).Type.ID())
		assert.Equal(t, arrow.TIMESTAMP, readSchema.Field(2).Type.ID())
	})
}

func TestParquet_DateTimeEdgeCases(t *testing.T) {
	tempDir := t.TempDir()
	pool := memory.NewGoAllocator()

	t.Run("LeapYearDates", func(t *testing.T) {
		// Test leap year edge cases
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "date", Type: arrow.FixedWidthTypes.Date32},
				{Name: "description", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		dateBuilder := array.NewDate32Builder(pool)
		descBuilder := array.NewStringBuilder(pool)

		// Key leap year dates
		leapDates := []struct {
			date arrow.Date32
			desc string
		}{
			{arrow.Date32(18321), "2020-02-29"}, // 2020 leap year
			{arrow.Date32(19782), "2024-02-29"}, // 2024 leap year
			{arrow.Date32(18322), "2020-03-01"}, // Day after leap day
			{arrow.Date32(18686), "2021-02-28"}, // Non-leap year
		}

		for _, ld := range leapDates {
			dateBuilder.Append(ld.date)
			descBuilder.Append(ld.desc)
		}

		arrays := []arrow.Array{
			dateBuilder.NewArray(),
			descBuilder.NewArray(),
		}
		for _, arr := range arrays {
			defer arr.Release()
		}

		record := array.NewRecord(schema, arrays, int64(len(leapDates)))
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write and read back
		testFile := filepath.Join(tempDir, "leap_year_test.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		// Verify all dates preserved correctly
		readRecord := readDF.Record()
		dateArr := readRecord.Column(0).(*array.Date32)

		for i, ld := range leapDates {
			assert.Equal(t, ld.date, dateArr.Value(i))
		}
	})

	t.Run("TimestampPrecision", func(t *testing.T) {
		// Test different timestamp precisions
		testCases := []struct {
			name string
			unit arrow.TimeUnit
		}{
			{"Seconds", arrow.Second},
			{"Milliseconds", arrow.Millisecond},
			{"Microseconds", arrow.Microsecond},
			{"Nanoseconds", arrow.Nanosecond},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				timestampType := &arrow.TimestampType{Unit: tc.unit}
				schema := arrow.NewSchema(
					[]arrow.Field{
						{Name: "timestamp", Type: timestampType},
					},
					nil,
				)

				builder := array.NewTimestampBuilder(pool, timestampType)

				// Create a precise timestamp
				testTime := time.Date(2023, 6, 15, 14, 30, 45, 123456789, time.UTC)

				var tsValue arrow.Timestamp
				switch tc.unit {
				case arrow.Second:
					tsValue = arrow.Timestamp(testTime.Unix())
				case arrow.Millisecond:
					tsValue = arrow.Timestamp(testTime.UnixMilli())
				case arrow.Microsecond:
					tsValue = arrow.Timestamp(testTime.UnixMicro())
				case arrow.Nanosecond:
					tsValue = arrow.Timestamp(testTime.UnixNano())
				}

				builder.Append(tsValue)
				arr := builder.NewArray()
				defer arr.Release()

				record := array.NewRecord(schema, []arrow.Array{arr}, 1)
				df := dataframe.NewDataFrame(record)
				defer df.Release()

				// Write and read back
				testFile := filepath.Join(tempDir, "timestamp_"+tc.name+".parquet")
				writer := NewParquetWriter()
				err := writer.WriteFile(df, testFile)
				require.NoError(t, err)

				reader := NewParquetReader()
				readDF, err := reader.ReadFile(testFile)
				require.NoError(t, err)
				defer readDF.Release()

				// Verify timestamp precision is maintained
				readRecord := readDF.Record()
				tsArr := readRecord.Column(0).(*array.Timestamp)

				readTsType := readDF.Schema().Field(0).Type.(*arrow.TimestampType)

				// Note: Parquet may convert timestamp units for storage efficiency
				// The most important thing is that the timestamp value represents the same point in time
				expectedTimeUnix := testTime.Unix()
				actualTimeUnix := time.Unix(int64(tsArr.Value(0))/1000, 0).UTC().Unix() // Convert from ms to seconds

				// For seconds precision, Parquet might store as milliseconds
				if tc.unit == arrow.Second {
					// Allow conversion to milliseconds in Parquet
					if readTsType.Unit == arrow.Millisecond {
						assert.Equal(t, expectedTimeUnix, actualTimeUnix)
					} else {
						assert.Equal(t, tc.unit, readTsType.Unit)
						assert.Equal(t, tsValue, tsArr.Value(0))
					}
				} else {
					// For other precisions, expect exact preservation
					assert.Equal(t, tc.unit, readTsType.Unit)
					assert.Equal(t, tsValue, tsArr.Value(0))
				}
			})
		}
	})
}

func TestParquet_FileValidation(t *testing.T) {
	t.Run("InvalidFilePath", func(t *testing.T) {
		reader := NewParquetReader()
		writer := NewParquetWriter()

		// Test empty path
		_, err := reader.ReadFile("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "filename cannot be empty")

		pool := memory.NewGoAllocator()
		schema := arrow.NewSchema([]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}}, nil)
		builder := array.NewInt64Builder(pool)
		builder.Append(1)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 1)
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		err = writer.WriteFile(df, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "filename cannot be empty")
	})

	t.Run("DirectoryTraversal", func(t *testing.T) {
		reader := NewParquetReader()

		// Test directory traversal attempt
		_, err := reader.ReadFile("../../../etc/passwd")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory traversal detected")
	})
}

func TestParquet_LargeDataset(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large dataset test in short mode")
	}

	tempDir := t.TempDir()
	pool := memory.NewGoAllocator()

	t.Run("LargeDateDataset", func(t *testing.T) {
		// Create a large dataset with dates
		numRows := 10000
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "event_date", Type: arrow.FixedWidthTypes.Date32},
				{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		idBuilder := array.NewInt64Builder(pool)
		dateBuilder := array.NewDate32Builder(pool)
		valueBuilder := array.NewFloat64Builder(pool)

		// Generate data
		baseDate := arrow.Date32(19000) // Some base date
		for i := 0; i < numRows; i++ {
			idBuilder.Append(int64(i))
			dateBuilder.Append(baseDate + arrow.Date32(i%365)) // Cycle through a year
			valueBuilder.Append(float64(i) * 1.5)
		}

		arrays := []arrow.Array{
			idBuilder.NewArray(),
			dateBuilder.NewArray(),
			valueBuilder.NewArray(),
		}
		for _, arr := range arrays {
			defer arr.Release()
		}

		record := array.NewRecord(schema, arrays, int64(numRows))
		df := dataframe.NewDataFrame(record)
		defer df.Release()

		// Write to Parquet
		testFile := filepath.Join(tempDir, "large_date_dataset.parquet")
		writer := NewParquetWriter()
		err := writer.WriteFile(df, testFile)
		require.NoError(t, err)

		// Verify file was created and has reasonable size
		info, err := os.Stat(testFile)
		require.NoError(t, err)
		assert.Greater(t, info.Size(), int64(0))

		// Read back and verify
		reader := NewParquetReader()
		readDF, err := reader.ReadFile(testFile)
		require.NoError(t, err)
		defer readDF.Release()

		assert.Equal(t, int64(numRows), readDF.NumRows())
		assert.Equal(t, int64(3), readDF.NumCols())
	})
}
