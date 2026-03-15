package gopherframe

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// DataFrameIterator allows processing DataFrames in chunks to handle
// datasets larger than available memory.
type DataFrameIterator struct {
	chunks []*DataFrame
	index  int
}

// Next returns the next chunk DataFrame, or nil when exhausted.
func (it *DataFrameIterator) Next() *DataFrame {
	if it.index >= len(it.chunks) {
		return nil
	}
	df := it.chunks[it.index]
	it.index++
	return df
}

// HasNext returns true if there are more chunks available.
func (it *DataFrameIterator) HasNext() bool {
	return it.index < len(it.chunks)
}

// Reset resets the iterator to the beginning.
func (it *DataFrameIterator) Reset() {
	it.index = 0
}

// Len returns the total number of chunks.
func (it *DataFrameIterator) Len() int {
	return len(it.chunks)
}

// ReadCSVChunked reads a CSV file in chunks of the specified size.
// Returns an iterator that yields DataFrames one chunk at a time.
func ReadCSVChunked(filename string, chunkSize int) (*DataFrameIterator, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}
	if chunkSize <= 0 {
		return nil, fmt.Errorf("chunk size must be positive")
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer func() { _ = f.Close() }()

	reader := csv.NewReader(f)

	// Read header
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	var chunks []*DataFrame

	for {
		// Read chunkSize rows
		var rows [][]string
		for i := 0; i < chunkSize; i++ {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("failed to read CSV row: %w", err)
			}
			rows = append(rows, row)
		}

		if len(rows) == 0 {
			break
		}

		df, err := csvRowsToDataFrame(header, rows)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, df)
	}

	return &DataFrameIterator{chunks: chunks}, nil
}

// ForEachChunk applies a function to each chunk in the iterator.
// This is a convenience method for processing all chunks.
func (it *DataFrameIterator) ForEachChunk(fn func(chunk *DataFrame) error) error {
	it.Reset()
	for it.HasNext() {
		chunk := it.Next()
		if err := fn(chunk); err != nil {
			return err
		}
	}
	return nil
}

// Collect reads all chunks and concatenates them into a single DataFrame.
// Use with caution for large datasets — this loads everything into memory.
func (it *DataFrameIterator) Collect() (*DataFrame, error) {
	it.Reset()

	if !it.HasNext() {
		return nil, fmt.Errorf("no chunks to collect")
	}

	first := it.Next()
	if !it.HasNext() {
		return first, nil
	}

	// For simplicity, collect all rows and rebuild
	// A more efficient implementation would use Arrow's RecordBatch concatenation
	pool := memory.NewGoAllocator()
	schema := first.coreDF.Schema()
	numCols := int(first.NumCols())

	// Collect all values per column across chunks
	colValues := make([][]interface{}, numCols)
	for i := range colValues {
		colValues[i] = make([]interface{}, 0)
	}

	// Helper to extract all values from a DataFrame
	extractValues := func(df *DataFrame) {
		record := df.coreDF.Record()
		for row := 0; row < int(df.NumRows()); row++ {
			for col := 0; col < numCols; col++ {
				arr := record.Column(col)
				if arr.IsNull(row) {
					colValues[col] = append(colValues[col], nil)
				} else {
					colValues[col] = append(colValues[col], getTypedValue(arr, row))
				}
			}
		}
	}

	// Process first chunk
	extractValues(first)

	// Process remaining chunks
	for it.HasNext() {
		chunk := it.Next()
		extractValues(chunk)
	}

	totalRows := len(colValues[0])

	// Build result arrays
	var resultColumns []arrow.Array
	for i, field := range schema.Fields() {
		switch field.Type.ID() {
		case arrow.FLOAT64:
			b := array.NewFloat64Builder(pool)
			for _, v := range colValues[i] {
				if v == nil {
					b.AppendNull()
				} else {
					b.Append(v.(float64))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		case arrow.INT64:
			b := array.NewInt64Builder(pool)
			for _, v := range colValues[i] {
				if v == nil {
					b.AppendNull()
				} else {
					b.Append(v.(int64))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		case arrow.STRING:
			b := array.NewStringBuilder(pool)
			for _, v := range colValues[i] {
				if v == nil {
					b.AppendNull()
				} else {
					b.Append(v.(string))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		default:
			b := array.NewStringBuilder(pool)
			for _, v := range colValues[i] {
				if v == nil {
					b.AppendNull()
				} else {
					b.Append(fmt.Sprintf("%v", v))
				}
			}
			resultColumns = append(resultColumns, b.NewArray())
			b.Release()
		}
	}

	record := array.NewRecord(schema, resultColumns, int64(totalRows))
	return NewDataFrame(record), nil
}

// csvRowsToDataFrame converts CSV rows to a DataFrame with type inference.
func csvRowsToDataFrame(header []string, rows [][]string) (*DataFrame, error) {
	pool := memory.NewGoAllocator()
	numCols := len(header)

	// Infer types from first row
	types := make([]arrow.DataType, numCols)
	for i := range types {
		types[i] = arrow.BinaryTypes.String // default
		if len(rows) > 0 && i < len(rows[0]) {
			if _, err := strconv.ParseFloat(rows[0][i], 64); err == nil {
				types[i] = arrow.PrimitiveTypes.Float64
			}
		}
	}

	var fields []arrow.Field
	var columns []arrow.Array

	for i := 0; i < numCols; i++ {
		fields = append(fields, arrow.Field{Name: header[i], Type: types[i]})

		switch types[i].ID() {
		case arrow.FLOAT64:
			b := array.NewFloat64Builder(pool)
			for _, row := range rows {
				if i < len(row) {
					if v, err := strconv.ParseFloat(row[i], 64); err == nil {
						b.Append(v)
					} else {
						b.AppendNull()
					}
				} else {
					b.AppendNull()
				}
			}
			columns = append(columns, b.NewArray())
			b.Release()
		default:
			b := array.NewStringBuilder(pool)
			for _, row := range rows {
				if i < len(row) {
					b.Append(row[i])
				} else {
					b.AppendNull()
				}
			}
			columns = append(columns, b.NewArray())
			b.Release()
		}
	}

	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, columns, int64(len(rows)))
	return NewDataFrame(record), nil
}
