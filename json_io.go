package gopherframe

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// ReadJSON reads a JSON file containing an array of objects into a DataFrame.
func ReadJSON(filename string) (*DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	var records []map[string]interface{}
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return jsonRecordsToDataFrame(records)
}

// WriteJSON writes a DataFrame to a JSON file as an array of objects.
func WriteJSON(df *DataFrame, filename string) error {
	if err := validateFilePath(filename); err != nil {
		return err
	}
	if df.err != nil {
		return df.err
	}

	records := dataFrameToJSONRecords(df)

	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filename, data, 0600); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

// ReadNDJSON reads a newline-delimited JSON file into a DataFrame.
func ReadNDJSON(filename string) (*DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open NDJSON file: %w", err)
	}
	defer func() { _ = f.Close() }()

	var records []map[string]interface{}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}
		var obj map[string]interface{}
		if err := json.Unmarshal([]byte(line), &obj); err != nil {
			return nil, fmt.Errorf("failed to parse NDJSON at line %d: %w", lineNum, err)
		}
		records = append(records, obj)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading NDJSON file: %w", err)
	}

	return jsonRecordsToDataFrame(records)
}

// WriteNDJSON writes a DataFrame as newline-delimited JSON.
func WriteNDJSON(df *DataFrame, filename string) error {
	if err := validateFilePath(filename); err != nil {
		return err
	}
	if df.err != nil {
		return df.err
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create NDJSON file: %w", err)
	}
	defer func() { _ = f.Close() }()

	records := dataFrameToJSONRecords(df)
	writer := bufio.NewWriter(f)
	for _, rec := range records {
		data, err := json.Marshal(rec)
		if err != nil {
			return fmt.Errorf("failed to marshal NDJSON record: %w", err)
		}
		if _, err := writer.Write(data); err != nil {
			return err
		}
		if _, err := writer.WriteString("\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}

// jsonRecordsToDataFrame converts parsed JSON records to a DataFrame.
func jsonRecordsToDataFrame(records []map[string]interface{}) (*DataFrame, error) {
	if len(records) == 0 {
		// Empty DataFrame
		schema := arrow.NewSchema([]arrow.Field{}, nil)
		record := array.NewRecord(schema, nil, 0)
		return NewDataFrame(record), nil
	}

	// Scan all records for column names and infer types
	columnNames := make(map[string]bool)
	columnTypes := make(map[string]arrow.DataType)

	for _, rec := range records {
		for key, val := range rec {
			columnNames[key] = true
			if val == nil {
				continue
			}
			inferredType := inferArrowType(val)
			existing, ok := columnTypes[key]
			if !ok {
				columnTypes[key] = inferredType
			} else if existing.ID() != inferredType.ID() {
				// Mixed types default to string
				columnTypes[key] = arrow.BinaryTypes.String
			}
		}
	}

	// Sort column names for deterministic output
	var sortedCols []string
	for col := range columnNames {
		sortedCols = append(sortedCols, col)
	}
	sort.Strings(sortedCols)

	pool := memory.NewGoAllocator()
	numRows := len(records)

	var fields []arrow.Field
	var columns []arrow.Array

	for _, colName := range sortedCols {
		colType, ok := columnTypes[colName]
		if !ok {
			colType = arrow.BinaryTypes.String
		}
		fields = append(fields, arrow.Field{Name: colName, Type: colType})

		switch colType.ID() {
		case arrow.FLOAT64:
			b := array.NewFloat64Builder(pool)
			for _, rec := range records {
				val, ok := rec[colName]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					switch v := val.(type) {
					case float64:
						b.Append(v)
					case json.Number:
						f, _ := v.Float64()
						b.Append(f)
					default:
						b.AppendNull()
					}
				}
			}
			columns = append(columns, b.NewArray())
			b.Release()

		case arrow.INT64:
			b := array.NewInt64Builder(pool)
			for _, rec := range records {
				val, ok := rec[colName]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					switch v := val.(type) {
					case float64:
						b.Append(int64(v))
					case json.Number:
						i, _ := v.Int64()
						b.Append(i)
					default:
						b.AppendNull()
					}
				}
			}
			columns = append(columns, b.NewArray())
			b.Release()

		case arrow.BOOL:
			b := array.NewBooleanBuilder(pool)
			for _, rec := range records {
				val, ok := rec[colName]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					bv, ok := val.(bool)
					if ok {
						b.Append(bv)
					} else {
						b.AppendNull()
					}
				}
			}
			columns = append(columns, b.NewArray())
			b.Release()

		default: // STRING
			b := array.NewStringBuilder(pool)
			for _, rec := range records {
				val, ok := rec[colName]
				if !ok || val == nil {
					b.AppendNull()
				} else {
					b.Append(fmt.Sprintf("%v", val))
				}
			}
			columns = append(columns, b.NewArray())
			b.Release()
		}
	}

	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, columns, int64(numRows))
	return NewDataFrame(record), nil
}

// inferArrowType infers an Arrow type from a Go value.
func inferArrowType(val interface{}) arrow.DataType {
	switch v := val.(type) {
	case bool:
		return arrow.FixedWidthTypes.Boolean
	case float64:
		// JSON numbers are float64 by default; check if it's actually an integer
		if v == math.Trunc(v) && !math.IsInf(v, 0) && !math.IsNaN(v) {
			return arrow.PrimitiveTypes.Float64 // Keep as float64 since JSON doesn't distinguish
		}
		return arrow.PrimitiveTypes.Float64
	case string:
		return arrow.BinaryTypes.String
	default:
		return arrow.BinaryTypes.String
	}
}

// dataFrameToJSONRecords converts a DataFrame to a slice of maps.
func dataFrameToJSONRecords(df *DataFrame) []map[string]interface{} {
	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())

	records := make([]map[string]interface{}, numRows)
	for i := 0; i < numRows; i++ {
		records[i] = make(map[string]interface{})
		for j, field := range schema.Fields() {
			col := record.Column(j)
			if col.IsNull(i) {
				records[i][field.Name] = nil
				continue
			}
			switch a := col.(type) {
			case *array.Float64:
				records[i][field.Name] = a.Value(i)
			case *array.Int64:
				records[i][field.Name] = a.Value(i)
			case *array.String:
				records[i][field.Name] = a.Value(i)
			case *array.Boolean:
				records[i][field.Name] = a.Value(i)
			default:
				records[i][field.Name] = fmt.Sprintf("%v", a)
			}
		}
	}
	return records
}
