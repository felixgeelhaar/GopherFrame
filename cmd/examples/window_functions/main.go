package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
)

// Window Functions Example
//
// This example demonstrates all GopherFrame window function capabilities:
// - Analytical functions (RowNumber, Rank, DenseRank, Lag, Lead)
// - Rolling aggregations (RollingSum, RollingMean, RollingMin, RollingMax)
// - Cumulative operations (CumSum, CumMax, CumMin, CumProd)
//
// Use cases:
// - Sales analytics with running totals and moving averages
// - Employee ranking within departments
// - Time series analysis with previous period comparisons

func main() {
	fmt.Println("=== GopherFrame Window Functions Example ===")

	pool := memory.NewGoAllocator()

	// Example 1: Sales Analytics with Window Functions
	fmt.Println("Example 1: Sales Analytics")
	fmt.Println("---------------------------")
	salesDF := createSalesData(pool)
	defer salesDF.Release()

	// Display original data
	fmt.Println("\nOriginal Sales Data:")
	printDataFrame(salesDF, 10)

	// Apply multiple window functions at once
	result, err := salesDF.Window().
		PartitionBy("region").
		OrderBy("date").
		Rows(7).
		Over(
			core.RowNumber().As("day_number"),
			core.RollingSum("sales").As("sales_7d"),
			core.RollingMean("sales").As("avg_sales_7d"),
			core.CumSum("sales").As("cumulative_sales"),
			core.Lag("sales", 1).As("previous_day_sales"),
		)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Release()

	fmt.Println("\nWith Window Functions Applied:")
	printDataFrame(result, 15)

	// Example 2: Employee Ranking
	fmt.Println("\n\nExample 2: Employee Ranking by Department")
	fmt.Println("------------------------------------------")
	employeeDF := createEmployeeData(pool)
	defer employeeDF.Release()

	fmt.Println("\nOriginal Employee Data:")
	printDataFrame(employeeDF, 10)

	// Rank employees by salary within each department
	ranked, err := employeeDF.Window().
		PartitionBy("department").
		OrderBy("salary").
		Over(
			core.RowNumber().As("row_num"),
			core.Rank().As("salary_rank"),
			core.DenseRank().As("dense_rank"),
		)
	if err != nil {
		log.Fatal(err)
	}
	defer ranked.Release()

	fmt.Println("\nEmployees Ranked by Salary (within department):")
	printDataFrame(ranked, 15)

	// Example 3: Time Series with Moving Averages
	fmt.Println("\n\nExample 3: Stock Price Analysis")
	fmt.Println("--------------------------------")
	stockDF := createStockData(pool)
	defer stockDF.Release()

	fmt.Println("\nOriginal Stock Data:")
	printDataFrame(stockDF, 10)

	// Calculate various technical indicators using window functions
	// Note: For production use, apply different windows in separate operations
	indicators, err := stockDF.Window().
		OrderBy("date").
		Rows(14). // 14-day window
		Over(
			core.RollingMin("low").As("min_14d"),  // 14-day low
			core.RollingMax("high").As("max_14d"), // 14-day high
		)
	if err != nil {
		log.Fatal(err)
	}
	defer indicators.Release()

	// Add cumulative and lag functions
	final, err := indicators.Window().
		OrderBy("date").
		Over(
			core.CumMax("high").As("all_time_high"), // All-time high
			core.CumMin("low").As("all_time_low"),   // All-time low
			core.Lag("close", 1).As("prev_close"),   // Previous close
			core.Lead("close", 1).As("next_close"),  // Next close
		)
	if err != nil {
		log.Fatal(err)
	}
	defer final.Release()

	fmt.Println("\nStock Data with Technical Indicators:")
	printDataFrame(final, 25)

	// Example 4: Period-over-Period Comparison
	fmt.Println("\n\nExample 4: Month-over-Month Revenue Analysis")
	fmt.Println("--------------------------------------------")
	revenueDF := createMonthlyRevenueData(pool)
	defer revenueDF.Release()

	fmt.Println("\nOriginal Monthly Revenue Data:")
	printDataFrame(revenueDF, 12)

	// Calculate period-over-period changes
	analysis, err := revenueDF.Window().
		PartitionBy("product").
		OrderBy("month").
		Rows(3).
		Over(
			core.Lag("revenue", 1).As("prev_month_revenue"),
			core.Lag("revenue", 12).As("year_ago_revenue"),
			core.CumSum("revenue").As("ytd_revenue"),
			core.RollingMean("revenue").As("3month_avg"),
		)
	if err != nil {
		log.Fatal(err)
	}
	defer analysis.Release()

	fmt.Println("\nRevenue Analysis with Period Comparisons:")
	printDataFrame(analysis, 15)

	fmt.Println("\n=== All Examples Completed Successfully ===")
}

// createSalesData creates sample sales data for demonstration
func createSalesData(pool memory.Allocator) *core.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "region", Type: arrow.BinaryTypes.String},
			{Name: "date", Type: arrow.PrimitiveTypes.Int64}, // Days since epoch
			{Name: "sales", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	regionBuilder := array.NewStringBuilder(pool)
	defer regionBuilder.Release()
	dateBuilder := array.NewInt64Builder(pool)
	defer dateBuilder.Release()
	salesBuilder := array.NewFloat64Builder(pool)
	defer salesBuilder.Release()

	// Generate sample data: 30 days of sales for 2 regions
	baseDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix() / 86400

	for region := 0; region < 2; region++ {
		regionName := []string{"North", "South"}[region]
		baseAmount := []float64{100.0, 150.0}[region]

		for day := 0; day < 30; day++ {
			regionBuilder.Append(regionName)
			dateBuilder.Append(baseDate + int64(day))
			// Simulate varying sales with trend
			sales := baseAmount + float64(day)*2.0 + float64((day%7)*10)
			salesBuilder.Append(sales)
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		regionBuilder.NewArray(),
		dateBuilder.NewArray(),
		salesBuilder.NewArray(),
	}, int64(60))

	return core.NewDataFrame(record)
}

// createEmployeeData creates sample employee data for ranking demonstration
func createEmployeeData(pool memory.Allocator) *core.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "department", Type: arrow.BinaryTypes.String},
			{Name: "salary", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	nameBuilder := array.NewStringBuilder(pool)
	defer nameBuilder.Release()
	deptBuilder := array.NewStringBuilder(pool)
	defer deptBuilder.Release()
	salaryBuilder := array.NewFloat64Builder(pool)
	defer salaryBuilder.Release()

	// Sample employee data
	employees := []struct {
		name   string
		dept   string
		salary float64
	}{
		{"Alice", "Engineering", 120000},
		{"Bob", "Engineering", 115000},
		{"Carol", "Engineering", 130000},
		{"David", "Engineering", 115000}, // Tie with Bob
		{"Eve", "Sales", 95000},
		{"Frank", "Sales", 105000},
		{"Grace", "Sales", 100000},
		{"Henry", "Marketing", 85000},
		{"Ivy", "Marketing", 90000},
		{"Jack", "Marketing", 88000},
	}

	for _, emp := range employees {
		nameBuilder.Append(emp.name)
		deptBuilder.Append(emp.dept)
		salaryBuilder.Append(emp.salary)
	}

	record := array.NewRecord(schema, []arrow.Array{
		nameBuilder.NewArray(),
		deptBuilder.NewArray(),
		salaryBuilder.NewArray(),
	}, int64(len(employees)))

	return core.NewDataFrame(record)
}

// createStockData creates sample stock price data for technical analysis
func createStockData(pool memory.Allocator) *core.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "date", Type: arrow.PrimitiveTypes.Int64},
			{Name: "open", Type: arrow.PrimitiveTypes.Float64},
			{Name: "high", Type: arrow.PrimitiveTypes.Float64},
			{Name: "low", Type: arrow.PrimitiveTypes.Float64},
			{Name: "close", Type: arrow.PrimitiveTypes.Float64},
			{Name: "volume", Type: arrow.PrimitiveTypes.Int64},
		},
		nil,
	)

	dateBuilder := array.NewInt64Builder(pool)
	defer dateBuilder.Release()
	openBuilder := array.NewFloat64Builder(pool)
	defer openBuilder.Release()
	highBuilder := array.NewFloat64Builder(pool)
	defer highBuilder.Release()
	lowBuilder := array.NewFloat64Builder(pool)
	defer lowBuilder.Release()
	closeBuilder := array.NewFloat64Builder(pool)
	defer closeBuilder.Release()
	volumeBuilder := array.NewInt64Builder(pool)
	defer volumeBuilder.Release()

	// Generate 30 days of simulated stock data
	baseDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix() / 86400
	basePrice := 100.0

	for day := 0; day < 30; day++ {
		dateBuilder.Append(baseDate + int64(day))

		// Simulate price movement with trend and volatility
		trend := float64(day) * 0.5
		volatility := float64((day % 5) * 2)
		open := basePrice + trend + volatility
		high := open + 2.0 + float64(day%3)
		low := open - 2.0 - float64(day%3)
		close := open + float64(day%7) - 3.0

		openBuilder.Append(open)
		highBuilder.Append(high)
		lowBuilder.Append(low)
		closeBuilder.Append(close)
		volumeBuilder.Append(int64(1000000 + day*10000))
	}

	record := array.NewRecord(schema, []arrow.Array{
		dateBuilder.NewArray(),
		openBuilder.NewArray(),
		highBuilder.NewArray(),
		lowBuilder.NewArray(),
		closeBuilder.NewArray(),
		volumeBuilder.NewArray(),
	}, 30)

	return core.NewDataFrame(record)
}

// createMonthlyRevenueData creates sample monthly revenue data for period comparison
func createMonthlyRevenueData(pool memory.Allocator) *core.DataFrame {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "product", Type: arrow.BinaryTypes.String},
			{Name: "month", Type: arrow.PrimitiveTypes.Int64}, // YYYYMM format
			{Name: "revenue", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	productBuilder := array.NewStringBuilder(pool)
	defer productBuilder.Release()
	monthBuilder := array.NewInt64Builder(pool)
	defer monthBuilder.Release()
	revenueBuilder := array.NewFloat64Builder(pool)
	defer revenueBuilder.Release()

	// Generate 12 months of revenue data for 2 products
	products := []string{"Product A", "Product B"}
	baseRevenue := []float64{50000.0, 75000.0}

	for prodIdx, product := range products {
		for month := 1; month <= 12; month++ {
			productBuilder.Append(product)
			monthBuilder.Append(int64(202500 + month)) // 202501, 202502, etc.

			// Simulate seasonal revenue pattern
			seasonality := 1.0 + 0.2*float64(month%4)
			revenue := baseRevenue[prodIdx] * seasonality
			revenueBuilder.Append(revenue)
		}
	}

	record := array.NewRecord(schema, []arrow.Array{
		productBuilder.NewArray(),
		monthBuilder.NewArray(),
		revenueBuilder.NewArray(),
	}, 24)

	return core.NewDataFrame(record)
}

// printDataFrame prints a DataFrame in a formatted table
func printDataFrame(df *core.DataFrame, maxRows int) {
	if df == nil {
		fmt.Println("DataFrame is nil")
		return
	}

	// Print column names
	names := df.ColumnNames()
	for i, name := range names {
		if i > 0 {
			fmt.Print(" | ")
		}
		fmt.Printf("%-15s", name)
	}
	fmt.Println()

	// Print separator
	for i := range names {
		if i > 0 {
			fmt.Print("-+-")
		}
		fmt.Print("---------------")
	}
	fmt.Println()

	// Print rows
	numRows := int(df.NumRows())
	if numRows > maxRows {
		numRows = maxRows
	}

	for row := 0; row < numRows; row++ {
		for colIdx, name := range names {
			if colIdx > 0 {
				fmt.Print(" | ")
			}

			series, err := df.Column(name)
			if err != nil {
				fmt.Printf("%-15s", "ERROR")
				continue
			}

			arr := series.Array()
			if arr.IsNull(row) {
				fmt.Printf("%-15s", "null")
			} else {
				switch a := arr.(type) {
				case *array.String:
					fmt.Printf("%-15s", a.Value(row))
				case *array.Int64:
					fmt.Printf("%-15d", a.Value(row))
				case *array.Float64:
					fmt.Printf("%-15.2f", a.Value(row))
				default:
					fmt.Printf("%-15v", "?")
				}
			}
			series.Release()
		}
		fmt.Println()
	}

	if int(df.NumRows()) > maxRows {
		fmt.Printf("... (%d more rows)\n", int(df.NumRows())-maxRows)
	}
}
