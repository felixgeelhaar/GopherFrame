// Package gopherframe implements performance comparison benchmarks against other Go data libraries
// This demonstrates GopherFrame's performance advantages over existing solutions
package gopherframe

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// Benchmark data sizes for comprehensive testing
var benchmarkSizes = []int{1000, 10000, 100000}

// Test data generation for consistent comparison
func generateTestDataCSV(filename string, rows int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write CSV header
	f.WriteString("id,name,age,salary,department,active\n")

	// Generate consistent test data
	rand.Seed(42) // Fixed seed for reproducible results
	departments := []string{"Engineering", "Sales", "Marketing", "HR", "Operations"}

	for i := 0; i < rows; i++ {
		f.WriteString(fmt.Sprintf("%d,Employee_%d,%d,%.2f,%s,%t\n",
			i,
			i,
			20+rand.Intn(45),            // age 20-65
			30000+rand.Float64()*120000, // salary 30k-150k
			departments[rand.Intn(len(departments))],
			rand.Float64() > 0.1, // 90% active
		))
	}

	return nil
}

// GopherFrame benchmarks
func BenchmarkGopherFrameRead(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		b.Run(fmt.Sprintf("GopherFrame_Read_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				df, err := ReadCSV(filename)
				if err != nil {
					b.Fatalf("Failed to read CSV: %v", err)
				}
				df.Release()
			}
		})
	}
}

func BenchmarkGopherFrameFilter(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		df, err := ReadCSV(filename)
		if err != nil {
			b.Fatalf("Failed to read CSV: %v", err)
		}
		defer df.Release()

		b.Run(fmt.Sprintf("GopherFrame_Filter_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				filtered := df.Filter(Col("age").Gt(Lit(30)))
				filtered.Release()
			}
		})
	}
}

func BenchmarkGopherFrameGroupBy(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		df, err := ReadCSV(filename)
		if err != nil {
			b.Fatalf("Failed to read CSV: %v", err)
		}
		defer df.Release()

		b.Run(fmt.Sprintf("GopherFrame_GroupBy_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				grouped := df.GroupBy("department").Agg(Mean("salary"))
				grouped.Release()
			}
		})
	}
}

func BenchmarkGopherFrameSort(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		df, err := ReadCSV(filename)
		if err != nil {
			b.Fatalf("Failed to read CSV: %v", err)
		}
		defer df.Release()

		b.Run(fmt.Sprintf("GopherFrame_Sort_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sorted := df.Sort("salary", false)
				sorted.Release()
			}
		})
	}
}

// Gota benchmarks for comparison
func BenchmarkGotaRead(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		b.Run(fmt.Sprintf("Gota_Read_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f, err := os.Open(filename)
				if err != nil {
					b.Fatalf("Failed to open file: %v", err)
				}
				df := dataframe.ReadCSV(f)
				f.Close()
				_ = df // Use the dataframe
			}
		})
	}
}

func BenchmarkGotaFilter(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		f, err := os.Open(filename)
		if err != nil {
			b.Fatalf("Failed to open file: %v", err)
		}
		df := dataframe.ReadCSV(f)
		f.Close()

		b.Run(fmt.Sprintf("Gota_Filter_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				filtered := df.Filter(dataframe.F{Colname: "age", Comparator: series.Greater, Comparando: 30})
				_ = filtered
			}
		})
	}
}

// GroupBy benchmark for Gota (simplified - Gota doesn't have native GroupBy)
func BenchmarkGotaGroupBy(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		f, err := os.Open(filename)
		if err != nil {
			b.Fatalf("Failed to open file: %v", err)
		}
		df := dataframe.ReadCSV(f)
		f.Close()

		b.Run(fmt.Sprintf("Gota_GroupBy_Simulation_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Gota doesn't have native GroupBy, so we simulate with filtering
				// This is not a fair comparison but shows the need for native operations
				departments := []string{"Engineering", "Sales", "Marketing", "HR", "Operations"}
				results := make(map[string]float64)

				for _, dept := range departments {
					filtered := df.Filter(dataframe.F{Colname: "department", Comparator: series.Eq, Comparando: dept})
					if filtered.Nrow() > 0 {
						salaryCol := filtered.Col("salary")
						// Calculate mean manually
						sum := 0.0
						count := 0
						for i := 0; i < salaryCol.Len(); i++ {
							if val := salaryCol.Elem(i); val.Type() == series.Float {
								sum += val.Float()
								count++
							}
						}
						if count > 0 {
							results[dept] = sum / float64(count)
						}
					}
				}
				_ = results
			}
		})
	}
}

func BenchmarkGotaSort(b *testing.B) {
	for _, size := range benchmarkSizes {
		filename := fmt.Sprintf("benchmark_data_%d.csv", size)
		if err := generateTestDataCSV(filename, size); err != nil {
			b.Fatalf("Failed to generate test data: %v", err)
		}
		defer os.Remove(filename)

		f, err := os.Open(filename)
		if err != nil {
			b.Fatalf("Failed to open file: %v", err)
		}
		df := dataframe.ReadCSV(f)
		f.Close()

		b.Run(fmt.Sprintf("Gota_Sort_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sorted := df.Arrange(dataframe.RevSort("salary"))
				_ = sorted
			}
		})
	}
}

// Memory usage comparison
func BenchmarkGopherFrameMemory(b *testing.B) {
	size := 10000
	filename := fmt.Sprintf("benchmark_memory_%d.csv", size)
	if err := generateTestDataCSV(filename, size); err != nil {
		b.Fatalf("Failed to generate test data: %v", err)
	}
	defer os.Remove(filename)

	b.Run("GopherFrame_Memory_Usage", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			df, _ := ReadCSV(filename)
			filtered := df.Filter(Col("age").Gt(Lit(30)))
			grouped := filtered.GroupBy("department").Agg(Mean("salary"))
			sorted := grouped.Sort("salary", false)

			// Clean up
			sorted.Release()
			grouped.Release()
			filtered.Release()
			df.Release()
		}
	})
}

func BenchmarkGotaMemory(b *testing.B) {
	size := 10000
	filename := fmt.Sprintf("benchmark_memory_%d.csv", size)
	if err := generateTestDataCSV(filename, size); err != nil {
		b.Fatalf("Failed to generate test data: %v", err)
	}
	defer os.Remove(filename)

	b.Run("Gota_Memory_Usage", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			f, _ := os.Open(filename)
			df := dataframe.ReadCSV(f)
			f.Close()

			filtered := df.Filter(dataframe.F{Colname: "age", Comparator: series.Greater, Comparando: 30})
			sorted := filtered.Arrange(dataframe.RevSort("salary"))
			_ = sorted
		}
	})
}

// Comparative analysis function
func BenchmarkComparativeAnalysis(b *testing.B) {
	b.Run("Analysis", func(b *testing.B) {
		b.Skip("Run individual benchmarks to see comparison results")
	})
}

// Example usage comparison for documentation
func ExamplePerformanceComparison() {
	// GopherFrame approach
	df, _ := ReadCSV("data.csv")
	defer df.Release()

	result := df.
		Filter(Col("age").Gt(Lit(25))).
		GroupBy("department").
		Agg(Mean("salary")).
		Sort("salary", false)
	defer result.Release()

	fmt.Printf("GopherFrame result: %d rows\n", result.NumRows())

	// Gota approach (more verbose, no native GroupBy)
	f, _ := os.Open("data.csv")
	gotaDF := dataframe.ReadCSV(f)
	f.Close()

	filtered := gotaDF.Filter(dataframe.F{Colname: "age", Comparator: series.Greater, Comparando: 25})
	sorted := filtered.Arrange(dataframe.RevSort("salary"))

	fmt.Printf("Gota result: %d rows\n", sorted.Nrow())
}

// Performance test runner for automated testing
func TestPerformanceComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance comparison in short mode")
	}

	// Create test data
	filename := "perf_test_data.csv"
	if err := generateTestDataCSV(filename, 1000); err != nil {
		t.Fatalf("Failed to generate test data: %v", err)
	}
	defer os.Remove(filename)

	// Test GopherFrame
	start := time.Now()
	df, err := ReadCSV(filename)
	if err != nil {
		t.Fatalf("GopherFrame failed to read CSV: %v", err)
	}

	result := df.Filter(Col("age").Gt(Lit(30))).GroupBy("department").Agg(Mean("salary"))
	gopherFrameTime := time.Since(start)

	result.Release()
	df.Release()

	t.Logf("GopherFrame processing time: %v", gopherFrameTime)

	// Test Gota
	start = time.Now()
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open file for Gota: %v", err)
	}
	gotaDF := dataframe.ReadCSV(f)
	f.Close()

	filtered := gotaDF.Filter(dataframe.F{Colname: "age", Comparator: series.Greater, Comparando: 30})
	_ = filtered
	gotaTime := time.Since(start)

	t.Logf("Gota processing time: %v", gotaTime)

	// Performance comparison
	if gopherFrameTime < gotaTime {
		t.Logf("✅ GopherFrame is %.2fx faster than Gota", float64(gotaTime)/float64(gopherFrameTime))
	} else {
		t.Logf("⚠️  Gota is %.2fx faster than GopherFrame", float64(gopherFrameTime)/float64(gotaTime))
	}
}
