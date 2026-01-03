package gopherframe

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestParquetRoundTrip(t *testing.T) {
	// Create test DataFrame
	df := createSampleDataFrame()
	defer df.Release()

	// Create temp directory for test files
	tempDir, err := os.MkdirTemp("", "gopherframe_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	parquetFile := filepath.Join(tempDir, "test.parquet")

	// Test WriteParquet
	err = WriteParquet(df, parquetFile)
	if err != nil {
		t.Fatalf("Failed to write Parquet: %v", err)
	}

	// Verify file exists
	if _, statErr := os.Stat(parquetFile); os.IsNotExist(statErr) {
		t.Fatal("Parquet file was not created")
	}

	// Test ReadParquet
	readDF, err := ReadParquet(parquetFile)
	if err != nil {
		t.Fatalf("Failed to read Parquet: %v", err)
	}
	defer readDF.Release()

	// Verify data integrity
	if readDF.NumRows() != df.NumRows() {
		t.Errorf("Row count mismatch: expected %d, got %d", df.NumRows(), readDF.NumRows())
	}

	if readDF.NumCols() != df.NumCols() {
		t.Errorf("Column count mismatch: expected %d, got %d", df.NumCols(), readDF.NumCols())
	}

	// Verify column names
	originalNames := df.ColumnNames()
	readNames := readDF.ColumnNames()

	for i, originalName := range originalNames {
		if readNames[i] != originalName {
			t.Errorf("Column name mismatch at index %d: expected %s, got %s", i, originalName, readNames[i])
		}
	}
}

func TestParquetWithRealData(t *testing.T) {
	// Create a larger DataFrame for real-world testing
	df := createLargerDataFrame(1000)
	defer df.Release()

	tempDir, err := os.MkdirTemp("", "gopherframe_large_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	parquetFile := filepath.Join(tempDir, "large.parquet")

	// Write and read back
	err = WriteParquet(df, parquetFile)
	if err != nil {
		t.Fatalf("Failed to write large Parquet: %v", err)
	}

	readDF, err := ReadParquet(parquetFile)
	if err != nil {
		t.Fatalf("Failed to read large Parquet: %v", err)
	}
	defer readDF.Release()

	// Test operations on read DataFrame
	filtered := readDF.Filter(Col("score").Gt(Lit(90.0)))
	defer filtered.Release()

	if filtered.Err() != nil {
		t.Errorf("Filter operation failed on read DataFrame: %v", filtered.Err())
	}

	if filtered.NumRows() == 0 {
		t.Error("No rows after filter - expected some high-scoring records")
	}
}

func TestCSVRoundTrip(t *testing.T) {
	// Create test DataFrame
	df := createSampleDataFrame()
	defer df.Release()

	tempDir, err := os.MkdirTemp("", "gopherframe_csv_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	csvFile := filepath.Join(tempDir, "test.csv")

	// Test WriteCSV
	err = WriteCSV(df, csvFile)
	if err != nil {
		t.Fatalf("Failed to write CSV: %v", err)
	}

	// Test ReadCSV
	readDF, err := ReadCSV(csvFile)
	if err != nil {
		t.Fatalf("Failed to read CSV: %v", err)
	}
	defer readDF.Release()

	// Verify basic structure (CSV may have type differences)
	if readDF.NumRows() != df.NumRows() {
		t.Errorf("Row count mismatch: expected %d, got %d", df.NumRows(), readDF.NumRows())
	}

	if readDF.NumCols() != df.NumCols() {
		t.Errorf("Column count mismatch: expected %d, got %d", df.NumCols(), readDF.NumCols())
	}
}

func TestStorageBackendIntegration(t *testing.T) {
	// Create DataFrame
	df := createSampleDataFrame()
	defer df.Release()

	// Test using storage backend directly
	tempDir, err := os.MkdirTemp("", "gopherframe_storage_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// Test with Arrow storage backend
	arrowFile := filepath.Join(tempDir, "test.arrow")
	err = WriteArrowIPC(df, arrowFile)
	if err != nil {
		t.Fatalf("Failed to write Arrow IPC: %v", err)
	}

	readDF, err := ReadArrowIPC(arrowFile)
	if err != nil {
		t.Fatalf("Failed to read Arrow IPC: %v", err)
	}
	defer readDF.Release()

	// Verify perfect fidelity for Arrow format
	if readDF.NumRows() != df.NumRows() || readDF.NumCols() != df.NumCols() {
		t.Error("Arrow IPC should preserve exact structure")
	}
}

// Helper function to create larger DataFrame for testing
func createLargerDataFrame(numRows int) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	pool := memory.NewGoAllocator()

	// Create ID column
	idBuilder := array.NewInt64Builder(pool)
	for i := 0; i < numRows; i++ {
		idBuilder.Append(int64(i + 1))
	}
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	// Create score column with realistic distribution
	scoreBuilder := array.NewFloat64Builder(pool)
	for i := 0; i < numRows; i++ {
		// Generate scores between 70-100 with some high performers
		score := 70.0 + float64(i%31)
		if i%10 == 0 {
			score += 20.0 // Make some high performers
		}
		scoreBuilder.Append(score)
	}
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	// Create category column
	categoryBuilder := array.NewStringBuilder(pool)
	categories := []string{"A", "B", "C", "D"}
	for i := 0; i < numRows; i++ {
		categoryBuilder.Append(categories[i%len(categories)])
	}
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	// Create record
	record := array.NewRecord(schema, []arrow.Array{idArray, scoreArray, categoryArray}, int64(numRows))
	defer record.Release()

	return NewDataFrame(record)
}
