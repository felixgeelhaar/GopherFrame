package persistence

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
)

func TestSourceInfo(t *testing.T) {
	info := SourceInfo{
		Name: "test.parquet",
		Size: 1024,
		Path: "/path/to/test.parquet",
		Rows: 100,
	}

	if info.Name != "test.parquet" {
		t.Errorf("Expected name 'test.parquet', got %s", info.Name)
	}

	if info.Size != 1024 {
		t.Errorf("Expected size 1024, got %d", info.Size)
	}

	if info.Path != "/path/to/test.parquet" {
		t.Errorf("Expected path '/path/to/test.parquet', got %s", info.Path)
	}

	if info.Rows != 100 {
		t.Errorf("Expected rows 100, got %d", info.Rows)
	}
}

func TestReadWriteOptions(t *testing.T) {
	readOpts := ReadOptions{
		Columns: []string{"id", "name"},
		Filter:  "id > 100",
		Limit:   500,
	}

	if len(readOpts.Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(readOpts.Columns))
	}

	if readOpts.Filter != "id > 100" {
		t.Errorf("Expected filter 'id > 100', got %s", readOpts.Filter)
	}

	if readOpts.Limit != 500 {
		t.Errorf("Expected limit 500, got %d", readOpts.Limit)
	}

	writeOpts := WriteOptions{
		Compression: "snappy",
		BatchSize:   1000,
		Overwrite:   true,
	}

	if writeOpts.Compression != "snappy" {
		t.Errorf("Expected compression 'snappy', got %s", writeOpts.Compression)
	}

	if writeOpts.BatchSize != 1000 {
		t.Errorf("Expected batch size 1000, got %d", writeOpts.BatchSize)
	}

	if !writeOpts.Overwrite {
		t.Error("Expected overwrite to be true")
	}
}

// MemoryBackend is a simple in-memory implementation for testing
type MemoryBackend struct {
	data map[string][]byte
}

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{
		data: make(map[string][]byte),
	}
}

func (m *MemoryBackend) Read(ctx context.Context, source string, opts ReadOptions) (RecordReader, error) {
	if _, exists := m.data[source]; !exists {
		return nil, ErrSourceNotFound
	}
	return &mockRecordReader{}, nil
}

func (m *MemoryBackend) Write(ctx context.Context, destination string, records RecordReader, opts WriteOptions) error {
	if records == nil {
		return ErrInvalidSource
	}
	m.data[destination] = []byte("mock data")
	return nil
}

func (m *MemoryBackend) Scan(ctx context.Context, pattern string) ([]SourceInfo, error) {
	var results []SourceInfo
	for name := range m.data {
		results = append(results, SourceInfo{
			Name: name,
			Size: int64(len(m.data[name])),
			Path: "/memory/" + name,
		})
	}
	return results, nil
}

func (m *MemoryBackend) Schema(ctx context.Context, source string) (*arrow.Schema, error) {
	if _, exists := m.data[source]; !exists {
		return nil, ErrSourceNotFound
	}
	// Return a simple schema for testing
	return arrow.NewSchema([]arrow.Field{
		{Name: "test", Type: arrow.PrimitiveTypes.Int64},
	}, nil), nil
}

func (m *MemoryBackend) Close() error {
	m.data = nil
	return nil
}

// mockRecordReader for testing
type mockRecordReader struct{}

func (m *mockRecordReader) Next() bool            { return false }
func (m *mockRecordReader) Record() arrow.Record  { return nil }
func (m *mockRecordReader) Schema() *arrow.Schema { return nil }
func (m *mockRecordReader) Err() error            { return nil }
func (m *mockRecordReader) Close() error          { return nil }

func TestMemoryBackend(t *testing.T) {
	backend := NewMemoryBackend()
	defer backend.Close()

	ctx := context.Background()

	// Test Write
	mockReader := &mockRecordReader{}
	err := backend.Write(ctx, "test.data", mockReader, WriteOptions{})
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Test Read
	reader, err := backend.Read(ctx, "test.data", ReadOptions{})
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	reader.Close()

	// Test Read non-existent
	_, err = backend.Read(ctx, "nonexistent", ReadOptions{})
	if err != ErrSourceNotFound {
		t.Errorf("Expected ErrSourceNotFound, got %v", err)
	}

	// Test Schema
	schema, err := backend.Schema(ctx, "test.data")
	if err != nil {
		t.Fatalf("Schema failed: %v", err)
	}
	if schema == nil {
		t.Error("Expected schema, got nil")
	}

	// Test Scan
	sources, err := backend.Scan(ctx, "*")
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	if len(sources) != 1 {
		t.Errorf("Expected 1 source, got %d", len(sources))
	}

	if sources[0].Name != "test.data" {
		t.Errorf("Expected source name 'test.data', got %s", sources[0].Name)
	}
}

func TestRegistry(t *testing.T) {
	registry := NewRegistry()

	// Test empty registry
	backends := registry.List()
	if len(backends) != 0 {
		t.Errorf("Expected empty registry, got %d backends", len(backends))
	}

	// Test Register
	registry.Register("memory", func() Backend {
		return NewMemoryBackend()
	})

	backends = registry.List()
	if len(backends) != 1 {
		t.Errorf("Expected 1 backend, got %d", len(backends))
	}

	// Test Create
	backend, err := registry.Create("memory")
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if backend == nil {
		t.Error("Expected backend, got nil")
	}
	defer backend.Close()

	// Test Create non-existent
	_, err = registry.Create("nonexistent")
	if err != ErrBackendNotFound {
		t.Errorf("Expected ErrBackendNotFound, got %v", err)
	}
}
