# Troubleshooting Guide

Common issues and solutions when using GopherFrame.

## Memory Issues

### "memory limit exceeded" error

**Cause**: The `LimitedAllocator` has reached its configured memory cap.

**Solution**:
```go
// Increase memory limit
limited := core.NewLimitedAllocator(pool, 2*1024*1024*1024) // 2GB

// Or check pressure before operations
if limited.MemoryPressureLevel() == "high" {
    // Process in smaller batches
}
```

### OOM (Out of Memory) crashes

**Cause**: Processing datasets larger than available RAM without memory limits.

**Solution**:
1. Use `LimitedAllocator` with appropriate limits
2. Use `ReadCSVChunked()` for streaming processing
3. Always call `defer df.Release()` to free memory promptly
4. Process data in batches with `DataFrameIterator`

```go
it, _ := gf.ReadCSVChunked("large.csv", 10000)
it.ForEachChunk(func(chunk *gf.DataFrame) error {
    // Process chunk
    return nil
})
```

### Memory leaks

**Cause**: Not calling `Release()` on DataFrames or intermediate results.

**Solution**:
```go
df, _ := gf.ReadParquet("data.parquet")
defer df.Release()

result := df.Filter(df.Col("x").Gt(gf.Lit(0.0)))
defer result.Release() // Don't forget intermediate results!
```

## Type Errors

### "aggregation only supports float64"

**Cause**: Using numeric aggregations (Sum, Mean, Variance, etc.) on non-float64 columns.

**Solution**: Ensure your numeric columns are `float64` when creating DataFrames:
```go
builder := array.NewFloat64Builder(pool) // Use Float64, not Int64
```

### "unsupported data type for operation"

**Cause**: Applying an operation to an incompatible column type (e.g., `Upper()` on a numeric column).

**Solution**: Check column types before operations:
```go
schema := df.Schema()
for _, field := range schema.Fields() {
    fmt.Printf("%s: %s\n", field.Name, field.Type)
}
```

## Join Issues

### Unexpected null values after join

**Cause**: Key columns have null values, which never match in joins (SQL semantics).

**Solution**: Filter nulls before joining:
```go
clean := df.Filter(df.Col("key").Eq(df.Col("key"))) // Removes null keys
```

### Wrong number of result rows

**Cause**: Duplicate keys in either table create multiple matches (Cartesian product per key).

**Solution**: Verify key uniqueness:
```go
result := df.Validate(gf.UniqueValues("key"))
if !result.Valid {
    // Handle duplicates
}
```

## I/O Issues

### "directory traversal detected"

**Cause**: File path contains `..` components, which GopherFrame blocks for security.

**Solution**: Use absolute paths or paths without `..`:
```go
df, err := gf.ReadParquet("/data/sales.parquet") // Absolute path
```

### Parquet read/write errors

**Cause**: File corruption, incompatible schema, or missing file.

**Solution**:
1. Verify file exists: `os.Stat(filename)`
2. Check file permissions
3. Verify Parquet schema compatibility

### JSON parse errors

**Cause**: Malformed JSON or inconsistent schemas across records.

**Solution**:
```go
// Ensure valid JSON array of objects
// [{"col1": val1, "col2": val2}, ...]

// For NDJSON, each line must be a valid JSON object
// {"col1": val1}\n{"col2": val2}\n
```

## Expression Issues

### "column not found" in expressions

**Cause**: Column name doesn't exist in the DataFrame.

**Solution**: Check available columns first:
```go
fmt.Println(df.ColumnNames())
if df.HasColumn("my_col") {
    // Safe to use
}
```

### UDF returns wrong type

**Cause**: The UDF function returns a different type than the declared `outputType`.

**Solution**: Ensure your UDF returns the exact type specified:
```go
// If outputType is arrow.PrimitiveTypes.Float64
// The function MUST return float64, not int or int64
ScalarUDF(cols, arrow.PrimitiveTypes.Float64, func(row map[string]interface{}) (interface{}, error) {
    return float64(42), nil // Explicit float64
})
```

## Performance Issues

### Slow GroupBy operations

**Solution**:
1. Reduce the number of group columns
2. Use `CustomAgg` for complex aggregations instead of multiple passes
3. Filter data before grouping

### Slow joins

**Solution**:
1. Join on smaller tables when possible
2. Filter data before joining
3. Use the smaller table as the right side (hash table is built from right)

### Slow iteration

**Solution**: Avoid row-by-row access. Use vectorized operations:
```go
// Slow: row-by-row
for i := 0; i < int(df.NumRows()); i++ { ... }

// Fast: vectorized
result := df.WithColumn("doubled", df.Col("x").Mul(gf.Lit(2.0)))
```

## Getting Help

- **Issues**: [GitHub Issues](https://github.com/felixgeelhaar/GopherFrame/issues)
- **Discussions**: [GitHub Discussions](https://github.com/felixgeelhaar/GopherFrame/discussions)
- **API Reference**: [pkg.go.dev](https://pkg.go.dev/github.com/felixgeelhaar/GopherFrame)
