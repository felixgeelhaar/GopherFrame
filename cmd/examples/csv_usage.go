package main

import (
	"fmt"
	"log"
	"os"

	gf "github.com/felixgeelhaar/gopherFrame"
)

func main() {
	// Create a sample DataFrame
	df := createTestDataFrame()
	defer df.Release()

	// Write to CSV
	csvFile := "test_output.csv"
	fmt.Println("Writing DataFrame to CSV...")
	if err := gf.WriteCSV(df, csvFile); err != nil {
		log.Fatalf("Failed to write CSV: %v", err)
	}
	fmt.Printf("CSV written to %s\n", csvFile)

	// Read back from CSV
	fmt.Println("\nReading DataFrame from CSV...")
	readDF, err := gf.ReadCSV(csvFile)
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}
	defer readDF.Release()

	fmt.Printf("Read DataFrame with %d rows and %d columns\n", 
		readDF.NumRows(), readDF.NumCols())
	fmt.Printf("Columns: %v\n", readDF.ColumnNames())

	// Perform operations on the read DataFrame
	fmt.Println("\nPerforming operations on CSV data...")
	result := readDF.
		Filter(gf.Col("score").Gt(gf.Lit(80.0))).
		Select("name", "score")
	
	if result.Err() != nil {
		log.Fatalf("Operation failed: %v", result.Err())
	}
	defer result.Release()

	fmt.Printf("Filtered DataFrame has %d rows\n", result.NumRows())

	// Clean up
	os.Remove(csvFile)
	fmt.Println("\nCSV example completed successfully!")
}

func createTestDataFrame() *gf.DataFrame {
	// This is a helper function to create test data
	// In production, you would load from real sources
	df, err := gf.ReadParquet("../../internal/testdata/sample.parquet")
	if err != nil {
		// Fallback: create manually if test file doesn't exist
		log.Println("Creating sample data manually...")
		// Implementation would go here
		// For now, we'll just note this is where data creation would happen
		panic("Sample data file not found - please create test data first")
	}
	return df
}