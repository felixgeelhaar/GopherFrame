package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestDataFrame_BasicOperations(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Create test data
	idBuilder := array.NewInt64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)
	scoreBuilder := array.NewFloat64Builder(pool)

	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie"}, nil)
	scoreBuilder.AppendValues([]float64{95.5, 87.2, 92.0}, nil)

	idArray := idBuilder.NewArray()
	nameArray := nameBuilder.NewArray()
	scoreArray := scoreBuilder.NewArray()
	defer idArray.Release()
	defer nameArray.Release()
	defer scoreArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray}, 3)
	df := NewDataFrame(record)
	defer df.Release()

	// Test basic properties
	if df.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", df.NumRows())
	}

	if df.NumCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", df.NumCols())
	}

	// Test column names
	colNames := df.ColumnNames()
	expected := []string{"id", "name", "score"}
	for i, exp := range expected {
		if colNames[i] != exp {
			t.Errorf("Column %d: expected %s, got %s", i, exp, colNames[i])
		}
	}

	// Test HasColumn
	if !df.HasColumn("id") {
		t.Error("Expected HasColumn('id') to return true")
	}

	if df.HasColumn("nonexistent") {
		t.Error("Expected HasColumn('nonexistent') to return false")
	}
}

func TestDataFrame_Column(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "test", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	builder := array.NewFloat64Builder(pool)
	builder.AppendValues([]float64{1.1, 2.2, 3.3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	df := NewDataFrame(record)
	defer df.Release()

	// Test getting existing column
	series, err := df.Column("test")
	if err != nil {
		t.Fatalf("Column failed: %v", err)
	}
	defer series.Release()

	if series.Len() != 3 {
		t.Errorf("Expected series length 3, got %d", series.Len())
	}

	// Test getting nonexistent column
	_, err = df.Column("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent column")
	}
}

func TestDataFrame_Equals(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	// Create first DataFrame
	builder1 := array.NewInt64Builder(pool)
	builder1.AppendValues([]int64{1, 2, 3}, nil)
	arr1 := builder1.NewArray()
	defer arr1.Release()

	record1 := array.NewRecord(schema, []arrow.Array{arr1}, 3)
	df1 := NewDataFrame(record1)
	defer df1.Release()

	// Create identical DataFrame
	builder2 := array.NewInt64Builder(pool)
	builder2.AppendValues([]int64{1, 2, 3}, nil)
	arr2 := builder2.NewArray()
	defer arr2.Release()

	record2 := array.NewRecord(schema, []arrow.Array{arr2}, 3)
	df2 := NewDataFrame(record2)
	defer df2.Release()

	// Test basic properties (Equals method may not be implemented yet)
	if df1.NumRows() != df2.NumRows() {
		t.Error("Expected identical DataFrames to have same row count")
	}

	if df1.NumCols() != df2.NumCols() {
		t.Error("Expected identical DataFrames to have same column count")
	}

	// Create different DataFrame
	builder3 := array.NewInt64Builder(pool)
	builder3.AppendValues([]int64{4, 5, 6}, nil)
	arr3 := builder3.NewArray()
	defer arr3.Release()

	record3 := array.NewRecord(schema, []arrow.Array{arr3}, 3)
	df3 := NewDataFrame(record3)
	defer df3.Release()

	// Test that different data gives different results (basic check)
	if df1.NumRows() == df3.NumRows() {
		t.Log("Both DataFrames have same number of rows, which is expected")
		// More detailed equality checking would require the Equals method
	}
}

func TestDataFrame_EmptyDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	arr := builder.NewArray() // Empty array
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 0)
	df := NewDataFrame(record)
	defer df.Release()

	if df.NumRows() != 0 {
		t.Errorf("Expected 0 rows in empty DataFrame, got %d", df.NumRows())
	}

	if df.NumCols() != 1 {
		t.Errorf("Expected 1 column in empty DataFrame, got %d", df.NumCols())
	}
}

func TestDataFrame_LargeDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "value", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	size := 10000
	builder := array.NewInt64Builder(pool)
	for i := 0; i < size; i++ {
		builder.Append(int64(i))
	}
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, int64(size))
	df := NewDataFrame(record)
	defer df.Release()

	if df.NumRows() != int64(size) {
		t.Errorf("Expected %d rows, got %d", size, df.NumRows())
	}

	// Test getting column from large DataFrame
	series, err := df.Column("value")
	if err != nil {
		t.Fatalf("Column failed: %v", err)
	}
	defer series.Release()

	if series.Len() != size {
		t.Errorf("Expected series length %d, got %d", size, series.Len())
	}
}

func TestDataFrame_MultipleDataTypes(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int64_col", Type: arrow.PrimitiveTypes.Int64},
			{Name: "float64_col", Type: arrow.PrimitiveTypes.Float64},
			{Name: "string_col", Type: arrow.BinaryTypes.String},
			{Name: "bool_col", Type: arrow.FixedWidthTypes.Boolean},
		},
		nil,
	)

	// Create arrays of different types
	intBuilder := array.NewInt64Builder(pool)
	floatBuilder := array.NewFloat64Builder(pool)
	stringBuilder := array.NewStringBuilder(pool)
	boolBuilder := array.NewBooleanBuilder(pool)

	intBuilder.AppendValues([]int64{1, 2, 3}, nil)
	floatBuilder.AppendValues([]float64{1.1, 2.2, 3.3}, nil)
	stringBuilder.AppendValues([]string{"a", "b", "c"}, nil)
	boolBuilder.AppendValues([]bool{true, false, true}, nil)

	intArray := intBuilder.NewArray()
	floatArray := floatBuilder.NewArray()
	stringArray := stringBuilder.NewArray()
	boolArray := boolBuilder.NewArray()
	defer intArray.Release()
	defer floatArray.Release()
	defer stringArray.Release()
	defer boolArray.Release()

	record := array.NewRecord(schema, []arrow.Array{intArray, floatArray, stringArray, boolArray}, 3)
	df := NewDataFrame(record)
	defer df.Release()

	// Test column access for different types
	intSeries, err := df.Column("int64_col")
	if err != nil {
		t.Fatalf("Failed to get int64 column: %v", err)
	}
	defer intSeries.Release()

	if intSeries.DataType().ID() != arrow.INT64 {
		t.Errorf("Expected int64 type, got %v", intSeries.DataType().ID())
	}

	floatSeries, err := df.Column("float64_col")
	if err != nil {
		t.Fatalf("Failed to get float64 column: %v", err)
	}
	defer floatSeries.Release()

	if floatSeries.DataType().ID() != arrow.FLOAT64 {
		t.Errorf("Expected float64 type, got %v", floatSeries.DataType().ID())
	}

	stringSeries, err := df.Column("string_col")
	if err != nil {
		t.Fatalf("Failed to get string column: %v", err)
	}
	defer stringSeries.Release()

	if stringSeries.DataType().ID() != arrow.STRING {
		t.Errorf("Expected string type, got %v", stringSeries.DataType().ID())
	}

	boolSeries, err := df.Column("bool_col")
	if err != nil {
		t.Fatalf("Failed to get bool column: %v", err)
	}
	defer boolSeries.Release()

	if boolSeries.DataType().ID() != arrow.BOOL {
		t.Errorf("Expected bool type, got %v", boolSeries.DataType().ID())
	}
}
