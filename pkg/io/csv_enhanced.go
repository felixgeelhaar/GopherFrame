package io

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// CSVReadOptions provides configuration for enhanced CSV reading with date support
type CSVReadOptions struct {
	HasHeader         bool
	Delimiter         rune
	Comment           rune
	SampleRows        int
	DateFormats       []string // Custom date formats to try (Go layout strings)
	AutoDetectDates   bool     // Whether to automatically detect date columns
	StrictDateParsing bool     // Whether to require all values in a date column to parse
	NullValues        []string // Values to treat as null
	SkipRows          int      // Number of rows to skip at the beginning
	MaxRows           int      // Maximum number of rows to read (0 = all)
	DefaultTimezone   string   // Default timezone for timestamp parsing
	TimestampFormats  []string // Custom timestamp formats with timezone info
}

// DefaultCSVReadOptions returns sensible defaults for CSV reading
func DefaultCSVReadOptions() *CSVReadOptions {
	return &CSVReadOptions{
		HasHeader:         true,
		Delimiter:         ',',
		Comment:           '#',
		SampleRows:        100,
		AutoDetectDates:   true,
		StrictDateParsing: false,
		NullValues:        []string{"", "NULL", "null", "N/A", "n/a", "NA", "na"},
		SkipRows:          0,
		MaxRows:           0,
		DefaultTimezone:   "UTC",
		TimestampFormats:  []string{"2006-01-02T15:04:05Z", "2006-01-02 15:04:05", "2006-01-02T15:04:05"},
	}
}

// CSVWriteOptions provides configuration for enhanced CSV writing with date support
type CSVWriteOptions struct {
	WriteHeader     bool
	Delimiter       rune
	DateFormat      string // Go layout string for date formatting
	TimestampFormat string // Go layout string for timestamp formatting
	NullValue       string // String to use for null values
	QuoteStyle      QuoteStyle
	OutputTimezone  string // Timezone for timestamp output (empty means keep original)
	IncludeTimezone bool   // Whether to include timezone info in timestamp output
}

// QuoteStyle defines how to quote CSV fields
type QuoteStyle int

const (
	QuoteMinimal    QuoteStyle = iota // Quote only when necessary
	QuoteAll                          // Quote all fields
	QuoteNone                         // Never quote (may cause issues)
	QuoteNonNumeric                   // Quote all non-numeric fields
)

// DefaultCSVWriteOptions returns sensible defaults for CSV writing
func DefaultCSVWriteOptions() *CSVWriteOptions {
	return &CSVWriteOptions{
		WriteHeader:     true,
		Delimiter:       ',',
		DateFormat:      "2006-01-02",          // ISO 8601 date
		TimestampFormat: "2006-01-02 15:04:05", // ISO 8601 timestamp
		NullValue:       "",
		QuoteStyle:      QuoteMinimal,
		OutputTimezone:  "",    // Keep original timezone
		IncludeTimezone: false, // Don't include timezone by default
	}
}

// EnhancedCSVReader provides advanced CSV reading with date parsing capabilities
type EnhancedCSVReader struct {
	dateParser *DateParser
	options    *CSVReadOptions
}

// NewEnhancedCSVReader creates a new enhanced CSV reader
func NewEnhancedCSVReader(options *CSVReadOptions) *EnhancedCSVReader {
	if options == nil {
		options = DefaultCSVReadOptions()
	}

	var dateParser *DateParser
	if options.DefaultTimezone != "" {
		var err error
		dateParser, err = NewDateParserWithTimezone(options.DefaultTimezone)
		if err != nil {
			// Fall back to default parser if timezone is invalid
			dateParser = NewDateParser()
		}
	} else {
		dateParser = NewDateParser()
	}

	reader := &EnhancedCSVReader{
		dateParser: dateParser,
		options:    options,
	}

	// Add custom date formats if provided
	for _, format := range options.DateFormats {
		reader.dateParser.AddDateFormat("custom", format, "User provided")
	}

	// Add custom timestamp formats if provided
	for _, format := range options.TimestampFormats {
		reader.dateParser.AddDateTimeFormat("custom", format, "User provided timestamp")
	}

	if options.StrictDateParsing {
		if options.DefaultTimezone != "" {
			strictParser, err := NewDateParserWithTimezone(options.DefaultTimezone)
			if err == nil {
				reader.dateParser = strictParser
				reader.dateParser.strictMode = true
			}
		} else {
			reader.dateParser = NewStrictDateParser()
		}
	}

	return reader
}

// ReadCSVWithDates reads a CSV file with enhanced date parsing capabilities
func (reader *EnhancedCSVReader) ReadCSVWithDates(filePath string, pool memory.Allocator) (*core.DataFrame, error) {
	if pool == nil {
		pool = memory.NewGoAllocator()
	}

	// Validate file path for security
	if err := validateFilePath(filePath); err != nil {
		return nil, fmt.Errorf("invalid file path: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return reader.ReadCSVFromReaderWithDates(file, pool)
}

// ReadCSVFromReaderWithDates reads CSV data from an io.Reader with date parsing
func (reader *EnhancedCSVReader) ReadCSVFromReaderWithDates(r io.Reader, pool memory.Allocator) (*core.DataFrame, error) {
	csvReader := csv.NewReader(r)
	csvReader.Comma = reader.options.Delimiter
	csvReader.Comment = reader.options.Comment

	// Read all records
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	// Skip initial rows if requested
	if reader.options.SkipRows > 0 && len(records) > reader.options.SkipRows {
		records = records[reader.options.SkipRows:]
	}

	// Extract headers
	var headers []string
	startIndex := 0

	if reader.options.HasHeader && len(records) > 0 {
		headers = records[0]
		startIndex = 1
	} else {
		// Generate default headers
		if len(records) > 0 {
			headers = make([]string, len(records[0]))
			for i := range headers {
				headers[i] = fmt.Sprintf("column_%d", i)
			}
		}
	}

	if startIndex >= len(records) {
		return nil, fmt.Errorf("no data rows found after header")
	}

	// Limit rows if requested
	endIndex := len(records)
	if reader.options.MaxRows > 0 && startIndex+reader.options.MaxRows < len(records) {
		endIndex = startIndex + reader.options.MaxRows
	}

	dataRecords := records[startIndex:endIndex]

	// Prepare column data
	if len(dataRecords) == 0 {
		return nil, fmt.Errorf("no data rows to process")
	}

	numCols := len(headers)
	columnData := make([][]string, numCols)

	for _, record := range dataRecords {
		// Pad short records with empty strings
		for i := 0; i < numCols; i++ {
			if i < len(record) {
				columnData[i] = append(columnData[i], record[i])
			} else {
				columnData[i] = append(columnData[i], "")
			}
		}
	}

	// Infer column types with date detection
	columnTypes := make([]ColumnTypeInfo, numCols)
	for i, data := range columnData {
		columnTypes[i] = reader.inferColumnTypeWithDates(data)
	}

	// Build Arrow arrays
	arrays := make([]arrow.Array, numCols)
	fields := make([]arrow.Field, numCols)

	for i, colType := range columnTypes {
		fields[i] = arrow.Field{Name: headers[i], Type: colType.ArrowType}

		var err error
		arrays[i], err = reader.buildColumnArrayWithDates(columnData[i], colType, pool)
		if err != nil {
			// Clean up arrays created so far
			for j := 0; j < i; j++ {
				arrays[j].Release()
			}
			return nil, fmt.Errorf("failed to build array for column '%s': %w", headers[i], err)
		}
	}

	// Create schema and record
	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, arrays, int64(len(dataRecords)))

	// Clean up arrays (record retains them)
	for _, arr := range arrays {
		arr.Release()
	}

	return core.NewDataFrame(record), nil
}

// ColumnTypeInfo holds information about a column's inferred type
type ColumnTypeInfo struct {
	BaseType       string // "int64", "float64", "string", "date", "timestamp"
	ArrowType      arrow.DataType
	DateFormat     string  // Go layout string for dates
	DateConfidence float64 // Confidence score for date detection
	NullCount      int     // Number of null values detected
}

// inferColumnTypeWithDates infers the type of a column including date detection
func (reader *EnhancedCSVReader) inferColumnTypeWithDates(values []string) ColumnTypeInfo {
	if len(values) == 0 {
		return ColumnTypeInfo{
			BaseType:  "string",
			ArrowType: arrow.BinaryTypes.String,
		}
	}

	// Filter out null values for type inference
	nonNullValues := make([]string, 0, len(values))
	nullCount := 0

	for _, value := range values {
		if reader.isNullValue(value) {
			nullCount++
		} else {
			nonNullValues = append(nonNullValues, value)
		}
	}

	if len(nonNullValues) == 0 {
		return ColumnTypeInfo{
			BaseType:  "string",
			ArrowType: arrow.BinaryTypes.String,
			NullCount: nullCount,
		}
	}

	// Try date detection first if enabled
	if reader.options.AutoDetectDates {
		if dateResult, err := reader.dateParser.InferDateType(nonNullValues); err == nil && dateResult.IsDate {
			return ColumnTypeInfo{
				BaseType:       getBaseTypeFromArrow(dateResult.ArrowType),
				ArrowType:      dateResult.ArrowType,
				DateFormat:     dateResult.GoLayout,
				DateConfidence: dateResult.Confidence,
				NullCount:      nullCount,
			}
		}
	}

	// Fall back to numeric type inference
	sampleSize := min(len(nonNullValues), reader.options.SampleRows)
	sample := nonNullValues[:sampleSize]

	// Try int64
	intCount := 0
	for _, value := range sample {
		if _, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64); err == nil {
			intCount++
		}
	}

	if intCount == len(sample) {
		return ColumnTypeInfo{
			BaseType:  "int64",
			ArrowType: arrow.PrimitiveTypes.Int64,
			NullCount: nullCount,
		}
	}

	// Try float64
	floatCount := 0
	for _, value := range sample {
		if _, err := strconv.ParseFloat(strings.TrimSpace(value), 64); err == nil {
			floatCount++
		}
	}

	if floatCount == len(sample) {
		return ColumnTypeInfo{
			BaseType:  "float64",
			ArrowType: arrow.PrimitiveTypes.Float64,
			NullCount: nullCount,
		}
	}

	// Default to string
	return ColumnTypeInfo{
		BaseType:  "string",
		ArrowType: arrow.BinaryTypes.String,
		NullCount: nullCount,
	}
}

// buildColumnArrayWithDates builds an Arrow array with support for date types
func (reader *EnhancedCSVReader) buildColumnArrayWithDates(values []string, colType ColumnTypeInfo, pool memory.Allocator) (arrow.Array, error) {
	switch colType.BaseType {
	case "date":
		return reader.dateParser.ParseDateColumnWithNullChecker(values, colType.DateFormat, pool, reader.isNullValue)
	case "timestamp":
		return reader.dateParser.ParseTimestampColumnWithNullChecker(values, colType.DateFormat, pool, reader.isNullValue)
	case "int64":
		return reader.buildInt64Array(values, pool)
	case "float64":
		return reader.buildFloat64Array(values, pool)
	default:
		return reader.buildStringArray(values, pool)
	}
}

// isNullValue checks if a value should be treated as null
func (reader *EnhancedCSVReader) isNullValue(value string) bool {
	trimmed := strings.TrimSpace(value)
	for _, nullVal := range reader.options.NullValues {
		if trimmed == nullVal {
			return true
		}
	}
	return false
}

// Helper method implementations for basic types
func (reader *EnhancedCSVReader) buildInt64Array(values []string, pool memory.Allocator) (arrow.Array, error) {
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	builder.Reserve(len(values))

	for _, value := range values {
		if reader.isNullValue(value) {
			builder.AppendNull()
		} else {
			parsed, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse int64 value '%s': %w", value, err)
			}
			builder.Append(parsed)
		}
	}

	return builder.NewArray(), nil
}

func (reader *EnhancedCSVReader) buildFloat64Array(values []string, pool memory.Allocator) (arrow.Array, error) {
	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	builder.Reserve(len(values))

	for _, value := range values {
		if reader.isNullValue(value) {
			builder.AppendNull()
		} else {
			parsed, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float64 value '%s': %w", value, err)
			}
			builder.Append(parsed)
		}
	}

	return builder.NewArray(), nil
}

func (reader *EnhancedCSVReader) buildStringArray(values []string, pool memory.Allocator) (arrow.Array, error) {
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Reserve(len(values))

	for _, value := range values {
		if reader.isNullValue(value) {
			builder.AppendNull()
		} else {
			builder.Append(value)
		}
	}

	return builder.NewArray(), nil
}

// getBaseTypeFromArrow converts Arrow type to base type string
func getBaseTypeFromArrow(arrowType arrow.DataType) string {
	switch arrowType.ID() {
	case arrow.DATE32, arrow.DATE64:
		return "date"
	case arrow.TIMESTAMP:
		return "timestamp"
	case arrow.INT64:
		return "int64"
	case arrow.FLOAT64:
		return "float64"
	default:
		return "string"
	}
}

// validateFilePath provides basic path validation (placeholder for existing function)
func validateFilePath(path string) error {
	// This should reference the existing validation function
	// For now, just check if path is not empty
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("file path cannot be empty")
	}
	return nil
}
