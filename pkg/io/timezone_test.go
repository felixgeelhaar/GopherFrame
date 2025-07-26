package io

import (
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimezoneHandler_Basic(t *testing.T) {
	t.Run("NewTimezoneHandler", func(t *testing.T) {
		handler := NewTimezoneHandler()
		assert.NotNil(t, handler)
		assert.Equal(t, time.UTC, handler.GetDefaultTimezone())
	})

	t.Run("NewTimezoneHandlerWithDefault", func(t *testing.T) {
		handler, err := NewTimezoneHandlerWithDefault("America/New_York")
		require.NoError(t, err)
		assert.NotNil(t, handler)

		defaultTz := handler.GetDefaultTimezone()
		assert.Equal(t, "America/New_York", defaultTz.String())
	})

	t.Run("InvalidDefaultTimezone", func(t *testing.T) {
		_, err := NewTimezoneHandlerWithDefault("Invalid/Timezone")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid default timezone")
	})
}

func TestTimezoneHandler_SetDefaultTimezone(t *testing.T) {
	handler := NewTimezoneHandler()

	t.Run("ValidTimezone", func(t *testing.T) {
		err := handler.SetDefaultTimezone("Europe/London")
		assert.NoError(t, err)
		assert.Equal(t, "Europe/London", handler.GetDefaultTimezone().String())
	})

	t.Run("InvalidTimezone", func(t *testing.T) {
		err := handler.SetDefaultTimezone("Invalid/Timezone")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to set default timezone")
	})
}

func TestTimezoneHandler_ParseTimestampWithTimezone(t *testing.T) {
	handler := NewTimezoneHandler()

	testCases := []struct {
		name      string
		value     string
		layout    string
		timezone  string
		expectErr bool
		checkFunc func(t *testing.T, result time.Time)
	}{
		{
			name:     "UTC timestamp",
			value:    "2023-06-15 14:30:00",
			layout:   "2006-01-02 15:04:05",
			timezone: "UTC",
			checkFunc: func(t *testing.T, result time.Time) {
				expected := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
				assert.Equal(t, expected, result)
			},
		},
		{
			name:     "EST timestamp",
			value:    "2023-06-15 14:30:00",
			layout:   "2006-01-02 15:04:05",
			timezone: "America/New_York",
			checkFunc: func(t *testing.T, result time.Time) {
				est, _ := time.LoadLocation("America/New_York")
				expected := time.Date(2023, 6, 15, 14, 30, 0, 0, est)
				assert.Equal(t, expected, result)
			},
		},
		{
			name:     "Timestamp with Z suffix",
			value:    "2023-06-15T14:30:00Z",
			layout:   "2006-01-02T15:04:05Z",
			timezone: "",
			checkFunc: func(t *testing.T, result time.Time) {
				expected := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
				assert.Equal(t, expected, result)
			},
		},
		{
			name:      "Invalid timezone",
			value:     "2023-06-15 14:30:00",
			layout:    "2006-01-02 15:04:05",
			timezone:  "Invalid/Timezone",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := handler.ParseTimestampWithTimezone(tc.value, tc.layout, tc.timezone)

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tc.checkFunc != nil {
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestTimezoneHandler_ConvertToTimezone(t *testing.T) {
	handler := NewTimezoneHandler()

	// Create a test time in UTC
	utcTime := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)

	testCases := []struct {
		name         string
		timezone     string
		expectErr    bool
		expectedHour int
	}{
		{
			name:         "Convert to EST",
			timezone:     "America/New_York",
			expectedHour: 10, // UTC 14:30 = EST 10:30 (during DST)
		},
		{
			name:         "Convert to PST",
			timezone:     "America/Los_Angeles",
			expectedHour: 7, // UTC 14:30 = PST 07:30 (during DST)
		},
		{
			name:         "Convert to JST",
			timezone:     "Asia/Tokyo",
			expectedHour: 23, // UTC 14:30 = JST 23:30
		},
		{
			name:     "Empty timezone (no conversion)",
			timezone: "",
		},
		{
			name:      "Invalid timezone",
			timezone:  "Invalid/Timezone",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := handler.ConvertToTimezone(utcTime, tc.timezone)

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tc.timezone == "" {
					assert.Equal(t, utcTime, result)
				} else {
					assert.Equal(t, tc.expectedHour, result.Hour())
				}
			}
		})
	}
}

func TestTimezoneHandler_ParseTimestampColumnWithTimezone(t *testing.T) {
	pool := memory.NewGoAllocator()
	handler := NewTimezoneHandler()

	t.Run("BasicTimezoneAwareParsing", func(t *testing.T) {
		values := []string{
			"2023-06-15T14:30:00Z",
			"2023-12-31T23:59:59Z",
			"",
		}

		options := TimezoneAwareTimestampOptions{
			Layout:         "2006-01-02T15:04:05Z",
			Timezone:       "UTC",
			OutputTimezone: "America/New_York",
		}

		isNull := func(value string) bool {
			return strings.TrimSpace(value) == ""
		}

		result, err := handler.ParseTimestampColumnWithTimezone(values, options, pool, isNull)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.TIMESTAMP, result.DataType().ID())
		assert.Equal(t, 3, result.Len())

		tsArray := result.(*array.Timestamp)
		assert.False(t, tsArray.IsNull(0))
		assert.False(t, tsArray.IsNull(1))
		assert.True(t, tsArray.IsNull(2))

		// Verify timezone conversion happened
		timestampType := result.DataType().(*arrow.TimestampType)
		assert.Equal(t, "America/New_York", timestampType.TimeZone)
	})

	t.Run("MixedFormatsWithTimezone", func(t *testing.T) {
		values := []string{
			"2023-06-15 14:30:00",
			"2023-12-31 23:59:59",
		}

		options := TimezoneAwareTimestampOptions{
			Layout:   "2006-01-02 15:04:05",
			Timezone: "Europe/London",
		}

		isNull := func(value string) bool {
			return strings.TrimSpace(value) == ""
		}

		result, err := handler.ParseTimestampColumnWithTimezone(values, options, pool, isNull)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.TIMESTAMP, result.DataType().ID())
		assert.Equal(t, 2, result.Len())
	})
}

func TestTimezoneHandler_FormatTimestampArrayWithTimezone(t *testing.T) {
	pool := memory.NewGoAllocator()
	handler := NewTimezoneHandler()

	// Create test timestamp array
	timestampType := &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: "UTC"}
	builder := array.NewTimestampBuilder(pool, timestampType)
	defer builder.Release()

	// Add test timestamps
	testTime1 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	testTime2 := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	builder.Append(arrow.Timestamp(testTime1.UnixMilli()))
	builder.Append(arrow.Timestamp(testTime2.UnixMilli()))
	builder.AppendNull()

	tsArray := builder.NewArray().(*array.Timestamp)
	defer tsArray.Release()

	t.Run("FormatWithTimezoneConversion", func(t *testing.T) {
		results, err := handler.FormatTimestampArrayWithTimezone(
			tsArray,
			"2006-01-02 15:04:05 MST",
			"UTC",
			"America/New_York",
		)
		require.NoError(t, err)

		assert.Equal(t, 3, len(results))
		assert.Contains(t, results[0], "EDT") // Eastern Daylight Time
		assert.Contains(t, results[1], "EST") // Eastern Standard Time
		assert.Equal(t, "", results[2])       // Null value
	})

	t.Run("FormatWithoutTimezoneConversion", func(t *testing.T) {
		results, err := handler.FormatTimestampArrayWithTimezone(
			tsArray,
			"2006-01-02 15:04:05",
			"",
			"",
		)
		require.NoError(t, err)

		assert.Equal(t, 3, len(results))
		assert.Equal(t, "2023-06-15 14:30:00", results[0])
		assert.Equal(t, "2023-12-31 23:59:59", results[1])
		assert.Equal(t, "", results[2])
	})
}

func TestTimezoneHandler_ConvertTimestampArray(t *testing.T) {
	pool := memory.NewGoAllocator()
	handler := NewTimezoneHandler()

	// Create test timestamp array
	timestampType := &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: "UTC"}
	builder := array.NewTimestampBuilder(pool, timestampType)

	testTime := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	builder.Append(arrow.Timestamp(testTime.UnixMilli()))
	builder.AppendNull()

	originalArray := builder.NewArray().(*array.Timestamp)
	defer originalArray.Release()

	t.Run("ConvertToEST", func(t *testing.T) {
		convertedArray, err := handler.ConvertTimestampArray(originalArray, "America/New_York", pool)
		require.NoError(t, err)
		defer convertedArray.Release()

		convertedTsArray := convertedArray.(*array.Timestamp)

		// Verify timezone metadata
		convertedType := convertedArray.DataType().(*arrow.TimestampType)
		assert.Equal(t, "America/New_York", convertedType.TimeZone)

		// Verify null preservation
		assert.True(t, convertedTsArray.IsNull(1))
	})

	t.Run("NoConversionNeeded", func(t *testing.T) {
		result, err := handler.ConvertTimestampArray(originalArray, "", pool)
		require.NoError(t, err)

		// Should return the same array
		assert.Equal(t, originalArray, result)
	})
}

func TestTimezoneValidation(t *testing.T) {
	testCases := []struct {
		name      string
		timezone  string
		expectErr bool
	}{
		{"Valid UTC", "UTC", false},
		{"Valid EST", "America/New_York", false},
		{"Valid PST", "America/Los_Angeles", false},
		{"Valid GMT", "Europe/London", false},
		{"Empty timezone", "", false},
		{"Invalid timezone", "Invalid/Timezone", true},
		{"Malformed timezone", "Not/A/Real/Timezone", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateTimezone(tc.timezone)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTimezoneOffset(t *testing.T) {
	testCases := []struct {
		name     string
		timezone string
	}{
		{"UTC", "UTC"},
		{"EST", "America/New_York"},
		{"PST", "America/Los_Angeles"},
		{"JST", "Asia/Tokyo"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			offset, err := GetTimezoneOffset(tc.timezone)
			assert.NoError(t, err)

			// Offset should be reasonable (between -12 and +14 hours)
			assert.GreaterOrEqual(t, offset, -12*3600)
			assert.LessOrEqual(t, offset, 14*3600)
		})
	}

	t.Run("EmptyTimezone", func(t *testing.T) {
		offset, err := GetTimezoneOffset("")
		assert.NoError(t, err)
		assert.Equal(t, 0, offset)
	})

	t.Run("InvalidTimezone", func(t *testing.T) {
		_, err := GetTimezoneOffset("Invalid/Timezone")
		assert.Error(t, err)
	})
}

func TestGetCommonTimezones(t *testing.T) {
	timezones := GetCommonTimezones()

	assert.NotEmpty(t, timezones)
	assert.Contains(t, timezones, "UTC")
	assert.Contains(t, timezones, "America/New_York")
	assert.Contains(t, timezones, "America/Los_Angeles")
	assert.Contains(t, timezones, "Europe/London")

	// Verify all returned timezones are valid
	for _, tz := range timezones {
		if tz != "Local" { // "Local" is special and might not be loadable in some environments
			err := ValidateTimezone(tz)
			assert.NoError(t, err, "Invalid timezone in common list: %s", tz)
		}
	}
}

func TestDateParserTimezoneIntegration(t *testing.T) {
	t.Run("DateParserWithTimezone", func(t *testing.T) {
		parser, err := NewDateParserWithTimezone("America/New_York")
		require.NoError(t, err)

		defaultTz := parser.GetDefaultTimezone()
		assert.Equal(t, "America/New_York", defaultTz.String())
	})

	t.Run("DateParserSetTimezone", func(t *testing.T) {
		parser := NewDateParser()

		err := parser.SetDefaultTimezone("Europe/London")
		assert.NoError(t, err)

		defaultTz := parser.GetDefaultTimezone()
		assert.Equal(t, "Europe/London", defaultTz.String())
	})

	t.Run("DateParserInvalidTimezone", func(t *testing.T) {
		parser := NewDateParser()

		err := parser.SetDefaultTimezone("Invalid/Timezone")
		assert.Error(t, err)

		// Should keep original timezone
		defaultTz := parser.GetDefaultTimezone()
		assert.Equal(t, "UTC", defaultTz.String())
	})
}

func TestTimezoneAwareCSVOptions(t *testing.T) {
	t.Run("DefaultCSVReadOptionsWithTimezone", func(t *testing.T) {
		options := DefaultCSVReadOptions()

		assert.Equal(t, "UTC", options.DefaultTimezone)
		assert.NotEmpty(t, options.TimestampFormats)
		assert.Contains(t, options.TimestampFormats, "2006-01-02T15:04:05Z")
	})

	t.Run("DefaultCSVWriteOptionsWithTimezone", func(t *testing.T) {
		options := DefaultCSVWriteOptions()

		assert.Equal(t, "", options.OutputTimezone)
		assert.False(t, options.IncludeTimezone)
	})
}

func BenchmarkTimezoneConversion(b *testing.B) {
	handler := NewTimezoneHandler()
	testTime := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)

	b.Run("ConvertToTimezone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = handler.ConvertToTimezone(testTime, "America/New_York")
		}
	})

	b.Run("ParseTimestampWithTimezone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = handler.ParseTimestampWithTimezone("2023-06-15 14:30:00", "2006-01-02 15:04:05", "America/New_York")
		}
	})
}
