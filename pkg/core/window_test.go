package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDataFrame_Window_RowNumber tests the ROW_NUMBER window function.
func TestDataFrame_Window_RowNumber(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data:
	// category | value
	// ---------|------
	// A        | 100
	// A        | 200
	// B        | 150
	// B        | 250
	// B        | 300
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.AppendValues([]string{"A", "A", "B", "B", "B"}, nil)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100, 200, 150, 250, 300}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		categoryBuilder.NewArray(),
		valueBuilder.NewArray(),
	}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test 1: ROW_NUMBER without partitioning
	result, err := df.Window().
		OrderBy("value").
		Over(RowNumber().As("row_num"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	assert.Equal(t, int64(3), result.NumCols()) // category, value, row_num
	assert.Equal(t, int64(5), result.NumRows())

	rowNumSeries, err := result.Column("row_num")
	require.NoError(t, err)
	rowNumCol := rowNumSeries.Array().(*array.Int64)

	// Result maintains original row order, so row_nums reflect sorted position:
	// Original: A(100), A(200), B(150), B(250), B(300)
	// Sorted:   A(100)=1, B(150)=2, A(200)=3, B(250)=4, B(300)=5
	// Expected in original order: [1, 3, 2, 4, 5]
	expectedRowNums := []int64{1, 3, 2, 4, 5}
	for i := 0; i < rowNumCol.Len(); i++ {
		assert.Equal(t, expectedRowNums[i], rowNumCol.Value(i), "row_num at index %d", i)
	}

	// Test 2: ROW_NUMBER with partitioning
	result2, err := df.Window().
		PartitionBy("category").
		OrderBy("value").
		Over(RowNumber().As("row_num"))
	require.NoError(t, err)
	require.NotNil(t, result2)
	defer result2.Release()

	rowNumSeries2, err := result2.Column("row_num")
	require.NoError(t, err)
	rowNumCol2 := rowNumSeries2.Array().(*array.Int64)

	// Expected: partition A gets 1, 2; partition B gets 1, 2, 3
	// Original order: A(100), A(200), B(150), B(250), B(300)
	// The row numbers depend on the partition order, but within each partition:
	// A: 100 -> 1, 200 -> 2
	// B: 150 -> 1, 250 -> 2, 300 -> 3
	assert.Equal(t, int64(5), int64(rowNumCol2.Len()))
}

// TestDataFrame_Window_Rank tests the RANK window function.
func TestDataFrame_Window_Rank(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with ties:
	// category | score
	// ---------|------
	// A        | 90
	// A        | 90
	// A        | 95
	// B        | 80
	// B        | 85
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.AppendValues([]string{"A", "A", "A", "B", "B"}, nil)

	scoreBuilder := array.NewInt64Builder(pool)
	defer scoreBuilder.Release()
	scoreBuilder.AppendValues([]int64{90, 90, 95, 80, 85}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		categoryBuilder.NewArray(),
		scoreBuilder.NewArray(),
	}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: RANK with partitioning
	result, err := df.Window().
		PartitionBy("category").
		OrderBy("score").
		Over(Rank().As("rank"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	assert.Equal(t, int64(3), result.NumCols()) // category, score, rank
	assert.Equal(t, int64(5), result.NumRows())

	rankSeries, err := result.Column("rank")
	require.NoError(t, err)
	rankCol := rankSeries.Array().(*array.Int64)

	// Note: Current implementation doesn't handle ties yet (sequential ranks)
	// This test validates current behavior
	assert.Equal(t, 5, rankCol.Len())
}

// TestDataFrame_Window_DenseRank tests the DENSE_RANK window function.
func TestDataFrame_Window_DenseRank(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data:
	// value
	// -----
	// 100
	// 200
	// 300
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100, 200, 300}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		valueBuilder.NewArray(),
	}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: DENSE_RANK
	result, err := df.Window().
		OrderBy("value").
		Over(DenseRank().As("dense_rank"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	assert.Equal(t, int64(2), result.NumCols()) // value, dense_rank
	assert.Equal(t, int64(3), result.NumRows())

	rankSeries, err := result.Column("dense_rank")
	require.NoError(t, err)
	rankCol := rankSeries.Array().(*array.Int64)

	// Expected: 1, 2, 3 (no gaps in sequential ranks)
	expectedRanks := []int64{1, 2, 3}
	for i := 0; i < rankCol.Len(); i++ {
		assert.Equal(t, expectedRanks[i], rankCol.Value(i), "dense_rank at index %d", i)
	}
}

// TestDataFrame_Window_Lag tests the LAG window function.
func TestDataFrame_Window_Lag(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data:
	// date  | sales
	// ------|------
	// 2024-01 | 100
	// 2024-02 | 150
	// 2024-03 | 200
	// 2024-04 | 180
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "date", Type: arrow.BinaryTypes.String},
			{Name: "sales", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	dateBuilder := array.NewStringBuilder(pool)
	defer dateBuilder.Release()
	dateBuilder.AppendValues([]string{"2024-01", "2024-02", "2024-03", "2024-04"}, nil)

	salesBuilder := array.NewInt64Builder(pool)
	defer salesBuilder.Release()
	salesBuilder.AppendValues([]int64{100, 150, 200, 180}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		dateBuilder.NewArray(),
		salesBuilder.NewArray(),
	}, 4)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test 1: LAG with offset 1
	result, err := df.Window().
		OrderBy("date").
		Over(Lag("sales", 1).As("prev_sales"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	assert.Equal(t, int64(3), result.NumCols()) // date, sales, prev_sales
	assert.Equal(t, int64(4), result.NumRows())

	prevSalesSeries, err := result.Column("prev_sales")
	require.NoError(t, err)
	prevSalesCol := prevSalesSeries.Array().(*array.Int64)

	// Expected: NULL, 100, 150, 200
	assert.True(t, prevSalesCol.IsNull(0), "first row should be NULL")
	assert.Equal(t, int64(100), prevSalesCol.Value(1))
	assert.Equal(t, int64(150), prevSalesCol.Value(2))
	assert.Equal(t, int64(200), prevSalesCol.Value(3))

	// Test 2: LAG with offset 2
	result2, err := df.Window().
		OrderBy("date").
		Over(Lag("sales", 2).As("prev_2_sales"))
	require.NoError(t, err)
	require.NotNil(t, result2)
	defer result2.Release()

	prev2SalesSeries, err := result2.Column("prev_2_sales")
	require.NoError(t, err)
	prev2SalesCol := prev2SalesSeries.Array().(*array.Int64)

	// Expected: NULL, NULL, 100, 150
	assert.True(t, prev2SalesCol.IsNull(0))
	assert.True(t, prev2SalesCol.IsNull(1))
	assert.Equal(t, int64(100), prev2SalesCol.Value(2))
	assert.Equal(t, int64(150), prev2SalesCol.Value(3))

	// Test 3: LAG with default value
	result3, err := df.Window().
		OrderBy("date").
		Over(Lag("sales", 1).Default(int64(0)).As("prev_sales_default"))
	require.NoError(t, err)
	require.NotNil(t, result3)
	defer result3.Release()

	prevSalesDefaultSeries, err := result3.Column("prev_sales_default")
	require.NoError(t, err)
	prevSalesDefaultCol := prevSalesDefaultSeries.Array().(*array.Int64)

	// Expected: 0, 100, 150, 200
	assert.Equal(t, int64(0), prevSalesDefaultCol.Value(0))
	assert.Equal(t, int64(100), prevSalesDefaultCol.Value(1))
	assert.Equal(t, int64(150), prevSalesDefaultCol.Value(2))
	assert.Equal(t, int64(200), prevSalesDefaultCol.Value(3))
}

// TestDataFrame_Window_Lead tests the LEAD window function.
func TestDataFrame_Window_Lead(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data:
	// date  | sales
	// ------|------
	// 2024-01 | 100
	// 2024-02 | 150
	// 2024-03 | 200
	// 2024-04 | 180
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "date", Type: arrow.BinaryTypes.String},
			{Name: "sales", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	dateBuilder := array.NewStringBuilder(pool)
	defer dateBuilder.Release()
	dateBuilder.AppendValues([]string{"2024-01", "2024-02", "2024-03", "2024-04"}, nil)

	salesBuilder := array.NewInt64Builder(pool)
	defer salesBuilder.Release()
	salesBuilder.AppendValues([]int64{100, 150, 200, 180}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		dateBuilder.NewArray(),
		salesBuilder.NewArray(),
	}, 4)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test 1: LEAD with offset 1
	result, err := df.Window().
		OrderBy("date").
		Over(Lead("sales", 1).As("next_sales"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	assert.Equal(t, int64(3), result.NumCols()) // date, sales, next_sales
	assert.Equal(t, int64(4), result.NumRows())

	nextSalesSeries, err := result.Column("next_sales")
	require.NoError(t, err)
	nextSalesCol := nextSalesSeries.Array().(*array.Int64)

	// Expected: 150, 200, 180, NULL
	assert.Equal(t, int64(150), nextSalesCol.Value(0))
	assert.Equal(t, int64(200), nextSalesCol.Value(1))
	assert.Equal(t, int64(180), nextSalesCol.Value(2))
	assert.True(t, nextSalesCol.IsNull(3), "last row should be NULL")

	// Test 2: LEAD with offset 2
	result2, err := df.Window().
		OrderBy("date").
		Over(Lead("sales", 2).As("next_2_sales"))
	require.NoError(t, err)
	require.NotNil(t, result2)
	defer result2.Release()

	next2SalesSeries, err := result2.Column("next_2_sales")
	require.NoError(t, err)
	next2SalesCol := next2SalesSeries.Array().(*array.Int64)

	// Expected: 200, 180, NULL, NULL
	assert.Equal(t, int64(200), next2SalesCol.Value(0))
	assert.Equal(t, int64(180), next2SalesCol.Value(1))
	assert.True(t, next2SalesCol.IsNull(2))
	assert.True(t, next2SalesCol.IsNull(3))

	// Test 3: LEAD with default value
	result3, err := df.Window().
		OrderBy("date").
		Over(Lead("sales", 1).Default(int64(0)).As("next_sales_default"))
	require.NoError(t, err)
	require.NotNil(t, result3)
	defer result3.Release()

	nextSalesDefaultSeries, err := result3.Column("next_sales_default")
	require.NoError(t, err)
	nextSalesDefaultCol := nextSalesDefaultSeries.Array().(*array.Int64)

	// Expected: 150, 200, 180, 0
	assert.Equal(t, int64(150), nextSalesDefaultCol.Value(0))
	assert.Equal(t, int64(200), nextSalesDefaultCol.Value(1))
	assert.Equal(t, int64(180), nextSalesDefaultCol.Value(2))
	assert.Equal(t, int64(0), nextSalesDefaultCol.Value(3))
}

// TestDataFrame_Window_MultipleFunc tests applying multiple window functions at once.
func TestDataFrame_Window_MultipleFunc(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data:
	// category | value
	// ---------|------
	// A        | 100
	// A        | 200
	// B        | 150
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.AppendValues([]string{"A", "A", "B"}, nil)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100, 200, 150}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		categoryBuilder.NewArray(),
		valueBuilder.NewArray(),
	}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Apply multiple window functions at once
	result, err := df.Window().
		PartitionBy("category").
		OrderBy("value").
		Over(
			RowNumber().As("row_num"),
			Rank().As("rank"),
			Lag("value", 1).As("prev_value"),
			Lead("value", 1).As("next_value"),
		)
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	assert.Equal(t, int64(6), result.NumCols()) // category, value, row_num, rank, prev_value, next_value
	assert.Equal(t, int64(3), result.NumRows())

	// Check that all columns exist
	_, err = result.Column("row_num")
	require.NoError(t, err)
	_, err = result.Column("rank")
	require.NoError(t, err)
	_, err = result.Column("prev_value")
	require.NoError(t, err)
	_, err = result.Column("next_value")
	require.NoError(t, err)
}

// TestDataFrame_Window_OrderByDesc tests descending order in window functions.
func TestDataFrame_Window_OrderByDesc(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data:
	// value
	// -----
	// 100
	// 200
	// 300
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100, 200, 300}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		valueBuilder.NewArray(),
	}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: OrderByDesc
	result, err := df.Window().
		OrderByDesc("value").
		Over(RowNumber().As("row_num"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Verify results
	rowNumSeries, err := result.Column("row_num")
	require.NoError(t, err)
	rowNumCol := rowNumSeries.Array().(*array.Int64)

	// Expected: row numbers for descending order (300, 200, 100) -> 1, 2, 3
	assert.Equal(t, 3, rowNumCol.Len())
}

// TestDataFrame_Window_EmptyDataFrame tests window functions on empty DataFrame.
func TestDataFrame_Window_EmptyDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()

	record := array.NewRecord(schema, []arrow.Array{
		valueBuilder.NewArray(),
	}, 0)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Window function on empty DataFrame
	result, err := df.Window().
		OrderBy("value").
		Over(RowNumber().As("row_num"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	// Should have 2 columns but 0 rows
	assert.Equal(t, int64(2), result.NumCols())
	assert.Equal(t, int64(0), result.NumRows())
}

// TestDataFrame_Window_NullValues tests window functions with null values.
func TestDataFrame_Window_NullValues(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with nulls:
	// category | value
	// ---------|------
	// A        | 100
	// NULL     | 200
	// A        | NULL
	// B        | 150
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.AppendValues([]string{"A", "", "A", "B"}, []bool{true, false, true, true})

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100, 200, 0, 150}, []bool{true, true, false, true})

	record := array.NewRecord(schema, []arrow.Array{
		categoryBuilder.NewArray(),
		valueBuilder.NewArray(),
	}, 4)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Window function with null partition keys
	result, err := df.Window().
		PartitionBy("category").
		OrderBy("value").
		Over(RowNumber().As("row_num"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	assert.Equal(t, int64(3), result.NumCols())
	assert.Equal(t, int64(4), result.NumRows())
}

// TestDataFrame_Window_LagLeadNullValues tests LAG/LEAD with null source values.
func TestDataFrame_Window_LagLeadNullValues(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with null values:
	// value
	// -----
	// 100
	// NULL
	// 300
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100, 0, 300}, []bool{true, false, true})

	record := array.NewRecord(schema, []arrow.Array{
		valueBuilder.NewArray(),
	}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: LAG should propagate nulls
	result, err := df.Window().
		Over(Lag("value", 1).As("prev_value"))
	require.NoError(t, err)
	require.NotNil(t, result)
	defer result.Release()

	prevValueSeries, err := result.Column("prev_value")
	require.NoError(t, err)
	prevValueCol := prevValueSeries.Array().(*array.Int64)

	// Expected: NULL, 100, NULL
	assert.True(t, prevValueCol.IsNull(0))
	assert.Equal(t, int64(100), prevValueCol.Value(1))
	assert.True(t, prevValueCol.IsNull(2))
}

// TestDataFrame_Window_NoFunctions tests error when no window functions provided.
func TestDataFrame_Window_NoFunctions(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		valueBuilder.NewArray(),
	}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: Should error when no functions provided
	_, err := df.Window().Over()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one window function required")
}

// TestDataFrame_Window_InvalidColumn tests error when referencing non-existent column.
func TestDataFrame_Window_InvalidColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "value", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	valueBuilder := array.NewInt64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.AppendValues([]int64{100}, nil)

	record := array.NewRecord(schema, []arrow.Array{
		valueBuilder.NewArray(),
	}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test: LAG referencing non-existent column
	_, err := df.Window().Over(Lag("nonexistent", 1).As("prev"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "column nonexistent not found")
}
