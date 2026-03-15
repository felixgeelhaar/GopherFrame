package gopherframe

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateRange_Basic(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)

	df, err := DateRange("date", start, end, 24*time.Hour)
	require.NoError(t, err)
	assert.Equal(t, int64(7), df.NumRows()) // Jan 1-7
	assert.True(t, df.HasColumn("date"))
}

func TestDateRange_Hourly(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

	df, err := DateRange("ts", start, end, time.Hour)
	require.NoError(t, err)
	assert.Equal(t, int64(12), df.NumRows())
}

func TestDateRange_InvalidRange(t *testing.T) {
	start := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	_, err := DateRange("date", start, end, 24*time.Hour)
	assert.Error(t, err)
}

func TestDateRange_ZeroInterval(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)

	_, err := DateRange("date", start, end, 0)
	assert.Error(t, err)
}

func TestBusinessDaysBetween_Basic(t *testing.T) {
	// Mon Jan 1 to Fri Jan 5 = 4 business days
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)   // Friday

	assert.Equal(t, 4, BusinessDaysBetween(start, end))
}

func TestBusinessDaysBetween_AcrossWeekend(t *testing.T) {
	// Mon Jan 1 to Mon Jan 8 = 5 business days
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, 5, BusinessDaysBetween(start, end))
}

func TestBusinessDaysBetween_Reversed(t *testing.T) {
	start := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, -5, BusinessDaysBetween(start, end))
}

func TestBusinessDaysBetween_SameDay(t *testing.T) {
	d := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 0, BusinessDaysBetween(d, d))
}

func TestAddBusinessDays_Basic(t *testing.T) {
	// Monday + 5 business days = next Monday
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // Monday
	result := AddBusinessDays(start, 5)
	assert.Equal(t, time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC), result)
}

func TestAddBusinessDays_FromFriday(t *testing.T) {
	// Friday + 1 business day = Monday
	start := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC) // Friday
	result := AddBusinessDays(start, 1)
	assert.Equal(t, time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC), result)
}

func TestAddBusinessDays_Negative(t *testing.T) {
	// Monday - 1 business day = Friday
	start := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC) // Monday
	result := AddBusinessDays(start, -1)
	assert.Equal(t, time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC), result)
}
