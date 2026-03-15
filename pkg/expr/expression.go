// Package expr provides a simplified expression engine for DataFrame operations.
// This initial implementation focuses on basic column references and literals
// without complex compute operations.
package expr

import (
	"fmt"
	"regexp"
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
	Upper() Expr
	Lower() Expr
	Trim() Expr
	TrimLeft() Expr
	TrimRight() Expr
	Length() Expr
	Match(pattern Expr) Expr
	Replace(old, new Expr) Expr
	PadLeft(length, pad Expr) Expr
	PadRight(length, pad Expr) Expr
	SplitPart(separator Expr, index Expr) Expr

	// Temporal methods
	Year() Expr
	Month() Expr
	Day() Expr
	Hour() Expr
	Minute() Expr
	Second() Expr
	TruncateToYear() Expr
	TruncateToMonth() Expr
	TruncateToDay() Expr
	TruncateToHour() Expr
	AddDays(days Expr) Expr
	AddHours(hours Expr) Expr
	AddMinutes(minutes Expr) Expr
	AddSeconds(seconds Expr) Expr
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

// Upper converts string to uppercase.
func (c *ColumnExpr) Upper() Expr {
	return NewUnaryExpr(c, "upper")
}

// Lower converts string to lowercase.
func (c *ColumnExpr) Lower() Expr {
	return NewUnaryExpr(c, "lower")
}

// Trim removes leading and trailing whitespace.
func (c *ColumnExpr) Trim() Expr {
	return NewUnaryExpr(c, "trim")
}

// TrimLeft removes leading whitespace.
func (c *ColumnExpr) TrimLeft() Expr {
	return NewUnaryExpr(c, "trim_left")
}

// TrimRight removes trailing whitespace.
func (c *ColumnExpr) TrimRight() Expr {
	return NewUnaryExpr(c, "trim_right")
}

// Length returns the length of the string.
func (c *ColumnExpr) Length() Expr {
	return NewUnaryExpr(c, "length")
}

// Match tests if string matches a regular expression pattern.
func (c *ColumnExpr) Match(pattern Expr) Expr {
	return NewBinaryExpr(c, pattern, "match")
}

// Replace replaces all occurrences of old with new in the string.
func (c *ColumnExpr) Replace(old, new Expr) Expr {
	return NewTernaryExpr(c, old, new, "replace")
}

// PadLeft pads the string on the left to the given length with the given pad character.
func (c *ColumnExpr) PadLeft(length, pad Expr) Expr {
	return NewTernaryExpr(c, length, pad, "pad_left")
}

// SplitPart splits a string by separator and returns the element at the given index (0-based).
func (c *ColumnExpr) SplitPart(separator Expr, index Expr) Expr {
	return NewTernaryExpr(c, separator, index, "split_part")
}

// PadRight pads the string on the right to the given length with the given pad character.
func (c *ColumnExpr) PadRight(length, pad Expr) Expr {
	return NewTernaryExpr(c, length, pad, "pad_right")
}

// Temporal operations for ColumnExpr

// Year extracts the year component from a timestamp column.
func (c *ColumnExpr) Year() Expr {
	return NewUnaryExpr(c, "year")
}

// Month extracts the month component (1-12) from a timestamp column.
func (c *ColumnExpr) Month() Expr {
	return NewUnaryExpr(c, "month")
}

// Day extracts the day component (1-31) from a timestamp column.
func (c *ColumnExpr) Day() Expr {
	return NewUnaryExpr(c, "day")
}

// Hour extracts the hour component (0-23) from a timestamp column.
func (c *ColumnExpr) Hour() Expr {
	return NewUnaryExpr(c, "hour")
}

// Minute extracts the minute component (0-59) from a timestamp column.
func (c *ColumnExpr) Minute() Expr {
	return NewUnaryExpr(c, "minute")
}

// Second extracts the second component (0-59) from a timestamp column.
func (c *ColumnExpr) Second() Expr {
	return NewUnaryExpr(c, "second")
}

// TruncateToYear truncates timestamp to the start of the year.
func (c *ColumnExpr) TruncateToYear() Expr {
	return NewUnaryExpr(c, "trunc_year")
}

// TruncateToMonth truncates timestamp to the start of the month.
func (c *ColumnExpr) TruncateToMonth() Expr {
	return NewUnaryExpr(c, "trunc_month")
}

// TruncateToDay truncates timestamp to the start of the day.
func (c *ColumnExpr) TruncateToDay() Expr {
	return NewUnaryExpr(c, "trunc_day")
}

// TruncateToHour truncates timestamp to the start of the hour.
func (c *ColumnExpr) TruncateToHour() Expr {
	return NewUnaryExpr(c, "trunc_hour")
}

// AddDays adds a number of days to a timestamp column.
func (c *ColumnExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(c, days, "add_days")
}

// AddHours adds a number of hours to a timestamp column.
func (c *ColumnExpr) AddHours(hours Expr) Expr {
	return NewBinaryExpr(c, hours, "add_hours")
}

// AddMinutes adds a number of minutes to a timestamp column.
func (c *ColumnExpr) AddMinutes(minutes Expr) Expr {
	return NewBinaryExpr(c, minutes, "add_minutes")
}

// AddSeconds adds a number of seconds to a timestamp column.
func (c *ColumnExpr) AddSeconds(seconds Expr) Expr {
	return NewBinaryExpr(c, seconds, "add_seconds")
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
	numRows := int(df.NumRows())
	pool := memory.NewGoAllocator()

	// Create an array filled with the literal value
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

func (l *LiteralExpr) Upper() Expr {
	return NewUnaryExpr(l, "upper")
}

func (l *LiteralExpr) Lower() Expr {
	return NewUnaryExpr(l, "lower")
}

func (l *LiteralExpr) Trim() Expr {
	return NewUnaryExpr(l, "trim")
}

func (l *LiteralExpr) TrimLeft() Expr {
	return NewUnaryExpr(l, "trim_left")
}

func (l *LiteralExpr) TrimRight() Expr {
	return NewUnaryExpr(l, "trim_right")
}

func (l *LiteralExpr) Length() Expr {
	return NewUnaryExpr(l, "length")
}

func (l *LiteralExpr) Match(pattern Expr) Expr {
	return NewBinaryExpr(l, pattern, "match")
}

func (l *LiteralExpr) Replace(old, new Expr) Expr {
	return NewTernaryExpr(l, old, new, "replace")
}

func (l *LiteralExpr) PadLeft(length, pad Expr) Expr {
	return NewTernaryExpr(l, length, pad, "pad_left")
}

func (l *LiteralExpr) PadRight(length, pad Expr) Expr {
	return NewTernaryExpr(l, length, pad, "pad_right")
}

func (l *LiteralExpr) SplitPart(separator, index Expr) Expr {
	return NewTernaryExpr(l, separator, index, "split_part")
}

// Temporal methods for literals
func (l *LiteralExpr) Year() Expr {
	return NewUnaryExpr(l, "year")
}

func (l *LiteralExpr) Month() Expr {
	return NewUnaryExpr(l, "month")
}

func (l *LiteralExpr) Day() Expr {
	return NewUnaryExpr(l, "day")
}

func (l *LiteralExpr) Hour() Expr {
	return NewUnaryExpr(l, "hour")
}

func (l *LiteralExpr) Minute() Expr {
	return NewUnaryExpr(l, "minute")
}

func (l *LiteralExpr) Second() Expr {
	return NewUnaryExpr(l, "second")
}

func (l *LiteralExpr) TruncateToYear() Expr {
	return NewUnaryExpr(l, "trunc_year")
}

func (l *LiteralExpr) TruncateToMonth() Expr {
	return NewUnaryExpr(l, "trunc_month")
}

func (l *LiteralExpr) TruncateToDay() Expr {
	return NewUnaryExpr(l, "trunc_day")
}

func (l *LiteralExpr) TruncateToHour() Expr {
	return NewUnaryExpr(l, "trunc_hour")
}

func (l *LiteralExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(l, days, "add_days")
}

func (l *LiteralExpr) AddHours(hours Expr) Expr {
	return NewBinaryExpr(l, hours, "add_hours")
}

func (l *LiteralExpr) AddMinutes(minutes Expr) Expr {
	return NewBinaryExpr(l, minutes, "add_minutes")
}

func (l *LiteralExpr) AddSeconds(seconds Expr) Expr {
	return NewBinaryExpr(l, seconds, "add_seconds")
}

// BinaryExpr represents binary operations between two expressions.
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
	// Evaluate left and right operands
	leftArray, err := b.left.Evaluate(df)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate left operand: %w", err)
	}
	defer leftArray.Release()

	rightArray, err := b.right.Evaluate(df)
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
	case "match":
		return b.evaluateMatch(leftArray, rightArray)
	case "add_days":
		return b.evaluateAddDays(leftArray, rightArray)
	case "add_hours":
		return b.evaluateAddHours(leftArray, rightArray)
	case "add_minutes":
		return b.evaluateAddMinutes(leftArray, rightArray)
	case "add_seconds":
		return b.evaluateAddSeconds(leftArray, rightArray)
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

func (b *BinaryExpr) Upper() Expr {
	return NewUnaryExpr(b, "upper")
}

func (b *BinaryExpr) Lower() Expr {
	return NewUnaryExpr(b, "lower")
}

func (b *BinaryExpr) Trim() Expr {
	return NewUnaryExpr(b, "trim")
}

func (b *BinaryExpr) TrimLeft() Expr {
	return NewUnaryExpr(b, "trim_left")
}

func (b *BinaryExpr) TrimRight() Expr {
	return NewUnaryExpr(b, "trim_right")
}

func (b *BinaryExpr) Length() Expr {
	return NewUnaryExpr(b, "length")
}

func (b *BinaryExpr) Match(pattern Expr) Expr {
	return NewBinaryExpr(b, pattern, "match")
}

func (b *BinaryExpr) Replace(old, new Expr) Expr {
	return NewTernaryExpr(b, old, new, "replace")
}

func (b *BinaryExpr) PadLeft(length, pad Expr) Expr {
	return NewTernaryExpr(b, length, pad, "pad_left")
}

func (b *BinaryExpr) PadRight(length, pad Expr) Expr {
	return NewTernaryExpr(b, length, pad, "pad_right")
}

func (b *BinaryExpr) SplitPart(separator, index Expr) Expr {
	return NewTernaryExpr(b, separator, index, "split_part")
}

// Temporal methods for binary expressions
func (b *BinaryExpr) Year() Expr {
	return NewUnaryExpr(b, "year")
}

func (b *BinaryExpr) Month() Expr {
	return NewUnaryExpr(b, "month")
}

func (b *BinaryExpr) Day() Expr {
	return NewUnaryExpr(b, "day")
}

func (b *BinaryExpr) Hour() Expr {
	return NewUnaryExpr(b, "hour")
}

func (b *BinaryExpr) Minute() Expr {
	return NewUnaryExpr(b, "minute")
}

func (b *BinaryExpr) Second() Expr {
	return NewUnaryExpr(b, "second")
}

func (b *BinaryExpr) TruncateToYear() Expr {
	return NewUnaryExpr(b, "trunc_year")
}

func (b *BinaryExpr) TruncateToMonth() Expr {
	return NewUnaryExpr(b, "trunc_month")
}

func (b *BinaryExpr) TruncateToDay() Expr {
	return NewUnaryExpr(b, "trunc_day")
}

func (b *BinaryExpr) TruncateToHour() Expr {
	return NewUnaryExpr(b, "trunc_hour")
}

func (b *BinaryExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(b, days, "add_days")
}

func (b *BinaryExpr) AddHours(hours Expr) Expr {
	return NewBinaryExpr(b, hours, "add_hours")
}

func (b *BinaryExpr) AddMinutes(minutes Expr) Expr {
	return NewBinaryExpr(b, minutes, "add_minutes")
}

func (b *BinaryExpr) AddSeconds(seconds Expr) Expr {
	return NewBinaryExpr(b, seconds, "add_seconds")
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
	// Evaluate operand
	operandArray, err := u.operand.Evaluate(df)
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
	case "hour":
		return u.evaluateHour(operandArray)
	case "minute":
		return u.evaluateMinute(operandArray)
	case "second":
		return u.evaluateSecond(operandArray)
	case "trunc_year":
		return u.evaluateTruncateToYear(operandArray)
	case "trunc_month":
		return u.evaluateTruncateToMonth(operandArray)
	case "trunc_day":
		return u.evaluateTruncateToDay(operandArray)
	case "trunc_hour":
		return u.evaluateTruncateToHour(operandArray)
	case "upper":
		return u.evaluateUpper(operandArray)
	case "lower":
		return u.evaluateLower(operandArray)
	case "trim":
		return u.evaluateTrim(operandArray)
	case "trim_left":
		return u.evaluateTrimLeft(operandArray)
	case "trim_right":
		return u.evaluateTrimRight(operandArray)
	case "length":
		return u.evaluateLength(operandArray)
	default:
		return nil, fmt.Errorf("unsupported unary operator: %s", u.operator)
	}
}

// Name implements Expr.Name for unary operations.
func (u *UnaryExpr) Name() string {
	return fmt.Sprintf("%s(%s)", u.operator, u.operand.Name())
}

// String implements Expr.String for unary operations.
func (u *UnaryExpr) String() string {
	return fmt.Sprintf("%s(%s)", u.operator, u.operand.String())
}

// Fluent methods for UnaryExpr (delegate to BinaryExpr)
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

// String manipulation methods
func (u *UnaryExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(u, substring, "contains")
}

func (u *UnaryExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(u, prefix, "starts_with")
}

func (u *UnaryExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(u, suffix, "ends_with")
}

func (u *UnaryExpr) Upper() Expr {
	return NewUnaryExpr(u, "upper")
}

func (u *UnaryExpr) Lower() Expr {
	return NewUnaryExpr(u, "lower")
}

func (u *UnaryExpr) Trim() Expr {
	return NewUnaryExpr(u, "trim")
}

func (u *UnaryExpr) TrimLeft() Expr {
	return NewUnaryExpr(u, "trim_left")
}

func (u *UnaryExpr) TrimRight() Expr {
	return NewUnaryExpr(u, "trim_right")
}

func (u *UnaryExpr) Length() Expr {
	return NewUnaryExpr(u, "length")
}

func (u *UnaryExpr) Match(pattern Expr) Expr {
	return NewBinaryExpr(u, pattern, "match")
}

func (u *UnaryExpr) Replace(old, new Expr) Expr {
	return NewTernaryExpr(u, old, new, "replace")
}

func (u *UnaryExpr) PadLeft(length, pad Expr) Expr {
	return NewTernaryExpr(u, length, pad, "pad_left")
}

func (u *UnaryExpr) PadRight(length, pad Expr) Expr {
	return NewTernaryExpr(u, length, pad, "pad_right")
}

func (u *UnaryExpr) SplitPart(separator, index Expr) Expr {
	return NewTernaryExpr(u, separator, index, "split_part")
}

// Temporal methods for UnaryExpr
func (u *UnaryExpr) Year() Expr {
	return NewUnaryExpr(u, "year")
}

func (u *UnaryExpr) Month() Expr {
	return NewUnaryExpr(u, "month")
}

func (u *UnaryExpr) Day() Expr {
	return NewUnaryExpr(u, "day")
}

func (u *UnaryExpr) Hour() Expr {
	return NewUnaryExpr(u, "hour")
}

func (u *UnaryExpr) Minute() Expr {
	return NewUnaryExpr(u, "minute")
}

func (u *UnaryExpr) Second() Expr {
	return NewUnaryExpr(u, "second")
}

func (u *UnaryExpr) TruncateToYear() Expr {
	return NewUnaryExpr(u, "trunc_year")
}

func (u *UnaryExpr) TruncateToMonth() Expr {
	return NewUnaryExpr(u, "trunc_month")
}

func (u *UnaryExpr) TruncateToDay() Expr {
	return NewUnaryExpr(u, "trunc_day")
}

func (u *UnaryExpr) TruncateToHour() Expr {
	return NewUnaryExpr(u, "trunc_hour")
}

func (u *UnaryExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(u, days, "add_days")
}

func (u *UnaryExpr) AddHours(hours Expr) Expr {
	return NewBinaryExpr(u, hours, "add_hours")
}

func (u *UnaryExpr) AddMinutes(minutes Expr) Expr {
	return NewBinaryExpr(u, minutes, "add_minutes")
}

func (u *UnaryExpr) AddSeconds(seconds Expr) Expr {
	return NewBinaryExpr(u, seconds, "add_seconds")
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
	default:
		// Default to string for unknown types
		return arrow.BinaryTypes.String
	}
}

// ====================
// Temporal Operations
// ====================

// evaluateYear extracts the year component from timestamps
func (u *UnaryExpr) evaluateYear(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("year operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			builder.Append(int64(t.Year()))
		}
	}

	return builder.NewArray(), nil
}

// evaluateMonth extracts the month component (1-12) from timestamps
func (u *UnaryExpr) evaluateMonth(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("month operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			builder.Append(int64(t.Month()))
		}
	}

	return builder.NewArray(), nil
}

// evaluateDay extracts the day component (1-31) from timestamps
func (u *UnaryExpr) evaluateDay(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("day operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			builder.Append(int64(t.Day()))
		}
	}

	return builder.NewArray(), nil
}

// evaluateHour extracts the hour component (0-23) from timestamps
func (u *UnaryExpr) evaluateHour(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("hour operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			builder.Append(int64(t.Hour()))
		}
	}

	return builder.NewArray(), nil
}

// evaluateMinute extracts the minute component (0-59) from timestamps
func (u *UnaryExpr) evaluateMinute(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("minute operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			builder.Append(int64(t.Minute()))
		}
	}

	return builder.NewArray(), nil
}

// evaluateSecond extracts the second component (0-59) from timestamps
func (u *UnaryExpr) evaluateSecond(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("second operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			builder.Append(int64(t.Second()))
		}
	}

	return builder.NewArray(), nil
}

// evaluateTruncateToYear truncates timestamps to the start of the year
func (u *UnaryExpr) evaluateTruncateToYear(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("truncate operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			truncated := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
			ts, err := arrow.TimestampFromTime(truncated, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateTruncateToMonth truncates timestamps to the start of the month
func (u *UnaryExpr) evaluateTruncateToMonth(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("truncate operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			truncated := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
			ts, err := arrow.TimestampFromTime(truncated, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateTruncateToDay truncates timestamps to the start of the day
func (u *UnaryExpr) evaluateTruncateToDay(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("truncate operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			truncated := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
			ts, err := arrow.TimestampFromTime(truncated, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateTruncateToHour truncates timestamps to the start of the hour
func (u *UnaryExpr) evaluateTruncateToHour(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("truncate operation requires timestamp type, got %s", arr.DataType())
	}

	tsArray := arr.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if tsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			truncated := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
			ts, err := arrow.TimestampFromTime(truncated, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateAddDays adds a number of days to timestamps
func (b *BinaryExpr) evaluateAddDays(left, right arrow.Array) (arrow.Array, error) {
	if left.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("add_days left operand must be timestamp, got %s", left.DataType())
	}
	if right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("add_days right operand must be int64, got %s", right.DataType())
	}

	tsArray := left.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	daysArray := right.(*array.Int64)

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if tsArray.IsNull(i) || daysArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			days := daysArray.Value(i)
			newTime := t.AddDate(0, 0, int(days))
			ts, err := arrow.TimestampFromTime(newTime, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateAddHours adds a number of hours to timestamps
func (b *BinaryExpr) evaluateAddHours(left, right arrow.Array) (arrow.Array, error) {
	if left.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("add_hours left operand must be timestamp, got %s", left.DataType())
	}
	if right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("add_hours right operand must be int64, got %s", right.DataType())
	}

	tsArray := left.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	hoursArray := right.(*array.Int64)

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if tsArray.IsNull(i) || hoursArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			hours := hoursArray.Value(i)
			newTime := t.Add(time.Duration(hours) * time.Hour)
			ts, err := arrow.TimestampFromTime(newTime, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateAddMinutes adds a number of minutes to timestamps
func (b *BinaryExpr) evaluateAddMinutes(left, right arrow.Array) (arrow.Array, error) {
	if left.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("add_minutes left operand must be timestamp, got %s", left.DataType())
	}
	if right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("add_minutes right operand must be int64, got %s", right.DataType())
	}

	tsArray := left.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	minutesArray := right.(*array.Int64)

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if tsArray.IsNull(i) || minutesArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			minutes := minutesArray.Value(i)
			newTime := t.Add(time.Duration(minutes) * time.Minute)
			ts, err := arrow.TimestampFromTime(newTime, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// evaluateAddSeconds adds a number of seconds to timestamps
func (b *BinaryExpr) evaluateAddSeconds(left, right arrow.Array) (arrow.Array, error) {
	if left.DataType().ID() != arrow.TIMESTAMP {
		return nil, fmt.Errorf("add_seconds left operand must be timestamp, got %s", left.DataType())
	}
	if right.DataType().ID() != arrow.INT64 {
		return nil, fmt.Errorf("add_seconds right operand must be int64, got %s", right.DataType())
	}

	tsArray := left.(*array.Timestamp)
	tsType := tsArray.DataType().(*arrow.TimestampType)
	toTime, err := tsType.GetToTimeFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get time conversion function: %w", err)
	}

	secondsArray := right.(*array.Int64)

	pool := memory.NewGoAllocator()
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < left.Len(); i++ {
		if tsArray.IsNull(i) || secondsArray.IsNull(i) {
			builder.AppendNull()
		} else {
			t := toTime(tsArray.Value(i))
			seconds := secondsArray.Value(i)
			newTime := t.Add(time.Duration(seconds) * time.Second)
			ts, err := arrow.TimestampFromTime(newTime, tsType.Unit)
			if err != nil {
				return nil, fmt.Errorf("failed to convert time to timestamp: %w", err)
			}
			builder.Append(ts)
		}
	}

	return builder.NewArray(), nil
}

// ====================
// String Operations
// ====================

// evaluateUpper converts strings to uppercase
func (u *UnaryExpr) evaluateUpper(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("upper operation requires string type, got %s", arr.DataType())
	}

	strArray := arr.(*array.String)
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if strArray.IsNull(i) {
			builder.AppendNull()
		} else {
			builder.Append(strings.ToUpper(strArray.Value(i)))
		}
	}

	return builder.NewArray(), nil
}

// evaluateLower converts strings to lowercase
func (u *UnaryExpr) evaluateLower(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("lower operation requires string type, got %s", arr.DataType())
	}

	strArray := arr.(*array.String)
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if strArray.IsNull(i) {
			builder.AppendNull()
		} else {
			builder.Append(strings.ToLower(strArray.Value(i)))
		}
	}

	return builder.NewArray(), nil
}

// evaluateTrim removes leading and trailing whitespace
func (u *UnaryExpr) evaluateTrim(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("trim operation requires string type, got %s", arr.DataType())
	}

	strArray := arr.(*array.String)
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if strArray.IsNull(i) {
			builder.AppendNull()
		} else {
			builder.Append(strings.TrimSpace(strArray.Value(i)))
		}
	}

	return builder.NewArray(), nil
}

// evaluateTrimLeft removes leading whitespace
func (u *UnaryExpr) evaluateTrimLeft(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("trim_left operation requires string type, got %s", arr.DataType())
	}

	strArray := arr.(*array.String)
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if strArray.IsNull(i) {
			builder.AppendNull()
		} else {
			builder.Append(strings.TrimLeft(strArray.Value(i), " \t\n\r"))
		}
	}

	return builder.NewArray(), nil
}

// evaluateTrimRight removes trailing whitespace
func (u *UnaryExpr) evaluateTrimRight(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("trim_right operation requires string type, got %s", arr.DataType())
	}

	strArray := arr.(*array.String)
	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if strArray.IsNull(i) {
			builder.AppendNull()
		} else {
			builder.Append(strings.TrimRight(strArray.Value(i), " \t\n\r"))
		}
	}

	return builder.NewArray(), nil
}

// evaluateLength returns the length of strings
func (u *UnaryExpr) evaluateLength(arr arrow.Array) (arrow.Array, error) {
	if arr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("length operation requires string type, got %s", arr.DataType())
	}

	strArray := arr.(*array.String)
	pool := memory.NewGoAllocator()
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := 0; i < arr.Len(); i++ {
		if strArray.IsNull(i) {
			builder.AppendNull()
		} else {
			builder.Append(int64(len(strArray.Value(i))))
		}
	}

	return builder.NewArray(), nil
}

// evaluateMatch tests if strings match a regular expression pattern
func (b *BinaryExpr) evaluateMatch(left, right arrow.Array) (arrow.Array, error) {
	if left.Len() != right.Len() {
		return nil, fmt.Errorf("array length mismatch: %d vs %d", left.Len(), right.Len())
	}

	// Both operands must be strings
	if left.DataType().ID() != arrow.STRING || right.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("match operation requires string operands, got %s and %s", left.DataType(), right.DataType())
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

	// Cache compiled regexes for efficiency
	regexCache := make(map[string]*regexp.Regexp)

	for i := 0; i < left.Len(); i++ {
		if leftStr.IsNull(i) || rightStr.IsNull(i) {
			builder.AppendNull()
		} else {
			leftVal := leftStr.Value(i)
			pattern := rightStr.Value(i)

			// Check if regex is already compiled
			re, exists := regexCache[pattern]
			if !exists {
				var err error
				re, err = regexp.Compile(pattern)
				if err != nil {
					return nil, fmt.Errorf("invalid regex pattern '%s': %w", pattern, err)
				}
				regexCache[pattern] = re
			}

			builder.Append(re.MatchString(leftVal))
		}
	}

	return builder.NewArray(), nil
}

// ====================
// Ternary Expression
// ====================

// TernaryExpr represents operations that take three expression operands.
// Used for Replace(operand, old, new) and PadLeft/PadRight(operand, length, pad).
type TernaryExpr struct {
	first    Expr
	second   Expr
	third    Expr
	operator string
}

// NewTernaryExpr creates a new ternary expression.
func NewTernaryExpr(first, second, third Expr, operator string) Expr {
	return &TernaryExpr{
		first:    first,
		second:   second,
		third:    third,
		operator: operator,
	}
}

// Evaluate implements Expr.Evaluate for ternary operations.
func (te *TernaryExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	firstArray, err := te.first.Evaluate(df)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate first operand: %w", err)
	}
	defer firstArray.Release()

	secondArray, err := te.second.Evaluate(df)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate second operand: %w", err)
	}
	defer secondArray.Release()

	thirdArray, err := te.third.Evaluate(df)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate third operand: %w", err)
	}
	defer thirdArray.Release()

	switch te.operator {
	case "replace":
		return te.evaluateReplace(firstArray, secondArray, thirdArray)
	case "pad_left":
		return te.evaluatePadLeft(firstArray, secondArray, thirdArray)
	case "pad_right":
		return te.evaluatePadRight(firstArray, secondArray, thirdArray)
	case "split_part":
		return te.evaluateSplitPart(firstArray, secondArray, thirdArray)
	default:
		return nil, fmt.Errorf("unsupported ternary operator: %s", te.operator)
	}
}

// Name implements Expr.Name for ternary operations.
func (te *TernaryExpr) Name() string {
	return fmt.Sprintf("%s(%s, %s, %s)", te.operator, te.first.Name(), te.second.Name(), te.third.Name())
}

// String implements Expr.String for ternary operations.
func (te *TernaryExpr) String() string {
	return fmt.Sprintf("%s(%s, %s, %s)", te.operator, te.first.String(), te.second.String(), te.third.String())
}

// Fluent methods for TernaryExpr
func (te *TernaryExpr) Add(other Expr) Expr {
	return NewBinaryExpr(te, other, "add")
}

func (te *TernaryExpr) Sub(other Expr) Expr {
	return NewBinaryExpr(te, other, "subtract")
}

func (te *TernaryExpr) Mul(other Expr) Expr {
	return NewBinaryExpr(te, other, "multiply")
}

func (te *TernaryExpr) Div(other Expr) Expr {
	return NewBinaryExpr(te, other, "divide")
}

func (te *TernaryExpr) Gt(other Expr) Expr {
	return NewBinaryExpr(te, other, "greater")
}

func (te *TernaryExpr) Lt(other Expr) Expr {
	return NewBinaryExpr(te, other, "less")
}

func (te *TernaryExpr) Eq(other Expr) Expr {
	return NewBinaryExpr(te, other, "equal")
}

// String manipulation methods for TernaryExpr
func (te *TernaryExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(te, substring, "contains")
}

func (te *TernaryExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(te, prefix, "starts_with")
}

func (te *TernaryExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(te, suffix, "ends_with")
}

func (te *TernaryExpr) Upper() Expr {
	return NewUnaryExpr(te, "upper")
}

func (te *TernaryExpr) Lower() Expr {
	return NewUnaryExpr(te, "lower")
}

func (te *TernaryExpr) Trim() Expr {
	return NewUnaryExpr(te, "trim")
}

func (te *TernaryExpr) TrimLeft() Expr {
	return NewUnaryExpr(te, "trim_left")
}

func (te *TernaryExpr) TrimRight() Expr {
	return NewUnaryExpr(te, "trim_right")
}

func (te *TernaryExpr) Length() Expr {
	return NewUnaryExpr(te, "length")
}

func (te *TernaryExpr) Match(pattern Expr) Expr {
	return NewBinaryExpr(te, pattern, "match")
}

func (te *TernaryExpr) Replace(old, new Expr) Expr {
	return NewTernaryExpr(te, old, new, "replace")
}

func (te *TernaryExpr) PadLeft(length, pad Expr) Expr {
	return NewTernaryExpr(te, length, pad, "pad_left")
}

func (te *TernaryExpr) PadRight(length, pad Expr) Expr {
	return NewTernaryExpr(te, length, pad, "pad_right")
}

func (te *TernaryExpr) SplitPart(separator, index Expr) Expr {
	return NewTernaryExpr(te, separator, index, "split_part")
}

// Temporal methods for TernaryExpr
func (te *TernaryExpr) Year() Expr {
	return NewUnaryExpr(te, "year")
}

func (te *TernaryExpr) Month() Expr {
	return NewUnaryExpr(te, "month")
}

func (te *TernaryExpr) Day() Expr {
	return NewUnaryExpr(te, "day")
}

func (te *TernaryExpr) Hour() Expr {
	return NewUnaryExpr(te, "hour")
}

func (te *TernaryExpr) Minute() Expr {
	return NewUnaryExpr(te, "minute")
}

func (te *TernaryExpr) Second() Expr {
	return NewUnaryExpr(te, "second")
}

func (te *TernaryExpr) TruncateToYear() Expr {
	return NewUnaryExpr(te, "trunc_year")
}

func (te *TernaryExpr) TruncateToMonth() Expr {
	return NewUnaryExpr(te, "trunc_month")
}

func (te *TernaryExpr) TruncateToDay() Expr {
	return NewUnaryExpr(te, "trunc_day")
}

func (te *TernaryExpr) TruncateToHour() Expr {
	return NewUnaryExpr(te, "trunc_hour")
}

func (te *TernaryExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(te, days, "add_days")
}

func (te *TernaryExpr) AddHours(hours Expr) Expr {
	return NewBinaryExpr(te, hours, "add_hours")
}

func (te *TernaryExpr) AddMinutes(minutes Expr) Expr {
	return NewBinaryExpr(te, minutes, "add_minutes")
}

func (te *TernaryExpr) AddSeconds(seconds Expr) Expr {
	return NewBinaryExpr(te, seconds, "add_seconds")
}

// evaluateReplace replaces all occurrences of old with new in each string element.
func (te *TernaryExpr) evaluateReplace(operand, oldStr, newStr arrow.Array) (arrow.Array, error) {
	if operand.Len() != oldStr.Len() || operand.Len() != newStr.Len() {
		return nil, fmt.Errorf("array length mismatch in replace operation")
	}

	if operand.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("replace operation requires string type, got %s", operand.DataType())
	}
	if oldStr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("replace old parameter requires string type, got %s", oldStr.DataType())
	}
	if newStr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("replace new parameter requires string type, got %s", newStr.DataType())
	}

	strArray := operand.(*array.String)
	oldArray := oldStr.(*array.String)
	newArray := newStr.(*array.String)

	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < operand.Len(); i++ {
		if strArray.IsNull(i) || oldArray.IsNull(i) || newArray.IsNull(i) {
			builder.AppendNull()
		} else {
			result := strings.ReplaceAll(strArray.Value(i), oldArray.Value(i), newArray.Value(i))
			builder.Append(result)
		}
	}

	return builder.NewArray(), nil
}

// evaluatePadLeft pads strings on the left to the given length with the given pad string.
func (te *TernaryExpr) evaluatePadLeft(operand, lengthArr, padArr arrow.Array) (arrow.Array, error) {
	if operand.Len() != lengthArr.Len() || operand.Len() != padArr.Len() {
		return nil, fmt.Errorf("array length mismatch in pad_left operation")
	}

	if operand.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("pad_left operation requires string type, got %s", operand.DataType())
	}
	if padArr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("pad_left pad parameter requires string type, got %s", padArr.DataType())
	}

	strArray := operand.(*array.String)
	padStrArray := padArr.(*array.String)

	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < operand.Len(); i++ {
		if strArray.IsNull(i) || lengthArr.IsNull(i) || padStrArray.IsNull(i) {
			builder.AppendNull()
		} else {
			s := strArray.Value(i)
			targetLen := extractInt64Value(lengthArr, i)
			padChar := padStrArray.Value(i)

			if len(padChar) == 0 || int64(len(s)) >= targetLen {
				builder.Append(s)
			} else {
				needed := int(targetLen) - len(s)
				padding := strings.Repeat(padChar, (needed/len(padChar))+1)
				builder.Append(padding[:needed] + s)
			}
		}
	}

	return builder.NewArray(), nil
}

// evaluatePadRight pads strings on the right to the given length with the given pad string.
func (te *TernaryExpr) evaluatePadRight(operand, lengthArr, padArr arrow.Array) (arrow.Array, error) {
	if operand.Len() != lengthArr.Len() || operand.Len() != padArr.Len() {
		return nil, fmt.Errorf("array length mismatch in pad_right operation")
	}

	if operand.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("pad_right operation requires string type, got %s", operand.DataType())
	}
	if padArr.DataType().ID() != arrow.STRING {
		return nil, fmt.Errorf("pad_right pad parameter requires string type, got %s", padArr.DataType())
	}

	strArray := operand.(*array.String)
	padStrArray := padArr.(*array.String)

	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < operand.Len(); i++ {
		if strArray.IsNull(i) || lengthArr.IsNull(i) || padStrArray.IsNull(i) {
			builder.AppendNull()
		} else {
			s := strArray.Value(i)
			targetLen := extractInt64Value(lengthArr, i)
			padChar := padStrArray.Value(i)

			if len(padChar) == 0 || int64(len(s)) >= targetLen {
				builder.Append(s)
			} else {
				needed := int(targetLen) - len(s)
				padding := strings.Repeat(padChar, (needed/len(padChar))+1)
				builder.Append(s + padding[:needed])
			}
		}
	}

	return builder.NewArray(), nil
}

// extractInt64Value extracts an int64 value from an Arrow array at a given index.
// Supports Int64 and Float64 arrays (float64 is truncated to int64).
func extractInt64Value(arr arrow.Array, i int) int64 {
	switch a := arr.(type) {
	case *array.Int64:
		return a.Value(i)
	case *array.Float64:
		return int64(a.Value(i))
	case *array.Int32:
		return int64(a.Value(i))
	default:
		return 0
	}
}

// evaluateSplitPart splits strings by a separator and returns the part at the given index.
func (te *TernaryExpr) evaluateSplitPart(operand, separatorArr, indexArr arrow.Array) (arrow.Array, error) {
	strArr, ok := operand.(*array.String)
	if !ok {
		return nil, fmt.Errorf("split_part requires string operand, got %s", operand.DataType())
	}

	sepArr, ok := separatorArr.(*array.String)
	if !ok {
		return nil, fmt.Errorf("split_part separator must be string, got %s", separatorArr.DataType())
	}

	pool := memory.NewGoAllocator()
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for i := 0; i < strArr.Len(); i++ {
		if strArr.IsNull(i) || sepArr.IsNull(i) {
			builder.AppendNull()
			continue
		}

		s := strArr.Value(i)
		sep := sepArr.Value(i)
		idx := extractInt64Value(indexArr, i)

		parts := strings.Split(s, sep)
		if int(idx) >= 0 && int(idx) < len(parts) {
			builder.Append(parts[int(idx)])
		} else {
			builder.Append("")
		}
	}

	return builder.NewArray(), nil
}
