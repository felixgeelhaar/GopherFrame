# GopherFrame Performance Benchmarks

## Overview

GopherFrame is designed for production-first performance with Apache Arrow as its foundation. These benchmarks demonstrate the library's performance characteristics across various operations and data sizes.

## Benchmark Results

**Last Updated**: October 4, 2025
**Test Environment**: Apple M1, Go 1.23, Darwin ARM64

### Core Operations Performance

| Operation | 1K rows | 10K rows | 100K rows | Allocs (100K) | Memory (100K) |
|-----------|---------|----------|-----------|---------------|---------------|
| DataFrame Creation | 25µs | 257µs | 2.1ms | 111 | 5.8 MB |
| Filter | 35µs | 322µs | 2.9ms | 168 | 5.2 MB |
| Select | **765ns** | **751ns** | **687ns** | 17 | 1.6 KB |
| WithColumn | 16µs | 154µs | 1.1ms | 88 | 4.4 MB |
| GroupBy+Sum | 41µs | 339µs | 3.3ms | 255 | 3.6 MB |
| GroupBy (5 aggs) | 110µs | 985µs | - | 20K | 430 KB |
| Chained Ops | 60µs | 536µs | - | 223 | 1.0 MB |

**Key Observations**:
- **Select operation is O(1)**: ~700ns regardless of data size due to zero-copy columnar design
- **Linear scaling**: All operations scale linearly with data size
- **Memory efficiency**: Allocations grow linearly, leveraging Arrow's memory reuse

### I/O Performance

| Format | Operation | 1K rows | 10K rows | Throughput (10K) |
|--------|-----------|---------|----------|------------------|
| Parquet | Write | 397µs | 2.0ms | ~5M rows/s |
| Parquet | Read | 312µs | 930µs | ~10M rows/s |
| CSV | Write | 920µs | 10.4ms | ~960K rows/s |
| CSV | Read | - | - | - |

**Parquet Performance**:
- Excellent compression and read performance
- Write performance: 30-50 MB/s typical
- Read performance: 100+ MB/s typical

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

### Performance Claims Status

**⚠️ VALIDATION NEEDED**: Performance comparisons require independent verification

**Current Claims** (to be validated):
- **10x faster than Gota** for common operations - **NEEDS BENCHMARK COMPARISON**
- **Competitive with Python Polars** for data transformation - **NEEDS VALIDATION**
- **Native Go solution** - ✅ CONFIRMED (zero CGo, pure Go + Arrow Go)

### Competitive Advantages (Verified)

1. **Zero-Copy Operations**:
   - Select operation: ~700ns constant time regardless of data size
   - Arrow-native design eliminates serialization overhead

2. **Memory Efficiency**:
   - Columnar storage reduces memory footprint
   - Reference counting prevents duplicate allocations
   - Measured allocations scale linearly with data size

3. **Pure Go Implementation**:
   - No CGo overhead
   - Easy cross-compilation
   - Predictable performance characteristics

### Planned Comparison Benchmarks

To validate performance claims, we need:
1. Direct Gota comparison suite (same operations, same data)
2. Polars comparison (via Python interop or similar dataset)
3. Pandas comparison (baseline for data manipulation libraries)
4. Published results with reproducible test methodology

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