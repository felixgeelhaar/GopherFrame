package io

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// EnhancedCSVWriter provides advanced CSV writing with date formatting capabilities
type EnhancedCSVWriter struct {
	dateParser *DateParser
	options    *CSVWriteOptions
}

// NewEnhancedCSVWriter creates a new enhanced CSV writer
func NewEnhancedCSVWriter(options *CSVWriteOptions) *EnhancedCSVWriter {
	if options == nil {
		options = DefaultCSVWriteOptions()
	}

	return &EnhancedCSVWriter{
		dateParser: NewDateParser(),
		options:    options,
	}
}

// WriteCSVWithDates writes a DataFrame to CSV with enhanced date formatting
func (writer *EnhancedCSVWriter) WriteCSVWithDates(df *core.DataFrame, filePath string) error {
	// Validate file path
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("invalid file path: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return writer.WriteCSVToWriterWithDates(df, file)
}

// WriteCSVToWriterWithDates writes DataFrame to an io.Writer with date formatting
func (writer *EnhancedCSVWriter) WriteCSVToWriterWithDates(df *core.DataFrame, w io.Writer) error {
	csvWriter := csv.NewWriter(w)
	csvWriter.Comma = writer.options.Delimiter
	defer csvWriter.Flush()

	schema := df.Schema()
	numRows := int(df.NumRows())
	numCols := len(schema.Fields())

	// Write header if requested
	if writer.options.WriteHeader {
		headers := make([]string, numCols)
		for i, field := range schema.Fields() {
			headers[i] = field.Name
		}
		if err := csvWriter.Write(headers); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}
	}

	// Process each row
	for row := 0; row < numRows; row++ {
		record := make([]string, numCols)

		for col := 0; col < numCols; col++ {
			field := schema.Field(col)
			series, err := df.ColumnAt(col)
			if err != nil {
				return fmt.Errorf("failed to get column %d: %w", col, err)
			}
			column := series.Array()

			value, err := writer.formatCellValue(column, row, field.Type)
			if err != nil {
				return fmt.Errorf("failed to format value at row %d, column %d (%s): %w", row, col, field.Name, err)
			}

			record[col] = writer.applyCellQuoting(value, field.Type)
		}

		if err := csvWriter.Write(record); err != nil {
			return fmt.Errorf("failed to write row %d: %w", row, err)
		}
	}

	return nil
}

// formatCellValue formats a single cell value based on its Arrow type
func (writer *EnhancedCSVWriter) formatCellValue(column arrow.Array, row int, dataType arrow.DataType) (string, error) {
	if column.IsNull(row) {
		return writer.options.NullValue, nil
	}

	switch dataType.ID() {
	case arrow.INT64:
		intArray := column.(*array.Int64)
		return strconv.FormatInt(intArray.Value(row), 10), nil

	case arrow.FLOAT64:
		floatArray := column.(*array.Float64)
		return strconv.FormatFloat(floatArray.Value(row), 'f', -1, 64), nil

	case arrow.STRING:
		stringArray := column.(*array.String)
		return stringArray.Value(row), nil

	case arrow.BOOL:
		boolArray := column.(*array.Boolean)
		return strconv.FormatBool(boolArray.Value(row)), nil

	case arrow.DATE32:
		dateArray := column.(*array.Date32)
		days := dateArray.Value(row)
		return writer.dateParser.FormatDateValue(int32(days), writer.options.DateFormat), nil

	case arrow.DATE64:
		dateArray := column.(*array.Date64)
		ms := int64(dateArray.Value(row))
		// Convert milliseconds to days
		days := int32(ms / (1000 * 60 * 60 * 24))
		return writer.dateParser.FormatDateValue(days, writer.options.DateFormat), nil

	case arrow.TIMESTAMP:
		timestampArray := column.(*array.Timestamp)
		timestampType := dataType.(*arrow.TimestampType)
		timestamp := timestampArray.Value(row)
		return writer.dateParser.FormatTimestampValue(timestamp, timestampType.Unit, writer.options.TimestampFormat), nil

	default:
		return "", fmt.Errorf("unsupported Arrow type for CSV export: %s", dataType)
	}
}

// applyCellQuoting applies quoting rules based on the quote style and data type
func (writer *EnhancedCSVWriter) applyCellQuoting(value string, dataType arrow.DataType) string {
	switch writer.options.QuoteStyle {
	case QuoteAll:
		return value // csv.Writer will handle quoting

	case QuoteNone:
		return value // No additional quoting

	case QuoteNonNumeric:
		if isNumericType(dataType) {
			return value
		}
		return value // csv.Writer will handle quoting for non-numeric

	case QuoteMinimal:
		fallthrough
	default:
		return value // csv.Writer handles minimal quoting automatically
	}
}

// isNumericType checks if an Arrow type is numeric
func isNumericType(dataType arrow.DataType) bool {
	switch dataType.ID() {
	case arrow.INT8, arrow.INT16, arrow.INT32, arrow.INT64:
		return true
	case arrow.UINT8, arrow.UINT16, arrow.UINT32, arrow.UINT64:
		return true
	case arrow.FLOAT32, arrow.FLOAT64:
		return true
	default:
		return false
	}
}

// WriteCSVWithCustomFormats writes CSV with custom date/time formats for specific columns
func (writer *EnhancedCSVWriter) WriteCSVWithCustomFormats(df *core.DataFrame, filePath string, columnFormats map[string]string) error {
	// Store original formats
	originalDateFormat := writer.options.DateFormat
	originalTimestampFormat := writer.options.TimestampFormat

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = writer.options.Delimiter
	defer csvWriter.Flush()

	schema := df.Schema()
	numRows := int(df.NumRows())
	numCols := len(schema.Fields())

	// Write header if requested
	if writer.options.WriteHeader {
		headers := make([]string, numCols)
		for i, field := range schema.Fields() {
			headers[i] = field.Name
		}
		if err := csvWriter.Write(headers); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}
	}

	// Process each row
	for row := 0; row < numRows; row++ {
		record := make([]string, numCols)

		for col := 0; col < numCols; col++ {
			field := schema.Field(col)
			series, err := df.ColumnAt(col)
			if err != nil {
				return fmt.Errorf("failed to get column %d: %w", col, err)
			}
			column := series.Array()

			// Check for custom format for this column
			if customFormat, exists := columnFormats[field.Name]; exists {
				// Temporarily update format for this column
				switch field.Type.ID() {
				case arrow.DATE32, arrow.DATE64:
					writer.options.DateFormat = customFormat
				case arrow.TIMESTAMP:
					writer.options.TimestampFormat = customFormat
				}
			}

			value, err := writer.formatCellValue(column, row, field.Type)
			if err != nil {
				return fmt.Errorf("failed to format value at row %d, column %d (%s): %w", row, col, field.Name, err)
			}

			record[col] = writer.applyCellQuoting(value, field.Type)

			// Restore original formats
			writer.options.DateFormat = originalDateFormat
			writer.options.TimestampFormat = originalTimestampFormat
		}

		if err := csvWriter.Write(record); err != nil {
			return fmt.Errorf("failed to write row %d: %w", row, err)
		}
	}

	return nil
}

// GetColumnTypeInfo returns information about the data types in a DataFrame for CSV export planning
func (writer *EnhancedCSVWriter) GetColumnTypeInfo(df *core.DataFrame) []ColumnExportInfo {
	schema := df.Schema()
	info := make([]ColumnExportInfo, len(schema.Fields()))

	for i, field := range schema.Fields() {
		info[i] = ColumnExportInfo{
			Name:            field.Name,
			ArrowType:       field.Type,
			BaseType:        getBaseTypeFromArrow(field.Type),
			RequiresQuoting: !isNumericType(field.Type),
			SuggestedFormat: writer.getSuggestedFormat(field.Type),
		}
	}

	return info
}

// ColumnExportInfo provides information about a column for CSV export
type ColumnExportInfo struct {
	Name            string
	ArrowType       arrow.DataType
	BaseType        string
	RequiresQuoting bool
	SuggestedFormat string
}

// getSuggestedFormat returns a suggested format string for a given Arrow type
func (writer *EnhancedCSVWriter) getSuggestedFormat(dataType arrow.DataType) string {
	switch dataType.ID() {
	case arrow.DATE32, arrow.DATE64:
		return writer.options.DateFormat
	case arrow.TIMESTAMP:
		return writer.options.TimestampFormat
	case arrow.FLOAT64:
		return "%.6f" // Example precision format
	case arrow.INT64:
		return "%d"
	default:
		return "%s"
	}
}

// EstimateCSVSize estimates the output size of a DataFrame when written to CSV
func (writer *EnhancedCSVWriter) EstimateCSVSize(df *core.DataFrame) (int64, error) {
	schema := df.Schema()
	numRows := df.NumRows()
	numCols := int64(len(schema.Fields()))

	if numRows == 0 {
		return 0, nil
	}

	// Estimate average character width per column type
	var totalWidth int64

	for _, field := range schema.Fields() {
		var avgWidth int64

		switch field.Type.ID() {
		case arrow.INT64:
			avgWidth = 15 // Average int64 string length
		case arrow.FLOAT64:
			avgWidth = 20 // Average float64 string length with decimals
		case arrow.STRING:
			avgWidth = 25 // Conservative estimate for strings
		case arrow.DATE32, arrow.DATE64:
			avgWidth = int64(len(writer.options.DateFormat)) + 2
		case arrow.TIMESTAMP:
			avgWidth = int64(len(writer.options.TimestampFormat)) + 2
		case arrow.BOOL:
			avgWidth = 5 // "true" or "false"
		default:
			avgWidth = 20 // Conservative default
		}

		totalWidth += avgWidth
	}

	// Add overhead for delimiters, quotes, and newlines
	overhead := numCols + 2 // Delimiters and newline per row
	estimatedRowSize := totalWidth + overhead

	totalSize := estimatedRowSize * numRows

	// Add header size if applicable
	if writer.options.WriteHeader {
		headerSize := int64(0)
		for _, field := range schema.Fields() {
			headerSize += int64(len(field.Name)) + 1 // Name + delimiter/newline
		}
		totalSize += headerSize
	}

	return totalSize, nil
}
