package main

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	. "github.com/felixgeelhaar/GopherFrame/pkg/interfaces" //nolint:staticcheck // ST1001: dot import intentional for example readability
)

// Statistical Analysis Example
//
// This example demonstrates GopherFrame's advanced statistical aggregation
// capabilities including:
// - Percentile calculations (p50, p95, p99 for SLA monitoring)
// - Median for robust central tendency analysis
// - Mode for finding most common values
// - Correlation for relationship analysis
//
// Use cases:
// - Performance monitoring and SLA tracking
// - A/B testing statistical analysis
// - Market research and customer behavior analysis
// - Quality control and process monitoring

func main() {
	fmt.Println("=== GopherFrame Statistical Analysis Example ===")

	pool := memory.NewGoAllocator()

	// Example 1: API Response Time Analysis (SLA Monitoring)
	fmt.Println("Example 1: API Response Time Analysis")
	fmt.Println("-------------------------------------")
	responseDF := createResponseTimeData(pool)
	defer responseDF.Release()

	fmt.Println("\nSample Response Time Data:")
	printDataFrame(responseDF, 10)

	// Calculate percentiles for SLA monitoring
	slaMetrics := responseDF.GroupBy("endpoint").Agg(
		Count("response_time").As("request_count"),
		Mean("response_time").As("mean_ms"),
		Median("response_time").As("median_ms"),
		Percentile("response_time", 0.50).As("p50_ms"),
		Percentile("response_time", 0.95).As("p95_ms"),
		Percentile("response_time", 0.99).As("p99_ms"),
		Min("response_time").As("min_ms"),
		Max("response_time").As("max_ms"),
	)
	defer slaMetrics.Release()

	fmt.Println("\nSLA Metrics by Endpoint:")
	fmt.Println("(p50, p95, p99 are key performance indicators)")
	printDataFrame(slaMetrics, 10)

	// Example 2: Customer Segmentation Analysis
	fmt.Println("\n\nExample 2: Customer Purchase Behavior")
	fmt.Println("--------------------------------------")
	customerDF := createCustomerData(pool)
	defer customerDF.Release()

	fmt.Println("\nSample Customer Data:")
	printDataFrame(customerDF, 10)

	// Analyze customer segments
	segmentAnalysis := customerDF.GroupBy("segment").Agg(
		Count("customer_id").As("customer_count"),
		Mean("purchase_amount").As("avg_purchase"),
		Median("purchase_amount").As("median_purchase"),
		Percentile("purchase_amount", 0.75).As("p75_purchase"),
		Mode("payment_method").As("preferred_payment"),
		Min("purchase_amount").As("min_purchase"),
		Max("purchase_amount").As("max_purchase"),
	)
	defer segmentAnalysis.Release()

	fmt.Println("\nCustomer Segment Analysis:")
	printDataFrame(segmentAnalysis, 10)

	// Example 3: Correlation Analysis (Marketing Effectiveness)
	fmt.Println("\n\nExample 3: Marketing Campaign Correlation Analysis")
	fmt.Println("--------------------------------------------------")
	marketingDF := createMarketingData(pool)
	defer marketingDF.Release()

	fmt.Println("\nSample Marketing Data:")
	printDataFrame(marketingDF, 10)

	// Analyze correlations between different metrics
	correlationAnalysis := marketingDF.GroupBy("campaign").Agg(
		Count("week").As("weeks"),
		Mean("ad_spend").As("avg_ad_spend"),
		Mean("impressions").As("avg_impressions"),
		Mean("clicks").As("avg_clicks"),
		Mean("revenue").As("avg_revenue"),
		Correlation("ad_spend", "revenue").As("spend_revenue_corr"),
		Correlation("impressions", "clicks").As("impressions_clicks_corr"),
		Correlation("clicks", "revenue").As("clicks_revenue_corr"),
	)
	defer correlationAnalysis.Release()

	fmt.Println("\nMarketing Campaign Correlation Analysis:")
	fmt.Println("Correlation ranges from -1.0 (negative) to +1.0 (positive)")
	printDataFrame(correlationAnalysis, 10)

	// Example 4: Quality Control Process Monitoring
	fmt.Println("\n\nExample 4: Manufacturing Quality Control")
	fmt.Println("----------------------------------------")
	qualityDF := createQualityData(pool)
	defer qualityDF.Release()

	fmt.Println("\nSample Quality Measurements:")
	printDataFrame(qualityDF, 10)

	// Statistical process control metrics
	processControl := qualityDF.GroupBy("product_line").Agg(
		Count("measurement").As("sample_size"),
		Mean("measurement").As("process_mean"),
		Median("measurement").As("process_median"),
		Percentile("measurement", 0.05).As("lower_5pct"),
		Percentile("measurement", 0.95).As("upper_5pct"),
		Min("measurement").As("minimum"),
		Max("measurement").As("maximum"),
		Mode("defect_type").As("most_common_defect"),
	)
	defer processControl.Release()

	fmt.Println("\nStatistical Process Control Metrics:")
	fmt.Println("(5th and 95th percentiles define control limits)")
	printDataFrame(processControl, 10)

	// Example 5: Multi-Group Statistical Comparison
	fmt.Println("\n\nExample 5: A/B Testing Results Analysis")
	fmt.Println("----------------------------------------")
	abtestDF := createABTestData(pool)
	defer abtestDF.Release()

	fmt.Println("\nSample A/B Test Data:")
	printDataFrame(abtestDF, 10)

	// Compare variants across multiple metrics
	variantComparison := abtestDF.GroupBy("variant", "segment").Agg(
		Count("user_id").As("users"),
		Mean("conversion_rate").As("mean_cvr"),
		Median("conversion_rate").As("median_cvr"),
		Percentile("conversion_rate", 0.25).As("q1_cvr"),
		Percentile("conversion_rate", 0.75).As("q3_cvr"),
		Mean("revenue_per_user").As("mean_rpu"),
		Median("revenue_per_user").As("median_rpu"),
	)
	defer variantComparison.Release()

	fmt.Println("\nA/B Test Statistical Comparison:")
	fmt.Println("(Compare median and percentiles for robust insights)")
	printDataFrame(variantComparison, 12)

	fmt.Println("\n=== All Statistical Analysis Examples Completed ===")
}

// createResponseTimeData generates API response time data
func createResponseTimeData(pool memory.Allocator) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "endpoint", Type: arrow.BinaryTypes.String},
			{Name: "response_time", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	endpointBuilder := array.NewStringBuilder(pool)
	defer endpointBuilder.Release()
	timeBuilder := array.NewFloat64Builder(pool)
	defer timeBuilder.Release()

	endpoints := []string{"/api/users", "/api/orders", "/api/products"}
	baseTimes := []float64{50.0, 120.0, 80.0}

	// Generate 100 samples per endpoint with realistic distribution
	for i, endpoint := range endpoints {
		baseTime := baseTimes[i]
		for j := 0; j < 100; j++ {
			endpointBuilder.Append(endpoint)

			// Simulate realistic response time distribution
			// Most requests are fast, some are slower (long tail)
			variation := float64(j % 30)
			outlier := 0.0
			if j%20 == 0 { // 5% outliers
				outlier = float64(j * 5)
			}
			timeBuilder.Append(baseTime + variation + outlier)
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		endpointBuilder.NewArray(),
		timeBuilder.NewArray(),
	}, 300)

	return NewDataFrame(record)
}

// createCustomerData generates customer purchase data
func createCustomerData(pool memory.Allocator) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "customer_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "segment", Type: arrow.BinaryTypes.String},
			{Name: "purchase_amount", Type: arrow.PrimitiveTypes.Float64},
			{Name: "payment_method", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	defer idBuilder.Release()
	segmentBuilder := array.NewStringBuilder(pool)
	defer segmentBuilder.Release()
	amountBuilder := array.NewFloat64Builder(pool)
	defer amountBuilder.Release()
	methodBuilder := array.NewStringBuilder(pool)
	defer methodBuilder.Release()

	segments := []string{"Premium", "Standard", "Budget"}
	baseAmounts := []float64{500.0, 150.0, 50.0}
	methods := []string{"Credit", "Debit", "PayPal", "Cash"}

	customerID := int64(1)
	for segIdx, segment := range segments {
		for i := 0; i < 50; i++ {
			idBuilder.Append(customerID)
			segmentBuilder.Append(segment)

			// Purchase amount varies by segment
			amount := baseAmounts[segIdx] + float64(i*5)
			amountBuilder.Append(amount)

			// Payment method preference varies
			methodIdx := (segIdx*17 + i) % len(methods)
			methodBuilder.Append(methods[methodIdx])

			customerID++
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		idBuilder.NewArray(),
		segmentBuilder.NewArray(),
		amountBuilder.NewArray(),
		methodBuilder.NewArray(),
	}, 150)

	return NewDataFrame(record)
}

// createMarketingData generates marketing campaign data
func createMarketingData(pool memory.Allocator) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "campaign", Type: arrow.BinaryTypes.String},
			{Name: "week", Type: arrow.PrimitiveTypes.Int64},
			{Name: "ad_spend", Type: arrow.PrimitiveTypes.Float64},
			{Name: "impressions", Type: arrow.PrimitiveTypes.Float64},
			{Name: "clicks", Type: arrow.PrimitiveTypes.Float64},
			{Name: "revenue", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	campaignBuilder := array.NewStringBuilder(pool)
	defer campaignBuilder.Release()
	weekBuilder := array.NewInt64Builder(pool)
	defer weekBuilder.Release()
	spendBuilder := array.NewFloat64Builder(pool)
	defer spendBuilder.Release()
	impressionsBuilder := array.NewFloat64Builder(pool)
	defer impressionsBuilder.Release()
	clicksBuilder := array.NewFloat64Builder(pool)
	defer clicksBuilder.Release()
	revenueBuilder := array.NewFloat64Builder(pool)
	defer revenueBuilder.Release()

	campaigns := []string{"Campaign A", "Campaign B", "Campaign C"}

	// Generate 12 weeks of data per campaign
	for campIdx, campaign := range campaigns {
		effectiveness := []float64{1.5, 1.2, 0.8}[campIdx] // Different ROI

		for week := 1; week <= 12; week++ {
			campaignBuilder.Append(campaign)
			weekBuilder.Append(int64(week))

			// Ad spend varies by week
			adSpend := 5000.0 + float64(week)*200.0
			spendBuilder.Append(adSpend)

			// Impressions proportional to spend with noise
			impressions := adSpend * 100.0 * (1.0 + float64(week%3)*0.1)
			impressionsBuilder.Append(impressions)

			// Clicks proportional to impressions (CTR ~2%)
			clicks := impressions * 0.02 * (1.0 + float64(week%5)*0.05)
			clicksBuilder.Append(clicks)

			// Revenue correlated with clicks and spend
			revenue := clicks * 15.0 * effectiveness
			revenueBuilder.Append(revenue)
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		campaignBuilder.NewArray(),
		weekBuilder.NewArray(),
		spendBuilder.NewArray(),
		impressionsBuilder.NewArray(),
		clicksBuilder.NewArray(),
		revenueBuilder.NewArray(),
	}, 36)

	return NewDataFrame(record)
}

// createQualityData generates quality control measurements
func createQualityData(pool memory.Allocator) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "product_line", Type: arrow.BinaryTypes.String},
			{Name: "measurement", Type: arrow.PrimitiveTypes.Float64},
			{Name: "defect_type", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	lineBuilder := array.NewStringBuilder(pool)
	defer lineBuilder.Release()
	measurementBuilder := array.NewFloat64Builder(pool)
	defer measurementBuilder.Release()
	defectBuilder := array.NewStringBuilder(pool)
	defer defectBuilder.Release()

	productLines := []string{"Line A", "Line B", "Line C"}
	targetValues := []float64{100.0, 50.0, 75.0}
	defectTypes := []string{"None", "Minor", "Major"}

	for lineIdx, line := range productLines {
		target := targetValues[lineIdx]

		for i := 0; i < 80; i++ {
			lineBuilder.Append(line)

			// Measurements normally distributed around target
			// Most within ±5%, some outliers
			variation := float64((i%11)-5) * 0.5
			measurement := target + variation
			measurementBuilder.Append(measurement)

			// Defect distribution
			defectIdx := 0 // None
			if i%10 == 0 {
				defectIdx = 1 // Minor
			} else if i%20 == 0 {
				defectIdx = 2 // Major
			}
			defectBuilder.Append(defectTypes[defectIdx])
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		lineBuilder.NewArray(),
		measurementBuilder.NewArray(),
		defectBuilder.NewArray(),
	}, 240)

	return NewDataFrame(record)
}

// createABTestData generates A/B testing data
func createABTestData(pool memory.Allocator) *DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "user_id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "variant", Type: arrow.BinaryTypes.String},
			{Name: "segment", Type: arrow.BinaryTypes.String},
			{Name: "conversion_rate", Type: arrow.PrimitiveTypes.Float64},
			{Name: "revenue_per_user", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	idBuilder := array.NewInt64Builder(pool)
	defer idBuilder.Release()
	variantBuilder := array.NewStringBuilder(pool)
	defer variantBuilder.Release()
	segmentBuilder := array.NewStringBuilder(pool)
	defer segmentBuilder.Release()
	cvrBuilder := array.NewFloat64Builder(pool)
	defer cvrBuilder.Release()
	rpuBuilder := array.NewFloat64Builder(pool)
	defer rpuBuilder.Release()

	variants := []string{"Control", "Variant A", "Variant B"}
	segments := []string{"New Users", "Returning Users"}

	userID := int64(1)
	for _, variant := range variants {
		// Variant A performs better
		variantLift := map[string]float64{
			"Control":   1.0,
			"Variant A": 1.15, // 15% lift
			"Variant B": 1.08, // 8% lift
		}[variant]

		for _, segment := range segments {
			// Returning users convert better
			segmentBase := map[string]float64{
				"New Users":       0.05,
				"Returning Users": 0.12,
			}[segment]

			for i := 0; i < 30; i++ {
				idBuilder.Append(userID)
				variantBuilder.Append(variant)
				segmentBuilder.Append(segment)

				// Conversion rate with realistic variation
				cvr := segmentBase * variantLift * (1.0 + float64(i%5)*0.02)
				cvrBuilder.Append(cvr)

				// Revenue per user
				rpu := cvr * 50.0 * (1.0 + float64(i%7)*0.1)
				rpuBuilder.Append(rpu)

				userID++
			}
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		idBuilder.NewArray(),
		variantBuilder.NewArray(),
		segmentBuilder.NewArray(),
		cvrBuilder.NewArray(),
		rpuBuilder.NewArray(),
	}, 180)

	return NewDataFrame(record)
}

// printDataFrame prints a DataFrame in a formatted table
func printDataFrame(df *DataFrame, maxRows int) {
	if df == nil || df.Err() != nil {
		fmt.Println("Error:", df.Err())
		return
	}

	// Print column names
	names := df.ColumnNames()
	for i, name := range names {
		if i > 0 {
			fmt.Print(" | ")
		}
		fmt.Printf("%-18s", name)
	}
	fmt.Println()

	// Print separator
	for i := range names {
		if i > 0 {
			fmt.Print("-+-")
		}
		fmt.Print("------------------")
	}
	fmt.Println()

	// Print rows (would need actual data access implementation)
	fmt.Printf("... (%d rows with %d columns)\n", df.NumRows(), df.NumCols())

	if int(df.NumRows()) > maxRows {
		fmt.Printf("(showing first %d of %d rows)\n", maxRows, df.NumRows())
	}
}
