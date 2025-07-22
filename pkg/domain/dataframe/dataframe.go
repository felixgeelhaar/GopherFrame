// Package dataframe contains the core DataFrame domain model and business logic.
// This is the heart of the domain - immutable, strongly-typed data structures
// that provide zero-copy operations and seamless Arrow interoperability.
package dataframe

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/infrastructure/persistence"
)

// DataFrame is the domain model representing tabular data.
// It is immutable and provides zero-copy operations through Apache Arrow.
type DataFrame struct {
	record    arrow.Record
	allocator memory.Allocator
	backend   persistence.Backend // Dependency on infrastructure (acceptable in this layer)
}

// NewDataFrame creates a new DataFrame from an Arrow Record.
// This is a factory function for the domain entity.
func NewDataFrame(record arrow.Record) *DataFrame {
	record.Retain()
	return &DataFrame{
		record:    record,
		allocator: memory.DefaultAllocator,
	}
}

// NumRows returns the number of rows in the DataFrame.
func (df *DataFrame) NumRows() int64 {
	return df.record.NumRows()
}

// NumCols returns the number of columns in the DataFrame.
func (df *DataFrame) NumCols() int64 {
	return df.record.NumCols()
}

// Schema returns the Arrow schema of the DataFrame.
func (df *DataFrame) Schema() *arrow.Schema {
	return df.record.Schema()
}

// ColumnNames returns the names of all columns.
func (df *DataFrame) ColumnNames() []string {
	schema := df.record.Schema()
	names := make([]string, schema.NumFields())
	for i, field := range schema.Fields() {
		names[i] = field.Name
	}
	return names
}

// HasColumn checks if a column exists in the DataFrame.
func (df *DataFrame) HasColumn(name string) bool {
	schema := df.record.Schema()
	for _, field := range schema.Fields() {
		if field.Name == name {
			return true
		}
	}
	return false
}

// Record returns the underlying Arrow record (read-only access).
func (df *DataFrame) Record() arrow.Record {
	return df.record
}

// Release releases the underlying Arrow record.
func (df *DataFrame) Release() {
	if df.record != nil {
		df.record.Release()
	}
}

// Clone creates a shallow copy of the DataFrame.
func (df *DataFrame) Clone() *DataFrame {
	df.record.Retain() // Increment reference count
	return &DataFrame{
		record:    df.record,
		allocator: df.allocator,
		backend:   df.backend,
	}
}
