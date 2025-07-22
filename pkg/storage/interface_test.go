package storage

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// Simple in-memory backend implementation for testing
type MemoryBackend struct {
	records []arrow.Record
	schema  *arrow.Schema
}

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{
		records: make([]arrow.Record, 0),
	}
}

func (m *MemoryBackend) AddRecord(record arrow.Record) error {
	if m.schema == nil {
		m.schema = record.Schema()
	}
	record.Retain()
	m.records = append(m.records, record)
	return nil
}

func (m *MemoryBackend) GetRecord(index int) (arrow.Record, error) {
	if index < 0 || index >= len(m.records) {
		return nil, ErrSourceNotFound
	}
	return m.records[index], nil
}

func (m *MemoryBackend) NumRecords() int {
	return len(m.records)
}

func (m *MemoryBackend) Schema() *arrow.Schema {
	return m.schema
}

func (m *MemoryBackend) Read(ctx context.Context, source string, opts ReadOptions) (RecordReader, error) {
	return nil, ErrUnsupportedOperation
}

func (m *MemoryBackend) Write(ctx context.Context, destination string, records RecordReader, opts WriteOptions) error {
	return ErrUnsupportedOperation
}

func (m *MemoryBackend) Scan(ctx context.Context, pattern string) ([]SourceInfo, error) {
	return nil, ErrUnsupportedOperation
}

func (m *MemoryBackend) Close() error {
	for _, record := range m.records {
		record.Release()
	}
	m.records = nil
	return nil
}

type SimpleRecordReader struct {
	backend *MemoryBackend
	index   int
	err     error
}

func NewRecordReader(backend *MemoryBackend) *SimpleRecordReader {
	return &SimpleRecordReader{
		backend: backend,
		index:   -1,
	}
}

func (r *SimpleRecordReader) Next() bool {
	r.index++
	return r.index < r.backend.NumRecords()
}

func (r *SimpleRecordReader) Record() arrow.Record {
	if r.index < 0 || r.index >= r.backend.NumRecords() {
		return nil
	}
	record, err := r.backend.GetRecord(r.index)
	if err != nil {
		r.err = err
		return nil
	}
	return record
}

func (r *SimpleRecordReader) Schema() *arrow.Schema {
	return r.backend.Schema()
}

func (r *SimpleRecordReader) Err() error {
	return r.err
}

func (r *SimpleRecordReader) Close() error {
	return nil
}

func (r *SimpleRecordReader) Release() {
	// Nothing to release for our simple implementation
}

func TestMemoryBackend(t *testing.T) {
	backend := NewMemoryBackend()

	// Test initial state
	if backend.NumRecords() != 0 {
		t.Errorf("Expected 0 records, got %d", backend.NumRecords())
	}

	// Create test record
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	pool := memory.NewGoAllocator()
	idBuilder := array.NewInt64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)

	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie"}, nil)

	idArray := idBuilder.NewArray()
	nameArray := nameBuilder.NewArray()
	defer idArray.Release()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 3)

	// Test AddRecord
	err := backend.AddRecord(record)
	if err != nil {
		t.Fatalf("AddRecord failed: %v", err)
	}

	if backend.NumRecords() != 1 {
		t.Errorf("Expected 1 record, got %d", backend.NumRecords())
	}

	// Test GetRecord
	retrievedRecord, err := backend.GetRecord(0)
	if err != nil {
		t.Fatalf("GetRecord failed: %v", err)
	}

	if retrievedRecord.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", retrievedRecord.NumRows())
	}

	if retrievedRecord.NumCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", retrievedRecord.NumCols())
	}

	// Test GetRecord with invalid index
	_, err = backend.GetRecord(1)
	if err == nil {
		t.Error("Expected error for invalid index")
	}

	// Test GetRecord with negative index
	_, err = backend.GetRecord(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
}

func TestMemoryBackend_MultipleRecords(t *testing.T) {
	backend := NewMemoryBackend()

	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "value", Type: arrow.PrimitiveTypes.Float64}},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Add multiple records
	for i := 0; i < 5; i++ {
		builder := array.NewFloat64Builder(pool)
		builder.AppendValues([]float64{float64(i), float64(i + 1)}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 2)
		err := backend.AddRecord(record)
		if err != nil {
			t.Fatalf("AddRecord %d failed: %v", i, err)
		}
	}

	if backend.NumRecords() != 5 {
		t.Errorf("Expected 5 records, got %d", backend.NumRecords())
	}

	// Test retrieving all records
	for i := 0; i < 5; i++ {
		record, err := backend.GetRecord(i)
		if err != nil {
			t.Fatalf("GetRecord %d failed: %v", i, err)
		}

		if record.NumRows() != 2 {
			t.Errorf("Record %d: expected 2 rows, got %d", i, record.NumRows())
		}
	}
}

func TestMemoryBackend_Schema(t *testing.T) {
	backend := NewMemoryBackend()

	// Test empty backend schema
	schema := backend.Schema()
	if schema != nil {
		t.Error("Expected nil schema for empty backend")
	}

	// Add record and test schema
	testSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "test", Type: arrow.PrimitiveTypes.Int32},
		},
		nil,
	)

	pool := memory.NewGoAllocator()
	builder := array.NewInt32Builder(pool)
	builder.Append(42)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(testSchema, []arrow.Array{arr}, 1)
	err := backend.AddRecord(record)
	if err != nil {
		t.Fatalf("AddRecord failed: %v", err)
	}

	schema = backend.Schema()
	if schema == nil {
		t.Fatal("Expected non-nil schema after adding record")
	}

	if !schema.Equal(testSchema) {
		t.Error("Schema doesn't match expected schema")
	}
}

func TestNewRecordReader(t *testing.T) {
	// Create backend with test data
	backend := NewMemoryBackend()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "id", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	err := backend.AddRecord(record)
	if err != nil {
		t.Fatalf("AddRecord failed: %v", err)
	}

	// Test creating reader
	reader := NewRecordReader(backend)
	if reader == nil {
		t.Fatal("NewRecordReader returned nil")
	}

	// Test Schema method
	readerSchema := reader.Schema()
	if !readerSchema.Equal(schema) {
		t.Error("Reader schema doesn't match backend schema")
	}

	// Test Next method
	hasNext := reader.Next()
	if !hasNext {
		t.Error("Expected reader to have next record")
	}

	readRecord := reader.Record()
	if readRecord.NumRows() != 3 {
		t.Errorf("Expected 3 rows, got %d", readRecord.NumRows())
	}

	// Test no more records
	hasNext = reader.Next()
	if hasNext {
		t.Error("Expected no more records")
	}

	// Test Err method
	if reader.Err() != nil {
		t.Errorf("Unexpected error: %v", reader.Err())
	}

	// Test Release
	reader.Release()
}

func TestRecordReader_Empty(t *testing.T) {
	backend := NewMemoryBackend()
	reader := NewRecordReader(backend)

	// Test empty reader
	hasNext := reader.Next()
	if hasNext {
		t.Error("Expected no records in empty reader")
	}

	record := reader.Record()
	if record != nil {
		t.Error("Expected nil record from empty reader")
	}

	if reader.Schema() != nil {
		t.Error("Expected nil schema from empty reader")
	}

	reader.Release()
}

func TestRecordReader_MultipleRecords(t *testing.T) {
	backend := NewMemoryBackend()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "value", Type: arrow.PrimitiveTypes.Float64}},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Add 3 records
	for i := 0; i < 3; i++ {
		builder := array.NewFloat64Builder(pool)
		builder.Append(float64(i))
		arr := builder.NewArray()
		defer arr.Release()

		record := array.NewRecord(schema, []arrow.Array{arr}, 1)
		backend.AddRecord(record)
	}

	reader := NewRecordReader(backend)
	defer reader.Release()

	// Read all records
	recordCount := 0
	for reader.Next() {
		record := reader.Record()
		if record.NumRows() != 1 {
			t.Errorf("Record %d: expected 1 row, got %d", recordCount, record.NumRows())
		}
		recordCount++
	}

	if recordCount != 3 {
		t.Errorf("Expected 3 records, read %d", recordCount)
	}

	if reader.Err() != nil {
		t.Errorf("Unexpected error: %v", reader.Err())
	}
}
