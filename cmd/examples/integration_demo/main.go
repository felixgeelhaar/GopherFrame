// Package main demonstrates GopherFrame integration with popular Go tools.
// This example shows how GopherFrame fits into real-world Go applications
// alongside HTTP servers, databases, and observability tools.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
	fmt.Println("=== GopherFrame Integration Demo ===")
	fmt.Println()

	// 1. Simulate API metrics data (in production, this comes from Prometheus/metrics)
	fmt.Println("1. Creating API metrics DataFrame...")
	metricsDF := createMetricsData()
	defer metricsDF.Release()
	fmt.Printf("   Loaded %d metric records\n", metricsDF.NumRows())

	// 2. Real-time analytics pipeline
	fmt.Println("\n2. Running analytics pipeline...")

	// Filter to last hour
	recentMetrics := metricsDF.Filter(
		metricsDF.Col("latency_ms").Gt(gf.Lit(0.0)),
	)
	defer recentMetrics.Release()

	// Aggregate by endpoint
	summary := recentMetrics.GroupBy("endpoint").Agg(
		gf.Mean("latency_ms").As("avg_latency"),
		gf.Max("latency_ms").As("max_latency"),
		gf.Count("latency_ms").As("request_count"),
		gf.Variance("latency_ms").As("latency_variance"),
	)
	defer summary.Release()
	fmt.Printf("   Computed stats for %d endpoints\n", summary.NumRows())

	// 3. Outlier detection for alerting
	fmt.Println("\n3. Running outlier detection...")
	outliers, err := metricsDF.DetectOutliersIQR("latency_ms", 1.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Found %d outlier requests (potential alerts)\n", outliers.Count)
	fmt.Printf("   Normal range: [%.1f, %.1f] ms\n", outliers.LowerBound, outliers.UpperBound)

	// 4. Data quality validation
	fmt.Println("\n4. Validating data quality...")
	validation := metricsDF.Validate(
		gf.NotNull("endpoint"),
		gf.Positive("latency_ms"),
		gf.InRange("status_code", 100, 599),
	)
	if validation.Valid {
		fmt.Println("   All validation rules passed")
	} else {
		fmt.Printf("   Violations: %v\n", validation.Violations)
	}

	// 5. Export results
	fmt.Println("\n5. Export capabilities:")
	fmt.Println("   - WriteParquet(summary, 'metrics.parquet')")
	fmt.Println("   - WriteJSON(summary, 'metrics.json')")
	fmt.Println("   - WriteSQL(summary, db, 'metrics_summary')")
	fmt.Println("   - WritePartitioned(metricsDF, 'data/', []string{'endpoint'})")

	// 6. Descriptive statistics
	fmt.Println("\n6. Data profile:")
	fmt.Println(metricsDF.DescribeString())

	fmt.Println("=== Integration Demo Complete ===")
}

func createMetricsData() *gf.DataFrame {
	pool := memory.NewGoAllocator()

	endpoints := []string{"/api/users", "/api/orders", "/api/products", "/api/users", "/api/orders",
		"/api/products", "/api/users", "/api/orders", "/api/products", "/api/users"}
	latencies := []float64{12.5, 45.2, 8.1, 15.3, 52.8, 7.9, 250.0, 38.1, 9.2, 11.8}
	statusCodes := []float64{200, 200, 200, 200, 500, 200, 200, 200, 200, 200}

	_ = time.Now() // Would use for timestamp column in production

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "endpoint", Type: arrow.BinaryTypes.String},
		{Name: "latency_ms", Type: arrow.PrimitiveTypes.Float64},
		{Name: "status_code", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	eb := array.NewStringBuilder(pool)
	eb.AppendValues(endpoints, nil)
	ea := eb.NewArray()
	eb.Release()

	lb := array.NewFloat64Builder(pool)
	lb.AppendValues(latencies, nil)
	la := lb.NewArray()
	lb.Release()

	sb := array.NewFloat64Builder(pool)
	sb.AppendValues(statusCodes, nil)
	sa := sb.NewArray()
	sb.Release()

	record := array.NewRecord(schema, []arrow.Array{ea, la, sa}, int64(len(endpoints)))
	return gf.NewDataFrame(record)
}
