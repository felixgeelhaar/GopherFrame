package expr

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/compute"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/arrow/scalar"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// VectorizedExpr represents an expression that uses Arrow compute kernels
type VectorizedExpr struct {
	operands []Expr
	kernel   string
	options  compute.FunctionOptions
}

// NewVectorizedExpr creates a new vectorized expression
func NewVectorizedExpr(operands []Expr, kernel string, options compute.FunctionOptions) Expr {
	return &VectorizedExpr{
		operands: operands,
		kernel:   kernel,
		options:  options,
	}
}

// Evaluate implements Expr.Evaluate for vectorized operations
func (v *VectorizedExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	return v.EvaluateWithPool(df, memory.NewGoAllocator())
}

// EvaluateWithPool implements Expr.EvaluateWithPool for vectorized operations
func (v *VectorizedExpr) EvaluateWithPool(df *core.DataFrame, pool memory.Allocator) (arrow.Array, error) {
	ctx := context.Background()

	// Evaluate all operands
	operandArrays := make([]arrow.Array, len(v.operands))
	for i, operand := range v.operands {
		arr, err := operand.EvaluateWithPool(df, pool)
		if err != nil {
			// Clean up already evaluated arrays
			for j := 0; j < i; j++ {
				operandArrays[j].Release()
			}
			return nil, fmt.Errorf("failed to evaluate operand %d: %w", i, err)
		}
		operandArrays[i] = arr
	}

	// Ensure cleanup of operand arrays
	defer func() {
		for _, arr := range operandArrays {
			if arr != nil {
				arr.Release()
			}
		}
	}()

	// Create Arrow compute datums
	datums := make([]compute.Datum, len(operandArrays))
	for i, arr := range operandArrays {
		datums[i] = compute.NewDatum(arr)
	}

	// Execute the compute kernel
	result, err := compute.CallFunction(ctx, v.kernel, v.options, datums...)
	if err != nil {
		return nil, fmt.Errorf("compute kernel %s failed: %w", v.kernel, err)
	}

	// Extract result array
	switch resultDatum := result.(type) {
	case *compute.ArrayDatum:
		// Create array from ArrayData
		arr := array.MakeFromData(resultDatum.Value)
		arr.Retain() // Retain for caller
		return arr, nil
	case *compute.ScalarDatum:
		// For aggregations, we need to convert scalar to array
		return v.scalarToArray(resultDatum.Value, pool)
	default:
		return nil, fmt.Errorf("compute kernel %s returned unsupported result type", v.kernel)
	}
}

// Name implements Expr.Name for vectorized expressions
func (v *VectorizedExpr) Name() string {
	if len(v.operands) == 1 {
		return fmt.Sprintf("vectorized_%s(%s)", v.kernel, v.operands[0].Name())
	}
	names := make([]string, len(v.operands))
	for i, op := range v.operands {
		names[i] = op.Name()
	}
	return fmt.Sprintf("vectorized_%s(%v)", v.kernel, names)
}

// String implements Expr.String for vectorized expressions
func (v *VectorizedExpr) String() string {
	return v.Name()
}

// scalarToArray converts a scalar result to a single-element array
func (v *VectorizedExpr) scalarToArray(scal scalar.Scalar, pool memory.Allocator) (arrow.Array, error) {
	switch s := scal.(type) {
	case *scalar.Float64:
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()
		if s.IsValid() {
			builder.Append(s.Value)
		} else {
			builder.AppendNull()
		}
		return builder.NewArray(), nil
	case *scalar.Int64:
		builder := array.NewInt64Builder(pool)
		defer builder.Release()
		if s.IsValid() {
			builder.Append(s.Value)
		} else {
			builder.AppendNull()
		}
		return builder.NewArray(), nil
	case *scalar.Boolean:
		builder := array.NewBooleanBuilder(pool)
		defer builder.Release()
		if s.IsValid() {
			builder.Append(s.Value)
		} else {
			builder.AppendNull()
		}
		return builder.NewArray(), nil
	default:
		return nil, fmt.Errorf("unsupported scalar type: %T", scal)
	}
}

// Implement all interface methods for VectorizedExpr
func (v *VectorizedExpr) Add(other Expr) Expr {
	return NewBinaryExpr(v, other, "add")
}

func (v *VectorizedExpr) Sub(other Expr) Expr {
	return NewBinaryExpr(v, other, "subtract")
}

func (v *VectorizedExpr) Mul(other Expr) Expr {
	return NewBinaryExpr(v, other, "multiply")
}

func (v *VectorizedExpr) Div(other Expr) Expr {
	return NewBinaryExpr(v, other, "divide")
}

func (v *VectorizedExpr) Gt(other Expr) Expr {
	return NewBinaryExpr(v, other, "greater")
}

func (v *VectorizedExpr) Lt(other Expr) Expr {
	return NewBinaryExpr(v, other, "less")
}

func (v *VectorizedExpr) Eq(other Expr) Expr {
	return NewBinaryExpr(v, other, "equal")
}

func (v *VectorizedExpr) Contains(substring Expr) Expr {
	return NewBinaryExpr(v, substring, "contains")
}

func (v *VectorizedExpr) StartsWith(prefix Expr) Expr {
	return NewBinaryExpr(v, prefix, "starts_with")
}

func (v *VectorizedExpr) EndsWith(suffix Expr) Expr {
	return NewBinaryExpr(v, suffix, "ends_with")
}

func (v *VectorizedExpr) AddDays(days Expr) Expr {
	return NewBinaryExpr(v, days, "add_days")
}

func (v *VectorizedExpr) AddMonths(months Expr) Expr {
	return NewBinaryExpr(v, months, "add_months")
}

func (v *VectorizedExpr) AddYears(years Expr) Expr {
	return NewBinaryExpr(v, years, "add_years")
}

func (v *VectorizedExpr) Year() Expr {
	return NewUnaryExpr(v, "year")
}

func (v *VectorizedExpr) Month() Expr {
	return NewUnaryExpr(v, "month")
}

func (v *VectorizedExpr) Day() Expr {
	return NewUnaryExpr(v, "day")
}

func (v *VectorizedExpr) DayOfWeek() Expr {
	return NewUnaryExpr(v, "day_of_week")
}

func (v *VectorizedExpr) DateTrunc(unit string) Expr {
	return NewUnaryExpr(v, "date_trunc_"+unit)
}

// Vectorized operations for VectorizedExpr
func (v *VectorizedExpr) VectorizedAdd(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "add", nil)
}

func (v *VectorizedExpr) VectorizedSub(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "subtract", nil)
}

func (v *VectorizedExpr) VectorizedMul(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "multiply", nil)
}

func (v *VectorizedExpr) VectorizedDiv(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "divide", nil)
}

func (v *VectorizedExpr) VectorizedGt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "greater", nil)
}

func (v *VectorizedExpr) VectorizedLt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "less", nil)
}

func (v *VectorizedExpr) VectorizedEq(other Expr) Expr {
	return NewVectorizedExpr([]Expr{v, other}, "equal", nil)
}

func (v *VectorizedExpr) VectorizedSum() Expr {
	return NewVectorizedExpr([]Expr{v}, "sum", nil)
}

func (v *VectorizedExpr) VectorizedMean() Expr {
	return NewVectorizedExpr([]Expr{v}, "mean", nil)
}

func (v *VectorizedExpr) VectorizedMin() Expr {
	return NewVectorizedExpr([]Expr{v}, "min", nil)
}

func (v *VectorizedExpr) VectorizedMax() Expr {
	return NewVectorizedExpr([]Expr{v}, "max", nil)
}

func (v *VectorizedExpr) VectorizedStdDev() Expr {
	return NewVectorizedExpr([]Expr{v}, "stddev", nil)
}

func (v *VectorizedExpr) VectorizedVariance() Expr {
	return NewVectorizedExpr([]Expr{v}, "variance", nil)
}

// Vectorized operation methods for existing expression types
func (c *ColumnExpr) VectorizedAdd(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "add", nil)
}

func (c *ColumnExpr) VectorizedSub(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "subtract", nil)
}

func (c *ColumnExpr) VectorizedMul(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "multiply", nil)
}

func (c *ColumnExpr) VectorizedDiv(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "divide", nil)
}

func (c *ColumnExpr) VectorizedGt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "greater", nil)
}

func (c *ColumnExpr) VectorizedLt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "less", nil)
}

func (c *ColumnExpr) VectorizedEq(other Expr) Expr {
	return NewVectorizedExpr([]Expr{c, other}, "equal", nil)
}

// Vectorized aggregation methods - using correct Arrow compute function names
func (c *ColumnExpr) VectorizedSum() Expr {
	return NewVectorizedExpr([]Expr{c}, "sum", nil)
}

func (c *ColumnExpr) VectorizedMean() Expr {
	return NewVectorizedExpr([]Expr{c}, "mean", nil)
}

func (c *ColumnExpr) VectorizedMin() Expr {
	return NewVectorizedExpr([]Expr{c}, "min_max", nil) // Arrow uses min_max function
}

func (c *ColumnExpr) VectorizedMax() Expr {
	return NewVectorizedExpr([]Expr{c}, "min_max", nil) // Arrow uses min_max function
}

func (c *ColumnExpr) VectorizedStdDev() Expr {
	return NewVectorizedExpr([]Expr{c}, "stddev", nil)
}

func (c *ColumnExpr) VectorizedVariance() Expr {
	return NewVectorizedExpr([]Expr{c}, "variance", nil)
}

// Add vectorized methods to other expression types as well
func (l *LiteralExpr) VectorizedAdd(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "add", nil)
}

func (l *LiteralExpr) VectorizedSub(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "subtract", nil)
}

func (l *LiteralExpr) VectorizedMul(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "multiply", nil)
}

func (l *LiteralExpr) VectorizedDiv(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "divide", nil)
}

func (l *LiteralExpr) VectorizedGt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "greater", nil)
}

func (l *LiteralExpr) VectorizedLt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "less", nil)
}

func (l *LiteralExpr) VectorizedEq(other Expr) Expr {
	return NewVectorizedExpr([]Expr{l, other}, "equal", nil)
}

func (l *LiteralExpr) VectorizedSum() Expr {
	return NewVectorizedExpr([]Expr{l}, "sum", nil)
}

func (l *LiteralExpr) VectorizedMean() Expr {
	return NewVectorizedExpr([]Expr{l}, "mean", nil)
}

func (l *LiteralExpr) VectorizedMin() Expr {
	return NewVectorizedExpr([]Expr{l}, "min", nil)
}

func (l *LiteralExpr) VectorizedMax() Expr {
	return NewVectorizedExpr([]Expr{l}, "max", nil)
}

func (l *LiteralExpr) VectorizedStdDev() Expr {
	return NewVectorizedExpr([]Expr{l}, "stddev", nil)
}

func (l *LiteralExpr) VectorizedVariance() Expr {
	return NewVectorizedExpr([]Expr{l}, "variance", nil)
}

// Add vectorized methods to UnaryExpr and BinaryExpr as well
func (u *UnaryExpr) VectorizedAdd(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "add", nil)
}

func (u *UnaryExpr) VectorizedSub(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "subtract", nil)
}

func (u *UnaryExpr) VectorizedMul(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "multiply", nil)
}

func (u *UnaryExpr) VectorizedDiv(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "divide", nil)
}

func (u *UnaryExpr) VectorizedGt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "greater", nil)
}

func (u *UnaryExpr) VectorizedLt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "less", nil)
}

func (u *UnaryExpr) VectorizedEq(other Expr) Expr {
	return NewVectorizedExpr([]Expr{u, other}, "equal", nil)
}

func (u *UnaryExpr) VectorizedSum() Expr {
	return NewVectorizedExpr([]Expr{u}, "sum", nil)
}

func (u *UnaryExpr) VectorizedMean() Expr {
	return NewVectorizedExpr([]Expr{u}, "mean", nil)
}

func (u *UnaryExpr) VectorizedMin() Expr {
	return NewVectorizedExpr([]Expr{u}, "min", nil)
}

func (u *UnaryExpr) VectorizedMax() Expr {
	return NewVectorizedExpr([]Expr{u}, "max", nil)
}

func (u *UnaryExpr) VectorizedStdDev() Expr {
	return NewVectorizedExpr([]Expr{u}, "stddev", nil)
}

func (u *UnaryExpr) VectorizedVariance() Expr {
	return NewVectorizedExpr([]Expr{u}, "variance", nil)
}

func (b *BinaryExpr) VectorizedAdd(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "add", nil)
}

func (b *BinaryExpr) VectorizedSub(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "subtract", nil)
}

func (b *BinaryExpr) VectorizedMul(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "multiply", nil)
}

func (b *BinaryExpr) VectorizedDiv(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "divide", nil)
}

func (b *BinaryExpr) VectorizedGt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "greater", nil)
}

func (b *BinaryExpr) VectorizedLt(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "less", nil)
}

func (b *BinaryExpr) VectorizedEq(other Expr) Expr {
	return NewVectorizedExpr([]Expr{b, other}, "equal", nil)
}

func (b *BinaryExpr) VectorizedSum() Expr {
	return NewVectorizedExpr([]Expr{b}, "sum", nil)
}

func (b *BinaryExpr) VectorizedMean() Expr {
	return NewVectorizedExpr([]Expr{b}, "mean", nil)
}

func (b *BinaryExpr) VectorizedMin() Expr {
	return NewVectorizedExpr([]Expr{b}, "min", nil)
}

func (b *BinaryExpr) VectorizedMax() Expr {
	return NewVectorizedExpr([]Expr{b}, "max", nil)
}

func (b *BinaryExpr) VectorizedStdDev() Expr {
	return NewVectorizedExpr([]Expr{b}, "stddev", nil)
}

func (b *BinaryExpr) VectorizedVariance() Expr {
	return NewVectorizedExpr([]Expr{b}, "variance", nil)
}
