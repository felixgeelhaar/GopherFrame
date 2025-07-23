// Package persistence contains infrastructure for data persistence.
// This maps to the original storage package but follows DDD naming conventions.
package persistence

import (
	"context"

	"github.com/apache/arrow-go/v18/arrow"
)

// Backend defines the interface for pluggable persistence backends.
// This is the same as the original storage.Backend interface.
type Backend interface {
	Read(ctx context.Context, source string, opts ReadOptions) (RecordReader, error)
	Write(ctx context.Context, destination string, records RecordReader, opts WriteOptions) error
	Scan(ctx context.Context, pattern string) ([]SourceInfo, error)
	Schema(ctx context.Context, source string) (*arrow.Schema, error)
	Close() error
}

// RecordReader provides streaming access to Arrow Records.
type RecordReader interface {
	Next() bool
	Record() arrow.Record
	Schema() *arrow.Schema
	Err() error
	Close() error
}

// ReadOptions contains configuration for read operations.
type ReadOptions struct {
	Columns   []string
	Filter    string
	Limit     int64
	BatchSize int
	Schema    *arrow.Schema
	Options   map[string]interface{}
}

// WriteOptions contains configuration for write operations.
type WriteOptions struct {
	Overwrite        bool
	PartitionColumns []string
	Compression      string
	BatchSize        int
	Options          map[string]interface{}
}

// SourceInfo contains metadata about a data source.
type SourceInfo struct {
	Name     string
	Path     string
	Size     int64
	Rows     int64
	Schema   *arrow.Schema
	Modified int64
	Metadata map[string]string
}

// StreamReader adapts an io.Reader to RecordReader interface.
type StreamReader struct {
	// Implementation would contain Arrow IPC stream reader
}

// Registry manages available persistence backends.
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
