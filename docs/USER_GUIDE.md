# GopherFrame User Guide

## Quick Start

```bash
go get github.com/felixgeelhaar/GopherFrame
```

```go
import gf "github.com/felixgeelhaar/GopherFrame"

df, _ := gf.ReadCSV("data.csv")
defer df.Release()

result := df.Filter(df.Col("score").Gt(gf.Lit(90.0)))
defer result.Release()
```

## Core Concepts

### DataFrame

A DataFrame is an immutable table of columns backed by Apache Arrow. Every operation returns a **new** DataFrame — the original is never modified.

```go
df := gf.NewDataFrame(record) // Create from Arrow Record
defer df.Release()             // Always release when done

df.NumRows()      // Row count
df.NumCols()      // Column count
df.ColumnNames()  // All column names
df.HasColumn("x") // Check existence
df.Err()          // Check for errors
```

### Expressions

Expressions describe computations on columns:

```go
df.Col("price")                         // Column reference
gf.Lit(42.0)                            // Literal value
df.Col("price").Mul(gf.Lit(1.08))       // Arithmetic
df.Col("age").Gt(gf.Lit(int64(18)))     // Comparison
df.Col("name").Upper()                  // String ops
```

## Recipes

### ETL Pipeline

```go
raw, _ := gf.ReadCSV("orders.csv")
defer raw.Release()

clean := raw.Filter(raw.Col("amount").Gt(gf.Lit(0.0)))
defer clean.Release()

enriched := clean.WithColumn("tax", clean.Col("amount").Mul(gf.Lit(0.08)))
defer enriched.Release()

summary := enriched.GroupBy("region").Agg(
    gf.Sum("amount").As("total"),
    gf.Mean("amount").As("avg"),
    gf.Count("amount").As("n"),
)
defer summary.Release()

gf.WriteParquet(summary, "summary.parquet")
```

### Window Functions

```go
ranked := df.Window().
    PartitionBy("department").
    OrderBy("salary").
    Over(
        gf.RowNumber().As("rank"),
        gf.RollingMean("salary", 3).As("salary_3m_avg"),
    )
```

### Join Operations

```go
// Single key
result := users.InnerJoin(orders, "user_id", "customer_id")

// Multiple keys
result := left.InnerJoinMulti(right,
    []string{"dept", "role"},
    []string{"department", "position"},
)

// Optimized strategies
result := left.MergeJoin(right, "id", "id")       // Pre-sorted data
result := left.BroadcastJoin(right, "id", "id")    // Small right table
result := left.AutoJoin(right, "id", "id")         // Auto-select best
```

### User-Defined Functions

```go
// Scalar UDF (row by row)
df.WithColumn("bmi", gf.ScalarUDF(
    []string{"weight", "height"},
    arrow.PrimitiveTypes.Float64,
    func(row map[string]interface{}) (interface{}, error) {
        w := row["weight"].(float64)
        h := row["height"].(float64)
        return w / (h * h), nil
    },
))

// Vector UDF (columnar, fast)
df.WithColumn("doubled", gf.VectorUDF(
    []string{"x"},
    arrow.PrimitiveTypes.Float64,
    func(cols map[string]arrow.Array) (arrow.Array, error) {
        // Operate directly on Arrow arrays
    },
))
```

### Pivot / Unpivot

```go
// Long to wide
wide := df.Pivot([]string{"name"}, "metric", "value")

// Wide to long
long := df.Unpivot([]string{"name"}, []string{"height", "weight"}, "metric", "value")

// Cross-tabulation
ct := df.CrossTab("department", "status")
```

### Data Quality

```go
// Descriptive statistics
stats, _ := df.Describe()

// Null analysis
nulls := df.NullCount()
complete := df.IsComplete()

// Validation rules
result := df.Validate(
    gf.NotNull("id"),
    gf.Positive("revenue"),
    gf.InRange("score", 0, 100),
    gf.UniqueValues("email"),
)
if !result.Valid {
    for _, v := range result.Violations {
        log.Println(v)
    }
}

// Outlier detection
outliers, _ := df.DetectOutliersIQR("salary", 1.5)
fmt.Printf("Found %d outliers\n", outliers.Count)
```

### Streaming Large Files

```go
it, _ := gf.ReadCSVChunked("large.csv", 100000)

it.ForEachChunk(func(chunk *gf.DataFrame) error {
    // Process each chunk independently
    filtered := chunk.Filter(chunk.Col("valid").Eq(gf.Lit("true")))
    gf.WriteJSON(filtered, fmt.Sprintf("output_%d.json", i))
    return nil
})
```

### JSON I/O

```go
df, _ := gf.ReadJSON("data.json")       // Array of objects
df, _ := gf.ReadNDJSON("data.ndjson")   // One object per line

gf.WriteJSON(df, "output.json")
gf.WriteNDJSON(df, "output.ndjson")
```

### Date Parsing

```go
// Auto-detect format
result := df.ParseDateColumn("date_str", "date")

// Specific format
result := df.ParseDateWithFormat("ts", "timestamp", "2006-01-02 15:04:05")

// Date range generation
dates, _ := gf.DateRange("date",
    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
    24*time.Hour,
)

// Business days
days := gf.BusinessDaysBetween(start, end)
nextBizDay := gf.AddBusinessDays(start, 5)
```

### Query Optimization

```go
result := gf.NewQueryPlan(df).
    WithColumn("total", df.Col("price").Mul(df.Col("qty"))).
    Filter(df.Col("price").Gt(gf.Lit(10.0))).  // Pushed before WithColumn
    Select("product", "total").
    Execute()
```

## Best Practices

1. **Always `defer df.Release()`** — prevents memory leaks
2. **Use type-safe literals** — `gf.Lit(int64(18))` not `gf.Lit(18)`
3. **Set memory limits in production** — `core.NewLimitedAllocator(pool, limit)`
4. **Use VectorUDF over ScalarUDF** — 10-100x faster for large datasets
5. **Filter early** — reduce data size before joins and aggregations
6. **Use streaming for large files** — `ReadCSVChunked` prevents OOM
