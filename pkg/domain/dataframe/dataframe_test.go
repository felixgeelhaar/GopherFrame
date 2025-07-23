package dataframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestNewDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	df := NewDataFrame(record)
	defer df.Release()

	if df.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", df.NumRows())
	}

	if df.NumCols() != 1 {
		t.Errorf("Expected 1 column, got %d", df.NumCols())
	}
}

func TestDataFrameColumnNames(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)
	idBuilder.AppendValues([]int64{1, 2}, nil)
	nameBuilder.AppendValues([]string{"Alice", "Bob"}, nil)

	idArr := idBuilder.NewArray()
	nameArr := nameBuilder.NewArray()
	defer idArr.Release()
	defer nameArr.Release()

	record := array.NewRecord(schema, []arrow.Array{idArr, nameArr}, 2)
	df := NewDataFrame(record)
	defer df.Release()

	names := df.ColumnNames()
	expected := []string{"id", "name"}

	if len(names) != len(expected) {
		t.Fatalf("Expected %d columns, got %d", len(expected), len(names))
	}

	for i, exp := range expected {
		if names[i] != exp {
			t.Errorf("Column %d: expected %s, got %s", i, exp, names[i])
		}
	}
}

func TestDataFrameHasColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.Append(42)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	df := NewDataFrame(record)
	defer df.Release()

	if !df.HasColumn("test") {
		t.Error("Expected HasColumn('test') to return true")
	}

	if df.HasColumn("nonexistent") {
		t.Error("Expected HasColumn('nonexistent') to return false")
	}
}

func TestDataFrame_Record(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.Append(42)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	df := NewDataFrame(record)
	defer df.Release()

	resultRecord := df.Record()
	if resultRecord == nil {
		t.Error("Record() should not return nil")
	}

	if resultRecord.NumRows() != 1 {
		t.Errorf("Expected 1 row, got %d", resultRecord.NumRows())
	}
}

func TestDataFrameClone(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.Append(42)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	df := NewDataFrame(record)
	defer df.Release()

	clone := df.Clone()
	defer clone.Release()

	if clone.NumRows() != df.NumRows() {
		t.Error("Clone should have same number of rows")
	}

	if clone.NumCols() != df.NumCols() {
		t.Error("Clone should have same number of columns")
	}

	if !clone.Schema().Equal(df.Schema()) {
		t.Error("Clone should have same schema")
	}
}

// Test Series functionality
func TestNewSeries(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	series := NewSeries("test_series", arr)
	defer series.Release()

	if series.Name() != "test_series" {
		t.Errorf("Expected name 'test_series', got '%s'", series.Name())
	}

	if series.Len() != 3 {
		t.Errorf("Expected length 3, got %d", series.Len())
	}

	if series.DataType().ID() != arrow.INT64 {
		t.Errorf("Expected INT64 type, got %s", series.DataType())
	}
}

func TestSeries_GetInt64(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{10, 20, 30}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	series := NewSeries("int_series", arr)
	defer series.Release()

	value, err := series.GetInt64(1)
	if err != nil {
		t.Errorf("GetInt64 failed: %v", err)
	}

	if value != 20 {
		t.Errorf("Expected value 20, got %d", value)
	}

	// Test out of bounds
	_, err = series.GetInt64(10)
	if err == nil {
		t.Error("Expected error for out of bounds access")
	}
}

func TestSeries_GetFloat64(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewFloat64Builder(pool)
	builder.AppendValues([]float64{1.5, 2.5, 3.5}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	series := NewSeries("float_series", arr)
	defer series.Release()

	value, err := series.GetFloat64(0)
	if err != nil {
		t.Errorf("GetFloat64 failed: %v", err)
	}

	if value != 1.5 {
		t.Errorf("Expected value 1.5, got %f", value)
	}

	// Test with wrong type
	intBuilder := array.NewInt64Builder(pool)
	intBuilder.AppendValues([]int64{1, 2, 3}, nil)
	intArr := intBuilder.NewArray()
	defer intArr.Release()

	intSeries := NewSeries("int_series", intArr)
	defer intSeries.Release()

	_, err = intSeries.GetFloat64(0)
	if err == nil {
		t.Error("Expected error when getting float64 from int64 series")
	}
}

func TestSeries_GetString(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	builder.AppendValues([]string{"hello", "world", "test"}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	series := NewSeries("string_series", arr)
	defer series.Release()

	value, err := series.GetString(2)
	if err != nil {
		t.Errorf("GetString failed: %v", err)
	}

	if value != "test" {
		t.Errorf("Expected value 'test', got '%s'", value)
	}
}

func TestSeries_IsNull(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)

	// Add values with nulls: [1, null, 3]
	builder.Append(1)
	builder.AppendNull()
	builder.Append(3)

	arr := builder.NewArray()
	defer arr.Release()

	series := NewSeries("nullable_series", arr)
	defer series.Release()

	if series.IsNull(0) {
		t.Error("Index 0 should not be null")
	}

	if !series.IsNull(1) {
		t.Error("Index 1 should be null")
	}

	if series.IsNull(2) {
		t.Error("Index 2 should not be null")
	}
}

func TestSeries_Array(t *testing.T) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	series := NewSeries("test_series", arr)
	defer series.Release()

	resultArray := series.Array()
	if resultArray == nil {
		t.Error("Array() should not return nil")
	}

	if resultArray.Len() != 3 {
		t.Errorf("Expected array length 3, got %d", resultArray.Len())
	}
}
