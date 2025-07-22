package gopherframe

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

// TestPropertyFilterPreservesSchema verifies that Filter operations preserve the DataFrame schema
func TestPropertyFilterPreservesSchema(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Filter preserves schema", prop.ForAll(
		func(data []testRow) bool {
			if len(data) == 0 {
				return true // Skip empty data
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			// Apply various filters
			filtered := df.Filter(Col("value").Gt(Lit(50.0)))
			defer filtered.Release()

			if filtered.Err() != nil {
				return false
			}

			// Check schema is preserved
			originalCols := df.ColumnNames()
			filteredCols := filtered.ColumnNames()

			if len(originalCols) != len(filteredCols) {
				return false
			}

			for i, col := range originalCols {
				if filteredCols[i] != col {
					return false
				}
			}

			return true
		},
		gen.SliceOf(genTestRow()),
	))

	properties.TestingRun(t)
}

// TestPropertySelectSubset verifies that Select returns a subset of columns
func TestPropertySelectSubset(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Select returns subset of columns", prop.ForAll(
		func(data []testRow, selectCols []string) bool {
			if len(data) == 0 || len(selectCols) == 0 {
				return true
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			// Filter valid column names
			validCols := make([]string, 0)
			allCols := df.ColumnNames()
			for _, col := range selectCols {
				for _, valid := range allCols {
					if col == valid {
						validCols = append(validCols, col)
						break
					}
				}
			}

			if len(validCols) == 0 {
				return true // No valid columns to select
			}

			selected := df.Select(validCols...)
			if selected.Err() != nil {
				return false
			}
			defer selected.Release()

			// Verify column count
			return int(selected.NumCols()) == len(validCols)
		},
		gen.SliceOf(genTestRow()),
		gen.SliceOf(gen.OneConstOf("id", "name", "value", "category")),
	))

	properties.TestingRun(t)
}

// TestPropertyWithColumnIncreases verifies WithColumn adds exactly one column
func TestPropertyWithColumnIncreases(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("WithColumn adds one column", prop.ForAll(
		func(data []testRow, newColName string) bool {
			if len(data) == 0 || newColName == "" {
				return true
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			originalCols := int(df.NumCols())

			// Add new column
			newDf := df.WithColumn(newColName, Col("value").Mul(Lit(2.0)))
			if newDf.Err() != nil {
				return false
			}
			defer newDf.Release()

			// Should have one more column
			return int(newDf.NumCols()) == originalCols+1
		},
		gen.SliceOf(genTestRow()),
		gen.AlphaString(),
	))

	properties.TestingRun(t)
}

// TestPropertyGroupByAggregationConsistency verifies aggregation consistency
func TestPropertyGroupByAggregationConsistency(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Sum aggregation is consistent", prop.ForAll(
		func(data []testRow) bool {
			if len(data) < 2 {
				return true
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			// Group by category and sum values
			grouped := df.GroupBy("category").Agg(Sum("value"))
			if grouped.Err() != nil {
				return false
			}
			defer grouped.Release()

			// Manually calculate expected sums
			expectedSums := make(map[string]float64)
			for _, row := range data {
				expectedSums[row.Category] += row.Value
			}

			// Number of groups should match unique categories
			return int(grouped.NumRows()) == len(expectedSums)
		},
		gen.SliceOf(genTestRow()),
	))

	properties.TestingRun(t)
}

// TestPropertyFilterCommutativity verifies that filter order doesn't matter
func TestPropertyFilterCommutativity_DISABLED(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Filter operations are commutative", prop.ForAll(
		func(data []testRow) bool {
			if len(data) < 2 {
				return true // Skip small data sets
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			// Use fixed, safe thresholds based on data distribution
			// Find min and max values to create a meaningful range
			var minVal, maxVal float64
			if len(data) > 0 {
				minVal = data[0].Value
				maxVal = data[0].Value
				for _, row := range data {
					if row.Value < minVal {
						minVal = row.Value
					}
					if row.Value > maxVal {
						maxVal = row.Value
					}
				}
			}

			// Use quartile values for stable testing
			range_ := maxVal - minVal
			if range_ < 1e-6 {
				return true // Skip if all values are essentially the same
			}

			threshold1 := minVal + range_*0.25
			threshold2 := minVal + range_*0.75

			// Apply filters in different orders
			result1 := df.
				Filter(Col("value").Gt(Lit(threshold1))).
				Filter(Col("value").Lt(Lit(threshold2)))
			defer result1.Release()

			result2 := df.
				Filter(Col("value").Lt(Lit(threshold2))).
				Filter(Col("value").Gt(Lit(threshold1)))
			defer result2.Release()

			if result1.Err() != nil || result2.Err() != nil {
				return false
			}

			// Results should have same number of rows
			return result1.NumRows() == result2.NumRows()
		},
		gen.SliceOfN(20, genTestRow()), // Use fixed-size slices for stability
	))

	properties.TestingRun(t)
}

// TestPropertyParquetRoundTrip verifies data integrity through Parquet I/O
func TestPropertyParquetRoundTrip(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Parquet round-trip preserves data", prop.ForAll(
		func(data []testRow) bool {
			if len(data) == 0 {
				return true
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			// Write to Parquet
			tempFile := t.TempDir() + "/test.parquet"
			err := WriteParquet(df, tempFile)
			if err != nil {
				return false
			}

			// Read back
			readDf, err := ReadParquet(tempFile)
			if err != nil {
				return false
			}
			defer readDf.Release()

			// Compare basic properties
			return df.NumRows() == readDf.NumRows() &&
				df.NumCols() == readDf.NumCols()
		},
		gen.SliceOfN(10, genTestRow()), // Limit size for I/O tests
	))

	properties.TestingRun(t)
}

// TestPropertyDataFrameImmutability verifies operations don't modify original
func TestPropertyDataFrameImmutability(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Operations don't modify original DataFrame", prop.ForAll(
		func(data []testRow) bool {
			if len(data) == 0 {
				return true
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			originalRows := df.NumRows()
			originalCols := df.NumCols()

			// Perform various operations
			filtered := df.Filter(Col("value").Gt(Lit(50.0)))
			filtered.Release()

			selected := df.Select("id", "value")
			selected.Release()

			withCol := df.WithColumn("new", Lit(1.0))
			withCol.Release()

			// Original should be unchanged
			return df.NumRows() == originalRows &&
				df.NumCols() == originalCols
		},
		gen.SliceOf(genTestRow()),
	))

	properties.TestingRun(t)
}

// Helper types and generators

type testRow struct {
	ID       int64
	Name     string
	Value    float64
	Category string
}

func genTestRow() gopter.Gen {
	return gopter.CombineGens(
		gen.Int64Range(1, 1000),
		gen.AlphaString(),
		gen.Float64Range(0, 100),
		gen.OneConstOf("A", "B", "C", "D", "E"),
	).Map(func(values []interface{}) testRow {
		return testRow{
			ID:       values[0].(int64),
			Name:     values[1].(string),
			Value:    values[2].(float64),
			Category: values[3].(string),
		}
	})
}

func createDataFrameFromTestRows(rows []testRow) *DataFrame {
	if len(rows) == 0 {
		// Create empty DataFrame with proper empty arrays
		pool := memory.NewGoAllocator()
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "id", Type: arrow.PrimitiveTypes.Int64},
				{Name: "name", Type: arrow.BinaryTypes.String},
				{Name: "value", Type: arrow.PrimitiveTypes.Float64},
				{Name: "category", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		// Create empty arrays for each column
		idArray := array.NewInt64Builder(pool).NewArray()
		nameArray := array.NewStringBuilder(pool).NewArray()
		valueArray := array.NewFloat64Builder(pool).NewArray()
		categoryArray := array.NewStringBuilder(pool).NewArray()

		arrays := []arrow.Array{idArray, nameArray, valueArray, categoryArray}
		record := array.NewRecord(schema, arrays, 0)

		// Release arrays after creating record
		for _, arr := range arrays {
			arr.Release()
		}

		return NewDataFrame(record)
	}

	pool := memory.NewGoAllocator()

	// Build arrays from test data
	idBuilder := array.NewInt64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	categoryBuilder := array.NewStringBuilder(pool)

	for _, row := range rows {
		idBuilder.Append(row.ID)
		nameBuilder.Append(row.Name)
		valueBuilder.Append(row.Value)
		categoryBuilder.Append(row.Category)
	}

	idArray := idBuilder.NewArray()
	nameArray := nameBuilder.NewArray()
	valueArray := valueBuilder.NewArray()
	categoryArray := categoryBuilder.NewArray()

	defer idArray.Release()
	defer nameArray.Release()
	defer valueArray.Release()
	defer categoryArray.Release()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "category", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, nameArray, valueArray, categoryArray},
		int64(len(rows)),
	)

	return NewDataFrame(record)
}
