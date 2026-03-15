# Contributing to GopherFrame

Thank you for your interest in contributing to GopherFrame! This document provides guidelines and information for contributors.

## Getting Started

### Prerequisites

- Go 1.24 or later
- Git
- `golangci-lint` (for linting)
- `gosec` (for security scanning)

### Setup

```bash
git clone https://github.com/felixgeelhaar/GopherFrame.git
cd GopherFrame
go mod download
go test ./...
```

## Development Workflow

### Branch Strategy

- `main` — stable, release-ready code
- Feature branches — `feature/<name>` for new features
- Bug fix branches — `fix/<name>` for bug fixes

### Test-Driven Development (TDD)

GopherFrame follows TDD with the Red-Green-Refactor cycle:

1. **Red**: Write a failing test that defines the expected behavior
2. **Green**: Write the minimum code to make the test pass
3. **Refactor**: Improve the code while keeping tests green

### Running Tests

```bash
# All tests
go test ./...

# With race detection
go test -race ./...

# With coverage
go test -coverprofile=coverage.out ./pkg/...
go tool cover -html=coverage.out

# Benchmarks
go test -bench=. -benchmem ./pkg/core

# Property-based tests
go test -run TestProperty -count=3
```

### Code Quality

```bash
# Linting
golangci-lint run

# Security scanning
gosec -exclude=G304,G104 ./...
```

## Submitting Changes

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add Variance aggregation function
fix: handle null values in LeftJoin
docs: update API reference for window functions
test: add edge case tests for CrossJoin
perf: optimize hash join for large datasets
refactor: extract common join logic into helper
```

### Pull Request Process

1. Create a feature branch from `main`
2. Write tests first (TDD)
3. Implement the feature
4. Ensure all tests pass: `go test -race ./...`
5. Ensure linting passes: `golangci-lint run`
6. Run benchmarks to check for regressions: `go test -bench=. -benchmem ./pkg/core`
7. Update documentation if needed
8. Submit a PR with a clear description

### PR Checklist

- [ ] Tests written and passing
- [ ] No performance regressions (benchmarks run)
- [ ] Code follows Go conventions and project patterns
- [ ] Documentation updated (if applicable)
- [ ] No new security vulnerabilities (`gosec` clean)
- [ ] Coverage maintained at 80%+ for `pkg/`

## Architecture

GopherFrame follows Clean Architecture with Domain-Driven Design:

```
pkg/
├── core/           # Core DataFrame/Series implementations
├── expr/           # Expression engine and AST
├── domain/         # Domain models and business logic
├── application/    # Application services
├── infrastructure/ # I/O and persistence
├── interfaces/     # Public API facade
└── storage/        # Pluggable storage backends
```

### Key Principles

- **Immutability**: All data structures are immutable
- **Zero-Copy**: Leverage Apache Arrow's zero-copy capabilities
- **Type Safety**: Use Go generics and strong typing
- **Production-First**: Every feature must handle production data sizes

### Adding New Features

When adding a new feature (e.g., a new aggregation function):

1. Define the domain logic in `pkg/domain/`
2. Implement the core operation in `pkg/core/`
3. Add expression support in `pkg/expr/` if needed
4. Expose via the public API in root-level files
5. Write comprehensive tests including edge cases
6. Add benchmarks for performance-critical paths
7. Update examples if the feature is user-facing

## Code of Conduct

- Be respectful and constructive in all interactions
- Focus feedback on code, not the person
- Welcome contributors of all experience levels
- Follow Go community standards

## Questions?

- **Issues**: [GitHub Issues](https://github.com/felixgeelhaar/GopherFrame/issues)
- **Discussions**: [GitHub Discussions](https://github.com/felixgeelhaar/GopherFrame/discussions)
