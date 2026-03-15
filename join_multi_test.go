package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInnerJoinMulti(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Left DataFrame: employees with department and role
	leftData := map[string]interface{}{
		"dept": []string{"Engineering", "Engineering", "Sales", "Sales", "HR"},
		"name": []string{"Alice", "Bob", "Charlie", "David", "Eve"},
		"role": []string{"Senior", "Junior", "Senior", "Junior", "Senior"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Right DataFrame: budgets by department and role
	rightData := map[string]interface{}{
		"budget":     []float64{200000, 100000, 150000},
		"department": []string{"Engineering", "Engineering", "Sales"},
		"position":   []string{"Senior", "Junior", "Senior"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.InnerJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department", "position"},
	)
	require.NoError(t, result.Err())
	defer result.Release()

	// Should match: Alice(Eng,Senior), Bob(Eng,Junior), Charlie(Sales,Senior)
	assert.Equal(t, int64(3), result.NumRows())

	// Columns: dept, name, role, budget (department and position are right keys, skipped)
	actualCols := result.ColumnNames()
	assert.Contains(t, actualCols, "dept")
	assert.Contains(t, actualCols, "name")
	assert.Contains(t, actualCols, "role")
	assert.Contains(t, actualCols, "budget")
	// Right key columns should be excluded
	assert.NotContains(t, actualCols, "department")
	assert.NotContains(t, actualCols, "position")

	// Verify values
	record := result.Record()
	nameIdx := -1
	budgetIdx := -1
	for i, col := range actualCols {
		switch col {
		case "name":
			nameIdx = i
		case "budget":
			budgetIdx = i
		}
	}

	names := record.Column(nameIdx).(*array.String)
	budgets := record.Column(budgetIdx).(*array.Float64)

	// Collect results into a map for order-independent verification
	resultMap := make(map[string]float64)
	for i := 0; i < int(result.NumRows()); i++ {
		resultMap[names.Value(i)] = budgets.Value(i)
	}

	assert.Equal(t, 200000.0, resultMap["Alice"])
	assert.Equal(t, 100000.0, resultMap["Bob"])
	assert.Equal(t, 150000.0, resultMap["Charlie"])
}

func TestLeftJoinMulti(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Left DataFrame
	leftData := map[string]interface{}{
		"dept": []string{"Engineering", "Sales", "HR"},
		"name": []string{"Alice", "Bob", "Charlie"},
		"role": []string{"Senior", "Senior", "Senior"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Right DataFrame: only Engineering+Senior has a budget
	rightData := map[string]interface{}{
		"budget":     []float64{200000},
		"department": []string{"Engineering"},
		"position":   []string{"Senior"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.LeftJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department", "position"},
	)
	require.NoError(t, result.Err())
	defer result.Release()

	// All 3 left rows should be present
	assert.Equal(t, int64(3), result.NumRows())

	record := result.Record()
	actualCols := result.ColumnNames()

	nameIdx := -1
	budgetIdx := -1
	for i, col := range actualCols {
		switch col {
		case "name":
			nameIdx = i
		case "budget":
			budgetIdx = i
		}
	}

	names := record.Column(nameIdx).(*array.String)
	budgetCol := record.Column(budgetIdx).(*array.Float64)

	// Find Alice's row (the one with a match)
	for i := 0; i < int(result.NumRows()); i++ {
		if names.Value(i) == "Alice" {
			assert.False(t, budgetCol.IsNull(i), "Alice should have a budget")
			assert.Equal(t, 200000.0, budgetCol.Value(i))
		} else {
			assert.True(t, budgetCol.IsNull(i), "%s should have null budget", names.Value(i))
		}
	}
}

func TestRightJoinMulti(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Left DataFrame: only one employee
	leftData := map[string]interface{}{
		"dept": []string{"Engineering"},
		"name": []string{"Alice"},
		"role": []string{"Senior"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	// Right DataFrame: two budgets
	rightData := map[string]interface{}{
		"budget":     []float64{200000, 100000},
		"department": []string{"Engineering", "Sales"},
		"position":   []string{"Senior", "Senior"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.RightJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department", "position"},
	)
	require.NoError(t, result.Err())
	defer result.Release()

	// All 2 right rows should be present
	assert.Equal(t, int64(2), result.NumRows())

	record := result.Record()
	actualCols := result.ColumnNames()

	nameIdx := -1
	budgetIdx := -1
	for i, col := range actualCols {
		switch col {
		case "name":
			nameIdx = i
		case "budget":
			budgetIdx = i
		}
	}

	names := record.Column(nameIdx).(*array.String)
	budgets := record.Column(budgetIdx).(*array.Float64)

	// One row should have Alice, one should have null name
	foundAlice := false
	foundNull := false
	for i := 0; i < int(result.NumRows()); i++ {
		if !names.IsNull(i) && names.Value(i) == "Alice" {
			foundAlice = true
			assert.Equal(t, 200000.0, budgets.Value(i))
		}
		if names.IsNull(i) {
			foundNull = true
			assert.Equal(t, 100000.0, budgets.Value(i))
		}
	}
	assert.True(t, foundAlice, "should contain Alice")
	assert.True(t, foundNull, "should contain null name for unmatched right row")
}

func TestFullOuterJoinMulti(t *testing.T) {
	pool := memory.NewGoAllocator()

	leftData := map[string]interface{}{
		"dept": []string{"Engineering", "HR"},
		"name": []string{"Alice", "Bob"},
		"role": []string{"Senior", "Senior"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"budget":     []float64{200000, 100000},
		"department": []string{"Engineering", "Sales"},
		"position":   []string{"Senior", "Senior"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.FullOuterJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department", "position"},
	)
	require.NoError(t, result.Err())
	defer result.Release()

	// Alice+Eng matches, Bob+HR has no match, Sales+Senior has no match = 3 rows
	assert.Equal(t, int64(3), result.NumRows())
}

func TestJoinMultiKeyMismatch(t *testing.T) {
	pool := memory.NewGoAllocator()

	leftData := map[string]interface{}{
		"dept": []string{"Engineering"},
		"name": []string{"Alice"},
		"role": []string{"Senior"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"department": []string{"Engineering"},
		"position":   []string{"Senior"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Different number of keys
	result := leftDF.InnerJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department"},
	)
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "same length")
}

func TestJoinMultiKeyNotFound(t *testing.T) {
	pool := memory.NewGoAllocator()

	leftData := map[string]interface{}{
		"dept": []string{"Engineering"},
		"name": []string{"Alice"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"department": []string{"Engineering"},
		"budget":     []float64{200000},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Left key not found
	result := leftDF.InnerJoinMulti(
		rightDF,
		[]string{"dept", "missing_col"},
		[]string{"department", "budget"},
	)
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "left join key column not found: missing_col")

	// Right key not found
	result = leftDF.InnerJoinMulti(
		rightDF,
		[]string{"dept", "name"},
		[]string{"department", "missing_col"},
	)
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "right join key column not found: missing_col")
}

func TestJoinMultiEmptyKeys(t *testing.T) {
	pool := memory.NewGoAllocator()

	leftData := map[string]interface{}{
		"id": []int64{1},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"id": []int64{1},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.InnerJoinMulti(rightDF, []string{}, []string{})
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "join keys cannot be empty")
}

func TestJoinMultiNilDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()

	leftData := map[string]interface{}{
		"id": []int64{1},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	result := leftDF.InnerJoinMulti(nil, []string{"id"}, []string{"id"})
	assert.Error(t, result.Err())
	assert.Contains(t, result.Err().Error(), "cannot be nil")
}

func TestJoinMultiWithNullKeys(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame with a null in one key column
	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.Append("Engineering")
	deptBuilder.AppendNull() // null department
	deptBuilder.Append("Sales")
	deptArr := deptBuilder.NewArray()
	deptBuilder.Release()

	roleBuilder := array.NewStringBuilder(pool)
	roleBuilder.Append("Senior")
	roleBuilder.Append("Senior")
	roleBuilder.Append("Senior")
	roleArr := roleBuilder.NewArray()
	roleBuilder.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.Append("Alice")
	nameBuilder.Append("Bob")
	nameBuilder.Append("Charlie")
	nameArr := nameBuilder.NewArray()
	nameBuilder.Release()

	leftSchema := arrow.NewSchema([]arrow.Field{
		{Name: "dept", Type: arrow.BinaryTypes.String},
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "role", Type: arrow.BinaryTypes.String},
	}, nil)
	leftRecord := array.NewRecord(leftSchema, []arrow.Array{deptArr, nameArr, roleArr}, 3)
	leftDF := NewDataFrame(leftRecord)
	defer leftDF.Release()

	deptArr.Release()
	roleArr.Release()
	nameArr.Release()

	// Right DataFrame
	rightData := map[string]interface{}{
		"budget":     []float64{200000, 150000},
		"department": []string{"Engineering", "Sales"},
		"position":   []string{"Senior", "Senior"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	// Inner join: row with null dept should be excluded
	result := leftDF.InnerJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department", "position"},
	)
	require.NoError(t, result.Err())
	defer result.Release()

	// Only Alice(Eng,Senior) and Charlie(Sales,Senior) should match; Bob has null dept
	assert.Equal(t, int64(2), result.NumRows())

	// Left join: all 3 left rows, but Bob gets nulls for right side
	resultLeft := leftDF.LeftJoinMulti(
		rightDF,
		[]string{"dept", "role"},
		[]string{"department", "position"},
	)
	require.NoError(t, resultLeft.Err())
	defer resultLeft.Release()

	assert.Equal(t, int64(3), resultLeft.NumRows())
}

func TestJoinMultiDuplicateColumnNames(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Both DataFrames have a "value" column
	leftData := map[string]interface{}{
		"key1":  []string{"a", "b"},
		"key2":  []int64{1, 2},
		"value": []string{"left_a1", "left_b2"},
	}
	leftDF := createTestDataFrame(pool, leftData)
	defer leftDF.Release()

	rightData := map[string]interface{}{
		"rkey1": []string{"a", "b"},
		"rkey2": []int64{1, 2},
		"value": []string{"right_a1", "right_b2"},
	}
	rightDF := createTestDataFrame(pool, rightData)
	defer rightDF.Release()

	result := leftDF.InnerJoinMulti(
		rightDF,
		[]string{"key1", "key2"},
		[]string{"rkey1", "rkey2"},
	)
	require.NoError(t, result.Err())
	defer result.Release()

	assert.Equal(t, int64(2), result.NumRows())

	actualCols := result.ColumnNames()
	assert.Contains(t, actualCols, "value")
	assert.Contains(t, actualCols, "right_value")
	// Right keys should be excluded
	assert.NotContains(t, actualCols, "rkey1")
	assert.NotContains(t, actualCols, "rkey2")
}
