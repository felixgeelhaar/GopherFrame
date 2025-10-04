# Migrating from Polars to GopherFrame

## Overview

This guide helps Python developers using Polars migrate to GopherFrame for Go-based data processing. Both libraries share Apache Arrow as their foundation, making conceptual translation straightforward while adapting to Go's type system and idioms.

## Why Both Libraries Excel

### Shared Foundation: Apache Arrow

Both Polars and GopherFrame leverage Apache Arrow's columnar memory format:
- **Zero-copy operations** for column selection
- **Cache-friendly memory layout** for fast iteration
- **Interoperability** with Arrow ecosystem (Parquet, Flight RPC)
- **SIMD-optimized computations** via Arrow compute kernels

### Key Similarities

| Feature | Polars | GopherFrame |
|---------|--------|-------------|
| **Memory Layout** | Columnar (Arrow) | Columnar (Arrow) |
| **Performance** | Rust-optimized | Go-optimized |
| **Lazy Evaluation** | Supported | Planned (v0.2) |
| **Type System** | Static (Rust) | Static (Go) |
| **Null Handling** | Arrow null bitmaps | Arrow null bitmaps |
| **I/O** | Parquet, CSV, Arrow IPC | Parquet, CSV, Arrow IPC |

### Key Differences

| Aspect | Polars | GopherFrame |
|--------|--------|-------------|
| **Language** | Python (Rust core) | Go (pure Go + Arrow Go) |
| **Lazy Evaluation** | ✅ LazyFrame | ⏳ Planned v0.2 |
| **Deployment** | Python runtime | Compiled binary |
| **Concurrency** | Python threading limits | Native Go concurrency |
| **Integration** | Python ecosystem | Go services |

## Installation

### Polars
```bash
pip install polars
```

```python
import polars as pl
```

### GopherFrame
```bash
go get github.com/felixgeelhaar/GopherFrame
```

```go
import gf "github.com/felixgeelhaar/GopherFrame"
```

## Common Operations Comparison

### Creating DataFrames

**Polars**:
```python
import polars as pl

# From dictionary
df = pl.DataFrame({
    'id': [1, 2, 3, 4, 5],
    'name': ['Alice', 'Bob', 'Carol', 'David', 'Eve'],
    'age': [25, 30, 35, 40, 45],
    'salary': [50000.0, 60000.0, 70000.0, 80000.0, 90000.0]
})

# From CSV
df = pl.read_csv('data.csv')

# From Parquet
df = pl.read_parquet('data.parquet')
```

**GopherFrame**:
```go
import (
    gf "github.com/felixgeelhaar/GopherFrame"
    "github.com/apache/arrow-go/v18/arrow"
    "github.com/apache/arrow-go/v18/arrow/array"
    "github.com/apache/arrow-go/v18/arrow/memory"
)

// From arrays (manual construction required)
pool := memory.NewGoAllocator()
schema := arrow.NewSchema(
    []arrow.Field{
        {Name: "id", Type: arrow.PrimitiveTypes.Int64},
        {Name: "name", Type: arrow.BinaryTypes.String},
        {Name: "age", Type: arrow.PrimitiveTypes.Int64},
        {Name: "salary", Type: arrow.PrimitiveTypes.Float64},
    },
    nil,
)

// Build arrays (more verbose than Polars)
idBuilder := array.NewInt64Builder(pool)
idBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
// ... build other columns ...

record := array.NewRecord(schema, columns, 5)
df := gf.NewDataFrame(record)
defer df.Release()

// From CSV
df, err := gf.ReadCSV("data.csv")
defer df.Release()

// From Parquet
df, err := gf.ReadParquet("data.parquet")
defer df.Release()
```

### Selecting Columns

**Polars**:
```python
# Select columns
df.select(['name', 'age'])

# Select with expressions
df.select([
    pl.col('name'),
    pl.col('age') + 1
])
```

**GopherFrame**:
```go
// Select columns
selected := df.Select("name", "age")
defer selected.Release()

// Select with computed columns requires WithColumn
withAge := df.WithColumn("age_plus_one",
    df.Col("age").Add(gf.Lit(int64(1))),
)
defer withAge.Release()
```

### Filtering

**Polars**:
```python
# Simple filter
df.filter(pl.col('age') > 30)

# Multiple conditions
df.filter(
    (pl.col('age') > 30) & (pl.col('salary') > 60000)
)

# String operations
df.filter(pl.col('name').str.contains('Al'))
```

**GopherFrame**:
```go
// Simple filter
filtered := df.Filter(
    df.Col("age").Gt(gf.Lit(int64(30))),
)
defer filtered.Release()

// Multiple conditions (sequential filtering)
age30Plus := df.Filter(df.Col("age").Gt(gf.Lit(int64(30))))
defer age30Plus.Release()
highEarners := age30Plus.Filter(
    age30Plus.Col("salary").Gt(gf.Lit(60000.0)),
)
defer highEarners.Release()

// String operations
withAlice := df.Filter(
    df.Col("name").Contains(gf.Lit("Al")),
)
defer withAlice.Release()
```

### Adding/Modifying Columns

**Polars**:
```python
# Add column
df.with_columns([
    (pl.col('salary') * 0.1).alias('bonus')
])

# Multiple columns
df.with_columns([
    (pl.col('salary') * 0.1).alias('bonus'),
    (pl.col('age') + 1).alias('next_age')
])
```

**GopherFrame**:
```go
// Add column
withBonus := df.WithColumn("bonus",
    df.Col("salary").Mul(gf.Lit(0.1)),
)
defer withBonus.Release()

// Multiple columns (sequential)
withBonus := df.WithColumn("bonus",
    df.Col("salary").Mul(gf.Lit(0.1)),
)
defer withBonus.Release()

withNextAge := withBonus.WithColumn("next_age",
    withBonus.Col("age").Add(gf.Lit(int64(1))),
)
defer withNextAge.Release()
```

### GroupBy and Aggregation

**Polars**:
```python
# Group by with aggregations
df.group_by('department').agg([
    pl.col('salary').sum().alias('total_salary'),
    pl.col('salary').mean().alias('avg_salary'),
    pl.col('id').count().alias('employee_count')
])

# Multiple group keys
df.group_by(['department', 'level']).agg([
    pl.col('salary').mean()
])
```

**GopherFrame**:
```go
// Group by with aggregations
grouped := df.GroupBy("department").Agg(
    gf.Sum("salary").As("total_salary"),
    gf.Mean("salary").As("avg_salary"),
    gf.Count("id").As("employee_count"),
)
defer grouped.Release()

// Multiple group keys (not yet implemented in v0.1)
// Planned for v0.2
```

### Sorting

**Polars**:
```python
# Sort ascending
df.sort('age')

# Sort descending
df.sort('salary', descending=True)

# Multi-column sort
df.sort(['age', 'salary'], descending=[False, True])
```

**GopherFrame**:
```go
// Sort ascending
sorted := df.Sort("age", true)
defer sorted.Release()

// Sort descending
sorted := df.Sort("salary", false)
defer sorted.Release()

// Multi-column sort
sorted := df.SortMultiple(
    gf.SortKey{Column: "age", Ascending: true},
    gf.SortKey{Column: "salary", Ascending: false},
)
defer sorted.Release()
```

### Joins

**Polars**:
```python
# Inner join
df1.join(df2, on='id', how='inner')

# Left join with different key names
df1.join(df2, left_on='emp_id', right_on='id', how='left')

# Multiple keys
df1.join(df2, on=['dept', 'year'], how='inner')
```

**GopherFrame**:
```go
// Inner join
merged := df1.InnerJoin(df2, "id", "id")
defer merged.Release()

// Left join with different key names
merged := df1.LeftJoin(df2, "emp_id", "id")
defer merged.Release()

// Multiple keys (not yet implemented)
// Planned for v0.2
```

### Lazy Evaluation

**Polars**:
```python
# Create lazy frame
lf = pl.scan_csv('data.csv')

# Build query plan
result = (lf
    .filter(pl.col('age') > 30)
    .group_by('department')
    .agg(pl.col('salary').mean())
    .collect()  # Execute
)
```

**GopherFrame**:
```go
// Lazy evaluation not yet implemented
// v0.1 uses eager evaluation

// Planned for v0.2:
// query := gf.Scan("data.csv").
//     Filter(...).
//     GroupBy(...).
//     Collect()
```

## Migration Patterns

### Pattern 1: ETL Pipeline

**Polars**:
```python
import polars as pl

# Load
df = pl.read_csv('orders.csv')

# Transform
df = (df
    .filter(pl.col('amount') > 0)
    .with_columns([
        (pl.col('amount') * 0.08).alias('tax'),
    ])
    .with_columns([
        (pl.col('amount') + pl.col('tax')).alias('total')
    ])
)

# Join
customers = pl.read_csv('customers.csv')
merged = df.join(customers, on='customer_id', how='inner')

# Aggregate
summary = merged.group_by('region').agg([
    pl.col('total').sum().alias('total_revenue'),
    pl.col('order_id').count().alias('order_count')
])

# Save
summary.write_parquet('summary.parquet')
```

**GopherFrame**:
```go
import gf "github.com/felixgeelhaar/GopherFrame"

// Load
orders, err := gf.ReadCSV("orders.csv")
if err != nil {
    log.Fatal(err)
}
defer orders.Release()

// Transform
validOrders := orders.Filter(
    orders.Col("amount").Gt(gf.Lit(0.0)),
)
defer validOrders.Release()

withTax := validOrders.WithColumn("tax",
    validOrders.Col("amount").Mul(gf.Lit(0.08)),
)
defer withTax.Release()

withTotal := withTax.WithColumn("total",
    withTax.Col("amount").Add(withTax.Col("tax")),
)
defer withTotal.Release()

// Join
customers, err := gf.ReadCSV("customers.csv")
if err != nil {
    log.Fatal(err)
}
defer customers.Release()

merged := withTotal.InnerJoin(customers, "customer_id", "customer_id")
defer merged.Release()

// Aggregate
summary := merged.GroupBy("region").Agg(
    gf.Sum("total").As("total_revenue"),
    gf.Count("order_id").As("order_count"),
)
defer summary.Release()

// Save
if err := gf.WriteParquet(summary, "summary.parquet"); err != nil {
    log.Fatal(err)
}
```

### Pattern 2: ML Preprocessing

**Polars**:
```python
import polars as pl

# Load
df = pl.read_parquet('raw_features.parquet')

# Feature engineering
df = df.with_columns([
    (pl.col('revenue') - pl.col('cost')).alias('profit'),
    (pl.col('clicks') / pl.col('impressions')).alias('ctr'),
])

# Filter training data
train_df = df.filter(pl.col('date') < '2024-01-01')

# Select features
features = train_df.select([
    'user_id', 'profit', 'ctr', 'target'
])

# Save
features.write_parquet('train_features.parquet')
```

**GopherFrame**:
```go
// Load
df, err := gf.ReadParquet("raw_features.parquet")
defer df.Release()

// Feature engineering
withProfit := df.WithColumn("profit",
    df.Col("revenue").Sub(df.Col("cost")),
)
defer withProfit.Release()

withCTR := withProfit.WithColumn("ctr",
    withProfit.Col("clicks").Div(withProfit.Col("impressions")),
)
defer withCTR.Release()

// Filter training data
trainDF := withCTR.Filter(
    withCTR.Col("date").Lt(gf.Lit("2024-01-01")),
)
defer trainDF.Release()

// Select features
features := trainDF.Select("user_id", "profit", "ctr", "target")
defer features.Release()

// Save
if err := gf.WriteParquet(features, "train_features.parquet"); err != nil {
    log.Fatal(err)
}
```

## Key Migration Considerations

### 1. Memory Management

**Polars**: Automatic (Python GC + Rust ownership)
```python
df = pl.read_csv('data.csv')
# Automatic cleanup
```

**GopherFrame**: Explicit (Arrow reference counting)
```go
df, _ := gf.ReadCSV("data.csv")
defer df.Release() // REQUIRED

// Every operation creates new DataFrame
filtered := df.Filter(condition)
defer filtered.Release() // Also required
```

### 2. Type System

**Polars**: Python types with Rust backend
```python
# Automatic type inference
df = pl.DataFrame({'age': [25, 30, 35]})
```

**GopherFrame**: Go types with Arrow schema
```go
// Explicit type specification
schema := arrow.NewSchema(
    []arrow.Field{
        {Name: "age", Type: arrow.PrimitiveTypes.Int64},
    },
    nil,
)
```

### 3. Lazy vs Eager Evaluation

**Polars**: Lazy by default (with `scan_*`)
```python
lf = pl.scan_csv('data.csv')  # Lazy
lf = lf.filter(...)           # Builds query plan
df = lf.collect()              # Executes
```

**GopherFrame**: Eager in v0.1
```go
df, _ := gf.ReadCSV("data.csv")  // Eager - loads immediately
filtered := df.Filter(...)        // Executes immediately
```

### 4. Expression System

**Polars**: Rich expression DSL
```python
pl.col('age').filter(pl.col('age') > 18).mean()
```

**GopherFrame**: Basic expressions (v0.1)
```go
// Simpler expression system
df.Col("age").Gt(gf.Lit(int64(18)))

// Complex expressions planned for v0.2
```

## Performance Comparison

### Polars Performance
- **Rust-optimized** core with Python interface
- **Lazy evaluation** enables query optimization
- **Parallel execution** via Rust parallelism
- **SIMD operations** via Arrow compute kernels

### GopherFrame Performance
- **Go-optimized** with zero-copy Arrow operations
- **Eager evaluation** in v0.1 (lazy planned)
- **Native concurrency** via Go goroutines
- **SIMD potential** via Arrow compute kernels

### When to Use Each

**Use Polars when**:
- Python-based ML/data science workflows
- Need lazy evaluation and query optimization
- Rich ecosystem integration (pandas, sklearn)
- Team expertise in Python

**Use GopherFrame when**:
- Production Go services requiring data processing
- Need compiled binary deployment
- Native Go concurrency integration
- ML inference preprocessing in Go
- Microservices data pipelines

## Production Features

### Memory Limits (GopherFrame)

Unlike Polars, GopherFrame provides explicit memory management:

```go
import "github.com/felixgeelhaar/GopherFrame/pkg/core"

// Set 1GB memory limit
pool := memory.NewGoAllocator()
limited := core.NewLimitedAllocator(pool, 1024*1024*1024)

df := gf.NewDataFrameWithAllocator(record, limited)
defer df.Release()

// Monitor memory pressure
if limited.MemoryPressureLevel() == "high" {
    // Reduce batch size or trigger cleanup
}
```

This is useful for:
- Cloud deployments with memory constraints
- Multi-tenant services
- Preventing OOM crashes

## Roadmap Alignment

GopherFrame is actively developing features to match Polars:

| Feature | Polars | GopherFrame |
|---------|--------|-------------|
| **Basic Operations** | ✅ | ✅ v0.1 |
| **Joins** | ✅ | ✅ v0.1 |
| **GroupBy** | ✅ | ✅ v0.1 |
| **Window Functions** | ✅ | ⏳ v0.2 |
| **Lazy Evaluation** | ✅ | ⏳ v0.2 |
| **Multiple Group Keys** | ✅ | ⏳ v0.2 |
| **SQL Interface** | ✅ | ⏳ v0.3 |
| **UDFs** | ✅ | ⏳ v0.3 |

## Resources

- [GopherFrame Documentation](../README.md)
- [Polars Documentation](https://pola-rs.github.io/polars/)
- [Apache Arrow Documentation](https://arrow.apache.org/)

## Getting Help

- GitHub Issues: https://github.com/felixgeelhaar/GopherFrame/issues
- Discussions: https://github.com/felixgeelhaar/GopherFrame/discussions

## Conclusion

GopherFrame and Polars share Arrow as their foundation, providing similar performance characteristics. Choose GopherFrame when:
- Building production Go services
- Need compiled binary deployment
- Integrating with Go microservices
- Require explicit memory management

The migration requires adapting to Go's type system and explicit resource management, but both libraries provide excellent performance through their shared Arrow foundation.
