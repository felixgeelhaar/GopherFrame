// Package expr provides a simplified expression engine for DataFrame operations.
// This initial implementation focuses on basic column references and literals
// without complex compute operations.
package expr

import (
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// Expr is the interface that all expression types must implement.
type Expr interface {
	// Evaluate executes the expression against a DataFrame and returns the result.
	Evaluate(df *core.DataFrame) (arrow.Array, error)

	// EvaluateWithPool executes the expression using a specific memory pool.
	EvaluateWithPool(df *core.DataFrame, pool memory.Allocator) (arrow.Array, error)

	// Name returns the output name of this expression.
	Name() string

	// String returns a string representation for debugging.
	String() string

	// Fluent methods for building expressions
	Add(other Expr) Expr
	Sub(other Expr) Expr
	Mul(other Expr) Expr
	Div(other Expr) Expr
	Gt(other Expr) Expr
	Lt(other Expr) Expr
	Eq(other Expr) Expr

	// String manipulation methods
	Contains(substring Expr) Expr
	StartsWith(prefix Expr) Expr
	EndsWith(suffix Expr) Expr

	// Date/time manipulation methods
	AddDays(days Expr) Expr
	AddMonths(months Expr) Expr
	AddYears(years Expr) Expr
	Year() Expr
	Month() Expr
	Day() Expr
	DayOfWeek() Expr
	DateTrunc(unit string) Expr

	// Vectorized operations using Arrow compute kernels
	VectorizedAdd(other Expr) Expr
	VectorizedSub(other Expr) Expr
	VectorizedMul(other Expr) Expr
	VectorizedDiv(other Expr) Expr
	VectorizedGt(other Expr) Expr
	VectorizedLt(other Expr) Expr
	VectorizedEq(other Expr) Expr

	// Vectorized aggregations
	VectorizedSum() Expr
	VectorizedMean() Expr
	VectorizedMin() Expr
	VectorizedMax() Expr
	VectorizedStdDev() Expr
	VectorizedVariance() Expr
}

// ColumnExpr represents a reference to an existing column.
type ColumnExpr struct {
	columnName string
}

// NewColumnExpr creates a new column reference expression.
func NewColumnExpr(name string) Expr {
	return &ColumnExpr{columnName: name}
}

// Col creates a column reference expression.
func Col(name string) Expr {
	return NewColumnExpr(name)
}

// Evaluate implements Expr.Evaluate for column references.
func (c *ColumnExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	series, err := df.Column(c.columnName)
	if err != nil {
		return nil, fmt.Errorf("failed to get column %s: %w", c.columnName, err)
	}

	// Return a copy of the array to maintain immutability
	arr := series.Array()
	arr.Retain()
	return arr, nil
}

// EvaluateWithPool implements Expr.EvaluateWithPool for column references.
func (c *ColumnExpr) EvaluateWithPool(df *core.DataFrame, pool memory.Allocator) (arrow.Array, error) {
	// Column references don't need pooled allocation, just return the existing array
	return c.Evaluate(df)
}

// Name implements Expr.Name for column references.
func (c *ColumnExpr) Name() string {
	return c.columnName
}

// String implements Expr.String for column references.
func (c *ColumnExpr) String() string {
	return fmt.Sprintf("Col(%s)", c.columnName)
}

// Add creates a binary expression that adds this column to another expression.
func (c *ColumnExpr) Add(other Expr) Expr {
	return NewBinaryExpr(c, other, "add")
}

// Sub creates a binary expression that subtracts another expression from this column.
func (c *ColumnExpr) Sub(other Expr) Expr {
	return NewBinaryExpr(c, other, "subtract")
}

// Mul creates a binary expression that multiplies this column by another expression.
func (c *ColumnExpr) Mul(other Expr) Expr {
	return NewBinaryExpr(c, other, "multiply")
}

// Div creates a binary expression that divides this column by another expression.
func (c *ColumnExpr) Div(other Expr) Expr {
	return NewBinaryExpr(c, other, "divide")
}

// Gt creates a binary expression that tests if this column is greater than another expression.
func (c *ColumnExpr) Gt(other Expr) Expr {
	return NewBinaryExpr(c, other, "greater")
}

// Lt creates a binary expression that tests if this column is less than another expression.
func (c *ColumnExpr) Lt(other Expr) Expr {
	return NewBinaryExpr(c, other, "less")
}

// Eq creates a binary expression that tests if this column is equal to another expression.
func (c *ColumnExpr) Eq(other Expr) Expr {
	return NewBinaryExpr(c, other, "equal")
}

// Contains creates a binary expression that tests if this string column contains a substring.
func (c *ColumnExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(c, substring, "contains")
}

// StartsWith creates a binary expression that tests if this string column starts with a prefix.
func (c *ColumnExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(c, prefix, "starts_with")
}

// EndsWith creates a binary expression that tests if this string column ends with a suffix.
func (c *ColumnExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(c, suffix, "ends_with")
}

// Date/time methods for ColumnExpr
func (c *ColumnExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(c, days, "add_days")
}

func (c *ColumnExpr) AddMonths(months Expr) Expr {
	return NewBinaryExpr(c, months, "add_months")
}

func (c *ColumnExpr) AddYears(years Expr) Expr {
	return NewBinaryExpr(c, years, "add_years")
}

func (c *ColumnExpr) Year() Expr {
	return NewUnaryExpr(c, "year")
}

func (c *ColumnExpr) Month() Expr {
	return NewUnaryExpr(c, "month")
}

func (c *ColumnExpr) Day() Expr {
	return NewUnaryExpr(c, "day")
}

func (c *ColumnExpr) DayOfWeek() Expr {
	return NewUnaryExpr(c, "day_of_week")
}

func (c *ColumnExpr) DateTrunc(unit string) Expr {
	return NewUnaryExpr(c, "date_trunc_"+unit)
}

// LiteralExpr represents a literal value.
type LiteralExpr struct {
	value    interface{}
	dataType arrow.DataType
	name     string
}

// NewLiteralExpr creates a new literal expression.
func NewLiteralExpr(value interface{}) Expr {
	dataType := inferDataType(value)
	return &LiteralExpr{
		value:    value,
		dataType: dataType,
		name:     fmt.Sprintf("Lit(%v)", value),
	}
}

// Lit creates a literal value expression.
func Lit(value interface{}) Expr {
	return NewLiteralExpr(value)
}

// Evaluate implements Expr.Evaluate for literal values.
func (l *LiteralExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	return l.EvaluateWithPool(df, memory.NewGoAllocator())
}

// EvaluateWithPool implements Expr.EvaluateWithPool for literal values.
func (l *LiteralExpr) EvaluateWithPool(df *core.DataFrame, pool memory.Allocator) (arrow.Array, error) {
	numRows := int(df.NumRows())

	// Create an array filled with the literal value using the provided pool
	switch l.dataType.ID() {
	case arrow.INT64:
		builder := array.NewInt64Builder(pool)
		defer builder.Release()

		intVal, ok := l.value.(int64)
		if !ok {
			// Try to convert int to int64
			if intValue, ok := l.value.(int); ok {
				intVal = int64(intValue)
			} else {
				return nil, fmt.Errorf("type mismatch: expected int64, got %T", l.value)
			}
		}

		builder.Reserve(numRows)
		for i := 0; i < numRows; i++ {
			builder.Append(intVal)
		}
		return builder.NewArray(), nil

	case arrow.FLOAT64:
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

		floatVal, ok := l.value.(float64)
		if !ok {
			return nil, fmt.Errorf("type mismatch: expected float64, got %T", l.value)
		}

		builder.Reserve(numRows)
		for i := 0; i < numRows; i++ {
			builder.Append(floatVal)
		}
		return builder.NewArray(), nil

	case arrow.STRING:
		builder := array.NewStringBuilder(pool)
		defer builder.Release()

		strVal, ok := l.value.(string)
		if !ok {
			return nil, fmt.Errorf("type mismatch: expected string, got %T", l.value)
		}

		builder.Reserve(numRows)
		for i := 0; i < numRows; i++ {
			builder.Append(strVal)
		}
		return builder.NewArray(), nil

	case arrow.BOOL:
		builder := array.NewBooleanBuilder(pool)
		defer builder.Release()

		boolVal, ok := l.value.(bool)
		if !ok {
			return nil, fmt.Errorf("type mismatch: expected bool, got %T", l.value)
		}

		builder.Reserve(numRows)
		for i := 0; i < numRows; i++ {
			builder.Append(boolVal)
		}
		return builder.NewArray(), nil

	default:
		return nil, fmt.Errorf("unsupported literal type: %s", l.dataType)
	}
}

// Name implements Expr.Name for literals.
func (l *LiteralExpr) Name() string {
	return l.name
}

// String implements Expr.String for literals.
func (l *LiteralExpr) String() string {
	return l.name
}

// Add creates a binary expression that adds this literal to another expression.
func (l *LiteralExpr) Add(other Expr) Expr {
	return NewBinaryExpr(l, other, "add")
}

// Sub creates a binary expression that subtracts another expression from this literal.
func (l *LiteralExpr) Sub(other Expr) Expr {
	return NewBinaryExpr(l, other, "subtract")
}

// Mul creates a binary expression that multiplies this literal by another expression.
func (l *LiteralExpr) Mul(other Expr) Expr {
	return NewBinaryExpr(l, other, "multiply")
}

// Div creates a binary expression that divides this literal by another expression.
func (l *LiteralExpr) Div(other Expr) Expr {
	return NewBinaryExpr(l, other, "divide")
}

func (l *LiteralExpr) Gt(other Expr) Expr {
	return NewBinaryExpr(l, other, "greater")
}

func (l *LiteralExpr) Lt(other Expr) Expr {
	return NewBinaryExpr(l, other, "less")
}

func (l *LiteralExpr) Eq(other Expr) Expr {
	return NewBinaryExpr(l, other, "equal")
}

// String manipulation methods for literals
func (l *LiteralExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(l, substring, "contains")
}

func (l *LiteralExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(l, prefix, "starts_with")
}

func (l *LiteralExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(l, suffix, "ends_with")
}

// Date/time methods for LiteralExpr
func (l *LiteralExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(l, days, "add_days")
}

func (l *LiteralExpr) AddMonths(months Expr) Expr {
	return NewBinaryExpr(l, months, "add_months")
}

func (l *LiteralExpr) AddYears(years Expr) Expr {
	return NewBinaryExpr(l, years, "add_years")
}

func (l *LiteralExpr) Year() Expr {
	return NewUnaryExpr(l, "year")
}

func (l *LiteralExpr) Month() Expr {
	return NewUnaryExpr(l, "month")
}

func (l *LiteralExpr) Day() Expr {
	return NewUnaryExpr(l, "day")
}

func (l *LiteralExpr) DayOfWeek() Expr {
	return NewUnaryExpr(l, "day_of_week")
}

func (l *LiteralExpr) DateTrunc(unit string) Expr {
	return NewUnaryExpr(l, "date_trunc_"+unit)
}

// BinaryExpr represents binary operations between two expressions.
// UnaryExpr represents unary operations on a single expression.
type UnaryExpr struct {
	operand  Expr
	operator string
}

// NewUnaryExpr creates a new unary expression.
func NewUnaryExpr(operand Expr, operator string) Expr {
	return &UnaryExpr{
		operand:  operand,
		operator: operator,
	}
}

// Evaluate implements Expr.Evaluate for unary operations.
func (u *UnaryExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	return u.EvaluateWithPool(df, memory.NewGoAllocator())
}

// EvaluateWithPool implements Expr.EvaluateWithPool for unary operations.
func (u *UnaryExpr) EvaluateWithPool(df *core.DataFrame, pool memory.Allocator) (arrow.Array, error) {
	// Evaluate the operand with the same pool
	operandArray, err := u.operand.EvaluateWithPool(df, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate operand: %w", err)
	}
	defer operandArray.Release()

	// Implement unary operations
	switch u.operator {
	case "year":
		return u.evaluateYear(operandArray)
	case "month":
		return u.evaluateMonth(operandArray)
	case "day":
		return u.evaluateDay(operandArray)
	case "day_of_week":
		return u.evaluateDayOfWeek(operandArray)
	default:
		if strings.HasPrefix(u.operator, "date_trunc_") {
			unit := strings.TrimPrefix(u.operator, "date_trunc_")
			return u.evaluateDateTrunc(operandArray, unit)
		}
		return nil, fmt.Errorf("unsupported unary operator: %s", u.operator)
	}
}

// Name implements Expr.Name for unary expressions.
func (u *UnaryExpr) Name() string {
	return fmt.Sprintf("%s(%s)", u.operator, u.operand.Name())
}

// String implements Expr.String for unary expressions.
func (u *UnaryExpr) String() string {
	return fmt.Sprintf("%s(%s)", u.operator, u.operand.String())
}

// Implement all binary operation methods for UnaryExpr
func (u *UnaryExpr) Add(other Expr) Expr {
	return NewBinaryExpr(u, other, "add")
}

func (u *UnaryExpr) Sub(other Expr) Expr {
	return NewBinaryExpr(u, other, "subtract")
}

func (u *UnaryExpr) Mul(other Expr) Expr {
	return NewBinaryExpr(u, other, "multiply")
}

func (u *UnaryExpr) Div(other Expr) Expr {
	return NewBinaryExpr(u, other, "divide")
}

func (u *UnaryExpr) Gt(other Expr) Expr {
	return NewBinaryExpr(u, other, "greater")
}

func (u *UnaryExpr) Lt(other Expr) Expr {
	return NewBinaryExpr(u, other, "less")
}

func (u *UnaryExpr) Eq(other Expr) Expr {
	return NewBinaryExpr(u, other, "equal")
}

// String manipulation methods for UnaryExpr
func (u *UnaryExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(u, substring, "contains")
}

func (u *UnaryExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(u, prefix, "starts_with")
}

func (u *UnaryExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(u, suffix, "ends_with")
}

// Date/time methods for UnaryExpr
func (u *UnaryExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(u, days, "add_days")
}

func (u *UnaryExpr) AddMonths(months Expr) Expr {
	return NewBinaryExpr(u, months, "add_months")
}

func (u *UnaryExpr) AddYears(years Expr) Expr {
	return NewBinaryExpr(u, years, "add_years")
}

func (u *UnaryExpr) Year() Expr {
	return NewUnaryExpr(u, "year")
}

func (u *UnaryExpr) Month() Expr {
	return NewUnaryExpr(u, "month")
}

func (u *UnaryExpr) Day() Expr {
	return NewUnaryExpr(u, "day")
}

func (u *UnaryExpr) DayOfWeek() Expr {
	return NewUnaryExpr(u, "day_of_week")
}

func (u *UnaryExpr) DateTrunc(unit string) Expr {
	return NewUnaryExpr(u, "date_trunc_"+unit)
}

type BinaryExpr struct {
	left     Expr
	right    Expr
	operator string
}

// NewBinaryExpr creates a new binary expression.
func NewBinaryExpr(left, right Expr, operator string) Expr {
	return &BinaryExpr{
		left:     left,
		right:    right,
		operator: operator,
	}
}

// Evaluate implements Expr.Evaluate for binary operations.
func (b *BinaryExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	return b.EvaluateWithPool(df, memory.NewGoAllocator())
}

// EvaluateWithPool implements Expr.EvaluateWithPool for binary operations.
func (b *BinaryExpr) EvaluateWithPool(df *core.DataFrame, pool memory.Allocator) (arrow.Array, error) {
	// Evaluate left and right operands with the same pool
	leftArray, err := b.left.EvaluateWithPool(df, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate left operand: %w", err)
	}
	defer leftArray.Release()

	rightArray, err := b.right.EvaluateWithPool(df, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate right operand: %w", err)
	}
	defer rightArray.Release()

	// Implement basic binary operations
	switch b.operator {
	case "greater":
		return b.evaluateGreater(leftArray, rightArray)
	case "less":
		return b.evaluateLess(leftArray, rightArray)
	case "equal":
		return b.evaluateEqual(leftArray, rightArray)
	case "add":
		return b.evaluateAdd(leftArray, rightArray)
	case "subtract":
		return b.evaluateSubtract(leftArray, rightArray)
	case "multiply":
		return b.evaluateMultiply(leftArray, rightArray)
	case "divide":
		return b.evaluateDivide(leftArray, rightArray)
	case "contains":
		return b.evaluateContains(leftArray, rightArray)
	case "starts_with":
		return b.evaluateStartsWith(leftArray, rightArray)
	case "ends_with":
		return b.evaluateEndsWith(leftArray, rightArray)
	case "add_days":
		return b.evaluateAddDays(leftArray, rightArray)
	case "add_months":
		return b.evaluateAddMonths(leftArray, rightArray)
	case "add_years":
		return b.evaluateAddYears(leftArray, rightArray)
	default:
		return nil, fmt.Errorf("unsupported binary operator: %s", b.operator)
	}
}

// evaluateGreater implements greater-than comparison
func (b *BinaryExpr) evaluateGreater(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	// Handle Float64 comparisons (most common case for our test)
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		leftFloat, ok := asFloat64Array(left)
		if !ok {
			return nil, fmt.Errorf("failed to cast left array to Float64")
		}
		rightFloat, ok := asFloat64Array(right)
		if !ok {
			return nil, fmt.Errorf("failed to cast right array to Float64")
		}

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) > rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Int64 comparisons
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

	return nil, fmt.Errorf("unsupported types for greater comparison: %s > %s", left.DataType(), right.DataType())
}

// Add placeholder implementations for other operations
func (b *BinaryExpr) evaluateLess(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	// Handle Float64 comparisons
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		leftFloat := left.(*array.Float64)
		rightFloat := right.(*array.Float64)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) < rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Int64 comparisons
	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		leftInt := left.(*array.Int64)
		rightInt := right.(*array.Int64)

		for i := 0; i < left.Len(); i++ {
			if leftInt.IsNull(i) || rightInt.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftInt.Value(i) < rightInt.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	return nil, fmt.Errorf("unsupported types for less-than comparison: %s < %s", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateEqual(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	// Handle String comparisons
	if left.DataType().ID() == arrow.STRING && right.DataType().ID() == arrow.STRING {
		leftStr := left.(*array.String)
		rightStr := right.(*array.String)

		for i := 0; i < left.Len(); i++ {
			if leftStr.IsNull(i) || rightStr.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftStr.Value(i) == rightStr.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Int64 comparisons
	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		leftInt := left.(*array.Int64)
		rightInt := right.(*array.Int64)

		for i := 0; i < left.Len(); i++ {
			if leftInt.IsNull(i) || rightInt.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftInt.Value(i) == rightInt.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Float64 comparisons
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		leftFloat := left.(*array.Float64)
		rightFloat := right.(*array.Float64)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) == rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	return nil, fmt.Errorf("unsupported types for equality comparison: %s == %s", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateAdd(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Handle Float64 addition
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

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

	// Handle Int64 addition
	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		builder := array.NewInt64Builder(pool)
		defer builder.Release()

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

	return nil, fmt.Errorf("unsupported types for addition: %s + %s", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateSubtract(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Handle Float64 subtraction
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

		leftFloat := left.(*array.Float64)
		rightFloat := right.(*array.Float64)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftFloat.Value(i) - rightFloat.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Int64 subtraction
	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		builder := array.NewInt64Builder(pool)
		defer builder.Release()

		leftInt := left.(*array.Int64)
		rightInt := right.(*array.Int64)

		for i := 0; i < left.Len(); i++ {
			if leftInt.IsNull(i) || rightInt.IsNull(i) {
				builder.AppendNull()
			} else {
				builder.Append(leftInt.Value(i) - rightInt.Value(i))
			}
		}

		return builder.NewArray(), nil
	}

	return nil, fmt.Errorf("unsupported types for subtraction: %s - %s", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateMultiply(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Handle Float64 multiplication
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

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

	return nil, fmt.Errorf("unsupported types for multiply: %s * %s", left.DataType(), right.DataType())
}

func (b *BinaryExpr) evaluateDivide(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Handle Float64 division
	if left.DataType().ID() == arrow.FLOAT64 && right.DataType().ID() == arrow.FLOAT64 {
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()

		leftFloat := left.(*array.Float64)
		rightFloat := right.(*array.Float64)

		for i := 0; i < left.Len(); i++ {
			if leftFloat.IsNull(i) || rightFloat.IsNull(i) {
				builder.AppendNull()
			} else {
				rightVal := rightFloat.Value(i)
				if rightVal == 0.0 {
					return nil, fmt.Errorf("division by zero at index %d", i)
				}
				builder.Append(leftFloat.Value(i) / rightVal)
			}
		}

		return builder.NewArray(), nil
	}

	// Handle Int64 division (integer division)
	if left.DataType().ID() == arrow.INT64 && right.DataType().ID() == arrow.INT64 {
		builder := array.NewInt64Builder(pool)
		defer builder.Release()

		leftInt := left.(*array.Int64)
		rightInt := right.(*array.Int64)

		for i := 0; i < left.Len(); i++ {
			if leftInt.IsNull(i) || rightInt.IsNull(i) {
				builder.AppendNull()
			} else {
				rightVal := rightInt.Value(i)
				if rightVal == 0 {
					return nil, fmt.Errorf("division by zero at index %d", i)
				}
				builder.Append(leftInt.Value(i) / rightVal)
			}
		}

		return builder.NewArray(), nil
	}

	return nil, fmt.Errorf("unsupported types for division: %s / %s", left.DataType(), right.DataType())
}

// Name implements Expr.Name for binary operations.
func (b *BinaryExpr) Name() string {
	return fmt.Sprintf("(%s %s %s)", b.left.Name(), b.operator, b.right.Name())
}

// String implements Expr.String for binary operations.
func (b *BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.left.String(), b.operator, b.right.String())
}

// Fluent methods for BinaryExpr to enable further chaining
func (b *BinaryExpr) Add(other Expr) Expr {
	return NewBinaryExpr(b, other, "add")
}

func (b *BinaryExpr) Sub(other Expr) Expr {
	return NewBinaryExpr(b, other, "subtract")
}

func (b *BinaryExpr) Mul(other Expr) Expr {
	return NewBinaryExpr(b, other, "multiply")
}

func (b *BinaryExpr) Div(other Expr) Expr {
	return NewBinaryExpr(b, other, "divide")
}

func (b *BinaryExpr) Gt(other Expr) Expr {
	return NewBinaryExpr(b, other, "greater")
}

func (b *BinaryExpr) Lt(other Expr) Expr {
	return NewBinaryExpr(b, other, "less")
}

func (b *BinaryExpr) Eq(other Expr) Expr {
	return NewBinaryExpr(b, other, "equal")
}

// String manipulation methods for binary expressions
func (b *BinaryExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(b, substring, "contains")
}

func (b *BinaryExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(b, prefix, "starts_with")
}

func (b *BinaryExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(b, suffix, "ends_with")
}

// Date/time methods for BinaryExpr
func (b *BinaryExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(b, days, "add_days")
}

func (b *BinaryExpr) AddMonths(months Expr) Expr {
	return NewBinaryExpr(b, months, "add_months")
}

func (b *BinaryExpr) AddYears(years Expr) Expr {
	return NewBinaryExpr(b, years, "add_years")
}

func (b *BinaryExpr) Year() Expr {
	return NewUnaryExpr(b, "year")
}

func (b *BinaryExpr) Month() Expr {
	return NewUnaryExpr(b, "month")
}

func (b *BinaryExpr) Day() Expr {
	return NewUnaryExpr(b, "day")
}

func (b *BinaryExpr) DayOfWeek() Expr {
	return NewUnaryExpr(b, "day_of_week")
}

func (b *BinaryExpr) DateTrunc(unit string) Expr {
	return NewUnaryExpr(b, "date_trunc_"+unit)
}

// evaluateContains implements string contains comparison
func (b *BinaryExpr) evaluateContains(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	// Both operands must be strings
	if left.DataType().ID() != arrow.STRING || right.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("contains operation requires string operands, got %s and %s", left.DataType(), right.DataType())
	}

	leftStr, ok := asStringArray(left)
	if !ok {
		return nil, fmt.Errorf("failed to cast left array to String")
	}

	rightStr, ok := asStringArray(right)
	if !ok {
		return nil, fmt.Errorf("failed to cast right array to String")
	}

	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if leftStr.IsNull(i) || rightStr.IsNull(i) {
			builder.AppendNull()
		} else {
			leftVal := leftStr.Value(i)
			rightVal := rightStr.Value(i)
			result := strings.Contains(leftVal, rightVal)
			builder.Append(result)
		}
	}

	return builder.NewArray(), nil
}

// evaluateStartsWith implements string starts_with comparison
func (b *BinaryExpr) evaluateStartsWith(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	// Both operands must be strings
	if left.DataType().ID() != arrow.STRING || right.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("starts_with operation requires string operands, got %s and %s", left.DataType(), right.DataType())
	}

	leftStr, ok := asStringArray(left)
	if !ok {
		return nil, fmt.Errorf("failed to cast left array to String")
	}

	rightStr, ok := asStringArray(right)
	if !ok {
		return nil, fmt.Errorf("failed to cast right array to String")
	}

	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if leftStr.IsNull(i) || rightStr.IsNull(i) {
			builder.AppendNull()
		} else {
			leftVal := leftStr.Value(i)
			rightVal := rightStr.Value(i)
			result := strings.HasPrefix(leftVal, rightVal)
			builder.Append(result)
		}
	}

	return builder.NewArray(), nil
}

// evaluateEndsWith implements string ends_with comparison
func (b *BinaryExpr) evaluateEndsWith(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	// Both operands must be strings
	if left.DataType().ID() != arrow.STRING || right.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("ends_with operation requires string operands, got %s and %s", left.DataType(), right.DataType())
	}

	leftStr, ok := asStringArray(left)
	if !ok {
		return nil, fmt.Errorf("failed to cast left array to String")
	}

	rightStr, ok := asStringArray(right)
	if !ok {
		return nil, fmt.Errorf("failed to cast right array to String")
	}

	pool := memory.NewGoAllocator()
	builder := array.NewBooleanBuilder(pool)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if leftStr.IsNull(i) || rightStr.IsNull(i) {
			builder.AppendNull()
		} else {
			leftVal := leftStr.Value(i)
			rightVal := rightStr.Value(i)
			result := strings.HasSuffix(leftVal, rightVal)
			builder.Append(result)
		}
	}

	return builder.NewArray(), nil
}

// Helper functions for safe type assertions
func asFloat64Array(arr arrow.Array) (*array.Float64, bool) {
	f64arr, ok := arr.(*array.Float64)
	return f64arr, ok
}

func asInt64Array(arr arrow.Array) (*array.Int64, bool) {
	i64arr, ok := arr.(*array.Int64)
	return i64arr, ok
}

func asStringArray(arr arrow.Array) (*array.String, bool) {
	strArr, ok := arr.(*array.String)
	return strArr, ok
}

// Helper function to infer Arrow data type from Go value
func inferDataType(value interface{}) arrow.DataType {
	switch value.(type) {
	case bool:
		return arrow.FixedWidthTypes.Boolean
	case int8:
		return arrow.PrimitiveTypes.Int8
	case int16:
		return arrow.PrimitiveTypes.Int16
	case int32, int:
		return arrow.PrimitiveTypes.Int64 // Promote int to int64 for consistency
	case int64:
		return arrow.PrimitiveTypes.Int64
	case uint8:
		return arrow.PrimitiveTypes.Uint8
	case uint16:
		return arrow.PrimitiveTypes.Uint16
	case uint32, uint:
		return arrow.PrimitiveTypes.Uint32
	case uint64:
		return arrow.PrimitiveTypes.Uint64
	case float32:
		return arrow.PrimitiveTypes.Float32
	case float64:
		return arrow.PrimitiveTypes.Float64
	case string:
		return arrow.BinaryTypes.String
	case time.Time:
		return arrow.FixedWidthTypes.Timestamp_us
	default:
		// Default to string for unknown types
		return arrow.BinaryTypes.String
	}
}

// Date/time evaluation functions for UnaryExpr
func (u *UnaryExpr) evaluateYear(operand arrow.Array) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt32Builder(pool)
	defer builder.Release()

	switch operand.DataType().ID() {
	case arrow.DATE32:
		dateArray := operand.(*array.Date32)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				// Date32 is days since Unix epoch
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()
				year := t.Year()
				if year < -2147483648 || year > 2147483647 {
					return nil, fmt.Errorf("year %d out of int32 range", year)
				}
				builder.Append(int32(year))
			}
		}
	case arrow.DATE64:
		dateArray := operand.(*array.Date64)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				// Date64 is milliseconds since Unix epoch
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()
				year := t.Year()
				if year < -2147483648 || year > 2147483647 {
					return nil, fmt.Errorf("year %d out of int32 range", year)
				}
				builder.Append(int32(year))
			}
		}
	case arrow.TIMESTAMP:
		timestampArray := operand.(*array.Timestamp)
		for i := 0; i < operand.Len(); i++ {
			if timestampArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ts := int64(timestampArray.Value(i))
				t := time.Unix(ts/1000000, (ts%1000000)*1000).UTC()
				year := t.Year()
				if year < -2147483648 || year > 2147483647 {
					return nil, fmt.Errorf("year %d out of int32 range", year)
				}
				builder.Append(int32(year))
			}
		}
	default:
		return nil, fmt.Errorf("unsupported type for year extraction: %s", operand.DataType())
	}

	return builder.NewArray(), nil
}

func (u *UnaryExpr) evaluateMonth(operand arrow.Array) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt32Builder(pool)
	defer builder.Release()

	switch operand.DataType().ID() {
	case arrow.DATE32:
		dateArray := operand.(*array.Date32)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()
				builder.Append(int32(t.Month()))
			}
		}
	case arrow.DATE64:
		dateArray := operand.(*array.Date64)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()
				builder.Append(int32(t.Month()))
			}
		}
	case arrow.TIMESTAMP:
		timestampArray := operand.(*array.Timestamp)
		for i := 0; i < operand.Len(); i++ {
			if timestampArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ts := int64(timestampArray.Value(i))
				t := time.Unix(ts/1000000, (ts%1000000)*1000).UTC()
				builder.Append(int32(t.Month()))
			}
		}
	default:
		return nil, fmt.Errorf("unsupported type for month extraction: %s", operand.DataType())
	}

	return builder.NewArray(), nil
}

func (u *UnaryExpr) evaluateDay(operand arrow.Array) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt32Builder(pool)
	defer builder.Release()

	switch operand.DataType().ID() {
	case arrow.DATE32:
		dateArray := operand.(*array.Date32)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()
				builder.Append(int32(t.Day()))
			}
		}
	case arrow.DATE64:
		dateArray := operand.(*array.Date64)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()
				builder.Append(int32(t.Day()))
			}
		}
	case arrow.TIMESTAMP:
		timestampArray := operand.(*array.Timestamp)
		for i := 0; i < operand.Len(); i++ {
			if timestampArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ts := int64(timestampArray.Value(i))
				t := time.Unix(ts/1000000, (ts%1000000)*1000).UTC()
				builder.Append(int32(t.Day()))
			}
		}
	default:
		return nil, fmt.Errorf("unsupported type for day extraction: %s", operand.DataType())
	}

	return builder.NewArray(), nil
}

func (u *UnaryExpr) evaluateDayOfWeek(operand arrow.Array) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	builder := array.NewInt32Builder(pool)
	defer builder.Release()

	switch operand.DataType().ID() {
	case arrow.DATE32:
		dateArray := operand.(*array.Date32)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()
				// Convert Go's Sunday=0 to ISO's Monday=1
				weekday := int32(t.Weekday())
				if weekday == 0 {
					weekday = 7
				}
				builder.Append(weekday)
			}
		}
	case arrow.DATE64:
		dateArray := operand.(*array.Date64)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()
				weekday := int32(t.Weekday())
				if weekday == 0 {
					weekday = 7
				}
				builder.Append(weekday)
			}
		}
	case arrow.TIMESTAMP:
		timestampArray := operand.(*array.Timestamp)
		for i := 0; i < operand.Len(); i++ {
			if timestampArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ts := int64(timestampArray.Value(i))
				t := time.Unix(ts/1000000, (ts%1000000)*1000).UTC()
				weekday := int32(t.Weekday())
				if weekday == 0 {
					weekday = 7
				}
				builder.Append(weekday)
			}
		}
	default:
		return nil, fmt.Errorf("unsupported type for day_of_week extraction: %s", operand.DataType())
	}

	return builder.NewArray(), nil
}

func (u *UnaryExpr) evaluateDateTrunc(operand arrow.Array, unit string) (arrow.Array, error) {
	pool := memory.NewGoAllocator()

	switch operand.DataType().ID() {
	case arrow.DATE32:
		builder := array.NewDate32Builder(pool)
		defer builder.Release()

		dateArray := operand.(*array.Date32)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()
				truncated := truncateDate(t, unit)
				truncatedDays := int32(truncated.Unix() / 86400)
				builder.Append(arrow.Date32(truncatedDays))
			}
		}
		return builder.NewArray(), nil

	case arrow.DATE64:
		builder := array.NewDate64Builder(pool)
		defer builder.Release()

		dateArray := operand.(*array.Date64)
		for i := 0; i < operand.Len(); i++ {
			if dateArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()
				truncated := truncateDate(t, unit)
				truncatedMs := truncated.Unix() * 1000
				builder.Append(arrow.Date64(truncatedMs))
			}
		}
		return builder.NewArray(), nil

	case arrow.TIMESTAMP:
		builder := array.NewTimestampBuilder(pool, operand.DataType().(*arrow.TimestampType))
		defer builder.Release()

		timestampArray := operand.(*array.Timestamp)
		for i := 0; i < operand.Len(); i++ {
			if timestampArray.IsNull(i) {
				builder.AppendNull()
			} else {
				ts := int64(timestampArray.Value(i))
				t := time.Unix(ts/1000000, (ts%1000000)*1000).UTC()
				truncated := truncateDate(t, unit)
				truncatedTs := truncated.Unix() * 1000000
				builder.Append(arrow.Timestamp(truncatedTs))
			}
		}
		return builder.NewArray(), nil

	default:
		return nil, fmt.Errorf("unsupported type for date truncation: %s", operand.DataType())
	}
}

// Binary date operations
func (b *BinaryExpr) evaluateAddDays(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Right operand should be Int32 or Int64 (number of days)
	if right.DataType().ID() != arrow.INT32 && right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("AddDays requires integer days parameter, got %s", right.DataType())
	}

	switch left.DataType().ID() {
	case arrow.DATE32:
		builder := array.NewDate32Builder(pool)
		defer builder.Release()

		dateArray := left.(*array.Date32)
		var daysArray arrow.Array
		if right.DataType().ID() == arrow.INT32 {
			daysArray = right
		} else {
			// Convert Int64 to Int32
			int64Array := right.(*array.Int64)
			int32Builder := array.NewInt32Builder(pool)
			defer int32Builder.Release()
			for i := 0; i < right.Len(); i++ {
				if int64Array.IsNull(i) {
					int32Builder.AppendNull()
				} else {
					int32Builder.Append(int32(int64Array.Value(i)))
				}
			}
			daysArray = int32Builder.NewArray()
			defer daysArray.Release()
		}

		daysInt32 := daysArray.(*array.Int32)
		for i := 0; i < left.Len(); i++ {
			if dateArray.IsNull(i) || daysInt32.IsNull(i) {
				builder.AppendNull()
			} else {
				originalDays := dateArray.Value(i)
				addDays := daysInt32.Value(i)
				builder.Append(arrow.Date32(int32(originalDays) + addDays))
			}
		}
		return builder.NewArray(), nil

	case arrow.DATE64:
		builder := array.NewDate64Builder(pool)
		defer builder.Release()

		dateArray := left.(*array.Date64)
		for i := 0; i < left.Len(); i++ {
			if dateArray.IsNull(i) || right.IsNull(i) {
				builder.AppendNull()
			} else {
				originalMs := dateArray.Value(i)
				var addDays int64
				if right.DataType().ID() == arrow.INT32 {
					addDays = int64(right.(*array.Int32).Value(i))
				} else {
					addDays = right.(*array.Int64).Value(i)
				}
				newMs := int64(originalMs) + (addDays * 86400 * 1000) // days to milliseconds
				builder.Append(arrow.Date64(newMs))
			}
		}
		return builder.NewArray(), nil

	default:
		return nil, fmt.Errorf("unsupported type for AddDays: %s", left.DataType())
	}
}

func (b *BinaryExpr) evaluateAddMonths(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Right operand should be Int32 or Int64 (number of months)
	if right.DataType().ID() != arrow.INT32 && right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("AddMonths requires integer months parameter, got %s", right.DataType())
	}

	switch left.DataType().ID() {
	case arrow.DATE32:
		builder := array.NewDate32Builder(pool)
		defer builder.Release()

		dateArray := left.(*array.Date32)
		for i := 0; i < left.Len(); i++ {
			if dateArray.IsNull(i) || right.IsNull(i) {
				builder.AppendNull()
			} else {
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()

				var addMonths int
				if right.DataType().ID() == arrow.INT32 {
					addMonths = int(right.(*array.Int32).Value(i))
				} else {
					addMonths = int(right.(*array.Int64).Value(i))
				}

				newTime := t.AddDate(0, addMonths, 0)
				newDays := int32(newTime.Unix() / 86400)
				builder.Append(arrow.Date32(newDays))
			}
		}
		return builder.NewArray(), nil

	case arrow.DATE64:
		builder := array.NewDate64Builder(pool)
		defer builder.Release()

		dateArray := left.(*array.Date64)
		for i := 0; i < left.Len(); i++ {
			if dateArray.IsNull(i) || right.IsNull(i) {
				builder.AppendNull()
			} else {
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()

				var addMonths int
				if right.DataType().ID() == arrow.INT32 {
					addMonths = int(right.(*array.Int32).Value(i))
				} else {
					addMonths = int(right.(*array.Int64).Value(i))
				}

				newTime := t.AddDate(0, addMonths, 0)
				newMs := newTime.Unix() * 1000
				builder.Append(arrow.Date64(newMs))
			}
		}
		return builder.NewArray(), nil

	default:
		return nil, fmt.Errorf("unsupported type for AddMonths: %s", left.DataType())
	}
}

func (b *BinaryExpr) evaluateAddYears(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	pool := memory.NewGoAllocator()

	// Right operand should be Int32 or Int64 (number of years)
	if right.DataType().ID() != arrow.INT32 && right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("AddYears requires integer years parameter, got %s", right.DataType())
	}

	switch left.DataType().ID() {
	case arrow.DATE32:
		builder := array.NewDate32Builder(pool)
		defer builder.Release()

		dateArray := left.(*array.Date32)
		for i := 0; i < left.Len(); i++ {
			if dateArray.IsNull(i) || right.IsNull(i) {
				builder.AppendNull()
			} else {
				days := dateArray.Value(i)
				t := time.Unix(int64(days)*86400, 0).UTC()

				var addYears int
				if right.DataType().ID() == arrow.INT32 {
					addYears = int(right.(*array.Int32).Value(i))
				} else {
					addYears = int(right.(*array.Int64).Value(i))
				}

				newTime := t.AddDate(addYears, 0, 0)
				newDays := int32(newTime.Unix() / 86400)
				builder.Append(arrow.Date32(newDays))
			}
		}
		return builder.NewArray(), nil

	case arrow.DATE64:
		builder := array.NewDate64Builder(pool)
		defer builder.Release()

		dateArray := left.(*array.Date64)
		for i := 0; i < left.Len(); i++ {
			if dateArray.IsNull(i) || right.IsNull(i) {
				builder.AppendNull()
			} else {
				ms := int64(dateArray.Value(i))
				t := time.Unix(ms/1000, (ms%1000)*1000000).UTC()

				var addYears int
				if right.DataType().ID() == arrow.INT32 {
					addYears = int(right.(*array.Int32).Value(i))
				} else {
					addYears = int(right.(*array.Int64).Value(i))
				}

				newTime := t.AddDate(addYears, 0, 0)
				newMs := newTime.Unix() * 1000
				builder.Append(arrow.Date64(newMs))
			}
		}
		return builder.NewArray(), nil

	default:
		return nil, fmt.Errorf("unsupported type for AddYears: %s", left.DataType())
	}
}

// Helper function to truncate date based on unit
func truncateDate(t time.Time, unit string) time.Time {
	switch unit {
	case "year":
		return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	case "month":
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	case "day":
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	case "hour":
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	case "minute":
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	case "second":
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	default:
		return t // Return original if unknown unit
	}
}
