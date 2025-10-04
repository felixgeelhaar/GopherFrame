# GopherFrame API Reference

**Version:** 1.0.0
**Last Updated:** October 4, 2025

Complete API reference for GopherFrame - a production-ready DataFrame library for Go powered by Apache Arrow.

## Table of Contents

- [Data Structures](#data-structures)
  - [DataFrame](#dataframe)
  - [Series](#series)
  - [Expression](#expression)
- [I/O Operations](#io-operations)
  - [Reading Data](#reading-data)
  - [Writing Data](#writing-data)
- [DataFrame Operations](#dataframe-operations)
  - [Selection & Projection](#selection--projection)
  - [Filtering](#filtering)
  - [Column Operations](#column-operations)
  - [Sorting](#sorting)
  - [Joins](#joins)
- [Aggregations](#aggregations)
  - [GroupBy](#groupby)
  - [Basic Aggregations](#basic-aggregations)
  - [Statistical Aggregations](#statistical-aggregations)
- [Window Functions](#window-functions)
  - [Analytical Functions](#analytical-functions)
  - [Rolling Aggregations](#rolling-aggregations)
  - [Cumulative Operations](#cumulative-operations)
- [Expressions](#expressions)
  - [Literal Values](#literal-values)
  - [Column References](#column-references)
  - [Arithmetic Operations](#arithmetic-operations)
  - [Comparison Operations](#comparison-operations)
  - [Logical Operations](#logical-operations)
  - [String Operations](#string-operations)
  - [Temporal Operations](#temporal-operations)
- [Memory Management](#memory-management)

---

## Data Structures

### DataFrame

The core data structure representing a two-dimensional columnar table backed by Apache Arrow.

#### Creation

```go
// From Parquet file
df, err := gf.ReadParquet("data.parquet")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// From CSV file
df, err := gf.ReadCSV("data.csv")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// From Arrow Record
record := createArrowRecord() // your arrow.Record
df := gf.NewDataFrame(record)
defer df.Release()
```

#### Properties

```go
// Get number of rows
numRows := df.NumRows() // returns int64

// Get number of columns
numCols := df.NumCols() // returns int64

// Get column names
columnNames := df.ColumnNames() // returns []string

// Check if column exists
hasColumn := df.HasColumn("column_name") // returns bool

// Get error from chained operations
err := df.Err() // returns error
```

#### Memory Management

```go
// Release DataFrame resources
df.Release()

// IMPORTANT: Always defer Release() immediately after creation
df := gf.ReadParquet("data.parquet")
defer df.Release()
```

### Series

A single-column data structure representing a column from a DataFrame.

```go
// Get column from DataFrame
series, err := df.Column("column_name")
if err != nil {
    log.Fatal(err)
}
defer series.Release()

// Access underlying Arrow array
arr := series.Array() // returns arrow.Array

// Release Series resources
series.Release()
```

### Expression

Type-safe expression system for DataFrame operations.

```go
// Column reference
expr := df.Col("column_name")

// Literal value
expr := gf.Lit(42)
expr := gf.Lit("hello")
expr := gf.Lit(3.14)
```

---

## I/O Operations

### Reading Data

#### ReadParquet

```go
df, err := gf.ReadParquet("data.parquet")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Performance: 5-10M rows/second
// Features: Full type support, compression, metadata
```

#### ReadCSV

```go
df, err := gf.ReadCSV("data.csv")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Features: Header detection, type inference
```

#### ReadArrowIPC

```go
df, err := gf.ReadArrowIPC("data.arrow")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Features: Zero-copy deserialization, full Arrow type support
```

### Writing Data

#### WriteParquet

```go
err := gf.WriteParquet(df, "output.parquet")
if err != nil {
    log.Fatal(err)
}

// Features: Snappy compression, full type support
```

#### WriteCSV

```go
err := gf.WriteCSV(df, "output.csv")
if err != nil {
    log.Fatal(err)
}
```

#### WriteArrowIPC

```go
err := gf.WriteArrowIPC(df, "output.arrow")
if err != nil {
    log.Fatal(err)
}

// Features: Zero-copy serialization, smallest file size
```

---

## DataFrame Operations

### Selection & Projection

#### Select

Select specific columns from DataFrame.

```go
// Select single column
result := df.Select("name")
defer result.Release()

// Select multiple columns
result := df.Select("name", "age", "email")
defer result.Release()

// Performance: O(1) - zero-copy operation (~700ns)
```

### Filtering

#### Filter

Filter rows based on conditions.

```go
// Simple filter
result := df.Filter(df.Col("age").Gt(gf.Lit(18)))
defer result.Release()

// Complex filter with AND
result := df.Filter(
    df.Col("age").Gt(gf.Lit(18)).
        And(df.Col("country").Eq(gf.Lit("US"))),
)
defer result.Release()

// Complex filter with OR
result := df.Filter(
    df.Col("status").Eq(gf.Lit("active")).
        Or(df.Col("status").Eq(gf.Lit("pending"))),
)
defer result.Release()

// Performance: O(n) linear scan with vectorized operations
```

### Column Operations

#### WithColumn

Add or replace a column with computed values.

```go
// Add new column
result := df.WithColumn("profit",
    df.Col("revenue").Sub(df.Col("cost")),
)
defer result.Release()

// Replace existing column
result := df.WithColumn("price",
    df.Col("price").Mul(gf.Lit(1.1)), // 10% increase
)
defer result.Release()

// Chain multiple WithColumn operations
result := df.
    WithColumn("tax", df.Col("amount").Mul(gf.Lit(0.08))).
    WithColumn("total", df.Col("amount").Add(df.Col("tax")))
defer result.Release()
```

### Sorting

#### Sort

Sort DataFrame by column(s).

```go
// Sort by single column, ascending
result := df.Sort("age", true)
defer result.Release()

// Sort by single column, descending
result := df.Sort("salary", false)
defer result.Release()

// Sort by multiple columns
result := df.SortMultiple(
    []string{"department", "salary"},
    []bool{true, false}, // department ascending, salary descending
)
defer result.Release()

// Performance: O(n log n) with Arrow compute kernels
```

### Joins

#### InnerJoin

Inner join - returns only matching rows.

```go
// Join users and orders
users, _ := gf.ReadParquet("users.parquet")
defer users.Release()

orders, _ := gf.ReadParquet("orders.parquet")
defer orders.Release()

result, err := users.InnerJoin(orders, "user_id", "customer_id")
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Performance: O(n+m) hash-based join
// Benchmark: 85µs for 1K rows, 1.08ms for 10K rows
```

#### LeftJoin

Left outer join - preserves all left-side rows.

```go
result, err := users.LeftJoin(profiles, "id", "user_id")
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Rows without matches get null values for right-side columns
// Performance: O(n+m) hash-based join
// Benchmark: 103µs for 1K rows, 1.16ms for 10K rows
```

---

## Aggregations

### GroupBy

Group DataFrame by column(s) for aggregation.

```go
// Single column grouping
grouped := df.GroupBy("category")

// Multiple column grouping
grouped := df.GroupBy("region", "product")

// Apply aggregation
result := grouped.Agg(
    gf.Sum("revenue").As("total_revenue"),
    gf.Mean("price").As("avg_price"),
    gf.Count("order_id").As("order_count"),
)
defer result.Release()
```

### Basic Aggregations

#### Sum

```go
result := df.GroupBy("category").Agg(
    gf.Sum("revenue"),
)
defer result.Release()

// Default alias: "revenue_sum"
// Custom alias: gf.Sum("revenue").As("total")
```

#### Mean

```go
result := df.GroupBy("category").Agg(
    gf.Mean("price"),
)
defer result.Release()

// Default alias: "price_mean"
```

#### Count

```go
result := df.GroupBy("category").Agg(
    gf.Count("id"),
)
defer result.Release()

// Default alias: "id_count"
// Counts non-null values
```

#### Min / Max

```go
result := df.GroupBy("category").Agg(
    gf.Min("price"),
    gf.Max("price"),
)
defer result.Release()

// Default aliases: "price_min", "price_max"
```

### Statistical Aggregations

#### Percentile

Calculate any percentile (0.0-1.0).

```go
result := df.GroupBy("region").Agg(
    gf.Percentile("response_time", 0.50).As("p50"),
    gf.Percentile("response_time", 0.95).As("p95"),
    gf.Percentile("response_time", 0.99).As("p99"),
)
defer result.Release()

// Method: Linear interpolation
// Range: p ∈ [0.0, 1.0]
// Error if p outside range
```

#### Median

Calculate median (50th percentile).

```go
result := df.GroupBy("category").Agg(
    gf.Median("price"),
)
defer result.Release()

// Equivalent to Percentile(column, 0.5)
// Robust to outliers
```

#### Mode

Find most frequent value.

```go
result := df.GroupBy("store").Agg(
    gf.Mode("payment_method").As("most_common_payment"),
)
defer result.Release()

// Complexity: O(n) frequency counting
// Ties: Returns one value (deterministic)
```

#### Correlation

Calculate Pearson correlation coefficient between two columns.

```go
result := df.GroupBy("market").Agg(
    gf.Correlation("ad_spend", "revenue").As("correlation"),
)
defer result.Release()

// Range: [-1.0, 1.0]
//  1.0 = perfect positive correlation
//  0.0 = no correlation
// -1.0 = perfect negative correlation
// Returns null if insufficient data or no variance
```

---

## Window Functions

Window functions operate over a "window" of rows related to the current row.

```go
// Window builder
window := df.Window()

// With partitioning (groups)
window := df.Window().PartitionBy("category")

// With ordering
window := df.Window().OrderBy("date")

// With both
window := df.Window().PartitionBy("category").OrderBy("date")

// With frame specification (for rolling)
window := df.Window().Rows(7) // 7-row window
```

### Analytical Functions

#### RowNumber

Sequential numbering within each partition.

```go
result, err := df.Window().
    PartitionBy("department").
    OrderBy("hire_date").
    Over(gf.RowNumber().As("employee_number"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Result: 1, 2, 3, ... within each partition
// Benchmark: 381µs for 1K rows, 5.28ms for 10K rows
```

#### Rank

Ranking with gaps for ties.

```go
result, err := df.Window().
    PartitionBy("department").
    OrderBy("salary").
    Over(gf.Rank().As("salary_rank"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Result: 1, 2, 2, 4 (ties get same rank, next rank skipped)
// Benchmark: 351µs for 1K rows, 3.62ms for 10K rows
```

#### DenseRank

Ranking without gaps.

```go
result, err := df.Window().
    PartitionBy("department").
    OrderBy("salary").
    Over(gf.DenseRank().As("dense_rank"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Result: 1, 2, 2, 3 (no gaps)
```

#### Lag

Access previous row value.

```go
result, err := df.Window().
    OrderBy("date").
    Over(gf.Lag("sales", 1).As("prev_sales"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// With default value for first row
result, err := df.Window().
    OrderBy("date").
    Over(gf.Lag("sales", 1).Default(int64(0)).As("prev_sales"))

// offset: number of rows to look back
// Benchmark: 374µs for 1K rows, 4.01ms for 10K rows
```

#### Lead

Access next row value.

```go
result, err := df.Window().
    OrderBy("date").
    Over(gf.Lead("sales", 1).As("next_sales"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// With default value for last row
result, err := df.Window().
    OrderBy("date").
    Over(gf.Lead("sales", 1).Default(int64(0)).As("next_sales"))

// offset: number of rows to look ahead
```

### Rolling Aggregations

Rolling window aggregations with specified window size.

#### RollingSum

```go
result, err := df.Window().
    Rows(7). // 7-row window
    Over(gf.RollingSum("revenue").As("revenue_7d"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Benchmark: 152µs for 1K rows, 1.72ms for 10K rows
```

#### RollingMean

```go
result, err := df.Window().
    Rows(30).
    Over(gf.RollingMean("price").As("price_30d_avg"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Benchmark: 158µs for 1K rows, 1.79ms for 10K rows
```

#### RollingMin / RollingMax

```go
result, err := df.Window().
    Rows(14).
    Over(
        gf.RollingMin("price").As("min_14d"),
        gf.RollingMax("price").As("max_14d"),
    )
if err != nil {
    log.Fatal(err)
}
defer result.Release()
```

#### RollingCount

```go
result, err := df.Window().
    Rows(7).
    Over(gf.RollingCount("event").As("events_7d"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()
```

### Cumulative Operations

Cumulative aggregations from partition start to current row.

#### CumSum

```go
result, err := df.Window().
    PartitionBy("account").
    OrderBy("date").
    Over(gf.CumSum("amount").As("running_balance"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Benchmark: 117µs for 1K rows, 1.15ms for 10K rows
```

#### CumMax

```go
result, err := df.Window().
    PartitionBy("stock").
    OrderBy("date").
    Over(gf.CumMax("high").As("all_time_high"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()

// Benchmark: 105µs for 1K rows, 1.15ms for 10K rows
```

#### CumMin

```go
result, err := df.Window().
    Over(gf.CumMin("price").As("all_time_low"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()
```

#### CumProd

```go
result, err := df.Window().
    Over(gf.CumProd("multiplier").As("cumulative_product"))
if err != nil {
    log.Fatal(err)
}
defer result.Release()
```

---

## Expressions

Type-safe expression system for DataFrame operations.

### Literal Values

```go
// Integer literal
expr := gf.Lit(42)

// Float literal
expr := gf.Lit(3.14)

// String literal
expr := gf.Lit("hello")

// Boolean literal
expr := gf.Lit(true)
```

### Column References

```go
// Column reference
expr := df.Col("column_name")

// Used in operations
result := df.Filter(df.Col("age").Gt(gf.Lit(18)))
```

### Arithmetic Operations

```go
// Addition
expr := df.Col("a").Add(df.Col("b"))
expr := df.Col("price").Add(gf.Lit(10.0))

// Subtraction
expr := df.Col("revenue").Sub(df.Col("cost"))

// Multiplication
expr := df.Col("quantity").Mul(df.Col("price"))

// Division
expr := df.Col("total").Div(df.Col("count"))
```

### Comparison Operations

```go
// Equal
expr := df.Col("status").Eq(gf.Lit("active"))

// Not Equal
expr := df.Col("status").Ne(gf.Lit("deleted"))

// Greater Than
expr := df.Col("age").Gt(gf.Lit(18))

// Greater Than or Equal
expr := df.Col("score").Ge(gf.Lit(90))

// Less Than
expr := df.Col("price").Lt(gf.Lit(100.0))

// Less Than or Equal
expr := df.Col("quantity").Le(gf.Lit(10))
```

### Logical Operations

```go
// AND
expr := df.Col("age").Gt(gf.Lit(18)).
    And(df.Col("country").Eq(gf.Lit("US")))

// OR
expr := df.Col("status").Eq(gf.Lit("active")).
    Or(df.Col("status").Eq(gf.Lit("pending")))

// NOT (use Ne for negation)
expr := df.Col("deleted").Ne(gf.Lit(true))
```

### String Operations

#### Upper / Lower

```go
// Convert to uppercase
result := df.WithColumn("name_upper",
    df.Col("name").Upper(),
)
defer result.Release()

// Convert to lowercase
result := df.WithColumn("email_lower",
    df.Col("email").Lower(),
)
defer result.Release()
```

#### Trim

```go
// Remove leading and trailing whitespace
result := df.WithColumn("clean_name",
    df.Col("name").Trim(),
)
defer result.Release()

// Remove leading whitespace only
result := df.WithColumn("clean_name",
    df.Col("name").TrimLeft(),
)
defer result.Release()

// Remove trailing whitespace only
result := df.WithColumn("clean_name",
    df.Col("name").TrimRight(),
)
defer result.Release()
```

#### Length

```go
// Get string length
result := df.WithColumn("name_length",
    df.Col("name").Length(),
)
defer result.Release()

// Filter by length
result := df.Filter(
    df.Col("password").Length().Gt(gf.Lit(8)),
)
defer result.Release()
```

#### Match

```go
// Regex pattern matching
result := df.Filter(
    df.Col("email").Match(gf.Lit(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)),
)
defer result.Release()

// Pattern caching for performance
// Invalid regex returns error
```

#### Contains / StartsWith / EndsWith

```go
// Contains substring
result := df.Filter(
    df.Col("text").Contains(gf.Lit("keyword")),
)
defer result.Release()

// Starts with prefix
result := df.Filter(
    df.Col("filename").StartsWith(gf.Lit("report_")),
)
defer result.Release()

// Ends with suffix
result := df.Filter(
    df.Col("filename").EndsWith(gf.Lit(".csv")),
)
defer result.Release()
```

### Temporal Operations

#### Component Extraction

```go
// Extract year
result := df.WithColumn("year",
    df.Col("timestamp").Year(),
)
defer result.Release()

// Extract month (1-12)
result := df.WithColumn("month",
    df.Col("timestamp").Month(),
)
defer result.Release()

// Extract day of month
result := df.WithColumn("day",
    df.Col("timestamp").Day(),
)
defer result.Release()

// Extract hour (0-23)
result := df.WithColumn("hour",
    df.Col("timestamp").Hour(),
)
defer result.Release()

// Extract minute (0-59)
result := df.WithColumn("minute",
    df.Col("timestamp").Minute(),
)
defer result.Release()

// Extract second (0-59)
result := df.WithColumn("second",
    df.Col("timestamp").Second(),
)
defer result.Release()
```

#### Truncation Operations

```go
// Truncate to year start
result := df.WithColumn("year_start",
    df.Col("timestamp").TruncateToYear(),
)
defer result.Release()

// Truncate to month start
result := df.WithColumn("month_start",
    df.Col("timestamp").TruncateToMonth(),
)
defer result.Release()

// Truncate to day start
result := df.WithColumn("day_start",
    df.Col("timestamp").TruncateToDay(),
)
defer result.Release()

// Truncate to hour start
result := df.WithColumn("hour_start",
    df.Col("timestamp").TruncateToHour(),
)
defer result.Release()
```

#### Arithmetic Operations

```go
// Add days
result := df.WithColumn("next_week",
    df.Col("date").AddDays(gf.Lit(7)),
)
defer result.Release()

// Subtract days
result := df.WithColumn("last_week",
    df.Col("date").AddDays(gf.Lit(-7)),
)
defer result.Release()

// Add hours
result := df.WithColumn("plus_2h",
    df.Col("timestamp").AddHours(gf.Lit(2)),
)
defer result.Release()

// Add minutes
result := df.WithColumn("plus_30m",
    df.Col("timestamp").AddMinutes(gf.Lit(30)),
)
defer result.Release()

// Add seconds
result := df.WithColumn("plus_5s",
    df.Col("timestamp").AddSeconds(gf.Lit(5)),
)
defer result.Release()
```

---

## Memory Management

GopherFrame uses reference counting for deterministic resource management.

### Basic Principles

```go
// 1. Always defer Release() immediately after creation
df := gf.ReadParquet("data.parquet")
defer df.Release()

// 2. Release intermediate results
filtered := df.Filter(df.Col("age").Gt(gf.Lit(18)))
defer filtered.Release()

sorted := filtered.Sort("name", true)
defer sorted.Release()

// 3. Check for errors before Release
result := df.Select("invalid_column")
if result.Err() != nil {
    result.Release() // Still release even on error
    log.Fatal(result.Err())
}
defer result.Release()
```

### Memory Limits

```go
import (
    "github.com/apache/arrow-go/v18/arrow/memory"
    gf "github.com/felixgeelhaar/GopherFrame"
)

// Create limited allocator (2GB limit)
allocator := memory.NewLimitedAllocator(
    memory.NewGoAllocator(),
    2 * 1024 * 1024 * 1024, // 2GB
)

// Monitor memory pressure
pressure := allocator.MemoryPressure()
switch pressure {
case memory.Low:
    // < 50% of limit
case memory.Medium:
    // 50-75% of limit
case memory.High:
    // 75-90% of limit
case memory.Critical:
    // > 90% of limit - consider freeing resources
}

// Get current usage
bytesUsed := allocator.BytesAllocated()
```

### Best Practices

```go
// ❌ BAD: Resource leak
func processData() {
    df := gf.ReadParquet("data.parquet")
    // Missing defer df.Release()
    result := df.Filter(df.Col("age").Gt(gf.Lit(18)))
    // Resources leaked!
}

// ✅ GOOD: Proper cleanup
func processData() {
    df := gf.ReadParquet("data.parquet")
    defer df.Release()

    result := df.Filter(df.Col("age").Gt(gf.Lit(18)))
    defer result.Release()

    // Use result...
}

// ✅ GOOD: Chain with cleanup
func processData() {
    df := gf.ReadParquet("data.parquet")
    defer df.Release()

    // Create intermediate results with deferred cleanup
    filtered := df.Filter(df.Col("age").Gt(gf.Lit(18)))
    defer filtered.Release()

    sorted := filtered.Sort("name", true)
    defer sorted.Release()

    final := sorted.Select("name", "age")
    defer final.Release()

    return final
}
```

---

## Performance Benchmarks

All benchmarks run on Apple M1, Darwin/ARM64.

### Core Operations

| Operation | Size | Time | Memory | Allocs |
|-----------|------|------|--------|--------|
| Select | 1K | 800ns | 1.6KB | 16 |
| Select | 10K | 730ns | 1.6KB | 16 |
| Filter | 1K | 43µs | 54KB | 74 |
| Filter | 10K | 398µs | 798KB | 110 |
| Create | 1K | 31µs | 53KB | 68 |
| Create | 10K | 238µs | 789KB | 96 |

### Joins

| Operation | Size | Time | Memory | Allocs |
|-----------|------|------|--------|--------|
| InnerJoin | 1K | 85µs | 165KB | 1092 |
| InnerJoin | 10K | 1.08ms | 2.2MB | 19136 |
| LeftJoin | 1K | 103µs | 250KB | 1100 |
| LeftJoin | 10K | 1.16ms | 2.4MB | 19138 |

### Window Functions

| Operation | Size | Time | Memory | Allocs |
|-----------|------|------|--------|--------|
| RowNumber | 1K | 381µs | 601KB | 3306 |
| RowNumber | 10K | 5.28ms | 5.9MB | 37922 |
| Rank | 1K | 351µs | 601KB | 3306 |
| Rank | 10K | 3.62ms | 5.9MB | 37922 |
| Lag | 1K | 374µs | 623KB | 5397 |
| Lag | 10K | 4.01ms | 6.0MB | 51260 |

### Rolling Aggregations

| Operation | Size | Time | Memory | Allocs |
|-----------|------|------|--------|--------|
| RollingSum | 1K | 152µs | 186KB | 1092 |
| RollingSum | 10K | 1.72ms | 1.9MB | 10180 |
| RollingMean | 1K | 158µs | 186KB | 1092 |
| RollingMean | 10K | 1.79ms | 1.9MB | 10180 |

### Cumulative Aggregations

| Operation | Size | Time | Memory | Allocs |
|-----------|------|------|--------|--------|
| CumSum | 1K | 117µs | 186KB | 1092 |
| CumSum | 10K | 1.15ms | 1.9MB | 10180 |
| CumMax | 1K | 105µs | 186KB | 1092 |
| CumMax | 10K | 1.15ms | 1.9MB | 10180 |

---

## Error Handling

All operations that can fail return an error as the last return value.

```go
// I/O operations
df, err := gf.ReadParquet("data.parquet")
if err != nil {
    log.Fatalf("failed to read parquet: %v", err)
}
defer df.Release()

// Join operations
result, err := left.InnerJoin(right, "id", "user_id")
if err != nil {
    log.Fatalf("join failed: %v", err)
}
defer result.Release()

// Window operations
result, err := df.Window().Over(gf.RowNumber().As("row_num"))
if err != nil {
    log.Fatalf("window function failed: %v", err)
}
defer result.Release()

// Chained operations use Err()
result := df.
    Filter(df.Col("age").Gt(gf.Lit(18))).
    Sort("name", true).
    Select("name", "email")

if err := result.Err(); err != nil {
    result.Release()
    log.Fatalf("operation failed: %v", err)
}
defer result.Release()
```

---

## Type System

GopherFrame supports all Apache Arrow types:

### Numeric Types
- **Int8, Int16, Int32, Int64**: Signed integers
- **Uint8, Uint16, Uint32, Uint64**: Unsigned integers
- **Float32, Float64**: Floating point numbers

### String Types
- **String**: UTF-8 encoded strings
- **Binary**: Arbitrary byte arrays

### Temporal Types
- **Date32, Date64**: Date values
- **Timestamp**: Timestamp with timezone support
- **Time32, Time64**: Time of day
- **Duration**: Time duration

### Boolean Type
- **Boolean**: True/false values

### Complex Types
- **List**: Variable-length lists
- **Struct**: Nested structures
- **Map**: Key-value pairs

---

## Additional Resources

- **[Migration from Pandas](MIGRATION_FROM_PANDAS.md)**: Complete migration guide from pandas
- **[Migration from Polars](MIGRATION_FROM_POLARS.md)**: Complete migration guide from Polars
- **[Migration from Gota](MIGRATION_FROM_GOTA.md)**: Complete migration guide from Gota
- **[Benchmark Results](../BENCHMARKS.md)**: Detailed performance benchmarks
- **[Example Programs](../cmd/examples/)**: Complete example applications
- **[pkg.go.dev](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame)**: Auto-generated API documentation

---

**Last Updated:** October 4, 2025
**Version:** 1.0.0
**License:** Apache 2.0
