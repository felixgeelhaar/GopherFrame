# Migrating from Gota to GopherFrame

## Overview

This guide helps Go developers migrate from Gota to GopherFrame for high-performance data processing. GopherFrame provides **2-428x better performance** with Apache Arrow integration while maintaining familiar DataFrame operations.

## Why Migrate?

### Performance Advantages (Validated Benchmarks)

| Operation | Gota | GopherFrame | **Speedup** |
|-----------|------|-------------|-------------|
| DataFrame Creation (10K) | 747µs | 249µs | **3.0x faster** |
| Filter (10K) | 748µs | 404µs | **1.9x faster** |
| **Select (10K)** | 52.4ms | **772ns** | **🚀 67.8x faster** |
| Column Access (1K) | 3.4µs | 120ns | **🚀 28x faster** |
| Iteration (1K) | 166µs | 389ns | **🚀 428x faster** |

**Memory Efficiency**: 2-200x less memory usage across all operations

📊 **Full Benchmark Report**: [GOTA_COMPARISON_BENCHMARKS.md](GOTA_COMPARISON_BENCHMARKS.md)

### Additional Benefits

- **Apache Arrow Integration**: Zero-copy data exchange with Arrow ecosystem
- **Production Features**: Memory limits, OOM handling, pressure monitoring
- **Type Safety**: Compile-time type checking via Go generics
- **Active Development**: Rapid feature additions and optimizations
- **Better Documentation**: Comprehensive guides and examples

## Key Differences

| Aspect | Gota | GopherFrame |
|--------|------|-------------|
| **Backend** | Custom in-memory | Apache Arrow |
| **Memory Layout** | Row-based | Columnar (cache-optimized) |
| **Column Selection** | O(n) - copies data | **O(1) - zero-copy** |
| **Null Handling** | Go nil/zero values | Arrow null bitmaps |
| **Memory Management** | GC only | Reference counting + GC |
| **I/O Formats** | CSV | CSV, Parquet, Arrow IPC |
| **Performance** | Good | **2-428x faster** |

## Installation

### Gota
```go
go get github.com/go-gota/gota
```

### GopherFrame
```go
go get github.com/felixgeelhaar/GopherFrame
```

## Common Operations Comparison

### Creating DataFrames

**Gota**:
```go
import "github.com/go-gota/gota/dataframe"
import "github.com/go-gota/gota/series"

df := dataframe.New(
    series.New([]int{1, 2, 3}, series.Int, "id"),
    series.New([]string{"A", "B", "C"}, series.String, "name"),
    series.New([]float64{1.1, 2.2, 3.3}, series.Float, "value"),
)
```

**GopherFrame**:
```go
import (
    gf "github.com/felixgeelhaar/GopherFrame"
    "github.com/apache/arrow-go/v18/arrow"
    "github.com/apache/arrow-go/v18/arrow/array"
    "github.com/apache/arrow-go/v18/arrow/memory"
)

pool := memory.NewGoAllocator()

schema := arrow.NewSchema(
    []arrow.Field{
        {Name: "id", Type: arrow.PrimitiveTypes.Int64},
        {Name: "name", Type: arrow.BinaryTypes.String},
        {Name: "value", Type: arrow.PrimitiveTypes.Float64},
    },
    nil,
)

// Build arrays
idBuilder := array.NewInt64Builder(pool)
defer idBuilder.Release()
idBuilder.AppendValues([]int64{1, 2, 3}, nil)

nameBuilder := array.NewStringBuilder(pool)
defer nameBuilder.Release()
nameBuilder.AppendValues([]string{"A", "B", "C"}, nil)

valueBuilder := array.NewFloat64Builder(pool)
defer valueBuilder.Release()
valueBuilder.AppendValues([]float64{1.1, 2.2, 3.3}, nil)

idArray := idBuilder.NewArray()
defer idArray.Release()
nameArray := nameBuilder.NewArray()
defer nameArray.Release()
valueArray := valueBuilder.NewArray()
defer valueArray.Release()

record := array.NewRecord(
    schema,
    []arrow.Array{idArray, nameArray, valueArray},
    3,
)

df := gf.NewDataFrame(record)
defer df.Release()
```

**Performance**: DataFrame creation is **3x faster** in GopherFrame

### Loading from CSV

**Gota**:
```go
file, _ := os.Open("data.csv")
defer file.Close()

df := dataframe.ReadCSV(file)
```

**GopherFrame**:
```go
df, err := gf.ReadCSV("data.csv")
if err != nil {
    log.Fatal(err)
}
defer df.Release()
```

**Note**: GopherFrame requires explicit error handling (Go best practice)

### Selecting Columns

**Gota**:
```go
// Select columns
subset := df.Select([]string{"name", "value"})

// Access single column
col := df.Col("name")
```

**GopherFrame**:
```go
// Select columns
subset := df.Select("name", "value")
defer subset.Release()

if subset.Err() != nil {
    log.Fatal(subset.Err())
}

// Access single column
col, err := df.Column("name")
if err != nil {
    log.Fatal(err)
}
defer col.Release()
```

**Performance**: Column selection is **67.8x faster** in GopherFrame (772ns vs 52.4ms)

### Filtering Rows

**Gota**:
```go
import "github.com/go-gota/gota/series"

// Filter with comparator
filtered := df.Filter(
    dataframe.F{
        Colname:    "value",
        Comparator: series.Greater,
        Comparando: 2.0,
    },
)

// Multiple conditions
filtered := df.Filter(
    dataframe.F{Colname: "value", Comparator: series.Greater, Comparando: 2.0},
    dataframe.F{Colname: "name", Comparator: series.Eq, Comparando: "A"},
)
```

**GopherFrame**:
```go
// Filter with expression
filtered := df.Filter(
    df.Col("value").Gt(gf.Lit(2.0)),
)
defer filtered.Release()

if filtered.Err() != nil {
    log.Fatal(filtered.Err())
}

// Multiple conditions (sequential filtering)
filtered1 := df.Filter(df.Col("value").Gt(gf.Lit(2.0)))
defer filtered1.Release()

filtered2 := filtered1.Filter(filtered1.Col("name").Eq(gf.Lit("A")))
defer filtered2.Release()
```

**Performance**: Filtering is **1.9x faster** in GopherFrame at 10K rows

### Sorting

**Gota**:
```go
// Sort ascending
sorted := df.Arrange(
    dataframe.Sort("value"),
)

// Sort descending
sorted := df.Arrange(
    dataframe.RevSort("value"),
)

// Multi-column sort
sorted := df.Arrange(
    dataframe.Sort("name"),
    dataframe.RevSort("value"),
)
```

**GopherFrame**:
```go
// Sort ascending
sorted := df.Sort("value", true)
defer sorted.Release()

// Sort descending
sorted := df.Sort("value", false)
defer sorted.Release()

// Multi-column sort
sorted := df.SortMultiple(
    gf.SortKey{Column: "name", Ascending: true},
    gf.SortKey{Column: "value", Ascending: false},
)
defer sorted.Release()
```

### GroupBy and Aggregation

**Gota**:
```go
// Group by and aggregate
grouped := df.GroupBy("category").Aggregation(
    []dataframe.AggregationType{
        dataframe.Aggregation_SUM,
        dataframe.Aggregation_MEAN,
    },
    []string{"value"},
)
```

**GopherFrame**:
```go
// Group by and aggregate
grouped := df.GroupBy("category").Agg(
    gf.Sum("value").As("total_value"),
    gf.Mean("value").As("avg_value"),
    gf.Count("value").As("count"),
)
defer grouped.Release()

if grouped.Err() != nil {
    log.Fatal(grouped.Err())
}
```

### Mutating DataFrames

**Gota**:
```go
// Mutate creates new column
mutated := df.Mutate(
    series.New([]float64{...}, series.Float, "new_column"),
)

// Set replaces existing column
modified := df.Set(
    []int{0, 1, 2},  // row indices
    series.New([]float64{10, 20, 30}, series.Float, "value"),
)
```

**GopherFrame**:
```go
// Add/replace column with computed values
withNew := df.WithColumn("new_column",
    df.Col("value").Mul(gf.Lit(2.0)),
)
defer withNew.Release()

// GopherFrame is immutable - creates new DataFrame
// No in-place modification like Gota's Set()
```

### Joins

**Gota**: Not natively supported
```go
// Must implement manually or use third-party extensions
```

**GopherFrame**: Full join support
```go
// Inner join
merged := df1.InnerJoin(df2, "key_col", "key_col")
defer merged.Release()

// Left join
merged := df1.LeftJoin(df2, "left_key", "right_key")
defer merged.Release()

if merged.Err() != nil {
    log.Fatal(merged.Err())
}
```

**Performance**: GopherFrame uses hash-based joins (O(n+m) complexity)

## Migration Patterns

### Pattern 1: ETL Pipeline

**Gota**:
```go
file, _ := os.Open("data.csv")
df := dataframe.ReadCSV(file)

// Filter
filtered := df.Filter(
    dataframe.F{Colname: "amount", Comparator: series.Greater, Comparando: 0},
)

// Aggregate
grouped := filtered.GroupBy("region").Aggregation(
    []dataframe.AggregationType{dataframe.Aggregation_SUM},
    []string{"amount"},
)

// Save
csvfile, _ := os.Create("output.csv")
grouped.WriteCSV(csvfile)
```

**GopherFrame**:
```go
df, err := gf.ReadCSV("data.csv")
if err != nil {
    log.Fatal(err)
}
defer df.Release()

// Filter
filtered := df.Filter(df.Col("amount").Gt(gf.Lit(0.0)))
defer filtered.Release()
if filtered.Err() != nil {
    log.Fatal(filtered.Err())
}

// Aggregate
grouped := filtered.GroupBy("region").Agg(
    gf.Sum("amount").As("total_amount"),
)
defer grouped.Release()
if grouped.Err() != nil {
    log.Fatal(grouped.Err())
}

// Save (CSV or Parquet)
if err := gf.WriteCSV(grouped, "output.csv"); err != nil {
    log.Fatal(err)
}

// Or use Parquet for better performance
if err := gf.WriteParquet(grouped, "output.parquet"); err != nil {
    log.Fatal(err)
}
```

### Pattern 2: Data Transformation

**Gota**:
```go
// Load
df := dataframe.ReadCSV(file)

// Transform
df = df.Mutate(
    series.New(computeValues(), series.Float, "computed"),
)

df = df.Filter(
    dataframe.F{Colname: "computed", Comparator: series.Greater, Comparando: threshold},
)

df = df.Select([]string{"id", "name", "computed"})
```

**GopherFrame**:
```go
// Load
df, _ := gf.ReadCSV("data.csv")
defer df.Release()

// Transform - add computed column
withComputed := df.WithColumn("computed",
    df.Col("value").Mul(gf.Lit(2.0)),
)
defer withComputed.Release()

// Filter
filtered := withComputed.Filter(
    withComputed.Col("computed").Gt(gf.Lit(threshold)),
)
defer filtered.Release()

// Select columns
selected := filtered.Select("id", "name", "computed")
defer selected.Release()
```

## Key Migration Considerations

### 1. Memory Management

**Gota**: Automatic garbage collection only
```go
df := dataframe.ReadCSV(file)
// Memory freed when df goes out of scope
```

**GopherFrame**: Explicit resource management (Arrow requirement)
```go
df, err := gf.ReadCSV("data.csv")
defer df.Release() // REQUIRED: Must explicitly release

// Every DataFrame operation creates a new DataFrame
// that must also be released
filtered := df.Filter(condition)
defer filtered.Release()
```

**Best Practice**: Always use `defer df.Release()` immediately after creation

### 2. Error Handling

**Gota**: Panics on errors
```go
df := df.Filter(condition) // Panics if error
```

**GopherFrame**: Returns errors or accumulates them
```go
// Option 1: Check errors immediately
df, err := gf.ReadCSV("data.csv")
if err != nil {
    return err
}

// Option 2: Accumulate errors in chain
filtered := df.Filter(condition)
if filtered.Err() != nil {
    return filtered.Err()
}
```

### 3. Type System

**Gota**: Dynamic typing with series.Type enum
```go
s := series.New([]int{1, 2, 3}, series.Int, "col")
// Can change types at runtime
```

**GopherFrame**: Static typing via Arrow schema
```go
// Type is defined in schema and enforced at compile time
{Name: "col", Type: arrow.PrimitiveTypes.Int64}

// Type safety: wrong type literal causes compile error
df.Col("col").Gt(gf.Lit(int64(5))) // Correct
df.Col("col").Gt(gf.Lit(5))        // Wrong: compile error
```

### 4. Immutability

**Gota**: Both mutable and immutable operations
```go
df = df.Mutate(...)  // Returns new DataFrame
df.Set(...)          // Modifies in-place
```

**GopherFrame**: Fully immutable
```go
df2 := df.WithColumn(...) // Always returns new DataFrame
// Original df is unchanged
```

## Performance Migration Strategy

### Step 1: Identify Hot Paths

Profile your Gota code to find performance bottlenecks:
```go
import "runtime/pprof"

f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// Your Gota code here
```

### Step 2: Migrate High-Impact Operations First

Focus on operations where GopherFrame shows biggest advantages:

1. **Column Selection** (67.8x faster): Migrate `df.Select()` calls
2. **Iteration** (428x faster): Migrate row/column iteration loops
3. **Column Access** (28x faster): Migrate frequent column lookups
4. **Creation** (3x faster): Migrate DataFrame construction
5. **Filtering** (1.9x faster): Migrate filter operations

### Step 3: Measure Performance Gains

```go
import "time"

start := time.Now()
// GopherFrame code
duration := time.Since(start)
fmt.Printf("Operation took: %v\n", duration)
```

### Step 4: Optimize Memory Usage

Use LimitedAllocator for production deployments:
```go
import "github.com/felixgeelhaar/GopherFrame/pkg/core"

pool := memory.NewGoAllocator()
limited := core.NewLimitedAllocator(pool, 1024*1024*1024) // 1GB limit

df := gf.NewDataFrameWithAllocator(record, limited)
defer df.Release()

// Monitor memory pressure
if limited.MemoryPressureLevel() == "high" {
    // Reduce batch size or trigger cleanup
}
```

## Complete Migration Example

### Original Gota Code

```go
package main

import (
    "os"
    "github.com/go-gota/gota/dataframe"
    "github.com/go-gota/gota/series"
)

func main() {
    // Load
    file, _ := os.Open("sales.csv")
    df := dataframe.ReadCSV(file)

    // Clean
    df = df.Filter(
        dataframe.F{Colname: "amount", Comparator: series.Greater, Comparando: 0},
    )

    // Transform
    df = df.Mutate(
        series.New(computeTax(df), series.Float, "tax"),
    )

    // Aggregate
    summary := df.GroupBy("region").Aggregation(
        []dataframe.AggregationType{
            dataframe.Aggregation_SUM,
            dataframe.Aggregation_MEAN,
        },
        []string{"amount"},
    )

    // Save
    outfile, _ := os.Create("summary.csv")
    summary.WriteCSV(outfile)
}
```

### Migrated GopherFrame Code

```go
package main

import (
    "log"
    gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
    // Load
    df, err := gf.ReadCSV("sales.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer df.Release()

    // Clean
    cleaned := df.Filter(df.Col("amount").Gt(gf.Lit(0.0)))
    defer cleaned.Release()
    if cleaned.Err() != nil {
        log.Fatal(cleaned.Err())
    }

    // Transform - compute tax
    withTax := cleaned.WithColumn("tax",
        cleaned.Col("amount").Mul(gf.Lit(0.08)),
    )
    defer withTax.Release()

    // Aggregate
    summary := withTax.GroupBy("region").Agg(
        gf.Sum("amount").As("total_amount"),
        gf.Mean("amount").As("avg_amount"),
    )
    defer summary.Release()
    if summary.Err() != nil {
        log.Fatal(summary.Err())
    }

    // Save (Parquet recommended for performance)
    if err := gf.WriteParquet(summary, "summary.parquet"); err != nil {
        log.Fatal(err)
    }

    // Or CSV for compatibility
    if err := gf.WriteCSV(summary, "summary.csv"); err != nil {
        log.Fatal(err)
    }
}
```

**Performance Improvement**: Expect 2-10x faster overall execution, with specific operations showing 67-428x speedup.

## Common Migration Pitfalls

### 1. Forgetting to Release Resources

❌ **Wrong**:
```go
df, _ := gf.ReadCSV("data.csv")
// Memory leak - DataFrame never released
```

✅ **Correct**:
```go
df, _ := gf.ReadCSV("data.csv")
defer df.Release()
```

### 2. Not Checking Errors

❌ **Wrong**:
```go
filtered := df.Filter(condition)
// Silent error - operation may have failed
```

✅ **Correct**:
```go
filtered := df.Filter(condition)
if filtered.Err() != nil {
    return filtered.Err()
}
```

### 3. Type Mismatches

❌ **Wrong**:
```go
df.Col("age").Gt(gf.Lit(18)) // age is Int64, 18 is int
```

✅ **Correct**:
```go
df.Col("age").Gt(gf.Lit(int64(18)))
```

### 4. Modifying Original DataFrame

❌ **Wrong** (Gota thinking):
```go
df.Set(...) // Doesn't exist in GopherFrame
```

✅ **Correct**:
```go
newDF := df.WithColumn("col", expr)
defer newDF.Release()
```

## Resources

- [Performance Comparison](GOTA_COMPARISON_BENCHMARKS.md)
- [GopherFrame Documentation](../README.md)
- [API Reference](API.md)
- [Example Programs](../cmd/examples/)

## Getting Help

- GitHub Issues: https://github.com/felixgeelhaar/GopherFrame/issues
- Discussions: https://github.com/felixgeelhaar/GopherFrame/discussions

## Conclusion

Migrating from Gota to GopherFrame provides:
- **2-428x better performance** (validated benchmarks)
- **2-200x less memory usage**
- **Apache Arrow integration** for ecosystem compatibility
- **Production-grade features** (memory limits, monitoring)
- **Active development** and growing community

The migration requires explicit resource management and error handling, but these Go best practices lead to more robust production code. The performance gains make GopherFrame the clear choice for production data processing in Go.
