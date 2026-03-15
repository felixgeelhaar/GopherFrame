package gopherframe

import (
	"fmt"
	"math"
	"strings"

	"github.com/apache/arrow-go/v18/arrow/array"
)

// ColumnStats holds descriptive statistics for a single column.
type ColumnStats struct {
	Name       string
	Type       string
	Count      int64
	NullCount  int64
	NullPct    float64
	Unique     int64
	Min        interface{}
	Max        interface{}
	Mean       float64
	StdDev     float64
	HasNumeric bool
}

// Describe returns descriptive statistics for all columns in the DataFrame.
func (df *DataFrame) Describe() ([]ColumnStats, error) {
	if df.err != nil {
		return nil, df.err
	}
	if df.coreDF == nil {
		return nil, fmt.Errorf("cannot describe nil DataFrame")
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())
	var stats []ColumnStats

	for i, field := range schema.Fields() {
		col := record.Column(i)
		cs := ColumnStats{
			Name:  field.Name,
			Type:  field.Type.String(),
			Count: int64(numRows),
		}

		// Count nulls
		for j := 0; j < numRows; j++ {
			if col.IsNull(j) {
				cs.NullCount++
			}
		}
		if numRows > 0 {
			cs.NullPct = float64(cs.NullCount) / float64(numRows) * 100
		}

		// Count unique values
		unique := make(map[string]bool)
		for j := 0; j < numRows; j++ {
			if !col.IsNull(j) {
				unique[getStringValue(col, j)] = true
			}
		}
		cs.Unique = int64(len(unique))

		// Numeric stats
		switch a := col.(type) {
		case *array.Float64:
			cs.HasNumeric = true
			var sum, sumSq float64
			var min, max float64
			count := 0
			first := true
			for j := 0; j < numRows; j++ {
				if !a.IsNull(j) {
					v := a.Value(j)
					sum += v
					sumSq += v * v
					if first || v < min {
						min = v
					}
					if first || v > max {
						max = v
					}
					first = false
					count++
				}
			}
			if count > 0 {
				cs.Mean = sum / float64(count)
				cs.Min = min
				cs.Max = max
				if count > 1 {
					variance := (sumSq - sum*sum/float64(count)) / float64(count-1)
					if variance > 0 {
						cs.StdDev = math.Sqrt(variance)
					}
				}
			}
		case *array.Int64:
			cs.HasNumeric = true
			var sum float64
			var sumSq float64
			var min, max int64
			count := 0
			first := true
			for j := 0; j < numRows; j++ {
				if !a.IsNull(j) {
					v := a.Value(j)
					sum += float64(v)
					sumSq += float64(v) * float64(v)
					if first || v < min {
						min = v
					}
					if first || v > max {
						max = v
					}
					first = false
					count++
				}
			}
			if count > 0 {
				cs.Mean = sum / float64(count)
				cs.Min = min
				cs.Max = max
				if count > 1 {
					variance := (sumSq - sum*sum/float64(count)) / float64(count-1)
					if variance > 0 {
						cs.StdDev = math.Sqrt(variance)
					}
				}
			}
		case *array.String:
			if numRows > 0 {
				var min, max string
				first := true
				for j := 0; j < numRows; j++ {
					if !a.IsNull(j) {
						v := a.Value(j)
						if first || v < min {
							min = v
						}
						if first || v > max {
							max = v
						}
						first = false
					}
				}
				cs.Min = min
				cs.Max = max
			}
		}

		stats = append(stats, cs)
	}

	return stats, nil
}

// NullCount returns a map of column names to their null counts.
func (df *DataFrame) NullCount() map[string]int64 {
	if df.err != nil || df.coreDF == nil {
		return nil
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())
	result := make(map[string]int64)

	for i, field := range schema.Fields() {
		col := record.Column(i)
		count := int64(0)
		for j := 0; j < numRows; j++ {
			if col.IsNull(j) {
				count++
			}
		}
		result[field.Name] = count
	}

	return result
}

// IsComplete returns true if the DataFrame has no null values in any column.
func (df *DataFrame) IsComplete() bool {
	if df.err != nil || df.coreDF == nil {
		return false
	}

	for _, count := range df.NullCount() {
		if count > 0 {
			return false
		}
	}
	return true
}

// ValidationRule defines a rule for data validation.
type ValidationRule struct {
	Column string
	Rule   string // "not_null", "positive", "in_range", "regex", "unique"
	Params map[string]interface{}
}

// NotNull creates a validation rule that checks a column has no null values.
func NotNull(column string) ValidationRule {
	return ValidationRule{Column: column, Rule: "not_null"}
}

// Positive creates a validation rule that checks all numeric values are positive.
func Positive(column string) ValidationRule {
	return ValidationRule{Column: column, Rule: "positive"}
}

// InRange creates a validation rule that checks values are within [min, max].
func InRange(column string, min, max float64) ValidationRule {
	return ValidationRule{
		Column: column,
		Rule:   "in_range",
		Params: map[string]interface{}{"min": min, "max": max},
	}
}

// UniqueValues creates a validation rule that checks all values in a column are unique.
func UniqueValues(column string) ValidationRule {
	return ValidationRule{Column: column, Rule: "unique"}
}

// ValidationResult holds the result of a validation check.
type ValidationResult struct {
	Valid      bool
	Violations []string
}

// Validate checks the DataFrame against a set of validation rules.
func (df *DataFrame) Validate(rules ...ValidationRule) ValidationResult {
	result := ValidationResult{Valid: true}

	if df.err != nil {
		result.Valid = false
		result.Violations = append(result.Violations, fmt.Sprintf("DataFrame has error: %v", df.err))
		return result
	}
	if df.coreDF == nil {
		result.Valid = false
		result.Violations = append(result.Violations, "DataFrame is nil")
		return result
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())

	for _, rule := range rules {
		// Find column
		colIdx := -1
		for i, f := range schema.Fields() {
			if f.Name == rule.Column {
				colIdx = i
				break
			}
		}
		if colIdx < 0 {
			result.Valid = false
			result.Violations = append(result.Violations, fmt.Sprintf("column %q not found", rule.Column))
			continue
		}

		col := record.Column(colIdx)

		switch rule.Rule {
		case "not_null":
			for j := 0; j < numRows; j++ {
				if col.IsNull(j) {
					result.Valid = false
					result.Violations = append(result.Violations,
						fmt.Sprintf("column %q has null at row %d", rule.Column, j))
					break
				}
			}

		case "positive":
			switch a := col.(type) {
			case *array.Float64:
				for j := 0; j < numRows; j++ {
					if !a.IsNull(j) && a.Value(j) <= 0 {
						result.Valid = false
						result.Violations = append(result.Violations,
							fmt.Sprintf("column %q has non-positive value %g at row %d", rule.Column, a.Value(j), j))
						break
					}
				}
			case *array.Int64:
				for j := 0; j < numRows; j++ {
					if !a.IsNull(j) && a.Value(j) <= 0 {
						result.Valid = false
						result.Violations = append(result.Violations,
							fmt.Sprintf("column %q has non-positive value %d at row %d", rule.Column, a.Value(j), j))
						break
					}
				}
			default:
				result.Valid = false
				result.Violations = append(result.Violations,
					fmt.Sprintf("column %q is not numeric for 'positive' rule", rule.Column))
			}

		case "in_range":
			minVal := rule.Params["min"].(float64)
			maxVal := rule.Params["max"].(float64)
			switch a := col.(type) {
			case *array.Float64:
				for j := 0; j < numRows; j++ {
					if !a.IsNull(j) && (a.Value(j) < minVal || a.Value(j) > maxVal) {
						result.Valid = false
						result.Violations = append(result.Violations,
							fmt.Sprintf("column %q value %g at row %d out of range [%g, %g]",
								rule.Column, a.Value(j), j, minVal, maxVal))
						break
					}
				}
			case *array.Int64:
				for j := 0; j < numRows; j++ {
					v := float64(a.Value(j))
					if !a.IsNull(j) && (v < minVal || v > maxVal) {
						result.Valid = false
						result.Violations = append(result.Violations,
							fmt.Sprintf("column %q value %g at row %d out of range [%g, %g]",
								rule.Column, v, j, minVal, maxVal))
						break
					}
				}
			}

		case "unique":
			seen := make(map[string]bool)
			for j := 0; j < numRows; j++ {
				if !col.IsNull(j) {
					val := getStringValue(col, j)
					if seen[val] {
						result.Valid = false
						result.Violations = append(result.Violations,
							fmt.Sprintf("column %q has duplicate value %q at row %d", rule.Column, val, j))
						break
					}
					seen[val] = true
				}
			}
		}
	}

	return result
}

// DescribeString returns a formatted string representation of Describe() output.
func (df *DataFrame) DescribeString() string {
	stats, err := df.Describe()
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "%-20s %-10s %8s %8s %8s %8s\n", "Column", "Type", "Count", "Nulls", "Null%", "Unique")
	sb.WriteString(strings.Repeat("-", 78) + "\n")
	for _, s := range stats {
		fmt.Fprintf(&sb, "%-20s %-10s %8d %8d %7.1f%% %8d\n",
			s.Name, s.Type, s.Count, s.NullCount, s.NullPct, s.Unique)
	}
	return sb.String()
}
