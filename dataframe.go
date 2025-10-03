// Package gopherframe provides a production-first DataFrame library for Go.
// Built on Apache Arrow, it delivers high-performance data manipulation
// with zero-copy operations and seamless interoperability.
package gopherframe

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
)

// DataFrame is the public-facing DataFrame that provides a chainable,
// user-friendly API for data manipulation operations.
type DataFrame struct {
	coreDF *core.DataFrame
	err    error
}

// NewDataFrame creates a new public DataFrame from an Arrow Record.
func NewDataFrame(record arrow.Record) *DataFrame {
	coreDF := core.NewDataFrame(record)
	return &DataFrame{
		coreDF: coreDF,
	}
}

// NumRows returns the number of rows in the DataFrame.
func (df *DataFrame) NumRows() int64 {
	if df.err != nil || df.coreDF == nil {
		return 0
	}

	return df.coreDF.NumRows()
}

// NumCols returns the number of columns in the DataFrame.
func (df *DataFrame) NumCols() int64 {
	if df.err != nil || df.coreDF == nil {
		return 0
	}
	return df.coreDF.NumCols()
}

// ColumnNames returns the names of all columns.
func (df *DataFrame) ColumnNames() []string {
	if df.err != nil || df.coreDF == nil {
		return nil
	}
	return df.coreDF.ColumnNames()
}

// HasColumn returns true if a column with the given name exists.
func (df *DataFrame) HasColumn(name string) bool {
	if df.err != nil || df.coreDF == nil {
		return false
	}
	return df.coreDF.HasColumn(name)
}

// Schema returns the Arrow schema of the DataFrame.
func (df *DataFrame) Schema() *arrow.Schema {
	if df.err != nil || df.coreDF == nil {
		return nil
	}
	return df.coreDF.Schema()
}

// Err returns any accumulated error from chained operations.
func (df *DataFrame) Err() error {
	return df.err
}

// Filter returns a new DataFrame containing only rows that match the predicate.
// Uses lazy evaluation - the predicate expression is built but not executed
// until the result is materialized.
func (df *DataFrame) Filter(predicate expr.Expr) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	// Evaluate the predicate expression to get a boolean array
	predicateArray, err := predicate.Evaluate(df.coreDF)
	if err != nil {
		return &DataFrame{err: err}
	}
	defer predicateArray.Release()

	// Use the core DataFrame's Filter method
	filteredCoreDF, err := df.coreDF.Filter(predicateArray)
	if err != nil {
		return &DataFrame{err: err}
	}

	return &DataFrame{coreDF: filteredCoreDF}
}

// Select returns a new DataFrame with only the specified columns.
func (df *DataFrame) Select(columnNames ...string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	// Use the core DataFrame's Select method
	newCoreDF, err := df.coreDF.Select(columnNames)
	if err != nil {
		return &DataFrame{err: err}
	}

	return &DataFrame{coreDF: newCoreDF}
}

// WithColumn returns a new DataFrame with an additional or replaced column.
func (df *DataFrame) WithColumn(name string, expression expr.Expr) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	// Evaluate the expression to get the new column
	newColumnArray, err := expression.Evaluate(df.coreDF)
	if err != nil {
		return &DataFrame{err: err}
	}

	// Use the core DataFrame's WithColumn method
	newCoreDF, err := df.coreDF.WithColumn(name, newColumnArray)
	if err != nil {
		newColumnArray.Release()
		return &DataFrame{err: err}
	}

	newColumnArray.Release() // Core DataFrame took ownership
	return &DataFrame{coreDF: newCoreDF}
}

// Col creates a column expression for use in operations like Filter and WithColumn.
// Example: df.Col("age") returns an expression representing the "age" column
func (df *DataFrame) Col(name string) expr.Expr {
	return expr.Col(name)
}

// Record returns the underlying Arrow Record.
func (df *DataFrame) Record() arrow.Record {
	if df.err != nil || df.coreDF == nil {
		return nil
	}
	return df.coreDF.Record()
}

// Release releases the underlying Arrow resources.
// The DataFrame should not be used after calling Release.
func (df *DataFrame) Release() {
	if df.coreDF != nil {
		df.coreDF.Release()
		df.coreDF = nil
	}
}

// SortKey represents a sorting specification for multi-column sorts
type SortKey struct {
	Column    string
	Ascending bool
}

// Sort returns a new DataFrame sorted by the specified column.
func (df *DataFrame) Sort(columnName string, ascending bool) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	sortedCoreDF, err := df.coreDF.Sort(columnName, ascending)
	if err != nil {
		return &DataFrame{err: err}
	}

	return &DataFrame{coreDF: sortedCoreDF}
}

// SortMultiple returns a new DataFrame sorted by multiple columns in the specified order.
func (df *DataFrame) SortMultiple(sortKeys []SortKey) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	// Convert public SortKey to core.SortKey
	coreSortKeys := make([]core.SortKey, len(sortKeys))
	for i, key := range sortKeys {
		coreSortKeys[i] = core.SortKey{
			Column:    key.Column,
			Ascending: key.Ascending,
		}
	}

	sortedCoreDF, err := df.coreDF.SortMultiple(coreSortKeys)
	if err != nil {
		return &DataFrame{err: err}
	}

	return &DataFrame{coreDF: sortedCoreDF}
}

// By creates a sort specification for a column.
// Example: By("age", true) for ascending, By("name", false) for descending
func By(column string, ascending bool) SortKey {
	return SortKey{
		Column:    column,
		Ascending: ascending,
	}
}

// InnerJoin performs an inner join with another DataFrame.
// Returns rows that have matching values in both DataFrames.
// Example: df.InnerJoin(other, "user_id", "id")
func (df *DataFrame) InnerJoin(other *DataFrame, leftKey, rightKey string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if other == nil {
		return &DataFrame{err: fmt.Errorf("other DataFrame cannot be nil")}
	}
	if other.err != nil {
		return &DataFrame{err: other.err}
	}

	joinedCoreDF, err := df.coreDF.InnerJoin(other.coreDF, leftKey, rightKey)
	if err != nil {
		return &DataFrame{err: err}
	}

	return &DataFrame{coreDF: joinedCoreDF}
}

// LeftJoin performs a left join with another DataFrame.
// Returns all rows from the left DataFrame, and matching rows from the right DataFrame.
// Non-matching rows from the right will have null values.
// Example: df.LeftJoin(other, "user_id", "id")
func (df *DataFrame) LeftJoin(other *DataFrame, leftKey, rightKey string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if other == nil {
		return &DataFrame{err: fmt.Errorf("other DataFrame cannot be nil")}
	}
	if other.err != nil {
		return &DataFrame{err: other.err}
	}

	joinedCoreDF, err := df.coreDF.LeftJoin(other.coreDF, leftKey, rightKey)
	if err != nil {
		return &DataFrame{err: err}
	}

	return &DataFrame{coreDF: joinedCoreDF}
}
