# GopherFrame

**A production-first DataFrame library for Go, powered by Apache Arrow.**

GopherFrame is an Apache Arrow-backed DataFrame library designed from the ground up for production use. It bridges the gap between Python's rich data science ecosystem and Go's performance and operational simplicity.

## Project Status

🚧 **Early Development** - Target v0.1 release: Q1 2025

This is a greenfield project currently in active development. The API is not stable and breaking changes are expected.

## Vision

Empower Go developers to build fast, reliable, and scalable data engineering pipelines without leaving the Go ecosystem.

### Core Principles

- **Production-First**: Every API decision optimized for performance, reliability, and production use
- **Apache Arrow Integration**: Native zero-copy interoperability with the modern data ecosystem
- **Idiomatic Go**: Strongly-typed API using generics, explicit error handling, composable design
- **No Compromises**: Real algorithms, real data, real performance from day one

## Planned Features (v0.1 MVP)

- **Core DataFrame/Series**: Immutable, strongly-typed structures backed by Arrow
- **High-Performance I/O**: Concurrent Parquet, CSV, and Arrow IPC readers/writers
- **Expression Engine**: Lazy evaluation with chainable operations
- **Data Transformations**: Selection, filtering, column operations, sorting
- **Aggregations**: GroupBy operations with multiple aggregation functions

### Example Usage (Planned API)

```go
package main

import (
    "github.com/felixgeelhaar/gopherFrame"
)

func main() {
    // Read from Parquet with zero-copy efficiency
    df, err := gopherframe.ReadParquet("data.parquet")
    if err != nil {
        panic(err)
    }

    // Chain transformations with lazy evaluation
    result := df.
        Filter(df.Col("country").Eq(df.Lit("US"))).
        WithColumn("profit", df.Col("revenue").Sub(df.Col("cost"))).
        GroupBy("region").
        Aggregate(
            df.Sum("profit").As("total_profit"),
            df.Mean("revenue").As("avg_revenue"),
        ).
        Sort(df.By("total_profit", df.Desc))

    if result.Err() != nil {
        panic(result.Err())
    }

    // Write results efficiently
    err = result.WriteParquet("output.parquet")
    if err != nil {
        panic(err)
    }
}
```

## Development Setup

### Prerequisites

- Go 1.22+
- Apache Arrow Go v18.0.0 (automatically managed)

### Quick Setup

```bash
# Clone and setup development environment
git clone https://github.com/felixgeelhaar/gopherFrame.git
cd gopherFrame

# Install pre-commit hooks and development tools
./scripts/setup-hooks.sh

# Verify installation
go test ./...
```

The setup script installs:
- Pre-commit hooks (automatic formatting and linting)
- golangci-lint (code quality)
- goimports (import management)

### Building

```bash
git clone https://github.com/felixgeelhaar/gopherFrame.git
cd gopherFrame
go mod tidy
go build ./...
```

### Testing

```bash
# Unit tests
go test ./...

# Benchmarks
go test -bench=. ./...

# Integration tests with real data
go test -tags=integration ./...
```

## Architecture

The library follows a clean three-layer architecture:

```
pkg/
├── core/          # Internal DataFrame/Series implementations
├── expr/          # Expression engine and AST structures  
└── io/            # Parquet, CSV, and Arrow IPC readers/writers
```

- **Public API Layer**: User-facing DataFrame and Series types
- **Core Engine**: Internal implementations managing transformations
- **Subsystems**: Expression evaluation, I/O operations, Arrow integration

## Target Users

- **Data Engineers**: Building production ETL/ELT pipelines
- **ML Engineers**: Deploying models requiring identical data transformations between Python training and Go inference  
- **Go Developers**: Building backend services with non-trivial data analysis requirements

## Performance Goals

- **10x faster** than existing Go DataFrame libraries (Gota)
- **Competitive** with Python Polars on multi-core hardware
- **Zero-copy** interoperability with Python ecosystem via Arrow
- **Multi-gigabyte** dataset processing capabilities

## Contributing

This project is in early development. Contributions and feedback are welcome, but please note that APIs may change significantly before v1.0.

## License

Apache 2.0 License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built on the excellent [Apache Arrow Go](https://github.com/apache/arrow-go) library
- Inspired by the performance and design of [Polars](https://github.com/pola-rs/polars)
- Following Go's principles of simplicity and explicitness