// Package gopherframe provides I/O functions for reading and writing DataFrames
// in various formats including Parquet, CSV, and Arrow IPC.
package gopherframe

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/ipc"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet"
	"github.com/apache/arrow-go/v18/parquet/file"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
)

// validateFilePath performs basic security validation on file paths
// to prevent directory traversal attacks
func validateFilePath(filename string) error {
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	// Clean the path to resolve any .. components
	cleanPath := filepath.Clean(filename)

	// Check for directory traversal attempts
	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("invalid file path: directory traversal detected")
	}

	return nil
}

// ReadParquet reads a DataFrame from a Parquet file.
// Returns a new DataFrame with the data from the file.
func ReadParquet(filename string) (*DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	// Open the Parquet file
	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, fmt.Errorf("failed to open Parquet file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Skip file info - not needed for reading

	// Create Parquet file reader
	parquetReader, err := file.NewParquetReader(f, file.WithReadProps(parquet.NewReaderProperties(memory.DefaultAllocator)))
	if err != nil {
		return nil, fmt.Errorf("failed to create Parquet reader: %w", err)
	}
	defer func() { _ = parquetReader.Close() }()

	// Create Arrow file reader from Parquet
	arrowReader, err := pqarrow.NewFileReader(parquetReader, pqarrow.ArrowReadProperties{}, memory.DefaultAllocator)
	if err != nil {
		return nil, fmt.Errorf("failed to create Arrow reader: %w", err)
	}

	// Read all row groups into a table with context
	ctx := context.Background()
	table, err := arrowReader.ReadTable(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read table: %w", err)
	}
	defer table.Release()

	// Convert table to record
	tr := array.NewTableReader(table, -1)
	defer tr.Release()

	if !tr.Next() {
		return nil, fmt.Errorf("no records in Parquet file")
	}

	record := tr.Record()
	record.Retain() // Keep the record alive

	return NewDataFrame(record), nil
}

// WriteParquet writes a DataFrame to a Parquet file.
func WriteParquet(df *DataFrame, filename string) error {
	if df.Err() != nil {
		return fmt.Errorf("DataFrame has error: %w", df.Err())
	}

	if err := validateFilePath(filename); err != nil {
		return err
	}

	// Create the output file
	f, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Get the Arrow record from the DataFrame
	record := df.coreDF.Record()

	// Create Arrow table from record
	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})
	defer table.Release()

	// Set up Parquet writer properties
	writerProps := parquet.NewWriterProperties()
	arrowProps := pqarrow.DefaultWriterProps()

	// Create Parquet writer
	writer, err := pqarrow.NewFileWriter(record.Schema(), f, writerProps, arrowProps)
	if err != nil {
		return fmt.Errorf("failed to create Parquet writer: %w", err)
	}
	defer func() {
		if err := writer.Close(); err != nil {
			// Log error but don't override the main error
			fmt.Fprintf(os.Stderr, "warning: failed to close Parquet writer: %v\n", err)
		}
	}()

	// Write the table with proper chunk size
	chunkSize := table.NumRows()
	if chunkSize <= 0 {
		chunkSize = 1000 // Default chunk size
	}
	if err := writer.WriteTable(table, chunkSize); err != nil {
		return fmt.Errorf("failed to write table: %w", err)
	}

	return nil
}

// ReadCSV reads a DataFrame from a CSV file.
// This implementation attempts to infer column types from the data.
func ReadCSV(filename string) (*DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	// Read CSV data
	header, records, err := readCSVData(filename)
	if err != nil {
		return nil, err
	}

	// Infer column types by checking first non-empty value in each column
	columnTypes := make([]arrow.DataType, len(header))
	for i := range header {
		columnTypes[i] = inferColumnType(records, i)
	}

	// Build Arrow arrays
	pool := memory.NewGoAllocator()
	arrays := make([]arrow.Array, len(header))
	fields := make([]arrow.Field, len(header))

	for colIdx, colName := range header {
		fields[colIdx] = arrow.Field{Name: colName, Type: columnTypes[colIdx]}
		arrays[colIdx] = buildColumnArray(pool, columnTypes[colIdx], records, colIdx)
	}

	// Create schema and record
	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, arrays, int64(len(records)))

	// Release arrays as they're now owned by the record
	for _, arr := range arrays {
		arr.Release()
	}

	return NewDataFrame(record), nil
}

// inferColumnType attempts to determine the column type by sampling values
func inferColumnType(records [][]string, colIdx int) arrow.DataType {
	// Sample up to 100 non-empty values to determine type
	sampleSize := 100
	if len(records) < sampleSize {
		sampleSize = len(records)
	}

	hasInt := true
	hasFloat := true

	for i := 0; i < sampleSize; i++ {
		if colIdx >= len(records[i]) || records[i][colIdx] == "" {
			continue
		}

		val := records[i][colIdx]

		// Try parsing as int
		if _, err := strconv.ParseInt(val, 10, 64); err != nil {
			hasInt = false
		}

		// Try parsing as float
		if _, err := strconv.ParseFloat(val, 64); err != nil {
			hasFloat = false
		}

		// If neither int nor float, it's a string
		if !hasInt && !hasFloat {
			return arrow.BinaryTypes.String
		}
	}

	// Prefer int over float if possible
	if hasInt {
		return arrow.PrimitiveTypes.Int64
	}
	if hasFloat {
		return arrow.PrimitiveTypes.Float64
	}

	// Default to string
	return arrow.BinaryTypes.String
}

// WriteCSV writes a DataFrame to a CSV file.
func WriteCSV(df *DataFrame, filename string) error {
	if df.Err() != nil {
		return fmt.Errorf("DataFrame has error: %w", df.Err())
	}

	if err := validateFilePath(filename); err != nil {
		return err
	}

	// Create the output file
	f, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Create CSV writer
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Write header row with column names
	header := df.ColumnNames()
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Get number of rows
	numRows := int(df.NumRows())
	numCols := int(df.NumCols())

	// Write data rows
	for i := 0; i < numRows; i++ {
		row := make([]string, numCols)

		// Convert each column value to string
		for j := 0; j < numCols; j++ {
			series, err := df.coreDF.Column(header[j])
			if err != nil {
				return fmt.Errorf("failed to get column %s: %w", header[j], err)
			}
			defer series.Release()

			// Check if value is null
			if series.IsNull(i) {
				row[j] = ""
				continue
			}

			// Convert based on type
			switch series.DataType().ID() {
			case arrow.INT64:
				val, err := series.GetInt64(i)
				if err != nil {
					return fmt.Errorf("failed to get int64 value: %w", err)
				}
				row[j] = strconv.FormatInt(val, 10)
			case arrow.FLOAT64:
				val, err := series.GetFloat64(i)
				if err != nil {
					return fmt.Errorf("failed to get float64 value: %w", err)
				}
				row[j] = strconv.FormatFloat(val, 'f', -1, 64)
			case arrow.STRING:
				val, err := series.GetString(i)
				if err != nil {
					return fmt.Errorf("failed to get string value: %w", err)
				}
				row[j] = val
			default:
				return fmt.Errorf("unsupported type for CSV: %s", series.DataType())
			}
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	return nil
}

// ReadArrowIPC reads a DataFrame from an Arrow IPC file.
func ReadArrowIPC(filename string) (*DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	// Open the file
	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, fmt.Errorf("failed to open Arrow IPC file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Create Arrow IPC file reader
	reader, err := ipc.NewFileReader(f)
	if err != nil {
		return nil, fmt.Errorf("failed to create Arrow IPC reader: %w", err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close Arrow reader: %v\n", err)
		}
	}()

	// Read the first record
	if reader.NumRecords() == 0 {
		return nil, fmt.Errorf("no records in Arrow IPC file")
	}

	record, err := reader.RecordAt(0)
	if err != nil {
		return nil, fmt.Errorf("failed to read record: %w", err)
	}

	return NewDataFrame(record), nil
}

// WriteArrowIPC writes a DataFrame to an Arrow IPC file.
func WriteArrowIPC(df *DataFrame, filename string) error {
	if df.Err() != nil {
		return fmt.Errorf("DataFrame has error: %w", df.Err())
	}

	if err := validateFilePath(filename); err != nil {
		return err
	}

	// Create the output file
	f, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Get the Arrow record from the DataFrame
	record := df.coreDF.Record()

	// Create Arrow IPC file writer
	writer, err := ipc.NewFileWriter(f, ipc.WithSchema(record.Schema()))
	if err != nil {
		return fmt.Errorf("failed to create Arrow IPC writer: %w", err)
	}
	defer func() {
		if err := writer.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close Arrow writer: %v\n", err)
		}
	}()

	// Write the record
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write record: %w", err)
	}

	return nil
}

// readCSVData reads and parses CSV file data
func readCSVData(filename string) ([]string, [][]string, error) {
	// Open the CSV file
	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Create CSV reader
	reader := csv.NewReader(f)

	// Read header row
	header, err := reader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Read all records to determine types and build columns
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV records: %w", err)
	}

	if len(records) == 0 {
		return nil, nil, fmt.Errorf("CSV file has no data rows")
	}

	return header, records, nil
}

// buildColumnArray creates an Arrow array for a specific column
func buildColumnArray(pool memory.Allocator, dataType arrow.DataType, records [][]string, colIdx int) arrow.Array {
	switch dataType.ID() {
	case arrow.INT64:
		return buildInt64Array(pool, records, colIdx)
	case arrow.FLOAT64:
		return buildFloat64Array(pool, records, colIdx)
	default: // STRING
		return buildStringArray(pool, records, colIdx)
	}
}

// buildInt64Array creates an int64 Arrow array
func buildInt64Array(pool memory.Allocator, records [][]string, colIdx int) arrow.Array {
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for _, row := range records {
		if colIdx < len(row) && row[colIdx] != "" {
			val, err := strconv.ParseInt(row[colIdx], 10, 64)
			if err != nil {
				builder.AppendNull()
			} else {
				builder.Append(val)
			}
		} else {
			builder.AppendNull()
		}
	}

	return builder.NewArray()
}

// buildFloat64Array creates a float64 Arrow array
func buildFloat64Array(pool memory.Allocator, records [][]string, colIdx int) arrow.Array {
	builder := array.NewFloat64Builder(pool)
	defer builder.Release()

	for _, row := range records {
		if colIdx < len(row) && row[colIdx] != "" {
			val, err := strconv.ParseFloat(row[colIdx], 64)
			if err != nil {
				builder.AppendNull()
			} else {
				builder.Append(val)
			}
		} else {
			builder.AppendNull()
		}
	}

	return builder.NewArray()
}

// buildStringArray creates a string Arrow array
func buildStringArray(pool memory.Allocator, records [][]string, colIdx int) arrow.Array {
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for _, row := range records {
		if colIdx < len(row) {
			builder.Append(row[colIdx])
		} else {
			builder.Append("")
		}
	}

	return builder.NewArray()
}
