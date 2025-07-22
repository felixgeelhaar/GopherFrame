// Package dataframe contains the Series value object.
package dataframe

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

// Series is a value object representing a single column of data.
// It wraps an Arrow Array to provide type-safe access to column values.
type Series struct {
	name  string
	array arrow.Array
}

// NewSeries creates a new Series from an Arrow Array.
func NewSeries(name string, array arrow.Array) *Series {
	array.Retain()
	return &Series{
		name:  name,
		array: array,
	}
}

// Name returns the name of the series.
func (s *Series) Name() string {
	return s.name
}

// Len returns the length of the series.
func (s *Series) Len() int {
	return s.array.Len()
}

// DataType returns the Arrow data type of the series.
func (s *Series) DataType() arrow.DataType {
	return s.array.DataType()
}

// IsNull checks if the value at the given index is null.
func (s *Series) IsNull(i int) bool {
	if i < 0 || i >= s.array.Len() {
		return false // Out of bounds is treated as non-null
	}
	return s.array.IsNull(i)
}

// GetFloat64 returns the float64 value at the given index.
func (s *Series) GetFloat64(i int) (float64, error) {
	if i < 0 || i >= s.array.Len() {
		return 0, fmt.Errorf("index %d out of bounds for series of length %d", i, s.array.Len())
	}

	switch arr := s.array.(type) {
	case *array.Float64:
		return arr.Value(i), nil
	default:
		return 0, fmt.Errorf("series is not float64 type, got %s", s.array.DataType())
	}
}

// GetInt64 returns the int64 value at the given index.
func (s *Series) GetInt64(i int) (int64, error) {
	if i < 0 || i >= s.array.Len() {
		return 0, fmt.Errorf("index %d out of bounds for series of length %d", i, s.array.Len())
	}

	switch arr := s.array.(type) {
	case *array.Int64:
		return arr.Value(i), nil
	default:
		return 0, fmt.Errorf("series is not int64 type, got %s", s.array.DataType())
	}
}

// GetString returns the string value at the given index.
func (s *Series) GetString(i int) (string, error) {
	if i < 0 || i >= s.array.Len() {
		return "", fmt.Errorf("index %d out of bounds for series of length %d", i, s.array.Len())
	}

	switch arr := s.array.(type) {
	case *array.String:
		return arr.Value(i), nil
	default:
		return "", fmt.Errorf("series is not string type, got %s", s.array.DataType())
	}
}

// Array returns the underlying Arrow array (read-only access).
func (s *Series) Array() arrow.Array {
	return s.array
}

// Release releases the underlying Arrow array.
func (s *Series) Release() {
	if s.array != nil {
		s.array.Release()
	}
}
