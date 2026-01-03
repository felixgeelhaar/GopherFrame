package core

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/storage"
	arrowbackend "github.com/felixgeelhaar/GopherFrame/pkg/storage/arrow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDataFrame_WriteToStorage tests writing DataFrame to storage
func TestDataFrame_WriteToStorage(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Carol"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{95.5, 87.3, 92.1}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Create temporary directory for test
	tmpDir, err := os.MkdirTemp("", "gopherframe-storage-test-*")
	require.NoError(t, err)
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testFile := filepath.Join(tmpDir, "test.arrow")

	// Test: Write to Arrow IPC format
	backend := arrowbackend.NewBackend()
	ctx := context.Background()

	err = df.WriteToStorage(ctx, backend, testFile, storage.WriteOptions{})
	require.NoError(t, err)

	err = backend.Close()
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(testFile)
	assert.NoError(t, err, "Output file should exist")

	// Test: Read back and verify
	readBackend := arrowbackend.NewBackend()
	defer func() { _ = readBackend.Close() }()

	readDf, err := NewDataFrameFromStorage(ctx, readBackend, testFile, storage.ReadOptions{})
	require.NoError(t, err)
	defer readDf.Release()

	// Verify data matches
	assert.Equal(t, df.NumRows(), readDf.NumRows())
	assert.Equal(t, df.NumCols(), readDf.NumCols())

	// Verify specific values
	readIdSeries, err := readDf.Column("id")
	require.NoError(t, err)
	readIdCol := readIdSeries.Array().(*array.Int64)
	assert.Equal(t, int64(1), readIdCol.Value(0))
	assert.Equal(t, int64(2), readIdCol.Value(1))
	assert.Equal(t, int64(3), readIdCol.Value(2))
}

// TestDataFrame_WriteToStorageWithOptions tests writing with various options
func TestDataFrame_WriteToStorageWithOptions(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "product", Type: arrow.BinaryTypes.String},
			{Name: "quantity", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	productBuilder := array.NewStringBuilder(pool)
	productBuilder.AppendValues([]string{"Widget", "Gadget"}, nil)
	productArray := productBuilder.NewArray()
	defer productArray.Release()

	qtyBuilder := array.NewInt64Builder(pool)
	qtyBuilder.AppendValues([]int64{100, 200}, nil)
	qtyArray := qtyBuilder.NewArray()
	defer qtyArray.Release()

	record := array.NewRecord(schema, []arrow.Array{productArray, qtyArray}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "gopherframe-options-test-*")
	require.NoError(t, err)
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testFile := filepath.Join(tmpDir, "test.arrow")

	backend := arrowbackend.NewBackend()
	ctx := context.Background()

	// Test with overwrite option
	opts := storage.WriteOptions{
		Overwrite: true,
		BatchSize: 1000,
	}

	err = df.WriteToStorage(ctx, backend, testFile, opts)
	require.NoError(t, err)
	_ = backend.Close()

	// Verify file exists
	_, err = os.Stat(testFile)
	assert.NoError(t, err)
}

// TestNewDataFrameFromStorage_Error tests error handling in storage reading
func TestNewDataFrameFromStorage_Error(t *testing.T) {
	ctx := context.Background()

	// Test with nil backend
	_, err := NewDataFrameFromStorage(ctx, nil, "test.arrow", storage.ReadOptions{})
	assert.Error(t, err)

	// Test with non-existent file
	backend := arrowbackend.NewBackend()
	defer func() { _ = backend.Close() }()
	_, err = NewDataFrameFromStorage(ctx, backend, "/nonexistent/path/file.arrow", storage.ReadOptions{})
	assert.Error(t, err)
}

// TestSingleRecordReader tests the singleRecordReader implementation
func TestSingleRecordReader(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "x", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	defer record.Release()

	// Create reader
	reader := &singleRecordReader{
		record: record,
		schema: schema,
	}

	// Test Next() - should return true first time
	assert.True(t, reader.Next())

	// Test Record() - should return the record
	rec := reader.Record()
	assert.NotNil(t, rec)
	assert.Equal(t, int64(3), rec.NumRows())

	// Test Next() again - should return false (already consumed)
	assert.False(t, reader.Next())

	// Test Record() after consumed - should return nil
	rec = reader.Record()
	assert.Nil(t, rec)

	// Test Schema()
	sch := reader.Schema()
	assert.Equal(t, schema, sch)

	// Test Err() - should be nil
	assert.NoError(t, reader.Err())

	// Test Close() - should succeed
	err := reader.Close()
	assert.NoError(t, err)
}

// TestSingleRecordReader_WithError tests singleRecordReader with error
func TestSingleRecordReader_WithError(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "x", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	defer record.Release()

	// Create reader with error
	testErr := assert.AnError
	reader := &singleRecordReader{
		record: record,
		schema: schema,
		err:    testErr,
	}

	// Test Next() with error - should return false
	assert.False(t, reader.Next())

	// Test Err() - should return the error
	assert.Equal(t, testErr, reader.Err())
}

// TestDataFrame_WriteReadRoundTrip tests complete write-read cycle
func TestDataFrame_WriteReadRoundTrip(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with various data types
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "int_col", Type: arrow.PrimitiveTypes.Int64},
			{Name: "float_col", Type: arrow.PrimitiveTypes.Float64},
			{Name: "string_col", Type: arrow.BinaryTypes.String},
			{Name: "bool_col", Type: arrow.FixedWidthTypes.Boolean},
		},
		nil,
	)

	intBuilder := array.NewInt64Builder(pool)
	intBuilder.AppendValues([]int64{10, 20, 30}, nil)
	intArray := intBuilder.NewArray()
	defer intArray.Release()

	floatBuilder := array.NewFloat64Builder(pool)
	floatBuilder.AppendValues([]float64{1.5, 2.5, 3.5}, nil)
	floatArray := floatBuilder.NewArray()
	defer floatArray.Release()

	stringBuilder := array.NewStringBuilder(pool)
	stringBuilder.AppendValues([]string{"a", "b", "c"}, nil)
	stringArray := stringBuilder.NewArray()
	defer stringArray.Release()

	boolBuilder := array.NewBooleanBuilder(pool)
	boolBuilder.AppendValues([]bool{true, false, true}, nil)
	boolArray := boolBuilder.NewArray()
	defer boolArray.Release()

	record := array.NewRecord(schema, []arrow.Array{intArray, floatArray, stringArray, boolArray}, 3)
	defer record.Release()

	originalDf := NewDataFrame(record)
	defer originalDf.Release()

	// Write and read back
	tmpDir, err := os.MkdirTemp("", "gopherframe-roundtrip-test-*")
	require.NoError(t, err)
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testFile := filepath.Join(tmpDir, "roundtrip.arrow")
	ctx := context.Background()

	writeBackend := arrowbackend.NewBackend()
	err = originalDf.WriteToStorage(ctx, writeBackend, testFile, storage.WriteOptions{})
	require.NoError(t, err)
	_ = writeBackend.Close()

	readBackend := arrowbackend.NewBackend()
	defer func() { _ = readBackend.Close() }()

	readDf, err := NewDataFrameFromStorage(ctx, readBackend, testFile, storage.ReadOptions{})
	require.NoError(t, err)
	defer readDf.Release()

	// Verify all column types and values
	assert.Equal(t, originalDf.NumRows(), readDf.NumRows())
	assert.Equal(t, originalDf.NumCols(), readDf.NumCols())

	// Verify int column
	intSeries, err := readDf.Column("int_col")
	require.NoError(t, err)
	intCol := intSeries.Array().(*array.Int64)
	assert.Equal(t, int64(10), intCol.Value(0))
	assert.Equal(t, int64(20), intCol.Value(1))
	assert.Equal(t, int64(30), intCol.Value(2))

	// Verify float column
	floatSeries, err := readDf.Column("float_col")
	require.NoError(t, err)
	floatCol := floatSeries.Array().(*array.Float64)
	assert.Equal(t, 1.5, floatCol.Value(0))

	// Verify string column
	stringSeries, err := readDf.Column("string_col")
	require.NoError(t, err)
	stringCol := stringSeries.Array().(*array.String)
	assert.Equal(t, "a", stringCol.Value(0))

	// Verify bool column
	boolSeries, err := readDf.Column("bool_col")
	require.NoError(t, err)
	boolCol := boolSeries.Array().(*array.Boolean)
	assert.True(t, boolCol.Value(0))
	assert.False(t, boolCol.Value(1))
}

// TestDataFrame_WriteToStorage_NilBackend tests writing with nil backend
func TestDataFrame_WriteToStorage_NilBackend(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "x", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	ctx := context.Background()

	// Try to write with nil backend (should error)
	err := df.WriteToStorage(ctx, nil, "test.arrow", storage.WriteOptions{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no storage backend available")
}
