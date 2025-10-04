package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDataFrame_InnerJoin tests the InnerJoin method
func TestDataFrame_InnerJoin(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame: users
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "age", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	userIdBuilder := array.NewInt64Builder(pool)
	userIdBuilder.AppendValues([]int64{1, 2, 3, 4}, nil)
	userIdArray := userIdBuilder.NewArray()
	defer userIdArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Carol", "Dave"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	ageBuilder := array.NewInt64Builder(pool)
	ageBuilder.AppendValues([]int64{25, 30, 28, 35}, nil)
	ageArray := ageBuilder.NewArray()
	defer ageArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{userIdArray, nameArray, ageArray}, 4)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame: orders
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "order_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "customer_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	orderIdBuilder := array.NewInt64Builder(pool)
	orderIdBuilder.AppendValues([]int64{101, 102, 103, 104}, nil)
	orderIdArray := orderIdBuilder.NewArray()
	defer orderIdArray.Release()

	customerIdBuilder := array.NewInt64Builder(pool)
	customerIdBuilder.AppendValues([]int64{1, 2, 2, 5}, nil) // user 3 and 4 have no orders, user 5 doesn't exist
	customerIdArray := customerIdBuilder.NewArray()
	defer customerIdArray.Release()

	amountBuilder := array.NewFloat64Builder(pool)
	amountBuilder.AppendValues([]float64{100.50, 200.75, 150.25, 300.00}, nil)
	amountArray := amountBuilder.NewArray()
	defer amountArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{orderIdArray, customerIdArray, amountArray}, 4)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform inner join
	result, err := leftDf.InnerJoin(rightDf, "user_id", "customer_id")
	require.NoError(t, err)
	defer result.Release()

	// Verify result
	assert.Equal(t, int64(3), result.NumRows(), "Inner join should return 3 rows (users 1, 2, 2)")
	assert.Equal(t, int64(5), result.NumCols(), "Result should have 5 columns (3 from left + 2 from right, excluding duplicate key)")

	// Verify column names
	assert.True(t, result.HasColumn("user_id"))
	assert.True(t, result.HasColumn("name"))
	assert.True(t, result.HasColumn("age"))
	assert.True(t, result.HasColumn("order_id"))
	assert.True(t, result.HasColumn("amount"))
	assert.False(t, result.HasColumn("customer_id"), "customer_id should be excluded (duplicate key)")

	// Verify data: first row should be user 1
	userIdSeries, err := result.Column("user_id")
	require.NoError(t, err)
	userIdCol := userIdSeries.Array().(*array.Int64)
	assert.Equal(t, int64(1), userIdCol.Value(0))

	nameSeries, err := result.Column("name")
	require.NoError(t, err)
	nameCol := nameSeries.Array().(*array.String)
	assert.Equal(t, "Alice", nameCol.Value(0))

	// Verify second and third rows are user 2 (two orders)
	assert.Equal(t, int64(2), userIdCol.Value(1))
	assert.Equal(t, int64(2), userIdCol.Value(2))
	assert.Equal(t, "Bob", nameCol.Value(1))
	assert.Equal(t, "Bob", nameCol.Value(2))
}

// TestDataFrame_LeftJoin tests the LeftJoin method
func TestDataFrame_LeftJoin(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2, 3}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueBuilder := array.NewStringBuilder(pool)
	valueBuilder.AppendValues([]string{"A", "B", "C"}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{idArray, valueArray}, 3)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame (only has matches for 1 and 2)
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "ref_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	refIdBuilder := array.NewInt64Builder(pool)
	refIdBuilder.AppendValues([]int64{1, 2}, nil)
	refIdArray := refIdBuilder.NewArray()
	defer refIdArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{95.5, 87.3}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{refIdArray, scoreArray}, 2)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform left join
	result, err := leftDf.LeftJoin(rightDf, "id", "ref_id")
	require.NoError(t, err)
	defer result.Release()

	// Verify result
	assert.Equal(t, int64(3), result.NumRows(), "Left join should preserve all left rows")
	assert.Equal(t, int64(3), result.NumCols(), "Result should have 3 columns")

	// Verify that row 3 (id=3) has null score
	scoreSeries, err := result.Column("score")
	require.NoError(t, err)
	scoreCol := scoreSeries.Array().(*array.Float64)

	assert.False(t, scoreCol.IsNull(0), "First row should have score")
	assert.False(t, scoreCol.IsNull(1), "Second row should have score")
	assert.True(t, scoreCol.IsNull(2), "Third row should have null score (no match)")
}

// TestDataFrame_JoinWithNulls tests join behavior with null values
func TestDataFrame_JoinWithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame with null key
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "key", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
			{Name: "data", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	keyBuilder := array.NewInt64Builder(pool)
	keyBuilder.AppendValues([]int64{1, 2, 3}, []bool{true, false, true}) // second value is null
	keyArray := keyBuilder.NewArray()
	defer keyArray.Release()

	dataBuilder := array.NewStringBuilder(pool)
	dataBuilder.AppendValues([]string{"A", "B", "C"}, nil)
	dataArray := dataBuilder.NewArray()
	defer dataArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{keyArray, dataArray}, 3)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "key", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	rightKeyBuilder := array.NewInt64Builder(pool)
	rightKeyBuilder.AppendValues([]int64{1, 3}, nil)
	rightKeyArray := rightKeyBuilder.NewArray()
	defer rightKeyArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{10.5, 30.5}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{rightKeyArray, valueArray}, 2)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Inner join should skip null keys
	innerResult, err := leftDf.InnerJoin(rightDf, "key", "key")
	require.NoError(t, err)
	defer innerResult.Release()
	assert.Equal(t, int64(2), innerResult.NumRows(), "Inner join should skip null keys")

	// Left join should include null key row with null right values
	leftResult, err := leftDf.LeftJoin(rightDf, "key", "key")
	require.NoError(t, err)
	defer leftResult.Release()
	assert.Equal(t, int64(3), leftResult.NumRows(), "Left join should include null key rows")
}

// TestDataFrame_JoinErrors tests error cases for join operations
func TestDataFrame_JoinErrors(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueBuilder := array.NewStringBuilder(pool)
	valueBuilder.AppendValues([]string{"A", "B"}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, valueArray}, 2)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: nil other DataFrame
	_, err := df.InnerJoin(nil, "id", "id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be nil")

	// Test: left key not found
	otherDf := NewDataFrame(record)
	defer otherDf.Release()
	_, err = df.InnerJoin(otherDf, "nonexistent", "id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "left join key column not found")

	// Test: right key not found
	_, err = df.InnerJoin(otherDf, "id", "nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "right join key column not found")
}

// TestDataFrame_JoinWithDifferentTypes tests join with different data types
func TestDataFrame_JoinWithDifferentTypes(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("StringKeys", func(t *testing.T) {
		// Left DataFrame
		leftSchema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "code", Type: arrow.BinaryTypes.String},
				{Name: "name", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		codeBuilder := array.NewStringBuilder(pool)
		codeBuilder.AppendValues([]string{"US", "UK", "FR"}, nil)
		codeArray := codeBuilder.NewArray()
		defer codeArray.Release()

		nameBuilder := array.NewStringBuilder(pool)
		nameBuilder.AppendValues([]string{"United States", "United Kingdom", "France"}, nil)
		nameArray := nameBuilder.NewArray()
		defer nameArray.Release()

		leftRecord := array.NewRecord(leftSchema, []arrow.Array{codeArray, nameArray}, 3)
		defer leftRecord.Release()

		leftDf := NewDataFrame(leftRecord)
		defer leftDf.Release()

		// Right DataFrame
		rightSchema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "country_code", Type: arrow.BinaryTypes.String},
				{Name: "population", Type: arrow.PrimitiveTypes.Int64},
			},
			nil,
		)

		countryCodeBuilder := array.NewStringBuilder(pool)
		countryCodeBuilder.AppendValues([]string{"US", "UK"}, nil)
		countryCodeArray := countryCodeBuilder.NewArray()
		defer countryCodeArray.Release()

		popBuilder := array.NewInt64Builder(pool)
		popBuilder.AppendValues([]int64{331000000, 67000000}, nil)
		popArray := popBuilder.NewArray()
		defer popArray.Release()

		rightRecord := array.NewRecord(rightSchema, []arrow.Array{countryCodeArray, popArray}, 2)
		defer rightRecord.Release()

		rightDf := NewDataFrame(rightRecord)
		defer rightDf.Release()

		// Inner join on string keys
		result, err := leftDf.InnerJoin(rightDf, "code", "country_code")
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, int64(2), result.NumRows())
		assert.True(t, result.HasColumn("code"))
		assert.True(t, result.HasColumn("name"))
		assert.True(t, result.HasColumn("population"))
	})

	t.Run("Float64Keys", func(t *testing.T) {
		// Left DataFrame
		leftSchema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "price", Type: arrow.PrimitiveTypes.Float64},
				{Name: "product", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		priceBuilder := array.NewFloat64Builder(pool)
		priceBuilder.AppendValues([]float64{9.99, 19.99, 29.99}, nil)
		priceArray := priceBuilder.NewArray()
		defer priceArray.Release()

		productBuilder := array.NewStringBuilder(pool)
		productBuilder.AppendValues([]string{"Book", "DVD", "Game"}, nil)
		productArray := productBuilder.NewArray()
		defer productArray.Release()

		leftRecord := array.NewRecord(leftSchema, []arrow.Array{priceArray, productArray}, 3)
		defer leftRecord.Release()

		leftDf := NewDataFrame(leftRecord)
		defer leftDf.Release()

		// Right DataFrame
		rightSchema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "cost", Type: arrow.PrimitiveTypes.Float64},
				{Name: "quantity", Type: arrow.PrimitiveTypes.Int64},
			},
			nil,
		)

		costBuilder := array.NewFloat64Builder(pool)
		costBuilder.AppendValues([]float64{9.99, 29.99}, nil)
		costArray := costBuilder.NewArray()
		defer costArray.Release()

		qtyBuilder := array.NewInt64Builder(pool)
		qtyBuilder.AppendValues([]int64{100, 50}, nil)
		qtyArray := qtyBuilder.NewArray()
		defer qtyArray.Release()

		rightRecord := array.NewRecord(rightSchema, []arrow.Array{costArray, qtyArray}, 2)
		defer rightRecord.Release()

		rightDf := NewDataFrame(rightRecord)
		defer rightDf.Release()

		// Inner join on float64 keys
		result, err := leftDf.InnerJoin(rightDf, "price", "cost")
		require.NoError(t, err)
		defer result.Release()

		assert.Equal(t, int64(2), result.NumRows())
	})
}

// TestDataFrame_JoinWithColumnNameConflicts tests handling of duplicate column names
func TestDataFrame_JoinWithColumnNameConflicts(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Both DataFrames have a "value" column
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	leftIdBuilder := array.NewInt64Builder(pool)
	leftIdBuilder.AppendValues([]int64{1, 2}, nil)
	leftIdArray := leftIdBuilder.NewArray()
	defer leftIdArray.Release()

	leftValueBuilder := array.NewStringBuilder(pool)
	leftValueBuilder.AppendValues([]string{"A", "B"}, nil)
	leftValueArray := leftValueBuilder.NewArray()
	defer leftValueArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{leftIdArray, leftValueArray}, 2)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Right DataFrame also has "value" column
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "ref_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64}, // Different type but same name
		},
		nil,
	)

	rightIdBuilder := array.NewInt64Builder(pool)
	rightIdBuilder.AppendValues([]int64{1, 2}, nil)
	rightIdArray := rightIdBuilder.NewArray()
	defer rightIdArray.Release()

	rightValueBuilder := array.NewFloat64Builder(pool)
	rightValueBuilder.AppendValues([]float64{10.5, 20.5}, nil)
	rightValueArray := rightValueBuilder.NewArray()
	defer rightValueArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{rightIdArray, rightValueArray}, 2)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform join
	result, err := leftDf.InnerJoin(rightDf, "id", "ref_id")
	require.NoError(t, err)
	defer result.Release()

	// Verify that conflicting column is prefixed
	assert.True(t, result.HasColumn("value"), "Left 'value' column should exist")
	assert.True(t, result.HasColumn("right_value"), "Right 'value' column should be prefixed")

	// Verify correct types
	valueSeries, err := result.Column("value")
	require.NoError(t, err)
	assert.Equal(t, arrow.BinaryTypes.String, valueSeries.Field().Type, "Left value should be string")

	rightValueSeries, err := result.Column("right_value")
	require.NoError(t, err)
	assert.Equal(t, arrow.PrimitiveTypes.Float64, rightValueSeries.Field().Type, "Right value should be float64")
}

// TestDataFrame_JoinOneToMany tests one-to-many join scenarios
func TestDataFrame_JoinOneToMany(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Left: One user
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	userIdBuilder := array.NewInt64Builder(pool)
	userIdBuilder.AppendValues([]int64{1}, nil)
	userIdArray := userIdBuilder.NewArray()
	defer userIdArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{userIdArray, nameArray}, 1)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Right: Multiple orders for same user
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "customer_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "order_id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	customerIdBuilder := array.NewInt64Builder(pool)
	customerIdBuilder.AppendValues([]int64{1, 1, 1}, nil) // Three orders for user 1
	customerIdArray := customerIdBuilder.NewArray()
	defer customerIdArray.Release()

	orderIdBuilder := array.NewInt64Builder(pool)
	orderIdBuilder.AppendValues([]int64{101, 102, 103}, nil)
	orderIdArray := orderIdBuilder.NewArray()
	defer orderIdArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{customerIdArray, orderIdArray}, 3)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Inner join should produce 3 rows
	result, err := leftDf.InnerJoin(rightDf, "user_id", "customer_id")
	require.NoError(t, err)
	defer result.Release()

	assert.Equal(t, int64(3), result.NumRows(), "One-to-many join should produce 3 rows")

	// Verify all rows have same user name
	nameSeries, err := result.Column("name")
	require.NoError(t, err)
	nameCol := nameSeries.Array().(*array.String)

	for i := 0; i < 3; i++ {
		assert.Equal(t, "Alice", nameCol.Value(i))
	}
}

// TestDataFrame_getColumnIndex tests the getColumnIndex helper
func TestDataFrame_getColumnIndex(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "first", Type: arrow.PrimitiveTypes.Int64},
			{Name: "second", Type: arrow.BinaryTypes.String},
			{Name: "third", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	firstBuilder := array.NewInt64Builder(pool)
	firstBuilder.AppendValues([]int64{1}, nil)
	firstArray := firstBuilder.NewArray()
	defer firstArray.Release()

	secondBuilder := array.NewStringBuilder(pool)
	secondBuilder.AppendValues([]string{"A"}, nil)
	secondArray := secondBuilder.NewArray()
	defer secondArray.Release()

	thirdBuilder := array.NewFloat64Builder(pool)
	thirdBuilder.AppendValues([]float64{1.5}, nil)
	thirdArray := thirdBuilder.NewArray()
	defer thirdArray.Release()

	record := array.NewRecord(schema, []arrow.Array{firstArray, secondArray, thirdArray}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test existing columns
	assert.Equal(t, 0, df.getColumnIndex("first"))
	assert.Equal(t, 1, df.getColumnIndex("second"))
	assert.Equal(t, 2, df.getColumnIndex("third"))

	// Test non-existent column
	assert.Equal(t, -1, df.getColumnIndex("nonexistent"))
}

// TestDataFrame_RightJoin tests the RightJoin method
func TestDataFrame_RightJoin(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame: users (some will not match)
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	userIdBuilder := array.NewInt64Builder(pool)
	userIdBuilder.AppendValues([]int64{1, 2, 3}, nil)
	userIdArray := userIdBuilder.NewArray()
	defer userIdArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Carol"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{userIdArray, nameArray}, 3)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame: orders (user_id 4 and 5 have no matching users)
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "order_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "customer_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	orderIdBuilder := array.NewInt64Builder(pool)
	orderIdBuilder.AppendValues([]int64{101, 102, 103, 104}, nil)
	orderIdArray := orderIdBuilder.NewArray()
	defer orderIdArray.Release()

	customerIdBuilder := array.NewInt64Builder(pool)
	customerIdBuilder.AppendValues([]int64{1, 2, 4, 5}, nil) // 4 and 5 don't exist in left
	customerIdArray := customerIdBuilder.NewArray()
	defer customerIdArray.Release()

	amountBuilder := array.NewFloat64Builder(pool)
	amountBuilder.AppendValues([]float64{100.0, 200.0, 300.0, 400.0}, nil)
	amountArray := amountBuilder.NewArray()
	defer amountArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{orderIdArray, customerIdArray, amountArray}, 4)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform right join
	result, err := leftDf.RightJoin(rightDf, "user_id", "customer_id")
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify all right rows are included (4 rows from right)
	assert.Equal(t, int64(4), result.NumRows())

	// Verify schema: user_id (from left key), name, order_id, amount
	// Note: customer_id is excluded as it's the right join key
	assert.Equal(t, int64(4), result.NumCols())
	assert.True(t, result.HasColumn("user_id"))
	assert.True(t, result.HasColumn("name"))
	assert.True(t, result.HasColumn("order_id"))
	assert.True(t, result.HasColumn("amount"))

	// Verify data: rows with customer_id 4 and 5 should have null user names
	nameSeries, err := result.Column("name")
	require.NoError(t, err)
	nameCol := nameSeries.Array().(*array.String)

	orderIdSeries, err := result.Column("order_id")
	require.NoError(t, err)
	orderIdCol := orderIdSeries.Array().(*array.Int64)

	// First two orders (101, 102) have matching users
	assert.False(t, nameCol.IsNull(0)) // order 101 -> user 1 (Alice)
	assert.False(t, nameCol.IsNull(1)) // order 102 -> user 2 (Bob)

	// Last two orders (103, 104) have no matching users
	assert.True(t, nameCol.IsNull(2)) // order 103 -> no user (null)
	assert.True(t, nameCol.IsNull(3)) // order 104 -> no user (null)

	// All order IDs should be present
	assert.Equal(t, int64(101), orderIdCol.Value(0))
	assert.Equal(t, int64(102), orderIdCol.Value(1))
	assert.Equal(t, int64(103), orderIdCol.Value(2))
	assert.Equal(t, int64(104), orderIdCol.Value(3))
}

// TestDataFrame_FullOuterJoin tests the FullOuterJoin method
func TestDataFrame_FullOuterJoin(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame: users (user 3 has no orders)
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	userIdBuilder := array.NewInt64Builder(pool)
	userIdBuilder.AppendValues([]int64{1, 2, 3}, nil)
	userIdArray := userIdBuilder.NewArray()
	defer userIdArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Carol"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{userIdArray, nameArray}, 3)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame: orders (order 103 has no matching user)
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "order_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "customer_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	orderIdBuilder := array.NewInt64Builder(pool)
	orderIdBuilder.AppendValues([]int64{101, 102, 103}, nil)
	orderIdArray := orderIdBuilder.NewArray()
	defer orderIdArray.Release()

	customerIdBuilder := array.NewInt64Builder(pool)
	customerIdBuilder.AppendValues([]int64{1, 2, 99}, nil) // 99 doesn't exist in left
	customerIdArray := customerIdBuilder.NewArray()
	defer customerIdArray.Release()

	amountBuilder := array.NewFloat64Builder(pool)
	amountBuilder.AppendValues([]float64{100.0, 200.0, 300.0}, nil)
	amountArray := amountBuilder.NewArray()
	defer amountArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{orderIdArray, customerIdArray, amountArray}, 3)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform full outer join
	result, err := leftDf.FullOuterJoin(rightDf, "user_id", "customer_id")
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify result includes:
	// - User 1 with order 101 (matched)
	// - User 2 with order 102 (matched)
	// - User 3 with no order (left only)
	// - Order 103 with no user (right only)
	// Total: 4 rows
	assert.Equal(t, int64(4), result.NumRows())

	// Verify schema
	assert.Equal(t, int64(4), result.NumCols())
	assert.True(t, result.HasColumn("user_id"))
	assert.True(t, result.HasColumn("name"))
	assert.True(t, result.HasColumn("order_id"))
	assert.True(t, result.HasColumn("amount"))

	// Verify data
	nameSeries, err := result.Column("name")
	require.NoError(t, err)
	nameCol := nameSeries.Array().(*array.String)

	orderIdSeries, err := result.Column("order_id")
	require.NoError(t, err)
	orderIdCol := orderIdSeries.Array().(*array.Int64)

	// Count nulls in each column
	nameNulls := 0
	orderNulls := 0
	for i := 0; i < int(result.NumRows()); i++ {
		if nameCol.IsNull(i) {
			nameNulls++
		}
		if orderIdCol.IsNull(i) {
			orderNulls++
		}
	}

	// Should have 1 null name (order 103 with no user)
	assert.Equal(t, 1, nameNulls, "Expected 1 null name (unmatched order)")

	// Should have 1 null order (user 3 with no orders)
	assert.Equal(t, 1, orderNulls, "Expected 1 null order (unmatched user)")
}

// TestDataFrame_CrossJoin tests the CrossJoin method
func TestDataFrame_CrossJoin(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame: products
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "product_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "product_name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	productIdBuilder := array.NewInt64Builder(pool)
	productIdBuilder.AppendValues([]int64{1, 2}, nil)
	productIdArray := productIdBuilder.NewArray()
	defer productIdArray.Release()

	productNameBuilder := array.NewStringBuilder(pool)
	productNameBuilder.AppendValues([]string{"Widget", "Gadget"}, nil)
	productNameArray := productNameBuilder.NewArray()
	defer productNameArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{productIdArray, productNameArray}, 2)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame: colors
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "color_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "color_name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	colorIdBuilder := array.NewInt64Builder(pool)
	colorIdBuilder.AppendValues([]int64{1, 2, 3}, nil)
	colorIdArray := colorIdBuilder.NewArray()
	defer colorIdArray.Release()

	colorNameBuilder := array.NewStringBuilder(pool)
	colorNameBuilder.AppendValues([]string{"Red", "Blue", "Green"}, nil)
	colorNameArray := colorNameBuilder.NewArray()
	defer colorNameArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{colorIdArray, colorNameArray}, 3)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform cross join (Cartesian product)
	result, err := leftDf.CrossJoin(rightDf, "", "")
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify result has 2 * 3 = 6 rows (Cartesian product)
	assert.Equal(t, int64(6), result.NumRows())

	// Verify schema has all columns from both DataFrames
	assert.Equal(t, int64(4), result.NumCols())
	assert.True(t, result.HasColumn("product_id"))
	assert.True(t, result.HasColumn("product_name"))
	assert.True(t, result.HasColumn("color_id"))
	assert.True(t, result.HasColumn("color_name"))

	// Verify data - each product should appear with each color
	productNameSeries, err := result.Column("product_name")
	require.NoError(t, err)
	productNameCol := productNameSeries.Array().(*array.String)

	colorNameSeries, err := result.Column("color_name")
	require.NoError(t, err)
	colorNameCol := colorNameSeries.Array().(*array.String)

	// Expected combinations:
	// Widget-Red, Widget-Blue, Widget-Green,
	// Gadget-Red, Gadget-Blue, Gadget-Green
	expectedCombos := map[string]bool{
		"Widget-Red":   false,
		"Widget-Blue":  false,
		"Widget-Green": false,
		"Gadget-Red":   false,
		"Gadget-Blue":  false,
		"Gadget-Green": false,
	}

	for i := 0; i < int(result.NumRows()); i++ {
		combo := productNameCol.Value(i) + "-" + colorNameCol.Value(i)
		expectedCombos[combo] = true
	}

	// All combinations should be present
	for combo, found := range expectedCombos {
		assert.True(t, found, "Expected combination %s not found", combo)
	}
}

// TestDataFrame_CrossJoin_LargeResult tests cross join overflow detection
func TestDataFrame_CrossJoin_LargeResult(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create large left DataFrame
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	// Create array with 50,000 rows
	leftSize := 50000
	leftBuilder := array.NewInt64Builder(pool)
	for i := 0; i < leftSize; i++ {
		leftBuilder.Append(int64(i))
	}
	leftArray := leftBuilder.NewArray()
	defer leftArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{leftArray}, int64(leftSize))
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create large right DataFrame
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	// Create array with 50,000 rows
	rightSize := 50000
	rightBuilder := array.NewInt64Builder(pool)
	for i := 0; i < rightSize; i++ {
		rightBuilder.Append(int64(i))
	}
	rightArray := rightBuilder.NewArray()
	defer rightArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{rightArray}, int64(rightSize))
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Attempt cross join - should fail due to overflow
	// 50,000 * 50,000 = 2.5 billion rows (exceeds 2 billion limit)
	result, err := leftDf.CrossJoin(rightDf, "", "")

	// Should return error
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "cross join result too large")
}

// TestDataFrame_RightJoin_WithNulls tests right join with null values in join keys
func TestDataFrame_RightJoin_WithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create left DataFrame with some null keys
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
			{Name: "value", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{1, 2}, []bool{true, true})
	idBuilder.AppendNull() // Null key
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	valueBuilder := array.NewStringBuilder(pool)
	valueBuilder.AppendValues([]string{"A", "B", "C"}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{idArray, valueArray}, 3)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create right DataFrame with null key
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "ref_id", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	refIdBuilder := array.NewInt64Builder(pool)
	refIdBuilder.Append(1)
	refIdBuilder.AppendNull() // Null key
	refIdArray := refIdBuilder.NewArray()
	defer refIdArray.Release()

	amountBuilder := array.NewFloat64Builder(pool)
	amountBuilder.AppendValues([]float64{100.0, 200.0}, nil)
	amountArray := amountBuilder.NewArray()
	defer amountArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{refIdArray, amountArray}, 2)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform right join
	result, err := leftDf.RightJoin(rightDf, "id", "ref_id")
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Should have 2 rows (all from right)
	assert.Equal(t, int64(2), result.NumRows())

	// Verify the matched row has non-null value
	valueSeries, err := result.Column("value")
	require.NoError(t, err)
	valueCol := valueSeries.Array().(*array.String)
	assert.False(t, valueCol.IsNull(0), "Row 0 should have matched value from left")

	// Null key row should have null value
	assert.True(t, valueCol.IsNull(1), "Row 1 (null key) should have null value")
}

// TestDataFrame_FullOuterJoin_EmptyDataFrames tests full outer join with empty DataFrames
func TestDataFrame_FullOuterJoin_EmptyDataFrames(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create empty left DataFrame
	leftSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	leftRecord := array.NewRecord(leftSchema, []arrow.Array{idArray, nameArray}, 0)
	defer leftRecord.Release()

	leftDf := NewDataFrame(leftRecord)
	defer leftDf.Release()

	// Create non-empty right DataFrame
	rightSchema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "ref_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	refIdBuilder := array.NewInt64Builder(pool)
	refIdBuilder.AppendValues([]int64{1, 2}, nil)
	refIdArray := refIdBuilder.NewArray()
	defer refIdArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{100.0, 200.0}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	rightRecord := array.NewRecord(rightSchema, []arrow.Array{refIdArray, valueArray}, 2)
	defer rightRecord.Release()

	rightDf := NewDataFrame(rightRecord)
	defer rightDf.Release()

	// Perform full outer join
	result, err := leftDf.FullOuterJoin(rightDf, "id", "ref_id")
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Should have 2 rows (all from right since left is empty)
	assert.Equal(t, int64(2), result.NumRows())

	// All name values should be null (no left matches)
	nameSeries, err := result.Column("name")
	require.NoError(t, err)
	nameCol := nameSeries.Array().(*array.String)
	for i := 0; i < int(result.NumRows()); i++ {
		assert.True(t, nameCol.IsNull(i), "Row %d should have null name", i)
	}
}
