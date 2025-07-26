# GopherFrame API Documentation

## Table of Contents
- [Quick Start](#quick-start)
- [Core Concepts](#core-concepts)
- [DataFrame Operations](#dataframe-operations)
- [Expression System](#expression-system)
- [I/O Operations](#io-operations)
- [GroupBy and Aggregations](#groupby-and-aggregations)
- [Performance Guide](#performance-guide)
- [Migration from Other Libraries](#migration-from-other-libraries)

## Quick Start

### Installation

```bash
go get github.com/felixgeelhaar/GopherFrame
```

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
    // Read a Parquet file
    df, err := gf.ReadParquet("data.parquet")
    if err != nil {
        log.Fatal(err)
    }
    defer df.Release() // Always release resources
    
    // Basic operations
    result := df.
        Filter(gf.Col("sales").Gt(gf.Lit(1000.0))).
        WithColumn("profit_margin", gf.Col("profit").Div(gf.Col("sales"))).
        Select("product", "sales", "profit_margin").
        Sort("sales", false) // descending
    defer result.Release()
    
    fmt.Printf("Result: %d rows, %d columns\n", result.NumRows(), result.NumCols())
}
```

## Core Concepts

### DataFrames

A DataFrame is GopherFrame's core data structure, backed by Apache Arrow for high performance:

```go
// Creating DataFrames
df, err := gf.ReadParquet("file.parquet")
df, err := gf.ReadCSV("file.csv")
df := gf.NewDataFrame(arrowRecord) // From Arrow Record

// Always release when done
defer df.Release()
```

### Memory Management

GopherFrame uses Apache Arrow's memory management. **Always call `Release()` on DataFrames** when done:

```go
df := gf.ReadParquet("data.parquet")
defer df.Release() // Critical for preventing memory leaks

// For chained operations
result := df.Filter(...).Select(...).Sort(...)
defer result.Release() // Only need to release the final result
```

### Type System

GopherFrame leverages Arrow's type system for high performance:

- **Numeric**: `int64`, `float64`, `int32`, `uint64`
- **String**: `string`, `binary`, `large_string`
- **Temporal**: `date32`, `date64`, `timestamp`
- **Boolean**: `bool`

## DataFrame Operations

### Inspection

```go
// Basic info
rows := df.NumRows()         // int64
cols := df.NumCols()         // int64
names := df.ColumnNames()    // []string
schema := df.Schema()        // *arrow.Schema

// Column access
series, err := df.Column("column_name")  // *Series
series, err := df.ColumnAt(0)            // *Series by index

// Data preview
fmt.Println(df.String()) // Print DataFrame contents
```

### Selection and Projection

```go
// Select specific columns
subset := df.Select("name", "age", "salary")
defer subset.Release()

// Select with expressions
calculated := df.Select(
    "name",
    gf.Col("salary").Mul(gf.Lit(1.1)).As("salary_increase"),
    gf.Col("age").Add(gf.Lit(1)).As("next_year_age"),
)
defer calculated.Release()
```

### Filtering

```go
// Simple filters
adults := df.Filter(gf.Col("age").Gt(gf.Lit(18)))
defer adults.Release()

// Complex filters with logical operations
qualified := df.Filter(
    gf.Col("age").Gt(gf.Lit(25)).
    And(gf.Col("salary").Gt(gf.Lit(50000.0))).
    And(gf.Col("department").Eq(gf.Lit("Engineering"))),
)
defer qualified.Release()

// String operations
engineers := df.Filter(gf.Col("title").Contains(gf.Lit("Engineer")))
defer engineers.Release()
```

### Adding Columns

```go
// Add computed columns
enriched := df.WithColumn("bonus", gf.Col("salary").Mul(gf.Lit(0.1)))
defer enriched.Release()

// Multiple columns
enhanced := df.
    WithColumn("annual_salary", gf.Col("monthly_salary").Mul(gf.Lit(12))).
    WithColumn("tax_rate", gf.Lit(0.25)).
    WithColumn("net_salary", gf.Col("annual_salary").Mul(gf.Lit(0.75)))
defer enhanced.Release()
```

### Sorting

```go
// Single column sort
sorted := df.Sort("salary", false) // false = descending
defer sorted.Release()

// Multi-column sort
multiSort := df.SortMultiple([]gf.SortKey{
    {Column: "department", Ascending: true},
    {Column: "salary", Ascending: false},
})
defer multiSort.Release()
```

### Joins

```go
// Inner join
employees, _ := gf.ReadParquet("employees.parquet")
departments, _ := gf.ReadParquet("departments.parquet")
defer employees.Release()
defer departments.Release()

joined := employees.Join(departments, "dept_id", "id", gf.InnerJoin)
defer joined.Release()

// Left join
leftJoined := employees.Join(departments, "dept_id", "id", gf.LeftJoin)
defer leftJoined.Release()
```

## Expression System

### Column References

```go
// Simple column reference
col := gf.Col("column_name")

// Arithmetic operations
sum := gf.Col("a").Add(gf.Col("b"))
difference := gf.Col("revenue").Sub(gf.Col("cost"))
ratio := gf.Col("profit").Div(gf.Col("revenue"))
```

### Literals

```go
// Numeric literals
intLit := gf.Lit(42)
floatLit := gf.Lit(3.14159)

// String literals
stringLit := gf.Lit("Hello, World!")

// Boolean literals
boolLit := gf.Lit(true)
```

### Comparisons

```go
// Comparison operations
gt := gf.Col("age").Gt(gf.Lit(21))           // Greater than
lt := gf.Col("score").Lt(gf.Lit(100))        // Less than  
eq := gf.Col("status").Eq(gf.Lit("active"))  // Equal
ne := gf.Col("type").Ne(gf.Lit("test"))      // Not equal
```

### String Operations

```go
// String matching
contains := gf.Col("email").Contains(gf.Lit("@gmail.com"))
startsWith := gf.Col("name").StartsWith(gf.Lit("John"))
endsWith := gf.Col("file").EndsWith(gf.Lit(".pdf"))

// Case conversion (if implemented)
upper := gf.Col("name").Upper()
lower := gf.Col("email").Lower()
```

### Date/Time Operations

```go
// Date extraction
year := gf.Col("created_date").Year()
month := gf.Col("created_date").Month()
day := gf.Col("created_date").Day()
dayOfWeek := gf.Col("created_date").DayOfWeek()

// Date arithmetic
nextWeek := gf.Col("start_date").AddDays(gf.Lit(7))
nextMonth := gf.Col("start_date").AddMonths(gf.Lit(1))
nextYear := gf.Col("start_date").AddYears(gf.Lit(1))

// Date truncation
monthStart := gf.Col("timestamp").DateTrunc("month")
dayStart := gf.Col("timestamp").DateTrunc("day")
```

## I/O Operations

### Parquet Files

```go
// Reading Parquet (recommended for performance)
df, err := gf.ReadParquet("data.parquet")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Writing Parquet
err = gf.WriteParquet(df, "output.parquet")
if err != nil {
    log.Fatal(err)
}
```

### CSV Files

```go
// Reading CSV
df, err := gf.ReadCSV("data.csv")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Writing CSV
err = gf.WriteCSV(df, "output.csv")
if err != nil {
    log.Fatal(err)
}
```

### Arrow IPC

```go
// Reading Arrow IPC format
df, err := gf.ReadArrowIPC("data.arrow")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Writing Arrow IPC
err = gf.WriteArrowIPC(df, "output.arrow")
if err != nil {
    log.Fatal(err)
}
```

## GroupBy and Aggregations

### Basic GroupBy

```go
// Group by single column
grouped := df.GroupBy("department").Agg(gf.Sum("salary"))
defer grouped.Release()

// Group by multiple columns
multiGrouped := df.GroupBy("department", "level").Agg(
    gf.Mean("salary"),
    gf.Count("employee_id"),
)
defer multiGrouped.Release()
```

### Aggregation Functions

```go
// Available aggregations
result := df.GroupBy("category").Agg(
    gf.Sum("amount").As("total_amount"),
    gf.Mean("price").As("avg_price"),
    gf.Count("item_id").As("item_count"),
    gf.Min("date").As("earliest_date"),
    gf.Max("date").As("latest_date"),
)
defer result.Release()
```

### Advanced Aggregations

```go
// Multiple aggregations on same column
stats := df.GroupBy("product").Agg(
    gf.Sum("sales"),
    gf.Mean("sales").As("avg_sales"),
    gf.Min("sales").As("min_sales"),
    gf.Max("sales").As("max_sales"),
    gf.Count("sales").As("sales_count"),
)
defer stats.Release()
```

## Performance Guide

### Best Practices

1. **Always Release DataFrames**
```go
df := gf.ReadParquet("data.parquet")
defer df.Release() // Critical for memory management
```

2. **Use Parquet for Best Performance**
```go
// Preferred: Fast, compressed, columnar
df, _ := gf.ReadParquet("data.parquet")

// Slower: Row-based format
df, _ := gf.ReadCSV("data.csv")
```

3. **Chain Operations for Efficiency**
```go
// Efficient: Single pass through data
result := df.
    Filter(gf.Col("active").Eq(gf.Lit(true))).
    WithColumn("profit", gf.Col("revenue").Sub(gf.Col("cost"))).
    Select("id", "profit").
    Sort("profit", false)
defer result.Release()
```

4. **Filter Early**
```go
// Good: Filter reduces subsequent processing
result := df.
    Filter(gf.Col("date").Gt(gf.Lit("2023-01-01"))). // Filter first
    GroupBy("category").
    Agg(gf.Sum("amount"))
```

5. **Use Appropriate Data Types**
```go
// Arrow types are optimized for performance
// int64, float64, string are most efficient
```

### Performance Characteristics

Based on benchmarks (Apple M1):

| Operation | 1K rows/sec | 10K rows/sec | 100K rows/sec |
|-----------|-------------|--------------|---------------|
| Filter    | 28,142      | 3,092        | 347           |
| GroupBy   | 24,131      | 2,878        | 284           |
| Sort      | ~20,000     | ~2,000       | ~200          |
| I/O Read  | 7,556       | 1,857        | ~186          |

### Memory Usage

- **DataFrame Creation**: ~52KB per 1K rows
- **Filter Operations**: ~39KB per 1K rows  
- **Low Allocation Count**: 60-200 allocations regardless of size
- **Linear Memory Scaling**: Predictable memory usage

## Migration from Other Libraries

### From Standard Library

```go
// Before: maps and slices
data := []map[string]interface{}{
    {"name": "John", "age": 30, "salary": 50000},
    {"name": "Jane", "age": 25, "salary": 60000},
}

// After: GopherFrame
df, _ := gf.ReadParquet("employees.parquet")
defer df.Release()

filtered := df.Filter(gf.Col("age").Gt(gf.Lit(25)))
defer filtered.Release()
```

### From Gota

```go
// Gota style
import "github.com/go-gota/gota/dataframe"

gotaDF := dataframe.LoadFromCSV("data.csv")
filtered := gotaDF.Filter(dataframe.F{"age", series.Greater, 25})

// GopherFrame style (similar but more performant)
import gf "github.com/felixgeelhaar/GopherFrame"

df, _ := gf.ReadCSV("data.csv")
defer df.Release()
filtered := df.Filter(gf.Col("age").Gt(gf.Lit(25)))
defer filtered.Release()
```

## Error Handling

### Common Patterns

```go
// Always check errors for I/O
df, err := gf.ReadParquet("data.parquet")
if err != nil {
    return fmt.Errorf("failed to read parquet: %w", err)
}
defer df.Release()

// Column operations can fail
series, err := df.Column("nonexistent_column")
if err != nil {
    log.Printf("Column not found: %v", err)
    return
}

// Validate before operations
if df.NumRows() == 0 {
    log.Println("Warning: empty DataFrame")
    return
}
```

### Resource Cleanup

```go
func processData(filename string) error {
    df, err := gf.ReadParquet(filename)
    if err != nil {
        return err
    }
    defer df.Release() // Cleanup even on error
    
    result := df.Filter(gf.Col("active").Eq(gf.Lit(true)))
    defer result.Release() // Cleanup intermediate results
    
    return gf.WriteParquet(result, "output.parquet")
}
```

## Advanced Usage

### Custom Memory Management

```go
import "github.com/apache/arrow-go/v18/arrow/memory"

// Custom memory allocator for fine-grained control
pool := memory.NewGoAllocator()
// Use with custom DataFrame creation...
```

### Concurrent Operations

```go
// Read-only operations are thread-safe
var wg sync.WaitGroup
results := make([]*gf.DataFrame, 10)

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(index int) {
        defer wg.Done()
        // Each goroutine gets its own filtered DataFrame
        filtered := df.Filter(gf.Col("category").Eq(gf.Lit(fmt.Sprintf("cat_%d", index))))
        results[index] = filtered
        // Remember to release in calling code
    }(i)
}
wg.Wait()

// Clean up results
for _, result := range results {
    if result != nil {
        result.Release()
    }
}
```

This API documentation provides comprehensive coverage of GopherFrame's capabilities with practical examples for production use.