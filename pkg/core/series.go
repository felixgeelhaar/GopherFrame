// Package core provides Series implementation wrapping Apache Arrow Arrays.
package core

import (
	"fmt"
	"math"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

// Series represents an immutable, one-dimensional array of data.
// It wraps an Arrow Array to provide type-safe operations and
// seamless interoperability with the Arrow ecosystem.
//
// Thread Safety: Series is immutable and safe for concurrent reads.
// All transformation methods return new Series instances without modifying
// the original. However, Release() is not thread-safe and should only be called
// when you're certain no other goroutines are using the Series.
type Series struct {
	// array is the underlying Arrow Array containing the actual data.
	array arrow.Array

	// field contains the column metadata (name, type, nullable).
	field arrow.Field
}

// NewSeries creates a new Series from an Arrow Array and Field.
// The Series takes ownership of the array and will release it when closed.
func NewSeries(arr arrow.Array, field arrow.Field) *Series {
	arr.Retain() // Increment reference count
	return &Series{
		array: arr,
		field: field,
	}
}

// NewSeriesFromData creates a Series from raw data using Arrow builders.
// T must be a supported Arrow data type.
func NewSeriesFromData[T any](_ string, _ []T) (*Series, error) {
	// This is a simplified version. Production implementation would
	// handle all Arrow data types and use proper builders.
	return nil, fmt.Errorf("NewSeriesFromData not yet implemented")
}

// Name returns the name of the Series (column name).
func (s *Series) Name() string {
	return s.field.Name
}

// DataType returns the Arrow data type of the Series.
func (s *Series) DataType() arrow.DataType {
	return s.field.Type
}

// Field returns the Arrow field metadata.
func (s *Series) Field() arrow.Field {
	return s.field
}

// Array returns the underlying Arrow Array.
// This is used internally by operations that need direct Arrow access.
func (s *Series) Array() arrow.Array {
	return s.array
}

// Len returns the number of elements in the Series.
func (s *Series) Len() int {
	return s.array.Len()
}

// Null returns the number of null values in the Series.
func (s *Series) Null() int {
	return s.array.NullN()
}

// IsNull returns true if the value at index i is null.
func (s *Series) IsNull(i int) bool {
	return s.array.IsNull(i)
}

// IsValid returns true if the value at index i is not null.
func (s *Series) IsValid(i int) bool {
	return s.array.IsValid(i)
}

// Nullable returns true if the Series can contain null values.
func (s *Series) Nullable() bool {
	return s.field.Nullable
}

// GetValue returns the value at index i as an interface{}.
// Returns nil if the value is null.
func (s *Series) GetValue(i int) interface{} {
	if i < 0 || i >= s.Len() {
		return nil
	}

	if s.IsNull(i) {
		return nil
	}

	// Extract value based on array type
	switch arr := s.array.(type) {
	case *array.Int64:
		return arr.Value(i)
	case *array.Float64:
		return arr.Value(i)
	case *array.String:
		return arr.Value(i)
	case *array.Boolean:
		return arr.Value(i)
	default:
		// Fallback: convert to string representation
		return fmt.Sprintf("<%s>", s.array.DataType())
	}
}

// GetString returns the value at index i as a string.
// This works for string types and provides string representation for others.
func (s *Series) GetString(i int) (string, error) {
	if i < 0 || i >= s.Len() {
		return "", fmt.Errorf("index out of range: %d", i)
	}

	if s.IsNull(i) {
		return "", nil
	}

	// Convert value to string based on array type
	switch arr := s.array.(type) {
	case *array.String:
		return arr.Value(i), nil
	case *array.Int64:
		return fmt.Sprintf("%d", arr.Value(i)), nil
	case *array.Float64:
		return fmt.Sprintf("%g", arr.Value(i)), nil
	case *array.Boolean:
		return fmt.Sprintf("%t", arr.Value(i)), nil
	default:
		return fmt.Sprintf("<%s>", s.array.DataType()), nil
	}
}

// GetInt64 returns the value at index i as an int64.
// This works for integer types and attempts conversion for others.
// For uint64 values exceeding math.MaxInt64, returns an error to prevent silent overflow.
func (s *Series) GetInt64(i int) (int64, error) {
	if i < 0 || i >= s.Len() {
		return 0, fmt.Errorf("index out of range: %d", i)
	}

	if s.IsNull(i) {
		return 0, fmt.Errorf("null value at index %d", i)
	}

	switch arr := s.array.(type) {
	case *array.Int64:
		return arr.Value(i), nil
	case *array.Int32:
		return int64(arr.Value(i)), nil
	case *array.Int16:
		return int64(arr.Value(i)), nil
	case *array.Int8:
		return int64(arr.Value(i)), nil
	case *array.Uint64:
		val := arr.Value(i)
		if val > math.MaxInt64 {
			return 0, fmt.Errorf("uint64 value %d exceeds int64 range [0, %d]", val, int64(math.MaxInt64))
		}
		return int64(val), nil
	case *array.Uint32:
		return int64(arr.Value(i)), nil
	case *array.Uint16:
		return int64(arr.Value(i)), nil
	case *array.Uint8:
		return int64(arr.Value(i)), nil
	default:
		return 0, fmt.Errorf("cannot convert %s to int64", s.DataType())
	}
}

// GetFloat64 returns the value at index i as a float64.
func (s *Series) GetFloat64(i int) (float64, error) {
	if i < 0 || i >= s.Len() {
		return 0, fmt.Errorf("index out of range: %d", i)
	}

	if s.IsNull(i) {
		return 0, fmt.Errorf("null value at index %d", i)
	}

	switch arr := s.array.(type) {
	case *array.Float64:
		return arr.Value(i), nil
	case *array.Float32:
		return float64(arr.Value(i)), nil
	default:
		// Try to convert from integer types
		if intVal, err := s.GetInt64(i); err == nil {
			return float64(intVal), nil
		}
		return 0, fmt.Errorf("cannot convert %s to float64", s.DataType())
	}
}

// GetBool returns the value at index i as a boolean.
func (s *Series) GetBool(i int) (bool, error) {
	if i < 0 || i >= s.Len() {
		return false, fmt.Errorf("index out of range: %d", i)
	}

	if s.IsNull(i) {
		return false, fmt.Errorf("null value at index %d", i)
	}

	if arr, ok := s.array.(*array.Boolean); ok {
		return arr.Value(i), nil
	}

	return false, fmt.Errorf("cannot convert %s to bool", s.DataType())
}

// Equal compares two Series for equality.
// Returns true if they have the same type, length, and data.
func (s *Series) Equal(other *Series) bool {
	if s == other {
		return true
	}

	if other == nil {
		return false
	}

	// Check field equality (name and type)
	if !s.field.Equal(other.field) {
		return false
	}

	// Use Arrow's built-in array equality
	return array.Equal(s.array, other.array)
}

// Validate checks the Series for consistency and data integrity.
func (s *Series) Validate() error {
	if s.array == nil {
		return fmt.Errorf("Series has no underlying array")
	}

	// Validate that the field type matches the array type
	if !arrow.TypeEqual(s.field.Type, s.array.DataType()) {
		return fmt.Errorf("field type %s does not match array type %s",
			s.field.Type, s.array.DataType())
	}

	return nil
}

// String returns a string representation of the Series.
// This is primarily for debugging and should not be used for large Series.
func (s *Series) String() string {
	if s.array == nil {
		return fmt.Sprintf("Series{%s: <empty>}", s.Name())
	}

	return fmt.Sprintf("Series{%s: %s, len: %d, nulls: %d}",
		s.Name(), s.DataType(), s.Len(), s.Null())
}

// Clone creates a shallow copy of the Series.
// The underlying Arrow data is shared (copy-on-write semantics).
func (s *Series) Clone() *Series {
	s.array.Retain() // Increment reference count
	return &Series{
		array: s.array,
		field: s.field,
	}
}

// Slice returns a new Series containing a subset of the data.
// The underlying Arrow data is shared when possible.
func (s *Series) Slice(offset, length int64) (*Series, error) {
	if offset < 0 || length < 0 {
		return nil, fmt.Errorf("offset and length must be non-negative")
	}

	if offset+length > int64(s.Len()) {
		return nil, fmt.Errorf("slice bounds out of range")
	}

	sliced := array.NewSlice(s.array, offset, offset+length)
	return NewSeries(sliced, s.field), nil
}

// Head returns the first n elements as a new Series.
func (s *Series) Head(n int) (*Series, error) {
	if n < 0 {
		return nil, fmt.Errorf("n must be non-negative")
	}

	if n == 0 {
		return s.Slice(0, 0)
	}

	length := int64(n)
	if length > int64(s.Len()) {
		length = int64(s.Len())
	}

	return s.Slice(0, length)
}

// Tail returns the last n elements as a new Series.
func (s *Series) Tail(n int) (*Series, error) {
	if n < 0 {
		return nil, fmt.Errorf("n must be non-negative")
	}

	if n == 0 {
		return s.Slice(int64(s.Len()), 0)
	}

	length := int64(n)
	offset := int64(s.Len()) - length

	if offset < 0 {
		offset = 0
		length = int64(s.Len())
	}

	return s.Slice(offset, length)
}

// Release decrements the reference count of the underlying Arrow Array.
// The Series should not be used after calling Release().
// It is safe to call Release multiple times.
//
// Thread Safety: Release() is NOT thread-safe. Only call Release() when you're
// certain no other goroutines are accessing this Series.
func (s *Series) Release() {
	if s.array != nil {
		s.array.Release()
		s.array = nil
	}
}
