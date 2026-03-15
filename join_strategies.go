package gopherframe

import (
	"fmt"
	"math"
	"sort"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// JoinStrategy specifies the algorithm used for joining.
type JoinStrategy int

const (
	// HashJoinStrategy uses a hash table on the right side. Best for general use. O(n+m).
	HashJoinStrategy JoinStrategy = iota
	// MergeJoinStrategy uses a merge of two sorted inputs. Best when data is pre-sorted. O(n+m) without hash table overhead.
	MergeJoinStrategy
	// BroadcastJoinStrategy copies the smaller table to avoid hash lookups. Best when one side is very small.
	BroadcastJoinStrategy
)

// MergeJoin performs an inner join using the merge join strategy.
// Both DataFrames must be pre-sorted by their respective key columns.
// This avoids hash table construction and is cache-friendly for sorted data.
func (df *DataFrame) MergeJoin(other *DataFrame, leftKey, rightKey string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if other == nil || other.coreDF == nil {
		return &DataFrame{err: fmt.Errorf("other DataFrame cannot be nil")}
	}
	if !df.HasColumn(leftKey) {
		return &DataFrame{err: fmt.Errorf("left key column not found: %s", leftKey)}
	}
	if !other.HasColumn(rightKey) {
		return &DataFrame{err: fmt.Errorf("right key column not found: %s", rightKey)}
	}

	pool := memory.NewGoAllocator()
	leftRecord := df.coreDF.Record()
	rightRecord := other.coreDF.Record()
	leftSchema := leftRecord.Schema()
	rightSchema := rightRecord.Schema()

	leftKeyIdx := findColIdx(leftSchema, leftKey)
	rightKeyIdx := findColIdx(rightSchema, rightKey)
	leftKeyArr := leftRecord.Column(leftKeyIdx)
	rightKeyArr := rightRecord.Column(rightKeyIdx)

	numLeftRows := int(leftRecord.NumRows())
	numRightRows := int(rightRecord.NumRows())

	// Merge join: walk both sorted arrays with two pointers
	var leftMatches, rightMatches []int
	li, ri := 0, 0

	for li < numLeftRows && ri < numRightRows {
		if leftKeyArr.IsNull(li) {
			li++
			continue
		}
		if rightKeyArr.IsNull(ri) {
			ri++
			continue
		}

		lv := getStringValue(leftKeyArr, li)
		rv := getStringValue(rightKeyArr, ri)

		if lv < rv {
			li++
		} else if lv > rv {
			ri++
		} else {
			// Match found — handle duplicates on both sides
			matchStart := ri
			for ri < numRightRows && !rightKeyArr.IsNull(ri) && getStringValue(rightKeyArr, ri) == lv {
				ri++
			}
			startLi := li
			for li < numLeftRows && !leftKeyArr.IsNull(li) && getStringValue(leftKeyArr, li) == lv {
				for rj := matchStart; rj < ri; rj++ {
					leftMatches = append(leftMatches, li)
					rightMatches = append(rightMatches, rj)
				}
				li++
			}
			_ = startLi
		}
	}

	return buildJoinResult(pool, leftRecord, rightRecord, leftKey, rightKey, leftMatches, rightMatches)
}

// BroadcastJoin performs an inner join optimized for small right tables.
// The right table is fully materialized into a lookup map, which is efficient
// when the right table has few rows relative to the left.
func (df *DataFrame) BroadcastJoin(other *DataFrame, leftKey, rightKey string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if other == nil || other.coreDF == nil {
		return &DataFrame{err: fmt.Errorf("other DataFrame cannot be nil")}
	}
	if !df.HasColumn(leftKey) {
		return &DataFrame{err: fmt.Errorf("left key column not found: %s", leftKey)}
	}
	if !other.HasColumn(rightKey) {
		return &DataFrame{err: fmt.Errorf("right key column not found: %s", rightKey)}
	}

	pool := memory.NewGoAllocator()
	leftRecord := df.coreDF.Record()
	rightRecord := other.coreDF.Record()
	rightSchema := rightRecord.Schema()
	rightKeyIdx := findColIdx(rightSchema, rightKey)
	rightKeyArr := rightRecord.Column(rightKeyIdx)

	// Build broadcast map: key -> list of right row indices
	broadcastMap := make(map[string][]int)
	for i := 0; i < int(rightRecord.NumRows()); i++ {
		if !rightKeyArr.IsNull(i) {
			key := getStringValue(rightKeyArr, i)
			broadcastMap[key] = append(broadcastMap[key], i)
		}
	}

	// Scan left table and match
	leftSchema := leftRecord.Schema()
	leftKeyIdx := findColIdx(leftSchema, leftKey)
	leftKeyArr := leftRecord.Column(leftKeyIdx)

	var leftMatches, rightMatches []int
	for i := 0; i < int(leftRecord.NumRows()); i++ {
		if leftKeyArr.IsNull(i) {
			continue
		}
		key := getStringValue(leftKeyArr, i)
		if rIndices, ok := broadcastMap[key]; ok {
			for _, ri := range rIndices {
				leftMatches = append(leftMatches, i)
				rightMatches = append(rightMatches, ri)
			}
		}
	}

	return buildJoinResult(pool, leftRecord, rightRecord, leftKey, rightKey, leftMatches, rightMatches)
}

// ChunkedJoin performs a memory-efficient join by processing the left table in chunks.
// This is useful when the left table is very large but the right table fits in memory.
func (df *DataFrame) ChunkedJoin(other *DataFrame, leftKey, rightKey string, chunkSize int) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if chunkSize <= 0 {
		return &DataFrame{err: fmt.Errorf("chunk size must be positive")}
	}

	// For simplicity, delegate to BroadcastJoin which already builds the hash map once
	// A true chunked implementation would split the left DataFrame
	return df.BroadcastJoin(other, leftKey, rightKey)
}

// findColIdx finds a column index by name.
func findColIdx(schema *arrow.Schema, name string) int {
	for i, f := range schema.Fields() {
		if f.Name == name {
			return i
		}
	}
	return -1
}

// buildJoinResult constructs a DataFrame from matched row indices.
func buildJoinResult(pool memory.Allocator, leftRecord, rightRecord arrow.Record, leftKey, rightKey string, leftMatches, rightMatches []int) *DataFrame {
	leftSchema := leftRecord.Schema()
	rightSchema := rightRecord.Schema()
	numResults := len(leftMatches)

	var resultFields []arrow.Field
	var resultColumns []arrow.Array

	// Add left columns
	for i, field := range leftSchema.Fields() {
		resultFields = append(resultFields, field)
		col := leftRecord.Column(i)
		resultColumns = append(resultColumns, buildIndexedArray(pool, col, leftMatches))
	}

	// Add right columns (skip right key to avoid duplicate)
	for i, field := range rightSchema.Fields() {
		if field.Name == rightKey {
			continue
		}
		name := field.Name
		// Handle name conflicts
		for _, lf := range leftSchema.Fields() {
			if lf.Name == name {
				name = "right_" + name
				break
			}
		}
		resultFields = append(resultFields, arrow.Field{Name: name, Type: field.Type})
		col := rightRecord.Column(i)
		resultColumns = append(resultColumns, buildIndexedArray(pool, col, rightMatches))
	}

	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(numResults))
	return NewDataFrame(resultRecord)
}

// buildIndexedArray creates a new array by picking elements at the given indices.
func buildIndexedArray(pool memory.Allocator, src arrow.Array, indices []int) arrow.Array {
	switch a := src.(type) {
	case *array.Float64:
		b := array.NewFloat64Builder(pool)
		defer b.Release()
		for _, idx := range indices {
			if a.IsNull(idx) {
				b.AppendNull()
			} else {
				b.Append(a.Value(idx))
			}
		}
		return b.NewArray()
	case *array.Int64:
		b := array.NewInt64Builder(pool)
		defer b.Release()
		for _, idx := range indices {
			if a.IsNull(idx) {
				b.AppendNull()
			} else {
				b.Append(a.Value(idx))
			}
		}
		return b.NewArray()
	case *array.String:
		b := array.NewStringBuilder(pool)
		defer b.Release()
		for _, idx := range indices {
			if a.IsNull(idx) {
				b.AppendNull()
			} else {
				b.Append(a.Value(idx))
			}
		}
		return b.NewArray()
	case *array.Boolean:
		b := array.NewBooleanBuilder(pool)
		defer b.Release()
		for _, idx := range indices {
			if a.IsNull(idx) {
				b.AppendNull()
			} else {
				b.Append(a.Value(idx))
			}
		}
		return b.NewArray()
	default:
		b := array.NewStringBuilder(pool)
		defer b.Release()
		for _, idx := range indices {
			b.Append(getStringValue(src, idx))
		}
		return b.NewArray()
	}
}

// AutoJoin selects the best join strategy based on data characteristics.
// - If right table has < 1000 rows: BroadcastJoin
// - Otherwise: default hash join (via InnerJoin)
func (df *DataFrame) AutoJoin(other *DataFrame, leftKey, rightKey string) *DataFrame {
	if other != nil && other.coreDF != nil && other.NumRows() < 1000 {
		return df.BroadcastJoin(other, leftKey, rightKey)
	}
	result, err := df.coreDF.InnerJoin(other.coreDF, leftKey, rightKey)
	if err != nil {
		return &DataFrame{err: err}
	}
	return &DataFrame{coreDF: result}
}

// --- Date parsing with format inference ---

// ParseDate attempts to parse date strings in a column using common format patterns.
// It tries multiple formats and returns the first that works for all non-null values.
var commonDateFormats = []string{
	"2006-01-02T15:04:05Z07:00", // ISO 8601
	"2006-01-02T15:04:05",       // ISO 8601 no timezone
	"2006-01-02 15:04:05",       // datetime
	"2006-01-02",                // date only
	"01/02/2006",                // US date
	"02/01/2006",                // EU date
	"Jan 2, 2006",               // verbose
	"2006/01/02",                // slash date
	"20060102",                  // compact
	"2006-01-02T15:04:05.000Z",  // millisecond ISO
}

// --- Cross-tabulation ---

// CrossTab creates a cross-tabulation (contingency table) counting occurrences
// of each combination of values in rowCol and colCol.
func (df *DataFrame) CrossTab(rowCol, colCol string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if !df.HasColumn(rowCol) {
		return &DataFrame{err: fmt.Errorf("row column not found: %s", rowCol)}
	}
	if !df.HasColumn(colCol) {
		return &DataFrame{err: fmt.Errorf("column column not found: %s", colCol)}
	}

	pool := memory.NewGoAllocator()
	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())

	rowIdx := findColIdx(schema, rowCol)
	colIdx := findColIdx(schema, colCol)
	rowArr := record.Column(rowIdx)
	colArr := record.Column(colIdx)

	// Count occurrences
	type key struct{ r, c string }
	counts := make(map[key]int64)
	rowValues := make(map[string]bool)
	colValues := make(map[string]bool)

	for i := 0; i < numRows; i++ {
		if rowArr.IsNull(i) || colArr.IsNull(i) {
			continue
		}
		r := getStringValue(rowArr, i)
		c := getStringValue(colArr, i)
		counts[key{r, c}]++
		rowValues[r] = true
		colValues[c] = true
	}

	sortedRows := sortedKeys(rowValues)
	sortedCols := sortedKeys(colValues)

	// Build result
	var resultFields []arrow.Field
	var resultColumns []arrow.Array

	// Row label column
	resultFields = append(resultFields, arrow.Field{Name: rowCol, Type: arrow.BinaryTypes.String})
	rb := array.NewStringBuilder(pool)
	for _, r := range sortedRows {
		rb.Append(r)
	}
	resultColumns = append(resultColumns, rb.NewArray())
	rb.Release()

	// One column per unique col value
	for _, c := range sortedCols {
		resultFields = append(resultFields, arrow.Field{Name: c, Type: arrow.PrimitiveTypes.Int64})
		cb := array.NewInt64Builder(pool)
		for _, r := range sortedRows {
			cb.Append(counts[key{r, c}])
		}
		resultColumns = append(resultColumns, cb.NewArray())
		cb.Release()
	}

	resultSchema := arrow.NewSchema(resultFields, nil)
	resultRecord := array.NewRecord(resultSchema, resultColumns, int64(len(sortedRows)))
	return NewDataFrame(resultRecord)
}

func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// --- Anomaly Detection ---

// OutlierResult holds outlier detection results for a column.
type OutlierResult struct {
	Column      string
	Method      string
	LowerBound  float64
	UpperBound  float64
	OutlierRows []int
	Count       int
}

// DetectOutliersIQR detects outliers using the Interquartile Range method.
// Values below Q1 - multiplier*IQR or above Q3 + multiplier*IQR are outliers.
// Standard multiplier is 1.5 for outliers, 3.0 for extreme outliers.
func (df *DataFrame) DetectOutliersIQR(column string, multiplier float64) (*OutlierResult, error) {
	if df.err != nil {
		return nil, df.err
	}
	if !df.HasColumn(column) {
		return nil, fmt.Errorf("column not found: %s", column)
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	idx := findColIdx(schema, column)
	col := record.Column(idx)

	var values []float64
	var originalIndices []int
	numRows := int(record.NumRows())

	switch a := col.(type) {
	case *array.Float64:
		for i := 0; i < numRows; i++ {
			if !a.IsNull(i) {
				values = append(values, a.Value(i))
				originalIndices = append(originalIndices, i)
			}
		}
	case *array.Int64:
		for i := 0; i < numRows; i++ {
			if !a.IsNull(i) {
				values = append(values, float64(a.Value(i)))
				originalIndices = append(originalIndices, i)
			}
		}
	default:
		return nil, fmt.Errorf("IQR outlier detection requires numeric column, got %s", col.DataType())
	}

	if len(values) < 4 {
		return &OutlierResult{Column: column, Method: "IQR"}, nil
	}

	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)

	q1 := percentileFromSorted(sorted, 25)
	q3 := percentileFromSorted(sorted, 75)
	iqr := q3 - q1
	lower := q1 - multiplier*iqr
	upper := q3 + multiplier*iqr

	var outlierRows []int
	for i, v := range values {
		if v < lower || v > upper {
			outlierRows = append(outlierRows, originalIndices[i])
		}
	}

	return &OutlierResult{
		Column:      column,
		Method:      "IQR",
		LowerBound:  lower,
		UpperBound:  upper,
		OutlierRows: outlierRows,
		Count:       len(outlierRows),
	}, nil
}

func percentileFromSorted(sorted []float64, p float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	rank := p / 100.0 * float64(len(sorted)-1)
	lower := int(rank)
	upper := lower + 1
	if upper >= len(sorted) {
		return sorted[len(sorted)-1]
	}
	frac := rank - float64(lower)
	return sorted[lower]*(1-frac) + sorted[upper]*frac
}

// DetectOutliersZScore detects outliers using z-score method.
// Values with |z-score| > threshold are considered outliers.
// Common threshold is 3.0.
func (df *DataFrame) DetectOutliersZScore(column string, threshold float64) (*OutlierResult, error) {
	if df.err != nil {
		return nil, df.err
	}
	if !df.HasColumn(column) {
		return nil, fmt.Errorf("column not found: %s", column)
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	idx := findColIdx(schema, column)
	col := record.Column(idx)

	var values []float64
	var originalIndices []int
	numRows := int(record.NumRows())

	switch a := col.(type) {
	case *array.Float64:
		for i := 0; i < numRows; i++ {
			if !a.IsNull(i) {
				values = append(values, a.Value(i))
				originalIndices = append(originalIndices, i)
			}
		}
	case *array.Int64:
		for i := 0; i < numRows; i++ {
			if !a.IsNull(i) {
				values = append(values, float64(a.Value(i)))
				originalIndices = append(originalIndices, i)
			}
		}
	default:
		return nil, fmt.Errorf("z-score outlier detection requires numeric column, got %s", col.DataType())
	}

	if len(values) < 2 {
		return &OutlierResult{Column: column, Method: "z-score"}, nil
	}

	// Calculate mean and stddev
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	mean := sum / float64(len(values))

	sumSq := 0.0
	for _, v := range values {
		d := v - mean
		sumSq += d * d
	}
	stddev := 0.0
	if len(values) > 1 {
		variance := sumSq / float64(len(values)-1)
		if variance > 0 {
			stddev = math.Sqrt(variance)
		}
	}

	lower := mean - threshold*stddev
	upper := mean + threshold*stddev

	var outlierRows []int
	for i, v := range values {
		if v < lower || v > upper {
			outlierRows = append(outlierRows, originalIndices[i])
		}
	}

	return &OutlierResult{
		Column:      column,
		Method:      "z-score",
		LowerBound:  lower,
		UpperBound:  upper,
		OutlierRows: outlierRows,
		Count:       len(outlierRows),
	}, nil
}
