package aggregation

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupBy_Percentile(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with groups
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	// Group A: 10, 20, 30 -> 75th percentile should be 25
	// Group B: 40, 50, 60 -> 75th percentile should be 55
	valueBuilder.Append(10)
	valueBuilder.Append(20)
	valueBuilder.Append(30)
	valueBuilder.Append(40)
	valueBuilder.Append(50)
	valueBuilder.Append(60)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 6)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test 75th percentile
	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column:     "value",
				Type:       Percentile,
				Alias:      "value_p75",
				Percentile: 0.75,
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	require.NotNil(t, result.DataFrame)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	assert.Equal(t, int64(2), resultRecord.NumRows()) // Two groups: A and B

	// Check results
	percentileCol := resultRecord.Column(1).(*array.Float64)
	assert.InDelta(t, 25.0, percentileCol.Value(0), 0.01) // Group A
	assert.InDelta(t, 55.0, percentileCol.Value(1), 0.01) // Group B
}

func TestGroupBy_Median(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with groups
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	// Group A: 10, 20, 30 -> median should be 20
	// Group B: 40, 50 -> median should be 45
	valueBuilder.Append(10)
	valueBuilder.Append(20)
	valueBuilder.Append(30)
	valueBuilder.Append(40)
	valueBuilder.Append(50)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 5)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test median
	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column: "value",
				Type:   Median,
				Alias:  "value_median",
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	require.NotNil(t, result.DataFrame)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	assert.Equal(t, int64(2), resultRecord.NumRows())

	// Check results
	medianCol := resultRecord.Column(1).(*array.Float64)
	assert.InDelta(t, 20.0, medianCol.Value(0), 0.01) // Group A
	assert.InDelta(t, 45.0, medianCol.Value(1), 0.01) // Group B
}

func TestGroupBy_Mode(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with groups
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	// Group A: 10, 10, 10, 20 -> mode should be 10
	// Group B: 30, 30, 40 -> mode should be 30
	valueBuilder.Append(10)
	valueBuilder.Append(10)
	valueBuilder.Append(10)
	valueBuilder.Append(20)
	valueBuilder.Append(30)
	valueBuilder.Append(30)
	valueBuilder.Append(40)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 7)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test mode
	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column: "value",
				Type:   Mode,
				Alias:  "value_mode",
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	require.NotNil(t, result.DataFrame)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	assert.Equal(t, int64(2), resultRecord.NumRows())

	// Check results
	modeCol := resultRecord.Column(1).(*array.Float64)
	assert.Equal(t, 10.0, modeCol.Value(0)) // Group A
	assert.Equal(t, 30.0, modeCol.Value(1)) // Group B
}

func TestGroupBy_Correlation(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with groups
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryBuilder.Append("B")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	xBuilder := array.NewFloat64Builder(pool)
	defer xBuilder.Release()
	// Group A: perfect positive correlation
	xBuilder.Append(1)
	xBuilder.Append(2)
	xBuilder.Append(3)
	// Group B: perfect negative correlation
	xBuilder.Append(1)
	xBuilder.Append(2)
	xBuilder.Append(3)
	xArray := xBuilder.NewArray()
	defer xArray.Release()

	yBuilder := array.NewFloat64Builder(pool)
	defer yBuilder.Release()
	// Group A: y = 2*x -> perfect positive correlation (r = 1.0)
	yBuilder.Append(2)
	yBuilder.Append(4)
	yBuilder.Append(6)
	// Group B: y = 4 - x -> perfect negative correlation (r = -1.0)
	yBuilder.Append(3)
	yBuilder.Append(2)
	yBuilder.Append(1)
	yArray := yBuilder.NewArray()
	defer yArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "x", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
		{Name: "y", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, xArray, yArray}, 6)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test correlation
	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column:       "x",
				SecondColumn: "y",
				Type:         Correlation,
				Alias:        "corr_xy",
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	require.NotNil(t, result.DataFrame)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	assert.Equal(t, int64(2), resultRecord.NumRows())

	// Check results
	corrCol := resultRecord.Column(1).(*array.Float64)
	assert.InDelta(t, 1.0, corrCol.Value(0), 0.0001)  // Group A: perfect positive
	assert.InDelta(t, -1.0, corrCol.Value(1), 0.0001) // Group B: perfect negative
}

func TestGroupBy_Percentile_WithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with nulls
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	// Group A: 10, null, 20, 30 -> skip null, 75th percentile of [10, 20, 30] should be 25
	valueBuilder.Append(10)
	valueBuilder.AppendNull()
	valueBuilder.Append(20)
	valueBuilder.Append(30)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 4)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column:     "value",
				Type:       Percentile,
				Alias:      "value_p75",
				Percentile: 0.75,
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	percentileCol := resultRecord.Column(1).(*array.Float64)
	assert.InDelta(t, 25.0, percentileCol.Value(0), 0.01)
}

func TestGroupBy_Mode_AllSameFrequency(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data where all values have same frequency
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	// Group A: 10, 20, 30 -> all have frequency 1, should return one of them
	valueBuilder.Append(10)
	valueBuilder.Append(20)
	valueBuilder.Append(30)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 3)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column: "value",
				Type:   Mode,
				Alias:  "value_mode",
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	modeCol := resultRecord.Column(1).(*array.Float64)
	// Should be one of the values
	mode := modeCol.Value(0)
	assert.True(t, mode == 10.0 || mode == 20.0 || mode == 30.0)
}

func TestGroupBy_Correlation_InsufficientData(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with only 1 pair
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	xBuilder := array.NewFloat64Builder(pool)
	defer xBuilder.Release()
	xBuilder.Append(1)
	xArray := xBuilder.NewArray()
	defer xArray.Release()

	yBuilder := array.NewFloat64Builder(pool)
	defer yBuilder.Release()
	yBuilder.Append(2)
	yArray := yBuilder.NewArray()
	defer yArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "x", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
		{Name: "y", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, xArray, yArray}, 1)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column:       "x",
				SecondColumn: "y",
				Type:         Correlation,
				Alias:        "corr_xy",
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	corrCol := resultRecord.Column(1).(*array.Float64)
	// Should be null because we need at least 2 pairs
	assert.True(t, corrCol.IsNull(0))
}

func TestGroupBy_Correlation_NoVariance(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with no variance in one variable
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryBuilder.Append("A")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	xBuilder := array.NewFloat64Builder(pool)
	defer xBuilder.Release()
	// x has no variance (all same value)
	xBuilder.Append(5)
	xBuilder.Append(5)
	xBuilder.Append(5)
	xArray := xBuilder.NewArray()
	defer xArray.Release()

	yBuilder := array.NewFloat64Builder(pool)
	defer yBuilder.Release()
	yBuilder.Append(1)
	yBuilder.Append(2)
	yBuilder.Append(3)
	yArray := yBuilder.NewArray()
	defer yArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "x", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
		{Name: "y", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, xArray, yArray}, 3)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{
				Column:       "x",
				SecondColumn: "y",
				Type:         Correlation,
				Alias:        "corr_xy",
			},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	corrCol := resultRecord.Column(1).(*array.Float64)
	// Should be null because correlation is undefined when one variable has no variance
	assert.True(t, corrCol.IsNull(0))
}

func TestGroupBy_MultipleStatisticalAggregations(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	for i := 0; i < 5; i++ {
		categoryBuilder.Append("A")
	}
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	// Values: 10, 20, 30, 40, 50
	for i := 1; i <= 5; i++ {
		valueBuilder.Append(float64(i * 10))
	}
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 5)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test multiple aggregations
	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{Column: "value", Type: Mean, Alias: "mean"},
			{Column: "value", Type: Median, Alias: "median"},
			{Column: "value", Type: Mode, Alias: "mode"},
			{Column: "value", Type: Percentile, Alias: "p25", Percentile: 0.25},
			{Column: "value", Type: Percentile, Alias: "p75", Percentile: 0.75},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	assert.Equal(t, int64(1), resultRecord.NumRows())
	assert.Equal(t, int64(6), resultRecord.NumCols()) // category + 5 aggregations

	// Verify results
	meanCol := resultRecord.Column(1).(*array.Float64)
	assert.InDelta(t, 30.0, meanCol.Value(0), 0.01)

	medianCol := resultRecord.Column(2).(*array.Float64)
	assert.InDelta(t, 30.0, medianCol.Value(0), 0.01)

	p25Col := resultRecord.Column(4).(*array.Float64)
	assert.InDelta(t, 20.0, p25Col.Value(0), 0.01)

	p75Col := resultRecord.Column(5).(*array.Float64)
	assert.InDelta(t, 40.0, p75Col.Value(0), 0.01)
}

func TestGroupBy_Percentile_InvalidRange(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create simple test data
	categoryBuilder := array.NewStringBuilder(pool)
	defer categoryBuilder.Release()
	categoryBuilder.Append("A")
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.Append(10)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 1)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	service := NewGroupByService()

	// Test percentile > 1.0
	request := GroupByRequest{
		GroupColumns: []string{"category"},
		Aggregations: []AggregationSpec{
			{Column: "value", Type: Percentile, Alias: "p150", Percentile: 1.5},
		},
	}

	result := service.Execute(df, request)
	assert.Error(t, result.Error)
	assert.Contains(t, result.Error.Error(), "percentile must be between 0.0 and 1.0")

	// Test percentile < 0.0
	request.Aggregations[0].Percentile = -0.1
	result = service.Execute(df, request)
	assert.Error(t, result.Error)
	assert.Contains(t, result.Error.Error(), "percentile must be between 0.0 and 1.0")
}

func TestGroupBy_MultiColumnGroup_WithStatistical(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with multi-column groups
	col1Builder := array.NewStringBuilder(pool)
	defer col1Builder.Release()
	col1Builder.Append("A")
	col1Builder.Append("A")
	col1Builder.Append("B")
	col1Builder.Append("B")
	col1Array := col1Builder.NewArray()
	defer col1Array.Release()

	col2Builder := array.NewStringBuilder(pool)
	defer col2Builder.Release()
	col2Builder.Append("X")
	col2Builder.Append("Y")
	col2Builder.Append("X")
	col2Builder.Append("Y")
	col2Array := col2Builder.NewArray()
	defer col2Array.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	defer valueBuilder.Release()
	valueBuilder.Append(10)
	valueBuilder.Append(20)
	valueBuilder.Append(30)
	valueBuilder.Append(40)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "col1", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "col2", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "value", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{col1Array, col2Array, valueArray}, 4)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	service := NewGroupByService()
	request := GroupByRequest{
		GroupColumns: []string{"col1", "col2"},
		Aggregations: []AggregationSpec{
			{Column: "value", Type: Median, Alias: "median"},
		},
	}

	result := service.Execute(df, request)
	require.NoError(t, result.Error)
	defer result.DataFrame.Release()

	resultRecord := result.DataFrame.Record()
	assert.Equal(t, int64(4), resultRecord.NumRows()) // 4 unique combinations

	// Each group has 1 value, so median equals that value
	medianCol := resultRecord.Column(2).(*array.Float64)
	assert.Equal(t, 10.0, medianCol.Value(0)) // A-X
	assert.Equal(t, 20.0, medianCol.Value(1)) // A-Y
	assert.Equal(t, 30.0, medianCol.Value(2)) // B-X
	assert.Equal(t, 40.0, medianCol.Value(3)) // B-Y
}
