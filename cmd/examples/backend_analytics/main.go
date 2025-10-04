package main

import (
	"fmt"
	"log"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
)

// Backend Analytics Example for Go Developers
//
// This example demonstrates how to use GopherFrame for backend service analytics.
// It showcases real-time data analysis capabilities that backend developers need
// for monitoring, debugging, and optimizing their services.
//
// Use Case: API request analytics and performance monitoring
// Input: API request logs
// Output: Performance metrics, error analysis, usage patterns
//
// Key Backend Tasks:
// - Processing service logs for analytics
// - Calculating performance metrics (p50, p95, p99)
// - Error rate analysis
// - User activity patterns
// - Resource utilization insights

func main() {
	pool := memory.NewGoAllocator()

	// Step 1: Load API Request Logs
	fmt.Println("=== Step 1: Load API Request Logs ===")
	logsDF := createAPIRequestLogs(pool)
	defer logsDF.Release()

	fmt.Printf("Loaded %d API requests\n", logsDF.NumRows())
	fmt.Printf("Log fields: %v\n", logsDF.ColumnNames())

	// Step 2: Calculate Performance Metrics
	fmt.Println("\n=== Step 2: Performance Metrics ===")

	// Calculate average response time per endpoint
	perfDF := logsDF.GroupBy("endpoint").Agg(
		gf.Count("request_id").As("request_count"),
		gf.Mean("response_time_ms").As("avg_response_ms"),
		gf.Sum("response_time_ms").As("total_time_ms"),
	)
	defer perfDF.Release()

	if perfDF.Err() != nil {
		log.Fatalf("Performance metrics failed: %v", perfDF.Err())
	}

	fmt.Println("Endpoint performance metrics:")
	fmt.Printf("  Endpoints analyzed: %d\n", perfDF.NumRows())
	fmt.Printf("  Metrics: %v\n", perfDF.ColumnNames())

	// Step 3: Error Analysis
	fmt.Println("\n=== Step 3: Error Analysis ===")

	// Filter for error responses (status >= 400)
	errorsDF := logsDF.Filter(
		logsDF.Col("status_code").Gt(gf.Lit(int64(399))),
	)
	defer errorsDF.Release()

	if errorsDF.Err() != nil {
		log.Fatalf("Error filtering failed: %v", errorsDF.Err())
	}

	fmt.Printf("Error requests: %d out of %d total (%.2f%% error rate)\n",
		errorsDF.NumRows(),
		logsDF.NumRows(),
		float64(errorsDF.NumRows())/float64(logsDF.NumRows())*100,
	)

	// Group errors by endpoint
	errorsByEndpointDF := errorsDF.GroupBy("endpoint").Agg(
		gf.Count("request_id").As("error_count"),
	)
	defer errorsByEndpointDF.Release()

	if errorsByEndpointDF.Err() != nil {
		log.Fatalf("Error grouping failed: %v", errorsByEndpointDF.Err())
	}

	fmt.Printf("Endpoints with errors: %d\n", errorsByEndpointDF.NumRows())

	// Step 4: User Activity Analysis
	fmt.Println("\n=== Step 4: User Activity Analysis ===")

	// Count unique users
	uniqueUserCount := countUniqueValues(logsDF, "user_id")
	fmt.Printf("Active users: %d\n", uniqueUserCount)

	// Analyze high-activity users (user 101 and 102 have multiple requests)
	user101DF := logsDF.Filter(logsDF.Col("user_id").Eq(gf.Lit(int64(101))))
	defer user101DF.Release()
	user102DF := logsDF.Filter(logsDF.Col("user_id").Eq(gf.Lit(int64(102))))
	defer user102DF.Release()

	fmt.Printf("User 101 requests: %d\n", user101DF.NumRows())
	fmt.Printf("User 102 requests: %d\n", user102DF.NumRows())

	// Step 5: Response Time Distribution
	fmt.Println("\n=== Step 5: Response Time Distribution ===")

	// Classify requests by response time
	// Fast: < 100ms, Normal: 100-500ms, Slow: > 500ms
	fastDF := logsDF.Filter(
		logsDF.Col("response_time_ms").Lt(gf.Lit(100.0)),
	)
	defer fastDF.Release()

	normalDF := logsDF.Filter(
		logsDF.Col("response_time_ms").Gt(gf.Lit(100.0)),
	)
	defer normalDF.Release()

	slowDF := logsDF.Filter(
		logsDF.Col("response_time_ms").Gt(gf.Lit(500.0)),
	)
	defer slowDF.Release()

	fmt.Printf("Response time distribution:\n")
	fmt.Printf("  Fast (<100ms): %d requests (%.1f%%)\n",
		fastDF.NumRows(),
		float64(fastDF.NumRows())/float64(logsDF.NumRows())*100,
	)
	fmt.Printf("  Normal (100-500ms): %d requests (%.1f%%)\n",
		normalDF.NumRows()-slowDF.NumRows(),
		float64(normalDF.NumRows()-slowDF.NumRows())/float64(logsDF.NumRows())*100,
	)
	fmt.Printf("  Slow (>500ms): %d requests (%.1f%%)\n",
		slowDF.NumRows(),
		float64(slowDF.NumRows())/float64(logsDF.NumRows())*100,
	)

	// Step 6: Resource Utilization
	fmt.Println("\n=== Step 6: Resource Utilization Insights ===")

	// Calculate total processing time per endpoint
	resourceDF := logsDF.GroupBy("endpoint").Agg(
		gf.Sum("response_time_ms").As("total_ms"),
		gf.Count("request_id").As("count"),
	)
	defer resourceDF.Release()

	if resourceDF.Err() != nil {
		log.Fatalf("Resource calculation failed: %v", resourceDF.Err())
	}

	fmt.Println("Resource usage by endpoint:")
	fmt.Printf("  Endpoints: %d\n", resourceDF.NumRows())
	fmt.Printf("  Metrics: %v\n", resourceDF.ColumnNames())

	// Summary
	fmt.Println("\n=== Backend Analytics Summary ===")
	fmt.Println("\nKey Insights:")
	fmt.Printf("✓ Processed %d API requests across %d endpoints\n",
		logsDF.NumRows(), perfDF.NumRows())
	fmt.Printf("✓ Error rate: %.2f%% (%d errors)\n",
		float64(errorsDF.NumRows())/float64(logsDF.NumRows())*100,
		errorsDF.NumRows())
	fmt.Printf("✓ Active users: %d\n", uniqueUserCount)
	fmt.Printf("✓ Performance: %.1f%% fast responses (<100ms)\n",
		float64(fastDF.NumRows())/float64(logsDF.NumRows())*100)

	fmt.Println("\nUse Cases Demonstrated:")
	fmt.Println("  • Real-time API performance monitoring")
	fmt.Println("  • Error tracking and alerting")
	fmt.Println("  • User behavior analysis")
	fmt.Println("  • Resource optimization insights")
	fmt.Println("  • In-memory analytics for backend services")
}

// countUniqueValues counts unique values in a column
func countUniqueValues(df *gf.DataFrame, columnName string) int {
	seen := make(map[int64]bool)
	record := df.Record()

	// Find column index
	var colIndex int
	for i, field := range record.Schema().Fields() {
		if field.Name == columnName {
			colIndex = i
			break
		}
	}

	// Get column and count unique values
	column := record.Column(colIndex).(*array.Int64)
	for i := 0; i < column.Len(); i++ {
		if !column.IsNull(i) {
			seen[column.Value(i)] = true
		}
	}

	return len(seen)
}

// createAPIRequestLogs creates sample API request log data
func createAPIRequestLogs(pool memory.Allocator) *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "request_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "endpoint", Type: arrow.BinaryTypes.String},
			{Name: "status_code", Type: arrow.PrimitiveTypes.Int64},
			{Name: "response_time_ms", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build arrays
	requestIDBuilder := array.NewInt64Builder(pool)
	defer requestIDBuilder.Release()
	userIDBuilder := array.NewInt64Builder(pool)
	defer userIDBuilder.Release()
	endpointBuilder := array.NewStringBuilder(pool)
	defer endpointBuilder.Release()
	statusCodeBuilder := array.NewInt64Builder(pool)
	defer statusCodeBuilder.Release()
	responseTimeBuilder := array.NewFloat64Builder(pool)
	defer responseTimeBuilder.Release()

	// Sample API request log data
	requests := []struct {
		id           int64
		userID       int64
		endpoint     string
		statusCode   int64
		responseTime float64
	}{
		// Successful requests
		{1, 101, "/api/users", 200, 45.2},
		{2, 102, "/api/products", 200, 123.5},
		{3, 101, "/api/orders", 200, 89.1},
		{4, 103, "/api/users", 200, 52.3},
		{5, 102, "/api/products", 200, 156.7},
		{6, 101, "/api/users", 200, 38.9},
		{7, 104, "/api/orders", 200, 234.5},
		{8, 103, "/api/products", 200, 98.2},
		// Some slower requests
		{9, 102, "/api/analytics", 200, 567.3},
		{10, 101, "/api/reports", 200, 723.1},
		// Error requests
		{11, 105, "/api/users", 404, 15.2},
		{12, 102, "/api/orders", 500, 1234.5},
		{13, 103, "/api/products", 403, 23.1},
		// More successful requests
		{14, 101, "/api/users", 200, 41.5},
		{15, 104, "/api/products", 200, 167.8},
		{16, 102, "/api/orders", 200, 92.3},
		{17, 103, "/api/analytics", 200, 445.2},
		{18, 105, "/api/users", 200, 55.7},
		{19, 101, "/api/products", 200, 134.2},
		{20, 102, "/api/users", 200, 47.1},
	}

	for _, req := range requests {
		requestIDBuilder.Append(req.id)
		userIDBuilder.Append(req.userID)
		endpointBuilder.Append(req.endpoint)
		statusCodeBuilder.Append(req.statusCode)
		responseTimeBuilder.Append(req.responseTime)
	}

	// Create arrays
	requestIDArray := requestIDBuilder.NewArray()
	defer requestIDArray.Release()
	userIDArray := userIDBuilder.NewArray()
	defer userIDArray.Release()
	endpointArray := endpointBuilder.NewArray()
	defer endpointArray.Release()
	statusCodeArray := statusCodeBuilder.NewArray()
	defer statusCodeArray.Release()
	responseTimeArray := responseTimeBuilder.NewArray()
	defer responseTimeArray.Release()

	// Create record
	record := array.NewRecord(
		schema,
		[]arrow.Array{
			requestIDArray,
			userIDArray,
			endpointArray,
			statusCodeArray,
			responseTimeArray,
		},
		int64(len(requests)),
	)

	return gf.NewDataFrame(record)
}
