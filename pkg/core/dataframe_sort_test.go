package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestDataFrame_Sort_SingleColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with mixed order
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "score", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Create arrays with unsorted data
	idBuilder := array.NewInt64Builder(pool)
	idBuilder.AppendValues([]int64{3, 1, 2}, nil)
	idArray := idBuilder.NewArray()
	defer idArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Charlie", "Alice", "Bob"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	scoreBuilder := array.NewFloat64Builder(pool)
	scoreBuilder.AppendValues([]float64{85.5, 95.0, 90.2}, nil)
	scoreArray := scoreBuilder.NewArray()
	defer scoreArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, nameArray, scoreArray}, 3)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	testCases := []struct {
		name           string
		sortColumn     string
		ascending      bool
		expectedIds    []int64
		expectedNames  []string
		expectedScores []float64
	}{
		{
			name:           "Sort by ID ascending",
			sortColumn:     "id",
			ascending:      true,
			expectedIds:    []int64{1, 2, 3},
			expectedNames:  []string{"Alice", "Bob", "Charlie"},
			expectedScores: []float64{95.0, 90.2, 85.5},
		},
		{
			name:           "Sort by ID descending",
			sortColumn:     "id",
			ascending:      false,
			expectedIds:    []int64{3, 2, 1},
			expectedNames:  []string{"Charlie", "Bob", "Alice"},
			expectedScores: []float64{85.5, 90.2, 95.0},
		},
		{
			name:           "Sort by name ascending",
			sortColumn:     "name",
			ascending:      true,
			expectedIds:    []int64{1, 2, 3},
			expectedNames:  []string{"Alice", "Bob", "Charlie"},
			expectedScores: []float64{95.0, 90.2, 85.5},
		},
		{
			name:           "Sort by score descending",
			sortColumn:     "score",
			ascending:      false,
			expectedIds:    []int64{1, 2, 3},
			expectedNames:  []string{"Alice", "Bob", "Charlie"},
			expectedScores: []float64{95.0, 90.2, 85.5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sorted, err := df.Sort(tc.sortColumn, tc.ascending)
			if err != nil {
				t.Fatalf("Sort failed: %v", err)
			}
			defer sorted.Release()

			// Verify row count is preserved
			if sorted.NumRows() != df.NumRows() {
				t.Errorf("Expected %d rows after sort, got %d", df.NumRows(), sorted.NumRows())
			}

			// Verify column count is preserved
			if sorted.NumCols() != df.NumCols() {
				t.Errorf("Expected %d columns after sort, got %d", df.NumCols(), sorted.NumCols())
			}

			// Verify data is sorted correctly
			idSeries, err := sorted.Column("id")
			if err != nil {
				t.Fatalf("Failed to get id column: %v", err)
			}
			idArr := idSeries.Array().(*array.Int64)
			for i, expected := range tc.expectedIds {
				if idArr.Value(i) != expected {
					t.Errorf("Row %d: expected id %d, got %d", i, expected, idArr.Value(i))
				}
			}

			nameSeries, err := sorted.Column("name")
			if err != nil {
				t.Fatalf("Failed to get name column: %v", err)
			}
			nameArr := nameSeries.Array().(*array.String)
			for i, expected := range tc.expectedNames {
				if nameArr.Value(i) != expected {
					t.Errorf("Row %d: expected name %s, got %s", i, expected, nameArr.Value(i))
				}
			}

			scoreSeries, err := sorted.Column("score")
			if err != nil {
				t.Fatalf("Failed to get score column: %v", err)
			}
			scoreArr := scoreSeries.Array().(*array.Float64)
			for i, expected := range tc.expectedScores {
				if scoreArr.Value(i) != expected {
					t.Errorf("Row %d: expected score %.1f, got %.1f", i, expected, scoreArr.Value(i))
				}
			}
		})
	}
}

func TestDataFrame_Sort_MultiColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test data with duplicate values to test multi-column sorting
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	deptBuilder := array.NewStringBuilder(pool)
	deptBuilder.AppendValues([]string{"IT", "HR", "IT", "HR", "IT"}, nil)
	deptArray := deptBuilder.NewArray()
	defer deptArray.Release()

	salaryBuilder := array.NewInt64Builder(pool)
	salaryBuilder.AppendValues([]int64{70000, 60000, 80000, 60000, 70000}, nil)
	salaryArray := salaryBuilder.NewArray()
	defer salaryArray.Release()

	nameBuilder := array.NewStringBuilder(pool)
	nameBuilder.AppendValues([]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}, nil)
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{deptArray, salaryArray, nameArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test multi-column sort: department ascending, then salary descending
	sorted, err := df.SortMultiple([]SortKey{
		{Column: "department", Ascending: true},
		{Column: "salary", Ascending: false},
	})
	if err != nil {
		t.Fatalf("Multi-column sort failed: %v", err)
	}
	defer sorted.Release()

	// Expected order: HR-60000-Bob, HR-60000-Diana, IT-80000-Charlie, IT-70000-Alice, IT-70000-Eve
	expectedDepts := []string{"HR", "HR", "IT", "IT", "IT"}
	expectedSalaries := []int64{60000, 60000, 80000, 70000, 70000}
	expectedNames := []string{"Bob", "Diana", "Charlie", "Alice", "Eve"}

	deptSeries, _ := sorted.Column("department")
	deptArr := deptSeries.Array().(*array.String)
	salarySeries, _ := sorted.Column("salary")
	salaryArr := salarySeries.Array().(*array.Int64)
	nameSeries, _ := sorted.Column("name")
	nameArr := nameSeries.Array().(*array.String)

	for i := 0; i < int(sorted.NumRows()); i++ {
		if deptArr.Value(i) != expectedDepts[i] {
			t.Errorf("Row %d: expected dept %s, got %s", i, expectedDepts[i], deptArr.Value(i))
		}
		if salaryArr.Value(i) != expectedSalaries[i] {
			t.Errorf("Row %d: expected salary %d, got %d", i, expectedSalaries[i], salaryArr.Value(i))
		}
		if nameArr.Value(i) != expectedNames[i] {
			t.Errorf("Row %d: expected name %s, got %s", i, expectedNames[i], nameArr.Value(i))
		}
	}
}

func TestDataFrame_Sort_ErrorCases(t *testing.T) {
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

	df := NewDataFrame(record)
	defer df.Release()

	// Test sort on non-existent column
	_, err := df.Sort("nonexistent", true)
	if err == nil {
		t.Error("Expected error for non-existent column")
	}

	// Test multi-column sort with non-existent column
	_, err = df.SortMultiple([]SortKey{
		{Column: "test", Ascending: true},
		{Column: "nonexistent", Ascending: false},
	})
	if err == nil {
		t.Error("Expected error for non-existent column in multi-sort")
	}

	// Test empty sort keys
	_, err = df.SortMultiple([]SortKey{})
	if err == nil {
		t.Error("Expected error for empty sort keys")
	}
}

func TestDataFrame_Sort_PreservesImmutability(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "values", Type: arrow.PrimitiveTypes.Int64}},
		nil,
	)

	builder := array.NewInt64Builder(pool)
	builder.AppendValues([]int64{3, 1, 2}, nil)
	arr := builder.NewArray()
	defer arr.Release()

	record := array.NewRecord(schema, []arrow.Array{arr}, 3)
	defer record.Release()

	original := NewDataFrame(record)
	defer original.Release()

	// Sort the DataFrame
	sorted, err := original.Sort("values", true)
	if err != nil {
		t.Fatalf("Sort failed: %v", err)
	}
	defer sorted.Release()

	// Verify original is unchanged
	originalSeries, _ := original.Column("values")
	originalArr := originalSeries.Array().(*array.Int64)
	if originalArr.Value(0) != 3 || originalArr.Value(1) != 1 || originalArr.Value(2) != 2 {
		t.Error("Original DataFrame was modified by sort operation")
	}

	// Verify sorted has correct order
	sortedSeries, _ := sorted.Column("values")
	sortedArr := sortedSeries.Array().(*array.Int64)
	if sortedArr.Value(0) != 1 || sortedArr.Value(1) != 2 || sortedArr.Value(2) != 3 {
		t.Error("Sorted DataFrame does not have correct order")
	}
}
