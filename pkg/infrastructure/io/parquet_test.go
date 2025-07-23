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

func TestParquetReader_ReadFile_NonExistentFile(t *testing.T) {
	reader := NewParquetReader()
	_, err := reader.ReadFile("nonexistent.parquet")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
	// Should get a file not found error
}

func TestParquetWriter_WriteFile_NilDataFrame(t *testing.T) {
	writer := NewParquetWriter()

	err := writer.WriteFile(nil, "test.parquet")
	if err == nil {
		t.Error("Expected error for nil DataFrame")
	}
	if err.Error() != "DataFrame cannot be nil" {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestParquetAPI_WorksCorrectly(t *testing.T) {
	// Test that the API exists and works correctly
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

	// Test that reading non-existent file returns error
	_, err := reader.ReadFile(testFile)
	if err == nil {
		t.Error("Expected ReadFile to return error for non-existent file")
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

	// Test writing works correctly
	err = writer.WriteFile(df, testFile)
	if err != nil {
		t.Errorf("WriteFile failed: %v", err)
	}

	// Test reading the written file works
	readDF, err := reader.ReadFile(testFile)
	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}
	if readDF != nil {
		defer readDF.Release()
		if readDF.NumRows() != 1 {
			t.Errorf("Expected 1 row, got %d", readDF.NumRows())
		}
	}
}
