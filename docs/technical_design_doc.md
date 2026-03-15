# GopherFrame Technical Design Document

## Overview

GopherFrame is a production-grade DataFrame library for Go, built on Apache Arrow for zero-copy columnar data operations. This document describes the architecture, key design decisions, and implementation details.

## Architecture

### Layer Architecture

GopherFrame follows Clean Architecture with Domain-Driven Design:

```
┌─────────────────────────────────────────────┐
│              Public API Layer               │
│  (dataframe.go, groupby.go, io.go, etc.)   │
├─────────────────────────────────────────────┤
│           Application Services              │
│        (pkg/application/)                   │
├─────────────────────────────────────────────┤
│            Domain Layer                     │
│  (pkg/domain/dataframe, aggregation/)       │
├──────────────┬──────────────────────────────┤
│  Expression  │        Core Engine           │
│   Engine     │  (pkg/core/ - DataFrame,     │
│ (pkg/expr/)  │   Series, Window, Memory)    │
├──────────────┴──────────────────────────────┤
│        Storage Abstraction                  │
│   (pkg/storage/ - Backend interface)        │
├─────────────────────────────────────────────┤
│          Infrastructure                     │
│  (pkg/infrastructure/io, persistence/)      │
├─────────────────────────────────────────────┤
│           Apache Arrow Go                   │
│   (arrow.Record, arrow.Array, Compute)      │
└─────────────────────────────────────────────┘
```

### Data Flow

1. **Input**: Data enters via I/O layer (Parquet/CSV/Arrow IPC) or programmatic construction
2. **Storage**: Stored as `arrow.Record` in the Arrow backend
3. **Operations**: Transformations create new immutable records via expression evaluation
4. **Output**: Results written via I/O layer or accessed via Series/Record API

## Key Design Decisions

### ADR-001: Apache Arrow as Storage Backend

**Context**: Need a columnar in-memory format for high-performance data operations.

**Decision**: Use Apache Arrow Go v18 as the sole storage backend.

**Rationale**:
- Zero-copy column selection (O(1) instead of O(n))
- Cache-friendly columnar memory layout
- Native Parquet/IPC support without serialization overhead
- Interoperability with the broader Arrow ecosystem (Flight, Spark, Pandas)
- Production-proven at scale (used by Databricks, Snowflake, DuckDB)

**Consequences**:
- Heavy dependency on Arrow Go library
- Must work within Arrow's type system
- Memory management requires explicit `Release()` calls

### ADR-002: Immutable DataFrames

**Context**: Mutable state leads to bugs in concurrent environments and makes reasoning about data pipelines difficult.

**Decision**: All DataFrame and Series operations return new instances; originals are never modified.

**Rationale**:
- Thread-safe by default
- Enables method chaining without side effects
- Arrow Records are naturally immutable
- Simpler reasoning about data transformations

**Consequences**:
- Memory allocation for each operation (mitigated by Arrow's reference counting)
- Users must manage lifecycle with `Release()`

### ADR-003: Expression-Based Filtering

**Context**: Need a way to express complex predicates for filtering and column operations.

**Decision**: Build an expression AST (`pkg/expr/`) with lazy evaluation.

**Rationale**:
- Type-safe at compile time via Go interfaces
- Composable: `Col("a").Gt(Lit(5)).And(Col("b").Lt(Lit(10)))`
- Extensible for new operations (temporal, string, custom)
- Future optimization potential (predicate pushdown, constant folding)

**Consequences**:
- More complex than simple callback-based filtering
- Expression evaluation overhead (mitigated by operating on Arrow arrays, not rows)

### ADR-004: Hash-Based Joins

**Context**: Need efficient join implementations for production data sizes.

**Decision**: Use hash join as the default (and currently only) join strategy.

**Rationale**:
- O(n+m) time complexity for equi-joins
- Handles null values correctly
- Well-suited for in-memory operations
- Simpler implementation than sort-merge

**Consequences**:
- Memory overhead for hash table construction
- Not optimal for pre-sorted data (merge join would be O(n+m) without hash table)

### ADR-005: Domain-Driven Design

**Context**: Need clear separation of concerns as the library grows.

**Decision**: Organize code using DDD patterns with bounded contexts.

**Rationale**:
- Clear ownership boundaries (aggregation, expression, I/O)
- Testable in isolation
- Clean dependency direction (domain doesn't depend on infrastructure)

## Core Components

### DataFrame (`pkg/core/dataframe.go`)

The central data structure wrapping an `arrow.Record`:

- **Immutable**: All operations return new DataFrames
- **Error Accumulation**: Chainable operations accumulate errors checked via `Err()`
- **Resource Management**: Reference counting via Arrow's `Retain()`/`Release()`

### Expression Engine (`pkg/expr/expression.go`)

AST-based expression system:

- **Expr interface**: Common interface for all expressions
- **Column references**: `ColExpr` wraps column name lookups
- **Literals**: `LitExpr` wraps typed constant values
- **Binary operations**: Arithmetic (`Add`, `Sub`, `Mul`, `Div`), comparison (`Eq`, `Gt`, `Lt`), logical (`And`, `Or`)
- **Unary operations**: String (`Upper`, `Lower`, `Trim`), temporal (`Year`, `Month`), metrics (`Length`)

### Window Functions (`pkg/core/window.go`)

Flexible window function framework:

- **WindowBuilder**: Fluent API for `PartitionBy()`, `OrderBy()`, `Over()`
- **Partitioning**: Groups rows by column values before applying functions
- **Ordering**: Sorts within partitions for rank/lag/lead operations
- **Frame Types**: Row-based frames for rolling aggregations

### Memory Management (`pkg/core/memory_limit.go`)

Production-grade memory controls:

- **LimitedAllocator**: Wraps Arrow's allocator with configurable byte limits
- **Atomic operations**: Thread-safe allocation tracking
- **Pressure levels**: Low (<50%), Medium (50-75%), High (75-90%), Critical (>90%)
- **Pre-flight checks**: `CheckCanAllocate()` before expensive operations

## Type System

GopherFrame operates within Arrow's type system:

| Go Type | Arrow Type | Usage |
|---------|------------|-------|
| `int64` | `arrow.INT64` | Integer columns |
| `float64` | `arrow.FLOAT64` | Numeric columns |
| `string` | `arrow.STRING` | Text columns |
| `bool` | `arrow.BOOLEAN` | Boolean columns |
| `time.Time` | `arrow.TIMESTAMP` | Temporal columns |

Null values are handled via Arrow's validity bitmaps — every array tracks which indices are null without sentinel values.

## Performance Characteristics

### Complexity Analysis

| Operation | Time | Space |
|-----------|------|-------|
| Select | O(1) | O(1) — zero-copy pointer update |
| Filter | O(n) | O(k) — k matching rows |
| WithColumn | O(n) | O(n) — new column array |
| Sort | O(n log n) | O(n) |
| GroupBy+Agg | O(n) | O(g) — g groups |
| InnerJoin | O(n+m) | O(min(n,m)) — hash table |
| Window (RowNumber) | O(n log n) | O(n) — sort + scan |
| Rolling (window=w) | O(n) | O(1) — sliding window |

### Memory Model

- Arrow arrays use contiguous memory buffers
- Reference counting avoids copies for derived views
- `LimitedAllocator` prevents OOM in production
- Explicit `Release()` for deterministic cleanup

## Testing Strategy

- **Unit Tests**: Component isolation with table-driven tests
- **Integration Tests**: End-to-end workflow validation
- **Property-Based Tests**: Invariant verification via gopter
- **Benchmark Tests**: Performance regression detection in CI
- **Fuzz Tests**: Parser robustness for I/O operations
- **Memory Tests**: Leak detection and pressure testing

## Future Architecture Considerations

- **Query Optimization**: Expression tree rewriting (predicate pushdown, constant folding)
- **Streaming**: Iterator-based processing for datasets larger than memory
- **Partitioned I/O**: Hive-style partitioned Parquet reads with pruning
- **UDFs**: User-defined functions operating on Arrow arrays for extensibility
