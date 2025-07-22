# Property-Based Testing in GopherFrame

## Overview

GopherFrame uses property-based testing with the `gopter` library to ensure correctness across a wide range of inputs. Property tests complement unit tests by automatically generating test cases and finding edge cases.

## Test Categories

### 1. Schema Preservation Tests
- **FilterPreservesSchema**: Verifies that filter operations maintain the DataFrame's schema
- **SelectSubset**: Ensures Select returns exactly the requested columns
- **WithColumnIncreases**: Confirms WithColumn adds exactly one column

### 2. Operation Consistency Tests
- **GroupByAggregationConsistency**: Validates aggregation results match manual calculations
- **FilterCommutativity**: Checks that filter order doesn't affect results (for valid ranges)
- **ChainedOperationsAssociativity**: Ensures operation chains produce consistent results

### 3. Data Integrity Tests
- **ParquetRoundTrip**: Verifies data survives Parquet write/read cycles
- **DataFrameImmutability**: Confirms operations don't modify original DataFrames
- **CSVStringHandling**: Tests special character handling in CSV I/O

### 4. Edge Case Tests
- **NullHandling**: Validates null value processing across operations
- **EmptyDataFrame**: Ensures operations handle empty DataFrames gracefully
- **LargeValues**: Tests extreme floating-point values (Inf, NaN, Max/Min)

## Running Property Tests

```bash
# Run all property tests
go test -run TestProperty -v

# Run specific property test
go test -run TestPropertyFilterPreservesSchema -v

# Run with custom seed (for reproducibility)
GOPTER_SEED=1234567890 go test -run TestProperty -v
```

## Property Test Structure

```go
properties.Property("Property name", prop.ForAll(
    func(input Type) bool {
        // Setup
        df := createDataFrame(input)
        defer df.Release()
        
        // Operation
        result := df.SomeOperation()
        defer result.Release()
        
        // Assertion
        return checkProperty(result)
    },
    generator, // Input generator
))
```

## Key Findings

Property testing has uncovered several important edge cases:

1. **Empty DataFrames**: Initial implementation didn't handle empty records correctly
2. **Filter Commutativity**: Filters with equal or near-equal thresholds need special handling
3. **Null Propagation**: Null values must be handled consistently across all operations
4. **Floating Point Edge Cases**: Operations should handle Inf, -Inf, and NaN gracefully

## Generator Patterns

### Basic Generators
```go
// Generate test rows with controlled data
gen.SliceOf(genTestRow())

// Generate specific value ranges
gen.Float64Range(0, 100)
gen.Int64Range(1, 1000)
gen.AlphaString()
```

### Custom Generators
```go
// Generate rows with potential nulls
gen.PtrOf(gen.Float64Range(0, 100))
gen.OneGenOf(gen.PtrOf(gen.AlphaString()), gen.Const((*string)(nil)))

// Generate extreme values
gen.OneConstOf(0.0, math.Inf(1), math.Inf(-1), math.NaN())
```

## Best Practices

1. **Start Simple**: Begin with basic properties, then add complexity
2. **Shrinking**: Gopter automatically shrinks failing cases to minimal examples
3. **Determinism**: Use seeds for reproducible test failures
4. **Performance**: Limit data size for I/O-intensive tests
5. **Coverage**: Combine with traditional unit tests for comprehensive coverage

## Future Improvements

1. **Concurrent Operations**: Test thread safety with parallel property tests
2. **Performance Properties**: Ensure operations complete within time bounds
3. **Memory Properties**: Verify no memory leaks across operations
4. **Custom Generators**: Domain-specific data generators for realistic test data

## Integration with CI/CD

Property tests should be part of the continuous integration pipeline:

```yaml
# Example GitHub Actions configuration
- name: Run Property Tests
  run: |
    go test -run TestProperty -v -count=3
    go test -run TestProperty -v -race
```

The `-count=3` ensures tests run multiple times with different seeds, increasing the chance of finding edge cases.