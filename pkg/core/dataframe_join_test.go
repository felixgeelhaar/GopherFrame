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
