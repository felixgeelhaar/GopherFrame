// Package gopherframe provides expression builders for DataFrame operations.
package gopherframe

import (
	"github.com/felixgeelhaar/gopherFrame/pkg/expr"
)

// Col creates a column reference expression.
// This is the primary way to reference DataFrame columns in operations.
func Col(name string) expr.Expr {
	return expr.Col(name)
}

// Lit creates a literal value expression.
// The literal value will be broadcast to match the DataFrame length.
func Lit(value interface{}) expr.Expr {
	return expr.Lit(value)
}

// Expression builder methods
// These will be added to the core expression types to enable fluent chaining

// Extend ColumnExpr with fluent methods
type ColumnExpr interface {
	expr.Expr
	// Add creates an addition expression
	Add(other expr.Expr) expr.Expr
	// Sub creates a subtraction expression  
	Sub(other expr.Expr) expr.Expr
	// Mul creates a multiplication expression
	Mul(other expr.Expr) expr.Expr
	// Div creates a division expression
	Div(other expr.Expr) expr.Expr
	// Gt creates a greater-than comparison
	Gt(other expr.Expr) expr.Expr
	// Lt creates a less-than comparison
	Lt(other expr.Expr) expr.Expr
	// Eq creates an equality comparison
	Eq(other expr.Expr) expr.Expr
}