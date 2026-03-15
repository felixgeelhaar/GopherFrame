package gopherframe

import (
	"fmt"
	"sort"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// Pivot transforms long-format data to wide-format.
// indexCols: columns to keep as row identifiers
// pivotCol: column whose unique values become new column names
// valueCol: column whose values fill the new columns
func (df *DataFrame) Pivot(indexCols []string, pivotCol, valueCol string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if df.coreDF == nil {
		return &DataFrame{err: fmt.Errorf("cannot pivot nil DataFrame")}
	}

	// Validate columns
	for _, col := range indexCols {
		if !df.HasColumn(col) {
			return &DataFrame{err: fmt.Errorf("index column not found: %s", col)}
		}
	}
	if !df.HasColumn(pivotCol) {
		return &DataFrame{err: fmt.Errorf("pivot column not found: %s", pivotCol)}
	}
	if !df.HasColumn(valueCol) {
		return &DataFrame{err: fmt.Errorf("value column not found: %s", valueCol)}
	}

	pool := memory.NewGoAllocator()
	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())

	// Find column indices
	colIdx := func(name string) int {
		for i, f := range schema.Fields() {
			if f.Name == name {
				return i
			}
		}
		return -1
	}

	pivotIdx := colIdx(pivotCol)
	valueIdx := colIdx(valueCol)
	var indexIndices []int
	for _, col := range indexCols {
		indexIndices = append(indexIndices, colIdx(col))
	}

	// Extract unique pivot values (sorted for deterministic output)
	pivotValues := make(map[string]bool)
	pivotArr := record.Column(pivotIdx)
	for i := 0; i < numRows; i++ {
		if !pivotArr.IsNull(i) {
			pivotValues[getStringValue(pivotArr, i)] = true
		}
	}
	var pivotKeys []string
	for k := range pivotValues {
		pivotKeys = append(pivotKeys, k)
	}
	sort.Strings(pivotKeys)

	// Build lookup: compositeIndexKey + pivotValue -> value
	type cellKey struct {
		indexKey string
		pivotVal string
	}
	lookup := make(map[cellKey]interface{})
	indexKeySet := make(map[string]bool)
	var indexKeyOrder []string

	for i := 0; i < numRows; i++ {
		// Build composite index key
		var parts []string
		for _, idx := range indexIndices {
			parts = append(parts, getStringValue(record.Column(idx), i))
		}
		ik := fmt.Sprintf("%v", parts)

		if !indexKeySet[ik] {
			indexKeySet[ik] = true
			indexKeyOrder = append(indexKeyOrder, ik)
		}

		pv := getStringValue(pivotArr, i)
		val := getTypedValue(record.Column(valueIdx), i)
		lookup[cellKey{indexKey: ik, pivotVal: pv}] = val
	}

	// Build index column -> first row index mapping
	indexFirstRow := make(map[string]int)
	for i := 0; i < numRows; i++ {
		var parts []string
		for _, idx := range indexIndices {
			parts = append(parts, getStringValue(record.Column(idx), i))
		}
		ik := fmt.Sprintf("%v", parts)
		if _, exists := indexFirstRow[ik]; !exists {
			indexFirstRow[ik] = i
		}
	}

	// Build result
	resultRows := len(indexKeyOrder)
	var resultFields []arrow.Field
	var resultColumns []arrow.Array

	// Add index columns
	for _, ic := range indexCols {
		idx := colIdx(ic)
		resultFields = append(resultFields, schema.Field(idx))
		srcArr := record.Column(idx)

		switch a := srcArr.(type) {
		case *array.String:
			b := array.NewStringBuilder(pool)
			for _, ik := range indexKeyOrder {
				ri := indexFirstRow[ik]
				if a.IsNull(ri) {
					b.AppendNull()
				} else {
					b.Append(a.Value(ri))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		case *array.Int64:
			b := array.NewInt64Builder(pool)
			for _, ik := range indexKeyOrder {
				ri := indexFirstRow[ik]
				if a.IsNull(ri) {
					b.AppendNull()
				} else {
					b.Append(a.Value(ri))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		default:
			b := array.NewStringBuilder(pool)
			for _, ik := range indexKeyOrder {
				ri := indexFirstRow[ik]
				b.Append(getStringValue(srcArr, ri))
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		}
	}

	// Add pivot value columns
	valueType := schema.Field(valueIdx).Type
	for _, pv := range pivotKeys {
		resultFields = append(resultFields, arrow.Field{Name: pv, Type: valueType})

		switch valueType.ID() {
		case arrow.FLOAT64:
			b := array.NewFloat64Builder(pool)
			for _, ik := range indexKeyOrder {
				val, ok := lookup[cellKey{indexKey: ik, pivotVal: pv}]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					b.Append(val.(float64))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		case arrow.INT64:
			b := array.NewInt64Builder(pool)
			for _, ik := range indexKeyOrder {
				val, ok := lookup[cellKey{indexKey: ik, pivotVal: pv}]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					b.Append(val.(int64))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		default:
			b := array.NewStringBuilder(pool)
			for _, ik := range indexKeyOrder {
				val, ok := lookup[cellKey{indexKey: ik, pivotVal: pv}]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					b.Append(fmt.Sprintf("%v", val))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		}
	}

	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(resultRows))
	return NewDataFrame(resultRecord)
}

// Unpivot (Melt) transforms wide-format data to long-format.
// idCols: columns to keep as row identifiers
// valueCols: columns to unpivot into rows
// variableName: name for the column containing original column names
// valueName: name for the column containing the values
func (df *DataFrame) Unpivot(idCols, valueCols []string, variableName, valueName string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if df.coreDF == nil {
		return &DataFrame{err: fmt.Errorf("cannot unpivot nil DataFrame")}
	}

	for _, col := range idCols {
		if !df.HasColumn(col) {
			return &DataFrame{err: fmt.Errorf("id column not found: %s", col)}
		}
	}
	for _, col := range valueCols {
		if !df.HasColumn(col) {
			return &DataFrame{err: fmt.Errorf("value column not found: %s", col)}
		}
	}
	if len(valueCols) == 0 {
		return &DataFrame{err: fmt.Errorf("no value columns specified for unpivot")}
	}

	pool := memory.NewGoAllocator()
	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())
	outputRows := numRows * len(valueCols)

	colIdx := func(name string) int {
		for i, f := range schema.Fields() {
			if f.Name == name {
				return i
			}
		}
		return -1
	}

	var resultFields []arrow.Field
	var resultColumns []arrow.Array

	// Build ID columns (each value repeated len(valueCols) times)
	for _, idCol := range idCols {
		idx := colIdx(idCol)
		field := schema.Field(idx)
		resultFields = append(resultFields, field)
		srcArr := record.Column(idx)

		switch a := srcArr.(type) {
		case *array.String:
			b := array.NewStringBuilder(pool)
			for i := 0; i < numRows; i++ {
				for range valueCols {
					if a.IsNull(i) {
						b.AppendNull()
					} else {
						b.Append(a.Value(i))
					}
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		case *array.Int64:
			b := array.NewInt64Builder(pool)
			for i := 0; i < numRows; i++ {
				for range valueCols {
					if a.IsNull(i) {
						b.AppendNull()
					} else {
						b.Append(a.Value(i))
					}
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		case *array.Float64:
			b := array.NewFloat64Builder(pool)
			for i := 0; i < numRows; i++ {
				for range valueCols {
					if a.IsNull(i) {
						b.AppendNull()
					} else {
						b.Append(a.Value(i))
					}
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		default:
			b := array.NewStringBuilder(pool)
			for i := 0; i < numRows; i++ {
				for range valueCols {
					b.Append(getStringValue(srcArr, i))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		}
	}

	// Build variable column (column names)
	resultFields = append(resultFields, arrow.Field{Name: variableName, Type: arrow.BinaryTypes.String})
	varBuilder := array.NewStringBuilder(pool)
	for i := 0; i < numRows; i++ {
		for _, vc := range valueCols {
			varBuilder.Append(vc)
		}
	}
	resultColumns = append(resultColumns, varBuilder.NewArray())
	varBuilder.Release()

	// Build value column - use float64 as common type
	resultFields = append(resultFields, arrow.Field{Name: valueName, Type: arrow.PrimitiveTypes.Float64})
	valBuilder := array.NewFloat64Builder(pool)
	for i := 0; i < numRows; i++ {
		for _, vc := range valueCols {
			idx := colIdx(vc)
			arr := record.Column(idx)
			if arr.IsNull(i) {
				valBuilder.AppendNull()
			} else {
				switch a := arr.(type) {
				case *array.Float64:
					valBuilder.Append(a.Value(i))
				case *array.Int64:
					valBuilder.Append(float64(a.Value(i)))
				default:
					valBuilder.AppendNull()
				}
			}
		}
	}
	resultColumns = append(resultColumns, valBuilder.NewArray())
	valBuilder.Release()

	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(outputRows))
	return NewDataFrame(resultRecord)
}

// getStringValue extracts a string representation of a value from an Arrow array.
func getStringValue(arr arrow.Array, i int) string {
	if arr.IsNull(i) {
		return ""
	}
	switch a := arr.(type) {
	case *array.String:
		return a.Value(i)
	case *array.Int64:
		return fmt.Sprintf("%d", a.Value(i))
	case *array.Float64:
		return fmt.Sprintf("%g", a.Value(i))
	case *array.Boolean:
		return fmt.Sprintf("%t", a.Value(i))
	default:
		return fmt.Sprintf("%v", a)
	}
}

// getTypedValue extracts a typed value from an Arrow array.
func getTypedValue(arr arrow.Array, i int) interface{} {
	if arr.IsNull(i) {
		return nil
	}
	switch a := arr.(type) {
	case *array.Float64:
		return a.Value(i)
	case *array.Int64:
		return a.Value(i)
	case *array.String:
		return a.Value(i)
	case *array.Boolean:
		return a.Value(i)
	default:
		return nil
	}
}
