# GopherFrame v0.1 Release Notes

**Release Date**: January 2025
**Status**: Production Ready
**Download**: `go get github.com/felixgeelhaar/GopherFrame@v0.1.0`

---

## 🎉 What's New in v0.1

GopherFrame v0.1 marks the first production-ready release of a high-performance DataFrame library for Go, built on Apache Arrow. This release delivers on our promise of a **production-first** DataFrame library with validated performance claims, comprehensive documentation, and enterprise-grade features.

### Performance Validated: 2-428x Faster Than Gota

We've conducted comprehensive benchmarks comparing GopherFrame to Gota (the leading Go DataFrame library):

| Operation | GopherFrame | Gota | **Speedup** | Use Case |
|-----------|-------------|------|-------------|----------|
| **Select (10K)** | 772ns | 52.4ms | **🚀 67.8x** | Column projection |
| **Column Access** | 120ns | 3.4µs | **🚀 28x** | Single column lookup |
| **Iteration (1K)** | 389ns | 166µs | **🚀 428x** | Row-by-row processing |
| **Filter (10K)** | 404µs | 748µs | **1.9x** | Conditional filtering |
| **Creation (10K)** | 249µs | 747µs | **3.0x** | DataFrame construction |

**Memory Efficiency**: 2-200x less memory usage across all operations

📊 **[Full Benchmark Report](docs/GOTA_COMPARISON_BENCHMARKS.md)**

---

## ✨ Key Features

### Core DataFrame Operations

- ✅ **DataFrame/Series**: Immutable, strongly-typed structures with Apache Arrow backend
- ✅ **High-Performance I/O**: Parquet (5-10M rows/sec), CSV, Arrow IPC
- ✅ **Transformations**: Select, Filter, WithColumn, Sort (single/multi-column)
- ✅ **Joins**: InnerJoin, LeftJoin with hash-based O(n+m) implementation
- ✅ **GroupBy/Aggregation**: Sum, Mean, Count with multiple aggregations
- ✅ **String Operations**: Contains, StartsWith, EndsWith
- ✅ **Expression Engine**: Type-safe column operations and predicates

### Production-Grade Features

#### Memory Management (NEW in v0.1)

Production-ready memory management system with configurable limits and OOM prevention:

```go
import "github.com/felixgeelhaar/GopherFrame/pkg/core"

// Configure 512MB memory limit
pool := memory.NewGoAllocator()
limited := core.NewLimitedAllocator(pool, 512*1024*1024)

df := gf.NewDataFrameWithAllocator(record, limited)
defer df.Release()

// Monitor memory pressure
switch limited.MemoryPressureLevel() {
case "critical":
    return errors.New("memory critical, aborting")
case "high":
    batchSize = batchSize / 2  // Reduce batch size
}

// Pre-flight check before expensive operations
if err := limited.CheckCanAllocate(estimatedSize); err != nil {
    // Handle insufficient memory
}
```

**Features**:
- 🔒 **Thread-safe** atomic memory tracking
- 📊 **Memory pressure monitoring** (low/medium/high/critical levels)
- 🛡️ **Pre-flight allocation checks** to prevent OOM
- ⚡ **Graceful degradation** for memory-constrained environments

📚 **[Production Memory Management Guide](docs/PRODUCTION_MEMORY_MANAGEMENT.md)**

#### Benchmark Regression Testing (NEW in v0.1)

Automated performance monitoring to prevent regressions:

- **PR-based comparison**: Automatic benchmark comparison on every pull request
- **Statistical analysis**: Uses Go's official `benchstat` tool with p-value testing
- **10% regression threshold**: CI fails if performance degrades >10%
- **PR comments**: Detailed performance comparison posted to pull requests
- **Historical tracking**: Performance trends tracked in GitHub Pages

```yaml
# Example PR comment
## 📊 Benchmark Comparison Results

name                old time/op    new time/op    delta
Filter_10K          404µs ± 1%     380µs ± 2%    -5.94%  ✅
Select_10K          772ns ± 5%     801ns ± 4%    +3.76%  ⚠️

### Performance Regression Policy:
- ✅ No action if ~0% or improvements
- ⚠️  Review required if 5-10%
- ❌ Must fix if >10%
```

📚 **[Benchmark Regression Testing Guide](docs/BENCHMARK_REGRESSION_TESTING.md)**

#### Security & Reliability

- 🔒 **Zero security vulnerabilities** (gosec validated)
- 🛡️ **Path traversal protection** for file I/O operations
- ⚡ **Memory-safe resource management** with explicit Release()
- 🧪 **200+ tests** with 82-86% code coverage
- ✅ **100% test pass rate** with comprehensive edge case coverage

---

## 📚 Documentation

### Migration Guides (NEW in v0.1)

Comprehensive guides to help you migrate from other DataFrame libraries:

#### [From Pandas](docs/MIGRATION_FROM_PANDAS.md) (713 lines)
Complete migration guide for Python/Pandas users moving to Go:
- Side-by-side code comparisons for all common operations
- DataFrame creation, selection, filtering, sorting, groupby, joins
- ETL pipeline migration patterns
- ML preprocessing workflow examples
- Common gotchas (memory management, null handling, type conversion)
- Complete example migration from Pandas to GopherFrame

#### [From Polars](docs/MIGRATION_FROM_POLARS.md) (592 lines)
Arrow-to-Arrow migration guide for Polars users:
- Shared Apache Arrow foundation between both libraries
- Lazy vs eager evaluation differences
- Python-to-Go adaptation patterns
- When to use each library (production Go vs Python ML)
- Complete migration examples

#### [From Gota](docs/MIGRATION_FROM_GOTA.md) (784 lines)
Performance-focused migration guide for Gota users:
- Validated benchmark comparisons (2-428x faster)
- Operation-by-operation translation guide
- Memory management differences (GC only vs reference counting)
- Performance migration strategy (identify hot paths first)
- Complete example migration with performance improvements

### Example Programs (NEW in v0.1)

Production-ready examples demonstrating real-world use cases:

#### [ETL Pipeline](cmd/examples/etl_pipeline/main.go) (279 lines)
Complete data engineering workflow:
```go
// Load, clean, transform, join, aggregate, and save
orders := gf.ReadCSV("orders.csv")
validOrders := orders.Filter(orders.Col("amount").Gt(gf.Lit(0.0)))
withTax := validOrders.WithColumn("tax", ...)
merged := withTax.InnerJoin(customers, "customer_id", "customer_id")
summary := merged.GroupBy("region").Agg(...)
gf.WriteParquet(summary, "regional_summary.parquet")
```

#### [ML Preprocessing](cmd/examples/ml_preprocessing/main.go) (275 lines)
Machine learning data preparation pipeline:
```go
// Data quality checks, feature engineering, train/test split
validData := rawDF.Filter(rawDF.Col("score").Lt(gf.Lit(101.0)))
withFeatures := validData.WithColumn("bonus", ...)
features := withFeatures.Select("feature1", "feature2", ..., "target")
gf.WriteParquet(features, "train_features.parquet")
```

#### [Backend Analytics](cmd/examples/backend_analytics/main.go) (284 lines)
Real-time API monitoring and performance metrics:
```go
// Real-time API monitoring and performance metrics
perfMetrics := logs.GroupBy("endpoint").Agg(
    gf.Mean("response_time_ms").As("avg_response"),
    gf.Count("request_id").As("request_count"),
)
errors := logs.Filter(logs.Col("status_code").Gt(gf.Lit(int64(399))))
```

#### [Production Memory Management](cmd/examples/production_memory/main.go) (213 lines)
Memory limits, pressure monitoring, OOM handling:
```go
// Memory limits, pressure monitoring, OOM handling
limited := core.NewLimitedAllocator(base, 512*1024*1024) // 512MB limit
df := gf.NewDataFrameWithAllocator(record, limited)

if limited.MemoryPressureLevel() == "critical" {
    // Handle memory pressure
}
```

---

## 🎯 Use Cases

### Perfect For:

✅ **Production Go services** requiring data processing
✅ **High-performance ETL/ELT pipelines** (5-10M rows/sec I/O)
✅ **Real-time analytics** in Go backends
✅ **ML model inference preprocessing** (identical to Python training)
✅ **Memory-constrained environments** (2-200x less memory than alternatives)
✅ **Arrow ecosystem integration** (Parquet, Flight RPC)

### When to Use Alternatives:

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

---

## 📦 Installation & Getting Started

### Installation

```bash
go get github.com/felixgeelhaar/GopherFrame@v0.1.0
```

### Quick Start

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

---

## 🔄 Upgrade Guide

### From Pre-release Versions

If you've been using development versions, upgrade to v0.1.0:

```bash
go get -u github.com/felixgeelhaar/GopherFrame@v0.1.0
go mod tidy
```

### Breaking Changes

**None** - This is the first production release. All future v0.x releases will maintain API compatibility.

### Best Practices for v0.1

1. **Always `defer df.Release()`** - Explicit resource management prevents memory leaks
2. **Check errors** - Handle `df.Err()` after operations
3. **Use type-safe literals** - `gf.Lit(int64(18))` not `gf.Lit(18)`
4. **Set memory limits** - Use `LimitedAllocator` in production (see examples)
5. **Monitor pressure** - React to memory pressure levels (`low/medium/high/critical`)

---

## 📊 Quality Metrics

- ✅ **200+ tests** with 100% pass rate
- ✅ **82-86% test coverage** across core packages
- ✅ **Zero security vulnerabilities** (gosec validated)
- ✅ **Automated benchmark regression** detection in CI
- ✅ **Cross-platform testing** (Linux, macOS, Windows)
- ✅ **Go version support**: 1.21, 1.22, 1.23

---

## 🗺️ Roadmap

### v0.2 (Planned - Q1 2025)

- Window functions and rolling aggregations
- Advanced date/time operations
- Enhanced string operations (regex, splitting)
- Additional aggregations (percentile, correlation, covariance)

### v0.3 (Future)

- User-defined functions (UDFs)
- Streaming operations
- Partitioned datasets
- SQL interface

**[Full Roadmap](ROADMAP.md)**

---

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

---

## 🙏 Acknowledgments

- Built on [Apache Arrow Go](https://github.com/apache/arrow-go) for zero-copy operations
- Inspired by [Polars](https://github.com/pola-rs/polars) performance and design
- Benchmarked against [Gota](https://github.com/go-gota/gota) for Go DataFrame operations
- Following Go's principles of simplicity, explicitness, and composition

---

## 📧 Contact & Support

- **Issues**: [GitHub Issues](https://github.com/felixgeelhaar/GopherFrame/issues)
- **Discussions**: [GitHub Discussions](https://github.com/felixgeelhaar/GopherFrame/discussions)
- **Author**: Felix Geelhaar ([@felixgeelhaar](https://github.com/felixgeelhaar))

---

## 📄 License

Apache 2.0 License - see [LICENSE](LICENSE) file for details.

---

**Built with ❤️ for the Go community**

*Making data processing in Go fast, safe, and production-ready.*
