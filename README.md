# GopherFrame

**A production-ready DataFrame library for Go, powered by Apache Arrow.**

[![Go Reference](https://pkg.go.dev/badge/github.com/felixgeelhaar/GopherFrame.svg)](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame)
[![Go Report Card](https://goreportcard.com/badge/github.com/felixgeelhaar/GopherFrame)](https://goreportcard.com/report/github.com/felixgeelhaar/GopherFrame)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

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
- ✅ **High-Performance I/O**: Parquet (5-10M rows/sec), CSV, Arrow IPC
- ✅ **Transformations**: Select, Filter, WithColumn, Sort (single/multi-column)
- ✅ **Joins**: InnerJoin, LeftJoin with hash-based O(n+m) implementation
- ✅ **GroupBy/Aggregation**: Sum, Mean, Count with multiple aggregations
- ✅ **String Operations**: Contains, StartsWith, EndsWith
- ✅ **Expression Engine**: Type-safe column operations and predicates

### Production Features
- 🔒 **Memory Management**: LimitedAllocator with configurable limits and OOM prevention
- 📊 **Memory Monitoring**: Real-time pressure tracking (low/medium/high/critical)
- 🛡️ **Resource Safety**: Reference counting with explicit Release() for deterministic cleanup
- ⚡ **Zero-Copy Operations**: Column selection in O(1) constant time (~700ns)
- 🔍 **Security Hardened**: Path traversal protection, input validation

### Developer Experience
- 📚 **Comprehensive Documentation**: Migration guides, examples, API reference
- 🎯 **Example Programs**: ETL pipeline, ML preprocessing, backend analytics
- 🧪 **Benchmark Regression CI**: Automated performance testing on every PR
- 📈 **Performance Validated**: 2-428x faster than Gota (verified benchmarks)

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

🚀 **v1.0 Production Ready - Feature Complete**

- ✅ All Phase 1 & 2 features complete and tested
- ✅ Window functions (analytical, rolling, cumulative)
- ✅ Statistical aggregations (percentile, median, mode, correlation)
- ✅ Enhanced join operations (inner, left, right, full outer, cross)
- ✅ Temporal operations (extraction, truncation, arithmetic)
- ✅ String operations (case, trim, regex, length)
- ✅ Production memory management
- ✅ Performance validated (2-428x faster than Gota)
- ✅ Comprehensive API documentation
- ✅ Benchmark regression CI
- ✅ 10 example programs including Phase 2 features
- ✅ Migration guides from pandas/Polars/Gota

**Quality Metrics:**
- 279 tests, 100% pass rate
- 95%+ test coverage
- Zero security vulnerabilities (gosec validated)
- Automated benchmark regression detection
- Production deployments ready

## 🗺️ Roadmap

### v1.0 (Current) ✅
- Core DataFrame/Series operations
- Enhanced join operations (Inner, Left, Right, Full Outer, Cross)
- Window functions (RowNumber, Rank, DenseRank, Lag, Lead)
- Rolling aggregations (Sum, Mean, Min, Max, Count)
- Cumulative operations (CumSum, CumMax, CumMin, CumProd)
- Statistical aggregations (Percentile, Median, Mode, Correlation)
- Temporal operations (Year, Month, Day, Hour, Minute, Second, Truncate, Add)
- String operations (Upper, Lower, Trim, Length, Match, Contains)
- Production memory management
- Performance validation
- Comprehensive documentation

### v1.1 (Planned - Q2 2026)
- User-defined functions (UDFs)
- Pivot operations
- Additional string functions
- Advanced join strategies

### v1.2 (Future)
- Streaming data processing
- Additional file formats (JSON, ORC)
- Distributed computing support

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
