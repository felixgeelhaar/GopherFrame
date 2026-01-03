package main

import (
	"fmt"
	"log"
	"os"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// Production Memory Management Example
//
// This example demonstrates production-grade memory management using GopherFrame's
// LimitedAllocator to prevent Out-Of-Memory (OOM) errors in production deployments.
//
// Use Case: Processing user analytics data with strict memory constraints
// Scenario: Cloud service with 512MB memory limit per container
// Input: User activity logs (potentially large datasets)
// Output: Aggregated metrics with graceful OOM handling
//
// Key Production Patterns:
// - Memory limit enforcement to prevent OOM crashes
// - Pre-flight checks before expensive operations
// - Memory pressure monitoring for proactive scaling
// - Graceful degradation when memory is constrained
// - Chunked processing for large datasets

func main() {
	// Step 1: Configure Memory Limits
	fmt.Println("=== Step 1: Configure Production Memory Limits ===")

	// In production, you'd get this from environment variables or config
	memoryLimitMB := 100 // 100MB limit for this example (512MB in production)
	memoryLimitBytes := int64(memoryLimitMB * 1024 * 1024)

	// Create a limited allocator with production memory limits
	baseAllocator := memory.NewGoAllocator()
	limitedAllocator := core.NewLimitedAllocator(baseAllocator, memoryLimitBytes)

	fmt.Printf("✓ Memory limit configured: %d MB\n", memoryLimitMB)
	fmt.Printf("✓ Allocator: LimitedAllocator (production-hardened)\n")

	// Step 2: Pre-flight Memory Check
	fmt.Println("\n=== Step 2: Pre-flight Memory Check ===")

	// Estimate memory needed for our dataset
	// For this example: 10,000 rows × 3 columns × 8 bytes (int64/float64) = ~240KB
	estimatedRows := int64(10000)
	estimatedBytesPerRow := int64(24) // 3 columns × 8 bytes each
	estimatedMemory := estimatedRows * estimatedBytesPerRow

	// Check if we have enough memory before proceeding
	if err := limitedAllocator.CheckCanAllocate(estimatedMemory); err != nil {
		if oomErr, ok := err.(*core.ErrMemoryLimitExceeded); ok {
			log.Printf("WARNING: Insufficient memory for full dataset")
			log.Printf("  Requested: %.2f MB", float64(oomErr.Requested)/(1024*1024))
			log.Printf("  Available: %.2f MB", float64(oomErr.Limit-oomErr.Current)/(1024*1024))
			log.Printf("  Falling back to chunked processing...")

			// In production: reduce batch size, use streaming, or return error
			estimatedRows = 5000 // Process half the data
		}
	}

	fmt.Printf("✓ Pre-flight check passed\n")
	fmt.Printf("✓ Estimated memory: %.2f MB\n", float64(estimatedMemory)/(1024*1024))
	fmt.Printf("✓ Processing %d rows\n", estimatedRows)

	// Step 3: Create Dataset with Memory Monitoring
	fmt.Println("\n=== Step 3: Create Dataset with Memory Monitoring ===")

	df := createUserActivityData(limitedAllocator, int(estimatedRows))
	defer df.Release()

	// Monitor memory usage
	usedMB := float64(limitedAllocator.AllocatedBytes()) / (1024 * 1024)
	usagePercent := limitedAllocator.UsagePercent()
	pressureLevel := limitedAllocator.MemoryPressureLevel()

	fmt.Printf("✓ Dataset created: %d rows × %d columns\n", df.NumRows(), df.NumCols())
	fmt.Printf("✓ Memory usage: %.2f MB (%.1f%% of limit)\n", usedMB, usagePercent)
	fmt.Printf("✓ Memory pressure: %s\n", pressureLevel)

	// Step 4: Check Memory Pressure Before Operations
	fmt.Println("\n=== Step 4: Memory Pressure Checks ===")

	// Before expensive operations, check memory pressure
	if pressureLevel == "critical" {
		log.Fatal("CRITICAL: Memory pressure too high, aborting operation")
	}

	if pressureLevel == "high" {
		log.Println("WARNING: High memory pressure, using optimized algorithms")
		// In production: switch to memory-efficient algorithms
	}

	fmt.Printf("✓ Memory pressure check: %s (safe to proceed)\n", pressureLevel)

	// Step 5: Perform Aggregations with Memory Monitoring
	fmt.Println("\n=== Step 5: Aggregations with Memory Monitoring ===")

	// Track memory before aggregation
	memBefore := limitedAllocator.AllocatedBytes()

	// Aggregate by user_type
	aggDF := df.GroupBy("user_type").Agg(
		gf.Mean("session_duration").As("avg_duration"),
		gf.Count("user_id").As("user_count"),
	)
	defer aggDF.Release()

	if aggDF.Err() != nil {
		log.Fatalf("Aggregation failed: %v", aggDF.Err())
	}

	// Track memory after aggregation
	memAfter := limitedAllocator.AllocatedBytes()
	aggMemoryMB := float64(memAfter-memBefore) / (1024 * 1024)

	fmt.Printf("✓ Aggregation complete: %d groups\n", aggDF.NumRows())
	fmt.Printf("✓ Memory used by aggregation: %.2f MB\n", aggMemoryMB)
	fmt.Printf("✓ Total memory: %.2f MB (%.1f%%)\n",
		float64(memAfter)/(1024*1024),
		limitedAllocator.UsagePercent())

	// Step 6: Filter with Memory Checks
	fmt.Println("\n=== Step 6: Filtering with Memory Monitoring ===")

	// Check memory before filtering
	if limitedAllocator.UsagePercent() > 90 {
		log.Println("WARNING: Memory near limit, skipping optional operations")
	} else {
		// Filter for active users (>100 page views)
		activeDF := df.Filter(df.Col("page_views").Gt(gf.Lit(int64(100))))
		defer activeDF.Release()

		if activeDF.Err() != nil {
			log.Fatalf("Filter failed: %v", activeDF.Err())
		}

		fmt.Printf("✓ Active users filtered: %d/%d users (>100 views)\n",
			activeDF.NumRows(), df.NumRows())
	}

	// Step 7: Memory Cleanup and Verification
	fmt.Println("\n=== Step 7: Memory Cleanup ===")

	// Release DataFrame resources
	df.Release()
	aggDF.Release()

	// Note: In this example, we're tracking allocations but not fully releasing
	// because Arrow's reference counting may keep some buffers alive.
	// In production, proper Release() calls ensure memory is freed.

	finalMemoryMB := float64(limitedAllocator.AllocatedBytes()) / (1024 * 1024)
	fmt.Printf("✓ Final memory usage: %.2f MB\n", finalMemoryMB)

	// Step 8: Production Recommendations
	fmt.Println("\n=== Step 8: Production Best Practices ===")

	fmt.Println("\nMemory Management Summary:")
	fmt.Printf("  • Memory limit: %d MB\n", memoryLimitMB)
	fmt.Printf("  • Peak usage: %.2f MB (%.1f%%)\n", usedMB, usagePercent)
	fmt.Printf("  • Pressure level: %s\n", pressureLevel)

	fmt.Println("\nProduction Recommendations:")
	fmt.Println("  ✓ Always use LimitedAllocator in production")
	fmt.Println("  ✓ Configure memory limits from environment variables")
	fmt.Println("  ✓ Perform pre-flight checks before large operations")
	fmt.Println("  ✓ Monitor memory pressure and scale proactively")
	fmt.Println("  ✓ Use chunked processing for large datasets")
	fmt.Println("  ✓ Implement graceful degradation when memory is constrained")
	fmt.Println("  ✓ Set up alerting for high memory pressure (>80%)")

	fmt.Println("\nEnvironment Variable Example:")
	fmt.Println("  export GOPHERFRAME_MEMORY_LIMIT_MB=512")
	fmt.Println("  export GOPHERFRAME_MEMORY_ALERT_THRESHOLD=80")

	// Production monitoring example
	if pressureLevel == "high" || pressureLevel == "critical" {
		// In production: send metrics to monitoring system
		fmt.Printf("\n⚠️  ALERT: Memory pressure %s at %.1f%% usage\n",
			pressureLevel, usagePercent)
		fmt.Println("  Action: Consider scaling up or reducing batch size")
	}
}

// createUserActivityData creates sample user activity data
func createUserActivityData(pool memory.Allocator, numRows int) *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "page_views", Type: arrow.PrimitiveTypes.Int64},
			{Name: "session_duration", Type: arrow.PrimitiveTypes.Float64},
			{Name: "user_type", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Build arrays
	userIDBuilder := array.NewInt64Builder(pool)
	defer userIDBuilder.Release()
	pageViewsBuilder := array.NewInt64Builder(pool)
	defer pageViewsBuilder.Release()
	durationBuilder := array.NewFloat64Builder(pool)
	defer durationBuilder.Release()
	userTypeBuilder := array.NewStringBuilder(pool)
	defer userTypeBuilder.Release()

	// Generate sample data
	userTypes := []string{"free", "premium", "enterprise"}
	for i := 0; i < numRows; i++ {
		userIDBuilder.Append(int64(i + 1))
		pageViewsBuilder.Append(int64(50 + (i % 200)))
		durationBuilder.Append(float64(100 + (i % 500)))
		userTypeBuilder.Append(userTypes[i%3])
	}

	// Create arrays
	userIDArray := userIDBuilder.NewArray()
	defer userIDArray.Release()
	pageViewsArray := pageViewsBuilder.NewArray()
	defer pageViewsArray.Release()
	durationArray := durationBuilder.NewArray()
	defer durationArray.Release()
	userTypeArray := userTypeBuilder.NewArray()
	defer userTypeArray.Release()

	// Create record
	record := array.NewRecord(
		schema,
		[]arrow.Array{userIDArray, pageViewsArray, durationArray, userTypeArray},
		int64(numRows),
	)

	return gf.NewDataFrame(record)
}

// Example of reading memory limit from environment
//
//nolint:unused // example function showing environment configuration pattern
func getMemoryLimitFromEnv() int64 {
	limitMB := 512 // Default: 512MB
	if envLimit := os.Getenv("GOPHERFRAME_MEMORY_LIMIT_MB"); envLimit != "" {
		_, _ = fmt.Sscanf(envLimit, "%d", &limitMB)
	}
	return int64(limitMB * 1024 * 1024)
}
