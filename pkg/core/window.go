// Package core provides window function operations for DataFrames.
package core

import (
	"fmt"
	"sort"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// WindowSpec defines a window function specification.
//
// Window functions operate over a "window" of rows defined by partitioning
// and ordering criteria. This enables analytical operations like ranking,
// lag/lead calculations, and cumulative aggregations.
//
// Example:
//
//	df.Window().
//	  PartitionBy("category").
//	  OrderBy("date").
//	  Over(Lag("sales", 1), RowNumber())
type WindowSpec struct {
	df             *DataFrame
	partitionCols  []string
	orderCols      []string
	orderAscending []bool
}

// WindowFunc represents a window function to be applied.
//
// Window functions compute values based on a set of rows related to the
// current row, such as ranking, offset access (lag/lead), or aggregation.
type WindowFunc interface {
	// Name returns the result column name for this window function
	Name() string

	// Compute calculates the window function over the given partition
	Compute(partition []int, df *DataFrame) (arrow.Array, error)
}

// Window initiates a window function specification.
//
// Window functions allow you to perform calculations across a set of rows
// that are related to the current row. Common use cases include:
//   - Ranking rows within partitions
//   - Accessing previous/next row values (lag/lead)
//   - Computing running totals
//   - Calculating moving averages
//
// The window specification follows a builder pattern:
//  1. Optionally partition data with PartitionBy()
//  2. Optionally order data with OrderBy()
//  3. Apply window functions with Over()
//
// Returns:
//   - *WindowSpec: Window specification builder
//
// Example:
//
//	// Rank products by price within each category
//	result := df.Window().
//	    PartitionBy("category").
//	    OrderBy("price").
//	    Over(Rank().As("price_rank"))
//
//	// Calculate 7-day moving average of sales
//	result := df.Window().
//	    OrderBy("date").
//	    Over(
//	        Lag("sales", 1).As("prev_day_sales"),
//	        Lead("sales", 1).As("next_day_sales"),
//	    )
//
// Memory: All window functions return new arrays that must be released
func (df *DataFrame) Window() *WindowSpec {
	return &WindowSpec{
		df:             df,
		partitionCols:  []string{},
		orderCols:      []string{},
		orderAscending: []bool{},
	}
}

// PartitionBy specifies columns to partition the data.
//
// Partitioning divides the DataFrame into groups. Window functions are
// computed separately within each partition. For example, ranking is
// reset for each partition.
//
// Parameters:
//   - cols: Column names to partition by
//
// Returns:
//   - *WindowSpec: Updated window specification (builder pattern)
//
// Example:
//
//	// Rank within each category and region
//	df.Window().
//	    PartitionBy("category", "region").
//	    OrderBy("sales").
//	    Over(Rank())
//
// Note: Multiple PartitionBy calls replace previous partitions, not append
func (ws *WindowSpec) PartitionBy(cols ...string) *WindowSpec {
	ws.partitionCols = cols
	return ws
}

// OrderBy specifies columns to order the data within partitions.
//
// Ordering determines the sequence in which window functions process rows.
// This is critical for functions like Lag, Lead, Rank, and RowNumber.
// By default, ordering is ascending.
//
// Parameters:
//   - cols: Column names to order by (ascending)
//
// Returns:
//   - *WindowSpec: Updated window specification (builder pattern)
//
// Example:
//
//	// Order by date ascending (oldest first)
//	df.Window().OrderBy("date").Over(Lag("value", 1))
//
//	// Order by multiple columns
//	df.Window().OrderBy("year", "month", "day").Over(RowNumber())
//
// Note: For descending order, use OrderByDesc()
// Multiple OrderBy calls replace previous ordering, not append
func (ws *WindowSpec) OrderBy(cols ...string) *WindowSpec {
	ws.orderCols = cols
	ws.orderAscending = make([]bool, len(cols))
	for i := range ws.orderAscending {
		ws.orderAscending[i] = true // ascending by default
	}
	return ws
}

// OrderByDesc specifies columns to order the data in descending order.
//
// Same as OrderBy but sorts in descending order (largest to smallest).
//
// Parameters:
//   - cols: Column names to order by (descending)
//
// Returns:
//   - *WindowSpec: Updated window specification (builder pattern)
//
// Example:
//
//	// Order by sales descending (highest first)
//	df.Window().
//	    PartitionBy("category").
//	    OrderByDesc("sales").
//	    Over(Rank().As("sales_rank"))
func (ws *WindowSpec) OrderByDesc(cols ...string) *WindowSpec {
	ws.orderCols = cols
	ws.orderAscending = make([]bool, len(cols))
	for i := range ws.orderAscending {
		ws.orderAscending[i] = false // descending
	}
	return ws
}

// Over applies window functions to the DataFrame.
//
// This is the terminal operation that executes the window computation and
// returns a new DataFrame with additional columns for each window function.
//
// Parameters:
//   - funcs: One or more window functions to apply
//
// Returns:
//   - *DataFrame: New DataFrame with original columns plus window function results
//   - error: Returns error if window computation fails
//
// Memory: Caller must call Release() on the returned DataFrame
//
// Example:
//
//	// Add multiple window function columns
//	result, err := df.Window().
//	    PartitionBy("category").
//	    OrderBy("date").
//	    Over(
//	        RowNumber().As("row_num"),
//	        Rank().As("rank"),
//	        Lag("value", 1).As("prev_value"),
//	    )
//	defer result.Release()
//
// Complexity: O(n log n) for sorting + O(n) for window computation
func (ws *WindowSpec) Over(funcs ...WindowFunc) (*DataFrame, error) {
	if len(funcs) == 0 {
		return nil, fmt.Errorf("at least one window function required")
	}

	// Get partitions based on partition columns
	partitions, err := ws.getPartitions()
	if err != nil {
		return nil, fmt.Errorf("failed to create partitions: %w", err)
	}

	// Sort each partition by order columns
	for i := range partitions {
		if err := ws.sortPartition(&partitions[i]); err != nil {
			return nil, fmt.Errorf("failed to sort partition %d: %w", i, err)
		}
	}

	// Compute each window function
	pool := ws.df.allocator
	var newColumns []arrow.Array
	var newFields []arrow.Field

	// Start with existing columns
	for i := 0; i < int(ws.df.NumCols()); i++ {
		newColumns = append(newColumns, ws.df.record.Column(i))
		newFields = append(newFields, ws.df.record.Schema().Field(i))
	}

	// Add window function result columns
	for _, fn := range funcs {
		resultArray, err := ws.computeWindowFunc(fn, partitions, pool)
		if err != nil {
			// Release any arrays we've created
			for _, arr := range newColumns[int(ws.df.NumCols()):] {
				arr.Release()
			}
			return nil, fmt.Errorf("failed to compute window function %s: %w", fn.Name(), err)
		}

		newColumns = append(newColumns, resultArray)

		// Determine field type from result array
		newFields = append(newFields, arrow.Field{
			Name:     fn.Name(),
			Type:     resultArray.DataType(),
			Nullable: true,
		})
	}

	// Create new record with window function columns
	newSchema := arrow.NewSchema(newFields, nil)
	newRecord := array.NewRecord(newSchema, newColumns, ws.df.record.NumRows())

	// Release window function result arrays (record retains them)
	for i := int(ws.df.NumCols()); i < len(newColumns); i++ {
		newColumns[i].Release()
	}

	return NewDataFrameWithAllocator(newRecord, pool), nil
}

// partition represents a group of row indices within the same partition
type partition struct {
	rows []int // row indices in this partition
}

// getPartitions divides the DataFrame into partitions based on partition columns.
func (ws *WindowSpec) getPartitions() ([]partition, error) {
	if len(ws.partitionCols) == 0 {
		// No partitioning - all rows in one partition
		allRows := make([]int, ws.df.record.NumRows())
		for i := range allRows {
			allRows[i] = i
		}
		return []partition{{rows: allRows}}, nil
	}

	// Build partition key for each row
	partitionMap := make(map[string][]int)

	for rowIdx := 0; rowIdx < int(ws.df.record.NumRows()); rowIdx++ {
		key, err := ws.getPartitionKey(rowIdx)
		if err != nil {
			return nil, err
		}
		partitionMap[key] = append(partitionMap[key], rowIdx)
	}

	// Convert map to slice of partitions
	partitions := make([]partition, 0, len(partitionMap))
	for _, rows := range partitionMap {
		partitions = append(partitions, partition{rows: rows})
	}

	return partitions, nil
}

// getPartitionKey creates a unique key string for a row based on partition columns.
func (ws *WindowSpec) getPartitionKey(rowIdx int) (string, error) {
	var key string
	for i, colName := range ws.partitionCols {
		series, err := ws.df.Column(colName)
		if err != nil {
			return "", err
		}

		arr := series.Array()
		if arr.IsNull(rowIdx) {
			key += "<NULL>"
		} else {
			key += fmt.Sprintf("%v", arr.GetOneForMarshal(rowIdx))
		}

		if i < len(ws.partitionCols)-1 {
			key += "|"
		}
	}
	return key, nil
}

// sortPartition sorts rows within a partition based on order columns.
func (ws *WindowSpec) sortPartition(p *partition) error {
	if len(ws.orderCols) == 0 {
		return nil // No ordering specified
	}

	// Get arrays for order columns
	orderArrays := make([]arrow.Array, len(ws.orderCols))
	for i, colName := range ws.orderCols {
		series, err := ws.df.Column(colName)
		if err != nil {
			return err
		}
		orderArrays[i] = series.Array()
	}

	// Sort partition rows
	sort.Slice(p.rows, func(i, j int) bool {
		rowI := p.rows[i]
		rowJ := p.rows[j]

		for colIdx, arr := range orderArrays {
			// Handle nulls (nulls sort last)
			nullI := arr.IsNull(rowI)
			nullJ := arr.IsNull(rowJ)

			if nullI && nullJ {
				continue
			}
			if nullI {
				return false // null sorts last
			}
			if nullJ {
				return true
			}

			// Compare values
			cmp := compareValues(arr, rowI, rowJ)
			if cmp == 0 {
				continue // equal, check next column
			}

			ascending := ws.orderAscending[colIdx]
			if ascending {
				return cmp < 0
			}
			return cmp > 0
		}

		return false // all columns equal
	})

	return nil
}

// compareValues compares two values in an Arrow array.
// Returns: -1 if a < b, 0 if a == b, 1 if a > b
func compareValues(arr arrow.Array, idxA, idxB int) int {
	switch typedArr := arr.(type) {
	case *array.Int64:
		a, b := typedArr.Value(idxA), typedArr.Value(idxB)
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	case *array.Float64:
		a, b := typedArr.Value(idxA), typedArr.Value(idxB)
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	case *array.String:
		a, b := typedArr.Value(idxA), typedArr.Value(idxB)
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	case *array.Boolean:
		a, b := typedArr.Value(idxA), typedArr.Value(idxB)
		if !a && b {
			return -1
		}
		if a && !b {
			return 1
		}
		return 0
	default:
		// For other types, use string comparison
		a := fmt.Sprintf("%v", arr.GetOneForMarshal(idxA))
		b := fmt.Sprintf("%v", arr.GetOneForMarshal(idxB))
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
}

// computeWindowFunc computes a window function across all partitions.
func (ws *WindowSpec) computeWindowFunc(fn WindowFunc, partitions []partition, pool memory.Allocator) (arrow.Array, error) {
	// Create result array builder based on function type
	// For now, we'll use a generic approach and determine type from first partition

	// Compute for first partition to determine result type
	if len(partitions) == 0 {
		return nil, fmt.Errorf("no partitions to compute")
	}

	firstResult, err := fn.Compute(partitions[0].rows, ws.df)
	if err != nil {
		return nil, err
	}
	defer firstResult.Release()

	resultType := firstResult.DataType()

	// Create builder for full result
	builder := array.NewBuilder(pool, resultType)
	defer builder.Release()

	// Build result array with proper ordering
	resultMap := make(map[int]interface{})

	// Compute for all partitions
	for _, part := range partitions {
		partResult, err := fn.Compute(part.rows, ws.df)
		if err != nil {
			return nil, err
		}

		// Map results back to original row indices
		for i, rowIdx := range part.rows {
			if partResult.IsNull(i) {
				resultMap[rowIdx] = nil
			} else {
				resultMap[rowIdx] = partResult.GetOneForMarshal(i)
			}
		}
		partResult.Release()
	}

	// Build final array in original row order
	for rowIdx := 0; rowIdx < int(ws.df.record.NumRows()); rowIdx++ {
		val := resultMap[rowIdx]
		if val == nil {
			builder.AppendNull()
		} else {
			appendValue(builder, val)
		}
	}

	return builder.NewArray(), nil
}

// appendValue appends a value to an array builder based on its type.
func appendValue(builder array.Builder, val interface{}) {
	switch b := builder.(type) {
	case *array.Int64Builder:
		b.Append(val.(int64))
	case *array.Float64Builder:
		b.Append(val.(float64))
	case *array.StringBuilder:
		b.Append(val.(string))
	case *array.BooleanBuilder:
		b.Append(val.(bool))
	default:
		// For other types, append null as fallback
		builder.AppendNull()
	}
}
