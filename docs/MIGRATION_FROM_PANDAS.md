# Migrating from Pandas to GopherFrame

## Overview

This guide helps Python developers using Pandas migrate to GopherFrame for Go-based data processing. GopherFrame provides similar functionality with Go's type safety, performance advantages, and seamless integration with Go services.

## Key Differences

| Aspect | Pandas | GopherFrame |
|--------|--------|-------------|
| **Language** | Python | Go |
| **Type System** | Dynamic (runtime) | Static (compile-time) |
| **Memory Model** | Row-based (sometimes columnar) | Columnar (Apache Arrow) |
| **Performance** | Interpreted, some C extensions | Compiled, zero-copy operations |
| **Null Handling** | NaN, None, pd.NA | Arrow null bitmaps |
| **Memory Management** | Garbage collected | Reference counted + GC |
| **API Style** | Chained methods, in-place options | Chained methods, immutable |

## Installation and Setup

### Pandas
```python
pip install pandas
import pandas as pd
```

### GopherFrame
```go
go get github.com/felixgeelhaar/GopherFrame

import (
    gf "github.com/felixgeelhaar/GopherFrame"
    "github.com/apache/arrow-go/v18/arrow"
    "github.com/apache/arrow-go/v18/arrow/array"
    "github.com/apache/arrow-go/v18/arrow/memory"
)
```

## Common Operations Comparison

### Creating DataFrames

**Pandas**:
```python
# From dictionary
df = pd.DataFrame({
    'id': [1, 2, 3, 4, 5],
    'name': ['Alice', 'Bob', 'Carol', 'David', 'Eve'],
    'age': [25, 30, 35, 40, 45],
    'salary': [50000, 60000, 70000, 80000, 90000]
})

# From CSV
df = pd.read_csv('data.csv')

# From Parquet
df = pd.read_parquet('data.parquet')
```

**GopherFrame**:
```go
// From arrays (Arrow-native)
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

// Build arrays
idBuilder := array.NewInt64Builder(pool)
defer idBuilder.Release()
idBuilder.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
// ... build other columns ...

record := array.NewRecord(schema, columns, numRows)
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

**Pandas**:
```python
# Single column
ages = df['age']

# Multiple columns
subset = df[['name', 'age']]

# Column selection by condition
numeric_cols = df.select_dtypes(include=['number'])
```

**GopherFrame**:
```go
// Single column
ageCol, err := df.Column("age")
defer ageCol.Release()

// Multiple columns
subset := df.Select("name", "age")
defer subset.Release()

// Note: Type-based selection not yet implemented
// Workaround: manually specify column names
```

### Filtering Rows

**Pandas**:
```python
# Simple condition
adults = df[df['age'] >= 18]

# Multiple conditions
high_earners = df[(df['age'] >= 30) & (df['salary'] > 70000)]

# String matching
alice_records = df[df['name'] == 'Alice']
```

**GopherFrame**:
```go
// Simple condition
adults := df.Filter(
    df.Col("age").Gt(gf.Lit(int64(18))),
)
defer adults.Release()

// Multiple conditions (manual combination)
age30Plus := df.Filter(df.Col("age").Gt(gf.Lit(int64(30))))
defer age30Plus.Release()
highEarners := age30Plus.Filter(
    age30Plus.Col("salary").Gt(gf.Lit(70000.0)),
)
defer highEarners.Release()

// String matching
aliceRecords := df.Filter(
    df.Col("name").Eq(gf.Lit("Alice")),
)
defer aliceRecords.Release()
```

### Adding/Modifying Columns

**Pandas**:
```python
# Add new column
df['bonus'] = df['salary'] * 0.1

# Modify existing column
df['age'] = df['age'] + 1

# Computed column
df['annual_income'] = df['salary'] * 12
```

**GopherFrame**:
```go
// Add new column
withBonus := df.WithColumn("bonus",
    df.Col("salary").Mul(gf.Lit(0.1)),
)
defer withBonus.Release()

// Modify existing column (replace)
olderDF := df.WithColumn("age",
    df.Col("age").Add(gf.Lit(int64(1))),
)
defer olderDF.Release()

// Computed column
withAnnual := df.WithColumn("annual_income",
    df.Col("salary").Mul(gf.Lit(12.0)),
)
defer withAnnual.Release()
```

### Sorting

**Pandas**:
```python
# Sort by single column
sorted_df = df.sort_values('age')

# Sort descending
sorted_df = df.sort_values('salary', ascending=False)

# Sort by multiple columns
sorted_df = df.sort_values(['age', 'salary'], ascending=[True, False])
```

**GopherFrame**:
```go
// Sort by single column
sorted := df.Sort("age", true) // true = ascending
defer sorted.Release()

// Sort descending
sortedDesc := df.Sort("salary", false) // false = descending
defer sortedDesc.Release()

// Sort by multiple columns
sorted := df.SortMultiple(
    gf.SortKey{Column: "age", Ascending: true},
    gf.SortKey{Column: "salary", Ascending: false},
)
defer sorted.Release()
```

### GroupBy and Aggregation

**Pandas**:
```python
# Group by single column
grouped = df.groupby('department').agg({
    'salary': ['mean', 'sum', 'count'],
    'age': 'mean'
})

# Multiple grouping columns
grouped = df.groupby(['department', 'level']).agg({
    'salary': 'mean'
})
```

**GopherFrame**:
```go
// Group by single column
grouped := df.GroupBy("department").Agg(
    gf.Mean("salary").As("avg_salary"),
    gf.Sum("salary").As("total_salary"),
    gf.Count("salary").As("employee_count"),
    gf.Mean("age").As("avg_age"),
)
defer grouped.Release()

// Multiple grouping columns (not yet implemented)
// Workaround: create composite key column first
```

### Joins

**Pandas**:
```python
# Inner join
merged = df1.merge(df2, on='id', how='inner')

# Left join
merged = df1.merge(df2, left_on='emp_id', right_on='id', how='left')

# Multiple key columns
merged = df1.merge(df2, on=['dept', 'year'])
```

**GopherFrame**:
```go
// Inner join
merged := df1.InnerJoin(df2, "id", "id")
defer merged.Release()

// Left join
merged := df1.LeftJoin(df2, "emp_id", "id")
defer merged.Release()

// Multiple key columns (not yet implemented)
// Workaround: create composite key column
```

### I/O Operations

**Pandas**:
```python
# CSV
df = pd.read_csv('input.csv')
df.to_csv('output.csv', index=False)

# Parquet
df = pd.read_parquet('input.parquet')
df.to_parquet('output.parquet')

# Excel (not supported in GopherFrame)
df = pd.read_excel('input.xlsx')
df.to_excel('output.xlsx')
```

**GopherFrame**:
```go
// CSV
df, err := gf.ReadCSV("input.csv")
defer df.Release()
err = gf.WriteCSV(df, "output.csv")

// Parquet (native Arrow format - highly optimized)
df, err := gf.ReadParquet("input.parquet")
defer df.Release()
err = gf.WriteParquet(df, "output.parquet")

// Excel not supported - use CSV as interchange format
```

## Migration Patterns

### Pattern 1: ETL Pipeline

**Pandas**:
```python
# Load
df = pd.read_csv('orders.csv')

# Transform
df = df[df['amount'] > 0]  # Filter invalid
df['tax'] = df['amount'] * 0.08
df['total'] = df['amount'] + df['tax']

# Aggregate
summary = df.groupby('region').agg({
    'total': 'sum',
    'order_id': 'count'
})

# Save
summary.to_parquet('summary.parquet')
```

**GopherFrame**:
```go
// Load
df, err := gf.ReadCSV("orders.csv")
defer df.Release()

// Transform
validDF := df.Filter(df.Col("amount").Gt(gf.Lit(0.0)))
defer validDF.Release()

withTax := validDF.WithColumn("tax",
    validDF.Col("amount").Mul(gf.Lit(0.08)),
)
defer withTax.Release()

withTotal := withTax.WithColumn("total",
    withTax.Col("amount").Add(withTax.Col("tax")),
)
defer withTotal.Release()

// Aggregate
summary := withTotal.GroupBy("region").Agg(
    gf.Sum("total").As("total_revenue"),
    gf.Count("order_id").As("order_count"),
)
defer summary.Release()

// Save
err = gf.WriteParquet(summary, "summary.parquet")
```

### Pattern 2: Data Cleaning

**Pandas**:
```python
# Remove duplicates
df = df.drop_duplicates()

# Handle missing values
df = df.dropna()  # or df.fillna(0)

# Type conversion
df['age'] = df['age'].astype(int)
df['date'] = pd.to_datetime(df['date'])
```

**GopherFrame**:
```go
// Remove duplicates (not yet implemented)
// Workaround: GroupBy + First aggregation

// Handle missing values
// Filter out rows with nulls in specific column
nonNullDF := df.Filter(/* manual null check */)

// Type conversion
// Must be done at DataFrame creation time via schema
// GopherFrame is strongly typed - schema defines types upfront
```

### Pattern 3: Time Series Analysis

**Pandas**:
```python
# Set datetime index
df['date'] = pd.to_datetime(df['date'])
df = df.set_index('date')

# Resample
daily = df.resample('D').mean()

# Rolling window
df['ma_7'] = df['value'].rolling(window=7).mean()
```

**GopherFrame**:
```go
// Time series operations not yet implemented in v0.1
// Planned for v0.2 with window functions

// Current workaround: preprocess in Python or use custom logic
```

## Performance Considerations

### Pandas Performance
- Interpreted Python overhead
- Some operations use C extensions (NumPy, Cython)
- Copy-on-write for some operations
- Memory usage can be high for large datasets

### GopherFrame Performance
- Compiled Go code (no interpreter overhead)
- Zero-copy operations via Apache Arrow
- Columnar memory layout (cache-efficient)
- 2-428x faster than similar Go library (Gota)
- **67x faster** for column selection
- **428x faster** for iteration

### When to Use Each

**Use Pandas when**:
- Rapid prototyping and exploration
- Rich ecosystem integration (matplotlib, scikit-learn)
- Complex time series analysis
- Jupyter notebook workflows
- Team primarily Python-based

**Use GopherFrame when**:
- Production Go services requiring data processing
- High-performance ETL pipelines
- Real-time analytics in Go backends
- ML model inference in Go (preprocessing)
- Memory-constrained environments
- Need for compile-time type safety

## Common Gotchas

### 1. Memory Management

**Pandas**: Automatic garbage collection
```python
df = pd.read_csv('large.csv')
# Memory automatically released when df goes out of scope
```

**GopherFrame**: Explicit resource management
```go
df, err := gf.ReadCSV("large.csv")
defer df.Release() // MUST explicitly release!

// Or immediate cleanup
df.Release()
```

### 2. Null Handling

**Pandas**: Multiple null representations (None, NaN, pd.NA)
```python
df['age'].fillna(0)
df['name'].isna()
```

**GopherFrame**: Arrow null bitmaps
```go
// Check if value is null
col, _ := df.Column("age")
defer col.Release()
isNull := col.IsNull(rowIndex)

// Null filtering requires manual null checking
```

### 3. Type Conversion

**Pandas**: Runtime type conversion
```python
df['age'] = df['age'].astype(float)
```

**GopherFrame**: Types defined at schema creation
```go
// Must define correct type in schema from start
schema := arrow.NewSchema(
    []arrow.Field{
        {Name: "age", Type: arrow.PrimitiveTypes.Float64},
    },
    nil,
)
```

### 4. Chaining Operations

**Pandas**: Method chaining
```python
result = (df
    .query('age > 30')
    .groupby('department')
    .agg({'salary': 'mean'})
    .sort_values('salary', ascending=False)
)
```

**GopherFrame**: Chaining with explicit cleanup
```go
// Intermediate results need cleanup
filtered := df.Filter(df.Col("age").Gt(gf.Lit(int64(30))))
defer filtered.Release()

grouped := filtered.GroupBy("department").Agg(
    gf.Mean("salary").As("avg_salary"),
)
defer grouped.Release()

result := grouped.Sort("avg_salary", false)
defer result.Release()
```

## Best Practices

### 1. Defer Release Calls

```go
func processData() error {
    df, err := gf.ReadCSV("data.csv")
    if err != nil {
        return err
    }
    defer df.Release() // Ensures cleanup even if errors occur

    // ... processing ...
    return nil
}
```

### 2. Check Errors

```go
df := df.Filter(condition)
if df.Err() != nil {
    return fmt.Errorf("filter failed: %w", df.Err())
}
```

### 3. Use Type-Safe Literals

```go
// Correct: type matches column type
df.Col("age").Gt(gf.Lit(int64(18)))  // age is Int64

// Incorrect: type mismatch will cause errors
df.Col("age").Gt(gf.Lit(18))  // int, not int64
```

### 4. Pre-allocate Memory for Large Operations

```go
// Use LimitedAllocator for production
pool := memory.NewGoAllocator()
limited := core.NewLimitedAllocator(pool, 1024*1024*1024) // 1GB limit
df := gf.NewDataFrameWithAllocator(record, limited)
```

## Example: Complete Migration

### Original Pandas Code

```python
import pandas as pd

# Load data
orders = pd.read_csv('orders.csv')
customers = pd.read_csv('customers.csv')

# Clean data
orders = orders[orders['amount'] > 0]
orders = orders.dropna()

# Feature engineering
orders['tax'] = orders['amount'] * 0.08
orders['total'] = orders['amount'] + orders['tax']

# Join with customers
merged = orders.merge(customers, on='customer_id', how='inner')

# Aggregate
summary = merged.groupby('region').agg({
    'total': ['sum', 'mean'],
    'order_id': 'count'
})

# Save
summary.to_parquet('regional_summary.parquet')
```

### Migrated GopherFrame Code

```go
package main

import (
    "log"
    gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
    // Load data
    orders, err := gf.ReadCSV("orders.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer orders.Release()

    customers, err := gf.ReadCSV("customers.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer customers.Release()

    // Clean data - filter invalid amounts
    validOrders := orders.Filter(
        orders.Col("amount").Gt(gf.Lit(0.0)),
    )
    defer validOrders.Release()
    if validOrders.Err() != nil {
        log.Fatal(validOrders.Err())
    }

    // Feature engineering - add tax
    withTax := validOrders.WithColumn("tax",
        validOrders.Col("amount").Mul(gf.Lit(0.08)),
    )
    defer withTax.Release()

    // Add total
    withTotal := withTax.WithColumn("total",
        withTax.Col("amount").Add(withTax.Col("tax")),
    )
    defer withTotal.Release()

    // Join with customers
    merged := withTotal.InnerJoin(customers, "customer_id", "customer_id")
    defer merged.Release()
    if merged.Err() != nil {
        log.Fatal(merged.Err())
    }

    // Aggregate by region
    summary := merged.GroupBy("region").Agg(
        gf.Sum("total").As("total_revenue"),
        gf.Mean("total").As("avg_order_value"),
        gf.Count("order_id").As("order_count"),
    )
    defer summary.Release()
    if summary.Err() != nil {
        log.Fatal(summary.Err())
    }

    // Save
    if err := gf.WriteParquet(summary, "regional_summary.parquet"); err != nil {
        log.Fatal(err)
    }
}
```

## Resources

- [GopherFrame Documentation](../README.md)
- [API Reference](API.md)
- [Performance Benchmarks](../BENCHMARKS.md)
- [Example Programs](../cmd/examples/)

## Getting Help

- GitHub Issues: https://github.com/felixgeelhaar/GopherFrame/issues
- Discussions: https://github.com/felixgeelhaar/GopherFrame/discussions
