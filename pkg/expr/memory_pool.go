package expr

import (
	"fmt"
	"sync"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// MemoryPool provides thread-safe memory pool management for expressions
type MemoryPool struct {
	pool memory.Allocator
	mu   sync.RWMutex
}

// NewMemoryPool creates a new memory pool for expression evaluation
func NewMemoryPool() *MemoryPool {
	return &MemoryPool{
		pool: memory.NewGoAllocator(),
	}
}

// GetAllocator returns the underlying allocator
func (m *MemoryPool) GetAllocator() memory.Allocator {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.pool
}

// ExprContext holds shared context for expression evaluation
type ExprContext struct {
	Pool *MemoryPool
}

// NewExprContext creates a new expression evaluation context
func NewExprContext() *ExprContext {
	return &ExprContext{
		Pool: NewMemoryPool(),
	}
}

// DefaultContext is a shared context for simple use cases
var DefaultContext = NewExprContext()

// Helper function to evaluate any expression with a pool
func evaluateWithPool(expr Expr, df *core.DataFrame, pool memory.Allocator) (arrow.Array, error) {
	return expr.EvaluateWithPool(df, pool)
}

// Update existing evaluation methods to use pools internally
func (u *UnaryExpr) evaluateWithPool(operand arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	switch u.operator {
	case "year":
		return u.evaluateYearWithPool(operand, pool)
	case "month":
		return u.evaluateMonthWithPool(operand, pool)
	case "day":
		return u.evaluateDayWithPool(operand, pool)
	case "day_of_week":
		return u.evaluateDayOfWeekWithPool(operand, pool)
	default:
		// For operations not yet optimized, use regular evaluation
		return nil, fmt.Errorf("unimplemented pool evaluation for operator: %s", u.operator)
	}
}

func (b *BinaryExpr) evaluateWithPool(left, right arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	switch b.operator {
	case "add":
		return b.evaluateAddWithPool(left, right, pool)
	case "multiply":
		return b.evaluateMultiplyWithPool(left, right, pool)
	case "greater":
		return b.evaluateGreaterWithPool(left, right, pool)
	default:
		// For operations not yet optimized, use regular evaluation
		return nil, fmt.Errorf("unimplemented pool evaluation for operator: %s", b.operator)
	}
}

// Optimized evaluation methods using pools
func (u *UnaryExpr) evaluateYearWithPool(operand arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	// For now, delegate to existing implementation
	return u.evaluateYear(operand)
}

func (u *UnaryExpr) evaluateMonthWithPool(operand arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	return u.evaluateMonth(operand)
}

func (u *UnaryExpr) evaluateDayWithPool(operand arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	return u.evaluateDay(operand)
}

func (u *UnaryExpr) evaluateDayOfWeekWithPool(operand arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	return u.evaluateDayOfWeek(operand)
}

func (b *BinaryExpr) evaluateAddWithPool(left, right arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, errLengthMismatch(left.Len(), right.Len())
	}

	// Handle Float64 addition with pool
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

		// Reserve capacity
		builder.Reserve(left.Len())

		leftFloat := left.(*array.Float64)
		rightFloat := right.(*array.Float64)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) + rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Int64 addition with pool
	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		builder := array.NewInt64Builder(pool)
		defer builder.Release()

		builder.Reserve(left.Len())

		leftInt := left.(*array.Int64)
		rightInt := right.(*array.Int64)

		for i := 0; i < left.Len(); i++ {
			if leftInt.IsNull(i) || rightInt.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftInt.Value(i) + rightInt.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	return nil, errUnsupportedTypes("addition", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateMultiplyWithPool(left, right arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, errLengthMismatch(left.Len(), right.Len())
	}

	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

		builder.Reserve(left.Len())

		leftFloat := left.(*array.Float64)
		rightFloat := right.(*array.Float64)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) * rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	return nil, errUnsupportedTypes("multiply", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateGreaterWithPool(left, right arrow.Array, pool memory.Allocator) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, errLengthMismatch(left.Len(), right.Len())
	}

	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	builder.Reserve(left.Len())

	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		leftFloat, _ := asFloat64Array(left)
		rightFloat, _ := asFloat64Array(right)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) > rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		leftInt := left.(*array.Int64)
		rightInt := right.(*array.Int64)

		for i := 0; i < left.Len(); i++ {
			if leftInt.IsNull(i) || rightInt.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftInt.Value(i) > rightInt.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	return nil, errUnsupportedTypes("greater", left.DataType(), right.DataType())
}
