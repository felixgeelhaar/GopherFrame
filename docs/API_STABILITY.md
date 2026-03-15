# API Stability Guarantees

## Semantic Versioning

GopherFrame follows [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR** (v2.0.0): Breaking API changes
- **MINOR** (v1.x.0): New features, backward compatible
- **PATCH** (v1.x.y): Bug fixes, backward compatible

## Stability Promise

Starting with v1.0.0:

1. **No breaking changes until v2.0** — existing code will continue to compile and work correctly
2. **New features via minor versions** — added as backward-compatible extensions
3. **Bug fixes via patch versions** — no API changes

## Deprecation Policy

When an API needs to change:

1. The old API is marked with `// Deprecated:` godoc comment
2. The new API is introduced alongside the old one
3. The old API remains functional for **2 minor versions**
4. The old API is removed in the next major version

Example:
```go
// Deprecated: Use InnerJoinMulti instead. Will be removed in v2.0.
func (df *DataFrame) InnerJoin(other *DataFrame, leftKey, rightKey string) *DataFrame
```

## What is Covered

The following are part of the stable API:

- All exported types, functions, and methods in the `gopherframe` package
- All exported types in `pkg/expr` (the `Expr` interface and constructors)
- All I/O functions (`ReadParquet`, `WriteCSV`, `ReadJSON`, etc.)
- The `core.DataFrame` and `core.Series` types
- Error types and error messages (structure, not exact wording)

## What is NOT Covered

- Internal packages (`pkg/domain/`, `pkg/infrastructure/`, `pkg/storage/`)
- Benchmark performance numbers (may vary with optimizations)
- Debug output format (`String()` methods)
- Test utilities

## Go Version Support

- **Minimum**: Go 1.24
- **Tested**: Go 1.24, 1.25, 1.26
- **Policy**: Support the 3 most recent Go minor versions

## Compatibility Testing

The CI pipeline tests against all supported Go versions on both Linux and macOS to ensure cross-platform compatibility.
