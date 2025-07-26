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
	"github.com/felixgeelhaar/GopherFrame/pkg/storage"
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
	defer func() { _ = reader.Close() }()

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
			case arrow.BOOL:
				builder := array.NewBooleanBuilder(pool)
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
			srcArray, ok := column.(*array.Int64)
			if !ok {
				return nil, fmt.Errorf("expected Int64 array for column %d", colIdx)
			}
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
			srcArray, ok := column.(*array.Float64)
			if !ok {
				return nil, fmt.Errorf("expected Float64 array for column %d", colIdx)
			}
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
			srcArray, ok := column.(*array.String)
			if !ok {
				return nil, fmt.Errorf("expected String array for column %d", colIdx)
			}
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

		case arrow.BOOL:
			srcArray, ok := column.(*array.Boolean)
			if !ok {
				return nil, fmt.Errorf("expected Boolean array for column %d", colIdx)
			}
			builder := array.NewBooleanBuilder(pool)

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

// SortKey represents a sorting specification for multi-column sorts
type SortKey struct {
	Column    string
	Ascending bool
}

// Sort returns a new DataFrame sorted by the specified column.
func (df *DataFrame) Sort(columnName string, ascending bool) (*DataFrame, error) {
	return df.SortMultiple([]SortKey{{Column: columnName, Ascending: ascending}})
}

// SortMultiple returns a new DataFrame sorted by multiple columns in the specified order.
func (df *DataFrame) SortMultiple(sortKeys []SortKey) (*DataFrame, error) {
	if len(sortKeys) == 0 {
		return nil, fmt.Errorf("no sort keys provided")
	}

	// Validate all columns exist
	schema := df.record.Schema()
	for _, key := range sortKeys {
		found := false
		for _, field := range schema.Fields() {
			if field.Name == key.Column {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("column not found: %s", key.Column)
		}
	}

	// Get the number of rows
	numRows := int(df.NumRows())
	if numRows == 0 {
		return df.Clone(), nil
	}

	// Create row indices
	indices := make([]int, numRows)
	for i := 0; i < numRows; i++ {
		indices[i] = i
	}

	// Sort indices based on the column values
	err := df.sortIndices(indices, sortKeys)
	if err != nil {
		return nil, fmt.Errorf("failed to sort indices: %w", err)
	}

	// Create new arrays with sorted data
	newColumns := make([]arrow.Array, df.NumCols())
	pool := memory.NewGoAllocator()

	for colIdx, field := range schema.Fields() {
		column := df.record.Column(colIdx)

		switch field.Type.ID() {
		case arrow.INT64:
			srcArray, ok := column.(*array.Int64)
			if !ok {
				return nil, fmt.Errorf("expected Int64 array for column %d", colIdx)
			}
			builder := array.NewInt64Builder(pool)
			defer builder.Release()

			for _, idx := range indices {
				if srcArray.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(srcArray.Value(idx))
				}
			}
			newColumns[colIdx] = builder.NewArray()

		case arrow.FLOAT64:
			srcArray, ok := column.(*array.Float64)
			if !ok {
				return nil, fmt.Errorf("expected Float64 array for column %d", colIdx)
			}
			builder := array.NewFloat64Builder(pool)
			defer builder.Release()

			for _, idx := range indices {
				if srcArray.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(srcArray.Value(idx))
				}
			}
			newColumns[colIdx] = builder.NewArray()

		case arrow.STRING:
			srcArray, ok := column.(*array.String)
			if !ok {
				return nil, fmt.Errorf("expected String array for column %d", colIdx)
			}
			builder := array.NewStringBuilder(pool)
			defer builder.Release()

			for _, idx := range indices {
				if srcArray.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(srcArray.Value(idx))
				}
			}
			newColumns[colIdx] = builder.NewArray()

		case arrow.BOOL:
			srcArray, ok := column.(*array.Boolean)
			if !ok {
				return nil, fmt.Errorf("expected Boolean array for column %d", colIdx)
			}
			builder := array.NewBooleanBuilder(pool)
			defer builder.Release()

			for _, idx := range indices {
				if srcArray.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(srcArray.Value(idx))
				}
			}
			newColumns[colIdx] = builder.NewArray()

		default:
			return nil, fmt.Errorf("unsupported data type for sorting: %s", field.Type)
		}
	}

	// Create new record with sorted data
	sortedRecord := array.NewRecord(schema, newColumns, df.record.NumRows())
	return NewDataFrame(sortedRecord), nil
}

// sortIndices sorts the indices array based on the specified sort keys using a stable sort
func (df *DataFrame) sortIndices(indices []int, sortKeys []SortKey) error {
	schema := df.record.Schema()

	// Get column indices for sort keys
	sortColumnIndices := make([]int, len(sortKeys))
	for i, key := range sortKeys {
		found := false
		for j, field := range schema.Fields() {
			if field.Name == key.Column {
				sortColumnIndices[i] = j
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("column not found: %s", key.Column)
		}
	}

	// Custom comparison function for sorting
	less := func(i, j int) bool {
		idxI, idxJ := indices[i], indices[j]

		for k, key := range sortKeys {
			colIdx := sortColumnIndices[k]
			column := df.record.Column(colIdx)
			field := schema.Field(colIdx)

			cmp := df.compareValues(column, field.Type, idxI, idxJ)
			if cmp != 0 {
				if key.Ascending {
					return cmp < 0
				} else {
					return cmp > 0
				}
			}
		}
		return false // Equal values
	}

	// Use a stable sort to maintain relative order for equal values
	df.stableSort(indices, less)
	return nil
}

// compareValues compares two values at given indices in a column
// Returns -1 if left < right, 0 if equal, 1 if left > right
func (df *DataFrame) compareValues(column arrow.Array, dataType arrow.DataType, leftIdx, rightIdx int) int {
	// Handle null values (nulls sort last)
	leftNull := column.IsNull(leftIdx)
	rightNull := column.IsNull(rightIdx)

	if leftNull && rightNull {
		return 0
	}
	if leftNull {
		return 1 // null sorts after non-null
	}
	if rightNull {
		return -1 // non-null sorts before null
	}

	switch dataType.ID() {
	case arrow.INT64:
		arr, ok := column.(*array.Int64)
		if !ok {
			return 0
		}
		left, right := arr.Value(leftIdx), arr.Value(rightIdx)
		if left < right {
			return -1
		} else if left > right {
			return 1
		}
		return 0

	case arrow.FLOAT64:
		arr, ok := column.(*array.Float64)
		if !ok {
			return 0
		}
		left, right := arr.Value(leftIdx), arr.Value(rightIdx)
		if left < right {
			return -1
		} else if left > right {
			return 1
		}
		return 0

	case arrow.STRING:
		arr, ok := column.(*array.String)
		if !ok {
			return 0
		}
		left, right := arr.Value(leftIdx), arr.Value(rightIdx)
		if left < right {
			return -1
		} else if left > right {
			return 1
		}
		return 0

	case arrow.BOOL:
		arr, ok := column.(*array.Boolean)
		if !ok {
			return 0
		}
		left, right := arr.Value(leftIdx), arr.Value(rightIdx)
		// false < true
		if !left && right {
			return -1
		} else if left && !right {
			return 1
		}
		return 0

	default:
		return 0 // Unsupported types are considered equal
	}
}

// stableSort implements a stable sort algorithm (insertion sort for simplicity)
func (df *DataFrame) stableSort(indices []int, less func(i, j int) bool) {
	for i := 1; i < len(indices); i++ {
		for j := i; j > 0 && less(j, j-1); j-- {
			indices[j], indices[j-1] = indices[j-1], indices[j]
		}
	}
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

// JoinType represents the type of join operation
type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
)

// Join performs a join operation between this DataFrame and another DataFrame.
// Uses hash join algorithm for efficient processing.
func (df *DataFrame) Join(other *DataFrame, leftKey, rightKey string, joinType JoinType) (*DataFrame, error) {
	if other == nil {
		return nil, fmt.Errorf("other DataFrame cannot be nil")
	}

	// Validate join keys exist
	if !df.HasColumn(leftKey) {
		return nil, fmt.Errorf("left join key column not found: %s", leftKey)
	}
	if !other.HasColumn(rightKey) {
		return nil, fmt.Errorf("right join key column not found: %s", rightKey)
	}

	// Get join key arrays
	leftKeyArray := df.record.Column(df.getColumnIndex(leftKey))
	rightKeyArray := other.record.Column(other.getColumnIndex(rightKey))

	// Perform the join based on type
	switch joinType {
	case InnerJoin:
		return df.performInnerJoin(other, leftKey, rightKey, leftKeyArray, rightKeyArray)
	case LeftJoin:
		return df.performLeftJoin(other, leftKey, rightKey, leftKeyArray, rightKeyArray)
	default:
		return nil, fmt.Errorf("unsupported join type: %d", joinType)
	}
}

// InnerJoin is a convenience method for inner joins
func (df *DataFrame) InnerJoin(other *DataFrame, leftKey, rightKey string) (*DataFrame, error) {
	return df.Join(other, leftKey, rightKey, InnerJoin)
}

// LeftJoin is a convenience method for left joins
func (df *DataFrame) LeftJoin(other *DataFrame, leftKey, rightKey string) (*DataFrame, error) {
	return df.Join(other, leftKey, rightKey, LeftJoin)
}

// getColumnIndex returns the index of a column by name
func (df *DataFrame) getColumnIndex(columnName string) int {
	schema := df.record.Schema()
	for i, field := range schema.Fields() {
		if field.Name == columnName {
			return i
		}
	}
	return -1
}

// performInnerJoin implements the inner join logic
func (df *DataFrame) performInnerJoin(other *DataFrame, leftKey, rightKey string, leftKeyArray, rightKeyArray arrow.Array) (*DataFrame, error) {
	// Build hash map from right DataFrame for efficient lookups
	rightHashMap := make(map[interface{}][]int)

	// Populate hash map with right side values
	for i := 0; i < rightKeyArray.Len(); i++ {
		if rightKeyArray.IsNull(i) {
			continue // Skip null values in joins
		}

		key := extractValue(rightKeyArray, i)
		if key != nil {
			rightHashMap[key] = append(rightHashMap[key], i)
		}
	}

	// Find matching rows
	var leftIndices, rightIndices []int
	for i := 0; i < leftKeyArray.Len(); i++ {
		if leftKeyArray.IsNull(i) {
			continue // Skip null values
		}

		leftValue := extractValue(leftKeyArray, i)
		if leftValue == nil {
			continue
		}

		if rightRows, exists := rightHashMap[leftValue]; exists {
			for _, rightRow := range rightRows {
				leftIndices = append(leftIndices, i)
				rightIndices = append(rightIndices, rightRow)
			}
		}
	}

	return df.buildJoinResult(other, leftKey, rightKey, leftIndices, rightIndices, false)
}

// performLeftJoin implements the left join logic
func (df *DataFrame) performLeftJoin(other *DataFrame, leftKey, rightKey string, leftKeyArray, rightKeyArray arrow.Array) (*DataFrame, error) {
	// Build hash map from right DataFrame
	rightHashMap := make(map[interface{}][]int)

	for i := 0; i < rightKeyArray.Len(); i++ {
		if rightKeyArray.IsNull(i) {
			continue
		}

		key := extractValue(rightKeyArray, i)
		if key != nil {
			rightHashMap[key] = append(rightHashMap[key], i)
		}
	}

	// Find matching rows, including unmatched left rows
	var leftIndices, rightIndices []int
	for i := 0; i < leftKeyArray.Len(); i++ {
		if leftKeyArray.IsNull(i) {
			// Include left row with null values for right side
			leftIndices = append(leftIndices, i)
			rightIndices = append(rightIndices, -1) // -1 indicates no match
			continue
		}

		leftValue := extractValue(leftKeyArray, i)
		if leftValue == nil {
			leftIndices = append(leftIndices, i)
			rightIndices = append(rightIndices, -1)
			continue
		}

		if rightRows, exists := rightHashMap[leftValue]; exists {
			for _, rightRow := range rightRows {
				leftIndices = append(leftIndices, i)
				rightIndices = append(rightIndices, rightRow)
			}
		} else {
			// No match found, include left row with nulls for right side
			leftIndices = append(leftIndices, i)
			rightIndices = append(rightIndices, -1)
		}
	}

	return df.buildJoinResult(other, leftKey, rightKey, leftIndices, rightIndices, true)
}

// extractValue extracts a comparable value from an Arrow array at given index
func extractValue(arr arrow.Array, index int) interface{} {
	switch typedArr := arr.(type) {
	case *array.Int64:
		return typedArr.Value(index)
	case *array.Float64:
		return typedArr.Value(index)
	case *array.String:
		return typedArr.Value(index)
	case *array.Boolean:
		return typedArr.Value(index)
	default:
		// For other types, convert to string representation
		return fmt.Sprintf("%v", arr.GetOneForMarshal(index))
	}
}

// buildJoinResult constructs the final joined DataFrame
func (df *DataFrame) buildJoinResult(other *DataFrame, leftKey, rightKey string, leftIndices, rightIndices []int, isLeftJoin bool) (*DataFrame, error) {
	if len(leftIndices) != len(rightIndices) {
		return nil, fmt.Errorf("internal error: index arrays length mismatch")
	}

	leftSchema := df.record.Schema()
	rightSchema := other.record.Schema()

	// Build result schema, avoiding duplicate column names
	var resultFields []arrow.Field
	columnNameMap := make(map[string]bool)

	// Add all left columns
	for _, field := range leftSchema.Fields() {
		resultFields = append(resultFields, field)
		columnNameMap[field.Name] = true
	}

	// Add right columns, skipping the join key and handling name conflicts
	for _, field := range rightSchema.Fields() {
		if field.Name == rightKey {
			continue // Skip right join key to avoid duplication
		}

		fieldName := field.Name
		if columnNameMap[fieldName] {
			fieldName = "right_" + fieldName // Prefix to avoid conflicts
		}

		resultFields = append(resultFields, arrow.Field{
			Name: fieldName,
			Type: field.Type,
		})
		columnNameMap[fieldName] = true
	}

	resultSchema := arrow.NewSchema(resultFields, nil)

	// Build result arrays
	resultArrays := make([]arrow.Array, len(resultFields))
	pool := memory.NewGoAllocator()

	fieldIndex := 0

	// Process left columns
	for i, field := range leftSchema.Fields() {
		resultArrays[fieldIndex] = df.buildJoinedArray(pool, df.record.Column(i), leftIndices, field.Type, false, int(df.record.NumRows()))
		fieldIndex++
	}

	// Process right columns (excluding join key)
	for i, field := range rightSchema.Fields() {
		if field.Name == rightKey {
			continue
		}

		resultArrays[fieldIndex] = df.buildJoinedArray(pool, other.record.Column(i), rightIndices, field.Type, isLeftJoin, int(other.record.NumRows()))
		fieldIndex++
	}

	// Create result record
	resultRecord := array.NewRecord(resultSchema, resultArrays, int64(len(leftIndices)))

	// Release arrays as they're now owned by the record
	for _, arr := range resultArrays {
		arr.Release()
	}

	return NewDataFrame(resultRecord), nil
}

// buildJoinedArray creates an array for the join result based on the provided indices
func (df *DataFrame) buildJoinedArray(pool memory.Allocator, sourceArray arrow.Array, indices []int, dataType arrow.DataType, allowNulls bool, sourceLength int) arrow.Array {
	switch dataType.ID() {
	case arrow.INT64:
		builder := array.NewInt64Builder(pool)
		defer builder.Release()

		sourceTyped := sourceArray.(*array.Int64)
		for _, idx := range indices {
			if idx == -1 && allowNulls {
				builder.AppendNull()
			} else if idx >= 0 && idx < sourceLength {
				if sourceTyped.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(sourceTyped.Value(idx))
				}
			} else {
				builder.AppendNull()
			}
		}
		return builder.NewArray()

	case arrow.FLOAT64:
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

		sourceTyped := sourceArray.(*array.Float64)
		for _, idx := range indices {
			if idx == -1 && allowNulls {
				builder.AppendNull()
			} else if idx >= 0 && idx < sourceLength {
				if sourceTyped.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(sourceTyped.Value(idx))
				}
			} else {
				builder.AppendNull()
			}
		}
		return builder.NewArray()

	case arrow.STRING:
		builder := array.NewStringBuilder(pool)
		defer builder.Release()

		sourceTyped := sourceArray.(*array.String)
		for _, idx := range indices {
			if idx == -1 && allowNulls {
				builder.AppendNull()
			} else if idx >= 0 && idx < sourceLength {
				if sourceTyped.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(sourceTyped.Value(idx))
				}
			} else {
				builder.AppendNull()
			}
		}
		return builder.NewArray()

	case arrow.BOOL:
		builder := array.NewBooleanBuilder(pool)
		defer builder.Release()

		sourceTyped := sourceArray.(*array.Boolean)
		for _, idx := range indices {
			if idx == -1 && allowNulls {
				builder.AppendNull()
			} else if idx >= 0 && idx < sourceLength {
				if sourceTyped.IsNull(idx) {
					builder.AppendNull()
				} else {
					builder.Append(sourceTyped.Value(idx))
				}
			} else {
				builder.AppendNull()
			}
		}
		return builder.NewArray()

	default:
		// For unsupported types, create a null array
		builder := array.NewNullBuilder(pool)
		defer builder.Release()

		for range indices {
			builder.AppendNull()
		}
		return builder.NewArray()
	}
}
