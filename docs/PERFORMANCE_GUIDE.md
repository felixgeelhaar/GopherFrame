# Performance Guide

## Architecture for Speed

GopherFrame achieves 2-428x performance over alternatives through:

1. **Apache Arrow columnar format** — CPU cache-friendly contiguous memory
2. **Zero-copy operations** — Select returns a view, not a copy
3. **Hash-based joins** — O(n+m) time complexity
4. **Vectorized processing** — operations on entire columns, not rows

## Optimization Strategies

### Memory Management

```go
// Set memory limits to prevent OOM
pool := memory.NewGoAllocator()
limited := core.NewLimitedAllocator(pool, 2*1024*1024*1024) // 2GB

// Monitor memory pressure
switch limited.MemoryPressureLevel() {
case "critical": // abort
case "high":     // reduce batch size
case "medium":   // continue with caution
case "low":      // normal operation
}
```

### Query Optimization

```go
// BAD: filter after expensive join
result := big.InnerJoin(other, "id", "id")
filtered := result.Filter(...)

// GOOD: filter before join
filtered := big.Filter(...)
result := filtered.InnerJoin(other, "id", "id")

// BEST: use QueryPlan for automatic optimization
result := gf.NewQueryPlan(big).
    Filter(big.Col("x").Gt(gf.Lit(0.0))).
    WithColumn("y", big.Col("x").Mul(gf.Lit(2.0))).
    Execute()
```

### Join Strategy Selection

| Strategy | Best For | Complexity | Memory |
|----------|----------|-----------|--------|
| Hash Join (default) | General use | O(n+m) | O(min(n,m)) |
| Merge Join | Pre-sorted data | O(n+m) | O(1) |
| Broadcast Join | Small right table | O(n*m_small) | O(m) |
| AutoJoin | Auto-select | Varies | Varies |

```go
df.AutoJoin(other, "id", "id") // Picks best strategy automatically
```

### Streaming for Large Data

```go
// Process 100K rows at a time instead of loading everything
it, _ := gf.ReadCSVChunked("10gb_file.csv", 100000)
it.ForEachChunk(func(chunk *gf.DataFrame) error {
    // Process chunk — only 100K rows in memory at a time
    return nil
})
```

### VectorUDF vs ScalarUDF

```go
// ScalarUDF: ~1M rows/sec (row-by-row with map allocation)
gf.ScalarUDF(cols, outType, func(row map[string]interface{}) (interface{}, error) { ... })

// VectorUDF: ~100M rows/sec (direct Arrow array access)
gf.VectorUDF(cols, outType, func(cols map[string]arrow.Array) (arrow.Array, error) { ... })
```

Always prefer VectorUDF for performance-critical paths.

## Benchmarks

| Operation | 1K rows | 10K rows | 100K rows |
|-----------|---------|----------|-----------|
| Select | ~700ns | ~800ns | ~900ns |
| Filter | ~40µs | ~400µs | ~4ms |
| GroupBy+Sum | ~100µs | ~500µs | ~5ms |
| InnerJoin | ~85µs | ~1ms | ~10ms |
| Window (RowNumber) | ~380µs | ~5ms | ~50ms |
| RollingSum(7) | ~150µs | ~1.7ms | ~17ms |

## Profiling

```bash
# CPU profile
go test -bench=BenchmarkFilter -cpuprofile=cpu.prof ./pkg/core
go tool pprof cpu.prof

# Memory profile
go test -bench=BenchmarkFilter -memprofile=mem.prof ./pkg/core
go tool pprof mem.prof

# Allocation tracking
go test -bench=. -benchmem ./pkg/core
```
