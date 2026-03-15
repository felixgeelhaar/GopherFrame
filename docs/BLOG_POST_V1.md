# GopherFrame v1.0: Production-Grade DataFrames for Go

*Making data processing in Go fast, safe, and production-ready.*

## The Problem

Go excels at building backend services, but when those services need non-trivial data processing — ETL pipelines, ML inference preprocessing, real-time analytics — teams often reach for Python. This creates operational complexity: two languages, two deployments, and a serialization boundary between them.

Existing Go DataFrame libraries like Gota fill the gap for simple cases, but they weren't designed for production workloads with millions of rows and strict memory constraints.

## The Solution

GopherFrame is an Apache Arrow-backed DataFrame library for Go that's designed production-first. Instead of optimizing for notebook exploration, every API decision prioritizes performance, reliability, and operational simplicity.

### Zero-Copy Performance

Built on Apache Arrow's columnar memory format, GopherFrame achieves:

- **67.8x faster column selection** than Gota (772ns vs 52.4ms)
- **428x faster iteration** (389ns vs 166µs)
- **2-200x less memory** across all operations

This isn't just benchmarketing — it's the natural result of Arrow's O(1) column selection (pointer update vs data copy) and cache-friendly columnar layout.

### Production Features

- **Memory limits**: `LimitedAllocator` prevents OOM with configurable caps and pressure monitoring
- **Streaming**: `ReadCSVChunked` processes files larger than RAM
- **Type safety**: Strong typing via Go generics and Arrow's type system
- **Zero security vulnerabilities**: gosec validated, path traversal protection

## What Can You Build?

### ETL Pipelines

```go
orders, _ := gf.ReadParquet("orders.parquet")
defer orders.Release()

result := orders.
    Filter(orders.Col("amount").Gt(gf.Lit(0.0))).
    WithColumn("tax", orders.Col("amount").Mul(gf.Lit(0.08)))

summary := result.GroupBy("region").Agg(
    gf.Sum("amount").As("total"),
    gf.Percentile("amount", 0.95).As("p95"),
)
gf.WriteParquet(summary, "regional.parquet")
```

### ML Inference Preprocessing

Use identical transformations in Go inference services as in Python training:

```go
features := raw.
    Filter(raw.Col("score").Gt(gf.Lit(0.0))).
    WithColumn("normalized", gf.ScalarUDF(
        []string{"score", "max_score"},
        arrow.PrimitiveTypes.Float64,
        func(row map[string]interface{}) (interface{}, error) {
            return row["score"].(float64) / row["max_score"].(float64), nil
        },
    ))
```

### Real-Time Analytics

```go
df.Window().
    PartitionBy("endpoint").
    OrderBy("timestamp").
    Over(
        gf.RollingMean("latency_ms", 100).As("p_avg"),
        gf.Rank().As("latency_rank"),
    )
```

## Complete Feature Set

- **Core**: Select, Filter, Sort, GroupBy, WithColumn
- **Joins**: Inner, Left, Right, Full Outer, Cross + multi-column keys + merge/broadcast strategies
- **Window**: RowNumber, Rank, DenseRank, Lag, Lead, Rolling, Cumulative
- **Temporal**: Year/Month/Day extraction, truncation, arithmetic, date parsing, business days
- **String**: Upper, Lower, Trim, Pad, Replace, SplitPart, Match (regex)
- **Statistics**: Mean, Median, Mode, Percentile, Variance, StdDev, Correlation
- **I/O**: Parquet, CSV, Arrow IPC, JSON, NDJSON, SQL, partitioned datasets
- **UDFs**: Scalar and vectorized user-defined functions
- **Pivot/Unpivot**: Wide-to-long, long-to-wide, cross-tabulation
- **Quality**: Validation rules, outlier detection, descriptive statistics
- **Optimization**: Filter pushdown, constant folding, query plans

## Getting Started

```bash
go get github.com/felixgeelhaar/GopherFrame
```

- [GitHub Repository](https://github.com/felixgeelhaar/GopherFrame)
- [API Reference](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame)
- [User Guide](docs/USER_GUIDE.md)
- [Migration from Pandas](docs/MIGRATION_FROM_PANDAS.md)
- [Performance Guide](docs/PERFORMANCE_GUIDE.md)

## What's Next

- **v1.1**: Avro support, partition pruning optimization, parallel reads
- **v1.2**: Query plan optimization, distributed execution

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) to get started.

---

*GopherFrame — because production data pipelines shouldn't require a second language.*
