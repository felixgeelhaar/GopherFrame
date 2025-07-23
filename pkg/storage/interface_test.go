package storage

import (
	"context"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
)

func TestSourceInfo(t *testing.T) {
	info := SourceInfo{
		Name:     "test.parquet",
		Path:     "/path/to/test.parquet",
		Size:     1024,
		Rows:     100,
		Modified: 1640995200, // Unix timestamp
		Metadata: map[string]string{
			"format":      "parquet",
			"compression": "snappy",
		},
	}

	if info.Name != "test.parquet" {
		t.Errorf("Expected name 'test.parquet', got %s", info.Name)
	}

	if info.Path != "/path/to/test.parquet" {
		t.Errorf("Expected path '/path/to/test.parquet', got %s", info.Path)
	}

	if info.Size != 1024 {
		t.Errorf("Expected size 1024, got %d", info.Size)
	}

	if info.Rows != 100 {
		t.Errorf("Expected rows 100, got %d", info.Rows)
	}

	if info.Modified != 1640995200 {
		t.Errorf("Expected modified 1640995200, got %d", info.Modified)
	}

	if len(info.Metadata) != 2 {
		t.Errorf("Expected 2 metadata entries, got %d", len(info.Metadata))
	}

	if info.Metadata["format"] != "parquet" {
		t.Errorf("Expected format 'parquet', got %s", info.Metadata["format"])
	}
}

func TestReadOptions(t *testing.T) {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	opts := ReadOptions{
		Columns:   []string{"id", "name"},
		Filter:    "id > 100",
		Limit:     500,
		BatchSize: 1000,
		Schema:    schema,
		Options: map[string]interface{}{
			"parallel": true,
		},
	}

	if len(opts.Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(opts.Columns))
	}

	if opts.Columns[0] != "id" || opts.Columns[1] != "name" {
		t.Errorf("Expected columns [id, name], got %v", opts.Columns)
	}

	if opts.Filter != "id > 100" {
		t.Errorf("Expected filter 'id > 100', got %s", opts.Filter)
	}

	if opts.Limit != 500 {
		t.Errorf("Expected limit 500, got %d", opts.Limit)
	}

	if opts.BatchSize != 1000 {
		t.Errorf("Expected batch size 1000, got %d", opts.BatchSize)
	}

	if opts.Schema == nil {
		t.Error("Expected schema to be set")
	}

	if len(opts.Options) != 1 {
		t.Errorf("Expected 1 option, got %d", len(opts.Options))
	}

	if parallel, ok := opts.Options["parallel"].(bool); !ok || !parallel {
		t.Error("Expected parallel option to be true")
	}
}

func TestWriteOptions(t *testing.T) {
	opts := WriteOptions{
		Overwrite:        true,
		PartitionColumns: []string{"year", "month"},
		Compression:      "snappy",
		BatchSize:        2000,
		Options: map[string]interface{}{
			"row_group_size": 50000,
		},
	}

	if !opts.Overwrite {
		t.Error("Expected overwrite to be true")
	}

	if len(opts.PartitionColumns) != 2 {
		t.Errorf("Expected 2 partition columns, got %d", len(opts.PartitionColumns))
	}

	if opts.PartitionColumns[0] != "year" || opts.PartitionColumns[1] != "month" {
		t.Errorf("Expected partition columns [year, month], got %v", opts.PartitionColumns)
	}

	if opts.Compression != "snappy" {
		t.Errorf("Expected compression 'snappy', got %s", opts.Compression)
	}

	if opts.BatchSize != 2000 {
		t.Errorf("Expected batch size 2000, got %d", opts.BatchSize)
	}

	if len(opts.Options) != 1 {
		t.Errorf("Expected 1 option, got %d", len(opts.Options))
	}

	if rowGroupSize, ok := opts.Options["row_group_size"].(int); !ok || rowGroupSize != 50000 {
		t.Error("Expected row_group_size option to be 50000")
	}
}

func TestStreamReader(_ *testing.T) {
	// Test that StreamReader struct can be instantiated
	// Since it's currently just a placeholder, we just test basic construction
	var reader StreamReader

	// StreamReader should be a valid struct (even if empty)
	_ = reader
}

func TestNewRegistry(t *testing.T) {
	registry := NewRegistry()

	if registry == nil {
		t.Fatal("NewRegistry should not return nil")
	}

	if registry.backends == nil {
		t.Error("Registry backends map should be initialized")
	}

	if len(registry.backends) != 0 {
		t.Error("New registry should start with no backends")
	}
}

func TestRegistry_Register(t *testing.T) {
	registry := NewRegistry()

	// Mock backend factory
	mockFactory := func() Backend {
		return &mockBackend{}
	}

	// Register backend
	registry.Register("mock", mockFactory)

	if len(registry.backends) != 1 {
		t.Errorf("Expected 1 backend after registration, got %d", len(registry.backends))
	}

	if _, exists := registry.backends["mock"]; !exists {
		t.Error("Backend should be registered with name 'mock'")
	}
}

func TestRegistry_Create(t *testing.T) {
	registry := NewRegistry()

	// Mock backend factory
	mockFactory := func() Backend {
		return &mockBackend{}
	}

	// Register backend
	registry.Register("mock", mockFactory)

	// Create backend
	backend, err := registry.Create("mock")
	if err != nil {
		t.Fatalf("Failed to create backend: %v", err)
	}

	if backend == nil {
		t.Error("Created backend should not be nil")
	}

	if _, ok := backend.(*mockBackend); !ok {
		t.Error("Created backend should be of type mockBackend")
	}

	// Test creating non-existent backend
	_, err = registry.Create("nonexistent")
	if err == nil {
		t.Error("Creating non-existent backend should return error")
	}
}

func TestRegistry_List(t *testing.T) {
	registry := NewRegistry()

	// Initially should be empty
	backends := registry.List()
	if len(backends) != 0 {
		t.Errorf("Expected 0 backends initially, got %d", len(backends))
	}

	// Register some backends
	registry.Register("arrow", func() Backend { return &mockBackend{} })
	registry.Register("parquet", func() Backend { return &mockBackend{} })
	registry.Register("csv", func() Backend { return &mockBackend{} })

	backends = registry.List()
	if len(backends) != 3 {
		t.Errorf("Expected 3 backends after registration, got %d", len(backends))
	}

	// Check that all expected backends are listed
	expectedBackends := map[string]bool{
		"arrow":   false,
		"parquet": false,
		"csv":     false,
	}

	for _, name := range backends {
		if _, exists := expectedBackends[name]; exists {
			expectedBackends[name] = true
		} else {
			t.Errorf("Unexpected backend in list: %s", name)
		}
	}

	for name, found := range expectedBackends {
		if !found {
			t.Errorf("Expected backend %s not found in list", name)
		}
	}
}

// mockBackend is a simple mock implementation for testing
type mockBackend struct{}

func (m *mockBackend) Read(_ context.Context, _ string, _ ReadOptions) (RecordReader, error) {
	return nil, nil
}

func (m *mockBackend) Write(_ context.Context, _ string, _ RecordReader, _ WriteOptions) error {
	return nil
}

func (m *mockBackend) Scan(_ context.Context, _ string) ([]SourceInfo, error) {
	return nil, nil
}

func (m *mockBackend) Schema(_ context.Context, _ string) (*arrow.Schema, error) {
	return nil, nil
}

func (m *mockBackend) Close() error {
	return nil
}
