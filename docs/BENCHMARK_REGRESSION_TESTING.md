# Benchmark Regression Testing

## Overview

GopherFrame uses automated benchmark regression testing to ensure performance remains stable or improves with each change. Two complementary CI workflows protect against performance degradation:

1. **Benchmark Regression Check** (`benchmark-regression.yml`) - Runs on every PR
2. **Performance Tracking** (`performance.yml`) - Continuous monitoring on main branch

## How It Works

### Pull Request Workflow

When you open a PR, the benchmark regression workflow automatically:

1. **Runs benchmarks on your PR branch** (5 iterations, 3 seconds each)
2. **Checks out the base branch** (usually `main`)
3. **Runs the same benchmarks on base branch**
4. **Compares results using `benchstat`** (Go's official statistical comparison tool)
5. **Posts a comment** on the PR with detailed results
6. **Fails the CI check** if regressions > 10% are detected

### Main Branch Workflow

After merging to main:

1. Benchmarks run on every commit
2. Results are tracked over time in `dev/bench`
3. Alerts triggered if performance degrades >10%
4. GitHub Pages shows historical performance trends

## Understanding Benchmark Results

### PR Comment Format

```
## 📊 Benchmark Comparison Results

This PR's performance compared to `main`:

name                    old time/op    new time/op    delta
DataFrameCreation_1K    24.3µs ± 2%    23.1µs ± 3%   -4.94%  (p=0.008 n=5+5)
Filter_10K              404µs ± 1%     380µs ± 2%    -5.94%  (p=0.008 n=5+5)
Select_10K              772ns ± 5%     801ns ± 4%    +3.76%  (p=0.016 n=5+5)
```

### How to Read Results

#### Delta Column
- **`~0%`** - No statistically significant change ✅
- **`-X%`** - Your PR is **faster** by X% 🎉
- **`+X%`** - Your PR is **slower** by X% ⚠️

#### Statistical Significance
- **`(p=0.008 n=5+5)`** - Statistical test results
  - `p < 0.05` - Change is statistically significant
  - `p >= 0.05` - Change may be noise (marked with `~`)
  - `n=5+5` - Number of samples compared

### Regression Thresholds

| Change | Action Required |
|--------|----------------|
| **< 5%** | ✅ No action needed (within normal variance) |
| **5-10%** | ⚠️  Review recommended - investigate if intentional |
| **> 10%** | ❌ CI fails - must fix or justify |

## Running Benchmarks Locally

### Quick Start

```bash
# Run all benchmarks
go test -bench=. -benchmem ./pkg/core

# Run specific benchmark
go test -bench=BenchmarkSelect -benchmem ./pkg/core

# Run with more iterations for accuracy
go test -bench=BenchmarkSelect -benchmem -count=10 ./pkg/core
```

### Comparing with Base Branch

```bash
# Benchmark base branch
git checkout main
go test -bench=. -benchtime=3s -count=5 ./pkg/core > base.txt

# Benchmark your changes
git checkout your-branch
go test -bench=. -benchtime=3s -count=5 ./pkg/core > pr.txt

# Compare with benchstat
go install golang.org/x/perf/cmd/benchstat@latest
benchstat base.txt pr.txt
```

### Example Comparison Session

```bash
$ benchstat base.txt pr.txt
name                    old time/op    new time/op    delta
DataFrameCreation_1K    24.3µs ± 2%    23.1µs ± 3%   -4.94%  (p=0.008 n=5+5)
Filter_10K              404µs ± 1%     380µs ± 2%    -5.94%  (p=0.008 n=5+5)
Select_10K              772ns ± 5%     801ns ± 4%    +3.76%  (p=0.016 n=5+5)

name                    old alloc/op   new alloc/op   delta
DataFrameCreation_1K    53.3kB ± 0%    51.2kB ± 0%   -3.94%  (p=0.008 n=5+5)
Filter_10K              798kB ± 0%     780kB ± 0%    -2.26%  (p=0.008 n=5+5)
Select_10K              1.62kB ± 0%    1.62kB ± 0%     ~     (all equal)
```

## Responding to Regressions

### Step 1: Understand the Regression

1. **Check the PR comment** for affected benchmarks
2. **Run benchmarks locally** to reproduce
3. **Profile the code** if regression is unexpected

```bash
# CPU profile
go test -bench=BenchmarkYourFunction -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profile
go test -bench=BenchmarkYourFunction -memprofile=mem.prof
go tool pprof mem.prof
```

### Step 2: Investigate Root Cause

Common causes of regressions:

- **Additional allocations** - Check `alloc/op` and `allocs/op`
- **Algorithmic changes** - O(n) → O(n²) complexity increase
- **Unnecessary copies** - Lost zero-copy optimization
- **Type conversions** - Added reflection or type assertions
- **Lock contention** - New synchronization overhead

### Step 3: Fix or Justify

#### Option A: Fix the Regression

```go
// Before (slow)
func Filter(df *DataFrame, condition Expr) *DataFrame {
    // Creates unnecessary copy
    rows := df.ToRows()  // ⚠️ Slow
    filtered := filterRows(rows)
    return FromRows(filtered)
}

// After (fast)
func Filter(df *DataFrame, condition Expr) *DataFrame {
    // Zero-copy using Arrow predicates
    predicateArray, _ := condition.Evaluate(df)  // ✅ Fast
    defer predicateArray.Release()
    return df.coreDF.Filter(predicateArray)
}
```

#### Option B: Justify the Regression

If regression is intentional (e.g., trading performance for correctness):

1. Document the reason in PR description
2. Add a comment explaining the tradeoff
3. Update relevant documentation
4. Get approval from maintainers

Example justification:
```
## Performance Regression Justification

The 15% regression in `Filter` is intentional and acceptable because:

1. **Correctness**: Previous implementation had incorrect null handling
2. **Safety**: Now properly validates all edge cases
3. **Impact**: Absolute time is still < 1ms for 10K rows
4. **Alternative**: Would require significant refactoring (tracked in #123)

Verified the tradeoff is acceptable for our use cases.
```

## Writing Performance-Conscious Code

### Best Practices

#### 1. Avoid Unnecessary Allocations

```go
// ❌ Bad: Creates new slice every time
func GetValues(df *DataFrame) []interface{} {
    values := make([]interface{}, df.NumRows())
    for i := 0; i < df.NumRows(); i++ {
        values[i] = df.GetValue(i)
    }
    return values
}

// ✅ Good: Reuses Arrow arrays (zero-copy)
func GetValues(df *DataFrame) arrow.Array {
    return df.Column(0).Array()  // Returns reference
}
```

#### 2. Use Arrow's Zero-Copy Operations

```go
// ❌ Bad: Copies data
func SelectColumns(df *DataFrame, cols []string) *DataFrame {
    newData := make([][]interface{}, len(cols))
    for i, col := range cols {
        newData[i] = copyColumn(df, col)  // Expensive copy
    }
    return NewDataFrame(newData)
}

// ✅ Good: Zero-copy column selection
func SelectColumns(df *DataFrame, cols []string) *DataFrame {
    indices := df.getColumnIndices(cols)
    return df.selectByIndices(indices)  // Just updates pointers
}
```

#### 3. Batch Operations

```go
// ❌ Bad: Row-by-row processing
for i := 0; i < df.NumRows(); i++ {
    df.SetValue(i, col, transform(df.GetValue(i, col)))
}

// ✅ Good: Vectorized processing
array := df.Column(col).Array()
transformed := computeKernel.Transform(array)
df.SetColumn(col, transformed)
```

### Performance Testing Checklist

Before submitting a PR:

- [ ] Run `go test -bench=. -benchmem ./...` locally
- [ ] Compare with base branch using `benchstat`
- [ ] Profile code if new hot paths added
- [ ] Check for unnecessary allocations (`allocs/op`)
- [ ] Verify O(n) complexity for data operations
- [ ] Test with realistic data sizes (1K, 10K, 100K rows)

## Benchmark Writing Guidelines

### Good Benchmark Structure

```go
func BenchmarkMyOperation(b *testing.B) {
    // Setup (not measured)
    pool := memory.NewGoAllocator()
    df := createTestDataFrame(pool, 1000)
    defer df.Release()

    // Reset timer to exclude setup
    b.ResetTimer()

    // Benchmark loop
    for i := 0; i < b.N; i++ {
        result := df.MyOperation()
        result.Release()  // Clean up each iteration
    }
}
```

### Benchmark Naming Convention

```go
BenchmarkOperationName_Size
BenchmarkFilter_1K        // Filter with 1,000 rows
BenchmarkFilter_10K       // Filter with 10,000 rows
BenchmarkFilter_100K      // Filter with 100,000 rows
```

### Testing Multiple Sizes

```go
func BenchmarkFilter(b *testing.B) {
    sizes := []int{1000, 10000, 100000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%dRows", size), func(b *testing.B) {
            pool := memory.NewGoAllocator()
            df := createTestDataFrame(pool, size)
            defer df.Release()

            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                result := df.Filter(condition)
                result.Release()
            }
        })
    }
}
```

## Continuous Monitoring

### GitHub Pages Dashboard

Performance trends are published to GitHub Pages at:
`https://felixgeelhaar.github.io/GopherFrame/dev/bench/`

Tracks:
- Operation times over commits
- Memory allocations trends
- Performance alerts
- Comparison charts

### Daily Performance Tests

Scheduled workflow runs daily at 2 AM UTC:
- Full benchmark suite
- Memory profiling
- Leak detection
- Historical comparison

## Troubleshooting

### Benchmark Failed to Run

**Symptom**: CI shows "Benchmark execution failed"

**Solutions**:
1. Check if benchmark code compiles: `go test -c ./pkg/core`
2. Verify test filters: `go test -bench=BenchmarkName -run=^$ ./pkg/core`
3. Look for panics in benchmark setup code

### Inconsistent Results

**Symptom**: Results vary widely between runs

**Solutions**:
1. Increase benchmark time: `-benchtime=10s`
2. Increase sample count: `-count=10`
3. Check for background processes affecting performance
4. Run on clean CI environment (GitHub Actions)

### Benchstat Shows No Data

**Symptom**: `benchstat base.txt pr.txt` shows empty output

**Solutions**:
1. Verify both files contain benchmark results
2. Check benchmark names match between files
3. Ensure output format is correct (don't capture `stderr`)

## References

- [Go Benchmarking Documentation](https://pkg.go.dev/testing#hdr-Benchmarks)
- [benchstat User Guide](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)
- [Performance Testing Best Practices](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [GopherFrame Performance Benchmarks](../BENCHMARKS.md)
- [Gota Comparison Benchmarks](GOTA_COMPARISON_BENCHMARKS.md)
