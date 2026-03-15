package gopherframe

import (
	"fmt"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// ParseDateColumn parses a string column into a timestamp column using automatic format inference.
// It tries common date formats and uses the first that successfully parses all non-null values.
func (df *DataFrame) ParseDateColumn(column, newColumn string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if !df.HasColumn(column) {
		return &DataFrame{err: fmt.Errorf("column not found: %s", column)}
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	colIdx := findColIdx(schema, column)
	col := record.Column(colIdx)

	strArr, ok := col.(*array.String)
	if !ok {
		return &DataFrame{err: fmt.Errorf("ParseDateColumn requires string column, got %s", col.DataType())}
	}

	numRows := int(record.NumRows())

	// Infer format from first non-null value
	format := ""
	for i := 0; i < numRows; i++ {
		if !strArr.IsNull(i) {
			format = inferDateFormat(strArr.Value(i))
			break
		}
	}
	if format == "" {
		return &DataFrame{err: fmt.Errorf("could not infer date format from column %s", column)}
	}

	// Parse all values
	pool := memory.NewGoAllocator()
	tsType := &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < numRows; i++ {
		if strArr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		t, err := time.Parse(format, strArr.Value(i))
		if err != nil {
			// Try other formats as fallback
			parsed := false
			for _, f := range commonDateFormats {
				if t2, err2 := time.Parse(f, strArr.Value(i)); err2 == nil {
					builder.Append(arrow.Timestamp(t2.UnixMicro()))
					parsed = true
					break
				}
			}
			if !parsed {
				builder.AppendNull()
			}
		} else {
			builder.Append(arrow.Timestamp(t.UnixMicro()))
		}
	}

	tsArr := builder.NewArray()

	// Build new record with additional column
	var newFields []arrow.Field
	var newColumns []arrow.Array
	for i, f := range schema.Fields() {
		newFields = append(newFields, f)
		newColumns = append(newColumns, record.Column(i))
	}
	newFields = append(newFields, arrow.Field{Name: newColumn, Type: tsType})
	newColumns = append(newColumns, tsArr)

	newSchema := arrow.NewSchema(newFields, nil)
	newRecord := array.NewRecord(newSchema, newColumns, int64(numRows))
	return NewDataFrame(newRecord)
}

// inferDateFormat tries common formats and returns the first that parses successfully.
func inferDateFormat(s string) string {
	for _, f := range commonDateFormats {
		if _, err := time.Parse(f, s); err == nil {
			return f
		}
	}
	return ""
}

// ParseDateWithFormat parses a string column into a timestamp column using a specific format.
func (df *DataFrame) ParseDateWithFormat(column, newColumn, format string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}
	if !df.HasColumn(column) {
		return &DataFrame{err: fmt.Errorf("column not found: %s", column)}
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	colIdx := findColIdx(schema, column)
	col := record.Column(colIdx)

	strArr, ok := col.(*array.String)
	if !ok {
		return &DataFrame{err: fmt.Errorf("ParseDateWithFormat requires string column, got %s", col.DataType())}
	}

	numRows := int(record.NumRows())
	pool := memory.NewGoAllocator()
	tsType := &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}
	builder := array.NewTimestampBuilder(pool, tsType)
	defer builder.Release()

	for i := 0; i < numRows; i++ {
		if strArr.IsNull(i) {
			builder.AppendNull()
			continue
		}
		t, err := time.Parse(format, strArr.Value(i))
		if err != nil {
			builder.AppendNull()
		} else {
			builder.Append(arrow.Timestamp(t.UnixMicro()))
		}
	}

	tsArr := builder.NewArray()

	var newFields []arrow.Field
	var newColumns []arrow.Array
	for i, f := range schema.Fields() {
		newFields = append(newFields, f)
		newColumns = append(newColumns, record.Column(i))
	}
	newFields = append(newFields, arrow.Field{Name: newColumn, Type: tsType})
	newColumns = append(newColumns, tsArr)

	newSchema := arrow.NewSchema(newFields, nil)
	newRecord := array.NewRecord(newSchema, newColumns, int64(numRows))
	return NewDataFrame(newRecord)
}
