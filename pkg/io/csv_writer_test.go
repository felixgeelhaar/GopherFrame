package io

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnhancedCSVWriter_WriteCSVWithDates(t *testing.T) {
	df := createTestDataFrameWithDates()
	defer df.Release()

	t.Run("BasicWriteWithDates", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "test_output_*.csv")
		require.NoError(t, err)
		defer os.Remove(tempFile.Name())
		tempFile.Close()

		writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())
		err = writer.WriteCSVWithDates(df, tempFile.Name())
		require.NoError(t, err)

		// Read back and verify
		content, err := os.ReadFile(tempFile.Name())
		require.NoError(t, err)

		lines := strings.Split(string(content), "\n")
		assert.Contains(t, lines[0], "name,birth_date,score,active") // Header
		assert.Contains(t, lines[1], "John,2023-06-15,85.5,true")    // First row
		assert.Contains(t, lines[2], "Jane,2023-12-31,92,false")     // Second row
	})

	t.Run("WriteToWriter", func(t *testing.T) {
		var buf bytes.Buffer
		writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())

		err := writer.WriteCSVToWriterWithDates(df, &buf)
		require.NoError(t, err)

		content := buf.String()
		lines := strings.Split(content, "\n")
		assert.Contains(t, lines[0], "name,birth_date,score,active")
		assert.Contains(t, lines[1], "John,2023-06-15,85.5,true")
	})

	t.Run("CustomDateFormat", func(t *testing.T) {
		var buf bytes.Buffer
		options := DefaultCSVWriteOptions()
		options.DateFormat = "01/02/2006" // US format

		writer := NewEnhancedCSVWriter(options)
		err := writer.WriteCSVToWriterWithDates(df, &buf)
		require.NoError(t, err)

		content := buf.String()
		assert.Contains(t, content, "06/15/2023") // US formatted date
		assert.Contains(t, content, "12/31/2023")
	})

	t.Run("NoHeader", func(t *testing.T) {
		var buf bytes.Buffer
		options := DefaultCSVWriteOptions()
		options.WriteHeader = false

		writer := NewEnhancedCSVWriter(options)
		err := writer.WriteCSVToWriterWithDates(df, &buf)
		require.NoError(t, err)

		content := buf.String()
		lines := strings.Split(content, "\n")
		// First line should be data, not header
		assert.Contains(t, lines[0], "John,2023-06-15,85.5,true")
		assert.NotContains(t, content, "name,birth_date")
	})

	t.Run("CustomDelimiter", func(t *testing.T) {
		var buf bytes.Buffer
		options := DefaultCSVWriteOptions()
		options.Delimiter = ';'

		writer := NewEnhancedCSVWriter(options)
		err := writer.WriteCSVToWriterWithDates(df, &buf)
		require.NoError(t, err)

		content := buf.String()
		assert.Contains(t, content, "name;birth_date;score;active")
		assert.Contains(t, content, "John;2023-06-15;85.5;true")
	})

	t.Run("CustomNullValue", func(t *testing.T) {
		// Create DataFrame with null values
		dfWithNulls := createTestDataFrameWithNulls()
		defer dfWithNulls.Release()

		var buf bytes.Buffer
		options := DefaultCSVWriteOptions()
		options.NullValue = "N/A"

		writer := NewEnhancedCSVWriter(options)
		err := writer.WriteCSVToWriterWithDates(dfWithNulls, &buf)
		require.NoError(t, err)

		content := buf.String()
		assert.Contains(t, content, "N/A") // Null values should be replaced
	})
}

func TestEnhancedCSVWriter_WriteCSVWithCustomFormats(t *testing.T) {
	df := createTestDataFrameWithTimestamps()
	defer df.Release()

	tempFile, err := os.CreateTemp("", "test_custom_*.csv")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())

	// Custom formats for specific columns
	columnFormats := map[string]string{
		"birth_date": "01/02/2006",           // US date format
		"created_at": "2006-01-02T15:04:05Z", // ISO 8601 timestamp
	}

	err = writer.WriteCSVWithCustomFormats(df, tempFile.Name(), columnFormats)
	require.NoError(t, err)

	// Read back and verify
	content, err := os.ReadFile(tempFile.Name())
	require.NoError(t, err)

	contentStr := string(content)
	assert.Contains(t, contentStr, "06/15/2023")           // US date format
	assert.Contains(t, contentStr, "2023-06-15T14:30:00Z") // ISO timestamp
}

func TestEnhancedCSVWriter_FormatCellValue(t *testing.T) {
	pool := memory.NewGoAllocator()
	writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())

	t.Run("FormatInt64", func(t *testing.T) {
		builder := array.NewInt64Builder(pool)
		builder.Append(12345)
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.PrimitiveTypes.Int64)
		require.NoError(t, err)
		assert.Equal(t, "12345", result)
	})

	t.Run("FormatFloat64", func(t *testing.T) {
		builder := array.NewFloat64Builder(pool)
		builder.Append(123.456)
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.PrimitiveTypes.Float64)
		require.NoError(t, err)
		assert.Equal(t, "123.456", result)
	})

	t.Run("FormatString", func(t *testing.T) {
		builder := array.NewStringBuilder(pool)
		builder.Append("hello world")
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.BinaryTypes.String)
		require.NoError(t, err)
		assert.Equal(t, "hello world", result)
	})

	t.Run("FormatBoolean", func(t *testing.T) {
		builder := array.NewBooleanBuilder(pool)
		builder.Append(true)
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.FixedWidthTypes.Boolean)
		require.NoError(t, err)
		assert.Equal(t, "true", result)
	})

	t.Run("FormatDate32", func(t *testing.T) {
		builder := array.NewDate32Builder(pool)
		// 2023-06-15 = 19523 days since epoch
		builder.Append(arrow.Date32(19523))
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.FixedWidthTypes.Date32)
		require.NoError(t, err)
		assert.Equal(t, "2023-06-15", result)
	})

	t.Run("FormatDate64", func(t *testing.T) {
		builder := array.NewDate64Builder(pool)
		// Convert days to milliseconds: 19523 * 24 * 60 * 60 * 1000
		ms := int64(19523) * 24 * 60 * 60 * 1000
		builder.Append(arrow.Date64(ms))
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.FixedWidthTypes.Date64)
		require.NoError(t, err)
		assert.Equal(t, "2023-06-15", result)
	})

	t.Run("FormatTimestamp", func(t *testing.T) {
		timestampType := &arrow.TimestampType{Unit: arrow.Millisecond}
		builder := array.NewTimestampBuilder(pool, timestampType)

		// 2023-06-15 14:30:00
		testTime := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
		builder.Append(arrow.Timestamp(testTime.UnixMilli()))
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, timestampType)
		require.NoError(t, err)
		assert.Equal(t, "2023-06-15 14:30:00", result)
	})

	t.Run("FormatNullValue", func(t *testing.T) {
		builder := array.NewStringBuilder(pool)
		builder.AppendNull()
		arr := builder.NewArray()
		defer arr.Release()

		result, err := writer.formatCellValue(arr, 0, arrow.BinaryTypes.String)
		require.NoError(t, err)
		assert.Equal(t, "", result) // Default null value
	})

	t.Run("UnsupportedType", func(t *testing.T) {
		builder := array.NewStringBuilder(pool)
		builder.Append("test")
		arr := builder.NewArray()
		defer arr.Release()

		// Use an unsupported type
		_, err := writer.formatCellValue(arr, 0, arrow.ListOf(arrow.PrimitiveTypes.Int32))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported Arrow type for CSV export")
	})
}

func TestEnhancedCSVWriter_QuoteStyles(t *testing.T) {
	writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())

	tests := []struct {
		name       string
		quoteStyle QuoteStyle
		value      string
		dataType   arrow.DataType
		expected   string
	}{
		{"QuoteAll", QuoteAll, "hello", arrow.BinaryTypes.String, "hello"},
		{"QuoteNone", QuoteNone, "hello", arrow.BinaryTypes.String, "hello"},
		{"QuoteNonNumeric_String", QuoteNonNumeric, "hello", arrow.BinaryTypes.String, "hello"},
		{"QuoteNonNumeric_Int", QuoteNonNumeric, "123", arrow.PrimitiveTypes.Int64, "123"},
		{"QuoteMinimal", QuoteMinimal, "hello", arrow.BinaryTypes.String, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer.options.QuoteStyle = tt.quoteStyle
			result := writer.applyCellQuoting(tt.value, tt.dataType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEnhancedCSVWriter_GetColumnTypeInfo(t *testing.T) {
	df := createTestDataFrameWithDates()
	defer df.Release()

	writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())
	typeInfo := writer.GetColumnTypeInfo(df)

	assert.Equal(t, 4, len(typeInfo))

	// Check name column (string)
	assert.Equal(t, "name", typeInfo[0].Name)
	assert.Equal(t, "string", typeInfo[0].BaseType)
	assert.True(t, typeInfo[0].RequiresQuoting)

	// Check birth_date column (date)
	assert.Equal(t, "birth_date", typeInfo[1].Name)
	assert.Equal(t, "date", typeInfo[1].BaseType)
	assert.True(t, typeInfo[1].RequiresQuoting)
	assert.Equal(t, "2006-01-02", typeInfo[1].SuggestedFormat)

	// Check score column (float)
	assert.Equal(t, "score", typeInfo[2].Name)
	assert.Equal(t, "float64", typeInfo[2].BaseType)
	assert.False(t, typeInfo[2].RequiresQuoting)

	// Check active column (bool)
	assert.Equal(t, "active", typeInfo[3].Name)
	assert.Equal(t, "string", typeInfo[3].BaseType) // Bool gets mapped to string base type
	assert.True(t, typeInfo[3].RequiresQuoting)
}

func TestEnhancedCSVWriter_EstimateCSVSize(t *testing.T) {
	df := createTestDataFrameWithDates()
	defer df.Release()

	writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())

	size, err := writer.EstimateCSVSize(df)
	require.NoError(t, err)

	assert.Greater(t, size, int64(0), "Estimated size should be positive")
	assert.Less(t, size, int64(10000), "Estimated size should be reasonable for test data")
}

func TestEnhancedCSVWriter_EstimateCSVSize_EmptyDataFrame(t *testing.T) {
	// Create empty DataFrame
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "col1", Type: arrow.BinaryTypes.String},
	}, nil)

	stringBuilder := array.NewStringBuilder(pool)
	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	record := array.NewRecord(schema, []arrow.Array{stringArray}, 0)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	writer := NewEnhancedCSVWriter(DefaultCSVWriteOptions())
	size, err := writer.EstimateCSVSize(df)
	require.NoError(t, err)

	assert.Equal(t, int64(0), size, "Empty DataFrame should have 0 estimated size")
}

func TestIsNumericType(t *testing.T) {
	tests := []struct {
		dataType arrow.DataType
		expected bool
	}{
		{arrow.PrimitiveTypes.Int8, true},
		{arrow.PrimitiveTypes.Int16, true},
		{arrow.PrimitiveTypes.Int32, true},
		{arrow.PrimitiveTypes.Int64, true},
		{arrow.PrimitiveTypes.Uint8, true},
		{arrow.PrimitiveTypes.Uint16, true},
		{arrow.PrimitiveTypes.Uint32, true},
		{arrow.PrimitiveTypes.Uint64, true},
		{arrow.PrimitiveTypes.Float32, true},
		{arrow.PrimitiveTypes.Float64, true},
		{arrow.BinaryTypes.String, false},
		{arrow.FixedWidthTypes.Boolean, false},
		{arrow.FixedWidthTypes.Date32, false},
	}

	for _, tt := range tests {
		t.Run(tt.dataType.String(), func(t *testing.T) {
			result := isNumericType(tt.dataType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDefaultCSVWriteOptions(t *testing.T) {
	options := DefaultCSVWriteOptions()

	assert.True(t, options.WriteHeader)
	assert.Equal(t, ',', options.Delimiter)
	assert.Equal(t, "2006-01-02", options.DateFormat)
	assert.Equal(t, "2006-01-02 15:04:05", options.TimestampFormat)
	assert.Equal(t, "", options.NullValue)
	assert.Equal(t, QuoteMinimal, options.QuoteStyle)
}

// Helper functions to create test DataFrames

func createTestDataFrameWithDates() *core.DataFrame {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "birth_date", Type: arrow.FixedWidthTypes.Date32},
		{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		{Name: "active", Type: arrow.FixedWidthTypes.Boolean},
	}, nil)

	// Create string array
	stringBuilder := array.NewStringBuilder(pool)
	stringBuilder.AppendValues([]string{"John", "Jane"}, nil)
	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	// Create date array (2023-06-15, 2023-12-31)
	dateBuilder := array.NewDate32Builder(pool)
	dateBuilder.AppendValues([]arrow.Date32{19523, 19722}, nil)
	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()

	// Create float array
	floatBuilder := array.NewFloat64Builder(pool)
	floatBuilder.AppendValues([]float64{85.5, 92.0}, nil)
	floatArray := floatBuilder.NewArray()
	defer floatArray.Release()

	// Create boolean array
	boolBuilder := array.NewBooleanBuilder(pool)
	boolBuilder.AppendValues([]bool{true, false}, nil)
	boolArray := boolBuilder.NewArray()
	defer boolArray.Release()

	record := array.NewRecord(schema, []arrow.Array{stringArray, dateArray, floatArray, boolArray}, 2)
	defer record.Release()

	return core.NewDataFrame(record)
}

func createTestDataFrameWithNulls() *core.DataFrame {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "birth_date", Type: arrow.FixedWidthTypes.Date32},
	}, nil)

	// Create string array with null
	stringBuilder := array.NewStringBuilder(pool)
	stringBuilder.Append("John")
	stringBuilder.AppendNull() // Second value is null
	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	// Create date array with null
	dateBuilder := array.NewDate32Builder(pool)
	dateBuilder.Append(arrow.Date32(19523))
	dateBuilder.AppendNull() // Second value is null
	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()

	record := array.NewRecord(schema, []arrow.Array{stringArray, dateArray}, 2)
	defer record.Release()

	return core.NewDataFrame(record)
}

func createTestDataFrameWithTimestamps() *core.DataFrame {
	pool := memory.NewGoAllocator()

	timestampType := &arrow.TimestampType{Unit: arrow.Millisecond}
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "birth_date", Type: arrow.FixedWidthTypes.Date32},
		{Name: "created_at", Type: timestampType},
	}, nil)

	// Create string array
	stringBuilder := array.NewStringBuilder(pool)
	stringBuilder.AppendValues([]string{"John", "Jane"}, nil)
	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	// Create date array
	dateBuilder := array.NewDate32Builder(pool)
	dateBuilder.AppendValues([]arrow.Date32{19523, 19722}, nil)
	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()

	// Create timestamp array
	timestampBuilder := array.NewTimestampBuilder(pool, timestampType)
	testTime1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	testTime2 := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
	timestampBuilder.AppendValues([]arrow.Timestamp{
		arrow.Timestamp(testTime1.UnixMilli()),
		arrow.Timestamp(testTime2.UnixMilli()),
	}, nil)
	timestampArray := timestampBuilder.NewArray()
	defer timestampArray.Release()

	record := array.NewRecord(schema, []arrow.Array{stringArray, dateArray, timestampArray}, 2)
	defer record.Release()

	return core.NewDataFrame(record)
}
