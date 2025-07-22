// Package gopherframe provides a production-first DataFrame library for Go.
// Built on Apache Arrow, it delivers high-performance data manipulation
// with zero-copy operations and seamless interoperability.
package gopherframe

import (
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

// Release releases the underlying Arrow resources.
// The DataFrame should not be used after calling Release.
func (df *DataFrame) Release() {
	if df.coreDF != nil {
		df.coreDF.Release()
		df.coreDF = nil
	}
}
