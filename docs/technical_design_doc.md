Technical Design Document: "GopherFrame"
Engineering a Production-First DataFrame Library for Go

Status:

Draft

Version:

2.0

Author:

Technical Leader

Last Updated:

July 22, 2024

Related PRD:

PRD: GopherFrame v2.0

Related Research:

Strategic Analysis

1. Overview & Goals
   This document provides the technical blueprint for building GopherFrame, a high-performance, Apache Arrow-backed DataFrame library for Go. It is the engineering response to the product vision and requirements outlined in PRD v2.0.

Our primary technical goals are derived directly from the PRD's success outcomes:

Performance: Achieve core operation speeds that are at least 10x faster than Gota and are competitive with Polars on multi-core hardware. This is our north star.

Zero-Copy Interoperability: Natively use the Apache Arrow memory format to enable frictionless, zero-copy data exchange with the Python ecosystem and other data systems, directly addressing the needs of our ML Engineer persona, Marcus.

Idiomatic & Type-Safe API: Create an API that feels natural to a Go developer. It must be explicit, composable, and leverage generics to provide compile-time type safety.

Production-Grade Reliability: Engineer a robust, thoroughly tested library capable of handling multi-gigabyte datasets in production pipelines, as required by our Data Engineer persona, Priya.

2. Domain-Driven Design (DDD)
   We will use DDD to model our library around the language and concepts of data engineering.

2.1. Ubiquitous Language
The terms defined here will be used consistently across all code, documentation, and team communication, aligning with the PRD.

DataFrame: The primary aggregate root. An immutable, in-memory representation of tabular data. It is conceptually a named collection of Series.

Series: An immutable, one-dimensional array of data representing a single column. All data in a Series shares a single, specific data type.

Schema: The metadata defining the structure of a DataFrame, consisting of a list of column names and their corresponding data types.

Expression: A declarative, lazy representation of a computation. Expressions build a syntax tree that describes an operation (e.g., (colA + colB) > 5) but do not execute it immediately.

Transformation: A high-level operation that takes one or more DataFrames and produces a new DataFrame (e.g., Filter, Select, GroupBy). Transformations are driven by Expressions.

Aggregation: A specific type of Expression that reduces a Series to a single value (e.g., Sum, Mean).

Kernel: A low-level, highly optimized compute function that operates directly on Arrow memory arrays. We will leverage Arrow's own compute kernels wherever possible.

2.2. Bounded Contexts & High-Level Architecture
The system is architected into three distinct bounded contexts, which will map to a clean package structure (pkg/core, pkg/expr, pkg/io). This separation of concerns is critical for maintainability and testability.

Diagram: A three-layer architecture.

Top Layer (Public API): gfr.DataFrame, gfr.Series, gfr.Col(), gfr.Lit(). This is the user-facing surface.

Middle Layer (Core Engine): core.DataFrame, core.Series. Manages transformations and executes expression plans.

Bottom Layer (Subsystems):

expr: Defines the expression tree structures.

io: Handles Parquet, CSV, and IPC readers/writers.

arrow: The foundational Apache Arrow library (compute kernels, memory management).

Dependencies flow strictly downwards. The public API orchestrates calls to the core engine, which in turn leverages the subsystems to perform its work.

3. Detailed Design (v0.1 MVP)
   This section details the implementation plan for the features defined in the PRD.

3.1. Core Data Structures (pkg/core)
Our core structures are thin, immutable wrappers around the powerful primitives provided by the official Apache Arrow Go library. This is a non-negotiable principle.

// pkg/core/dataframe.go

import "github.com/apache/arrow/go/v15/arrow"

// DataFrame is the internal, immutable representation of tabular data.
// It is fundamentally a wrapper around an arrow.Record, ensuring direct
// compatibility with the Arrow ecosystem.
type DataFrame struct {
// This is never exposed directly. Immutability is enforced by returning
// new DataFrame instances from all transformation methods.
record arrow.Record
}

// All public methods will be on a separate, user-facing struct in the
// root package, which holds this core.DataFrame.

3.2. Expressions Engine & Lazy Execution (pkg/expr)
To achieve high performance, we will implement a lazy execution engine. User operations build an Expression tree, which is then optimized and executed in a single pass.

// pkg/expr/expression.go

// Expr is the interface that all expression types (column references,
// literals, binary operations, function calls) must implement.
type Expr interface {
// ToArrow evaluates the expression against a DataFrame. It returns the
// resulting column as an arrow.Array. This is the entry point for the
// execution planner.
ToArrow(df core.DataFrame) (arrow.Array, error)

    // Name returns the output name of this expression, e.g., "revenue" or
    // "SUM(cost)".
    Name() string

}

// --- Example Expression Structs ---

// ColExpr represents a reference to an existing column.
type ColExpr struct { name string }

// LitExpr represents a literal value (e.g., a number or string).
type LitExpr[T any] struct { value T }

// BinaryExpr represents an operation like addition or equality.
type BinaryExpr struct {
Left, Right Expr
Op string // "add", "eq", "gt", etc.
}

// FunctionExpr represents a function call like SUM() or LOWER().
type FunctionExpr struct {
Name string
Args []Expr
}

Execution Flow Example: df.Filter(gfr.Col("A").Gt(gfr.Lit(100)))

The user's call constructs a BinaryExpr AST: BinaryExpr{Left: ColExpr{"A"}, Right: LitExpr{100}, Op: "gt"}.

The Filter method receives this Expr. It does not execute it immediately.

Instead, it passes the Expr to an internal execution planner.

The planner calls expr.ToArrow(df). This recursively traverses the tree, fetching Col("A") and resolving Lit(100).

It then dispatches the operation to Arrow's highly-optimized gt compute kernel.

The result is a boolean arrow.Array (the filter mask).

This mask is then used with Arrow's filter kernel to produce the final, filtered arrow.Record.

A new DataFrame is constructed from this record and returned.

3.3. I/O Subsystem (pkg/io)
The I/O subsystem must be highly concurrent to meet Priya's user story of reading multi-gigabyte files from object storage.

Readers (parquet, csv, ipc): Will be designed to accept an io.Reader. We will implement our own logic to parallelize reads from sources that also implement io.ReaderAt (like local files or S3 objects), reading chunks of the file concurrently.

Writers: Will accept an io.Writer. For partitioned Parquet output, the writer will manage the creation of multiple sub-directories and files.

Concurrency: We will use a pool of goroutines for decoding file chunks (e.g., row groups in Parquet) in parallel to maximize CPU utilization on multi-core machines.

3.4. Public API (gopherframe.go)
The public API will be clean, chainable, and focused on developer ergonomics.

// gopherframe.go (root package)

// DataFrame is the public-facing, user-friendly DataFrame object.
type DataFrame struct {
// The internal, immutable core DataFrame.
coreDF core.DataFrame
// An optional error field to allow for chained calls to fail gracefully.
err error
}

// Filter returns a new DataFrame containing only the rows matching the predicate.
// It follows the lazy execution model.
func (df DataFrame) Filter(predicate expr.Expr) DataFrame {
// ... implementation details ...
}

// GroupBy initiates a grouped operation.
func (df DataFrame) GroupBy(cols ...string) \*GroupedDataFrame {
// ... implementation details ...
}

// GroupedDataFrame represents a pending group-by operation.
type GroupedDataFrame struct {
df DataFrame
by []string
err error
}

// Aggregate performs the specified aggregations. This is an "action" that
// triggers the execution of the group-by plan.
func (gdf \*GroupedDataFrame) Aggregate(aggs ...expr.Expr) DataFrame {
// This will trigger a call to Arrow's powerful group-by kernel, which
// can perform multiple aggregations in a single pass over the data.
// ... implementation details ...
}

4. Test-Driven Development (TDD) Strategy
   Our development process will be strictly test-driven to ensure we meet our reliability and performance goals. We will not merge a feature until it has comprehensive test coverage.

Unit Tests (\_test.go):

Scope: Individual components in isolation (e.g., expression struct creation, a single I/O parser function).

Goal: Verify logical correctness of the smallest code units.

Integration Tests:

Scope: End-to-end user workflows as defined in the PRD.

Example Test Case: "TestPriyaETLWorkflow"

Create a sample multi-gigabyte Parquet file in a temporary location.

Use the library to read the file.

Perform a chain of transformations: Select, Filter, WithColumn.

Perform a GroupBy and Aggregate.

Write the result to a new Parquet file.

Read the output file back and assert that its contents are exactly as expected.

Goal: Verify that the system's components work together correctly to solve a real user problem.

Benchmark Tests (\_test.go with BenchmarkXxx):

Scope: Performance measurement of all critical paths.

Targets: BenchmarkReadParquet_1GB, BenchmarkFilter_100M_Rows, BenchmarkGroupBy_Complex, BenchmarkIPC_Roundtrip.

Goal: Continuously track our performance against the PRD's goals (10x Gota, competitive with Polars). These will be part of our CI pipeline to prevent performance regressions.

Property-Based Testing:

Tool: github.com/leanovate/gopter

Example Property: "For any DataFrame df and any sort expression s, the number of rows in df.Sort(s) is always equal to the number of rows in df."

Goal: Uncover edge cases and validate the logical invariants of our transformations that example-based testing might miss.

5. Roadmap Beyond MVP (v0.2+)
   The architecture is explicitly designed to be extensible.

Joins (v0.2): The PRD defers joins. Our Expression-based architecture is well-suited for this. A df.Join(otherDF) method will create a JoinPlan object. The execution engine will then choose the optimal physical join strategy (e.g., hash join, merge join) based on data size and sort order, and dispatch to the appropriate Arrow compute kernel.

User-Defined Functions (UDFs): We will explore a udf.Apply() expression that takes a Go function. The initial implementation will operate on chunks of data (arrow.Array) for safety. For higher performance, we will investigate a go:generate-based approach to create specialized, reflection-free UDF kernels, as inspired by the research document.

6. Risks & Mitigations
   Risk: Performance Parity with Polars is Extremely Difficult. Polars is a mature, highly-optimized Rust library. Matching it is a significant challenge.

Mitigation: We will not reinvent the wheel. We will lean heavily on the performance of the underlying Apache Arrow Go library and its C++ kernels (via CGo if necessary for specific, critical-path functions where the Go implementation lags). Our primary focus will be on building an efficient and intelligent execution planner on top of Arrow.

Risk: Arrow Go Library Complexity. The Arrow library is powerful but has a steep learning curve and its own set of bugs or performance quirks.

Mitigation: We will build a strong internal competency in the Arrow library. We will contribute upstream to the Arrow project to fix bugs or add features we need, benefiting the entire ecosystem. Our comprehensive integration tests will act as a safety net to detect issues in our usage of the library.

Risk: API Design Churn. Getting the "Go-idiomatic" feel right is subjective and can lead to breaking changes.

Mitigation: We will release early v0.x versions and actively solicit feedback from our target personas (Priya, Marcus, Alex) and the broader Go community. We will be clear that APIs are unstable until v1.0.0 but will strive to get the core concepts right from the start.
