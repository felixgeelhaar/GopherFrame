package io

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateParser_InferDateType(t *testing.T) {
	parser := NewDateParser()

	tests := []struct {
		name           string
		values         []string
		expectedIsDate bool
		expectedFormat string
		minConfidence  float64
	}{
		{
			name:           "ISO8601 dates",
			values:         []string{"2023-06-15", "2023-12-31", "2024-02-29"},
			expectedIsDate: true,
			expectedFormat: "YYYY-MM-DD",
			minConfidence:  1.0,
		},
		{
			name:           "US format dates",
			values:         []string{"06/15/2023", "12/31/2023", "02/29/2024"},
			expectedIsDate: true,
			expectedFormat: "MM/DD/YYYY",
			minConfidence:  1.0,
		},
		{
			name:           "European format dates",
			values:         []string{"15/06/2023", "31/12/2023", "29/02/2024"},
			expectedIsDate: true,
			expectedFormat: "DD/MM/YYYY",
			minConfidence:  1.0,
		},
		{
			name:           "Mixed valid and empty",
			values:         []string{"2023-06-15", "", "2023-12-31", "   ", "2024-02-29"},
			expectedIsDate: true,
			expectedFormat: "YYYY-MM-DD",
			minConfidence:  1.0,
		},
		{
			name:           "Datetime strings",
			values:         []string{"2023-06-15 14:30:00", "2023-12-31 23:59:59"},
			expectedIsDate: true,
			expectedFormat: "YYYY-MM-DD HH:MM:SS",
			minConfidence:  1.0,
		},
		{
			name:           "Non-date strings",
			values:         []string{"hello", "world", "not-a-date"},
			expectedIsDate: false,
		},
		{
			name:           "Numbers that look like dates",
			values:         []string{"20230615", "20231231", "20240229"},
			expectedIsDate: true,
			expectedFormat: "YYYYMMDD",
			minConfidence:  1.0,
		},
		{
			name:           "Partial matches below threshold",
			values:         []string{"2023-06-15", "not-a-date", "hello"},
			expectedIsDate: false, // Only 1/3 = 33% < 80% threshold
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.InferDateType(tt.values)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedIsDate, result.IsDate, "IsDate mismatch")
			if tt.expectedIsDate {
				assert.Equal(t, tt.expectedFormat, result.Format, "Format mismatch")
				assert.GreaterOrEqual(t, result.Confidence, tt.minConfidence, "Confidence too low")
				assert.NotNil(t, result.ArrowType, "ArrowType should be set")
				assert.False(t, result.SampleValue.IsZero(), "SampleValue should be set")
			}
		})
	}
}

func TestDateParser_ParseDateColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	parser := NewDateParser()

	t.Run("ValidISO8601Dates", func(t *testing.T) {
		values := []string{"2023-06-15", "2023-12-31", "2024-02-29"}

		result, err := parser.ParseDateColumn(values, "2006-01-02", pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.DATE32, result.DataType().ID())
		assert.Equal(t, 3, result.Len())

		dateArray := result.(*array.Date32)

		// Verify dates are correct
		expectedDates := []time.Time{
			time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		}

		for i, expected := range expectedDates {
			assert.False(t, dateArray.IsNull(i))
			days := dateArray.Value(i)
			actual := time.Unix(int64(days)*86400, 0).UTC()
			assert.Equal(t, expected.Year(), actual.Year())
			assert.Equal(t, expected.Month(), actual.Month())
			assert.Equal(t, expected.Day(), actual.Day())
		}
	})

	t.Run("WithNullValues", func(t *testing.T) {
		values := []string{"2023-06-15", "", "2023-12-31"}

		result, err := parser.ParseDateColumn(values, "2006-01-02", pool)
		require.NoError(t, err)
		defer result.Release()

		dateArray := result.(*array.Date32)
		assert.False(t, dateArray.IsNull(0))
		assert.True(t, dateArray.IsNull(1))
		assert.False(t, dateArray.IsNull(2))
	})

	t.Run("InvalidFormat", func(t *testing.T) {
		values := []string{"2023-06-15", "invalid-date"}

		_, err := parser.ParseDateColumn(values, "2006-01-02", pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse date value")
	})
}

func TestDateParser_ParseTimestampColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	parser := NewDateParser()

	t.Run("ValidTimestamps", func(t *testing.T) {
		values := []string{"2023-06-15 14:30:00", "2023-12-31 23:59:59"}

		result, err := parser.ParseTimestampColumn(values, "2006-01-02 15:04:05", pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.TIMESTAMP, result.DataType().ID())
		assert.Equal(t, 2, result.Len())

		timestampArray := result.(*array.Timestamp)

		// Verify timestamps are correct
		expectedTimes := []time.Time{
			time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
		}

		for i, expected := range expectedTimes {
			assert.False(t, timestampArray.IsNull(i))
			ms := int64(timestampArray.Value(i))
			actual := time.Unix(ms/1000, (ms%1000)*1000000).UTC()
			assert.Equal(t, expected.Year(), actual.Year())
			assert.Equal(t, expected.Month(), actual.Month())
			assert.Equal(t, expected.Day(), actual.Day())
			assert.Equal(t, expected.Hour(), actual.Hour())
			assert.Equal(t, expected.Minute(), actual.Minute())
			assert.Equal(t, expected.Second(), actual.Second())
		}
	})
}

func TestDateParser_FormatDateValue(t *testing.T) {
	parser := NewDateParser()

	// Test with known date: 2023-06-15 = 19523 days since epoch
	days := int32(19523)

	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{"ISO8601", "2006-01-02", "2023-06-15"},
		{"US", "01/02/2006", "06/15/2023"},
		{"European", "02/01/2006", "15/06/2023"},
		{"Long", "Jan 02, 2006", "Jun 15, 2023"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.FormatDateValue(days, tt.format)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDateParser_FormatTimestampValue(t *testing.T) {
	parser := NewDateParser()

	// Test with known timestamp: 2023-06-15 14:30:00 UTC
	testTime := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)

	tests := []struct {
		name      string
		timestamp arrow.Timestamp
		unit      arrow.TimeUnit
		format    string
		expected  string
	}{
		{
			name:      "Milliseconds",
			timestamp: arrow.Timestamp(testTime.UnixMilli()),
			unit:      arrow.Millisecond,
			format:    "2006-01-02 15:04:05",
			expected:  "2023-06-15 14:30:00",
		},
		{
			name:      "Seconds",
			timestamp: arrow.Timestamp(testTime.Unix()),
			unit:      arrow.Second,
			format:    "2006-01-02 15:04:05",
			expected:  "2023-06-15 14:30:00",
		},
		{
			name:      "ISO8601Format",
			timestamp: arrow.Timestamp(testTime.UnixMilli()),
			unit:      arrow.Millisecond,
			format:    "2006-01-02T15:04:05Z",
			expected:  "2023-06-15T14:30:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.FormatTimestampValue(tt.timestamp, tt.unit, tt.format)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDateParser_AddCustomFormats(t *testing.T) {
	parser := NewDateParser()

	// Add custom format
	parser.AddDateFormat("CUSTOM", "2006.01.02", "Dot-separated format")

	// Test inference with custom format
	values := []string{"2023.06.15", "2023.12.31", "2024.02.29"}
	result, err := parser.InferDateType(values)
	require.NoError(t, err)

	assert.True(t, result.IsDate)
	assert.Equal(t, "CUSTOM", result.Format)
	assert.Equal(t, "2006.01.02", result.GoLayout)
}

func TestDateParser_StrictMode(t *testing.T) {
	strictParser := NewStrictDateParser()

	// Mix of valid and invalid dates - should fail in strict mode
	values := []string{"2023-06-15", "invalid-date", "2023-12-31"}

	result, err := strictParser.InferDateType(values)
	require.NoError(t, err)

	assert.False(t, result.IsDate, "Strict mode should reject mixed valid/invalid dates")
}

func TestDateParser_ParseWithCustomFormat(t *testing.T) {
	pool := memory.NewGoAllocator()
	parser := NewDateParser()

	t.Run("CustomDateFormat", func(t *testing.T) {
		values := []string{"2023.06.15", "2023.12.31"}

		result, err := parser.ParseWithCustomFormat(values, "2006.01.02", pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.DATE32, result.DataType().ID())
		assert.Equal(t, 2, result.Len())
	})

	t.Run("CustomTimestampFormat", func(t *testing.T) {
		values := []string{"2023.06.15 14:30:00", "2023.12.31 23:59:59"}

		result, err := parser.ParseWithCustomFormat(values, "2006.01.02 15:04:05", pool)
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, arrow.TIMESTAMP, result.DataType().ID())
		assert.Equal(t, 2, result.Len())
	})

	t.Run("InvalidCustomFormat", func(t *testing.T) {
		values := []string{"2023-06-15"}

		_, err := parser.ParseWithCustomFormat(values, "invalid-layout", pool)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid date layout")
	})
}

func TestDateParser_GetSupportedFormats(t *testing.T) {
	parser := NewDateParser()

	dateFormats, dateTimeFormats := parser.GetSupportedFormats()

	assert.NotEmpty(t, dateFormats, "Should have date formats")
	assert.NotEmpty(t, dateTimeFormats, "Should have datetime formats")

	// Check some expected formats are present
	hasISO8601 := false
	for _, format := range dateFormats {
		if format.Pattern == "YYYY-MM-DD" {
			hasISO8601 = true
			break
		}
	}
	assert.True(t, hasISO8601, "Should have ISO 8601 format")
}

func TestDateParser_AutoDetectDateFormat(t *testing.T) {
	parser := NewDateParser()

	t.Run("SuccessfulDetection", func(t *testing.T) {
		values := []string{"2023-06-15", "2023-12-31", "2024-02-29"}

		result, err := parser.AutoDetectDateFormat(values)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.True(t, result.IsDate)
		assert.Equal(t, "YYYY-MM-DD", result.Format)
		assert.GreaterOrEqual(t, result.Confidence, 0.8)
	})

	t.Run("NoDateFormatFound", func(t *testing.T) {
		values := []string{"hello", "world", "not-a-date"}

		_, err := parser.AutoDetectDateFormat(values)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no recognizable date format found")
	})
}

func TestDateParser_ContainsTimeComponents(t *testing.T) {
	parser := NewDateParser()

	tests := []struct {
		layout   string
		expected bool
	}{
		{"2006-01-02", false},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02T15:04:05Z", true},
		{"Jan 02, 2006", false},
		{"2006-01-02 15:04", true},
		{"15:04:05", true},
	}

	for _, tt := range tests {
		t.Run(tt.layout, func(t *testing.T) {
			result := parser.containsTimeComponents(tt.layout)
			assert.Equal(t, tt.expected, result)
		})
	}
}
