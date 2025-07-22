// Package gopherframe provides GroupBy and aggregation functionality.
package gopherframe

import (
	"fmt"
	"sort"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// GroupedDataFrame represents a DataFrame grouped by one or more columns.
type GroupedDataFrame struct {
	df          *DataFrame
	groupByCols []string
	err         error
}

// Aggregation represents an aggregation operation on a column.
type Aggregation struct {
	column    string
	operation string
	alias     string
}

// GroupBy groups the DataFrame by the specified columns.
func (df *DataFrame) GroupBy(columns ...string) *GroupedDataFrame {
	if df.err != nil {
		return &GroupedDataFrame{err: df.err}
	}

	if len(columns) == 0 {
		return &GroupedDataFrame{err: fmt.Errorf("no columns specified for groupby")}
	}

	// Validate all columns exist
	for _, col := range columns {
		if !df.HasColumn(col) {
			return &GroupedDataFrame{err: fmt.Errorf("column not found: %s", col)}
		}
	}

	return &GroupedDataFrame{
		df:          df,
		groupByCols: columns,
	}
}

// Agg performs the specified aggregations on the grouped data.
func (gdf *GroupedDataFrame) Agg(aggregations ...Aggregation) *DataFrame {
	if gdf.err != nil {
		return &DataFrame{err: gdf.err}
	}

	if len(aggregations) == 0 {
		return &DataFrame{err: fmt.Errorf("no aggregations specified")}
	}

	// For now, implement a simple groupby for single column with basic aggregations
	result, err := gdf.performGroupBy(aggregations)
	if err != nil {
		return &DataFrame{err: err}
	}

	return result
}

// performGroupBy executes the actual groupby logic
func (gdf *GroupedDataFrame) performGroupBy(aggregations []Aggregation) (*DataFrame, error) {
	// Handle both single and multi-column groupby
	if len(gdf.groupByCols) == 1 {
		return gdf.performSingleColumnGroupBy(aggregations)
	}
	return gdf.performMultiColumnGroupBy(aggregations)
}

// performSingleColumnGroupBy handles single column grouping
func (gdf *GroupedDataFrame) performSingleColumnGroupBy(aggregations []Aggregation) (*DataFrame, error) {
	groupCol := gdf.groupByCols[0]

	// Get the grouping column
	groupSeries, err := gdf.df.coreDF.Column(groupCol)
	if err != nil {
		return nil, fmt.Errorf("failed to get group column: %w", err)
	}
	defer groupSeries.Release()

	// Extract unique groups and build group indices
	groups, groupIndices, err := gdf.extractGroups(groupSeries)
	if err != nil {
		return nil, fmt.Errorf("failed to extract groups: %w", err)
	}

	// Build result columns
	// First column: group keys
	resultFields := []arrow.Field{{Name: groupCol, Type: groupSeries.DataType()}}
	resultColumns := []arrow.Array{groups}

	// Add aggregation columns
	for _, agg := range aggregations {
		aggField, aggColumn, err := gdf.performAggregation(agg, groupIndices)
		if err != nil {
			return nil, fmt.Errorf("failed to perform aggregation %s: %w", agg.Name(), err)
		}

		resultFields = append(resultFields, aggField)
		resultColumns = append(resultColumns, aggColumn)
	}

	// Create result schema and record
	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(groups.Len()))

	return NewDataFrame(resultRecord), nil
}

// extractGroups finds unique values and their indices
func (gdf *GroupedDataFrame) extractGroups(groupSeries *core.Series) (arrow.Array, map[string][]int, error) {
	// Build map of group value to row indices
	groupMap := make(map[string][]int)

	for i := 0; i < groupSeries.Len(); i++ {
		if groupSeries.IsNull(i) {
			continue // Skip null values for now
		}

		value, err := groupSeries.GetString(i)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get string value: %w", err)
		}

		groupMap[value] = append(groupMap[value], i)
	}

	// Sort group keys for consistent output
	var groupKeys []string
	for key := range groupMap {
		groupKeys = append(groupKeys, key)
	}
	sort.Strings(groupKeys)

	// Build result array of group keys
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	// Reorganize indices by sorted keys
	sortedIndices := make(map[string][]int)
	for _, key := range groupKeys {
		builder.Append(key)
		sortedIndices[key] = groupMap[key]
	}

	groupArray := builder.NewArray()
	return groupArray, sortedIndices, nil
}

// performAggregation executes a single aggregation
func (gdf *GroupedDataFrame) performAggregation(agg Aggregation, groupIndices map[string][]int) (arrow.Field, arrow.Array, error) {
	// Get the column to aggregate
	aggSeries, err := gdf.df.coreDF.Column(agg.column)
	if err != nil {
		return arrow.Field{}, nil, fmt.Errorf("failed to get aggregation column: %w", err)
	}
	defer aggSeries.Release()

	pool := memory.NewGoAllocator()

	switch agg.operation {
	case "sum":
		return gdf.performSum(aggSeries, groupIndices, agg.Name(), pool)
	case "mean":
		return gdf.performMean(aggSeries, groupIndices, agg.Name(), pool)
	case "count":
		return gdf.performCount(aggSeries, groupIndices, agg.Name(), pool)
	case "min":
		return gdf.performMin(aggSeries, groupIndices, agg.Name(), pool)
	case "max":
		return gdf.performMax(aggSeries, groupIndices, agg.Name(), pool)
	default:
		return arrow.Field{}, nil, fmt.Errorf("unsupported aggregation: %s", agg.operation)
	}
}

// performSum calculates sum for each group
func (gdf *GroupedDataFrame) performSum(series *core.Series, groupIndices map[string][]int, name string, pool memory.Allocator) (arrow.Field, arrow.Array, error) {
	if series.DataType().ID() != arrow.FLOAT64 {
		return arrow.Field{}, nil, fmt.Errorf("sum aggregation only supports float64, got %s", series.DataType())
	}

	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	// Sort keys to maintain consistent order
	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		sum := 0.0

		for _, idx := range indices {
			if !series.IsNull(idx) {
				val, err := series.GetFloat64(idx)
				if err != nil {
					return arrow.Field{}, nil, fmt.Errorf("failed to get float64 value: %w", err)
				}
				sum += val
			}
		}

		builder.Append(sum)
	}

	field := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performMean calculates mean for each group
func (gdf *GroupedDataFrame) performMean(series *core.Series, groupIndices map[string][]int, name string, pool memory.Allocator) (arrow.Field, arrow.Array, error) {
	// Simplified implementation - reuse sum logic and divide by count
	_, sumArray, err := gdf.performSum(series, groupIndices, name, pool)
	if err != nil {
		return arrow.Field{}, nil, err
	}
	defer sumArray.Release()

	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	sumFloat := sumArray.(*array.Float64)

	// Sort keys to maintain consistent order
	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	i := 0
	for _, key := range keys {
		indices := groupIndices[key]
		count := 0

		// Count non-null values
		for _, idx := range indices {
			if !series.IsNull(idx) {
				count++
			}
		}

		if count > 0 {
			mean := sumFloat.Value(i) / float64(count)
			builder.Append(mean)
		} else {
			builder.AppendNull()
		}
		i++
	}

	field := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performCount counts non-null values for each group
func (gdf *GroupedDataFrame) performCount(series *core.Series, groupIndices map[string][]int, name string, pool memory.Allocator) (arrow.Field, arrow.Array, error) {
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	// Sort keys to maintain consistent order
	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		count := int64(0)

		for _, idx := range indices {
			if !series.IsNull(idx) {
				count++
			}
		}

		builder.Append(count)
	}

	field := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int64}
	return field, builder.NewArray(), nil
}

// performMin finds minimum value for each group
func (gdf *GroupedDataFrame) performMin(series *core.Series, groupIndices map[string][]int, name string, pool memory.Allocator) (arrow.Field, arrow.Array, error) {
	if series.DataType().ID() != arrow.FLOAT64 {
		return arrow.Field{}, nil, fmt.Errorf("min aggregation only supports float64, got %s", series.DataType())
	}

	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	// Sort keys to maintain consistent order
	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		var min *float64

		for _, idx := range indices {
			if !series.IsNull(idx) {
				val, err := series.GetFloat64(idx)
				if err != nil {
					return arrow.Field{}, nil, fmt.Errorf("failed to get float64 value: %w", err)
				}

				if min == nil || val < *min {
					min = &val
				}
			}
		}

		if min != nil {
			builder.Append(*min)
		} else {
			builder.AppendNull()
		}
	}

	field := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performMax finds maximum value for each group
func (gdf *GroupedDataFrame) performMax(series *core.Series, groupIndices map[string][]int, name string, pool memory.Allocator) (arrow.Field, arrow.Array, error) {
	if series.DataType().ID() != arrow.FLOAT64 {
		return arrow.Field{}, nil, fmt.Errorf("max aggregation only supports float64, got %s", series.DataType())
	}

	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	// Sort keys to maintain consistent order
	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		var max *float64

		for _, idx := range indices {
			if !series.IsNull(idx) {
				val, err := series.GetFloat64(idx)
				if err != nil {
					return arrow.Field{}, nil, fmt.Errorf("failed to get float64 value: %w", err)
				}

				if max == nil || val > *max {
					max = &val
				}
			}
		}

		if max != nil {
			builder.Append(*max)
		} else {
			builder.AppendNull()
		}
	}

	field := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// Aggregation builders

// Sum creates a sum aggregation.
func Sum(column string) Aggregation {
	return Aggregation{
		column:    column,
		operation: "sum",
		alias:     column + "_sum",
	}
}

// Mean creates a mean/average aggregation.
func Mean(column string) Aggregation {
	return Aggregation{
		column:    column,
		operation: "mean",
		alias:     column + "_mean",
	}
}

// Count creates a count aggregation.
func Count(column string) Aggregation {
	return Aggregation{
		column:    column,
		operation: "count",
		alias:     column + "_count",
	}
}

// Min creates a minimum aggregation.
func Min(column string) Aggregation {
	return Aggregation{
		column:    column,
		operation: "min",
		alias:     column + "_min",
	}
}

// Max creates a maximum aggregation.
func Max(column string) Aggregation {
	return Aggregation{
		column:    column,
		operation: "max",
		alias:     column + "_max",
	}
}

// As sets a custom name for the aggregation result.
func (a Aggregation) As(alias string) Aggregation {
	a.alias = alias
	return a
}

// Name returns the name of the aggregation result column.
func (a Aggregation) Name() string {
	return a.alias
}

// performMultiColumnGroupBy handles multi-column grouping
func (gdf *GroupedDataFrame) performMultiColumnGroupBy(aggregations []Aggregation) (*DataFrame, error) {
	// Get all grouping columns
	var groupSeries []*core.Series
	for _, col := range gdf.groupByCols {
		series, err := gdf.df.coreDF.Column(col)
		if err != nil {
			return nil, fmt.Errorf("failed to get group column %s: %w", col, err)
		}
		groupSeries = append(groupSeries, series)
	}
	defer func() {
		for _, series := range groupSeries {
			series.Release()
		}
	}()

	// Build composite keys and group indices
	groupKeys, groupIndices, err := gdf.extractMultiColumnGroups(groupSeries)
	if err != nil {
		return nil, fmt.Errorf("failed to extract multi-column groups: %w", err)
	}

	// Build result columns - start with group columns
	var resultFields []arrow.Field
	var resultColumns []arrow.Array

	// Add group key columns to result
	for i, col := range gdf.groupByCols {
		resultFields = append(resultFields, arrow.Field{Name: col, Type: groupSeries[i].DataType()})
	}

	// Create arrays for each group column
	pool := memory.NewGoAllocator()
	for i, col := range gdf.groupByCols {
		builder := createBuilderForType(groupSeries[i].DataType(), pool)
		defer builder.Release()

		for _, key := range groupKeys {
			if len(groupIndices[key]) > 0 {
				firstIdx := groupIndices[key][0] // Use first row in group for group key
				if err := appendValueFromSeries(builder, groupSeries[i], firstIdx); err != nil {
					return nil, fmt.Errorf("failed to append group value for column %s: %w", col, err)
				}
			}
		}
		resultColumns = append(resultColumns, builder.NewArray())
	}

	// Add aggregation columns
	for _, agg := range aggregations {
		aggField, aggColumn, err := gdf.performAggregation(agg, groupIndices)
		if err != nil {
			return nil, fmt.Errorf("failed to perform aggregation %s: %w", agg.Name(), err)
		}
		resultFields = append(resultFields, aggField)
		resultColumns = append(resultColumns, aggColumn)
	}

	// Create result schema and record
	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(len(groupKeys)))

	return NewDataFrame(resultRecord), nil
}

// extractMultiColumnGroups creates composite keys from multiple columns
func (gdf *GroupedDataFrame) extractMultiColumnGroups(groupSeries []*core.Series) ([]string, map[string][]int, error) {
	if len(groupSeries) == 0 {
		return nil, nil, fmt.Errorf("no group series provided")
	}

	numRows := groupSeries[0].Len()
	groupMap := make(map[string][]int)

	// Build composite keys by concatenating column values
	for i := 0; i < numRows; i++ {
		var keyParts []string
		skipRow := false

		for _, series := range groupSeries {
			if series.IsNull(i) {
				skipRow = true
				break
			}

			value, err := series.GetString(i)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to get string value: %w", err)
			}
			keyParts = append(keyParts, value)
		}

		if !skipRow {
			key := fmt.Sprintf("%v", keyParts) // Simple string representation of composite key
			groupMap[key] = append(groupMap[key], i)
		}
	}

	// Sort group keys for consistent output
	var groupKeys []string
	for key := range groupMap {
		groupKeys = append(groupKeys, key)
	}
	sort.Strings(groupKeys)

	return groupKeys, groupMap, nil
}

// createBuilderForType creates appropriate builder for Arrow data type
func createBuilderForType(dataType arrow.DataType, pool memory.Allocator) array.Builder {
	switch dataType.ID() {
	case arrow.STRING:
		return array.NewStringBuilder(pool)
	case arrow.INT64:
		return array.NewInt64Builder(pool)
	case arrow.FLOAT64:
		return array.NewFloat64Builder(pool)
	case arrow.BOOL:
		return array.NewBooleanBuilder(pool)
	default:
		return array.NewStringBuilder(pool) // Fallback to string
	}
}

// appendValueFromSeries appends value from series to appropriate builder
func appendValueFromSeries(builder array.Builder, series *core.Series, index int) error {
	if series.IsNull(index) {
		builder.AppendNull()
		return nil
	}

	switch b := builder.(type) {
	case *array.StringBuilder:
		val, err := series.GetString(index)
		if err != nil {
			return err
		}
		b.Append(val)
	case *array.Int64Builder:
		val, err := series.GetInt64(index)
		if err != nil {
			return err
		}
		b.Append(val)
	case *array.Float64Builder:
		val, err := series.GetFloat64(index)
		if err != nil {
			return err
		}
		b.Append(val)
	case *array.BooleanBuilder:
		val, err := series.GetBool(index)
		if err != nil {
			return err
		}
		b.Append(val)
	default:
		return fmt.Errorf("unsupported builder type")
	}
	return nil
}
