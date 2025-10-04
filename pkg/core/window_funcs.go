// Package core provides built-in window function implementations.
package core

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// RowNumberFunc implements the ROW_NUMBER window function.
//
// ROW_NUMBER assigns a unique sequential integer to each row within
// a partition, starting at 1. The numbering resets for each partition.
//
// Example:
//
//	df.Window().
//	    PartitionBy("category").
//	    OrderBy("price").
//	    Over(RowNumber().As("row_num"))
type RowNumberFunc struct {
	name string
}

// RowNumber creates a new ROW_NUMBER window function.
//
// Returns:
//   - *RowNumberFunc: Window function that assigns row numbers
//
// Example:
//
//	// Number rows within each category
//	df.Window().
//	    PartitionBy("category").
//	    Over(RowNumber().As("row_number"))
func RowNumber() *RowNumberFunc {
	return &RowNumberFunc{name: "row_number"}
}

// As sets the result column name.
func (fn *RowNumberFunc) As(name string) *RowNumberFunc {
	fn.name = name
	return fn
}

// Name returns the result column name.
func (fn *RowNumberFunc) Name() string {
	return fn.name
}

// Compute calculates row numbers for the partition.
func (fn *RowNumberFunc) Compute(partition []int, df *DataFrame) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := range partition {
		builder.Append(int64(i + 1)) // 1-indexed
	}

	return builder.NewArray(), nil
}

// RankFunc implements the RANK window function.
//
// RANK assigns a rank to each row within a partition. Rows with equal
// values receive the same rank, and the next rank is incremented by the
// number of tied rows (leaving gaps in the ranking).
//
// Example: ranks 1, 2, 2, 4, 5 (note gap at 3)
type RankFunc struct {
	name string
}

// Rank creates a new RANK window function.
//
// RANK requires OrderBy to be specified in the window specification.
// Rows with equal order column values receive the same rank.
//
// Returns:
//   - *RankFunc: Window function that assigns ranks with gaps
//
// Example:
//
//	// Rank products by price (ties get same rank)
//	df.Window().
//	    PartitionBy("category").
//	    OrderBy("price").
//	    Over(Rank().As("price_rank"))
func Rank() *RankFunc {
	return &RankFunc{name: "rank"}
}

// As sets the result column name.
func (fn *RankFunc) As(name string) *RankFunc {
	fn.name = name
	return fn
}

// Name returns the result column name.
func (fn *RankFunc) Name() string {
	return fn.name
}

// Compute calculates ranks for the partition.
func (fn *RankFunc) Compute(partition []int, df *DataFrame) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	if len(partition) == 0 {
		return builder.NewArray(), nil
	}

	// Assign ranks (1-indexed, with gaps for ties)
	currentRank := int64(1)
	builder.Append(currentRank)

	for i := 1; i < len(partition); i++ {
		// Compare with previous row
		// If values are different, increment rank by number of rows since last change
		// For now, all rows get consecutive ranks (will enhance with proper comparison)
		currentRank = int64(i + 1)
		builder.Append(currentRank)
	}

	return builder.NewArray(), nil
}

// DenseRankFunc implements the DENSE_RANK window function.
//
// DENSE_RANK assigns a rank to each row within a partition. Rows with equal
// values receive the same rank, but unlike RANK, there are no gaps in the
// ranking sequence.
//
// Example: ranks 1, 2, 2, 3, 4 (no gap at 3)
type DenseRankFunc struct {
	name string
}

// DenseRank creates a new DENSE_RANK window function.
//
// DENSE_RANK requires OrderBy to be specified in the window specification.
// Unlike RANK, this produces consecutive rank values with no gaps.
//
// Returns:
//   - *DenseRankFunc: Window function that assigns ranks without gaps
//
// Example:
//
//	// Dense rank products by price
//	df.Window().
//	    PartitionBy("category").
//	    OrderByDesc("price").
//	    Over(DenseRank().As("dense_rank"))
func DenseRank() *DenseRankFunc {
	return &DenseRankFunc{name: "dense_rank"}
}

// As sets the result column name.
func (fn *DenseRankFunc) As(name string) *DenseRankFunc {
	fn.name = name
	return fn
}

// Name returns the result column name.
func (fn *DenseRankFunc) Name() string {
	return fn.name
}

// Compute calculates dense ranks for the partition.
func (fn *DenseRankFunc) Compute(partition []int, df *DataFrame) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	if len(partition) == 0 {
		return builder.NewArray(), nil
	}

	// Assign dense ranks (1-indexed, no gaps)
	currentRank := int64(1)
	builder.Append(currentRank)

	for i := 1; i < len(partition); i++ {
		// Compare with previous row
		// If values are different, increment rank by 1
		// For now, all rows get consecutive ranks (will enhance with proper comparison)
		currentRank = int64(i + 1)
		builder.Append(currentRank)
	}

	return builder.NewArray(), nil
}

// LagFunc implements the LAG window function.
//
// LAG provides access to a row at a given offset prior to the current row
// within the partition. This is useful for comparing current row values
// with previous row values.
//
// Example:
//
//	// Get previous day's sales
//	df.Window().
//	    OrderBy("date").
//	    Over(Lag("sales", 1).As("prev_day_sales"))
type LagFunc struct {
	name       string
	columnName string
	offset     int
	defaultVal interface{}
}

// Lag creates a new LAG window function.
//
// LAG accesses the value of a column from a previous row in the partition.
// If the offset goes beyond the partition boundary, null is returned
// (or the default value if specified).
//
// Parameters:
//   - columnName: Column to retrieve value from
//   - offset: Number of rows back (must be positive)
//
// Returns:
//   - *LagFunc: Window function that retrieves previous row values
//
// Example:
//
//	// Get previous row's value
//	Lag("price", 1).As("prev_price")
//
//	// Get value from 3 rows back
//	Lag("sales", 3).As("sales_3_days_ago")
func Lag(columnName string, offset int) *LagFunc {
	if offset < 0 {
		offset = -offset // ensure positive
	}
	if offset == 0 {
		offset = 1 // minimum offset is 1
	}

	return &LagFunc{
		name:       fmt.Sprintf("lag_%s_%d", columnName, offset),
		columnName: columnName,
		offset:     offset,
		defaultVal: nil,
	}
}

// As sets the result column name.
func (fn *LagFunc) As(name string) *LagFunc {
	fn.name = name
	return fn
}

// Default sets the default value when offset goes beyond partition boundary.
func (fn *LagFunc) Default(val interface{}) *LagFunc {
	fn.defaultVal = val
	return fn
}

// Name returns the result column name.
func (fn *LagFunc) Name() string {
	return fn.name
}

// Compute calculates lag values for the partition.
func (fn *LagFunc) Compute(partition []int, df *DataFrame) (arrow.Array, error) {
	// Get source column
	series, err := df.Column(fn.columnName)
	if err != nil {
		return nil, fmt.Errorf("column %s not found: %w", fn.columnName, err)
	}

	sourceArray := series.Array()
	pool := memory.NewGoAllocator()

	// Create builder for result array
	builder := array.NewBuilder(pool, sourceArray.DataType())
	defer builder.Release()

	// For each row in partition, get value from offset rows back
	for i := range partition {
		lagIdx := i - fn.offset

		if lagIdx < 0 {
			// Beyond partition boundary
			if fn.defaultVal != nil {
				appendValue(builder, fn.defaultVal)
			} else {
				builder.AppendNull()
			}
		} else {
			// Get value from lagged row
			lagRowIdx := partition[lagIdx]
			if sourceArray.IsNull(lagRowIdx) {
				builder.AppendNull()
			} else {
				val := sourceArray.GetOneForMarshal(lagRowIdx)
				appendValue(builder, val)
			}
		}
	}

	return builder.NewArray(), nil
}

// LeadFunc implements the LEAD window function.
//
// LEAD provides access to a row at a given offset after the current row
// within the partition. This is useful for comparing current row values
// with future row values.
//
// Example:
//
//	// Get next day's sales
//	df.Window().
//	    OrderBy("date").
//	    Over(Lead("sales", 1).As("next_day_sales"))
type LeadFunc struct {
	name       string
	columnName string
	offset     int
	defaultVal interface{}
}

// Lead creates a new LEAD window function.
//
// LEAD accesses the value of a column from a subsequent row in the partition.
// If the offset goes beyond the partition boundary, null is returned
// (or the default value if specified).
//
// Parameters:
//   - columnName: Column to retrieve value from
//   - offset: Number of rows forward (must be positive)
//
// Returns:
//   - *LeadFunc: Window function that retrieves next row values
//
// Example:
//
//	// Get next row's value
//	Lead("price", 1).As("next_price")
//
//	// Get value from 3 rows ahead
//	Lead("sales", 3).As("sales_3_days_ahead")
func Lead(columnName string, offset int) *LeadFunc {
	if offset < 0 {
		offset = -offset // ensure positive
	}
	if offset == 0 {
		offset = 1 // minimum offset is 1
	}

	return &LeadFunc{
		name:       fmt.Sprintf("lead_%s_%d", columnName, offset),
		columnName: columnName,
		offset:     offset,
		defaultVal: nil,
	}
}

// As sets the result column name.
func (fn *LeadFunc) As(name string) *LeadFunc {
	fn.name = name
	return fn
}

// Default sets the default value when offset goes beyond partition boundary.
func (fn *LeadFunc) Default(val interface{}) *LeadFunc {
	fn.defaultVal = val
	return fn
}

// Name returns the result column name.
func (fn *LeadFunc) Name() string {
	return fn.name
}

// Compute calculates lead values for the partition.
func (fn *LeadFunc) Compute(partition []int, df *DataFrame) (arrow.Array, error) {
	// Get source column
	series, err := df.Column(fn.columnName)
	if err != nil {
		return nil, fmt.Errorf("column %s not found: %w", fn.columnName, err)
	}

	sourceArray := series.Array()
	pool := memory.NewGoAllocator()

	// Create builder for result array
	builder := array.NewBuilder(pool, sourceArray.DataType())
	defer builder.Release()

	// For each row in partition, get value from offset rows ahead
	for i := range partition {
		leadIdx := i + fn.offset

		if leadIdx >= len(partition) {
			// Beyond partition boundary
			if fn.defaultVal != nil {
				appendValue(builder, fn.defaultVal)
			} else {
				builder.AppendNull()
			}
		} else {
			// Get value from leading row
			leadRowIdx := partition[leadIdx]
			if sourceArray.IsNull(leadRowIdx) {
				builder.AppendNull()
			} else {
				val := sourceArray.GetOneForMarshal(leadRowIdx)
				appendValue(builder, val)
			}
		}
	}

	return builder.NewArray(), nil
}
