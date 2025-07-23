// Package storage defines interfaces for pluggable storage backends.
// This allows GopherFrame to support multiple data sources and formats
// while maintaining a consistent API.
package storage

import (
	"context"

	"github.com/apache/arrow-go/v18/arrow"
)

// Backend defines the interface for pluggable storage backends.
// All storage implementations must provide read/write capabilities
// and return data in Apache Arrow format for consistency.
type Backend interface {
	// Read loads data from a source and returns it as Arrow Records.
	// The source parameter format depends on the backend implementation
	// (e.g., file path, connection string, URL, etc.).
	Read(ctx context.Context, source string, opts ReadOptions) (RecordReader, error)

	// Write saves Arrow Records to a destination.
	// The destination parameter format depends on the backend implementation.
	Write(ctx context.Context, destination string, records RecordReader, opts WriteOptions) error

	// Scan returns available sources matching a pattern.
	// Used for listing tables, files, partitions, etc.
	Scan(ctx context.Context, pattern string) ([]SourceInfo, error)

	// Schema returns the schema for a given source without reading the data.
	Schema(ctx context.Context, source string) (*arrow.Schema, error)

	// Close releases any resources held by the backend.
	Close() error
}

// RecordReader provides streaming access to Arrow Records.
// This abstraction allows for memory-efficient processing of large datasets.
type RecordReader interface {
	// Next returns true if there are more records to read.
	Next() bool

	// Record returns the current Arrow Record.
	// Valid only after a successful call to Next().
	Record() arrow.Record

	// Schema returns the schema of the records.
	Schema() *arrow.Schema

	// Err returns any error that occurred during reading.
	Err() error

	// Close releases resources held by the reader.
	Close() error
}

// ReadOptions contains configuration for read operations.
type ReadOptions struct {
	// Columns specifies which columns to read (projection).
	// Empty slice means read all columns.
	Columns []string

	// Filter specifies row-level filtering predicate.
	// Backend implementations may push down filters for optimization.
	Filter string

	// Limit specifies maximum number of rows to read.
	// 0 means no limit.
	Limit int64

	// BatchSize specifies the preferred number of rows per batch.
	// Backends may use this for memory optimization.
	BatchSize int

	// Schema can be provided to override automatic schema detection.
	Schema *arrow.Schema

	// Additional backend-specific options.
	Options map[string]interface{}
}

// WriteOptions contains configuration for write operations.
type WriteOptions struct {
	// Overwrite specifies whether to overwrite existing data.
	Overwrite bool

	// PartitionColumns specifies columns to use for partitioning.
	PartitionColumns []string

	// Compression specifies the compression algorithm to use.
	Compression string

	// BatchSize specifies the number of rows to write per batch.
	BatchSize int

	// Additional backend-specific options.
	Options map[string]interface{}
}

// SourceInfo contains metadata about a data source.
type SourceInfo struct {
	// Name is the source identifier.
	Name string

	// Path is the full path or identifier for the source.
	Path string

	// Size is the size in bytes, if available.
	Size int64

	// Rows is the number of rows, if available.
	Rows int64

	// Schema is the data schema, if available.
	Schema *arrow.Schema

	// Modified is the last modification time, if available.
	Modified int64

	// Additional metadata.
	Metadata map[string]string
}

// StreamReader adapts an io.Reader to RecordReader interface.
// This is useful for backends that provide streaming data access.
type StreamReader struct {
	// Implementation would contain Arrow IPC stream reader
}

// Registry manages available storage backends.
// This allows for dynamic registration of new backend implementations.
type Registry struct {
	backends map[string]func() Backend
}

// NewRegistry creates a new backend registry.
func NewRegistry() *Registry {
	return &Registry{
		backends: make(map[string]func() Backend),
	}
}

// Register adds a backend factory function to the registry.
func (r *Registry) Register(name string, factory func() Backend) {
	r.backends[name] = factory
}

// Create instantiates a backend by name.
func (r *Registry) Create(name string) (Backend, error) {
	factory, exists := r.backends[name]
	if !exists {
		return nil, ErrBackendNotFound
	}
	return factory(), nil
}

// List returns names of all registered backends.
func (r *Registry) List() []string {
	names := make([]string, 0, len(r.backends))
	for name := range r.backends {
		names = append(names, name)
	}
	return names
}
