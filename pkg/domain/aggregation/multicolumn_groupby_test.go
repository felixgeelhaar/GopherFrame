package aggregation

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

func TestGroupByService_MultiColumnGroupBy(t *testing.T) {
	service := NewGroupByService()

	// Create test DataFrame with multi-dimensional data
	// Data represents employees with department, level, and salary
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "level", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
			{Name: "bonus", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build test data - employees in different departments and levels
	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{
		"Engineering", "Engineering", "Engineering", "Engineering",
		"Sales", "Sales", "Sales",
		"Marketing", "Marketing",
	}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	levelBuilder := array.NewStringBuilder(pool)
	levelBuilder.AppendValues([]string{
		"Senior", "Senior", "Junior", "Junior",
		"Senior", "Junior", "Junior",
		"Senior", "Junior",
	}, nil)
	levelArray := levelBuilder.NewArray()
	defer levelArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryBuilder.AppendValues([]float64{
		120000, 115000, 80000, 85000,
		100000, 70000, 75000,
		95000, 65000,
	}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	bonusBuilder := array.NewFloat64Builder(pool)
	bonusBuilder.AppendValues([]float64{
		15000, 12000, 8000, 10000,
		20000, 5000, 7000,
		12000, 4000,
	}, nil)
	bonusArray := bonusBuilder.NewArray()
	defer bonusArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, levelArray, salaryArray, bonusArray}, 9)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test multi-column GroupBy: department + level
	request := GroupByRequest{
		GroupColumns: []string{"department", "level"},
		Aggregations: []AggregationSpec{
			{Column: "salary", Type: Mean, Alias: "avg_salary"},
			{Column: "bonus", Type: Sum, Alias: "total_bonus"},
			{Column: "salary", Type: Count, Alias: "employee_count"},
		},
	}

	result := service.Execute(df, request)
	if result.Error != nil {
		t.Fatalf("Multi-column GroupBy failed: %v", result.Error)
	}
	defer result.DataFrame.Release()

	// Verify result structure
	if result.DataFrame.NumCols() != 5 { // 2 group cols + 3 aggregations
		t.Errorf("Expected 5 columns, got %d", result.DataFrame.NumCols())
	}

	// Expected groups: (Engineering,Junior), (Engineering,Senior), (Marketing,Junior), (Marketing,Senior), (Sales,Junior), (Sales,Senior)
	expectedRows := int64(6)
	if result.DataFrame.NumRows() != expectedRows {
		t.Errorf("Expected %d rows, got %d", expectedRows, result.DataFrame.NumRows())
	}

	// Verify column names
	expectedColumns := []string{"department", "level", "avg_salary", "total_bonus", "employee_count"}
	actualColumns := result.DataFrame.ColumnNames()
	if len(actualColumns) != len(expectedColumns) {
		t.Errorf("Expected %d columns, got %d", len(expectedColumns), len(actualColumns))
	}
	for i, expected := range expectedColumns {
		if i < len(actualColumns) && actualColumns[i] != expected {
			t.Errorf("Column %d: expected %s, got %s", i, expected, actualColumns[i])
		}
	}

	// Verify specific group results
	resultRecord := result.DataFrame.Record()
	deptCol := resultRecord.Column(0).(*array.String)
	levelCol := resultRecord.Column(1).(*array.String)
	avgSalaryCol := resultRecord.Column(2).(*array.Float64)
	totalBonusCol := resultRecord.Column(3).(*array.Float64)
	countCol := resultRecord.Column(4).(*array.Int64)

	// Test specific expected values for known groups
	testCases := []struct {
		dept           string
		level          string
		expectedAvgSal float64
		expectedBonus  float64
		expectedCount  int64
	}{
		{"Engineering", "Junior", 82500.0, 18000.0, 2},  // (80000+85000)/2, 8000+10000, 2 employees
		{"Engineering", "Senior", 117500.0, 27000.0, 2}, // (120000+115000)/2, 15000+12000, 2 employees
		{"Marketing", "Junior", 65000.0, 4000.0, 1},     // 65000, 4000, 1 employee
		{"Marketing", "Senior", 95000.0, 12000.0, 1},    // 95000, 12000, 1 employee
		{"Sales", "Junior", 72500.0, 12000.0, 2},        // (70000+75000)/2, 5000+7000, 2 employees
		{"Sales", "Senior", 100000.0, 20000.0, 1},       // 100000, 20000, 1 employee
	}

	// Find and verify each expected group
	for _, tc := range testCases {
		found := false
		for i := 0; i < int(result.DataFrame.NumRows()); i++ {
			if deptCol.Value(i) != tc.dept || levelCol.Value(i) != tc.level {
				continue
			}
			found = true

			if avgSalaryCol.Value(i) != tc.expectedAvgSal {
				t.Errorf("Group (%s,%s): expected avg salary %.1f, got %.1f",
					tc.dept, tc.level, tc.expectedAvgSal, avgSalaryCol.Value(i))
			}
			if totalBonusCol.Value(i) != tc.expectedBonus {
				t.Errorf("Group (%s,%s): expected total bonus %.1f, got %.1f",
					tc.dept, tc.level, tc.expectedBonus, totalBonusCol.Value(i))
			}
			if countCol.Value(i) != tc.expectedCount {
				t.Errorf("Group (%s,%s): expected count %d, got %d",
					tc.dept, tc.level, tc.expectedCount, countCol.Value(i))
			}
			break
		}
		if !found {
			t.Errorf("Expected group (%s,%s) not found in results", tc.dept, tc.level)
		}
	}
}

func TestGroupByService_ThreeColumnGroupBy(t *testing.T) {
	service := NewGroupByService()

	// Create test DataFrame with 3-dimensional grouping
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "region", Type: arrow.BinaryTypes.String},
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "level", Type: arrow.BinaryTypes.String},
			{Name: "sales", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build test data - sales by region, department, and level
	regionBuilder := array.NewStringBuilder(pool)
	regionBuilder.AppendValues([]string{
		"US", "US", "US", "US",
		"EU", "EU", "EU", "EU",
	}, nil)
	regionArray := regionBuilder.NewArray()
	defer regionArray.Release()

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{
		"Sales", "Sales", "Engineering", "Engineering",
		"Sales", "Sales", "Engineering", "Engineering",
	}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	levelBuilder := array.NewStringBuilder(pool)
	levelBuilder.AppendValues([]string{
		"Senior", "Junior", "Senior", "Junior",
		"Senior", "Junior", "Senior", "Junior",
	}, nil)
	levelArray := levelBuilder.NewArray()
	defer levelArray.Release()

	salesBuilder := array.NewFloat64Builder(pool)
	salesBuilder.AppendValues([]float64{
		150000, 100000, 80000, 60000,
		120000, 80000, 70000, 50000,
	}, nil)
	salesArray := salesBuilder.NewArray()
	defer salesArray.Release()

	record := array.NewRecord(schema, []arrow.Array{regionArray, deptArray, levelArray, salesArray}, 8)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test three-column GroupBy
	request := GroupByRequest{
		GroupColumns: []string{"region", "department", "level"},
		Aggregations: []AggregationSpec{
			{Column: "sales", Type: Sum, Alias: "total_sales"},
		},
	}

	result := service.Execute(df, request)
	if result.Error != nil {
		t.Fatalf("Three-column GroupBy failed: %v", result.Error)
	}
	defer result.DataFrame.Release()

	// Should have 8 unique groups (2 regions × 2 depts × 2 levels)
	if result.DataFrame.NumRows() != 8 {
		t.Errorf("Expected 8 groups, got %d", result.DataFrame.NumRows())
	}

	// Verify we have 4 columns: 3 group columns + 1 aggregation
	if result.DataFrame.NumCols() != 4 {
		t.Errorf("Expected 4 columns, got %d", result.DataFrame.NumCols())
	}
}

func TestGroupByService_MultiColumnGroupBy_ErrorCases(t *testing.T) {
	service := NewGroupByService()

	// Create simple test DataFrame
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "dept", Type: arrow.BinaryTypes.String},
			{Name: "level", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"IT", "HR"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	levelBuilder := array.NewStringBuilder(pool)
	levelBuilder.AppendValues([]string{"Senior", "Junior"}, nil)
	levelArray := levelBuilder.NewArray()
	defer levelArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryBuilder.AppendValues([]float64{100000, 80000}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, levelArray, salaryArray}, 2)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	// Test with non-existent group column
	request := GroupByRequest{
		GroupColumns: []string{"dept", "nonexistent"},
		Aggregations: []AggregationSpec{
			{Column: "salary", Type: Sum, Alias: "total_salary"},
		},
	}

	result := service.Execute(df, request)
	if result.Error == nil {
		t.Error("Expected error for non-existent group column")
	}

	// Test with non-existent aggregation column
	request = GroupByRequest{
		GroupColumns: []string{"dept", "level"},
		Aggregations: []AggregationSpec{
			{Column: "nonexistent", Type: Sum, Alias: "total"},
		},
	}

	result = service.Execute(df, request)
	if result.Error == nil {
		t.Error("Expected error for non-existent aggregation column")
	}
}

func TestGroupByService_MultiColumnGroupBy_EmptyGroups(t *testing.T) {
	service := NewGroupByService()

	// Create test DataFrame with no data
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "dept", Type: arrow.BinaryTypes.String},
			{Name: "level", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	deptBuilder := array.NewStringBuilder(pool)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	levelBuilder := array.NewStringBuilder(pool)
	levelArray := levelBuilder.NewArray()
	defer levelArray.Release()

	salaryBuilder := array.NewFloat64Builder(pool)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, levelArray, salaryArray}, 0)
	defer record.Release()

	df := dataframe.NewDataFrame(record)
	defer df.Release()

	request := GroupByRequest{
		GroupColumns: []string{"dept", "level"},
		Aggregations: []AggregationSpec{
			{Column: "salary", Type: Sum, Alias: "total_salary"},
		},
	}

	result := service.Execute(df, request)
	if result.Error != nil {
		t.Fatalf("Empty DataFrame GroupBy failed: %v", result.Error)
	}
	defer result.DataFrame.Release()

	// Should return empty result with correct schema
	if result.DataFrame.NumRows() != 0 {
		t.Errorf("Expected 0 rows for empty input, got %d", result.DataFrame.NumRows())
	}

	if result.DataFrame.NumCols() != 3 { // 2 group cols + 1 aggregation
		t.Errorf("Expected 3 columns, got %d", result.DataFrame.NumCols())
	}
}
