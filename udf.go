package gopherframe

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
)

// ScalarUDFFunc operates on a single row's values and returns a single value.
type ScalarUDFFunc func(row map[string]interface{}) (interface{}, error)

// VectorUDFFunc operates on entire Arrow arrays for high performance.
type VectorUDFFunc func(columns map[string]arrow.Array) (arrow.Array, error)

// ScalarUDF creates an expression that applies a scalar function row-by-row.
func ScalarUDF(inputCols []string, outputType arrow.DataType, fn ScalarUDFFunc) expr.Expr {
	return &scalarUDFExpr{
		inputCols:  inputCols,
		outputType: outputType,
		fn:         fn,
	}
}

// VectorUDF creates an expression that applies a vectorized function on entire columns.
func VectorUDF(inputCols []string, outputType arrow.DataType, fn VectorUDFFunc) expr.Expr {
	return &vectorUDFExpr{
		inputCols:  inputCols,
		outputType: outputType,
		fn:         fn,
	}
}

// scalarUDFExpr implements expr.Expr for scalar UDFs.
type scalarUDFExpr struct {
	inputCols  []string
	outputType arrow.DataType
	fn         ScalarUDFFunc
}

func (s *scalarUDFExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	pool := memory.NewGoAllocator()
	numRows := int(df.NumRows())

	// Resolve input column indices
	colIndices := make(map[string]int)
	schema := df.Schema()
	for _, col := range s.inputCols {
		idx := -1
		for i, f := range schema.Fields() {
			if f.Name == col {
				idx = i
				break
			}
		}
		if idx < 0 {
			return nil, fmt.Errorf("UDF input column not found: %s", col)
		}
		colIndices[col] = idx
	}

	record := df.Record()

	// Build result
	switch s.outputType.ID() {
	case arrow.FLOAT64:
		builder := array.NewFloat64Builder(pool)
		defer builder.Release()
		for i := 0; i < numRows; i++ {
			row := s.extractRow(record, colIndices, i)
			val, err := s.fn(row)
			if err != nil {
				return nil, fmt.Errorf("UDF error at row %d: %w", i, err)
			}
			if val == nil {
				builder.AppendNull()
			} else {
				switch v := val.(type) {
				case float64:
					builder.Append(v)
				case int:
					builder.Append(float64(v))
				case int64:
					builder.Append(float64(v))
				default:
					return nil, fmt.Errorf("UDF returned %T, expected float64", val)
				}
			}
		}
		return builder.NewArray(), nil

	case arrow.INT64:
		builder := array.NewInt64Builder(pool)
		defer builder.Release()
		for i := 0; i < numRows; i++ {
			row := s.extractRow(record, colIndices, i)
			val, err := s.fn(row)
			if err != nil {
				return nil, fmt.Errorf("UDF error at row %d: %w", i, err)
			}
			if val == nil {
				builder.AppendNull()
			} else {
				switch v := val.(type) {
				case int64:
					builder.Append(v)
				case int:
					builder.Append(int64(v))
				case float64:
					builder.Append(int64(v))
				default:
					return nil, fmt.Errorf("UDF returned %T, expected int64", val)
				}
			}
		}
		return builder.NewArray(), nil

	case arrow.STRING:
		builder := array.NewStringBuilder(pool)
		defer builder.Release()
		for i := 0; i < numRows; i++ {
			row := s.extractRow(record, colIndices, i)
			val, err := s.fn(row)
			if err != nil {
				return nil, fmt.Errorf("UDF error at row %d: %w", i, err)
			}
			if val == nil {
				builder.AppendNull()
			} else {
				builder.Append(fmt.Sprintf("%v", val))
			}
		}
		return builder.NewArray(), nil

	case arrow.BOOL:
		builder := array.NewBooleanBuilder(pool)
		defer builder.Release()
		for i := 0; i < numRows; i++ {
			row := s.extractRow(record, colIndices, i)
			val, err := s.fn(row)
			if err != nil {
				return nil, fmt.Errorf("UDF error at row %d: %w", i, err)
			}
			if val == nil {
				builder.AppendNull()
			} else {
				b, ok := val.(bool)
				if !ok {
					return nil, fmt.Errorf("UDF returned %T, expected bool", val)
				}
				builder.Append(b)
			}
		}
		return builder.NewArray(), nil

	default:
		return nil, fmt.Errorf("unsupported UDF output type: %s", s.outputType)
	}
}

func (s *scalarUDFExpr) extractRow(record arrow.Record, colIndices map[string]int, rowIdx int) map[string]interface{} {
	row := make(map[string]interface{}, len(colIndices))
	for col, idx := range colIndices {
		arr := record.Column(idx)
		if arr.IsNull(rowIdx) {
			row[col] = nil
			continue
		}
		switch a := arr.(type) {
		case *array.Float64:
			row[col] = a.Value(rowIdx)
		case *array.Int64:
			row[col] = a.Value(rowIdx)
		case *array.String:
			row[col] = a.Value(rowIdx)
		case *array.Boolean:
			row[col] = a.Value(rowIdx)
		default:
			row[col] = fmt.Sprintf("%v", a)
		}
	}
	return row
}

func (s *scalarUDFExpr) Name() string   { return "scalar_udf" }
func (s *scalarUDFExpr) String() string { return "ScalarUDF(...)" }
func (s *scalarUDFExpr) Add(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "add")
}
func (s *scalarUDFExpr) Sub(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "sub")
}
func (s *scalarUDFExpr) Mul(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "mul")
}
func (s *scalarUDFExpr) Div(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "div")
}
func (s *scalarUDFExpr) Gt(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "gt")
}
func (s *scalarUDFExpr) Lt(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "lt")
}
func (s *scalarUDFExpr) Eq(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, other, "eq")
}
func (s *scalarUDFExpr) Contains(sub expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, sub, "contains")
}
func (s *scalarUDFExpr) StartsWith(prefix expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, prefix, "starts_with")
}
func (s *scalarUDFExpr) EndsWith(suffix expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, suffix, "ends_with")
}
func (s *scalarUDFExpr) Upper() expr.Expr     { return expr.NewUnaryExpr(s, "upper") }
func (s *scalarUDFExpr) Lower() expr.Expr     { return expr.NewUnaryExpr(s, "lower") }
func (s *scalarUDFExpr) Trim() expr.Expr      { return expr.NewUnaryExpr(s, "trim") }
func (s *scalarUDFExpr) TrimLeft() expr.Expr  { return expr.NewUnaryExpr(s, "trim_left") }
func (s *scalarUDFExpr) TrimRight() expr.Expr { return expr.NewUnaryExpr(s, "trim_right") }
func (s *scalarUDFExpr) Length() expr.Expr    { return expr.NewUnaryExpr(s, "length") }
func (s *scalarUDFExpr) Match(pattern expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, pattern, "match")
}
func (s *scalarUDFExpr) Replace(old, new expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(s, old, new, "replace")
}
func (s *scalarUDFExpr) PadLeft(length, pad expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(s, length, pad, "pad_left")
}
func (s *scalarUDFExpr) PadRight(length, pad expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(s, length, pad, "pad_right")
}
func (s *scalarUDFExpr) SplitPart(separator, index expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(s, separator, index, "split_part")
}
func (s *scalarUDFExpr) Year() expr.Expr            { return expr.NewUnaryExpr(s, "year") }
func (s *scalarUDFExpr) Month() expr.Expr           { return expr.NewUnaryExpr(s, "month") }
func (s *scalarUDFExpr) Day() expr.Expr             { return expr.NewUnaryExpr(s, "day") }
func (s *scalarUDFExpr) Hour() expr.Expr            { return expr.NewUnaryExpr(s, "hour") }
func (s *scalarUDFExpr) Minute() expr.Expr          { return expr.NewUnaryExpr(s, "minute") }
func (s *scalarUDFExpr) Second() expr.Expr          { return expr.NewUnaryExpr(s, "second") }
func (s *scalarUDFExpr) TruncateToYear() expr.Expr  { return expr.NewUnaryExpr(s, "truncate_year") }
func (s *scalarUDFExpr) TruncateToMonth() expr.Expr { return expr.NewUnaryExpr(s, "truncate_month") }
func (s *scalarUDFExpr) TruncateToDay() expr.Expr   { return expr.NewUnaryExpr(s, "truncate_day") }
func (s *scalarUDFExpr) TruncateToHour() expr.Expr  { return expr.NewUnaryExpr(s, "truncate_hour") }
func (s *scalarUDFExpr) AddDays(d expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, d, "add_days")
}
func (s *scalarUDFExpr) AddHours(h expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, h, "add_hours")
}
func (s *scalarUDFExpr) AddMinutes(m expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, m, "add_minutes")
}
func (s *scalarUDFExpr) AddSeconds(sec expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(s, sec, "add_seconds")
}

// vectorUDFExpr implements expr.Expr for vectorized UDFs.
type vectorUDFExpr struct {
	inputCols  []string
	outputType arrow.DataType
	fn         VectorUDFFunc
}

func (v *vectorUDFExpr) Evaluate(df *core.DataFrame) (arrow.Array, error) {
	schema := df.Schema()
	record := df.Record()

	columns := make(map[string]arrow.Array, len(v.inputCols))
	for _, col := range v.inputCols {
		idx := -1
		for i, f := range schema.Fields() {
			if f.Name == col {
				idx = i
				break
			}
		}
		if idx < 0 {
			return nil, fmt.Errorf("VectorUDF input column not found: %s", col)
		}
		columns[col] = record.Column(idx)
	}

	return v.fn(columns)
}

func (v *vectorUDFExpr) Name() string   { return "vector_udf" }
func (v *vectorUDFExpr) String() string { return "VectorUDF(...)" }
func (v *vectorUDFExpr) Add(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "add")
}
func (v *vectorUDFExpr) Sub(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "sub")
}
func (v *vectorUDFExpr) Mul(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "mul")
}
func (v *vectorUDFExpr) Div(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "div")
}
func (v *vectorUDFExpr) Gt(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "gt")
}
func (v *vectorUDFExpr) Lt(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "lt")
}
func (v *vectorUDFExpr) Eq(other expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, other, "eq")
}
func (v *vectorUDFExpr) Contains(sub expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, sub, "contains")
}
func (v *vectorUDFExpr) StartsWith(prefix expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, prefix, "starts_with")
}
func (v *vectorUDFExpr) EndsWith(suffix expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, suffix, "ends_with")
}
func (v *vectorUDFExpr) Upper() expr.Expr     { return expr.NewUnaryExpr(v, "upper") }
func (v *vectorUDFExpr) Lower() expr.Expr     { return expr.NewUnaryExpr(v, "lower") }
func (v *vectorUDFExpr) Trim() expr.Expr      { return expr.NewUnaryExpr(v, "trim") }
func (v *vectorUDFExpr) TrimLeft() expr.Expr  { return expr.NewUnaryExpr(v, "trim_left") }
func (v *vectorUDFExpr) TrimRight() expr.Expr { return expr.NewUnaryExpr(v, "trim_right") }
func (v *vectorUDFExpr) Length() expr.Expr    { return expr.NewUnaryExpr(v, "length") }
func (v *vectorUDFExpr) Match(pattern expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, pattern, "match")
}
func (v *vectorUDFExpr) Replace(old, new expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(v, old, new, "replace")
}
func (v *vectorUDFExpr) PadLeft(length, pad expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(v, length, pad, "pad_left")
}
func (v *vectorUDFExpr) PadRight(length, pad expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(v, length, pad, "pad_right")
}
func (v *vectorUDFExpr) SplitPart(separator, index expr.Expr) expr.Expr {
	return expr.NewTernaryExpr(v, separator, index, "split_part")
}
func (v *vectorUDFExpr) Year() expr.Expr            { return expr.NewUnaryExpr(v, "year") }
func (v *vectorUDFExpr) Month() expr.Expr           { return expr.NewUnaryExpr(v, "month") }
func (v *vectorUDFExpr) Day() expr.Expr             { return expr.NewUnaryExpr(v, "day") }
func (v *vectorUDFExpr) Hour() expr.Expr            { return expr.NewUnaryExpr(v, "hour") }
func (v *vectorUDFExpr) Minute() expr.Expr          { return expr.NewUnaryExpr(v, "minute") }
func (v *vectorUDFExpr) Second() expr.Expr          { return expr.NewUnaryExpr(v, "second") }
func (v *vectorUDFExpr) TruncateToYear() expr.Expr  { return expr.NewUnaryExpr(v, "truncate_year") }
func (v *vectorUDFExpr) TruncateToMonth() expr.Expr { return expr.NewUnaryExpr(v, "truncate_month") }
func (v *vectorUDFExpr) TruncateToDay() expr.Expr   { return expr.NewUnaryExpr(v, "truncate_day") }
func (v *vectorUDFExpr) TruncateToHour() expr.Expr  { return expr.NewUnaryExpr(v, "truncate_hour") }
func (v *vectorUDFExpr) AddDays(d expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, d, "add_days")
}
func (v *vectorUDFExpr) AddHours(h expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, h, "add_hours")
}
func (v *vectorUDFExpr) AddMinutes(m expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, m, "add_minutes")
}
func (v *vectorUDFExpr) AddSeconds(sec expr.Expr) expr.Expr {
	return expr.NewBinaryExpr(v, sec, "add_seconds")
}
