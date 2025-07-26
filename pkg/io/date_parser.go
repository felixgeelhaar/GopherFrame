package io

import (
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// DateFormat represents different date format patterns for parsing
type DateFormat struct {
	Pattern     string
	GoLayout    string
	Description string
}

// Common date formats for CSV parsing
var CommonDateFormats = []DateFormat{
	{"YYYY-MM-DD", "2006-01-02", "ISO 8601 date"},
	{"MM/DD/YYYY", "01/02/2006", "US format"},
	{"DD/MM/YYYY", "02/01/2006", "European format"},
	{"DD-MM-YYYY", "02-01-2006", "European with dashes"},
	{"YYYY/MM/DD", "2006/01/02", "ISO with slashes"},
	{"Mon DD, YYYY", "Jan 02, 2006", "Long month name"},
	{"DD Mon YYYY", "02 Jan 2006", "Day month year"},
	{"YYYYMMDD", "20060102", "Compact format"},
	{"MM-DD-YYYY", "01-02-2006", "US with dashes"},
}

// DateTimeFormat represents date-time format patterns
type DateTimeFormat struct {
	Pattern     string
	GoLayout    string
	Description string
}

// Common date-time formats
var CommonDateTimeFormats = []DateTimeFormat{
	{"YYYY-MM-DD HH:MM:SS", "2006-01-02 15:04:05", "ISO 8601 datetime"},
	{"YYYY-MM-DD HH:MM:SS.sss", "2006-01-02 15:04:05.000", "ISO with milliseconds"},
	{"YYYY-MM-DDTHH:MM:SS", "2006-01-02T15:04:05", "ISO 8601 with T"},
	{"YYYY-MM-DDTHH:MM:SSZ", "2006-01-02T15:04:05Z", "ISO 8601 UTC"},
	{"MM/DD/YYYY HH:MM:SS", "01/02/2006 15:04:05", "US datetime"},
	{"DD/MM/YYYY HH:MM:SS", "02/01/2006 15:04:05", "European datetime"},
	{"YYYY-MM-DD HH:MM", "2006-01-02 15:04", "ISO without seconds"},
	{"Mon DD, YYYY HH:MM:SS", "Jan 02, 2006 15:04:05", "Long format"},
}

// DateParser provides comprehensive date parsing capabilities
type DateParser struct {
	dateFormats     []DateFormat
	dateTimeFormats []DateTimeFormat
	strictMode      bool
	timezoneHandler *TimezoneHandler
}

// NewDateParser creates a new date parser with default formats
func NewDateParser() *DateParser {
	return &DateParser{
		dateFormats:     CommonDateFormats,
		dateTimeFormats: CommonDateTimeFormats,
		strictMode:      false,
		timezoneHandler: NewTimezoneHandler(),
	}
}

// NewStrictDateParser creates a parser that requires exact format matches
func NewStrictDateParser() *DateParser {
	return &DateParser{
		dateFormats:     CommonDateFormats,
		dateTimeFormats: CommonDateTimeFormats,
		strictMode:      true,
		timezoneHandler: NewTimezoneHandler(),
	}
}

// NewDateParserWithTimezone creates a date parser with custom timezone settings
func NewDateParserWithTimezone(defaultTimezone string) (*DateParser, error) {
	tzHandler, err := NewTimezoneHandlerWithDefault(defaultTimezone)
	if err != nil {
		return nil, err
	}

	return &DateParser{
		dateFormats:     CommonDateFormats,
		dateTimeFormats: CommonDateTimeFormats,
		strictMode:      false,
		timezoneHandler: tzHandler,
	}, nil
}

// AddDateFormat adds a custom date format to the parser
func (dp *DateParser) AddDateFormat(pattern, goLayout, description string) {
	dp.dateFormats = append(dp.dateFormats, DateFormat{
		Pattern:     pattern,
		GoLayout:    goLayout,
		Description: description,
	})
}

// AddDateTimeFormat adds a custom date-time format to the parser
func (dp *DateParser) AddDateTimeFormat(pattern, goLayout, description string) {
	dp.dateTimeFormats = append(dp.dateTimeFormats, DateTimeFormat{
		Pattern:     pattern,
		GoLayout:    goLayout,
		Description: description,
	})
}

// InferDateType analyzes a column of string values to determine if they represent dates
func (dp *DateParser) InferDateType(values []string) (DateParseResult, error) {
	if len(values) == 0 {
		return DateParseResult{IsDate: false}, nil
	}

	// Sample up to 100 values for type inference
	sampleSize := min(len(values), 100)
	sample := values[:sampleSize]

	// Try each date format
	for _, dateFormat := range dp.dateFormats {
		if result := dp.tryDateFormat(sample, dateFormat, false); result.IsDate {
			return result, nil
		}
	}

	// Try each date-time format
	for _, dateTimeFormat := range dp.dateTimeFormats {
		if result := dp.tryDateTimeFormat(sample, dateTimeFormat); result.IsDate {
			return result, nil
		}
	}

	return DateParseResult{IsDate: false}, nil
}

// DateParseResult contains the results of date type inference
type DateParseResult struct {
	IsDate      bool
	IsDateTime  bool
	Format      string
	GoLayout    string
	ArrowType   arrow.DataType
	Confidence  float64
	SampleValue time.Time
}

// tryDateFormat attempts to parse a sample using a specific date format
func (dp *DateParser) tryDateFormat(sample []string, format DateFormat, allowPartial bool) DateParseResult {
	successCount := 0
	totalNonEmpty := 0
	var sampleTime time.Time

	for _, value := range sample {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue // Skip empty values
		}
		totalNonEmpty++

		parsed, err := time.Parse(format.GoLayout, trimmed)
		if err == nil {
			successCount++
			if sampleTime.IsZero() {
				sampleTime = parsed
			}
		} else if dp.strictMode {
			return DateParseResult{IsDate: false}
		}
	}

	if totalNonEmpty == 0 {
		return DateParseResult{IsDate: false}
	}

	confidence := float64(successCount) / float64(totalNonEmpty)
	threshold := 0.8 // 80% of values must parse successfully
	if allowPartial {
		threshold = 0.5 // 50% for partial matching
	}

	if confidence >= threshold {
		return DateParseResult{
			IsDate:      true,
			IsDateTime:  false,
			Format:      format.Pattern,
			GoLayout:    format.GoLayout,
			ArrowType:   arrow.FixedWidthTypes.Date32,
			Confidence:  confidence,
			SampleValue: sampleTime,
		}
	}

	return DateParseResult{IsDate: false}
}

// tryDateTimeFormat attempts to parse a sample using a specific date-time format
func (dp *DateParser) tryDateTimeFormat(sample []string, format DateTimeFormat) DateParseResult {
	successCount := 0
	totalNonEmpty := 0
	var sampleTime time.Time

	for _, value := range sample {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		totalNonEmpty++

		parsed, err := time.Parse(format.GoLayout, trimmed)
		if err == nil {
			successCount++
			if sampleTime.IsZero() {
				sampleTime = parsed
			}
		} else if dp.strictMode {
			return DateParseResult{IsDate: false}
		}
	}

	if totalNonEmpty == 0 {
		return DateParseResult{IsDate: false}
	}

	confidence := float64(successCount) / float64(totalNonEmpty)
	if confidence >= 0.8 {
		return DateParseResult{
			IsDate:      true,
			IsDateTime:  true,
			Format:      format.Pattern,
			GoLayout:    format.GoLayout,
			ArrowType:   arrow.FixedWidthTypes.Timestamp_ms,
			Confidence:  confidence,
			SampleValue: sampleTime,
		}
	}

	return DateParseResult{IsDate: false}
}

// ParseDateColumn converts a string column to Arrow date array using the specified format
func (dp *DateParser) ParseDateColumn(values []string, goLayout string, pool memory.Allocator) (arrow.Array, error) {
	return dp.ParseDateColumnWithNullChecker(values, goLayout, pool, dp.isNullValue)
}

// ParseDateColumnWithNullChecker converts a string column to Arrow date array with custom null detection
func (dp *DateParser) ParseDateColumnWithNullChecker(values []string, goLayout string, pool memory.Allocator, isNull func(string) bool) (arrow.Array, error) {
	builder := array.NewDate32Builder(pool)
	defer builder.Release()

	builder.Reserve(len(values))

	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if isNull(value) {
			builder.AppendNull()
			continue
		}

		parsed, err := time.Parse(goLayout, trimmed)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date value '%s' with layout '%s': %w", trimmed, goLayout, err)
		}

		// Convert to days since Unix epoch for Date32
		days := int32(parsed.Unix() / 86400)
		builder.Append(arrow.Date32(days))
	}

	return builder.NewArray(), nil
}

// isNullValue is the default null checker for date parser
func (dp *DateParser) isNullValue(value string) bool {
	trimmed := strings.TrimSpace(value)
	return trimmed == ""
}

// ParseTimestampColumn converts a string column to Arrow timestamp array
func (dp *DateParser) ParseTimestampColumn(values []string, goLayout string, pool memory.Allocator) (arrow.Array, error) {
	return dp.ParseTimestampColumnWithNullChecker(values, goLayout, pool, dp.isNullValue)
}

// ParseTimestampColumnWithNullChecker converts a string column to Arrow timestamp array with custom null detection
func (dp *DateParser) ParseTimestampColumnWithNullChecker(values []string, goLayout string, pool memory.Allocator, isNull func(string) bool) (arrow.Array, error) {
	timestampType := &arrow.TimestampType{Unit: arrow.Millisecond}
	builder := array.NewTimestampBuilder(pool, timestampType)
	defer builder.Release()

	builder.Reserve(len(values))

	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if isNull(value) {
			builder.AppendNull()
			continue
		}

		parsed, err := time.Parse(goLayout, trimmed)
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp value '%s' with layout '%s': %w", trimmed, goLayout, err)
		}

		// Convert to milliseconds since Unix epoch
		ms := parsed.UnixMilli()
		builder.Append(arrow.Timestamp(ms))
	}

	return builder.NewArray(), nil
}

// FormatDateValue converts a Date32 value to string using the specified format
func (dp *DateParser) FormatDateValue(days int32, goLayout string) string {
	// Convert days since epoch to time.Time
	unixTime := int64(days) * 86400
	t := time.Unix(unixTime, 0).UTC()
	return t.Format(goLayout)
}

// FormatTimestampValue converts a timestamp value to string using the specified format
func (dp *DateParser) FormatTimestampValue(ts arrow.Timestamp, unit arrow.TimeUnit, goLayout string) string {
	var unixTime int64

	switch unit {
	case arrow.Second:
		unixTime = int64(ts)
	case arrow.Millisecond:
		unixTime = int64(ts) / 1000
	case arrow.Microsecond:
		unixTime = int64(ts) / 1000000
	case arrow.Nanosecond:
		unixTime = int64(ts) / 1000000000
	default:
		unixTime = int64(ts) / 1000 // Default to milliseconds
	}

	t := time.Unix(unixTime, 0).UTC()
	return t.Format(goLayout)
}

// AutoDetectDateFormat attempts to automatically detect the date format in a column
func (dp *DateParser) AutoDetectDateFormat(values []string) (*DateParseResult, error) {
	result, err := dp.InferDateType(values)
	if err != nil {
		return nil, err
	}

	if !result.IsDate {
		return nil, fmt.Errorf("no recognizable date format found")
	}

	return &result, nil
}

// ParseWithCustomFormat parses dates using a custom Go layout string
func (dp *DateParser) ParseWithCustomFormat(values []string, goLayout string, pool memory.Allocator) (arrow.Array, error) {
	// First verify the format works with a sample
	var testValue string
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			testValue = v
			break
		}
	}

	if testValue != "" {
		_, err := time.Parse(goLayout, strings.TrimSpace(testValue))
		if err != nil {
			return nil, fmt.Errorf("invalid date layout '%s' for sample value '%s': %w", goLayout, testValue, err)
		}
	}

	// Determine if this should be a date or timestamp
	if dp.containsTimeComponents(goLayout) {
		return dp.ParseTimestampColumn(values, goLayout, pool)
	}
	return dp.ParseDateColumn(values, goLayout, pool)
}

// containsTimeComponents checks if a layout contains time components
func (dp *DateParser) containsTimeComponents(layout string) bool {
	timeComponents := []string{"15", "04", "05", "MST", "Z07", "Z0700", "Z07:00", "-07", "-0700", "-07:00"}
	for _, component := range timeComponents {
		if strings.Contains(layout, component) {
			return true
		}
	}
	return false
}

// GetSupportedFormats returns all supported date and datetime formats
func (dp *DateParser) GetSupportedFormats() ([]DateFormat, []DateTimeFormat) {
	return dp.dateFormats, dp.dateTimeFormats
}

// SetDefaultTimezone sets the default timezone for timestamp parsing
func (dp *DateParser) SetDefaultTimezone(timezone string) error {
	return dp.timezoneHandler.SetDefaultTimezone(timezone)
}

// GetDefaultTimezone returns the current default timezone
func (dp *DateParser) GetDefaultTimezone() *time.Location {
	return dp.timezoneHandler.GetDefaultTimezone()
}

// ParseTimestampColumnWithTimezoneParsing parses timestamps with timezone awareness
func (dp *DateParser) ParseTimestampColumnWithTimezoneParsing(
	values []string,
	options TimezoneAwareTimestampOptions,
	pool memory.Allocator,
	isNull func(string) bool,
) (arrow.Array, error) {
	return dp.timezoneHandler.ParseTimestampColumnWithTimezone(values, options, pool, isNull)
}

// FormatTimestampValueWithTimezone formats a timestamp with timezone conversion
func (dp *DateParser) FormatTimestampValueWithTimezone(
	ts arrow.Timestamp,
	unit arrow.TimeUnit,
	goLayout string,
	sourceTimezone string,
	targetTimezone string,
) (string, error) {
	// Convert Arrow timestamp to time.Time
	var t time.Time
	switch unit {
	case arrow.Second:
		t = time.Unix(int64(ts), 0)
	case arrow.Millisecond:
		t = time.Unix(int64(ts)/1000, (int64(ts)%1000)*1000000)
	case arrow.Microsecond:
		t = time.Unix(int64(ts)/1000000, (int64(ts)%1000000)*1000)
	case arrow.Nanosecond:
		t = time.Unix(int64(ts)/1000000000, int64(ts)%1000000000)
	default:
		t = time.Unix(int64(ts)/1000, (int64(ts)%1000)*1000000) // Default to milliseconds
	}

	// Apply source timezone
	if sourceTimezone != "" {
		convertedTime, err := dp.timezoneHandler.ConvertToTimezone(t, sourceTimezone)
		if err != nil {
			return "", err
		}
		t = convertedTime
	}

	// Format with target timezone
	return dp.timezoneHandler.FormatTimestampWithTimezone(t, goLayout, targetTimezone)
}

// Utility function for minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
