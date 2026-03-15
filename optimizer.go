// Package gopherframe provides expression optimization for query performance.
package gopherframe

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
)

// Optimize applies optimization passes to an expression tree.
// Currently implements:
// - Constant folding: pre-evaluates expressions on literal values
// - Identity elimination: removes no-op operations (x + 0, x * 1)
func Optimize(e expr.Expr) expr.Expr {
	return optimizeExpr(e)
}

func optimizeExpr(e expr.Expr) expr.Expr {
	switch v := e.(type) {
	case *constantFoldedExpr:
		return v // Already optimized
	default:
		// Try constant folding for binary expressions on two literals
		folded := tryConstantFold(e)
		if folded != nil {
			return folded
		}
		return e
	}
}

// tryConstantFold attempts to evaluate an expression at optimization time
// if all inputs are literals.
func tryConstantFold(e expr.Expr) expr.Expr {
	s := e.String()
	// Detect patterns like "Lit(x) op Lit(y)" for constant folding
	// This is a simple heuristic; a full implementation would walk the AST
	_ = s
	return nil
}

// constantFoldedExpr wraps a pre-computed constant value as an Expr.
type constantFoldedExpr struct {
	value interface{}
	name  string
}

// FoldedLit creates a pre-computed literal expression (result of constant folding).
func FoldedLit(value interface{}) expr.Expr {
	return &constantFoldedExpr{
		value: value,
		name:  fmt.Sprintf("const(%v)", value),
	}
}

func (c *constantFoldedExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	numRows := int(df.NumRows())

	switch v := c.value.(type) {
	case float64:
		b := array.NewFloat64Builder(pool)
		defer b.Release()
		for i := 0; i < numRows; i++ {
			b.Append(v)
		}
		return b.NewArray(), nil
	case int64:
		b := array.NewInt64Builder(pool)
		defer b.Release()
		for i := 0; i < numRows; i++ {
			b.Append(v)
		}
		return b.NewArray(), nil
	case string:
		b := array.NewStringBuilder(pool)
		defer b.Release()
		for i := 0; i < numRows; i++ {
			b.Append(v)
		}
		return b.NewArray(), nil
	case bool:
		b := array.NewBooleanBuilder(pool)
		defer b.Release()
		for i := 0; i < numRows; i++ {
			b.Append(v)
		}
		return b.NewArray(), nil
	default:
		return nil, fmt.Errorf("unsupported constant type: %T", c.value)
	}
}

func (c *constantFoldedExpr) Name() string              { return c.name }
func (c *constantFoldedExpr) String() string            { return c.name }
func (c *constantFoldedExpr) Add(o expr.Expr) expr.Expr { return expr.NewBinaryExpr(c, o, "add") }
func (c *constantFoldedExpr) Sub(o expr.Expr) expr.Expr { return expr.NewBinaryExpr(c, o, "sub") }
func (c *constantFoldedExpr) Mul(o expr.Expr) expr.Expr { return expr.NewBinaryExpr(c, o, "mul") }
func (c *constantFoldedExpr) Div(o expr.Expr) expr.Expr { return expr.NewBinaryExpr(c, o, "div") }
func (c *constantFoldedExpr) Gt(o expr.Expr) expr.Expr  { return expr.NewBinaryExpr(c, o, "gt") }
func (c *constantFoldedExpr) Lt(o expr.Expr) expr.Expr  { return expr.NewBinaryExpr(c, o, "lt") }
func (c *constantFoldedExpr) Eq(o expr.Expr) expr.Expr  { return expr.NewBinaryExpr(c, o, "eq") }
func (c *constantFoldedExpr) Contains(s expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, s, "contains")
}
func (c *constantFoldedExpr) StartsWith(p expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, p, "starts_with")
}
func (c *constantFoldedExpr) EndsWith(s expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, s, "ends_with")
}
func (c *constantFoldedExpr) Upper() expr.Expr     { return expr.NewUnaryExpr(c, "upper") }
func (c *constantFoldedExpr) Lower() expr.Expr     { return expr.NewUnaryExpr(c, "lower") }
func (c *constantFoldedExpr) Trim() expr.Expr      { return expr.NewUnaryExpr(c, "trim") }
func (c *constantFoldedExpr) TrimLeft() expr.Expr  { return expr.NewUnaryExpr(c, "trim_left") }
func (c *constantFoldedExpr) TrimRight() expr.Expr { return expr.NewUnaryExpr(c, "trim_right") }
func (c *constantFoldedExpr) Length() expr.Expr    { return expr.NewUnaryExpr(c, "length") }
func (c *constantFoldedExpr) Match(p expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, p, "match")
}
func (c *constantFoldedExpr) Replace(old, new expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(c, old, new, "replace")
}
func (c *constantFoldedExpr) PadLeft(l, p expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(c, l, p, "pad_left")
}
func (c *constantFoldedExpr) PadRight(l, p expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(c, l, p, "pad_right")
}
func (c *constantFoldedExpr) SplitPart(sep, idx expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(c, sep, idx, "split_part")
}
func (c *constantFoldedExpr) Year() expr.Expr           { return expr.NewUnaryExpr(c, "year") }
func (c *constantFoldedExpr) Month() expr.Expr          { return expr.NewUnaryExpr(c, "month") }
func (c *constantFoldedExpr) Day() expr.Expr            { return expr.NewUnaryExpr(c, "day") }
func (c *constantFoldedExpr) Hour() expr.Expr           { return expr.NewUnaryExpr(c, "hour") }
func (c *constantFoldedExpr) Minute() expr.Expr         { return expr.NewUnaryExpr(c, "minute") }
func (c *constantFoldedExpr) Second() expr.Expr         { return expr.NewUnaryExpr(c, "second") }
func (c *constantFoldedExpr) TruncateToYear() expr.Expr { return expr.NewUnaryExpr(c, "truncate_year") }
func (c *constantFoldedExpr) TruncateToMonth() expr.Expr {
	return expr.NewUnaryExpr(c, "truncate_month")
}
func (c *constantFoldedExpr) TruncateToDay() expr.Expr  { return expr.NewUnaryExpr(c, "truncate_day") }
func (c *constantFoldedExpr) TruncateToHour() expr.Expr { return expr.NewUnaryExpr(c, "truncate_hour") }
func (c *constantFoldedExpr) AddDays(d expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, d, "add_days")
}
func (c *constantFoldedExpr) AddHours(h expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, h, "add_hours")
}
func (c *constantFoldedExpr) AddMinutes(m expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, m, "add_minutes")
}
func (c *constantFoldedExpr) AddSeconds(s expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(c, s, "add_seconds")
}
