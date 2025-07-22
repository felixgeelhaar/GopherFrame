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
