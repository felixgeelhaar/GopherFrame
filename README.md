# GopherFrame

**A production-ready DataFrame library for Go, powered by Apache Arrow.**

[![CI](https://github.com/felixgeelhaar/GopherFrame/actions/workflows/ci.yml/badge.svg)](https://github.com/felixgeelhaar/GopherFrame/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/felixgeelhaar/GopherFrame.svg)](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame)
[![Go Report Card](https://goreportcard.com/badge/github.com/felixgeelhaar/GopherFrame)](https://goreportcard.com/report/github.com/felixgeelhaar/GopherFrame)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)](https://go.dev/)
[![Tests](https://img.shields.io/badge/tests-438%20passing-brightgreen.svg)](.github/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-81.4%25-brightgreen.svg)](.coverctl.yaml)
[![Security](https://img.shields.io/badge/security-0%20critical-brightgreen.svg)](SECURITY.md)

GopherFrame is a high-performance DataFrame library for Go, built on Apache Arrow for zero-copy data operations and seamless interoperability with the modern data ecosystem. Designed from the ground up for production use, it provides **2-428x better performance** than existing Go alternatives while maintaining type safety and idiomatic Go design.

## 🚀 Quick Start

```bash
go get github.com/felixgeelhaar/GopherFrame
```

```go
package main

import (
    "log"
    gf "github.com/felixgeelhaar/GopherFrame"
)

func main() {
    // Read from Parquet
    df, err := gf.ReadParquet("sales.parquet")
    if err != nil {
        log.Fatal(err)
    }
    defer df.Release()

    // Transform: filter, compute, aggregate
    result := df.Filter(df.Col("amount").Gt(gf.Lit(0.0)))
    defer result.Release()

    withTax := result.WithColumn("tax",
        result.Col("amount").Mul(gf.Lit(0.08)),
    )
    defer withTax.Release()

    summary := withTax.GroupBy("region").Agg(
        gf.Sum("amount").As("total_revenue"),
        gf.Mean("amount").As("avg_order_value"),
        gf.Count("order_id").As("order_count"),
    )
    defer summary.Release()

    // Save results
    if err := gf.WriteParquet(summary, "regional_summary.parquet"); err != nil {
        log.Fatal(err)
    }
}
```

## ✨ Features

### Core Operations
- ✅ **DataFrame/Series**: Immutable, strongly-typed structures with Apache Arrow backend
- ✅ **High-Performance I/O**: Parquet (5-10M rows/sec), CSV, Arrow IPC, JSON/NDJSON
- ✅ **Transformations**: Select, Filter, WithColumn, Sort (single/multi-column)
- ✅ **Joins**: InnerJoin, LeftJoin, RightJoin, FullOuterJoin, CrossJoin with hash-based O(n+m) implementation
- ✅ **GroupBy/Aggregation**: Sum, Mean, Count, Min, Max, Percentile, Median, Mode, Correlation, Variance, StdDev
- ✅ **Window Functions**: RowNumber, Rank, DenseRank, Lag, Lead, RollingSum/Mean/Min/Max, CumSum/Max/Min/Prod
- ✅ **Temporal Operations**: Year, Month, Day, Hour extraction; date truncation; date arithmetic
- ✅ **String Operations**: Upper, Lower, Trim, Length, Match (regex), Contains, StartsWith, EndsWith, Split, Replace, Pad
- ✅ **Expression Engine**: Type-safe column operations and predicates
- ✅ **UDFs**: Scalar and vectorized user-defined functions
- ✅ **Pivot/Unpivot**: Wide-to-long and long-to-wide transformations

### Advanced Features
- ✅ **Join Strategies**: Hash join, merge join (sorted data), broadcast join (small tables), auto-select
- ✅ **Multi-Column Joins**: InnerJoinMulti, LeftJoinMulti, RightJoinMulti, FullOuterJoinMulti
- ✅ **Custom Aggregations**: CustomAgg with user-defined functions, ConcatAgg for strings
- ✅ **Data Quality**: Validate (NotNull, Positive, InRange, UniqueValues), Describe, NullCount, IsComplete
- ✅ **Anomaly Detection**: DetectOutliersIQR, DetectOutliersZScore
- ✅ **Date Parsing**: ParseDateColumn with automatic format inference, DateRange, BusinessDays
- ✅ **Cross-Tabulation**: CrossTab for contingency tables
- ✅ **Query Optimization**: QueryPlan with filter pushdown, constant folding, CSE
- ✅ **Parallel Execution**: ParallelOps, ParallelAgg, ReadCSVParallel, ReadJSONParallel
- ✅ **Resource Forecasting**: EstimateResources, WillFitInMemory

### I/O Formats
- ✅ **Parquet**: High-performance columnar (5-10M rows/sec)
- ✅ **CSV**: With type inference and chunked streaming
- ✅ **Arrow IPC**: Zero-copy inter-process communication
- ✅ **JSON/NDJSON**: Array-of-objects and newline-delimited
- ✅ **Avro**: Object Container Format (OCF)
- ✅ **SQL**: ReadSQL/WriteSQL via database/sql (PostgreSQL, MySQL, SQLite)
- ✅ **Partitioned**: Hive-style partitioned datasets with pruning

### Production Features
- 🔒 **Memory Management**: LimitedAllocator with configurable limits and OOM prevention
- 📊 **Memory Monitoring**: Real-time pressure tracking (low/medium/high/critical)
- 🛡️ **Resource Safety**: Reference counting with explicit Release() for deterministic cleanup
- ⚡ **Zero-Copy Operations**: Column selection in O(1) constant time (~700ns)
- 🔍 **Security Hardened**: Path traversal protection, gosec validated, nox SCA scanned
- 🌊 **Streaming**: ReadCSVChunked, ReadCSVStreaming with backpressure

### Developer Experience
- 📚 **Comprehensive Documentation**: User guide, API reference, migration guides, troubleshooting
- 🎯 **12 Example Programs**: ETL, ML pipeline, analytics, integration demo, and more
- 🧪 **438 Tests**: Unit, integration, property-based, fuzz, chaos engineering
- 📈 **Performance Validated**: 2-428x faster than Gota (verified benchmarks)
- 🔄 **CI/CD**: Cross-platform (Linux + macOS), Go 1.24-1.26, SHA-pinned actions, coverctl + nox

## 📊 Performance

GopherFrame delivers exceptional performance through Apache Arrow's columnar memory layout and zero-copy operations:

| Operation | GopherFrame | Gota | **Speedup** | Use Case |
|-----------|-------------|------|-------------|----------|
| **Select (10K)** | 772ns | 52.4ms | **🚀 67.8x** | Column projection |
| **Column Access** | 120ns | 3.4µs | **🚀 28x** | Single column lookup |
| **Iteration (1K)** | 389ns | 166µs | **🚀 428x** | Row-by-row processing |
| **Filter (10K)** | 404µs | 748µs | **1.9x** | Conditional filtering |
| **Creation (10K)** | 249µs | 747µs | **3.0x** | DataFrame construction |

**Memory Efficiency**: 2-200x less memory usage across all operations

📈 **[Full Benchmark Report](docs/GOTA_COMPARISON_BENCHMARKS.md)** | **[Benchmarks](BENCHMARKS.md)**

### Why So Fast?

1. **O(1) Column Selection**: Zero-copy pointer updates vs O(n) data copying
2. **Columnar Memory Layout**: CPU cache-optimized contiguous storage
3. **Apache Arrow Integration**: Native support for vectorized operations
4. **Minimal Abstraction**: Direct access to Arrow arrays, no unnecessary conversions

## 🎓 Migration Guides

Comprehensive guides for migrating from popular data manipulation libraries:

- **[From Pandas](docs/MIGRATION_FROM_PANDAS.md)** - Python/Pandas users
- **[From Polars](docs/MIGRATION_FROM_POLARS.md)** - Polars (Arrow-to-Arrow migration)
- **[From Gota](docs/MIGRATION_FROM_GOTA.md)** - Go/Gota users (with performance comparison)

Each guide includes:
- Side-by-side code comparisons
- Common operation translations
- Migration patterns for ETL, ML preprocessing
- Performance considerations
- Complete example migrations

## 📖 Example Programs

Production-ready examples demonstrating real-world use cases:

### [ETL Pipeline](cmd/examples/etl_pipeline/main.go)
```go
// Load, clean, transform, join, aggregate, and save
orders := gf.ReadCSV("orders.csv")
validOrders := orders.Filter(orders.Col("amount").Gt(gf.Lit(0.0)))
withTax := validOrders.WithColumn("tax", ...)
merged := withTax.InnerJoin(customers, "customer_id", "customer_id")
summary := merged.GroupBy("region").Agg(...)
gf.WriteParquet(summary, "regional_summary.parquet")
```

### [ML Preprocessing](cmd/examples/ml_preprocessing/main.go)
```go
// Data quality checks, feature engineering, train/test split
validData := rawDF.Filter(rawDF.Col("score").Lt(gf.Lit(101.0)))
withFeatures := validData.WithColumn("bonus", ...)
features := withFeatures.Select("feature1", "feature2", ..., "target")
gf.WriteParquet(features, "train_features.parquet")
```

### [Backend Analytics](cmd/examples/backend_analytics/main.go)
```go
// Real-time API monitoring and performance metrics
perfMetrics := logs.GroupBy("endpoint").Agg(
    gf.Mean("response_time_ms").As("avg_response"),
    gf.Count("request_id").As("request_count"),
)
errors := logs.Filter(logs.Col("status_code").Gt(gf.Lit(int64(399))))
```

### [Production Memory Management](cmd/examples/production_memory/main.go)
```go
// Memory limits, pressure monitoring, OOM handling
limited := core.NewLimitedAllocator(base, 512*1024*1024) // 512MB limit
df := gf.NewDataFrameWithAllocator(record, limited)

if limited.MemoryPressureLevel() == "critical" {
    // Handle memory pressure
}
```

## 🛠️ Production Deployment

### Memory Management

GopherFrame provides production-grade memory management to prevent OOM crashes:

```go
import "github.com/felixgeelhaar/GopherFrame/pkg/core"

// Configure 1GB memory limit
pool := memory.NewGoAllocator()
limited := core.NewLimitedAllocator(pool, 1024*1024*1024)

// Create DataFrame with limited allocator
df := gf.NewDataFrameWithAllocator(record, limited)
defer df.Release()

// Pre-flight check before expensive operations
if err := limited.CheckCanAllocate(estimatedSize); err != nil {
    // Handle insufficient memory
}

// Monitor memory pressure
switch limited.MemoryPressureLevel() {
case "critical":
    return errors.New("memory critical, aborting")
case "high":
    batchSize = batchSize / 2  // Reduce batch size
}
```

**[Production Memory Management Guide](docs/PRODUCTION_MEMORY_MANAGEMENT.md)**

### Best Practices

1. **Always `defer df.Release()`** - Explicit resource management
2. **Check errors** - Handle `df.Err()` after operations
3. **Use type-safe literals** - `gf.Lit(int64(18))` not `gf.Lit(18)`
4. **Set memory limits** - Use LimitedAllocator in production
5. **Monitor pressure** - React to memory pressure levels

## 📚 Documentation

### Getting Started
- **[Quick Start Guide](#-quick-start)** - Get up and running in 5 minutes
- **[Example Programs](cmd/examples/)** - Production-ready examples
- **[API Reference](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame)** - Complete API documentation

### Migration
- **[From Pandas](docs/MIGRATION_FROM_PANDAS.md)** - Python/Pandas → GopherFrame
- **[From Polars](docs/MIGRATION_FROM_POLARS.md)** - Polars → GopherFrame
- **[From Gota](docs/MIGRATION_FROM_GOTA.md)** - Gota → GopherFrame

### Performance
- **[Benchmarks](BENCHMARKS.md)** - Performance characteristics
- **[Gota Comparison](docs/GOTA_COMPARISON_BENCHMARKS.md)** - Detailed comparison
- **[Benchmark Regression Testing](docs/BENCHMARK_REGRESSION_TESTING.md)** - CI/CD performance gates

### Production
- **[Production Memory Management](docs/PRODUCTION_MEMORY_MANAGEMENT.md)** - Memory limits and OOM handling
- **[Technical Design](docs/technical_design_doc.md)** - Architecture and design decisions

## 🎯 Use Cases

### When to Use GopherFrame

✅ **Perfect For:**
- Production Go services requiring data processing
- High-performance ETL/ELT pipelines
- Real-time analytics in Go backends
- ML model inference preprocessing (identical to Python training)
- Memory-constrained environments (2-200x less memory)
- Arrow ecosystem integration (Parquet, Flight RPC)

### When to Use Alternatives

**Use Pandas** when:
- Rapid prototyping in Jupyter notebooks
- Rich Python ecosystem integration (matplotlib, scikit-learn)
- Team expertise is primarily Python-based

**Use Polars** when:
- Python-based ML/data science workflows
- Need lazy evaluation and query optimization
- Team prefers Python over Go

**Use Gota** when:
- Simple data manipulation in Go
- Arrow dependency is unacceptable
- Performance is not critical

## 🏗️ Architecture

```
GopherFrame/
├── cmd/
│   ├── examples/          # Production-ready example programs
│   └── benchmark/         # Performance benchmarking tool
├── pkg/
│   ├── core/             # DataFrame/Series implementations
│   ├── expr/             # Expression engine and AST
│   ├── domain/           # Domain-driven design components
│   ├── application/      # Application services
│   ├── infrastructure/   # I/O and persistence
│   └── storage/          # Pluggable storage backends
├── docs/                 # Comprehensive documentation
└── .github/workflows/    # CI/CD with benchmark regression
```

**Design Principles:**
- **Clean Architecture**: Clear separation of concerns
- **Domain-Driven Design**: Rich domain models
- **Apache Arrow Foundation**: Zero-copy operations throughout
- **Production-First**: Every feature designed for production use

## 🧪 Development

### Setup

```bash
git clone https://github.com/felixgeelhaar/GopherFrame.git
cd GopherFrame
go mod download
```

### Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. -benchmem ./pkg/core

# Memory leak detection
go test -v -run TestMemoryLeak ./pkg/core
```

### Benchmarking

```bash
# Compare with Gota
go test -bench=. -benchmem ./pkg/core

# Run with benchstat for statistical comparison
go test -bench=. -count=5 ./pkg/core > new.txt
benchstat base.txt new.txt
```

## 🤝 Contributing

We welcome contributions! GopherFrame follows:

- **TDD**: Test-Driven Development with Red-Green-Refactor
- **Clean Code**: SOLID principles, readable, maintainable
- **Benchmark Regression**: Automated performance testing in CI
- **Documentation**: Every public API has godoc

**Before submitting a PR:**
1. Run tests: `go test ./...`
2. Run benchmarks: `go test -bench=. ./pkg/core`
3. Check for regressions: CI will compare with base branch
4. Update documentation if needed

**[Benchmark Regression Testing Guide](docs/BENCHMARK_REGRESSION_TESTING.md)**

## 📊 Project Status

🚀 **v1.0+ Feature Complete — All Roadmap Phases Delivered**

**Quality Metrics:**
- 438 tests passing, zero race conditions
- 81.4% pkg/ coverage (coverctl validated)
- Zero critical security findings (nox SCA + gosec)
- Cross-platform CI (Linux + macOS, Go 1.24-1.26)
- All GitHub Actions pinned to commit SHAs

**[Full Roadmap](ROADMAP.md)** — 96/96 items complete

## 📚 Documentation

| Guide | Description |
|-------|-------------|
| [User Guide](docs/USER_GUIDE.md) | Tutorials, recipes, and patterns |
| [API Reference](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame) | Complete API documentation |
| [Performance Guide](docs/PERFORMANCE_GUIDE.md) | Optimization and profiling |
| [Migration from Pandas](docs/MIGRATION_FROM_PANDAS.md) | Python/Pandas users |
| [Migration from Polars](docs/MIGRATION_FROM_POLARS.md) | Polars users |
| [Migration from Gota](docs/MIGRATION_FROM_GOTA.md) | Gota users |
| [Troubleshooting](docs/TROUBLESHOOTING.md) | Common issues and solutions |
| [Plugin API](docs/PLUGIN_API.md) | Extension and plugin development |
| [API Stability](docs/API_STABILITY.md) | Versioning and compatibility |
| [Technical Design](docs/technical_design_doc.md) | Architecture and ADRs |
| [Contributing](CONTRIBUTING.md) | How to contribute |
| [Security](SECURITY.md) | Vulnerability reporting |
| [Changelog](CHANGELOG.md) | Release history |

**[Full Roadmap](ROADMAP.md)**

## 📄 License

Apache 2.0 License - see [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built on [Apache Arrow Go](https://github.com/apache/arrow-go) for zero-copy operations
- Inspired by [Polars](https://github.com/pola-rs/polars) performance and design
- Benchmarked against [Gota](https://github.com/go-gota/gota) for Go DataFrame operations
- Following Go's principles of simplicity, explicitness, and composition

## 📧 Contact

- **Issues**: [GitHub Issues](https://github.com/felixgeelhaar/GopherFrame/issues)
- **Discussions**: [GitHub Discussions](https://github.com/felixgeelhaar/GopherFrame/discussions)
- **Author**: Felix Geelhaar ([@felixgeelhaar](https://github.com/felixgeelhaar))

---

**Built with ❤️ for the Go community**

*Making data processing in Go fast, safe, and production-ready.*
