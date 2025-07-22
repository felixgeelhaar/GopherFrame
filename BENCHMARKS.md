# GopherFrame Performance Benchmarks

## Overview

GopherFrame is designed for production-first performance with Apache Arrow as its foundation. These benchmarks demonstrate the library's performance characteristics across various operations and data sizes.

## Benchmark Results

### Core Operations Performance

Performance measured on Apple M1 processor:

| Operation | 1K rows | 10K rows | 100K rows | Throughput (rows/sec) |
|-----------|---------|----------|-----------|----------------------|
| Filter | 178µs | 882µs | 5.4ms | 18.5M |
| Select | 5.5µs | 6.5µs | 8.5µs | 11.7M+ |
| WithColumn | 35µs | 263µs | 1.8ms | 55.5K |
| GroupBy+Sum | 125µs | 704µs | 5.9ms | 16.9K |

### I/O Performance (10K rows)

| Format | Write Time | Read Time | Write Speed | Read Speed |
|--------|------------|-----------|-------------|------------|
| Parquet | 3.2ms | 1.3ms | 3.1M rows/s | 7.7M rows/s |
| CSV | 13.0ms | 1.9ms | 769K rows/s | 5.3M rows/s |
| Arrow IPC | ~1ms | ~1ms | 10M+ rows/s | 10M+ rows/s |

## Key Performance Characteristics

### 1. **Linear Scaling**
- Most operations scale linearly with data size
- Filter operation shows excellent scalability with increasing throughput on larger datasets

### 2. **Memory Efficiency**
- Leverages Apache Arrow's columnar format for cache-friendly access
- Zero-copy operations where possible (especially with Arrow IPC)
- Reference counting prevents unnecessary memory allocations

### 3. **I/O Performance**
- **Parquet**: Optimized for analytical workloads, excellent compression
- **CSV**: Good performance despite type inference overhead
- **Arrow IPC**: Near-zero overhead for Arrow-native data exchange

## Running Benchmarks

### Standard Go Benchmarks
```bash
# Run all benchmarks
go test -bench=. -benchtime=10s

# Run specific benchmark
go test -bench=BenchmarkFilter -benchtime=10s

# Memory profiling
go test -bench=BenchmarkMemoryUsage -benchmem
```

### Comprehensive Benchmark Suite
```bash
go run cmd/benchmark/main.go
```

## Optimization Strategies

### 1. **Arrow-Native Operations**
- Direct manipulation of Arrow arrays without intermediate conversions
- Leverages Arrow's optimized memory layout

### 2. **Lazy Evaluation**
- Operations build an execution plan rather than immediate execution
- Allows for query optimization in future versions

### 3. **Batch Processing**
- GroupBy operations process data in batches
- Minimizes memory allocations during aggregation

### 4. **Type-Specific Implementations**
- Specialized code paths for common types (Int64, Float64, String)
- Avoids reflection and interface conversions in hot paths

## Comparison with Other Libraries

While direct comparisons depend on specific use cases, GopherFrame aims to be:
- **10x faster than Gota** for common operations
- **Competitive with Python Polars** for data transformation tasks
- **Native Go solution** avoiding CGO overhead

## Future Optimizations

1. **Parallel Execution**: Leverage Go's concurrency for multi-core processing
2. **SIMD Operations**: Use Arrow's compute kernels for vectorized operations
3. **Query Optimization**: Implement predicate pushdown and column pruning
4. **Memory Pool Management**: Custom allocators for better cache locality

## Benchmark Development

The benchmark suite (`benchmark_test.go`) includes:
- Core operation benchmarks
- I/O performance tests
- Memory allocation tracking
- Chained operation performance
- Various data sizes (1K, 10K, 100K rows)

To add new benchmarks:
1. Add benchmark function to `benchmark_test.go`
2. Follow naming convention: `BenchmarkOperationName`
3. Test with multiple data sizes
4. Include memory allocation metrics where relevant