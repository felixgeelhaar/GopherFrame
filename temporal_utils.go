package gopherframe

import (
	"fmt"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// DateRange generates a DataFrame with a single timestamp column containing
// dates from start to end (exclusive) at the given interval.
func DateRange(columnName string, start, end time.Time, interval time.Duration) (*DataFrame, error) {
	if end.Before(start) {
		return nil, fmt.Errorf("end time must be after start time")
	}
	if interval <= 0 {
		return nil, fmt.Errorf("interval must be positive")
	}

	pool := memory.NewGoAllocator()

	// Count rows
	var timestamps []arrow.Timestamp
	for t := start; t.Before(end); t = t.Add(interval) {
		timestamps = append(timestamps, arrow.Timestamp(t.UnixMicro()))
	}

	tsType := &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for _, ts := range timestamps {
		builder.Append(ts)
	}

	tsArr := builder.NewArray()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: columnName, Type: tsType},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{tsArr}, int64(len(timestamps)))

	return NewDataFrame(record), nil
}

// BusinessDaysBetween calculates the number of business days (Mon-Fri) between two dates.
func BusinessDaysBetween(start, end time.Time) int {
	if end.Before(start) {
		return -BusinessDaysBetween(end, start)
	}

	count := 0
	current := start
	for current.Before(end) {
		wd := current.Weekday()
		if wd != time.Saturday && wd != time.Sunday {
			count++
		}
		current = current.AddDate(0, 0, 1)
	}
	return count
}

// AddBusinessDays adds the specified number of business days to a date.
func AddBusinessDays(start time.Time, days int) time.Time {
	if days < 0 {
		return subtractBusinessDays(start, -days)
	}

	current := start
	added := 0
	for added < days {
		current = current.AddDate(0, 0, 1)
		wd := current.Weekday()
		if wd != time.Saturday && wd != time.Sunday {
			added++
		}
	}
	return current
}

func subtractBusinessDays(start time.Time, days int) time.Time {
	current := start
	subtracted := 0
	for subtracted < days {
		current = current.AddDate(0, 0, -1)
		wd := current.Weekday()
		if wd != time.Saturday && wd != time.Sunday {
			subtracted++
		}
	}
	return current
}
