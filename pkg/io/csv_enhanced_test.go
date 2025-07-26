package io

import (
	"os"
	"strings"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnhancedCSVReader_ReadCSVWithDates(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("BasicDateParsing", func(t *testing.T) {
		// Create temporary CSV file with dates
		csvContent := `name,birth_date,score
John,2023-06-15,85.5
Jane,2023-12-31,92.0
Bob,2024-02-29,78.5`

		tempFile := createTempCSVFile(t, csvContent)
		defer os.Remove(tempFile)

		reader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		df, err := reader.ReadCSVWithDates(tempFile, pool)
		require.NoError(t, err)
		defer df.Release()

		// Verify structure
		assert.Equal(t, int64(3), df.NumRows())
		assert.Equal(t, int64(3), df.NumCols())

		schema := df.Schema()
		fields := schema.Fields()
		assert.Equal(t, "name", fields[0].Name)
		assert.Equal(t, "birth_date", fields[1].Name)
		assert.Equal(t, "score", fields[2].Name)

		// Verify date column was detected and parsed
		assert.Equal(t, arrow.STRING, fields[0].Type.ID())
		assert.Equal(t, arrow.DATE32, fields[1].Type.ID())
		assert.Equal(t, arrow.FLOAT64, fields[2].Type.ID())

		// Verify date values
		dateSeries, err := df.ColumnAt(1)
		require.NoError(t, err)
		dateColumn := dateSeries.Array().(*array.Date32)
		assert.False(t, dateColumn.IsNull(0))
		assert.False(t, dateColumn.IsNull(1))
		assert.False(t, dateColumn.IsNull(2))
	})

	t.Run("MixedDateFormats", func(t *testing.T) {
		csvContent := `event,us_date,euro_date
Event1,06/15/2023,15/06/2023
Event2,12/31/2023,31/12/2023`

		reader := strings.NewReader(csvContent)
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		schema := df.Schema()
		fields := schema.Fields()

		// Both date columns should be detected as dates
		assert.Equal(t, arrow.STRING, fields[0].Type.ID()) // event
		assert.Equal(t, arrow.DATE32, fields[1].Type.ID()) // us_date
		assert.Equal(t, arrow.DATE32, fields[2].Type.ID()) // euro_date
	})

	t.Run("WithNullValues", func(t *testing.T) {
		csvContent := `name,date,value
John,2023-06-15,100
Jane,,200
Bob,NULL,300
Alice,2023-12-31,400`

		reader := strings.NewReader(csvContent)
		options := DefaultCSVReadOptions()
		options.NullValues = append(options.NullValues, "NULL")

		csvReader := NewEnhancedCSVReader(options)
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		// Verify null handling in date column
		dateSeries, err := df.ColumnAt(1)
		require.NoError(t, err)
		dateColumn := dateSeries.Array().(*array.Date32)
		assert.False(t, dateColumn.IsNull(0)) // John
		assert.True(t, dateColumn.IsNull(1))  // Jane (empty)
		assert.True(t, dateColumn.IsNull(2))  // Bob (NULL)
		assert.False(t, dateColumn.IsNull(3)) // Alice
	})

	t.Run("CustomDateFormats", func(t *testing.T) {
		csvContent := `name,custom_date
John,2023.06.15
Jane,2023.12.31`

		reader := strings.NewReader(csvContent)
		options := DefaultCSVReadOptions()
		options.DateFormats = []string{"2006.01.02"}

		csvReader := NewEnhancedCSVReader(options)
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		schema := df.Schema()
		fields := schema.Fields()
		assert.Equal(t, arrow.DATE32, fields[1].Type.ID())
	})

	t.Run("DisabledAutoDetection", func(t *testing.T) {
		csvContent := `name,date_str
John,2023-06-15
Jane,2023-12-31`

		reader := strings.NewReader(csvContent)
		options := DefaultCSVReadOptions()
		options.AutoDetectDates = false

		csvReader := NewEnhancedCSVReader(options)
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		schema := df.Schema()
		fields := schema.Fields()
		// Date column should remain as string
		assert.Equal(t, arrow.STRING, fields[1].Type.ID())
	})

	t.Run("MaxRowsLimit", func(t *testing.T) {
		csvContent := `name,date
John,2023-06-15
Jane,2023-12-31
Bob,2024-02-29
Alice,2024-03-15`

		reader := strings.NewReader(csvContent)
		options := DefaultCSVReadOptions()
		options.MaxRows = 2

		csvReader := NewEnhancedCSVReader(options)
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		assert.Equal(t, int64(2), df.NumRows())
	})

	t.Run("SkipRows", func(t *testing.T) {
		// Test data: 1 extra row to skip, then header, then 2 data rows
		csvContent := `extra,row
name,date
John,2023-06-15
Jane,2023-12-31`

		reader := strings.NewReader(csvContent)
		options := DefaultCSVReadOptions()
		options.SkipRows = 1 // Skip the extra row

		csvReader := NewEnhancedCSVReader(options)
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		// Should have 2 data rows after header
		assert.Equal(t, int64(2), df.NumRows())

		// Verify we have the right column names from the header
		columnNames := df.ColumnNames()
		assert.Equal(t, []string{"name", "date"}, columnNames)
	})
}

func TestEnhancedCSVReader_TypeInference(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("NumericTypeInference", func(t *testing.T) {
		csvContent := `int_col,float_col,string_col
123,45.67,hello
456,78.90,world
789,12.34,test`

		reader := strings.NewReader(csvContent)
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		schema := df.Schema()
		fields := schema.Fields()

		assert.Equal(t, arrow.INT64, fields[0].Type.ID())
		assert.Equal(t, arrow.FLOAT64, fields[1].Type.ID())
		assert.Equal(t, arrow.STRING, fields[2].Type.ID())
	})

	t.Run("MixedTypesDefaultToString", func(t *testing.T) {
		csvContent := `mixed_col
123
hello
45.67
world`

		reader := strings.NewReader(csvContent)
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		require.NoError(t, err)
		defer df.Release()

		schema := df.Schema()
		fields := schema.Fields()

		// Mixed types should default to string
		assert.Equal(t, arrow.STRING, fields[0].Type.ID())
	})
}

func TestEnhancedCSVReader_ErrorHandling(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("EmptyFile", func(t *testing.T) {
		reader := strings.NewReader("")
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		_, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CSV file is empty")
	})

	t.Run("HeaderOnlyFile", func(t *testing.T) {
		reader := strings.NewReader("name,date")
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		_, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no data rows found after header")
	})

	t.Run("InvalidFilePath", func(t *testing.T) {
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		_, err := csvReader.ReadCSVWithDates("", pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "file path cannot be empty")
	})

	t.Run("NonexistentFile", func(t *testing.T) {
		csvReader := NewEnhancedCSVReader(DefaultCSVReadOptions())
		_, err := csvReader.ReadCSVWithDates("/nonexistent/file.csv", pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to open file")
	})
}

func TestEnhancedCSVReader_StrictDateParsing(t *testing.T) {
	pool := memory.NewGoAllocator()

	csvContent := `name,mixed_date
John,2023-06-15
Jane,not-a-date
Bob,2023-12-31`

	reader := strings.NewReader(csvContent)
	options := DefaultCSVReadOptions()
	options.StrictDateParsing = true

	csvReader := NewEnhancedCSVReader(options)
	df, err := csvReader.ReadCSVFromReaderWithDates(reader, pool)
	require.NoError(t, err)
	defer df.Release()

	schema := df.Schema()
	fields := schema.Fields()

	// In strict mode, mixed date/non-date should remain as string
	assert.Equal(t, arrow.STRING, fields[1].Type.ID())
}

func TestColumnTypeInfo(t *testing.T) {
	reader := NewEnhancedCSVReader(DefaultCSVReadOptions())

	t.Run("DateTypeInfo", func(t *testing.T) {
		values := []string{"2023-06-15", "2023-12-31", "2024-02-29"}
		typeInfo := reader.inferColumnTypeWithDates(values)

		assert.Equal(t, "date", typeInfo.BaseType)
		assert.Equal(t, arrow.DATE32, typeInfo.ArrowType.ID())
		assert.Equal(t, "2006-01-02", typeInfo.DateFormat)
		assert.GreaterOrEqual(t, typeInfo.DateConfidence, 0.8)
		assert.Equal(t, 0, typeInfo.NullCount)
	})

	t.Run("IntTypeInfo", func(t *testing.T) {
		values := []string{"123", "456", "789"}
		typeInfo := reader.inferColumnTypeWithDates(values)

		assert.Equal(t, "int64", typeInfo.BaseType)
		assert.Equal(t, arrow.INT64, typeInfo.ArrowType.ID())
		assert.Equal(t, 0, typeInfo.NullCount)
	})

	t.Run("FloatTypeInfo", func(t *testing.T) {
		values := []string{"12.34", "56.78", "90.12"}
		typeInfo := reader.inferColumnTypeWithDates(values)

		assert.Equal(t, "float64", typeInfo.BaseType)
		assert.Equal(t, arrow.FLOAT64, typeInfo.ArrowType.ID())
		assert.Equal(t, 0, typeInfo.NullCount)
	})

	t.Run("StringTypeInfo", func(t *testing.T) {
		values := []string{"hello", "world", "test"}
		typeInfo := reader.inferColumnTypeWithDates(values)

		assert.Equal(t, "string", typeInfo.BaseType)
		assert.Equal(t, arrow.STRING, typeInfo.ArrowType.ID())
		assert.Equal(t, 0, typeInfo.NullCount)
	})

	t.Run("WithNulls", func(t *testing.T) {
		values := []string{"123", "", "NULL", "456"}
		typeInfo := reader.inferColumnTypeWithDates(values)

		assert.Equal(t, "int64", typeInfo.BaseType)
		assert.Equal(t, 2, typeInfo.NullCount) // "" and "NULL" are nulls
	})
}

func TestGetBaseTypeFromArrow(t *testing.T) {
	tests := []struct {
		arrowType    arrow.DataType
		expectedBase string
	}{
		{arrow.FixedWidthTypes.Date32, "date"},
		{arrow.FixedWidthTypes.Date64, "date"},
		{arrow.FixedWidthTypes.Timestamp_ms, "timestamp"},
		{arrow.PrimitiveTypes.Int64, "int64"},
		{arrow.PrimitiveTypes.Float64, "float64"},
		{arrow.BinaryTypes.String, "string"},
	}

	for _, tt := range tests {
		t.Run(tt.expectedBase, func(t *testing.T) {
			result := getBaseTypeFromArrow(tt.arrowType)
			assert.Equal(t, tt.expectedBase, result)
		})
	}
}

// Helper function to create temporary CSV files for testing
func createTempCSVFile(t *testing.T, content string) string {
	tempFile, err := os.CreateTemp("", "test_*.csv")
	require.NoError(t, err)

	_, err = tempFile.WriteString(content)
	require.NoError(t, err)

	err = tempFile.Close()
	require.NoError(t, err)

	return tempFile.Name()
}

func TestDefaultCSVReadOptions(t *testing.T) {
	options := DefaultCSVReadOptions()

	assert.True(t, options.HasHeader)
	assert.Equal(t, ',', options.Delimiter)
	assert.Equal(t, '#', options.Comment)
	assert.Equal(t, 100, options.SampleRows)
	assert.True(t, options.AutoDetectDates)
	assert.False(t, options.StrictDateParsing)
	assert.Contains(t, options.NullValues, "")
	assert.Contains(t, options.NullValues, "NULL")
	assert.Equal(t, 0, options.SkipRows)
	assert.Equal(t, 0, options.MaxRows)
}

func TestEnhancedCSVReader_BuildArrays(t *testing.T) {
	pool := memory.NewGoAllocator()
	reader := NewEnhancedCSVReader(DefaultCSVReadOptions())

	t.Run("BuildInt64Array", func(t *testing.T) {
		values := []string{"123", "456", "", "789"}

		result, err := reader.buildInt64Array(values, pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.INT64, result.DataType().ID())
		assert.Equal(t, 4, result.Len())

		intArray := result.(*array.Int64)
		assert.False(t, intArray.IsNull(0))
		assert.Equal(t, int64(123), intArray.Value(0))
		assert.True(t, intArray.IsNull(2)) // Empty string should be null
	})

	t.Run("BuildFloat64Array", func(t *testing.T) {
		values := []string{"12.34", "56.78", "", "90.12"}

		result, err := reader.buildFloat64Array(values, pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.FLOAT64, result.DataType().ID())
		floatArray := result.(*array.Float64)
		assert.Equal(t, 12.34, floatArray.Value(0))
		assert.True(t, floatArray.IsNull(2))
	})

	t.Run("BuildStringArray", func(t *testing.T) {
		values := []string{"hello", "world", "", "test"}

		result, err := reader.buildStringArray(values, pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.STRING, result.DataType().ID())
		stringArray := result.(*array.String)
		assert.Equal(t, "hello", stringArray.Value(0))
		assert.True(t, stringArray.IsNull(2))
	})

	t.Run("BuildArrayWithInvalidData", func(t *testing.T) {
		values := []string{"123", "not-a-number", "456"}

		_, err := reader.buildInt64Array(values, pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse int64 value")
	})
}
