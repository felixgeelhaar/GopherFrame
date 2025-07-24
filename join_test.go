package gopherframe

import (
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a test DataFrame
func createTestDataFrame(pool memory.Allocator, data map[string]interface{}) *DataFrame {
	var fields []arrow.Field
	var arrays []arrow.Array

	for colName, values := range data {
		switch v := values.(type) {
		case []int64:
			builder := array.NewInt64Builder(pool)
			for _, val := range v {
				builder.Append(val)
			}
			arrays = append(arrays, builder.NewArray())
			fields = append(fields, arrow.Field{Name: colName, Type: arrow.PrimitiveTypes.Int64})
			builder.Release()

		case []string:
			builder := array.NewStringBuilder(pool)
			for _, val := range v {
				builder.Append(val)
			}
			arrays = append(arrays, builder.NewArray())
			fields = append(fields, arrow.Field{Name: colName, Type: arrow.BinaryTypes.String})
			builder.Release()

		case []float64:
			builder := array.NewFloat64Builder(pool)
			for _, val := range v {
				builder.Append(val)
			}
			arrays = append(arrays, builder.NewArray())
			fields = append(fields, arrow.Field{Name: colName, Type: arrow.PrimitiveTypes.Float64})
			builder.Release()
		}
	}

	schema := arrow.NewSchema(fields, nil)
	var numRows int64
	if len(arrays) > 0 {
		numRows = int64(arrays[0].Len())
	}
	record := array.NewRecord(schema, arrays, numRows)

	// Release arrays as they're now owned by the record
	for _, arr := range arrays {
		arr.Release()
	}

	return NewDataFrame(record)
}

func TestInnerJoin_BasicFunctionality(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame: users
	leftData := map[string]interface{}{
		"user_id": []int64{1, 2, 3, 4},
		"name":    []string{"Alice", "Bob", "Charlie", "David"},
		"age":     []int64{25, 30, 35, 40},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Create right DataFrame: orders
	rightData := map[string]interface{}{
		"user_id": []int64{1, 2, 2, 5},
		"product": []string{"Book", "Pen", "Notebook", "Phone"},
		"amount":  []float64{15.99, 2.99, 12.50, 599.99},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Perform inner join
	result := leftDF.InnerJoin(rightDF, "user_id", "user_id")
	require.NoError(t, result.Err())
	defer result.Release()

	// Verify the result
	assert.Equal(t, int64(3), result.NumRows()) // Only users 1 and 2 have orders
	assert.Equal(t, int64(5), result.NumCols()) // user_id, name, age, product, amount

	// Check column names (order may vary based on schema processing)
	actualCols := result.ColumnNames()
	assert.Contains(t, actualCols, "user_id")
	assert.Contains(t, actualCols, "name")
	assert.Contains(t, actualCols, "age")
	assert.Contains(t, actualCols, "product")
	assert.Contains(t, actualCols, "amount")

	// Verify specific values by getting correct column indices
	record := result.Record()

	// Find column indices
	userIdIdx := -1
	nameIdx := -1
	for i, colName := range actualCols {
		if colName == "user_id" {
			userIdIdx = i
		} else if colName == "name" {
			nameIdx = i
		}
	}

	// Check user_id column
	userIds := record.Column(userIdIdx).(*array.Int64)
	assert.Equal(t, int64(1), userIds.Value(0))
	assert.Equal(t, int64(2), userIds.Value(1))
	assert.Equal(t, int64(2), userIds.Value(2))

	// Check names
	names := record.Column(nameIdx).(*array.String)
	assert.Equal(t, "Alice", names.Value(0))
	assert.Equal(t, "Bob", names.Value(1))
	assert.Equal(t, "Bob", names.Value(2))
}

func TestLeftJoin_BasicFunctionality(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame: users
	leftData := map[string]interface{}{
		"user_id": []int64{1, 2, 3, 4},
		"name":    []string{"Alice", "Bob", "Charlie", "David"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Create right DataFrame: orders (only users 1 and 2 have orders)
	rightData := map[string]interface{}{
		"user_id": []int64{1, 2, 2},
		"product": []string{"Book", "Pen", "Notebook"},
		"amount":  []float64{15.99, 2.99, 12.50},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Perform left join
	result := leftDF.LeftJoin(rightDF, "user_id", "user_id")
	require.NoError(t, result.Err())
	defer result.Release()

	// Verify the result
	assert.Equal(t, int64(5), result.NumRows()) // All left rows + extra match for user 2
	assert.Equal(t, int64(4), result.NumCols()) // user_id, name, product, amount

	// Check that all users are present
	record := result.Record()
	userIds := record.Column(0).(*array.Int64)

	// Should have: 1, 2, 2, 3, 4 (user 2 appears twice due to multiple orders)
	expectedUserIds := []int64{1, 2, 2, 3, 4}
	for i, expected := range expectedUserIds {
		assert.Equal(t, expected, userIds.Value(i))
	}

	// Check that unmatched rows have null values
	// Find the product column index
	productIdx := -1
	actualCols := result.ColumnNames()
	for i, colName := range actualCols {
		if colName == "product" {
			productIdx = i
			break
		}
	}

	products := record.Column(productIdx).(*array.String)
	assert.False(t, products.IsNull(0)) // Alice has product
	assert.False(t, products.IsNull(1)) // Bob has product
	assert.False(t, products.IsNull(2)) // Bob has second product
	assert.True(t, products.IsNull(3))  // Charlie has no product (null)
	assert.True(t, products.IsNull(4))  // David has no product (null)
}

func TestJoin_ErrorCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	leftData := map[string]interface{}{
		"id":   []int64{1, 2, 3},
		"name": []string{"A", "B", "C"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"user_id": []int64{1, 2, 4},
		"value":   []string{"X", "Y", "Z"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Test nil DataFrame
	result := leftDF.InnerJoin(nil, "id", "user_id")
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "cannot be nil")

	// Test missing left key
	result = leftDF.InnerJoin(rightDF, "missing_key", "user_id")
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "left join key column not found")

	// Test missing right key
	result = leftDF.InnerJoin(rightDF, "id", "missing_key")
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "right join key column not found")
}

func TestJoin_DifferentDataTypes(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Left DataFrame with string keys
	leftData := map[string]interface{}{
		"code": []string{"A", "B", "C"},
		"desc": []string{"Alpha", "Beta", "Gamma"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Right DataFrame with string keys
	rightData := map[string]interface{}{
		"code":  []string{"A", "B", "D"},
		"value": []int64{100, 200, 300},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Perform inner join with string keys
	result := leftDF.InnerJoin(rightDF, "code", "code")
	require.NoError(t, result.Err())
	defer result.Release()

	// Should have 2 rows (A and B match)
	assert.Equal(t, int64(2), result.NumRows())

	record := result.Record()

	// Find the code column index
	codeIdx := -1
	actualCols := result.ColumnNames()
	for i, colName := range actualCols {
		if colName == "code" {
			codeIdx = i
			break
		}
	}

	codes := record.Column(codeIdx).(*array.String)
	assert.Equal(t, "A", codes.Value(0))
	assert.Equal(t, "B", codes.Value(1))
}

func TestJoin_EmptyDataFrames(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create empty left DataFrame
	leftData := map[string]interface{}{
		"id":   []int64{},
		"name": []string{},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Create non-empty right DataFrame
	rightData := map[string]interface{}{
		"id":    []int64{1, 2},
		"value": []string{"X", "Y"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Inner join with empty left should return empty result
	result := leftDF.InnerJoin(rightDF, "id", "id")
	require.NoError(t, result.Err())
	defer result.Release()

	assert.Equal(t, int64(0), result.NumRows())

	// Left join with empty left should return empty result
	result2 := leftDF.LeftJoin(rightDF, "id", "id")
	require.NoError(t, result2.Err())
	defer result2.Release()

	assert.Equal(t, int64(0), result2.NumRows())
}

func TestJoin_DuplicateColumnNames(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Both DataFrames have a "name" column
	leftData := map[string]interface{}{
		"id":   []int64{1, 2},
		"name": []string{"Alice", "Bob"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"id":   []int64{1, 2},
		"name": []string{"Smith", "Jones"}, // This should become "right_name"
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.InnerJoin(rightDF, "id", "id")
	require.NoError(t, result.Err())
	defer result.Release()

	// Check that duplicate column was renamed
	actualCols := result.ColumnNames()
	assert.Contains(t, actualCols, "id")
	assert.Contains(t, actualCols, "name")
	assert.Contains(t, actualCols, "right_name")

	// Verify values
	record := result.Record()

	// Find column indices
	nameIdx := -1
	rightNameIdx := -1
	for i, colName := range actualCols {
		if colName == "name" {
			nameIdx = i
		} else if colName == "right_name" {
			rightNameIdx = i
		}
	}

	leftNames := record.Column(nameIdx).(*array.String)
	rightNames := record.Column(rightNameIdx).(*array.String)

	assert.Equal(t, "Alice", leftNames.Value(0))
	assert.Equal(t, "Smith", rightNames.Value(0))
}

func TestJoin_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	pool := memory.NewGoAllocator()

	// Create larger DataFrames for performance testing
	const size = 10000

	leftIds := make([]int64, size)
	leftValues := make([]string, size)
	for i := 0; i < size; i++ {
		leftIds[i] = int64(i)
		leftValues[i] = fmt.Sprintf("left_%d", i)
	}

	rightIds := make([]int64, size/2)
	rightValues := make([]string, size/2)
	for i := 0; i < size/2; i++ {
		rightIds[i] = int64(i * 2) // Every other ID to test join selectivity
		rightValues[i] = fmt.Sprintf("right_%d", i*2)
	}

	leftData := map[string]interface{}{
		"id":    leftIds,
		"value": leftValues,
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"id":    rightIds,
		"value": rightValues,
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Perform inner join
	result := leftDF.InnerJoin(rightDF, "id", "id")
	require.NoError(t, result.Err())
	defer result.Release()

	// Should have size/2 rows (every other left row has a match)
	assert.Equal(t, int64(size/2), result.NumRows())
}

func TestJoin_NullValues(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create DataFrames with null values in join keys
	leftBuilder := array.NewInt64Builder(pool)
	leftBuilder.Append(1)
	leftBuilder.AppendNull() // null value
	leftBuilder.Append(3)
	leftIds := leftBuilder.NewArray()
	leftBuilder.Release()

	leftNames := []string{"Alice", "Bob", "Charlie"}
	leftNameBuilder := array.NewStringBuilder(pool)
	for _, name := range leftNames {
		leftNameBuilder.Append(name)
	}
	leftNameArray := leftNameBuilder.NewArray()
	leftNameBuilder.Release()

	leftSchema := arrow.NewSchema([]arrow.Field{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		{Name: "name", Type: arrow.BinaryTypes.String},
	}, nil)
	leftRecord := array.NewRecord(leftSchema, []arrow.Array{leftIds, leftNameArray}, 3)
	leftDF := NewDataFrame(leftRecord)
	defer leftDF.Release()

	// Right DataFrame
	rightData := map[string]interface{}{
		"id":    []int64{1, 2, 3},
		"value": []string{"X", "Y", "Z"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Inner join should skip null values
	result := leftDF.InnerJoin(rightDF, "id", "id")
	require.NoError(t, result.Err())
	defer result.Release()

	// Should have 2 rows (1 and 3 match, null is skipped)
	assert.Equal(t, int64(2), result.NumRows())

	// Left join should include null value row with nulls for right side
	resultLeft := leftDF.LeftJoin(rightDF, "id", "id")
	require.NoError(t, resultLeft.Err())
	defer resultLeft.Release()

	// Should have 3 rows (all left rows)
	assert.Equal(t, int64(3), resultLeft.NumRows())
}
