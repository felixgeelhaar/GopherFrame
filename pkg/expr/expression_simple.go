// Package expr provides a simplified expression engine for DataFrame operations.
// This initial implementation focuses on basic column references and literals
// without complex compute operations.
package expr

import (
	"fmt"

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
