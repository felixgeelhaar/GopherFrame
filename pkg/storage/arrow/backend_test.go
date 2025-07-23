package arrow

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/ipc"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/storage"
)

func TestNewBackend(t *testing.T) {
	backend := NewBackend()
	if backend == nil {
		t.Fatal("NewBackend should not return nil")
	}

	if _, ok := backend.(*Backend); !ok {
		t.Error("NewBackend should return *Backend type")
	}
}

func TestBackend_Close(t *testing.T) {
	backend := NewBackend()
	err := backend.Close()
	if err != nil {
		t.Errorf("Close should not return error, got: %v", err)
	}
}

func TestBackend_ReadWrite_RoundTrip(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create test data
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build test arrays
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{10.5, 20.3, 30.1, 40.7, 50.9}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	// Create record
	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, valueArray}, 5)
	defer record.Release()

	// Create temporary file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test.arrow")

	// Create mock record reader
	mockReader := &mockRecordReader{
		records: []arrow.Record{record},
		schema:  schema,
		index:   0,
	}

	// Test Write
	ctx := context.Background()
	writeOpts := storage.WriteOptions{
		Overwrite: true,
	}

	err := backend.Write(ctx, filename, mockReader, writeOpts)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Verify file exists
	if _, statErr := os.Stat(filename); os.IsNotExist(statErr) {
		t.Fatal("Output file was not created")
	}

	// Test Read
	readOpts := storage.ReadOptions{}
	reader, err := backend.Read(ctx, filename, readOpts)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	defer reader.Close()

	// Verify schema
	readSchema := reader.Schema()
	if !readSchema.Equal(schema) {
		t.Error("Read schema does not match original schema")
	}

	// Read records
	var readRecords []arrow.Record
	for reader.Next() {
		record := reader.Record()
		record.Retain()
		readRecords = append(readRecords, record)
	}

	if reader.Err() != nil {
		t.Fatalf("Reader error: %v", reader.Err())
	}

	// Verify we got records
	if len(readRecords) == 0 {
		t.Fatal("No records were read back")
	}

	// Verify first record content
	firstRecord := readRecords[0]
	defer firstRecord.Release()

	if firstRecord.NumRows() != 5 {
		t.Errorf("Expected 5 rows, got %d", firstRecord.NumRows())
	}

	if firstRecord.NumCols() != 3 {
		t.Errorf("Expected 3 columns, got %d", firstRecord.NumCols())
	}

	// Verify data in first column (id)
	idCol := firstRecord.Column(0).(*array.Int64)
	expectedIds := []int64{1, 2, 3, 4, 5}
	for i, expected := range expectedIds {
		if idCol.Value(i) != expected {
			t.Errorf("ID at index %d: expected %d, got %d", i, expected, idCol.Value(i))
		}
	}
}

func TestBackend_Schema(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create test data file first
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "schema_test.arrow")

	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "test_id", Type: arrow.PrimitiveTypes.Int32},
			{Name: "test_name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Create a simple Arrow file
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	writer, err := ipc.NewFileWriter(file, ipc.WithSchema(schema))
	if err != nil {
		file.Close()
		t.Fatalf("Failed to create writer: %v", err)
	}

	// Write an empty record to establish the schema
	idBuilder := array.NewInt32Builder(pool)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	emptyRecord := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 0)
	defer emptyRecord.Release()

	if err := writer.Write(emptyRecord); err != nil {
		writer.Close()
		file.Close()
		t.Fatalf("Failed to write record: %v", err)
	}

	// Properly close writer and file
	if closeErr := writer.Close(); closeErr != nil {
		file.Close()
		t.Fatalf("Failed to close writer: %v", closeErr)
	}

	if closeErr := file.Close(); closeErr != nil {
		t.Fatalf("Failed to close file: %v", closeErr)
	}

	// Test Schema method
	ctx := context.Background()
	readSchema, err := backend.Schema(ctx, filename)
	if err != nil {
		t.Fatalf("Schema failed: %v", err)
	}

	if !readSchema.Equal(schema) {
		t.Error("Read schema does not match original schema")
	}

	// Verify field details
	if len(readSchema.Fields()) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(readSchema.Fields()))
	}

	if readSchema.Field(0).Name != "test_id" {
		t.Errorf("Expected first field name 'test_id', got %s", readSchema.Field(0).Name)
	}

	if readSchema.Field(1).Name != "test_name" {
		t.Errorf("Expected second field name 'test_name', got %s", readSchema.Field(1).Name)
	}
}

func TestBackend_Read_NonExistentFile(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	ctx := context.Background()
	readOpts := storage.ReadOptions{}

	_, err := backend.Read(ctx, "/nonexistent/file.arrow", readOpts)
	if err == nil {
		t.Error("Reading non-existent file should return error")
	}

	if err != storage.ErrSourceNotFound {
		t.Errorf("Expected ErrSourceNotFound, got: %v", err)
	}
}

func TestBackend_Read_InvalidSource(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	ctx := context.Background()
	readOpts := storage.ReadOptions{}

	_, err := backend.Read(ctx, "", readOpts)
	if err == nil {
		t.Error("Reading with empty source should return error")
	}

	if err != storage.ErrInvalidSource {
		t.Errorf("Expected ErrInvalidSource, got: %v", err)
	}
}

func TestBackend_Write_InvalidDestination(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	ctx := context.Background()
	writeOpts := storage.WriteOptions{}
	mockReader := &mockRecordReader{}

	err := backend.Write(ctx, "", mockReader, writeOpts)
	if err == nil {
		t.Error("Writing with empty destination should return error")
	}

	if err != storage.ErrInvalidSource {
		t.Errorf("Expected ErrInvalidSource, got: %v", err)
	}
}

func TestBackend_Write_OverwriteProtection(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create existing file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "existing.arrow")

	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	ctx := context.Background()
	writeOpts := storage.WriteOptions{
		Overwrite: false, // Don't allow overwrite
	}
	mockReader := &mockRecordReader{}

	err = backend.Write(ctx, filename, mockReader, writeOpts)
	if err == nil {
		t.Error("Writing to existing file without overwrite should return error")
	}
}

func TestBackend_Scan(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create test directory with Arrow files
	tempDir := t.TempDir()

	// Create some test files
	arrowFile1 := filepath.Join(tempDir, "test1.arrow")
	arrowFile2 := filepath.Join(tempDir, "test2.ipc")
	nonArrowFile := filepath.Join(tempDir, "test.txt")

	for _, filename := range []string{arrowFile1, arrowFile2, nonArrowFile} {
		file, err := os.Create(filename)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", filename, err)
		}
		file.Close()
	}

	ctx := context.Background()
	pattern := filepath.Join(tempDir, "*")

	sources, err := backend.Scan(ctx, pattern)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	// Should find 2 Arrow files (.arrow and .ipc), not the .txt file
	if len(sources) != 2 {
		t.Errorf("Expected 2 sources, got %d", len(sources))
	}

	// Verify that non-Arrow file is not included
	for _, source := range sources {
		if filepath.Ext(source.Path) == ".txt" {
			t.Error("Non-Arrow file should not be included in scan results")
		}
	}
}

// mockRecordReader implements storage.RecordReader for testing
type mockRecordReader struct {
	records []arrow.Record
	schema  *arrow.Schema
	index   int
	err     error
}

func (m *mockRecordReader) Next() bool {
	if m.err != nil {
		return false
	}
	return m.index < len(m.records)
}

func (m *mockRecordReader) Record() arrow.Record {
	if m.index >= len(m.records) {
		return nil
	}
	record := m.records[m.index]
	m.index++
	return record
}

func (m *mockRecordReader) Schema() *arrow.Schema {
	return m.schema
}

func (m *mockRecordReader) Err() error {
	return m.err
}

func (m *mockRecordReader) Close() error {
	return nil
}

func TestBackend_Read_CorruptedFile(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create a corrupted Arrow file (not valid Arrow format)
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "corrupted.arrow")

	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Write invalid content
	_, err = file.WriteString("This is not a valid Arrow file")
	if err != nil {
		file.Close()
		t.Fatalf("Failed to write invalid content: %v", err)
	}
	file.Close()

	ctx := context.Background()
	readOpts := storage.ReadOptions{}

	_, err = backend.Read(ctx, filename, readOpts)
	if err == nil {
		t.Error("Reading corrupted file should return error")
	}
}

func TestBackend_Schema_NonExistentFile(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	ctx := context.Background()
	_, err := backend.Schema(ctx, "/nonexistent/file.arrow")
	if err == nil {
		t.Error("Schema on non-existent file should return error")
	}
}

func TestBackend_Schema_CorruptedFile(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create a corrupted Arrow file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "corrupted.arrow")

	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err = file.WriteString("Invalid Arrow content")
	if err != nil {
		file.Close()
		t.Fatalf("Failed to write invalid content: %v", err)
	}
	file.Close()

	ctx := context.Background()
	_, err = backend.Schema(ctx, filename)
	if err == nil {
		t.Error("Schema on corrupted file should return error")
	}
}

func TestBackend_Write_DirectoryCreation(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create test data
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray}, 2)
	defer record.Release()

	mockReader := &mockRecordReader{
		records: []arrow.Record{record},
		schema:  schema,
		index:   0,
	}

	// Test writing to a nested directory that doesn't exist
	tempDir := t.TempDir()
	nestedDir := filepath.Join(tempDir, "nested", "deeper")
	filename := filepath.Join(nestedDir, "test.arrow")

	ctx := context.Background()
	writeOpts := storage.WriteOptions{Overwrite: true}

	err := backend.Write(ctx, filename, mockReader, writeOpts)
	if err != nil {
		t.Fatalf("Write should succeed with directory creation: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("File should have been created in nested directory")
	}
}

func TestBackend_Write_FileCreationError(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	mockReader := &mockRecordReader{}

	// Try to write to a directory instead of a file (should fail)
	tempDir := t.TempDir()

	ctx := context.Background()
	writeOpts := storage.WriteOptions{}

	err := backend.Write(ctx, tempDir, mockReader, writeOpts) // tempDir is a directory, not a file
	if err == nil {
		t.Error("Writing to a directory should return error")
	}
}

func TestRecordReader_WithLimit(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create test data with multiple records
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1}, nil)
	idArray1 := idBuilder.NewArray()
	defer idArray1.Release()

	idBuilder2 := array.NewInt64Builder(pool)
	idBuilder2.AppendValues([]int64{2}, nil)
	idArray2 := idBuilder2.NewArray()
	defer idArray2.Release()

	idBuilder3 := array.NewInt64Builder(pool)
	idBuilder3.AppendValues([]int64{3}, nil)
	idArray3 := idBuilder3.NewArray()
	defer idArray3.Release()

	record1 := array.NewRecord(schema, []arrow.Array{idArray1}, 1)
	defer record1.Release()

	record2 := array.NewRecord(schema, []arrow.Array{idArray2}, 1)
	defer record2.Release()

	record3 := array.NewRecord(schema, []arrow.Array{idArray3}, 1)
	defer record3.Release()

	// Create test file with multiple records
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "multi_record.arrow")

	mockWriter := &mockRecordReader{
		records: []arrow.Record{record1, record2, record3},
		schema:  schema,
		index:   0,
	}

	ctx := context.Background()
	writeOpts := storage.WriteOptions{Overwrite: true}

	err := backend.Write(ctx, filename, mockWriter, writeOpts)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Read with limit
	readOpts := storage.ReadOptions{Limit: 2}
	reader, err := backend.Read(ctx, filename, readOpts)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	defer reader.Close()

	// Count records read (should be limited to 2)
	recordCount := 0
	for reader.Next() {
		record := reader.Record()
		record.Retain()
		defer record.Release()
		recordCount++
	}

	if reader.Err() != nil {
		t.Fatalf("Reader error: %v", reader.Err())
	}

	if recordCount != 2 {
		t.Errorf("Expected 2 records due to limit, got %d", recordCount)
	}
}

func TestRecordReader_ContextCancellation(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create test data
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray}, 1)
	defer record.Release()

	// Create test file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "context_test.arrow")

	mockWriter := &mockRecordReader{
		records: []arrow.Record{record},
		schema:  schema,
		index:   0,
	}

	ctx := context.Background()
	writeOpts := storage.WriteOptions{Overwrite: true}

	err := backend.Write(ctx, filename, mockWriter, writeOpts)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Create cancelled context
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	readOpts := storage.ReadOptions{}
	reader, err := backend.Read(cancelledCtx, filename, readOpts)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	defer reader.Close()

	// Try to read - should detect context cancellation
	hasNext := reader.Next()
	if hasNext {
		t.Error("Next should return false for cancelled context")
	}

	if reader.Err() != context.Canceled {
		t.Errorf("Expected context.Canceled error, got: %v", reader.Err())
	}
}

func TestRecordReader_RecordAtError(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	// Create a simple Arrow file
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray}, 1)
	defer record.Release()

	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "record_error_test.arrow")

	mockWriter := &mockRecordReader{
		records: []arrow.Record{record},
		schema:  schema,
		index:   0,
	}

	ctx := context.Background()
	writeOpts := storage.WriteOptions{Overwrite: true}

	err := backend.Write(ctx, filename, mockWriter, writeOpts)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	readOpts := storage.ReadOptions{}
	reader, err := backend.Read(ctx, filename, readOpts)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	defer reader.Close()

	// Read all available records to exhaust the reader
	for reader.Next() {
		_ = reader.Record()
	}

	// Try to read beyond available records
	nextRecord := reader.Record()
	if nextRecord != nil {
		t.Error("Record() should return nil when no more records available")
	}
}

func TestBackend_Scan_InvalidPattern(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	ctx := context.Background()

	// Use an invalid glob pattern
	_, err := backend.Scan(ctx, "[invalid")
	if err == nil {
		t.Error("Scan with invalid pattern should return error")
	}
}

func TestBackend_Scan_EmptyDirectory(t *testing.T) {
	backend := NewBackend()
	defer backend.Close()

	tempDir := t.TempDir()
	pattern := filepath.Join(tempDir, "*.arrow")

	ctx := context.Background()
	sources, err := backend.Scan(ctx, pattern)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	if len(sources) != 0 {
		t.Errorf("Expected 0 sources in empty directory, got %d", len(sources))
	}
}
