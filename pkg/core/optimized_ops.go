package core

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/compute"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// OptimizedOperations provides high-performance implementations of core DataFrame operations
type OptimizedOperations struct {
	pool memory.Allocator
}

// NewOptimizedOperations creates a new optimized operations context
func NewOptimizedOperations(pool memory.Allocator) *OptimizedOperations {
	if pool == nil {
		pool = memory.NewGoAllocator()
	}
	return &OptimizedOperations{pool: pool}
}

// FilterVectorized performs filtering using Arrow compute kernels for maximum performance
func (ops *OptimizedOperations) FilterVectorized(df *DataFrame, predicateArray arrow.Array) (*DataFrame, error) {
	ctx := context.Background()

	// Validate predicate
	if predicateArray.DataType().ID() != arrow.BOOL {
		return nil, fmt.Errorf("filter predicate must be boolean array, got %s", predicateArray.DataType())
	}

	if int64(predicateArray.Len()) != df.NumRows() {
		return nil, fmt.Errorf("predicate length %d does not match DataFrame rows %d", predicateArray.Len(), df.NumRows())
	}

	// Use Arrow compute filter kernel for vectorized operation
	schema := df.record.Schema()
	filteredColumns := make([]arrow.Array, len(schema.Fields()))

	for colIdx := range schema.Fields() {
		column := df.record.Column(colIdx)

		// Create datums for Arrow compute
		columnDatum := compute.NewDatum(column)
		predicateDatum := compute.NewDatum(predicateArray)

		// Use Arrow's filter compute kernel
		result, err := compute.CallFunction(ctx, "filter", nil, columnDatum, predicateDatum)
		if err != nil {
			// Clean up any already processed columns
			for i := 0; i < colIdx; i++ {
				if filteredColumns[i] != nil {
					filteredColumns[i].Release()
				}
			}
			return nil, fmt.Errorf("filter compute kernel failed on column %d: %w", colIdx, err)
		}

		// Extract result array
		switch resultDatum := result.(type) {
		case *compute.ArrayDatum:
			filteredArray := array.MakeFromData(resultDatum.Value)
			filteredArray.Retain()
			filteredColumns[colIdx] = filteredArray
		default:
			// Clean up
			for i := 0; i <= colIdx; i++ {
				if filteredColumns[i] != nil {
					filteredColumns[i].Release()
				}
			}
			return nil, fmt.Errorf("filter kernel returned unexpected result type for column %d", colIdx)
		}
	}

	// Count filtered rows
	var filteredRowCount int64
	if len(filteredColumns) > 0 {
		filteredRowCount = int64(filteredColumns[0].Len())
	}

	// Create new record
	filteredRecord := array.NewRecord(schema, filteredColumns, filteredRowCount)

	// Release filtered columns (record retains them)
	for _, arr := range filteredColumns {
		arr.Release()
	}

	return NewDataFrame(filteredRecord), nil
}

// FilterOptimized performs filtering with memory pool optimization and type-specific fast paths
func (ops *OptimizedOperations) FilterOptimized(df *DataFrame, predicateArray arrow.Array) (*DataFrame, error) {
	// Validate predicate
	if predicateArray.DataType().ID() != arrow.BOOL {
		return nil, fmt.Errorf("filter predicate must be boolean array, got %s", predicateArray.DataType())
	}

	if int64(predicateArray.Len()) != df.NumRows() {
		return nil, fmt.Errorf("predicate length %d does not match DataFrame rows %d", predicateArray.Len(), df.NumRows())
	}

	boolArray, ok := predicateArray.(*array.Boolean)
	if !ok {
		return nil, fmt.Errorf("failed to cast predicate to boolean array")
	}

	// Pre-compute indices for filtered rows to avoid multiple passes
	filteredIndices := make([]int, 0, predicateArray.Len())
	for i := 0; i < boolArray.Len(); i++ {
		if !boolArray.IsNull(i) && boolArray.Value(i) {
			filteredIndices = append(filteredIndices, i)
		}
	}

	trueCount := int64(len(filteredIndices))
	if trueCount == 0 {
		return ops.createEmptyDataFrame(df.record.Schema())
	}

	// Create filtered columns using optimized builders
	schema := df.record.Schema()
	filteredColumns := make([]arrow.Array, len(schema.Fields()))

	for colIdx, field := range schema.Fields() {
		column := df.record.Column(colIdx)

		switch field.Type.ID() {
		case arrow.INT64:
			filteredColumns[colIdx] = ops.filterInt64Column(column, filteredIndices)
		case arrow.FLOAT64:
			filteredColumns[colIdx] = ops.filterFloat64Column(column, filteredIndices)
		case arrow.STRING:
			filteredColumns[colIdx] = ops.filterStringColumn(column, filteredIndices)
		case arrow.BOOL:
			filteredColumns[colIdx] = ops.filterBooleanColumn(column, filteredIndices)
		default:
			// Clean up already processed columns
			for i := 0; i < colIdx; i++ {
				filteredColumns[i].Release()
			}
			return nil, fmt.Errorf("unsupported data type for optimized filtering: %s", field.Type)
		}
	}

	// Create new record
	filteredRecord := array.NewRecord(schema, filteredColumns, trueCount)

	// Release filtered columns (record retains them)
	for _, arr := range filteredColumns {
		arr.Release()
	}

	return NewDataFrame(filteredRecord), nil
}

// SelectOptimized performs column selection with zero-copy optimization
func (ops *OptimizedOperations) SelectOptimized(df *DataFrame, columnNames []string) (*DataFrame, error) {
	if len(columnNames) == 0 {
		return nil, fmt.Errorf("must specify at least one column")
	}

	schema := df.record.Schema()

	// Pre-validate all columns and build index map
	indices := make([]int, len(columnNames))
	fieldMap := make(map[string]int, len(schema.Fields()))

	// Build field map once
	for i, field := range schema.Fields() {
		fieldMap[field.Name] = i
	}

	// Validate all columns exist
	for i, columnName := range columnNames {
		idx, exists := fieldMap[columnName]
		if !exists {
			return nil, fmt.Errorf("column '%s' not found", columnName)
		}
		indices[i] = idx
	}

	// Create new schema with selected fields
	selectedFields := make([]arrow.Field, len(columnNames))
	selectedColumns := make([]arrow.Array, len(columnNames))

	for i, idx := range indices {
		selectedFields[i] = schema.Field(idx)
		column := df.record.Column(idx)
		column.Retain() // Zero-copy: just retain reference
		selectedColumns[i] = column
	}

	newSchema := arrow.NewSchema(selectedFields, nil)
	selectedRecord := array.NewRecord(newSchema, selectedColumns, df.NumRows())

	// Release columns (record retains them)
	for _, col := range selectedColumns {
		col.Release()
	}

	return NewDataFrame(selectedRecord), nil
}

// WithColumnOptimized performs column addition/replacement with memory pool optimization
func (ops *OptimizedOperations) WithColumnOptimized(df *DataFrame, columnName string, newColumn arrow.Array) (*DataFrame, error) {
	if int64(newColumn.Len()) != df.NumRows() {
		return nil, fmt.Errorf("new column length %d does not match DataFrame rows %d", newColumn.Len(), df.NumRows())
	}

	schema := df.record.Schema()
	existingColumnIndex := -1

	// Find if column already exists
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
				newFields[i] = arrow.Field{Name: columnName, Type: newColumn.DataType()}
				newColumn.Retain()
				newColumns[i] = newColumn
			} else {
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

		// Add new column at the end
		newFields[len(schema.Fields())] = arrow.Field{Name: columnName, Type: newColumn.DataType()}
		newColumn.Retain()
		newColumns[len(schema.Fields())] = newColumn
	}

	newSchema := arrow.NewSchema(newFields, nil)
	newRecord := array.NewRecord(newSchema, newColumns, df.NumRows())

	// Release columns (record retains them)
	for _, col := range newColumns {
		col.Release()
	}

	return NewDataFrame(newRecord), nil
}

// Helper methods for type-specific filtering

func (ops *OptimizedOperations) filterInt64Column(column arrow.Array, indices []int) arrow.Array {
	srcArray := column.(*array.Int64)
	builder := array.NewInt64Builder(ops.pool)
	defer builder.Release()

	builder.Reserve(len(indices))

	for _, idx := range indices {
		if srcArray.IsNull(idx) {
			builder.AppendNull()
		} else {
			builder.Append(srcArray.Value(idx))
		}
	}

	return builder.NewArray()
}

func (ops *OptimizedOperations) filterFloat64Column(column arrow.Array, indices []int) arrow.Array {
	srcArray := column.(*array.Float64)
	builder := array.NewFloat64Builder(ops.pool)
	defer builder.Release()

	builder.Reserve(len(indices))

	for _, idx := range indices {
		if srcArray.IsNull(idx) {
			builder.AppendNull()
		} else {
			builder.Append(srcArray.Value(idx))
		}
	}

	return builder.NewArray()
}

func (ops *OptimizedOperations) filterStringColumn(column arrow.Array, indices []int) arrow.Array {
	srcArray := column.(*array.String)
	builder := array.NewStringBuilder(ops.pool)
	defer builder.Release()

	builder.Reserve(len(indices))

	for _, idx := range indices {
		if srcArray.IsNull(idx) {
			builder.AppendNull()
		} else {
			builder.Append(srcArray.Value(idx))
		}
	}

	return builder.NewArray()
}

func (ops *OptimizedOperations) filterBooleanColumn(column arrow.Array, indices []int) arrow.Array {
	srcArray := column.(*array.Boolean)
	builder := array.NewBooleanBuilder(ops.pool)
	defer builder.Release()

	builder.Reserve(len(indices))

	for _, idx := range indices {
		if srcArray.IsNull(idx) {
			builder.AppendNull()
		} else {
			builder.Append(srcArray.Value(idx))
		}
	}

	return builder.NewArray()
}

func (ops *OptimizedOperations) createEmptyDataFrame(schema *arrow.Schema) (*DataFrame, error) {
	emptyColumns := make([]arrow.Array, len(schema.Fields()))

	for i, field := range schema.Fields() {
		switch field.Type.ID() {
		case arrow.INT64:
			builder := array.NewInt64Builder(ops.pool)
			emptyColumns[i] = builder.NewArray()
			builder.Release()
		case arrow.FLOAT64:
			builder := array.NewFloat64Builder(ops.pool)
			emptyColumns[i] = builder.NewArray()
			builder.Release()
		case arrow.STRING:
			builder := array.NewStringBuilder(ops.pool)
			emptyColumns[i] = builder.NewArray()
			builder.Release()
		case arrow.BOOL:
			builder := array.NewBooleanBuilder(ops.pool)
			emptyColumns[i] = builder.NewArray()
			builder.Release()
		default:
			return nil, fmt.Errorf("unsupported data type for empty DataFrame: %s", field.Type)
		}
	}

	emptyRecord := array.NewRecord(schema, emptyColumns, 0)

	// Release columns (record retains them)
	for _, col := range emptyColumns {
		col.Release()
	}

	return NewDataFrame(emptyRecord), nil
}

// Add optimized methods to DataFrame
func (df *DataFrame) FilterVectorized(predicateArray arrow.Array) (*DataFrame, error) {
	ops := NewOptimizedOperations(nil)
	return ops.FilterVectorized(df, predicateArray)
}

func (df *DataFrame) FilterOptimized(predicateArray arrow.Array) (*DataFrame, error) {
	ops := NewOptimizedOperations(nil)
	return ops.FilterOptimized(df, predicateArray)
}

func (df *DataFrame) FilterOptimizedWithPool(predicateArray arrow.Array, pool memory.Allocator) (*DataFrame, error) {
	ops := NewOptimizedOperations(pool)
	return ops.FilterOptimized(df, predicateArray)
}

func (df *DataFrame) SelectOptimized(columnNames []string) (*DataFrame, error) {
	ops := NewOptimizedOperations(nil)
	return ops.SelectOptimized(df, columnNames)
}

func (df *DataFrame) WithColumnOptimized(columnName string, newColumn arrow.Array) (*DataFrame, error) {
	ops := NewOptimizedOperations(nil)
	return ops.WithColumnOptimized(df, columnName, newColumn)
}

func (df *DataFrame) WithColumnOptimizedWithPool(columnName string, newColumn arrow.Array, pool memory.Allocator) (*DataFrame, error) {
	ops := NewOptimizedOperations(pool)
	return ops.WithColumnOptimized(df, columnName, newColumn)
}
