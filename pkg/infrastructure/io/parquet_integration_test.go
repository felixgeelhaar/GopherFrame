package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

// TestParquetRoundTrip tests writing and then reading a Parquet file
func TestParquetRoundTrip(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test DataFrame with various data types
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
			{Name: "active", Type: arrow.FixedWidthTypes.Boolean},
		},
		nil,
	)

	// Build test data
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{95.5, 87.2, 92.1, 88.9, 94.3}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	activeBuilder := array.NewBooleanBuilder(pool)
	activeBuilder.AppendValues([]bool{true, false, true, true, false}, nil)
	activeArray := activeBuilder.NewArray()
	defer activeArray.Release()

	// Create record and DataFrame
	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray, activeArray}, 5)
	defer record.Release()

	originalDF := dataframe.NewDataFrame(record)
	defer originalDF.Release()

	// Create temporary file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test.parquet")

	// Test writing
	writer := NewParquetWriter()
	err := writer.WriteFile(originalDF, filename)
	if err != nil {
		t.Fatalf("Failed to write Parquet file: %v", err)
	}

	// Verify file exists
	if _, statErr := os.Stat(filename); os.IsNotExist(statErr) {
		t.Fatal("Parquet file was not created")
	}

	// Test reading
	reader := NewParquetReader()
	readDF, err := reader.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read Parquet file: %v", err)
	}
	defer readDF.Release()

	// Verify data integrity
	if readDF.NumRows() != originalDF.NumRows() {
		t.Errorf("Row count mismatch: expected %d, got %d", originalDF.NumRows(), readDF.NumRows())
	}

	if readDF.NumCols() != originalDF.NumCols() {
		t.Errorf("Column count mismatch: expected %d, got %d", originalDF.NumCols(), readDF.NumCols())
	}

	// Verify schema
	originalSchema := originalDF.Schema()
	readSchema := readDF.Schema()

	if len(originalSchema.Fields()) != len(readSchema.Fields()) {
		t.Errorf("Field count mismatch: expected %d, got %d", len(originalSchema.Fields()), len(readSchema.Fields()))
	}

	for i, originalField := range originalSchema.Fields() {
		readField := readSchema.Field(i)
		if originalField.Name != readField.Name {
			t.Errorf("Field name mismatch at index %d: expected %s, got %s", i, originalField.Name, readField.Name)
		}
		if !arrow.TypeEqual(originalField.Type, readField.Type) {
			t.Errorf("Field type mismatch at index %d: expected %s, got %s", i, originalField.Type, readField.Type)
		}
	}

	// Verify data values for each column
	testCases := []struct {
		columnName      string
		expectedType    arrow.Type
		expectedInt64   []int64
		expectedString  []string
		expectedFloat64 []float64
		expectedBool    []bool
	}{
		{
			columnName:    "id",
			expectedType:  arrow.INT64,
			expectedInt64: []int64{1, 2, 3, 4, 5},
		},
		{
			columnName:     "name",
			expectedType:   arrow.STRING,
			expectedString: []string{"Alice", "Bob", "Charlie", "Diana", "Eve"},
		},
		{
			columnName:      "score",
			expectedType:    arrow.FLOAT64,
			expectedFloat64: []float64{95.5, 87.2, 92.1, 88.9, 94.3},
		},
		{
			columnName:   "active",
			expectedType: arrow.BOOL,
			expectedBool: []bool{true, false, true, true, false},
		},
	}

	for _, tc := range testCases {
		t.Run("verify_"+tc.columnName, func(t *testing.T) {
			record := readDF.Record()
			schema := record.Schema()

			// Find column index
			colIdx := -1
			for i, field := range schema.Fields() {
				if field.Name == tc.columnName {
					colIdx = i
					break
				}
			}
			if colIdx == -1 {
				t.Fatalf("Column %s not found", tc.columnName)
			}

			column := record.Column(colIdx)
			if column.DataType().ID() != tc.expectedType {
				t.Errorf("Column %s type mismatch: expected %s, got %s", tc.columnName, tc.expectedType, column.DataType().ID())
			}

			switch tc.expectedType {
			case arrow.INT64:
				int64Arr := column.(*array.Int64)
				for i, expected := range tc.expectedInt64 {
					if int64Arr.Value(i) != expected {
						t.Errorf("Column %s row %d: expected %d, got %d", tc.columnName, i, expected, int64Arr.Value(i))
					}
				}
			case arrow.STRING:
				stringArr := column.(*array.String)
				for i, expected := range tc.expectedString {
					if stringArr.Value(i) != expected {
						t.Errorf("Column %s row %d: expected %s, got %s", tc.columnName, i, expected, stringArr.Value(i))
					}
				}
			case arrow.FLOAT64:
				float64Arr := column.(*array.Float64)
				for i, expected := range tc.expectedFloat64 {
					if float64Arr.Value(i) != expected {
						t.Errorf("Column %s row %d: expected %f, got %f", tc.columnName, i, expected, float64Arr.Value(i))
					}
				}
			case arrow.BOOL:
				boolArr := column.(*array.Boolean)
				for i, expected := range tc.expectedBool {
					if boolArr.Value(i) != expected {
						t.Errorf("Column %s row %d: expected %t, got %t", tc.columnName, i, expected, boolArr.Value(i))
					}
				}
			}
		})
	}
}

func TestParquetReader_ErrorCases(t *testing.T) {
	reader := NewParquetReader()

	// Test reading non-existent file
	_, err := reader.ReadFile("nonexistent.parquet")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}

	// Test reading invalid file
	tempDir := t.TempDir()
	invalidFile := filepath.Join(tempDir, "invalid.parquet")

	// Create an invalid parquet file (just text)
	err = os.WriteFile(invalidFile, []byte("this is not a parquet file"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err = reader.ReadFile(invalidFile)
	if err == nil {
		t.Error("Expected error for invalid parquet file")
	}
}

func TestParquetWriter_ErrorCases(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "test", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{1, 2, 3}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	writer := NewParquetWriter()

	// Test writing to invalid path
	err := writer.WriteFile(df, "/invalid/path/test.parquet")
	if err == nil {
		t.Error("Expected error for invalid file path")
	}

	// Test writing nil DataFrame
	err = writer.WriteFile(nil, "test.parquet")
	if err == nil {
		t.Error("Expected error for nil DataFrame")
	}
}

func TestParquetRoundTrip_EmptyDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create empty DataFrame
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Create empty arrays
	idBuilder := array.NewInt64Builder(pool)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray}, 0)
	defer record.Release()

	emptyDF := dataframe.NewDataFrame(record)
	defer emptyDF.Release()

	// Test round trip
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "empty.parquet")

	writer := NewParquetWriter()
	err := writer.WriteFile(emptyDF, filename)
	if err != nil {
		t.Fatalf("Failed to write empty DataFrame: %v", err)
	}

	reader := NewParquetReader()
	readDF, err := reader.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read empty DataFrame: %v", err)
	}
	defer readDF.Release()

	// Verify empty DataFrame properties
	if readDF.NumRows() != 0 {
		t.Errorf("Expected 0 rows, got %d", readDF.NumRows())
	}

	if readDF.NumCols() != 2 {
		t.Errorf("Expected 2 columns, got %d", readDF.NumCols())
	}
}

func TestParquetRoundTrip_LargeDataSet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large dataset test in short mode")
	}

	pool := memory.NewGoAllocator()

	// Create larger DataFrame (1000 rows)
	numRows := 1000
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build large dataset
	idBuilder := array.NewInt64Builder(pool)
	valueBuilder := array.NewFloat64Builder(pool)

	for i := 0; i < numRows; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i) * 1.5)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, int64(numRows))
	defer record.Release()

	largeDF := dataframe.NewDataFrame(record)
	defer largeDF.Release()

	// Test round trip
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "large.parquet")

	writer := NewParquetWriter()
	err := writer.WriteFile(largeDF, filename)
	if err != nil {
		t.Fatalf("Failed to write large DataFrame: %v", err)
	}

	reader := NewParquetReader()
	readDF, err := reader.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read large DataFrame: %v", err)
	}
	defer readDF.Release()

	// Verify size
	if readDF.NumRows() != int64(numRows) {
		t.Errorf("Expected %d rows, got %d", numRows, readDF.NumRows())
	}

	// Verify a few sample values
	record = readDF.Record()
	idArr := record.Column(0).(*array.Int64) // id column is at index 0

	if idArr.Value(0) != 0 {
		t.Errorf("Expected first ID to be 0, got %d", idArr.Value(0))
	}
	if idArr.Value(numRows-1) != int64(numRows-1) {
		t.Errorf("Expected last ID to be %d, got %d", numRows-1, idArr.Value(numRows-1))
	}

	valueArr := record.Column(1).(*array.Float64) // value column is at index 1

	if valueArr.Value(0) != 0.0 {
		t.Errorf("Expected first value to be 0.0, got %f", valueArr.Value(0))
	}
	expectedLast := float64(numRows-1) * 1.5
	if valueArr.Value(numRows-1) != expectedLast {
		t.Errorf("Expected last value to be %f, got %f", expectedLast, valueArr.Value(numRows-1))
	}
}
