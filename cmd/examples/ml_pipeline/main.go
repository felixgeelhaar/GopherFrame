// Package main demonstrates a complete ML inference preprocessing pipeline using GopherFrame.
// This example shows how to prepare data for ML model inference in a Go service,
// using the same transformations that were applied during Python training.
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
	fmt.Println("=== ML Inference Pipeline Demo ===")
	fmt.Println()

	// 1. Load raw inference data
	fmt.Println("1. Loading raw inference data...")
	rawDF := createInferenceData()
	defer rawDF.Release()
	fmt.Printf("   %d samples, %d features\n", rawDF.NumRows(), rawDF.NumCols())

	// 2. Data validation
	fmt.Println("\n2. Validating input data...")
	validation := rawDF.Validate(
		gf.NotNull("age"),
		gf.NotNull("income"),
		gf.Positive("age"),
		gf.Positive("income"),
		gf.InRange("credit_score", 300, 850),
	)
	if !validation.Valid {
		log.Printf("   Warning: %d validation issues\n", len(validation.Violations))
	} else {
		fmt.Println("   All validations passed")
	}

	// 3. Feature engineering with UDFs
	fmt.Println("\n3. Engineering features...")

	// Log transform income (common ML preprocessing step)
	withLogIncome := rawDF.WithColumn("log_income",
		gf.ScalarUDF([]string{"income"}, arrow.PrimitiveTypes.Float64,
			func(row map[string]interface{}) (interface{}, error) {
				income := row["income"].(float64)
				if income <= 0 {
					return nil, nil
				}
				return math.Log(income), nil
			},
		),
	)
	defer withLogIncome.Release()

	// Debt-to-income ratio
	withDTI := withLogIncome.WithColumn("dti_ratio",
		gf.ScalarUDF([]string{"debt", "income"}, arrow.PrimitiveTypes.Float64,
			func(row map[string]interface{}) (interface{}, error) {
				debt := row["debt"].(float64)
				income := row["income"].(float64)
				if income == 0 {
					return nil, nil
				}
				return debt / income, nil
			},
		),
	)
	defer withDTI.Release()

	// Age bucket (categorical encoding)
	withAgeBucket := withDTI.WithColumn("age_bucket",
		gf.ScalarUDF([]string{"age"}, arrow.PrimitiveTypes.Float64,
			func(row map[string]interface{}) (interface{}, error) {
				age := row["age"].(float64)
				switch {
				case age < 25:
					return float64(0), nil
				case age < 35:
					return float64(1), nil
				case age < 50:
					return float64(2), nil
				default:
					return float64(3), nil
				}
			},
		),
	)
	defer withAgeBucket.Release()

	fmt.Printf("   Created %d features\n", withAgeBucket.NumCols())

	// 4. Outlier detection
	fmt.Println("\n4. Detecting outliers...")
	outliers, err := rawDF.DetectOutliersIQR("income", 1.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Found %d income outliers\n", outliers.Count)

	// 5. Feature selection
	fmt.Println("\n5. Selecting model features...")
	features := withAgeBucket.Select("age", "log_income", "credit_score", "dti_ratio", "age_bucket")
	defer features.Release()
	fmt.Printf("   Final feature matrix: %d x %d\n", features.NumRows(), features.NumCols())

	// 6. Data profile
	fmt.Println("\n6. Feature statistics:")
	fmt.Println(features.DescribeString())

	// 7. Export for model inference
	fmt.Println("7. Ready for inference:")
	fmt.Println("   - Pass to TensorFlow Serving via Arrow Flight")
	fmt.Println("   - Write to Parquet for batch prediction")
	fmt.Println("   - Stream via gRPC to model service")

	fmt.Println("\n=== ML Pipeline Complete ===")
}

func createInferenceData() *gf.DataFrame {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "age", Type: arrow.PrimitiveTypes.Float64},
		{Name: "income", Type: arrow.PrimitiveTypes.Float64},
		{Name: "debt", Type: arrow.PrimitiveTypes.Float64},
		{Name: "credit_score", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	ages := []float64{25, 34, 45, 52, 28, 61, 39, 22}
	incomes := []float64{50000, 75000, 120000, 95000, 42000, 150000, 88000, 35000}
	debts := []float64{15000, 25000, 45000, 30000, 20000, 10000, 35000, 18000}
	scores := []float64{720, 680, 750, 700, 650, 800, 710, 620}

	ab := array.NewFloat64Builder(pool)
	ab.AppendValues(ages, nil)
	aa := ab.NewArray()
	ab.Release()

	ib := array.NewFloat64Builder(pool)
	ib.AppendValues(incomes, nil)
	ia := ib.NewArray()
	ib.Release()

	db := array.NewFloat64Builder(pool)
	db.AppendValues(debts, nil)
	da := db.NewArray()
	db.Release()

	sb := array.NewFloat64Builder(pool)
	sb.AppendValues(scores, nil)
	sa := sb.NewArray()
	sb.Release()

	record := array.NewRecord(schema, []arrow.Array{aa, ia, da, sa}, int64(len(ages)))
	return gf.NewDataFrame(record)
}
