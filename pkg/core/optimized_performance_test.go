package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow/memory"
)

// BenchmarkOptimizedVsStandard compares optimized operations with standard implementations
func BenchmarkOptimizedVsStandard(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		// Filter benchmarks
		b.Run("Filter_Standard_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			mask := createBooleanMask(size, 0.5)
			defer mask.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Filter(mask)
				if err != nil {
					b.Fatalf("Filter failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Filter_Optimized_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			mask := createBooleanMask(size, 0.5)
			defer mask.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterOptimized(mask)
				if err != nil {
					b.Fatalf("FilterOptimized failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Filter_Vectorized_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			mask := createBooleanMask(size, 0.5)
			defer mask.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterVectorized(mask)
				if err != nil {
					b.Fatalf("FilterVectorized failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Filter_OptimizedPool_Size_"+formatSize(size), func(b *testing.B) {
			pool := memory.NewGoAllocator()
			df := createBenchmarkDataFrameWithPool(size, pool)
			defer df.Release()

			mask := createBooleanMaskWithPool(size, 0.5, pool)
			defer mask.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterOptimizedWithPool(mask, pool)
				if err != nil {
					b.Fatalf("FilterOptimizedWithPool failed: %v", err)
				}
				result.Release()
			}
		})

		// Select benchmarks
		b.Run("Select_Standard_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			columns := []string{"id", "value"}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Select(columns)
				if err != nil {
					b.Fatalf("Select failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Select_Optimized_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			columns := []string{"id", "value"}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.SelectOptimized(columns)
				if err != nil {
					b.Fatalf("SelectOptimized failed: %v", err)
				}
				result.Release()
			}
		})

		// WithColumn benchmarks
		b.Run("WithColumn_Standard_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			newColumn := createComputedColumn(size)
			defer newColumn.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.WithColumn("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumn failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("WithColumn_Optimized_Size_"+formatSize(size), func(b *testing.B) {
			df := createBenchmarkDataFrame(size)
			defer df.Release()

			newColumn := createComputedColumn(size)
			defer newColumn.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.WithColumnOptimized("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumnOptimized failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("WithColumn_OptimizedPool_Size_"+formatSize(size), func(b *testing.B) {
			pool := memory.NewGoAllocator()
			df := createBenchmarkDataFrameWithPool(size, pool)
			defer df.Release()

			newColumn := createComputedColumn(size)
			defer newColumn.Release()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.WithColumnOptimizedWithPool("computed", newColumn, pool)
				if err != nil {
					b.Fatalf("WithColumnOptimizedWithPool failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

// BenchmarkOptimizedFilterTypes tests optimized filtering across different data types
func BenchmarkOptimizedFilterTypes(b *testing.B) {
	size := 50000

	testCases := []struct {
		name     string
		dfFunc   func(int) *DataFrame
		dataType string
	}{
		{"Int64", createInt64DataFrame, "INT64"},
		{"Float64", createFloat64DataFrame, "FLOAT64"},
		{"String", createStringDataFrame, "STRING"},
		{"Mixed", createMixedTypeDataFrameNoBoolean, "MIXED"},
	}

	for _, tc := range testCases {
		df := tc.dfFunc(size)
		defer df.Release()

		mask := createBooleanMask(size, 0.3)
		defer mask.Release()

		b.Run("Standard_"+tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Filter(mask)
				if err != nil {
					b.Fatalf("Filter failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Optimized_"+tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterOptimized(mask)
				if err != nil {
					b.Fatalf("FilterOptimized failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Vectorized_"+tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterVectorized(mask)
				if err != nil {
					b.Fatalf("FilterVectorized failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

// BenchmarkSelectOptimizedColumns tests optimized column selection performance
func BenchmarkSelectOptimizedColumns(b *testing.B) {
	size := 50000
	df := createWideDataFrame(size, 20) // 20 columns
	defer df.Release()

	testCases := []struct {
		name    string
		columns []string
	}{
		{"1Col", []string{"col_00"}},
		{"5Cols", []string{"col_00", "col_01", "col_02", "col_03", "col_04"}},
		{"10Cols", []string{"col_00", "col_01", "col_02", "col_03", "col_04", "col_05", "col_06", "col_07", "col_08", "col_09"}},
		{"15Cols", []string{"col_00", "col_01", "col_02", "col_03", "col_04", "col_05", "col_06", "col_07", "col_08", "col_09", "col_10", "col_11", "col_12", "col_13", "col_14"}},
	}

	for _, tc := range testCases {
		b.Run("Standard_"+tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Select(tc.columns)
				if err != nil {
					b.Fatalf("Select failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Optimized_"+tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.SelectOptimized(tc.columns)
				if err != nil {
					b.Fatalf("SelectOptimized failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

// BenchmarkChainedOptimizedOperations tests performance of chained optimized operations
func BenchmarkChainedOptimizedOperations(b *testing.B) {
	sizes := []int{10000, 50000}

	for _, size := range sizes {
		df := createBenchmarkDataFrame(size)
		defer df.Release()

		mask := createBooleanMask(size, 0.3)
		defer mask.Release()

		newColumn := createComputedColumn(size)
		defer newColumn.Release()

		b.Run("Standard_Chain_Size_"+formatSize(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Chain: WithColumn -> Filter -> Select
				withCol, err := df.WithColumn("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumn failed: %v", err)
				}

				filtered, err := withCol.Filter(mask)
				withCol.Release()
				if err != nil {
					b.Fatalf("Filter failed: %v", err)
				}

				selected, err := filtered.Select([]string{"id", "computed"})
				filtered.Release()
				if err != nil {
					b.Fatalf("Select failed: %v", err)
				}

				selected.Release()
			}
		})

		b.Run("Optimized_Chain_Size_"+formatSize(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Chain: WithColumnOptimized -> FilterOptimized -> SelectOptimized
				withCol, err := df.WithColumnOptimized("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumnOptimized failed: %v", err)
				}

				filtered, err := withCol.FilterOptimized(mask)
				withCol.Release()
				if err != nil {
					b.Fatalf("FilterOptimized failed: %v", err)
				}

				selected, err := filtered.SelectOptimized([]string{"id", "computed"})
				filtered.Release()
				if err != nil {
					b.Fatalf("SelectOptimized failed: %v", err)
				}

				selected.Release()
			}
		})

		b.Run("Vectorized_Chain_Size_"+formatSize(size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Chain: WithColumnOptimized -> FilterVectorized -> SelectOptimized
				withCol, err := df.WithColumnOptimized("computed", newColumn)
				if err != nil {
					b.Fatalf("WithColumnOptimized failed: %v", err)
				}

				filtered, err := withCol.FilterVectorized(mask)
				withCol.Release()
				if err != nil {
					b.Fatalf("FilterVectorized failed: %v", err)
				}

				selected, err := filtered.SelectOptimized([]string{"id", "computed"})
				filtered.Release()
				if err != nil {
					b.Fatalf("SelectOptimized failed: %v", err)
				}

				selected.Release()
			}
		})

		// Pool-optimized chain
		b.Run("PoolOptimized_Chain_Size_"+formatSize(size), func(b *testing.B) {
			pool := memory.NewGoAllocator()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Chain with shared pool
				withCol, err := df.WithColumnOptimizedWithPool("computed", newColumn, pool)
				if err != nil {
					b.Fatalf("WithColumnOptimizedWithPool failed: %v", err)
				}

				filtered, err := withCol.FilterOptimizedWithPool(mask, pool)
				withCol.Release()
				if err != nil {
					b.Fatalf("FilterOptimizedWithPool failed: %v", err)
				}

				selected, err := filtered.SelectOptimized([]string{"id", "computed"})
				filtered.Release()
				if err != nil {
					b.Fatalf("SelectOptimized failed: %v", err)
				}

				selected.Release()
			}
		})
	}
}

// BenchmarkFilterSelectivity tests performance across different filter selectivities
func BenchmarkFilterSelectivity(b *testing.B) {
	size := 25000
	selectivities := []float64{0.01, 0.1, 0.3, 0.5, 0.8, 0.99}

	df := createBenchmarkDataFrame(size)
	defer df.Release()

	for _, selectivity := range selectivities {
		mask := createBooleanMask(size, selectivity)
		defer mask.Release()

		selectivityStr := formatFloat(selectivity)

		b.Run("Standard_Selectivity_"+selectivityStr, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.Filter(mask)
				if err != nil {
					b.Fatalf("Filter failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Optimized_Selectivity_"+selectivityStr, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterOptimized(mask)
				if err != nil {
					b.Fatalf("FilterOptimized failed: %v", err)
				}
				result.Release()
			}
		})

		b.Run("Vectorized_Selectivity_"+selectivityStr, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result, err := df.FilterVectorized(mask)
				if err != nil {
					b.Fatalf("FilterVectorized failed: %v", err)
				}
				result.Release()
			}
		})
	}
}

func formatFloat(f float64) string {
	if f == 0.01 {
		return "001"
	} else if f == 0.1 {
		return "010"
	} else if f == 0.3 {
		return "030"
	} else if f == 0.5 {
		return "050"
	} else if f == 0.8 {
		return "080"
	} else if f == 0.99 {
		return "099"
	}
	return "unknown"
}
