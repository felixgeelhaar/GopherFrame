// Package io contains infrastructure for file I/O operations.
// This is part of the infrastructure layer that handles external concerns.
package io

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet"
	"github.com/apache/arrow-go/v18/parquet/file"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
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

// ParquetReader provides functionality to read Parquet files into DataFrames.
type ParquetReader struct{}

// NewParquetReader creates a new ParquetReader.
func NewParquetReader() *ParquetReader {
	return &ParquetReader{}
}

// ReadFile reads a Parquet file and returns a DataFrame.
func (r *ParquetReader) ReadFile(filename string) (*dataframe.DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	// Open the Parquet file
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open Parquet file: %w", err)
	}
	defer func() { _ = f.Close() }()

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

	// Handle empty tables
	if table.NumRows() == 0 {
		// Create empty arrays for each column
		schema := table.Schema()
		emptyArrays := make([]arrow.Array, len(schema.Fields()))
		pool := memory.NewGoAllocator()

		for i, field := range schema.Fields() {
			switch field.Type.ID() {
			case arrow.INT64:
				builder := array.NewInt64Builder(pool)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			case arrow.FLOAT64:
				builder := array.NewFloat64Builder(pool)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			case arrow.STRING:
				builder := array.NewStringBuilder(pool)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			case arrow.BOOL:
				builder := array.NewBooleanBuilder(pool)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			case arrow.DATE32:
				builder := array.NewDate32Builder(pool)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			case arrow.DATE64:
				builder := array.NewDate64Builder(pool)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			case arrow.TIMESTAMP:
				timestampType := field.Type.(*arrow.TimestampType)
				builder := array.NewTimestampBuilder(pool, timestampType)
				emptyArrays[i] = builder.NewArray()
				builder.Release()
			default:
				return nil, fmt.Errorf("unsupported data type for empty table: %s", field.Type)
			}
		}

		emptyRecord := array.NewRecord(schema, emptyArrays, 0)
		return dataframe.NewDataFrame(emptyRecord), nil
	}

	// Convert table to record
	tr := array.NewTableReader(table, -1)
	defer tr.Release()

	if !tr.Next() {
		return nil, fmt.Errorf("no records in Parquet file")
	}

	record := tr.Record()
	record.Retain() // Keep the record alive

	return dataframe.NewDataFrame(record), nil
}

// ParquetWriter provides functionality to write DataFrames to Parquet files.
type ParquetWriter struct{}

// NewParquetWriter creates a new ParquetWriter.
func NewParquetWriter() *ParquetWriter {
	return &ParquetWriter{}
}

// WriteFile writes a DataFrame to a Parquet file.
func (w *ParquetWriter) WriteFile(df *dataframe.DataFrame, filename string) error {
	if df == nil {
		return fmt.Errorf("DataFrame cannot be nil")
	}

	if err := validateFilePath(filename); err != nil {
		return err
	}

	// Create the output file
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Get the Arrow record from the DataFrame
	record := df.Record()

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
	defer func() { _ = writer.Close() }()

	// Write the table with proper chunk size
	chunkSize := table.NumRows()
	if chunkSize <= 0 {
		chunkSize = 1000 // Default chunk size for empty tables
	}
	if err := writer.WriteTable(table, chunkSize); err != nil {
		return fmt.Errorf("failed to write table: %w", err)
	}

	return nil
}
