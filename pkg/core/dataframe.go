// Package core provides the internal DataFrame and Series implementations.
// These are the foundational data structures that wrap Apache Arrow Records
// and Arrays, providing immutable, strongly-typed operations.
package core

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/gopherFrame/pkg/storage"
)

// DataFrame is the internal, immutable representation of tabular data.
// It wraps an arrow.Record to provide zero-copy operations and seamless
// interoperability with the Arrow ecosystem.
type DataFrame struct {
	// record is the underlying Arrow Record containing the actual data.
	// This is never exposed directly to maintain immutability.
	record arrow.Record

	// allocator is used for memory management when creating new records.
	allocator memory.Allocator

	// storage backend for I/O operations (optional)
	backend storage.Backend
}

// NewDataFrame creates a new DataFrame from an Arrow Record.
// The DataFrame takes ownership of the record and will release it when closed.
func NewDataFrame(record arrow.Record) *DataFrame {
	record.Retain() // Increment reference count
	return &DataFrame{
		record:    record,
		allocator: memory.DefaultAllocator,
	}
}

// NewDataFrameWithAllocator creates a new DataFrame with a custom allocator.
func NewDataFrameWithAllocator(record arrow.Record, allocator memory.Allocator) *DataFrame {
	record.Retain()
	return &DataFrame{
		record:    record,
		allocator: allocator,
	}
}

// NewDataFrameFromStorage creates a DataFrame by reading from a storage backend.
func NewDataFrameFromStorage(ctx context.Context, backend storage.Backend, source string, opts storage.ReadOptions) (*DataFrame, error) {
	reader, err := backend.Read(ctx, source, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to read from storage: %w", err)
	}
	defer reader.Close()

	// For now, read the first record. Future versions will handle multiple records.
	if !reader.Next() {
		if err := reader.Err(); err != nil {
			return nil, fmt.Errorf("failed to read record: %w", err)
		}
		return nil, fmt.Errorf("no records found in source: %s", source)
	}

	record := reader.Record()
	record.Retain() // Take ownership

	df := &DataFrame{
		record:    record,
		allocator: memory.DefaultAllocator,
		backend:   backend,
	}

	return df, nil
}

// Schema returns the Arrow schema of the DataFrame.
func (df *DataFrame) Schema() *arrow.Schema {
	return df.record.Schema()
}

// NumRows returns the number of rows in the DataFrame.
func (df *DataFrame) NumRows() int64 {
	return df.record.NumRows()
}

// NumCols returns the number of columns in the DataFrame.
func (df *DataFrame) NumCols() int64 {
	return df.record.NumCols()
}

// Record returns the underlying Arrow Record.
// This is used internally by operations that need direct Arrow access.
func (df *DataFrame) Record() arrow.Record {
	return df.record
}

// ColumnNames returns the names of all columns in order.
func (df *DataFrame) ColumnNames() []string {
	schema := df.record.Schema()
	names := make([]string, schema.NumFields())
	for i, field := range schema.Fields() {
		names[i] = field.Name
	}
	return names
}

// Column returns a Series for the specified column name.
func (df *DataFrame) Column(name string) (*Series, error) {
	schema := df.record.Schema()
	fieldIndex := -1

	// Find the column index
	for i, field := range schema.Fields() {
		if field.Name == name {
			fieldIndex = i
			break
		}
	}

	if fieldIndex == -1 {
		return nil, fmt.Errorf("column not found: %s", name)
	}

	column := df.record.Column(fieldIndex)
	field := schema.Field(fieldIndex)

	return NewSeries(column, field), nil
}

// ColumnAt returns a Series for the column at the specified index.
func (df *DataFrame) ColumnAt(index int) (*Series, error) {
	if index < 0 || index >= int(df.NumCols()) {
		return nil, fmt.Errorf("column index out of range: %d", index)
	}

	column := df.record.Column(index)
	field := df.record.Schema().Field(index)

	return NewSeries(column, field), nil
}

// Columns returns all columns as a slice of Series.
func (df *DataFrame) Columns() []*Series {
	numCols := int(df.NumCols())
	series := make([]*Series, numCols)

	for i := 0; i < numCols; i++ {
		column := df.record.Column(i)
		field := df.record.Schema().Field(i)
		series[i] = NewSeries(column, field)
	}

	return series
}

// HasColumn checks if a column with the given name exists.
func (df *DataFrame) HasColumn(name string) bool {
	schema := df.record.Schema()
	for _, field := range schema.Fields() {
		if field.Name == name {
			return true
		}
	}
	return false
}

// Equal compares two DataFrames for equality.
// Returns true if they have the same schema and data.
func (df *DataFrame) Equal(other *DataFrame) bool {
	if df == other {
		return true
	}

	if other == nil {
		return false
	}

	// Compare schemas first
	if !df.record.Schema().Equal(other.record.Schema()) {
		return false
	}
	
	// Compare number of rows
	if df.record.NumRows() != other.record.NumRows() {
		return false
	}
	
	// Compare each column
	for i := 0; i < int(df.record.NumCols()); i++ {
		if !array.Equal(df.record.Column(i), other.record.Column(i)) {
			return false
		}
	}
	
	return true
}

// Validate checks the DataFrame for consistency and data integrity.
func (df *DataFrame) Validate() error {
	if df.record == nil {
		return fmt.Errorf("DataFrame has no underlying record")
	}

	// Validate schema and columns consistency
	schema := df.record.Schema()
	if int(df.record.NumCols()) != len(schema.Fields()) {
		return fmt.Errorf("column count mismatch: record has %d columns, schema has %d fields", 
			df.record.NumCols(), len(schema.Fields()))
	}
	
	// Validate each column against its field type
	for i, field := range schema.Fields() {
		column := df.record.Column(i)
		if !arrow.TypeEqual(column.DataType(), field.Type) {
			return fmt.Errorf("column %d type mismatch: expected %s, got %s", 
				i, field.Type, column.DataType())
		}
	}
	
	return nil
}

// String returns a string representation of the DataFrame.
// This is primarily for debugging and should not be used for large DataFrames.
func (df *DataFrame) String() string {
	if df.record == nil {
		return "DataFrame{<empty>}"
	}

	return fmt.Sprintf("DataFrame{rows: %d, cols: %d, schema: %s}",
		df.NumRows(), df.NumCols(), df.Schema())
}

// Clone creates a shallow copy of the DataFrame.
// The underlying Arrow data is shared (copy-on-write semantics).
func (df *DataFrame) Clone() *DataFrame {
	df.record.Retain() // Increment reference count
	return &DataFrame{
		record:    df.record,
		allocator: df.allocator,
		backend:   df.backend,
	}
}

// WriteToStorage saves the DataFrame to a storage backend.
func (df *DataFrame) WriteToStorage(ctx context.Context, backend storage.Backend, destination string, opts storage.WriteOptions) error {
	if backend == nil {
		if df.backend == nil {
			return fmt.Errorf("no storage backend available")
		}
		backend = df.backend
	}

	// Create a record reader from this single record
	reader := &singleRecordReader{
		record: df.record,
		schema: df.record.Schema(),
	}

	return backend.Write(ctx, destination, reader, opts)
}

// Select returns a new DataFrame with only the specified columns.
func (df *DataFrame) Select(columnNames []string) (*DataFrame, error) {
	if len(columnNames) == 0 {
		return nil, fmt.Errorf("no columns specified for selection")
	}
	
	schema := df.record.Schema()
	
	// Find column indices and validate columns exist
	indices := make([]int, len(columnNames))
	selectedFields := make([]arrow.Field, len(columnNames))
	
	for i, name := range columnNames {
		fieldIndex := -1
		for j, field := range schema.Fields() {
			if field.Name == name {
				fieldIndex = j
				selectedFields[i] = field
				break
			}
		}
		
		if fieldIndex == -1 {
			return nil, fmt.Errorf("column not found: %s", name)
		}
		indices[i] = fieldIndex
	}
	
	// Create new schema with selected fields
	newSchema := arrow.NewSchema(selectedFields, nil)
	
	// Extract selected columns
	selectedColumns := make([]arrow.Array, len(indices))
	for i, idx := range indices {
		column := df.record.Column(idx)
		column.Retain() // Retain reference for new record
		selectedColumns[i] = column
	}
	
	// Create new record with selected columns
	newRecord := array.NewRecord(newSchema, selectedColumns, df.record.NumRows())
	
	return NewDataFrame(newRecord), nil
}

// WithColumn returns a new DataFrame with an additional or replaced column.
func (df *DataFrame) WithColumn(columnName string, newColumn arrow.Array) (*DataFrame, error) {
	if newColumn == nil {
		return nil, fmt.Errorf("new column cannot be nil")
	}
	
	// Validate column length matches DataFrame
	if int64(newColumn.Len()) != df.NumRows() {
		return nil, fmt.Errorf("column length %d does not match DataFrame rows %d", newColumn.Len(), df.NumRows())
	}
	
	newColumn.Retain() // Take ownership
	
	schema := df.record.Schema()
	
	// Check if column already exists
	existingColumnIndex := -1
	for i, field := range schema.Fields() {
		if field.Name == columnName {
			existingColumnIndex = i
			break
		}
	}
	
	var newFields []arrow.Field
	var newColumns []arrow.Array
	
	if existingColumnIndex >= 0 {
		// Replace existing column
		newFields = make([]arrow.Field, len(schema.Fields()))
		newColumns = make([]arrow.Array, len(schema.Fields()))
		
		for i, field := range schema.Fields() {
			if i == existingColumnIndex {
				// Replace with new column
				newFields[i] = arrow.Field{Name: columnName, Type: newColumn.DataType()}
				newColumns[i] = newColumn
			} else {
				// Keep existing column
				newFields[i] = field
				column := df.record.Column(i)
				column.Retain()
				newColumns[i] = column
			}
		}
	} else {
		// Add new column
		newFields = make([]arrow.Field, len(schema.Fields())+1)
		newColumns = make([]arrow.Array, len(schema.Fields())+1)
		
		// Copy existing fields and columns
		for i, field := range schema.Fields() {
			newFields[i] = field
			column := df.record.Column(i)
			column.Retain()
			newColumns[i] = column
		}
		
		// Add new column
		newFields[len(schema.Fields())] = arrow.Field{Name: columnName, Type: newColumn.DataType()}
		newColumns[len(schema.Fields())] = newColumn
	}
	
	// Create new schema and record
	newSchema := arrow.NewSchema(newFields, nil)
	newRecord := array.NewRecord(newSchema, newColumns, df.record.NumRows())
	
	return NewDataFrame(newRecord), nil
}

// Filter returns a new DataFrame containing only rows where the predicate is true.
func (df *DataFrame) Filter(predicateArray arrow.Array) (*DataFrame, error) {
	// Validate that predicate is boolean array
	if predicateArray.DataType().ID() != arrow.BOOL {
		return nil, fmt.Errorf("filter predicate must be boolean array, got %s", predicateArray.DataType())
	}
	
	// Validate length matches DataFrame
	if int64(predicateArray.Len()) != df.NumRows() {
		return nil, fmt.Errorf("predicate length %d does not match DataFrame rows %d", predicateArray.Len(), df.NumRows())
	}
	
	boolArray, ok := predicateArray.(*array.Boolean)
	if !ok {
		return nil, fmt.Errorf("failed to cast predicate to boolean array")
	}
	
	// Count true values to determine result size
	trueCount := int64(0)
	for i := 0; i < boolArray.Len(); i++ {
		if !boolArray.IsNull(i) && boolArray.Value(i) {
			trueCount++
		}
	}
	
	if trueCount == 0 {
		// Return empty DataFrame with same schema
		schema := df.record.Schema()
		emptyColumns := make([]arrow.Array, len(schema.Fields()))
		
		pool := memory.NewGoAllocator()
		for i, field := range schema.Fields() {
			switch field.Type.ID() {
			case arrow.INT64:
				builder := array.NewInt64Builder(pool)
				emptyColumns[i] = builder.NewArray()
				builder.Release()
			case arrow.FLOAT64:
				builder := array.NewFloat64Builder(pool)
				emptyColumns[i] = builder.NewArray()
				builder.Release()
			case arrow.STRING:
				builder := array.NewStringBuilder(pool)
				emptyColumns[i] = builder.NewArray()
				builder.Release()
			default:
				return nil, fmt.Errorf("unsupported data type for empty filter: %s", field.Type)
			}
		}
		
		emptyRecord := array.NewRecord(schema, emptyColumns, 0)
		return NewDataFrame(emptyRecord), nil
	}
	
	// Create filtered columns
	schema := df.record.Schema()
	filteredColumns := make([]arrow.Array, len(schema.Fields()))
	
	pool := memory.NewGoAllocator()
	for colIdx, field := range schema.Fields() {
		column := df.record.Column(colIdx)
		
		switch field.Type.ID() {
		case arrow.INT64:
			srcArray := column.(*array.Int64)
			builder := array.NewInt64Builder(pool)
			
			for i := 0; i < boolArray.Len(); i++ {
				if !boolArray.IsNull(i) && boolArray.Value(i) {
					if srcArray.IsNull(i) {
						builder.AppendNull()
					} else {
						builder.Append(srcArray.Value(i))
					}
				}
			}
			filteredColumns[colIdx] = builder.NewArray()
			builder.Release()
			
		case arrow.FLOAT64:
			srcArray := column.(*array.Float64)
			builder := array.NewFloat64Builder(pool)
			
			for i := 0; i < boolArray.Len(); i++ {
				if !boolArray.IsNull(i) && boolArray.Value(i) {
					if srcArray.IsNull(i) {
						builder.AppendNull()
					} else {
						builder.Append(srcArray.Value(i))
					}
				}
			}
			filteredColumns[colIdx] = builder.NewArray()
			builder.Release()
			
		case arrow.STRING:
			srcArray := column.(*array.String)
			builder := array.NewStringBuilder(pool)
			
			for i := 0; i < boolArray.Len(); i++ {
				if !boolArray.IsNull(i) && boolArray.Value(i) {
					if srcArray.IsNull(i) {
						builder.AppendNull()
					} else {
						builder.Append(srcArray.Value(i))
					}
				}
			}
			filteredColumns[colIdx] = builder.NewArray()
			builder.Release()
			
		default:
			return nil, fmt.Errorf("unsupported data type for filtering: %s", field.Type)
		}
	}
	
	// Create new record with filtered data
	filteredRecord := array.NewRecord(schema, filteredColumns, trueCount)
	return NewDataFrame(filteredRecord), nil
}

// Release decrements the reference count of the underlying Arrow Record.
// The DataFrame should not be used after calling Release().
func (df *DataFrame) Release() {
	if df.record != nil {
		df.record.Release()
		df.record = nil
	}
}

// singleRecordReader implements storage.RecordReader for a single Arrow Record.
// This is used when writing a DataFrame to storage.
type singleRecordReader struct {
	record   arrow.Record
	schema   *arrow.Schema
	consumed bool
	err      error
}

func (r *singleRecordReader) Next() bool {
	return !r.consumed && r.err == nil
}

func (r *singleRecordReader) Record() arrow.Record {
	if r.consumed {
		return nil
	}
	r.consumed = true
	return r.record
}

func (r *singleRecordReader) Schema() *arrow.Schema {
	return r.schema
}

func (r *singleRecordReader) Err() error {
	return r.err
}

func (r *singleRecordReader) Close() error {
	return nil // No resources to clean up
}
