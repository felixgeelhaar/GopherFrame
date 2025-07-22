package io

import (
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

func TestParquetReader_NewParquetReader(t *testing.T) {
	reader := NewParquetReader()
	if reader == nil {
		t.Error("NewParquetReader should not return nil")
	}
}

func TestParquetWriter_NewParquetWriter(t *testing.T) {
	writer := NewParquetWriter()
	if writer == nil {
		t.Error("NewParquetWriter should not return nil")
	}
}

func TestParquetReader_ReadFile_NotImplemented(t *testing.T) {
	reader := NewParquetReader()
	_, err := reader.ReadFile("test.parquet")
	if err == nil {
		t.Error("Expected error for reading")
	}
	// Could fail at file open or at the not-implemented check, both are valid
	if err != nil {
		// Success - the functionality correctly returns an error
		return
	}
}

func TestParquetWriter_WriteFile_NotImplemented(t *testing.T) {
	writer := NewParquetWriter()

	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.Append(42)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	df := dataframe.NewDataFrame(record)
	defer df.Release()

	err := writer.WriteFile(df, "test.parquet")
	if err == nil {
		t.Error("Expected error for unimplemented functionality")
	}
	if err.Error() != "parquet writing not yet implemented in DDD structure" {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestParquetAPI_Exists(t *testing.T) {
	// Test that the API exists and can be instantiated
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.parquet")

	reader := NewParquetReader()
	writer := NewParquetWriter()

	// Both should be non-nil
	if reader == nil {
		t.Error("NewParquetReader returned nil")
	}
	if writer == nil {
		t.Error("NewParquetWriter returned nil")
	}

	// Test that the methods exist and return expected "not implemented" errors
	_, err := reader.ReadFile(testFile)
	if err == nil {
		t.Error("Expected ReadFile to return error for unimplemented functionality")
	}

	// Create a minimal DataFrame for testing the writer
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)
	builder := array.NewInt64Builder(pool)
	builder.Append(1)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 1)
	df := dataframe.NewDataFrame(record)
	defer df.Release()

	err = writer.WriteFile(df, testFile)
	if err == nil {
		t.Error("Expected WriteFile to return error for unimplemented functionality")
	}
}
