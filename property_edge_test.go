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

// TestPropertyNullHandling verifies null value handling across operations
func TestPropertyNullHandling(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Operations handle nulls correctly", prop.ForAll(
		func(data []testRowWithNulls) bool {
			if len(data) == 0 {
				return true
			}

			df := createDataFrameWithNulls(data)
			defer df.Release()

			// Filter should skip null values
			filtered := df.Filter(Col("value").Gt(Lit(50.0)))
			defer filtered.Release()

			if filtered.Err() != nil {
				return false
			}

			// Count non-null values in result
			series, err := filtered.coreDF.Column("value")
			if err != nil {
				return false
			}
			defer series.Release()

			nullCount := 0
			for i := 0; i < series.Len(); i++ {
				if series.IsNull(i) {
					nullCount++
				}
			}

			// All nulls should be filtered out in comparison operations
			return nullCount == 0
		},
		gen.SliceOf(genTestRowWithNulls()),
	))

	properties.TestingRun(t)
}

// TestPropertyEmptyDataFrame verifies operations on empty DataFrames
func TestPropertyEmptyDataFrame(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Empty DataFrame operations don't panic", prop.ForAll(
		func(threshold float64) bool {
			// Create empty DataFrame
			df := createDataFrameFromTestRows([]testRow{})
			defer df.Release()

			// All operations should work without panic
			filtered := df.Filter(Col("value").Gt(Lit(threshold)))
			defer filtered.Release()

			selected := df.Select("id")
			defer selected.Release()

			withCol := df.WithColumn("new", Lit(1.0))
			defer withCol.Release()

			grouped := df.GroupBy("category").Agg(Sum("value"))
			defer grouped.Release()

			// All should succeed with empty results
			return filtered.Err() == nil &&
				selected.Err() == nil &&
				withCol.Err() == nil &&
				grouped.Err() == nil &&
				filtered.NumRows() == 0
		},
		gen.Float64(),
	))

	properties.TestingRun(t)
}

// TestLargeValues verifies handling of extreme values
func TestLargeValues(t *testing.T) {
	testCases := []struct {
		name   string
		values []float64
	}{
		{"Zero values", []float64{0.0, 0.0, 0.0}},
		{"Small values", []float64{0.1, -0.1, 0.001}},
		{"Large values", []float64{1000.0, -1000.0, 10000.0}},
		{"Mixed values", []float64{0.0, 1.0, -1.0, 100.0, -100.0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := make([]testRow, len(tc.values))
			for i, v := range tc.values {
				rows[i] = testRow{
					ID:       int64(i),
					Name:     "test",
					Value:    v,
					Category: "A",
				}
			}

			df := createDataFrameFromTestRows(rows)
			defer df.Release()

			// Operations should handle reasonable arithmetic
			result := df.
				WithColumn("doubled", Col("value").Mul(Lit(2.0))).
				WithColumn("plus_one", Col("value").Add(Lit(1.0)))
			defer result.Release()

			if result.Err() != nil {
				t.Errorf("Operations failed: %v", result.Err())
			}

			if result.NumRows() != df.NumRows() {
				t.Errorf("Row count mismatch: expected %d, got %d", df.NumRows(), result.NumRows())
			}

			if result.NumCols() != df.NumCols()+2 {
				t.Errorf("Column count mismatch: expected %d, got %d", df.NumCols()+2, result.NumCols())
			}
		})
	}
}

// TestPropertyStringOperations verifies string column handling
func TestPropertyStringOperations(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("String columns handle special characters", prop.ForAll(
		func(strings []string) bool {
			if len(strings) == 0 {
				return true
			}

			// Create DataFrame with various strings
			rows := make([]testRow, len(strings))
			for i, s := range strings {
				rows[i] = testRow{
					ID:       int64(i),
					Name:     s,
					Value:    float64(i),
					Category: "A",
				}
			}

			df := createDataFrameFromTestRows(rows)
			defer df.Release()

			// CSV round-trip should preserve strings
			tempFile := "/tmp/test_strings.csv"
			err := WriteCSV(df, tempFile)
			if err != nil {
				return false
			}

			readDf, err := ReadCSV(tempFile)
			if err != nil {
				return false
			}
			defer readDf.Release()

			return df.NumRows() == readDf.NumRows()
		},
		gen.SliceOf(gen.OneConstOf(
			"", "normal", "with spaces", "with,comma",
			"with\"quote", "with\nnewline", "with\ttab",
			"unicode😀", "very long string with many characters",
		)),
	))

	properties.TestingRun(t)
}

// TestPropertyChainedOperationsAssociativity verifies operation associativity
func TestPropertyChainedOperationsAssociativity(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("Chained operations are associative", prop.ForAll(
		func(data []testRow) bool {
			if len(data) < 10 {
				return true // Need enough data for meaningful test
			}

			df := createDataFrameFromTestRows(data)
			defer df.Release()

			// Method 1: Filter then select
			result1 := df.
				Filter(Col("value").Gt(Lit(30.0))).
				Select("id", "value")
			defer result1.Release()

			// Method 2: Different operation order (where possible)
			result2 := df.
				Select("id", "value", "category").
				Filter(Col("value").Gt(Lit(30.0))).
				Select("id", "value")
			defer result2.Release()

			if result1.Err() != nil || result2.Err() != nil {
				return false
			}

			// Should produce same result
			return result1.NumRows() == result2.NumRows() &&
				result1.NumCols() == result2.NumCols()
		},
		gen.SliceOf(genTestRow()),
	))

	properties.TestingRun(t)
}

// Helper generators for edge cases

type testRowWithNulls struct {
	ID       int64
	Name     *string
	Value    *float64
	Category string
}

func genTestRowWithNulls() gopter.Gen {
	return gopter.CombineGens(
		gen.Int64Range(1, 1000),
		gen.OneGenOf(gen.PtrOf(gen.AlphaString()), gen.Const((*string)(nil))),
		gen.OneGenOf(gen.PtrOf(gen.Float64Range(0, 100)), gen.Const((*float64)(nil))),
		gen.OneConstOf("A", "B", "C"),
	).Map(func(values []interface{}) testRowWithNulls {
		var name *string
		var value *float64

		// Handle potential nil values
		if v, ok := values[1].(*string); ok {
			name = v
		}
		if v, ok := values[2].(*float64); ok {
			value = v
		}

		return testRowWithNulls{
			ID:       values[0].(int64),
			Name:     name,
			Value:    value,
			Category: values[3].(string),
		}
	})
}

func createDataFrameWithNulls(rows []testRowWithNulls) *DataFrame {
	pool := memory.NewGoAllocator()

	idBuilder := array.NewInt64Builder(pool)
	nameBuilder := array.NewStringBuilder(pool)
	valueBuilder := array.NewFloat64Builder(pool)
	categoryBuilder := array.NewStringBuilder(pool)

	for _, row := range rows {
		idBuilder.Append(row.ID)

		if row.Name != nil {
			nameBuilder.Append(*row.Name)
		} else {
			nameBuilder.AppendNull()
		}

		if row.Value != nil {
			valueBuilder.Append(*row.Value)
		} else {
			valueBuilder.AppendNull()
		}

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
