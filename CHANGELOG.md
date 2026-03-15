# Changelog

All notable changes to GopherFrame will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

#### New Aggregations
- `Variance()` sample variance aggregation (Bessel's correction, N-1 denominator)
- `StdDev()` standard deviation aggregation
- `ConcatAgg(column, separator)` string concatenation aggregation (concat_ws)
- `CustomAgg(column, alias, fn)` user-defined aggregation functions

#### String Operations
- `Replace(old, new)` string replacement expression
- `PadLeft(length, char)` / `PadRight(length, char)` string padding expressions
- `SplitPart(separator, index)` split string and return Nth part

#### Join Operations
- `InnerJoinMulti()` / `LeftJoinMulti()` / `RightJoinMulti()` / `FullOuterJoinMulti()` multi-column join keys

#### User-Defined Functions
- `ScalarUDF(inputCols, outputType, fn)` row-by-row user-defined functions
- `VectorUDF(inputCols, outputType, fn)` vectorized UDFs on Arrow arrays

#### Data Transformation
- `Pivot(indexCols, pivotCol, valueCol)` long-to-wide transformation
- `Unpivot(idCols, valueCols, varName, valName)` wide-to-long (melt) transformation

#### I/O
- `ReadJSON()` / `WriteJSON()` JSON array-of-objects I/O
- `ReadNDJSON()` / `WriteNDJSON()` newline-delimited JSON I/O
- `ReadCSVChunked(filename, chunkSize)` streaming/chunked CSV reading
- `DataFrameIterator` with `ForEachChunk()` and `Collect()` methods

#### Temporal Utilities
- `DateRange(column, start, end, interval)` timestamp series generation
- `BusinessDaysBetween(start, end)` business day calculation
- `AddBusinessDays(start, days)` business day arithmetic

#### Data Quality & Validation
- `Describe()` descriptive statistics for all columns
- `DescribeString()` formatted statistics output
- `NullCount()` per-column null counts
- `IsComplete()` check for any nulls
- `Validate(rules...)` rule-based validation with `NotNull`, `Positive`, `InRange`, `UniqueValues`

#### Expression Optimization
- `Optimize(expr)` expression optimization pass
- `FoldedLit(value)` pre-computed constant expression

### Changed
- CI now tests Go 1.24, 1.25, 1.26 on both Ubuntu and macOS
- All GitHub Actions pinned to commit SHAs (supply chain hardening)
- Upgraded: actions/cache v3→v5, codecov v3→v5, golangci-lint v4→v9, setup-go v5→v6

### Security
- Resolved all IAC-013 (GitHub Action mutable tag) supply chain findings
- Upgraded deprecated GitHub Action versions (IAC-157)
- Added `SECURITY.md` with vulnerability reporting policy

### Documentation
- Added `CONTRIBUTING.md` contributor guide
- Added `CHANGELOG.md` (this file)
- Added `docs/technical_design_doc.md` architecture document with ADRs
- Added `docs/TROUBLESHOOTING.md` common issues guide
- Updated `ROADMAP.md` with all completed features

### Testing
- Added fuzz tests: `FuzzReadCSV`, `FuzzReadJSON`, `FuzzReadNDJSON`, `FuzzJSONRoundtrip`
- Test count increased from 279 to 386+

## [1.0.0] - 2025-10-04

### Added

#### Core Operations
- DataFrame and Series types with Apache Arrow backend
- `Select()`, `Filter()`, `WithColumn()` for data transformation
- `Sort()` and `SortMultiple()` for single and multi-column sorting
- `GroupBy()` with `Sum`, `Mean`, `Count`, `Min`, `Max` aggregations
- High-performance I/O: `ReadParquet`/`WriteParquet`, `ReadCSV`/`WriteCSV`, `ReadArrowIPC`/`WriteArrowIPC`
- Expression engine with type-safe column operations and predicates

#### Enhanced Join Operations (Phase 2.1)
- `InnerJoin()` with hash-based O(n+m) implementation
- `LeftJoin()` with proper null handling
- `RightJoin()` for right outer joins
- `FullOuterJoin()` with both-side null support
- `CrossJoin()` for Cartesian product operations

#### Window Functions (Phase 2.2)
- Analytical: `RowNumber()`, `Rank()`, `DenseRank()`, `Lag()`, `Lead()`
- Rolling: `RollingSum()`, `RollingMean()`, `RollingMin()`, `RollingMax()`, `RollingCount()`
- Cumulative: `CumSum()`, `CumMax()`, `CumMin()`, `CumProd()`
- Flexible API with `PartitionBy()`, `OrderBy()`, `Over()`

#### Temporal Operations (Phase 2.3)
- Component extraction: `Year()`, `Month()`, `Day()`, `Hour()`, `Minute()`, `Second()`
- Truncation: `TruncateToYear()`, `TruncateToMonth()`, `TruncateToDay()`, `TruncateToHour()`
- Arithmetic: `AddDays()`, `AddHours()`, `AddMinutes()`, `AddSeconds()`
- Full timezone support with Arrow Timestamp types

#### String Operations (Phase 2.4)
- Case conversion: `Upper()`, `Lower()`
- Whitespace: `Trim()`, `TrimLeft()`, `TrimRight()`
- Metrics: `Length()`
- Pattern matching: `Match()` with regex and pattern caching
- Predicates: `Contains()`, `StartsWith()`, `EndsWith()`

#### Statistical Aggregations (Phase 2.5)
- `Percentile()` with linear interpolation (0-100)
- `Median()` (50th percentile)
- `Mode()` with O(n) frequency counting
- `Correlation()` for Pearson correlation coefficient

#### Production Features
- `LimitedAllocator` with configurable memory limits
- Memory pressure monitoring (low/medium/high/critical)
- Path traversal protection for all I/O operations
- Property-based testing with gopter

### Performance
- 2-428x faster than Gota across all operations
- 2-200x less memory usage
- Zero-copy Apache Arrow operations throughout
- Hash-based joins with O(n+m) complexity

### Quality
- 279 tests passing, 95%+ coverage
- Zero gosec security vulnerabilities
- Automated benchmark regression CI
- Cross-version testing (Go 1.24+)

## [0.1.0] - 2024-09-15

### Added
- Initial release with core DataFrame operations
- Parquet, CSV, Arrow IPC I/O support
- Basic filtering, selection, and aggregation
- Expression engine with lazy evaluation
- Apache Arrow Go v18 backend
