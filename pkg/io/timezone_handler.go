package io

import (
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// TimezoneHandler provides comprehensive timezone support for temporal data I/O
type TimezoneHandler struct {
	defaultTimezone *time.Location
	timezoneCache   map[string]*time.Location
}

// NewTimezoneHandler creates a new timezone handler with UTC as default
func NewTimezoneHandler() *TimezoneHandler {
	return &TimezoneHandler{
		defaultTimezone: time.UTC,
		timezoneCache:   make(map[string]*time.Location),
	}
}

// NewTimezoneHandlerWithDefault creates a timezone handler with a custom default timezone
func NewTimezoneHandlerWithDefault(defaultTz string) (*TimezoneHandler, error) {
	loc, err := time.LoadLocation(defaultTz)
	if err != nil {
		return nil, fmt.Errorf("invalid default timezone '%s': %w", defaultTz, err)
	}

	return &TimezoneHandler{
		defaultTimezone: loc,
		timezoneCache:   make(map[string]*time.Location),
	}, nil
}

// SetDefaultTimezone sets the default timezone for parsing timestamps without explicit timezone info
func (th *TimezoneHandler) SetDefaultTimezone(timezone string) error {
	loc, err := th.loadTimezone(timezone)
	if err != nil {
		return fmt.Errorf("failed to set default timezone: %w", err)
	}
	th.defaultTimezone = loc
	return nil
}

// GetDefaultTimezone returns the current default timezone
func (th *TimezoneHandler) GetDefaultTimezone() *time.Location {
	return th.defaultTimezone
}

// loadTimezone loads a timezone from the cache or system, with caching
func (th *TimezoneHandler) loadTimezone(timezone string) (*time.Location, error) {
	// Check cache first
	if loc, exists := th.timezoneCache[timezone]; exists {
		return loc, nil
	}

	// Load from system
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, fmt.Errorf("failed to load timezone '%s': %w", timezone, err)
	}

	// Cache for future use
	th.timezoneCache[timezone] = loc
	return loc, nil
}

// ParseTimestampWithTimezone parses a timestamp string with explicit timezone handling
func (th *TimezoneHandler) ParseTimestampWithTimezone(value, layout, timezone string) (time.Time, error) {
	// If timezone is specified, use it
	if timezone != "" {
		loc, err := th.loadTimezone(timezone)
		if err != nil {
			return time.Time{}, err
		}
		return time.ParseInLocation(layout, value, loc)
	}

	// Try parsing with timezone info in the string first
	if parsed, err := time.Parse(layout, value); err == nil {
		// If the layout includes timezone info, use the parsed timezone
		if strings.Contains(layout, "MST") || strings.Contains(layout, "Z07") || strings.Contains(layout, "-07") {
			return parsed, nil
		}
	}

	// Fall back to default timezone
	return time.ParseInLocation(layout, value, th.defaultTimezone)
}

// ConvertToTimezone converts a time to the specified timezone
func (th *TimezoneHandler) ConvertToTimezone(t time.Time, timezone string) (time.Time, error) {
	if timezone == "" {
		return t, nil
	}

	loc, err := th.loadTimezone(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}

// FormatTimestampWithTimezone formats a timestamp in the specified timezone
func (th *TimezoneHandler) FormatTimestampWithTimezone(t time.Time, layout, timezone string) (string, error) {
	if timezone == "" {
		return t.Format(layout), nil
	}

	convertedTime, err := th.ConvertToTimezone(t, timezone)
	if err != nil {
		return "", err
	}

	return convertedTime.Format(layout), nil
}

// TimezoneAwareTimestampOptions provides configuration for timezone-aware timestamp parsing
type TimezoneAwareTimestampOptions struct {
	Layout         string // Go time layout
	Timezone       string // Target timezone for interpretation
	OutputTimezone string // Timezone for output (empty means keep original)
	StrictTimezone bool   // Whether to require explicit timezone in input
}

// ParseTimestampColumnWithTimezone parses a column of timestamp strings with timezone support
func (th *TimezoneHandler) ParseTimestampColumnWithTimezone(
	values []string,
	options TimezoneAwareTimestampOptions,
	pool memory.Allocator,
	isNull func(string) bool,
) (arrow.Array, error) {
	timestampType := &arrow.TimestampType{
		Unit:     arrow.Millisecond,
		TimeZone: options.OutputTimezone,
	}

	builder := array.NewTimestampBuilder(pool, timestampType)
	defer builder.Release()

	builder.Reserve(len(values))

	for _, value := range values {
		if isNull(value) {
			builder.AppendNull()
			continue
		}

		trimmed := strings.TrimSpace(value)

		// Parse with timezone awareness
		parsed, err := th.ParseTimestampWithTimezone(trimmed, options.Layout, options.Timezone)
		if err != nil {
			if options.StrictTimezone {
				return nil, fmt.Errorf("failed to parse timestamp '%s' with timezone: %w", trimmed, err)
			}
			// Fall back to default parsing
			parsed, err = time.ParseInLocation(options.Layout, trimmed, th.defaultTimezone)
			if err != nil {
				return nil, fmt.Errorf("failed to parse timestamp '%s': %w", trimmed, err)
			}
		}

		// Convert to output timezone if specified
		if options.OutputTimezone != "" {
			parsed, err = th.ConvertToTimezone(parsed, options.OutputTimezone)
			if err != nil {
				return nil, fmt.Errorf("failed to convert timestamp to output timezone: %w", err)
			}
		}

		// Convert to milliseconds since Unix epoch
		ms := parsed.UnixMilli()
		builder.Append(arrow.Timestamp(ms))
	}

	return builder.NewArray(), nil
}

// FormatTimestampArrayWithTimezone formats timestamp array values with timezone conversion
func (th *TimezoneHandler) FormatTimestampArrayWithTimezone(
	arr *array.Timestamp,
	layout string,
	sourceTimezone string,
	targetTimezone string,
) ([]string, error) {
	results := make([]string, arr.Len())

	// Get timestamp type info
	timestampType := arr.DataType().(*arrow.TimestampType)

	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			results[i] = ""
			continue
		}

		// Convert Arrow timestamp to time.Time
		tsValue := arr.Value(i)
		var t time.Time

		switch timestampType.Unit {
		case arrow.Second:
			t = time.Unix(int64(tsValue), 0)
		case arrow.Millisecond:
			t = time.Unix(int64(tsValue)/1000, (int64(tsValue)%1000)*1000000)
		case arrow.Microsecond:
			t = time.Unix(int64(tsValue)/1000000, (int64(tsValue)%1000000)*1000)
		case arrow.Nanosecond:
			t = time.Unix(int64(tsValue)/1000000000, int64(tsValue)%1000000000)
		default:
			t = time.Unix(int64(tsValue)/1000, (int64(tsValue)%1000)*1000000) // Default to milliseconds
		}

		// Apply source timezone if timestamp type has timezone info
		if timestampType.TimeZone != "" {
			if sourceTimezone == "" {
				sourceTimezone = timestampType.TimeZone
			}
		}

		if sourceTimezone != "" {
			sourceLoc, err := th.loadTimezone(sourceTimezone)
			if err != nil {
				return nil, fmt.Errorf("invalid source timezone '%s': %w", sourceTimezone, err)
			}
			t = time.Unix(t.Unix(), t.UnixNano()%1000000000).In(sourceLoc)
		} else {
			t = t.UTC()
		}

		// Format with target timezone
		formatted, err := th.FormatTimestampWithTimezone(t, layout, targetTimezone)
		if err != nil {
			return nil, fmt.Errorf("failed to format timestamp at index %d: %w", i, err)
		}

		results[i] = formatted
	}

	return results, nil
}

// Common timezone constants for convenience
const (
	TimezoneUTC   = "UTC"
	TimezoneLocal = "Local"
	TimezoneEST   = "America/New_York"
	TimezonePST   = "America/Los_Angeles"
	TimezoneCST   = "America/Chicago"
	TimezoneMST   = "America/Denver"
	TimezoneGMT   = "Europe/London"
	TimezoneCET   = "Europe/Berlin"
	TimezoneJST   = "Asia/Tokyo"
	TimezoneIST   = "Asia/Kolkata"
	TimezoneAEST  = "Australia/Sydney"
)

// GetCommonTimezones returns a list of commonly used timezones
func GetCommonTimezones() []string {
	return []string{
		TimezoneUTC,
		TimezoneLocal,
		TimezoneEST,
		TimezonePST,
		TimezoneCST,
		TimezoneMST,
		TimezoneGMT,
		TimezoneCET,
		TimezoneJST,
		TimezoneIST,
		TimezoneAEST,
	}
}

// ValidateTimezone checks if a timezone string is valid
func ValidateTimezone(timezone string) error {
	if timezone == "" {
		return nil // Empty timezone is valid (means use default)
	}

	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone '%s': %w", timezone, err)
	}

	return nil
}

// GetTimezoneOffset returns the current offset for a timezone in seconds
func GetTimezoneOffset(timezone string) (int, error) {
	if timezone == "" {
		return 0, nil
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, fmt.Errorf("invalid timezone '%s': %w", timezone, err)
	}

	now := time.Now().In(loc)
	_, offset := now.Zone()
	return offset, nil
}

// ConvertTimestampArray converts all timestamps in an array to a different timezone
func (th *TimezoneHandler) ConvertTimestampArray(
	arr *array.Timestamp,
	targetTimezone string,
	pool memory.Allocator,
) (arrow.Array, error) {
	if targetTimezone == "" {
		// No conversion needed, return copy
		return arr, nil
	}

	targetLoc, err := th.loadTimezone(targetTimezone)
	if err != nil {
		return nil, err
	}

	// Create new timestamp type with target timezone
	timestampType := &arrow.TimestampType{
		Unit:     arr.DataType().(*arrow.TimestampType).Unit,
		TimeZone: targetTimezone,
	}

	builder := array.NewTimestampBuilder(pool, timestampType)
	defer builder.Release()

	builder.Reserve(arr.Len())

	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendNull()
			continue
		}

		// Get original timestamp
		tsValue := arr.Value(i)

		// Convert to time.Time (assuming milliseconds for simplicity)
		originalTime := time.Unix(int64(tsValue)/1000, (int64(tsValue)%1000)*1000000).UTC()

		// Convert to target timezone
		convertedTime := originalTime.In(targetLoc)

		// Convert back to timestamp value
		newTsValue := arrow.Timestamp(convertedTime.UnixMilli())
		builder.Append(newTsValue)
	}

	return builder.NewArray(), nil
}
