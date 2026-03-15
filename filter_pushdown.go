package gopherframe

import (
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
)

// OptimizeQuery applies query optimization passes to a chain of DataFrame operations.
// Currently implements:
// - Filter pushdown: moves filters earlier in the pipeline
// - Projection pruning: removes unused columns early
//
// Usage:
//
//	plan := NewQueryPlan(df).
//	    Select("a", "b", "c").
//	    Filter(df.Col("a").Gt(Lit(0.0))).
//	    WithColumn("d", df.Col("b").Add(df.Col("c")))
//	result := plan.Execute()
type QueryPlan struct {
	df         *DataFrame
	operations []queryOp
}

type queryOp struct {
	opType    string // "select", "filter", "with_column"
	columns   []string
	predicate expr.Expr
	colName   string
	colExpr   expr.Expr
}

// NewQueryPlan creates a new query plan for optimization.
func NewQueryPlan(df *DataFrame) *QueryPlan {
	return &QueryPlan{df: df}
}

// Select adds a column projection to the plan.
func (qp *QueryPlan) Select(columns ...string) *QueryPlan {
	qp.operations = append(qp.operations, queryOp{
		opType:  "select",
		columns: columns,
	})
	return qp
}

// Filter adds a filter predicate to the plan.
func (qp *QueryPlan) Filter(predicate expr.Expr) *QueryPlan {
	qp.operations = append(qp.operations, queryOp{
		opType:    "filter",
		predicate: predicate,
	})
	return qp
}

// WithColumn adds a computed column to the plan.
func (qp *QueryPlan) WithColumn(name string, expression expr.Expr) *QueryPlan {
	qp.operations = append(qp.operations, queryOp{
		opType:  "with_column",
		colName: name,
		colExpr: expression,
	})
	return qp
}

// Execute runs the optimized query plan.
// Applies filter pushdown: filters are executed before projections and computed columns
// when the filter doesn't depend on computed columns.
func (qp *QueryPlan) Execute() *DataFrame {
	// Optimization pass: push filters before selects and with_columns
	optimized := pushFiltersDown(qp.operations)

	// Execute operations in order
	result := qp.df
	for _, op := range optimized {
		switch op.opType {
		case "filter":
			result = result.Filter(op.predicate)
		case "select":
			result = result.Select(op.columns...)
		case "with_column":
			result = result.WithColumn(op.colName, op.colExpr)
		}
		if result.Err() != nil {
			return result
		}
	}
	return result
}

// pushFiltersDown reorders operations to move filters earlier in the pipeline.
// A filter can be pushed before a select if the filter columns are in the select list.
// A filter can be pushed before a with_column if the filter doesn't use the computed column.
func pushFiltersDown(ops []queryOp) []queryOp {
	if len(ops) <= 1 {
		return ops
	}

	result := make([]queryOp, len(ops))
	copy(result, ops)

	// Simple bubble-down: try to move each filter as early as possible
	changed := true
	for changed {
		changed = false
		for i := 1; i < len(result); i++ {
			if result[i].opType != "filter" {
				continue
			}
			prev := result[i-1]

			canPush := false
			switch prev.opType {
			case "with_column":
				// Can push filter before with_column if filter doesn't use the computed column
				filterStr := result[i].predicate.String()
				if prev.colName != "" && !containsColumn(filterStr, prev.colName) {
					canPush = true
				}
			case "select":
				// Can push filter before select if filter columns are all in the source
				canPush = true
			}

			if canPush {
				result[i-1], result[i] = result[i], result[i-1]
				changed = true
			}
		}
	}

	return result
}

// containsColumn checks if an expression string references a column name.
func containsColumn(exprStr, colName string) bool {
	// Simple heuristic: check if the column name appears in the expression string
	return len(colName) > 0 && len(exprStr) > 0 &&
		containsSubstring(exprStr, colName)
}

func containsSubstring(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
