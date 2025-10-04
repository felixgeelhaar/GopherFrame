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
//
// The Series takes ownership of the array by incrementing its reference count.
// The array will be released when the Series's Release() method is called.
//
// Parameters:
//   - arr: Arrow Array containing the series data
//   - field: Arrow Field containing metadata (name, type, nullability)
//
// Returns:
//   - *Series: A new Series wrapping the provided array and field
//
// Memory Management:
//   - The array's reference count is incremented (Retain called)
//   - Caller must call Release() on the Series when done
//
// Example:
//
//	builder := array.NewInt64Builder(pool)
//	builder.AppendValues([]int64{1, 2, 3}, nil)
//	arr := builder.NewArray()
//	field := arrow.Field{Name: "numbers", Type: arrow.PrimitiveTypes.Int64}
//	series := NewSeries(arr, field)
//	defer series.Release()
func NewSeries(arr arrow.Array, field arrow.Field) *Series {
	arr.Retain() // Increment reference count
	return &Series{
		array: arr,
		field: field,
	}
}

// NewSeriesFromData creates a Series from raw Go data using Arrow builders.
//
// This is a generic constructor that will build an Arrow array from Go slices.
// The type parameter T must be a supported Arrow data type (int64, float64, string, bool).
//
// Parameters:
//   - name: Name for the series (column name)
//   - data: Slice of values to populate the series
//
// Returns:
//   - *Series: A new Series containing the data
//   - error: Returns error if data type not supported or builder fails
//
// Note: This is currently unimplemented and will be added in a future version.
//
// Planned Example:
//
//	series, err := NewSeriesFromData("ages", []int64{25, 30, 35, 40})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer series.Release()
func NewSeriesFromData[T any](_ string, _ []T) (*Series, error) {
	// This is a simplified version. Production implementation would
	// handle all Arrow data types and use proper builders.
	return nil, fmt.Errorf("NewSeriesFromData not yet implemented")
}

// Name returns the name of the Series (column name).
//
// This is an O(1) operation that returns the field name from metadata.
//
// Returns:
//   - string: The series name (column name)
//
// Example:
//
//	name := series.Name()
//	fmt.Printf("Series name: %s\n", name)
func (s *Series) Name() string {
	return s.field.Name
}

// DataType returns the Arrow data type of the Series.
//
// This is an O(1) operation that returns the type from field metadata.
//
// Returns:
//   - arrow.DataType: The Arrow data type (e.g., INT64, FLOAT64, STRING)
//
// Example:
//
//	dataType := series.DataType()
//	fmt.Printf("Data type: %s\n", dataType)
func (s *Series) DataType() arrow.DataType {
	return s.field.Type
}

// Field returns the Arrow field metadata.
//
// The field contains the series name, data type, and nullability information.
// This method is useful for schema introspection and metadata access.
//
// Returns:
//   - arrow.Field: The field metadata structure
//
// Example:
//
//	field := series.Field()
//	fmt.Printf("Field: %s (%s), nullable: %v\n",
//	    field.Name, field.Type, field.Nullable)
func (s *Series) Field() arrow.Field {
	return s.field
}

// Array returns the underlying Arrow Array.
//
// This method provides direct access to the internal Arrow array for advanced use cases
// or interoperability with Arrow libraries. Use with caution as it exposes internal state.
//
// Returns:
//   - arrow.Array: The underlying Arrow array
//
// Note: This is used internally by operations that need direct Arrow access.
func (s *Series) Array() arrow.Array {
	return s.array
}

// Len returns the number of elements in the Series.
//
// This is an O(1) operation that returns the array length, including null values.
//
// Returns:
//   - int: Number of elements in the series (including nulls)
//
// Example:
//
//	fmt.Printf("Series has %d elements\n", series.Len())
//
// See also: Null for null value count
func (s *Series) Len() int {
	return s.array.Len()
}

// Null returns the number of null values in the Series.
//
// This is an O(1) operation that returns the null count from array metadata.
//
// Returns:
//   - int: Number of null values in the series
//
// Example:
//
//	nulls := series.Null()
//	fmt.Printf("Series has %d null values\n", nulls)
//
// See also: Len for total element count, IsNull to check specific index
func (s *Series) Null() int {
	return s.array.NullN()
}

// IsNull returns true if the value at index i is null.
//
// Parameters:
//   - i: Zero-based index to check (0 to Len()-1)
//
// Returns:
//   - bool: True if value at index is null, false otherwise
//
// Note: Does not validate index bounds - behavior undefined for out-of-range index.
//
// Example:
//
//	if series.IsNull(5) {
//	    fmt.Println("Value at index 5 is null")
//	}
//
// See also: IsValid for the inverse check
func (s *Series) IsNull(i int) bool {
	return s.array.IsNull(i)
}

// IsValid returns true if the value at index i is not null.
//
// This is the inverse of IsNull, provided for convenience and readability.
//
// Parameters:
//   - i: Zero-based index to check (0 to Len()-1)
//
// Returns:
//   - bool: True if value at index is valid (not null), false if null
//
// Note: Does not validate index bounds - behavior undefined for out-of-range index.
//
// Example:
//
//	if series.IsValid(5) {
//	    val := series.GetValue(5)
//	    fmt.Printf("Value: %v\n", val)
//	}
//
// See also: IsNull for null checking
func (s *Series) IsValid(i int) bool {
	return s.array.IsValid(i)
}

// Nullable returns true if the Series can contain null values.
//
// This is determined by the field's nullable flag in the schema metadata.
//
// Returns:
//   - bool: True if series can contain nulls, false otherwise
//
// Example:
//
//	if series.Nullable() {
//	    fmt.Println("This series may contain null values")
//	}
func (s *Series) Nullable() bool {
	return s.field.Nullable
}

// GetValue returns the value at index i as an interface{}.
//
// This method returns the value in its native Go type wrapped in interface{}.
// Returns nil for out-of-bounds indices or null values. For unsupported types,
// returns a string representation.
//
// Parameters:
//   - i: Zero-based index (0 to Len()-1)
//
// Returns:
//   - interface{}: Value in native type (int64, float64, string, bool) or nil
//
// Supported return types:
//   - int64 for INT64 arrays
//   - float64 for FLOAT64 arrays
//   - string for STRING arrays
//   - bool for BOOLEAN arrays
//   - nil for null values or out-of-bounds
//   - string representation for unsupported types
//
// Example:
//
//	val := series.GetValue(0)
//	if val != nil {
//	    fmt.Printf("Value: %v (type: %T)\n", val, val)
//	}
//
// See also: GetInt64, GetFloat64, GetString, GetBool for type-specific access
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
//
// For string arrays, returns the actual string value. For other types, provides
// a string representation using fmt formatting. Returns empty string for null values.
//
// Parameters:
//   - i: Zero-based index (0 to Len()-1)
//
// Returns:
//   - string: String value or representation
//   - error: Returns error if index out of bounds
//
// Type conversion:
//   - STRING: Direct string value
//   - INT64: Formatted as "%d"
//   - FLOAT64: Formatted as "%g" (minimal representation)
//   - BOOL: Formatted as "true" or "false"
//   - Others: Type name in angle brackets "<type>"
//   - NULL: Empty string ""
//
// Example:
//
//	str, err := series.GetString(0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(str)
//
// See also: GetValue for native type, GetInt64/GetFloat64/GetBool for specific types
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
//
// This method works for all integer types (signed and unsigned) with automatic conversion.
// For uint64 values exceeding math.MaxInt64, returns an error to prevent silent overflow.
//
// Parameters:
//   - i: Zero-based index (0 to Len()-1)
//
// Returns:
//   - int64: Integer value
//   - error: Returns error if index out of bounds, value is null, or conversion fails
//
// Supported conversions:
//   - INT64, INT32, INT16, INT8: Direct or widening conversion
//   - UINT64: Checked conversion (errors if > MaxInt64)
//   - UINT32, UINT16, UINT8: Safe widening conversion
//   - Others: Error (cannot convert type to int64)
//
// Error conditions:
//   - Index out of bounds
//   - Null value at index
//   - Uint64 overflow (value > 9223372036854775807)
//   - Incompatible type
//
// Example:
//
//	val, err := series.GetInt64(0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Int value: %d\n", val)
//
// See also: GetFloat64 for numeric conversion, GetValue for native type
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
//
// This method works for floating-point types and attempts conversion from integer types.
// Integer values are converted to float64 with potential precision loss for large values.
//
// Parameters:
//   - i: Zero-based index (0 to Len()-1)
//
// Returns:
//   - float64: Floating-point value
//   - error: Returns error if index out of bounds, value is null, or conversion fails
//
// Supported conversions:
//   - FLOAT64: Direct value
//   - FLOAT32: Widening conversion to float64
//   - Integer types: Conversion via GetInt64 then to float64
//   - Others: Error (cannot convert type to float64)
//
// Error conditions:
//   - Index out of bounds
//   - Null value at index
//   - Incompatible type (non-numeric)
//
// Example:
//
//	val, err := series.GetFloat64(0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Float value: %f\n", val)
//
// See also: GetInt64 for integer access, GetValue for native type
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
//
// This method only works for boolean arrays. No automatic conversion from other types.
//
// Parameters:
//   - i: Zero-based index (0 to Len()-1)
//
// Returns:
//   - bool: Boolean value (true or false)
//   - error: Returns error if index out of bounds, value is null, or type mismatch
//
// Error conditions:
//   - Index out of bounds
//   - Null value at index
//   - Not a boolean array (no automatic conversion)
//
// Example:
//
//	val, err := series.GetBool(0)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Boolean value: %v\n", val)
//
// See also: GetValue for native type with automatic type detection
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
//
// Two Series are considered equal if they have:
//   - Same field (name and type)
//   - Same length
//   - Same data values in the same order (including null positions)
//
// Parameters:
//   - other: Series to compare with (can be nil)
//
// Returns:
//   - bool: True if Series are equal, false otherwise
//
// Note: This performs a deep comparison of all data values. For large Series,
// this operation can be expensive. Uses Arrow's built-in array equality.
//
// Example:
//
//	if series1.Equal(series2) {
//	    fmt.Println("Series are identical")
//	}
//
// See also: Validate for integrity checking
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
//
// This method verifies that:
//   - The Series has an underlying array
//   - The field type matches the array data type
//
// Returns:
//   - error: Nil if validation passes, error describing the issue otherwise
//
// Use this method to verify Series integrity after complex operations or
// when loading data from untrusted sources.
//
// Example:
//
//	if err := series.Validate(); err != nil {
//	    log.Fatalf("Series validation failed: %v", err)
//	}
//
// See also: Equal for comparing Series
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
//
// Implements the fmt.Stringer interface. Returns a summary showing name, type,
// length, and null count. This is primarily for debugging and should not be used
// for large Series as it doesn't truncate output.
//
// Returns:
//   - string: Human-readable representation of the Series
//
// Example Output:
//
//	"Series{ages: int64, len: 1000, nulls: 5}"
//
// Note: For actual data values, use GetValue() or type-specific getters.
func (s *Series) String() string {
	if s.array == nil {
		return fmt.Sprintf("Series{%s: <empty>}", s.Name())
	}

	return fmt.Sprintf("Series{%s: %s, len: %d, nulls: %d}",
		s.Name(), s.DataType(), s.Len(), s.Null())
}

// Clone creates a shallow copy of the Series.
//
// The returned Series shares the underlying Arrow data with the original
// (copy-on-write semantics). The reference count is incremented to prevent
// premature deallocation. Both Series must be released independently.
//
// Returns:
//   - *Series: A new Series sharing the same underlying data
//
// Memory: Caller must call Release() on the cloned Series when done
//
// Example:
//
//	series2 := series.Clone()
//	defer series2.Release()
//	// series2 shares data with series but is an independent reference
//
// See also: Slice for creating a subset view
func (s *Series) Clone() *Series {
	s.array.Retain() // Increment reference count
	return &Series{
		array: s.array,
		field: s.field,
	}
}

// Slice returns a new Series containing a subset of the data.
//
// Creates a zero-copy view into the Series data using Arrow's slice functionality.
// The underlying Arrow data is shared between the original and sliced Series.
//
// Parameters:
//   - offset: Starting index for the slice (must be >= 0)
//   - length: Number of elements to include (must be >= 0)
//
// Returns:
//   - *Series: New Series viewing the specified subset
//   - error: Returns error if offset/length negative or out of bounds
//
// Memory: Caller must call Release() on the returned Series
//
// Example:
//
//	// Get elements 10-19 (10 elements starting at index 10)
//	subset, err := series.Slice(10, 10)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer subset.Release()
//
// Complexity: O(1) - zero-copy slice operation
//
// See also: Head for first n elements, Tail for last n elements
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
//
// This is a convenience method that delegates to Slice(0, n). If n exceeds the
// Series length, returns all elements. Creates a zero-copy view.
//
// Parameters:
//   - n: Number of elements to return from the beginning (must be >= 0)
//
// Returns:
//   - *Series: New Series containing the first n elements (or all if n > Len())
//   - error: Returns error if n is negative
//
// Memory: Caller must call Release() on the returned Series
//
// Example:
//
//	// Get first 10 elements
//	top10, err := series.Head(10)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer top10.Release()
//
// Complexity: O(1) - zero-copy operation
//
// See also: Tail for last n elements, Slice for arbitrary range
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
//
// This method calculates the appropriate offset and delegates to Slice(). If n exceeds
// the Series length, returns all elements. Creates a zero-copy view.
//
// Parameters:
//   - n: Number of elements to return from the end (must be >= 0)
//
// Returns:
//   - *Series: New Series containing the last n elements (or all if n > Len())
//   - error: Returns error if n is negative
//
// Memory: Caller must call Release() on the returned Series
//
// Example:
//
//	// Get last 10 elements
//	bottom10, err := series.Tail(10)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer bottom10.Release()
//
// Complexity: O(1) - zero-copy operation
//
// See also: Head for first n elements, Slice for arbitrary range
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
//
// This method must be called when you're done with the Series to prevent memory leaks.
// After calling Release(), the Series should not be used again. It is safe to call
// Release() multiple times - subsequent calls after the first will have no effect.
//
// Memory Management:
//   - Decrements the Arrow Array's reference count
//   - When reference count reaches zero, the underlying memory is freed
//   - Sets internal array to nil to prevent use-after-free
//
// Thread Safety: Release() is NOT thread-safe. Only call Release() when you're
// certain no other goroutines are accessing this Series. For concurrent usage,
// ensure proper synchronization or use Clone() to create independent references.
//
// Example:
//
//	series := NewSeries(arr, field)
//	defer series.Release()  // Ensure cleanup even on error
//	// ... use series ...
//
// Common Pattern with Multiple Operations:
//
//	series1, _ := df.Column("age")
//	defer series1.Release()
//
//	series2, _ := series1.Head(10)
//	defer series2.Release()  // Each Series needs its own Release()
//
// See also: Clone for creating independent references with separate lifecycle
func (s *Series) Release() {
	if s.array != nil {
		s.array.Release()
		s.array = nil
	}
}
