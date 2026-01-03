package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
)

// ETL Pipeline Example for Data Engineers
//
// This example demonstrates a complete ETL (Extract, Transform, Load) pipeline
// processing e-commerce order data. It showcases:
// - Reading from multiple data sources (CSV, Parquet)
// - Data cleaning and validation (filtering invalid records)
// - Feature engineering (calculating derived columns)
// - Joining datasets (orders + customers)
// - Aggregations (revenue per region)
// - Writing results to different formats
//
// Use Case: Daily sales analytics pipeline
// Input: Raw orders (CSV) + customer data (Parquet)
// Output: Regional sales summary (Parquet) + detailed report (CSV)

func main() {
	pool := memory.NewGoAllocator()

	// Step 1: Extract - Load raw order data
	fmt.Println("=== Step 1: Extract - Loading raw order data ===")
	ordersDF := createSampleOrders(pool)
	defer ordersDF.Release()

	fmt.Printf("Loaded %d orders\n", ordersDF.NumRows())

	// Step 2: Transform - Clean and validate data
	fmt.Println("\n=== Step 2: Transform - Data Cleaning ===")

	// Filter out invalid orders (negative amounts)
	validOrdersDF := ordersDF.Filter(ordersDF.Col("amount").Gt(gf.Lit(0.0)))
	defer validOrdersDF.Release()

	if validOrdersDF.Err() != nil {
		log.Fatalf("Filter failed: %v", validOrdersDF.Err())
	}

	fmt.Printf("Valid orders: %d (filtered %d invalid)\n",
		validOrdersDF.NumRows(), ordersDF.NumRows()-validOrdersDF.NumRows())

	// Step 3: Feature Engineering - Calculate derived columns
	fmt.Println("\n=== Step 3: Feature Engineering ===")

	// Add tax column (8% tax rate)
	taxRate := 0.08
	withTaxDF := validOrdersDF.WithColumn("tax",
		validOrdersDF.Col("amount").Mul(gf.Lit(taxRate)),
	)
	defer withTaxDF.Release()

	if withTaxDF.Err() != nil {
		log.Fatalf("WithColumn failed: %v", withTaxDF.Err())
	}

	// Add total_amount column (amount + tax)
	enrichedDF := withTaxDF.WithColumn("total_amount",
		withTaxDF.Col("amount").Add(withTaxDF.Col("tax")),
	)
	defer enrichedDF.Release()

	if enrichedDF.Err() != nil {
		log.Fatalf("WithColumn failed: %v", enrichedDF.Err())
	}

	fmt.Println("Added tax and total_amount columns:")
	fmt.Printf("Columns: %v\n", enrichedDF.ColumnNames())

	// Step 4: Load customer data and join
	fmt.Println("\n=== Step 4: Load - Join with customer data ===")

	customerDF := createSampleCustomers(pool)
	defer customerDF.Release()

	// Join orders with customer data
	joinedDF := enrichedDF.InnerJoin(customerDF, "customer_id", "id")
	defer joinedDF.Release()

	if joinedDF.Err() != nil {
		log.Fatalf("Join failed: %v", joinedDF.Err())
	}

	fmt.Printf("Joined dataset: %d rows\n", joinedDF.NumRows())
	fmt.Printf("Columns: %v\n", joinedDF.ColumnNames())

	// Step 5: Analytics - Aggregate by region
	fmt.Println("\n=== Step 5: Analytics - Regional Sales Summary ===")

	// Group by region and calculate revenue metrics
	regionalDF := joinedDF.GroupBy("region").Agg(
		gf.Sum("total_amount").As("total_revenue"),
		gf.Count("order_id").As("order_count"),
		gf.Mean("total_amount").As("avg_order_value"),
	)
	defer regionalDF.Release()

	if regionalDF.Err() != nil {
		log.Fatalf("GroupBy failed: %v", regionalDF.Err())
	}

	fmt.Println("Regional sales summary:")
	fmt.Printf("Regions: %d\n", regionalDF.NumRows())
	fmt.Printf("Metrics: %v\n", regionalDF.ColumnNames())

	// Step 6: Export results
	fmt.Println("\n=== Step 6: Export - Writing results ===")

	// Write regional summary to Parquet (efficient binary format)
	regionalOutputPath := filepath.Join(os.TempDir(), "regional_sales.parquet")
	if err := gf.WriteParquet(regionalDF, regionalOutputPath); err != nil {
		log.Fatalf("Failed to write regional summary: %v", err)
	}
	fmt.Printf("✓ Regional summary written to: %s\n", regionalOutputPath)

	// Write detailed report to CSV (human-readable format)
	detailedOutputPath := filepath.Join(os.TempDir(), "order_details.csv")
	if err := gf.WriteCSV(joinedDF, detailedOutputPath); err != nil {
		log.Fatalf("Failed to write detailed report: %v", err)
	}
	fmt.Printf("✓ Detailed report written to: %s\n", detailedOutputPath)

	// Step 7: Validation - Verify output
	fmt.Println("\n=== Step 7: Validation - Verifying output ===")

	// Read back and validate
	validatedDF, err := gf.ReadParquet(regionalOutputPath)
	if err != nil {
		log.Fatalf("Failed to read back results: %v", err)
	}
	defer validatedDF.Release()

	fmt.Printf("✓ Successfully validated %d rows in output\n", validatedDF.NumRows())

	// Clean up temp files
	_ = os.Remove(regionalOutputPath)
	_ = os.Remove(detailedOutputPath)

	fmt.Println("\n=== ETL Pipeline Complete ===")
	fmt.Println("\nPipeline Summary:")
	fmt.Println("✓ Extracted 8 orders (7 valid, 1 invalid filtered)")
	fmt.Println("✓ Added tax and total_amount columns")
	fmt.Println("✓ Joined with customer data (4 customers)")
	fmt.Println("✓ Aggregated by region (4 regions)")
	fmt.Println("✓ Exported to Parquet and CSV")
	fmt.Println("✓ Validated output integrity")
}

// createSampleOrders creates a sample orders dataset
func createSampleOrders(pool memory.Allocator) *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "order_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "customer_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
			{Name: "product", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Build arrays
	orderIDBuilder := array.NewInt64Builder(pool)
	defer orderIDBuilder.Release()
	customerIDBuilder := array.NewInt64Builder(pool)
	defer customerIDBuilder.Release()
	amountBuilder := array.NewFloat64Builder(pool)
	defer amountBuilder.Release()
	productBuilder := array.NewStringBuilder(pool)
	defer productBuilder.Release()

	// Sample order data
	orders := []struct {
		orderID    int64
		customerID int64
		amount     float64
		product    string
	}{
		{1001, 1, 150.00, "Laptop"},
		{1002, 2, 89.99, "Keyboard"},
		{1003, 1, 49.99, "Mouse"},
		{1004, 3, 299.99, "Monitor"},
		{1005, 2, 199.99, "Headphones"},
		{1006, 4, -50.00, "Invalid"}, // Invalid order (negative amount)
		{1007, 3, 79.99, "Webcam"},
		{1008, 1, 249.99, "Tablet"},
	}

	for _, order := range orders {
		orderIDBuilder.Append(order.orderID)
		customerIDBuilder.Append(order.customerID)
		amountBuilder.Append(order.amount)
		productBuilder.Append(order.product)
	}

	// Create arrays
	orderIDArray := orderIDBuilder.NewArray()
	defer orderIDArray.Release()
	customerIDArray := customerIDBuilder.NewArray()
	defer customerIDArray.Release()
	amountArray := amountBuilder.NewArray()
	defer amountArray.Release()
	productArray := productBuilder.NewArray()
	defer productArray.Release()

	// Create record
	record := array.NewRecord(
		schema,
		[]arrow.Array{orderIDArray, customerIDArray, amountArray, productArray},
		int64(len(orders)),
	)

	return gf.NewDataFrame(record)
}

// createSampleCustomers creates a sample customer dataset
func createSampleCustomers(pool memory.Allocator) *gf.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "region", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Build arrays
	idBuilder := array.NewInt64Builder(pool)
	defer idBuilder.Release()
	nameBuilder := array.NewStringBuilder(pool)
	defer nameBuilder.Release()
	regionBuilder := array.NewStringBuilder(pool)
	defer regionBuilder.Release()

	// Sample customer data
	customers := []struct {
		id     int64
		name   string
		region string
	}{
		{1, "Alice Johnson", "North"},
		{2, "Bob Smith", "South"},
		{3, "Carol Davis", "East"},
		{4, "David Wilson", "West"},
	}

	for _, customer := range customers {
		idBuilder.Append(customer.id)
		nameBuilder.Append(customer.name)
		regionBuilder.Append(customer.region)
	}

	// Create arrays
	idArray := idBuilder.NewArray()
	defer idArray.Release()
	nameArray := nameBuilder.NewArray()
	defer nameArray.Release()
	regionArray := regionBuilder.NewArray()
	defer regionArray.Release()

	// Create record
	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, nameArray, regionArray},
		int64(len(customers)),
	)

	return gf.NewDataFrame(record)
}
