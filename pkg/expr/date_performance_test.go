package expr

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// BenchmarkDateTimeOperations benchmarks date/time expression performance
func BenchmarkDateTimeOperations(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("Year_Size_"+string(rune(size/1000))+"K", func(b *testing.B) {
			df := createDateBenchmarkDataFrame(size)
			defer df.Release()

			yearExpr := Col("date_col").Year()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := yearExpr.Evaluate(df)
				if err != nil {
					b.Fatalf("Year extraction failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("AddDays_Size_"+string(rune(size/1000))+"K", func(b *testing.B) {
			df := createDateBenchmarkDataFrame(size)
			defer df.Release()

			addDaysExpr := Col("date_col").AddDays(Lit(30))

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := addDaysExpr.Evaluate(df)
				if err != nil {
					b.Fatalf("AddDays failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("AddMonths_Size_"+string(rune(size/1000))+"K", func(b *testing.B) {
			df := createDateBenchmarkDataFrame(size)
			defer df.Release()

			addMonthsExpr := Col("date_col").AddMonths(Lit(3))

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := addMonthsExpr.Evaluate(df)
				if err != nil {
					b.Fatalf("AddMonths failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

func createDateBenchmarkDataFrame(size int) *core.DataFrame {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "date_col", Type: arrow.FixedWidthTypes.Date32},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	dateBuilder := array.NewDate32Builder(pool)

	// Create varied dates spanning multiple years
	baseDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	baseDays := arrow.Date32(baseDate.Unix() / 86400)

	for i := 0; i < size; i++ {
		idBuilder.Append(int64(i))
		// Spread dates across 4 years (1460 days)
		dayOffset := int32(i % 1460)
		dateBuilder.Append(baseDays + arrow.Date32(dayOffset))
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()

	dateArray := dateBuilder.NewArray()
	defer dateArray.Release()

	record := array.NewRecord(schema, []arrow.Array{idArray, dateArray}, int64(size))
	defer record.Release()

	return core.NewDataFrame(record)
}
