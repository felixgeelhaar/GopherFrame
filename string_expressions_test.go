package gopherframe

import (
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a DataFrame with string data
func createStringDataFrame(pool memory.Allocator, columnName string, values []string) *DataFrame {
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	for _, val := range values {
		builder.Append(val)
	}

	strArray := builder.NewArray()
	fields := []arrow.Field{{Name: columnName, Type: arrow.BinaryTypes.String}}
	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, int64(len(values)))

	strArray.Release() // Record takes ownership
	return NewDataFrame(record)
}

func TestStringContains_BasicFunctionality(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with sample text data
	texts := []string{
		"Hello World",
		"Go programming",
		"Apache Arrow",
		"Data processing",
		"Machine learning",
	}
	df := createStringDataFrame(pool, "text", texts)
	defer df.Release()

	// Test contains with literal
	result := df.Filter(df.Col("text").Contains(expr.Lit("Go")))
	require.NoError(t, result.Err())
	defer result.Release()

	// Should only match "Go programming"
	assert.Equal(t, int64(1), result.NumRows())

	record := result.Record()
	textCol := record.Column(0).(*array.String)
	assert.Equal(t, "Go programming", textCol.Value(0))
}

func TestStringStartsWith_BasicFunctionality(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with sample data
	texts := []string{
		"Hello World",
		"Hello there",
		"Hi everyone",
		"Goodbye",
		"Hello again",
	}
	df := createStringDataFrame(pool, "greeting", texts)
	defer df.Release()

	// Test starts with literal
	result := df.Filter(df.Col("greeting").StartsWith(expr.Lit("Hello")))
	require.NoError(t, result.Err())
	defer result.Release()

	// Should match 3 entries that start with "Hello"
	assert.Equal(t, int64(3), result.NumRows())

	record := result.Record()
	textCol := record.Column(0).(*array.String)
	assert.Equal(t, "Hello World", textCol.Value(0))
	assert.Equal(t, "Hello there", textCol.Value(1))
	assert.Equal(t, "Hello again", textCol.Value(2))
}

func TestStringEndsWith_BasicFunctionality(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with file names
	filenames := []string{
		"document.pdf",
		"image.jpg",
		"script.py",
		"data.csv",
		"archive.zip",
		"photo.jpg",
	}
	df := createStringDataFrame(pool, "filename", filenames)
	defer df.Release()

	// Test ends with literal
	result := df.Filter(df.Col("filename").EndsWith(expr.Lit(".jpg")))
	require.NoError(t, result.Err())
	defer result.Release()

	// Should match 2 JPG files
	assert.Equal(t, int64(2), result.NumRows())

	record := result.Record()
	filenameCol := record.Column(0).(*array.String)
	assert.Equal(t, "image.jpg", filenameCol.Value(0))
	assert.Equal(t, "photo.jpg", filenameCol.Value(1))
}

func TestStringOperations_WithColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with text data
	texts := []string{
		"Hello World",
		"Go programming",
		"Apache Arrow",
		"Data processing",
	}
	df := createStringDataFrame(pool, "text", texts)
	defer df.Release()

	// Add a boolean column indicating if text contains "o"
	result := df.WithColumn("has_o", df.Col("text").Contains(expr.Lit("o")))
	require.NoError(t, result.Err())
	defer result.Release()

	assert.Equal(t, int64(4), result.NumRows())
	assert.Equal(t, int64(2), result.NumCols())

	record := result.Record()

	// Find the has_o column
	hasOIdx := -1
	for i, colName := range result.ColumnNames() {
		if colName == "has_o" {
			hasOIdx = i
			break
		}
	}
	require.NotEqual(t, -1, hasOIdx, "has_o column not found")

	hasOCol := record.Column(hasOIdx).(*array.Boolean)

	// Check results
	assert.True(t, hasOCol.Value(0)) // "Hello World" contains "o"
	assert.True(t, hasOCol.Value(1)) // "Go programming" contains "o"
	assert.True(t, hasOCol.Value(2)) // "Apache Arrow" contains "o"
	assert.True(t, hasOCol.Value(3)) // "Data processing" contains "o"
}

func TestStringOperations_CaseSensitive(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with mixed case data
	texts := []string{
		"Hello World",
		"HELLO WORLD",
		"hello world",
		"Hi There",
	}
	df := createStringDataFrame(pool, "text", texts)
	defer df.Release()

	// Test case-sensitive contains
	result := df.Filter(df.Col("text").Contains(expr.Lit("Hello")))
	require.NoError(t, result.Err())
	defer result.Release()

	// Should only match exact case "Hello World"
	assert.Equal(t, int64(1), result.NumRows())

	record := result.Record()
	textCol := record.Column(0).(*array.String)
	assert.Equal(t, "Hello World", textCol.Value(0))
}

func TestStringOperations_EmptyStrings(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with empty strings
	texts := []string{
		"",
		"Hello",
		"",
		"World",
	}
	df := createStringDataFrame(pool, "text", texts)
	defer df.Release()

	// Test contains with empty string (should match all)
	result := df.Filter(df.Col("text").Contains(expr.Lit("")))
	require.NoError(t, result.Err())
	defer result.Release()

	// Empty string is contained in all strings
	assert.Equal(t, int64(4), result.NumRows())
}

func TestStringOperations_NullValues(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with null values
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("Hello")
	builder.AppendNull()
	builder.Append("World")
	builder.AppendNull()

	strArray := builder.NewArray()
	fields := []arrow.Field{{Name: "text", Type: arrow.BinaryTypes.String}}
	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 4)

	strArray.Release()
	df := NewDataFrame(record)
	defer df.Release()

	// Create contains expression result
	containsResult := df.WithColumn("contains_l", df.Col("text").Contains(expr.Lit("l")))
	require.NoError(t, containsResult.Err())
	defer containsResult.Release()

	record2 := containsResult.Record()

	// Find contains_l column
	containsIdx := -1
	for i, colName := range containsResult.ColumnNames() {
		if colName == "contains_l" {
			containsIdx = i
			break
		}
	}

	containsCol := record2.Column(containsIdx).(*array.Boolean)

	// Check results: null inputs should produce null outputs
	assert.True(t, containsCol.Value(0))  // "Hello" contains "l" -> true
	assert.True(t, containsCol.IsNull(1)) // null -> null
	assert.True(t, containsCol.Value(2))  // "World" contains "l" -> true
	assert.True(t, containsCol.IsNull(3)) // null -> null
}

func TestStringOperations_ErrorCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with integer data (not strings)
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	for i := int64(1); i <= 4; i++ {
		builder.Append(i)
	}

	intArray := builder.NewArray()
	fields := []arrow.Field{{Name: "numbers", Type: arrow.PrimitiveTypes.Int64}}
	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, []arrow.Array{intArray}, 4)

	intArray.Release()
	df := NewDataFrame(record)
	defer df.Release()

	// Try to use string operations on integer column (should fail)
	result := df.Filter(df.Col("numbers").Contains(expr.Lit("1")))
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "requires string operands")
}

func TestStringOperations_ChainedExpressions(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrame with email-like data
	emails := []string{
		"user@example.com",
		"admin@test.org",
		"support@example.com",
		"info@company.net",
	}
	df := createStringDataFrame(pool, "email", emails)
	defer df.Release()

	// Filter emails that contain "@example" AND end with ".com"
	// First filter by contains, then by ends with
	result := df.Filter(
		df.Col("email").Contains(expr.Lit("@example")),
	).Filter(
		df.Col("email").EndsWith(expr.Lit(".com")),
	)
	require.NoError(t, result.Err())
	defer result.Release()

	// Should match 2 example.com emails
	assert.Equal(t, int64(2), result.NumRows())

	record := result.Record()
	emailCol := record.Column(0).(*array.String)
	assert.Equal(t, "user@example.com", emailCol.Value(0))
	assert.Equal(t, "support@example.com", emailCol.Value(1))
}

func TestStringOperations_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	pool := memory.NewGoAllocator()

	// Create large dataset
	const size = 10000
	texts := make([]string, size)
	for i := 0; i < size; i++ {
		if i%3 == 0 {
			texts[i] = "prefix_match_" + fmt.Sprintf("%d", i)
		} else if i%3 == 1 {
			texts[i] = fmt.Sprintf("%d", i) + "_suffix_match"
		} else {
			texts[i] = "no_match_" + fmt.Sprintf("%d", i)
		}
	}

	df := createStringDataFrame(pool, "text", texts)
	defer df.Release()

	// Test Contains performance
	result := df.Filter(df.Col("text").Contains(expr.Lit("prefix")))
	require.NoError(t, result.Err())
	defer result.Release()

	// Should find approximately size/3 matches
	expectedMatches := int64(size / 3)
	assert.InDelta(t, expectedMatches, result.NumRows(), 100) // Allow some variance
}
