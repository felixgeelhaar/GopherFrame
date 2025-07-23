// Package aggregation contains domain models and logic for data aggregation operations.
// This package encapsulates the business rules around grouping and summarizing data.
package aggregation

import (
	"fmt"
	"sort"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

// AggregationType represents the type of aggregation to perform.
type AggregationType int

const (
	Sum AggregationType = iota
	Mean
	Count
	Min
	Max
)

// AggregationSpec specifies an aggregation operation.
type AggregationSpec struct {
	Column string
	Type   AggregationType
	Alias  string
}

// GroupByRequest represents a request to group data by columns and perform aggregations.
type GroupByRequest struct {
	GroupColumns []string
	Aggregations []AggregationSpec
}

// GroupByResult represents the result of a group-by operation.
type GroupByResult struct {
	DataFrame *dataframe.DataFrame
	Error     error
}

// GroupByService is a domain service that handles grouping and aggregation logic.
type GroupByService struct {
	allocator memory.Allocator
}

// NewGroupByService creates a new GroupByService.
func NewGroupByService() *GroupByService {
	return &GroupByService{
		allocator: memory.DefaultAllocator,
	}
}

// Execute performs the group-by operation according to the specification.
func (s *GroupByService) Execute(df *dataframe.DataFrame, request GroupByRequest) GroupByResult {
	if len(request.GroupColumns) == 0 {
		return GroupByResult{Error: fmt.Errorf("no group columns specified")}
	}

	if len(request.Aggregations) == 0 {
		return GroupByResult{Error: fmt.Errorf("no aggregations specified")}
	}

	if len(request.GroupColumns) == 1 {
		result, err := s.performSingleColumnGroupBy(df, request)
		return GroupByResult{DataFrame: result, Error: err}
	} else {
		result, err := s.performMultiColumnGroupBy(df, request)
		return GroupByResult{DataFrame: result, Error: err}
	}
}

// performSingleColumnGroupBy handles single-column group-by operations.
func (s *GroupByService) performSingleColumnGroupBy(df *dataframe.DataFrame, request GroupByRequest) (*dataframe.DataFrame, error) {
	groupCol := request.GroupColumns[0]

	// Validate group column exists
	if !df.HasColumn(groupCol) {
		return nil, fmt.Errorf("group column not found: %s", groupCol)
	}

	// Extract group column data
	record := df.Record()
	schema := record.Schema()

	var groupColIndex int = -1
	for i, field := range schema.Fields() {
		if field.Name == groupCol {
			groupColIndex = i
			break
		}
	}

	if groupColIndex == -1 {
		return nil, fmt.Errorf("group column not found: %s", groupCol)
	}

	groupArray := record.Column(groupColIndex)

	// Build group mapping
	groups, groupIndices, err := s.extractGroups(groupArray)
	if err != nil {
		return nil, fmt.Errorf("failed to extract groups: %w", err)
	}
	defer groups.Release()

	// Build result schema and columns
	resultFields := []arrow.Field{{Name: groupCol, Type: groupArray.DataType()}}
	resultColumns := []arrow.Array{groups}

	// Perform aggregations
	for _, agg := range request.Aggregations {
		field, column, err := s.performAggregation(record, agg, groupIndices)
		if err != nil {
			return nil, fmt.Errorf("failed to perform aggregation %s: %w", agg.Alias, err)
		}
		defer column.Release()

		resultFields = append(resultFields, field)
		resultColumns = append(resultColumns, column)
	}

	// Create result DataFrame
	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(groups.Len()))

	return dataframe.NewDataFrame(resultRecord), nil
}

// extractGroups finds unique values in the group column and builds group indices.
func (s *GroupByService) extractGroups(groupArray arrow.Array) (arrow.Array, map[string][]int, error) {
	groupMap := make(map[string][]int)

	// Extract group values (simplified for string columns)
	stringArray, ok := groupArray.(*array.String)
	if !ok {
		return nil, nil, fmt.Errorf("only string group columns are currently supported")
	}

	for i := 0; i < stringArray.Len(); i++ {
		if !stringArray.IsNull(i) {
			value := stringArray.Value(i)
			groupMap[value] = append(groupMap[value], i)
		}
	}

	// Sort group keys for consistent output
	var groupKeys []string
	for key := range groupMap {
		groupKeys = append(groupKeys, key)
	}
	sort.Strings(groupKeys)

	// Build result array of group keys
	builder := array.NewStringBuilder(s.allocator)
	defer builder.Release()

	for _, key := range groupKeys {
		builder.Append(key)
	}

	groupKeysArray := builder.NewArray()
	return groupKeysArray, groupMap, nil
}

// performAggregation executes a single aggregation operation.
func (s *GroupByService) performAggregation(record arrow.Record, agg AggregationSpec, groupIndices map[string][]int) (arrow.Field, arrow.Array, error) {
	// Find the column to aggregate
	var aggColIndex int = -1
	schema := record.Schema()
	for i, field := range schema.Fields() {
		if field.Name == agg.Column {
			aggColIndex = i
			break
		}
	}

	if aggColIndex == -1 {
		return arrow.Field{}, nil, fmt.Errorf("aggregation column not found: %s", agg.Column)
	}

	aggArray := record.Column(aggColIndex)

	switch agg.Type {
	case Sum:
		return s.performSum(aggArray, groupIndices, agg.Alias)
	case Mean:
		return s.performMean(aggArray, groupIndices, agg.Alias)
	case Count:
		return s.performCount(aggArray, groupIndices, agg.Alias)
	case Min:
		return s.performMin(aggArray, groupIndices, agg.Alias)
	case Max:
		return s.performMax(aggArray, groupIndices, agg.Alias)
	default:
		return arrow.Field{}, nil, fmt.Errorf("unsupported aggregation type: %d", agg.Type)
	}
}

// performSum calculates sum for each group.
func (s *GroupByService) performSum(arr arrow.Array, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("sum aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	// Sort keys for consistent output
	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		sum := 0.0

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				sum += floatArray.Value(idx)
			}
		}

		builder.Append(sum)
	}

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performMean calculates mean for each group.
func (s *GroupByService) performMean(arr arrow.Array, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("mean aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		sum := 0.0
		count := 0

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				sum += floatArray.Value(idx)
				count++
			}
		}

		if count > 0 {
			builder.Append(sum / float64(count))
		} else {
			builder.AppendNull()
		}
	}

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performCount counts non-null values for each group.
func (s *GroupByService) performCount(arr arrow.Array, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	builder := array.NewInt64Builder(s.allocator)
	defer builder.Release()

	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		count := int64(0)

		for _, idx := range indices {
			if !arr.IsNull(idx) {
				count++
			}
		}

		builder.Append(count)
	}

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Int64}
	return field, builder.NewArray(), nil
}

// performMin finds minimum value for each group.
func (s *GroupByService) performMin(arr arrow.Array, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("min aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		var min *float64

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				val := floatArray.Value(idx)
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

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performMax finds maximum value for each group.
func (s *GroupByService) performMax(arr arrow.Array, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("max aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	var keys []string
	for key := range groupIndices {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		indices := groupIndices[key]
		var max *float64

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				val := floatArray.Value(idx)
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

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// performMultiColumnGroupBy handles multi-column group-by operations.
func (s *GroupByService) performMultiColumnGroupBy(df *dataframe.DataFrame, request GroupByRequest) (*dataframe.DataFrame, error) {
	record := df.Record()
	schema := record.Schema()

	// Validate all group columns exist
	groupColIndices := make([]int, len(request.GroupColumns))
	for i, groupCol := range request.GroupColumns {
		colIndex := -1
		for j, field := range schema.Fields() {
			if field.Name == groupCol {
				colIndex = j
				break
			}
		}
		if colIndex == -1 {
			return nil, fmt.Errorf("group column not found: %s", groupCol)
		}
		groupColIndices[i] = colIndex
	}

	// Extract multi-column groups
	groupKeys, groupIndices, err := s.extractMultiColumnGroups(record, request.GroupColumns, groupColIndices)
	if err != nil {
		return nil, fmt.Errorf("failed to extract multi-column groups: %w", err)
	}

	// Build result schema - start with group columns
	resultFields := make([]arrow.Field, len(request.GroupColumns))
	resultColumns := make([]arrow.Array, len(request.GroupColumns))

	for i, groupCol := range request.GroupColumns {
		colIndex := groupColIndices[i]
		field := schema.Field(colIndex)
		resultFields[i] = arrow.Field{Name: groupCol, Type: field.Type}

		// Build array of group key values for this column
		column, err := s.buildGroupKeyColumn(groupKeys, i, field.Type)
		if err != nil {
			return nil, fmt.Errorf("failed to build group key column %s: %w", groupCol, err)
		}
		resultColumns[i] = column
	}

	// Perform aggregations
	for _, agg := range request.Aggregations {
		field, column, err := s.performMultiColumnAggregation(record, agg, groupKeys, groupIndices)
		if err != nil {
			return nil, fmt.Errorf("failed to perform aggregation %s: %w", agg.Alias, err)
		}
		defer column.Release()

		resultFields = append(resultFields, field)
		resultColumns = append(resultColumns, column)
	}

	// Create result DataFrame
	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(len(groupKeys)))

	return dataframe.NewDataFrame(resultRecord), nil
}

// extractMultiColumnGroups finds unique combinations of values across multiple columns.
func (s *GroupByService) extractMultiColumnGroups(record arrow.Record, groupColumns []string, groupColIndices []int) ([][]string, map[string][]int, error) {
	groupMap := make(map[string][]int)
	numRows := int(record.NumRows())

	// Handle empty record
	if numRows == 0 {
		return [][]string{}, groupMap, nil
	}

	// Extract values from all group columns (simplified for string columns only)
	groupArrays := make([]*array.String, len(groupColIndices))
	for i, colIndex := range groupColIndices {
		stringArray, ok := record.Column(colIndex).(*array.String)
		if !ok {
			return nil, nil, fmt.Errorf("only string group columns are currently supported for column %s", groupColumns[i])
		}
		groupArrays[i] = stringArray
	}

	// Create composite keys for each row
	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		// Check if any group column value is null
		hasNull := false
		for _, arr := range groupArrays {
			if arr.IsNull(rowIdx) {
				hasNull = true
				break
			}
		}

		if !hasNull {
			// Build composite key by concatenating values with a separator
			keyParts := make([]string, len(groupArrays))
			for i, arr := range groupArrays {
				keyParts[i] = arr.Value(rowIdx)
			}

			// Use a separator that won't appear in data
			compositeKey := joinWithSeparator(keyParts, "|||")
			groupMap[compositeKey] = append(groupMap[compositeKey], rowIdx)
		}
	}

	// Sort group keys for consistent output and split back into individual values
	var compositeKeys []string
	for key := range groupMap {
		compositeKeys = append(compositeKeys, key)
	}
	sort.Strings(compositeKeys)

	// Convert back to individual group key values
	groupKeys := make([][]string, len(compositeKeys))
	for i, compositeKey := range compositeKeys {
		groupKeys[i] = splitBySeparator(compositeKey, "|||")
	}

	return groupKeys, groupMap, nil
}

// buildGroupKeyColumn creates an Arrow array for a specific group column.
func (s *GroupByService) buildGroupKeyColumn(groupKeys [][]string, columnIndex int, dataType arrow.DataType) (arrow.Array, error) {
	// For now, only support string columns
	if dataType.ID() != arrow.STRING {
		return nil, fmt.Errorf("only string group columns are currently supported")
	}

	builder := array.NewStringBuilder(s.allocator)
	defer builder.Release()

	for _, keyGroup := range groupKeys {
		if columnIndex < len(keyGroup) {
			builder.Append(keyGroup[columnIndex])
		} else {
			builder.AppendNull()
		}
	}

	return builder.NewArray(), nil
}

// performMultiColumnAggregation executes aggregation for multi-column groups.
func (s *GroupByService) performMultiColumnAggregation(record arrow.Record, agg AggregationSpec, groupKeys [][]string, groupIndices map[string][]int) (arrow.Field, arrow.Array, error) {
	// Find the column to aggregate
	var aggColIndex int = -1
	schema := record.Schema()
	for i, field := range schema.Fields() {
		if field.Name == agg.Column {
			aggColIndex = i
			break
		}
	}

	if aggColIndex == -1 {
		return arrow.Field{}, nil, fmt.Errorf("aggregation column not found: %s", agg.Column)
	}

	aggArray := record.Column(aggColIndex)

	// Convert group keys back to composite keys for lookup
	orderedGroupMap := make(map[string][]int)
	for _, keyGroup := range groupKeys {
		compositeKey := joinWithSeparator(keyGroup, "|||")
		if indices, exists := groupIndices[compositeKey]; exists {
			orderedGroupMap[compositeKey] = indices
		}
	}

	switch agg.Type {
	case Sum:
		return s.performSumMultiColumn(aggArray, groupKeys, orderedGroupMap, agg.Alias)
	case Mean:
		return s.performMeanMultiColumn(aggArray, groupKeys, orderedGroupMap, agg.Alias)
	case Count:
		return s.performCountMultiColumn(aggArray, groupKeys, orderedGroupMap, agg.Alias)
	case Min:
		return s.performMinMultiColumn(aggArray, groupKeys, orderedGroupMap, agg.Alias)
	case Max:
		return s.performMaxMultiColumn(aggArray, groupKeys, orderedGroupMap, agg.Alias)
	default:
		return arrow.Field{}, nil, fmt.Errorf("unsupported aggregation type: %d", agg.Type)
	}
}

// Multi-column aggregation helper functions
func (s *GroupByService) performSumMultiColumn(arr arrow.Array, groupKeys [][]string, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("sum aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	for _, keyGroup := range groupKeys {
		compositeKey := joinWithSeparator(keyGroup, "|||")
		indices := groupIndices[compositeKey]
		sum := 0.0

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				sum += floatArray.Value(idx)
			}
		}

		builder.Append(sum)
	}

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

func (s *GroupByService) performMeanMultiColumn(arr arrow.Array, groupKeys [][]string, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("mean aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	for _, keyGroup := range groupKeys {
		compositeKey := joinWithSeparator(keyGroup, "|||")
		indices := groupIndices[compositeKey]
		sum := 0.0
		count := 0

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				sum += floatArray.Value(idx)
				count++
			}
		}

		if count > 0 {
			builder.Append(sum / float64(count))
		} else {
			builder.AppendNull()
		}
	}

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

func (s *GroupByService) performCountMultiColumn(arr arrow.Array, groupKeys [][]string, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	builder := array.NewInt64Builder(s.allocator)
	defer builder.Release()

	for _, keyGroup := range groupKeys {
		compositeKey := joinWithSeparator(keyGroup, "|||")
		indices := groupIndices[compositeKey]
		count := int64(0)

		for _, idx := range indices {
			if !arr.IsNull(idx) {
				count++
			}
		}

		builder.Append(count)
	}

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Int64}
	return field, builder.NewArray(), nil
}

func (s *GroupByService) performMinMultiColumn(arr arrow.Array, groupKeys [][]string, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("min aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	for _, keyGroup := range groupKeys {
		compositeKey := joinWithSeparator(keyGroup, "|||")
		indices := groupIndices[compositeKey]
		var min *float64

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				val := floatArray.Value(idx)
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

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

func (s *GroupByService) performMaxMultiColumn(arr arrow.Array, groupKeys [][]string, groupIndices map[string][]int, alias string) (arrow.Field, arrow.Array, error) {
	floatArray, ok := arr.(*array.Float64)
	if !ok {
		return arrow.Field{}, nil, fmt.Errorf("max aggregation only supports float64 columns")
	}

	builder := array.NewFloat64Builder(s.allocator)
	defer builder.Release()

	for _, keyGroup := range groupKeys {
		compositeKey := joinWithSeparator(keyGroup, "|||")
		indices := groupIndices[compositeKey]
		var max *float64

		for _, idx := range indices {
			if !floatArray.IsNull(idx) {
				val := floatArray.Value(idx)
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

	field := arrow.Field{Name: alias, Type: arrow.PrimitiveTypes.Float64}
	return field, builder.NewArray(), nil
}

// Helper functions for string manipulation
func joinWithSeparator(parts []string, separator string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += separator + parts[i]
	}
	return result
}

func splitBySeparator(str, separator string) []string {
	if str == "" {
		return []string{}
	}

	parts := []string{}
	start := 0
	for i := 0; i <= len(str)-len(separator); i++ {
		if str[i:i+len(separator)] == separator {
			parts = append(parts, str[start:i])
			start = i + len(separator)
			i += len(separator) - 1 // Skip the separator
		}
	}
	parts = append(parts, str[start:]) // Add the last part
	return parts
}
