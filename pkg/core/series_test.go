package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestSeriesCreation(t *testing.T) {
	// Create test data
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test_column", Type: arrow.PrimitiveTypes.Int64}

	// Create Series
	series := NewSeries(arr, field)
	defer series.Release()

	// Test basic properties
	if series.Name() != "test_column" {
		t.Errorf("Expected name 'test_column', got '%s'", series.Name())
	}

	if series.Len() != 5 {
		t.Errorf("Expected length 5, got %d", series.Len())
	}

	if series.DataType().ID() != arrow.INT64 {
		t.Errorf("Expected INT64 type, got %s", series.DataType())
	}
}

func TestSeriesValueAccess(t *testing.T) {
	// Create test Series with mixed data including nulls
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)

	// Append values: [1, null, 3, 4, null]
	builder.Append(1)
	builder.AppendNull()
	builder.Append(3)
	builder.Append(4)
	builder.AppendNull()

	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test_column", Type: arrow.PrimitiveTypes.Int64, Nullable: true}
	series := NewSeries(arr, field)
	defer series.Release()

	// Test value access
	value := series.GetValue(0)
	if value != int64(1) {
		t.Errorf("Expected value 1, got %v", value)
	}

	// Test null value
	nullValue := series.GetValue(1)
	if nullValue != nil {
		t.Errorf("Expected nil for null value, got %v", nullValue)
	}

	// Test IsNull and IsValid
	if series.IsNull(0) {
		t.Error("Expected index 0 to not be null")
	}

	if !series.IsNull(1) {
		t.Error("Expected index 1 to be null")
	}

	if !series.IsValid(0) {
		t.Error("Expected index 0 to be valid")
	}

	if series.IsValid(1) {
		t.Error("Expected index 1 to be invalid")
	}

	// Test null count
	if series.Null() != 2 {
		t.Errorf("Expected 2 null values, got %d", series.Null())
	}
}

func TestSeriesStringRepresentation(t *testing.T) {
	// Create test Series
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	builder.AppendValues([]string{"hello", "world", "test"}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "text", Type: arrow.BinaryTypes.String}
	series := NewSeries(arr, field)
	defer series.Release()

	// Test GetString
	str, err := series.GetString(0)
	if err != nil {
		t.Errorf("Failed to get string value: %v", err)
	}
	if str != "hello" {
		t.Errorf("Expected 'hello', got '%s'", str)
	}

	// Test String method
	seriesStr := series.String()
	if seriesStr == "" {
		t.Error("Expected non-empty string representation")
	}
}

func TestSeriesValidation(t *testing.T) {
	// Create test Series
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64}
	series := NewSeries(arr, field)
	defer series.Release()

	// Test validation
	if err := series.Validate(); err != nil {
		t.Errorf("Expected valid Series, got error: %v", err)
	}
}

func TestNewSeriesFromData(t *testing.T) {
	// This function is not yet implemented according to the source code
	series, err := NewSeriesFromData("test", []int64{1, 2, 3, 4, 5})
	if err == nil {
		t.Error("Expected error for unimplemented function")
	}
	if series != nil {
		t.Error("Expected nil series for unimplemented function")
	}
}

func TestSeries_Field(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64, Nullable: true}
	series := NewSeries(arr, field)
	defer series.Release()

	resultField := series.Field()
	if resultField.Name != "test" {
		t.Errorf("Expected field name 'test', got '%s'", resultField.Name)
	}

	if !resultField.Nullable {
		t.Error("Expected field to be nullable")
	}
}

func TestSeries_Nullable(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64, Nullable: true}
	series := NewSeries(arr, field)
	defer series.Release()

	if !series.Nullable() {
		t.Error("Expected series to be nullable")
	}
}

func TestSeries_GetInt64(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{10, 20, 30}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64}
	series := NewSeries(arr, field)
	defer series.Release()

	value, err := series.GetInt64(1)
	if err != nil {
		t.Errorf("Failed to get int64 value: %v", err)
	}

	if value != 20 {
		t.Errorf("Expected value 20, got %d", value)
	}

	// Test with string series (should fail)
	stringBuilder := array.NewStringBuilder(pool)
	stringBuilder.AppendValues([]string{"hello"}, nil)
	stringArr := stringBuilder.NewArray()
	defer stringArr.Release()

	stringField := arrow.Field{Name: "str", Type: arrow.BinaryTypes.String}
	stringSeries := NewSeries(stringArr, stringField)
	defer stringSeries.Release()

	_, err = stringSeries.GetInt64(0)
	if err == nil {
		t.Error("Expected error when getting int64 from string series")
	}
}

func TestSeries_GetFloat64(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewFloat64Builder(pool)
	builder.AppendValues([]float64{1.5, 2.5, 3.5}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Float64}
	series := NewSeries(arr, field)
	defer series.Release()

	value, err := series.GetFloat64(1)
	if err != nil {
		t.Errorf("Failed to get float64 value: %v", err)
	}

	if value != 2.5 {
		t.Errorf("Expected value 2.5, got %f", value)
	}
}

func TestSeries_GetBool(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	builder.AppendValues([]bool{true, false, true}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.FixedWidthTypes.Boolean}
	series := NewSeries(arr, field)
	defer series.Release()

	value, err := series.GetBool(0)
	if err != nil {
		t.Errorf("Failed to get boolean value: %v", err)
	}

	if !value {
		t.Error("Expected true value")
	}

	value2, err := series.GetBool(1)
	if err != nil {
		t.Errorf("Failed to get boolean value: %v", err)
	}

	if value2 {
		t.Error("Expected false value")
	}
}

func TestSeries_Equal(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64}
	series1 := NewSeries(arr, field)
	defer series1.Release()

	// Create identical series
	builder2 := array.NewInt64Builder(pool)
	builder2.AppendValues([]int64{1, 2, 3}, nil)
	arr2 := builder2.NewArray()
	defer arr2.Release()

	series2 := NewSeries(arr2, field)
	defer series2.Release()

	if !series1.Equal(series2) {
		t.Error("Expected series to be equal")
	}

	// Create different series
	builder3 := array.NewInt64Builder(pool)
	builder3.AppendValues([]int64{4, 5, 6}, nil)
	arr3 := builder3.NewArray()
	defer arr3.Release()

	series3 := NewSeries(arr3, field)
	defer series3.Release()

	if series1.Equal(series3) {
		t.Error("Expected series to be different")
	}
}

func TestSeries_Clone(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64}
	series := NewSeries(arr, field)
	defer series.Release()

	cloned := series.Clone()
	defer cloned.Release()

	if !series.Equal(cloned) {
		t.Error("Expected cloned series to be equal to original")
	}

	if series.Len() != cloned.Len() {
		t.Errorf("Expected cloned length %d, got %d", series.Len(), cloned.Len())
	}
}

func TestSeriesSlicing(t *testing.T) {
	// Create test Series
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	field := arrow.Field{Name: "test", Type: arrow.PrimitiveTypes.Int64}
	series := NewSeries(arr, field)
	defer series.Release()

	// Test slice
	sliced, err := series.Slice(1, 3)
	if err != nil {
		t.Errorf("Failed to slice series: %v", err)
	}
	defer sliced.Release()

	if sliced.Len() != 3 {
		t.Errorf("Expected sliced length 3, got %d", sliced.Len())
	}

	// Check values in slice [2, 3, 4]
	value := sliced.GetValue(0)
	if value != int64(2) {
		t.Errorf("Expected first value in slice to be 2, got %v", value)
	}

	// Test head
	head, err := series.Head(2)
	if err != nil {
		t.Errorf("Failed to get head: %v", err)
	}
	defer head.Release()

	if head.Len() != 2 {
		t.Errorf("Expected head length 2, got %d", head.Len())
	}

	// Test tail
	tail, err := series.Tail(2)
	if err != nil {
		t.Errorf("Failed to get tail: %v", err)
	}
	defer tail.Release()

	if tail.Len() != 2 {
		t.Errorf("Expected tail length 2, got %d", tail.Len())
	}

	// Check last value is 5
	lastValue := tail.GetValue(1)
	if lastValue != int64(5) {
		t.Errorf("Expected last value in tail to be 5, got %v", lastValue)
	}
}
