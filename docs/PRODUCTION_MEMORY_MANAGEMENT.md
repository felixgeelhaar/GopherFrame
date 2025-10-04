# Production Memory Management

GopherFrame provides production-grade memory management through the `LimitedAllocator` to prevent Out-Of-Memory (OOM) errors in production deployments.

## Overview

The `LimitedAllocator` wraps Apache Arrow's memory allocator with configurable hard limits, memory pressure monitoring, and graceful OOM handling. This is essential for:

- **Cloud deployments** with strict memory constraints
- **Multi-tenant services** requiring resource quotas
- **High-availability systems** that must prevent OOM crashes
- **Cost optimization** by enforcing memory budgets

## Quick Start

```go
import (
    "github.com/apache/arrow-go/v18/arrow/memory"
    "github.com/felixgeelhaar/GopherFrame/pkg/core"
    gf "github.com/felixgeelhaar/GopherFrame"
)

// Configure 512MB memory limit
base := memory.NewGoAllocator()
limitedAllocator := core.NewLimitedAllocator(base, 512*1024*1024)

// Create DataFrame with limited allocator
df := gf.NewDataFrameWithAllocator(record, limitedAllocator)
defer df.Release()
```

## Key Features

### 1. Hard Memory Limits

Prevents allocations that would exceed the configured limit:

```go
allocator := core.NewLimitedAllocator(base, 100*1024*1024) // 100MB limit

// Pre-flight check before expensive operations
if err := allocator.CheckCanAllocate(estimatedSize); err != nil {
    if oomErr, ok := err.(*core.ErrMemoryLimitExceeded); ok {
        log.Printf("Insufficient memory: %v", oomErr)
        // Implement fallback strategy
    }
}
```

### 2. Memory Pressure Monitoring

Track memory usage in real-time:

```go
// Check current usage
usedBytes := allocator.AllocatedBytes()
usagePercent := allocator.UsagePercent()
pressureLevel := allocator.MemoryPressureLevel() // "low", "medium", "high", "critical"

// React to memory pressure
switch pressureLevel {
case "critical":
    return fmt.Errorf("memory pressure critical, aborting operation")
case "high":
    log.Warn("high memory pressure, reducing batch size")
    batchSize = batchSize / 2
}
```

### 3. Graceful OOM Handling

Handle memory exhaustion without crashing:

```go
// Allocate returns nil when limit would be exceeded
buf := allocator.Allocate(requestedSize)
if buf == nil {
    // Handle OOM gracefully
    log.Error("allocation failed: memory limit exceeded")
    // Options:
    // 1. Process in smaller chunks
    // 2. Trigger garbage collection
    // 3. Return error to client
    // 4. Scale horizontally
}
```

## Memory Pressure Levels

| Level | Usage | Action |
|-------|-------|--------|
| **low** | < 60% | Normal operation |
| **medium** | 60-80% | Monitor closely, consider reducing batch size |
| **high** | 80-95% | Use memory-efficient algorithms, trigger GC |
| **critical** | > 95% | Abort non-essential operations, scale up |

## Production Patterns

### Pattern 1: Pre-flight Checks

Always verify memory availability before expensive operations:

```go
func ProcessLargeDataset(df *gf.DataFrame, allocator *core.LimitedAllocator) error {
    // Estimate memory needed
    estimatedMemory := int64(df.NumRows()) * int64(df.NumCols()) * 8

    // Pre-flight check
    if err := allocator.CheckCanAllocate(estimatedMemory); err != nil {
        return fmt.Errorf("insufficient memory: %w", err)
    }

    // Safe to proceed
    result := df.GroupBy("category").Agg(...)
    defer result.Release()
    return nil
}
```

### Pattern 2: Chunked Processing

Process large datasets in memory-bounded chunks:

```go
func ProcessInChunks(totalRows int64, allocator *core.LimitedAllocator) error {
    chunkSize := int64(10000)

    for offset := int64(0); offset < totalRows; offset += chunkSize {
        // Check memory before processing chunk
        if allocator.UsagePercent() > 80 {
            // Trigger cleanup
            runtime.GC()

            // If still high, reduce chunk size
            if allocator.UsagePercent() > 80 {
                chunkSize = chunkSize / 2
            }
        }

        // Process chunk
        chunk := LoadChunk(offset, chunkSize)
        defer chunk.Release()

        // ... process chunk ...
    }
    return nil
}
```

### Pattern 3: Environment-Based Configuration

Configure limits from environment:

```go
func GetMemoryLimitFromEnv() int64 {
    limitMB := 512 // Default: 512MB
    if envLimit := os.Getenv("GOPHERFRAME_MEMORY_LIMIT_MB"); envLimit != "" {
        fmt.Sscanf(envLimit, "%d", &limitMB)
    }
    return int64(limitMB * 1024 * 1024)
}

// In main():
memLimit := GetMemoryLimitFromEnv()
allocator := core.NewLimitedAllocator(memory.NewGoAllocator(), memLimit)
```

### Pattern 4: Monitoring Integration

Export metrics to monitoring systems:

```go
// Prometheus example
memoryUsage := promauto.NewGauge(prometheus.GaugeOpts{
    Name: "gopherframe_memory_bytes",
    Help: "Current memory usage in bytes",
})

memoryPressure := promauto.NewGauge(prometheus.GaugeOpts{
    Name: "gopherframe_memory_pressure_percent",
    Help: "Memory usage as percentage of limit",
})

// Update metrics periodically
go func() {
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        memoryUsage.Set(float64(allocator.AllocatedBytes()))
        memoryPressure.Set(allocator.UsagePercent())
    }
}()
```

## Error Handling

### ErrMemoryLimitExceeded

Detailed error information when limits are exceeded:

```go
err := allocator.CheckCanAllocate(size)
if oomErr, ok := err.(*core.ErrMemoryLimitExceeded); ok {
    log.Printf("Memory limit exceeded:")
    log.Printf("  Requested: %d bytes", oomErr.Requested)
    log.Printf("  Limit: %d bytes", oomErr.Limit)
    log.Printf("  Current: %d bytes (%.1f%% used)",
        oomErr.Current,
        float64(oomErr.Current)/float64(oomErr.Limit)*100)

    // Calculate available memory
    available := oomErr.Limit - oomErr.Current
    log.Printf("  Available: %d bytes", available)
}
```

## Thread Safety

`LimitedAllocator` is thread-safe and uses atomic operations for all memory tracking:

```go
// Safe to use from multiple goroutines
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        buf := allocator.Allocate(1024)
        if buf != nil {
            defer allocator.Free(buf)
            // ... use buffer ...
        }
    }()
}
wg.Wait()
```

## Best Practices

1. **Always use LimitedAllocator in production**
   - Set limits based on container/pod memory limits
   - Leave 20-30% headroom for Go runtime and OS

2. **Perform pre-flight checks**
   - Estimate memory before expensive operations
   - Use `CheckCanAllocate()` before allocations

3. **Monitor memory pressure**
   - Check `MemoryPressureLevel()` regularly
   - Implement alerts for high/critical pressure

4. **Implement graceful degradation**
   - Reduce batch sizes when memory is constrained
   - Switch to streaming algorithms under pressure

5. **Clean up resources promptly**
   - Always defer `Release()` calls
   - Don't hold DataFrames longer than necessary

6. **Test OOM scenarios**
   - Test with realistic memory limits
   - Verify graceful handling of OOM conditions

## Example Application

See [cmd/examples/production_memory/main.go](../cmd/examples/production_memory/main.go) for a complete production example demonstrating:

- Memory limit configuration
- Pre-flight checks
- Memory pressure monitoring
- Graceful OOM handling
- Chunked processing
- Environment-based configuration

## Related Documentation

- [Memory Management Guide](MEMORY_MANAGEMENT.md) - General memory concepts
- [Performance Tuning](PERFORMANCE_TUNING.md) - Memory optimization strategies
- [Production Deployment](PRODUCTION_DEPLOYMENT.md) - Production best practices
