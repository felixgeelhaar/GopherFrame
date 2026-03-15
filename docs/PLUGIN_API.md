# GopherFrame Plugin & Extension API

## Overview

GopherFrame is designed for extensibility through several plugin points:

1. **User-Defined Functions (UDFs)** — custom row/column transformations
2. **Custom Aggregations** — user-defined aggregation logic
3. **Expression Extensions** — new expression types via the `Expr` interface
4. **I/O Formats** — custom readers/writers
5. **Storage Backends** — alternative storage implementations

## 1. User-Defined Functions

### Scalar UDF

```go
// Row-by-row function
gf.ScalarUDF(
    []string{"col1", "col2"},           // Input columns
    arrow.PrimitiveTypes.Float64,        // Output type
    func(row map[string]interface{}) (interface{}, error) {
        a := row["col1"].(float64)
        b := row["col2"].(float64)
        return math.Sqrt(a*a + b*b), nil // Custom computation
    },
)
```

### Vectorized UDF (High Performance)

```go
// Operates on entire Arrow arrays — 10-100x faster than scalar
gf.VectorUDF(
    []string{"x"},
    arrow.PrimitiveTypes.Float64,
    func(cols map[string]arrow.Array) (arrow.Array, error) {
        xArr := cols["x"].(*array.Float64)
        pool := memory.NewGoAllocator()
        builder := array.NewFloat64Builder(pool)
        defer builder.Release()
        for i := 0; i < xArr.Len(); i++ {
            builder.Append(xArr.Value(i) * 2)
        }
        return builder.NewArray(), nil
    },
)
```

### Code-Generated UDF

```go
// Generate optimized, type-safe UDF code at build time
code, _ := gf.GenerateUDFCode("MyTransform",
    []string{"price", "qty"},
    arrow.PrimitiveTypes.Float64,
)
// Write to file for go:generate workflow
```

## 2. Custom Aggregations

```go
// Define any aggregation function operating on float64 slices
rangeFn := func(values []float64) float64 {
    min, max := values[0], values[0]
    for _, v := range values[1:] {
        if v < min { min = v }
        if v > max { max = v }
    }
    return max - min
}

df.GroupBy("category").Agg(
    gf.CustomAgg("price", "price_range", rangeFn),
)
```

## 3. Expression Interface

Implement the `expr.Expr` interface to create custom expression types:

```go
type Expr interface {
    Evaluate(df *core.DataFrame) (arrow.Array, error)
    Name() string
    String() string
    // + fluent methods (Add, Sub, Gt, Lt, etc.)
}
```

Use `expr.NewBinaryExpr`, `expr.NewUnaryExpr`, `expr.NewTernaryExpr` constructors
to compose custom expressions with existing ones.

## 4. Custom I/O Formats

Follow the pattern of existing I/O functions:

```go
func ReadMyFormat(filename string) (*gf.DataFrame, error) {
    // 1. Validate path
    // 2. Parse file into []map[string]interface{}
    // 3. Infer Arrow schema
    // 4. Build Arrow arrays
    // 5. Create arrow.Record
    // 6. Return gf.NewDataFrame(record)
}
```

## 5. Storage Backend

Implement the `storage.Backend` interface in `pkg/storage/`:

```go
type Backend interface {
    Store(record arrow.Record) error
    Load() (arrow.Record, error)
    Release()
}
```

## Best Practices

- Prefer VectorUDF over ScalarUDF for production performance
- Always validate inputs and handle nulls
- Release Arrow arrays when done (`defer arr.Release()`)
- Use `GenerateUDFCode` for frequently-used UDFs in hot paths
- Test custom expressions with the existing test patterns
