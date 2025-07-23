// phase1_test.go tests all the Phase 1 MVP functionality to ensure it's complete
package gopherframe

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
)

// TestPhase1_CompleteWorkflow tests the entire Phase 1 MVP workflow
func TestPhase1_CompleteWorkflow(t *testing.T) {
	// Create test data
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "country", Type: arrow.BinaryTypes.String},
			{Name: "city", Type: arrow.BinaryTypes.String},
			{Name: "population", Type: arrow.PrimitiveTypes.Int64},
			{Name: "gdp", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build test arrays
	countryBuilder := array.NewStringBuilder(pool)
	countryBuilder.AppendValues([]string{"US", "US", "UK", "UK", "FR"}, nil)
	countryArray := countryBuilder.NewArray()
	defer countryArray.Release()

	cityBuilder := array.NewStringBuilder(pool)
	cityBuilder.AppendValues([]string{"NYC", "LA", "London", "Manchester", "Paris"}, nil)
	cityArray := cityBuilder.NewArray()
	defer cityArray.Release()

	populationBuilder := array.NewInt64Builder(pool)
	populationBuilder.AppendValues([]int64{8500000, 4000000, 9000000, 500000, 2200000}, nil)
	populationArray := populationBuilder.NewArray()
	defer populationArray.Release()

	gdpBuilder := array.NewFloat64Builder(pool)
	gdpBuilder.AppendValues([]float64{1200.5, 900.3, 800.7, 200.1, 750.9}, nil)
	gdpArray := gdpBuilder.NewArray()
	defer gdpArray.Release()

	// Create record and DataFrame
	record := array.NewRecord(schema, []arrow.Array{countryArray, cityArray, populationArray, gdpArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test 1: Selection/Projection - Select("col1", "col2")
	t.Run("Selection", func(t *testing.T) {
		selected := df.Select("country", "city")
		defer selected.Release()

		if selected.Err() != nil {
			t.Fatalf("Select failed: %v", selected.Err())
		}

		if selected.NumCols() != 2 {
			t.Errorf("Expected 2 columns after select, got %d", selected.NumCols())
		}

		expectedCols := []string{"country", "city"}
		actualCols := selected.ColumnNames()
		for i, expected := range expectedCols {
			if actualCols[i] != expected {
				t.Errorf("Column %d: expected %s, got %s", i, expected, actualCols[i])
			}
		}
	})

	// Test 2: Filtering - Filter(df.Col("country").Eq("US"))
	t.Run("Filtering", func(t *testing.T) {
		filtered := df.Filter(df.Col("country").Eq(expr.Lit("US")))
		defer filtered.Release()

		if filtered.Err() != nil {
			t.Fatalf("Filter failed: %v", filtered.Err())
		}

		if filtered.NumRows() != 2 {
			t.Errorf("Expected 2 rows after filter, got %d", filtered.NumRows())
		}
	})

	// Test 3: Column Operations - WithColumn("profit", df.Col("revenue").Sub(df.Col("cost")))
	t.Run("ColumnOperations", func(t *testing.T) {
		// Create a calculated column: boosted_gdp = gdp * 2
		// Keep it simple to avoid type conversion issues in the expression engine
		withCol := df.WithColumn("boosted_gdp",
			df.Col("gdp").Mul(expr.Lit(2.0)))
		defer withCol.Release()

		if withCol.Err() != nil {
			t.Fatalf("WithColumn failed: %v", withCol.Err())
		}

		if withCol.NumCols() != 5 { // original 4 + new 1
			t.Errorf("Expected 5 columns after WithColumn, got %d", withCol.NumCols())
		}

		cols := withCol.ColumnNames()
		found := false
		for _, col := range cols {
			if col == "boosted_gdp" {
				found = true
				break
			}
		}
		if !found {
			t.Error("New column 'boosted_gdp' not found")
		}
	})

	// Test 4: Sorting - Sort(df.By("signup_date", df.Desc))
	t.Run("Sorting", func(t *testing.T) {
		// Sort by population descending
		sorted := df.SortMultiple([]SortKey{
			By("population", false), // descending
		})
		defer sorted.Release()

		if sorted.Err() != nil {
			t.Fatalf("Sort failed: %v", sorted.Err())
		}

		// Verify sorting - population descending: 9000000(London), 8500000(NYC), 4000000(LA), 2200000(Paris), 500000(Manchester)
		record := sorted.Record()
		cityCol := record.Column(1).(*array.String) // city is column 1

		expectedOrder := []string{"London", "NYC", "LA", "Paris", "Manchester"}
		for i, expected := range expectedOrder {
			if cityCol.Value(i) != expected {
				t.Errorf("Row %d: expected city %s, got %s", i, expected, cityCol.Value(i))
			}
		}
	})

	// Test 5: I/O Operations (Parquet)
	t.Run("ParquetIO", func(t *testing.T) {
		tempDir := t.TempDir()
		filename := filepath.Join(tempDir, "test.parquet")

		// Write to Parquet
		err := WriteParquet(df, filename)
		if err != nil {
			t.Fatalf("WriteParquet failed: %v", err)
		}

		// Read from Parquet
		readDF, err := ReadParquet(filename)
		if err != nil {
			t.Fatalf("ReadParquet failed: %v", err)
		}
		defer readDF.Release()

		// Verify data integrity
		if readDF.NumRows() != df.NumRows() {
			t.Errorf("Row count mismatch: expected %d, got %d", df.NumRows(), readDF.NumRows())
		}
		if readDF.NumCols() != df.NumCols() {
			t.Errorf("Column count mismatch: expected %d, got %d", df.NumCols(), readDF.NumCols())
		}
	})

	// Test 6: CSV I/O
	t.Run("CSVIO", func(t *testing.T) {
		tempDir := t.TempDir()
		filename := filepath.Join(tempDir, "test.csv")

		// Write to CSV
		err := WriteCSV(df, filename)
		if err != nil {
			t.Fatalf("WriteCSV failed: %v", err)
		}

		// Verify file exists and has content
		if _, statErr := os.Stat(filename); os.IsNotExist(statErr) {
			t.Error("CSV file was not created")
		}

		// Read from CSV
		readDF, err := ReadCSV(filename)
		if err != nil {
			t.Fatalf("ReadCSV failed: %v", err)
		}
		defer readDF.Release()

		// Basic integrity check (CSV may have type inference differences)
		if readDF.NumRows() != df.NumRows() {
			t.Errorf("Row count mismatch: expected %d, got %d", df.NumRows(), readDF.NumRows())
		}
	})

	// Test 7: Arrow IPC I/O
	t.Run("ArrowIO", func(t *testing.T) {
		tempDir := t.TempDir()
		filename := filepath.Join(tempDir, "test.arrow")

		// Write to Arrow IPC
		err := WriteArrowIPC(df, filename)
		if err != nil {
			t.Fatalf("WriteArrowIPC failed: %v", err)
		}

		// Read from Arrow IPC
		readDF, err := ReadArrowIPC(filename)
		if err != nil {
			t.Fatalf("ReadArrowIPC failed: %v", err)
		}
		defer readDF.Release()

		// Verify perfect data integrity (Arrow should preserve everything)
		if readDF.NumRows() != df.NumRows() {
			t.Errorf("Row count mismatch: expected %d, got %d", df.NumRows(), readDF.NumRows())
		}
		if readDF.NumCols() != df.NumCols() {
			t.Errorf("Column count mismatch: expected %d, got %d", df.NumCols(), readDF.NumCols())
		}
	})
}

// TestPhase1_GroupByAggregation tests aggregation functionality
func TestPhase1_GroupByAggregation(t *testing.T) {
	// Create test data for aggregation
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "category", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	categoryBuilder := array.NewStringBuilder(pool)
	categoryBuilder.AppendValues([]string{"A", "B", "A", "B", "A"}, nil)
	categoryArray := categoryBuilder.NewArray()
	defer categoryArray.Release()

	valueBuilder := array.NewFloat64Builder(pool)
	valueBuilder.AppendValues([]float64{10.0, 20.0, 30.0, 40.0, 50.0}, nil)
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	record := array.NewRecord(schema, []arrow.Array{categoryArray, valueArray}, 5)
	defer record.Release()

	df := NewDataFrame(record)
	defer df.Release()

	// Test GroupBy/Agg: GroupBy("country").Agg(Sum("revenue"))
	grouped := df.GroupBy("category").Agg(Sum("value"))
	defer grouped.Release()

	if grouped.Err() != nil {
		t.Fatalf("GroupBy Sum failed: %v", grouped.Err())
	}

	// Should have 2 groups (A and B)
	if grouped.NumRows() != 2 {
		t.Errorf("Expected 2 groups, got %d", grouped.NumRows())
	}

	// Should have category + sum columns
	if grouped.NumCols() != 2 {
		t.Errorf("Expected 2 columns (category + value_sum), got %d", grouped.NumCols())
	}

	// Test multiple aggregations
	multiAgg := df.GroupBy("category").Agg(Sum("value"), Mean("value").As("avg_value"))
	defer multiAgg.Release()

	if multiAgg.Err() != nil {
		t.Fatalf("GroupBy Agg failed: %v", multiAgg.Err())
	}

	// Should have category + 2 aggregation columns
	if multiAgg.NumCols() != 3 {
		t.Errorf("Expected 3 columns (category + 2 aggs), got %d", multiAgg.NumCols())
	}
}
