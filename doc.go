// Package gopherframe provides a production-ready DataFrame library for Go.
//
// GopherFrame is built on Apache Arrow for zero-copy operations and delivers
// 2-428x better performance than existing Go alternatives while maintaining
// type safety and idiomatic Go design.
//
// # Installation
//
//	go get github.com/felixgeelhaar/GopherFrame
//
// # Quick Start
//
// Reading and transforming data:
//
//	import gf "github.com/felixgeelhaar/GopherFrame"
//
//	// Read from Parquet
//	df, err := gf.ReadParquet("sales.parquet")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer df.Release()
//
//	// Filter and transform
//	result := df.Filter(df.Col("amount").Gt(gf.Lit(0.0)))
//	defer result.Release()
//
//	withTax := result.WithColumn("tax", result.Col("amount").Mul(gf.Lit(0.08)))
//	defer withTax.Release()
//
// # Core Features
//
// DataFrame/Series Operations:
//   - Zero-copy column selection with Apache Arrow backend
//   - Immutable, strongly-typed data structures
//   - Type-safe expression system for transformations
//
// High-Performance I/O:
//   - Parquet (5-10M rows/sec with Snappy compression)
//   - CSV with header detection and type inference
//   - Arrow IPC for zero-copy serialization
//
// Data Transformations:
//   - Select: Project specific columns
//   - Filter: Row filtering with complex predicates
//   - WithColumn: Computed column creation
//   - Sort: Single and multi-column sorting
//   - GroupBy/Agg: Hash-based aggregation
//
// Join Operations:
//   - InnerJoin: Return only matching rows
//   - LeftJoin: Preserve all left-side rows
//   - RightJoin: Preserve all right-side rows
//   - FullOuterJoin: Preserve all rows from both sides
//   - CrossJoin: Cartesian product
//
// Window Functions:
//   - Analytical: RowNumber, Rank, DenseRank, Lag, Lead
//   - Rolling: RollingSum, RollingMean, RollingMin, RollingMax
//   - Cumulative: CumSum, CumMax, CumMin, CumProd
//
// Statistical Aggregations:
//   - Percentile: Any percentile (0.0-1.0)
//   - Median: 50th percentile
//   - Mode: Most frequent value
//   - Correlation: Pearson correlation coefficient
//
// String Operations:
//   - Case conversion: Upper, Lower
//   - Trimming: Trim, TrimLeft, TrimRight
//   - Pattern matching: Contains, StartsWith, EndsWith, Match (regex)
//   - Metrics: Length
//
// Temporal Operations:
//   - Component extraction: Year, Month, Day, Hour, Minute, Second
//   - Truncation: TruncateToYear, TruncateToMonth, TruncateToDay, TruncateToHour
//   - Arithmetic: AddDays, AddHours, AddMinutes, AddSeconds
//
// # Production Memory Management
//
// GopherFrame provides memory limits and pressure monitoring:
//
//	import "github.com/felixgeelhaar/GopherFrame/pkg/core"
//
//	// Configure 1GB memory limit
//	pool := memory.NewGoAllocator()
//	limited := core.NewLimitedAllocator(pool, 1024*1024*1024)
//
//	// Create DataFrame with limited allocator
//	df := gf.NewDataFrameWithAllocator(record, limited)
//	defer df.Release()
//
//	// Monitor memory pressure
//	switch limited.MemoryPressureLevel() {
//	case "critical":
//	    return errors.New("memory critical")
//	case "high":
//	    batchSize = batchSize / 2
//	}
//
// # Advanced Examples
//
// GroupBy and aggregation:
//
//	summary := df.GroupBy("region").Agg(
//	    gf.Sum("amount").As("total_revenue"),
//	    gf.Mean("amount").As("avg_order_value"),
//	    gf.Count("order_id").As("order_count"),
//	)
//	defer summary.Release()
//
// Join operations:
//
//	result, err := users.InnerJoin(orders, "user_id", "customer_id")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer result.Release()
//
// Window functions for analytics:
//
//	result, err := df.Window().
//	    PartitionBy("region").
//	    OrderBy("date").
//	    Rows(7).
//	    Over(
//	        gf.RollingSum("sales").As("sales_7d"),
//	        gf.CumSum("sales").As("cumulative_sales"),
//	    )
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer result.Release()
//
// Statistical analysis:
//
//	slaMetrics := df.GroupBy("endpoint").Agg(
//	    gf.Percentile("response_time", 0.50).As("p50_ms"),
//	    gf.Percentile("response_time", 0.95).As("p95_ms"),
//	    gf.Percentile("response_time", 0.99).As("p99_ms"),
//	)
//	defer slaMetrics.Release()
//
// # Performance
//
// GopherFrame delivers exceptional performance through Apache Arrow:
//
//   - Select (10K rows): 730ns (65x faster than Gota)
//   - Column access: 132ns (26x faster than Gota)
//   - Iteration (1K rows): 401ns (390x faster than Gota)
//   - Memory efficiency: 2-200x less memory usage
//
// # Best Practices
//
// Always release DataFrames:
//
//	df, err := gf.ReadParquet("data.parquet")
//	if err != nil {
//	    return err
//	}
//	defer df.Release()  // Explicit resource management
//
// Check for errors:
//
//	if df.Err() != nil {
//	    return df.Err()
//	}
//
// Use type-safe literals:
//
//	df.Filter(df.Col("age").Gt(gf.Lit(int64(18))))  // Correct
//	// Not: df.Filter(df.Col("age").Gt(gf.Lit(18)))  // Wrong type
//
// Set memory limits in production:
//
//	limited := core.NewLimitedAllocator(pool, maxBytes)
//	df := gf.NewDataFrameWithAllocator(record, limited)
//
// # Documentation
//
// For complete documentation and examples:
//   - API Reference: https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame
//   - GitHub: https://github.com/felixgeelhaar/GopherFrame
//   - Examples: https://github.com/felixgeelhaar/GopherFrame/tree/main/cmd/examples
//
// # License
//
// Apache License 2.0 - see LICENSE file for details.
package gopherframe
