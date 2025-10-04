# Phase 1: API Audit for Go Idioms

**Date**: October 4, 2025
**Status**: In Progress
**Goal**: Ensure all public APIs follow Go best practices and idioms

---

## Audit Scope

### Packages Reviewed
1. ✅ `pkg/core` - DataFrame and Series core types
2. ⏳ `pkg/expr` - Expression engine
3. ⏳ `pkg/domain/dataframe` - Domain layer
4. ⏳ `pkg/interfaces` - Public interfaces
5. ⏳ `pkg/storage` - Storage backends

---

## Go API Design Principles

### 1. Naming Conventions
- ✅ Use MixedCaps (not underscores)
- ✅ Acronyms should be consistent (ID not Id, HTTP not Http)
- ✅ Package names should be lowercase, single word
- ✅ Exported names start with capital letter
- ⚠️ Avoid stutter (e.g., `dataframe.DataFrame` OK, `dataframe.DataFrameType` redundant)

### 2. Error Handling
- ✅ Return error as last return value
- ⚠️ Should wrap errors with context using `fmt.Errorf("context: %w", err)`
- 📝 Consider custom error types for common errors
- 📝 Should not panic in library code (except for programmer errors)

### 3. Context Usage
- ✅ Context should be first parameter
- ⚠️ Currently only used in storage operations
- 📝 Should add context to long-running operations

### 4. Interfaces
- ✅ Small, focused interfaces (1-3 methods)
- ✅ Accept interfaces, return structs
- 📝 Consider adding more interfaces for testing/mocking

### 5. Options Pattern
- ⚠️ Currently using struct options (ReadOptions, WriteOptions)
- 📝 Consider functional options for builder pattern
- 📝 Add validation for option combinations

---

## pkg/core API Audit

### DataFrame Constructor Functions

#### ✅ `NewDataFrame(record arrow.Record) *DataFrame`
**Status**: Good
**Strengths**:
- Clear, concise naming
- Returns concrete type (appropriate for data structure)
- Proper ownership transfer with `record.Retain()`

**Recommendations**: None

---

#### ✅ `NewDataFrameWithAllocator(record arrow.Record, allocator memory.Allocator) *DataFrame`
**Status**: Good
**Strengths**:
- Clear intent from name
- Appropriate for advanced use cases

**Recommendations**: None

---

#### ⚠️ `NewDataFrameFromStorage(ctx context.Context, backend storage.Backend, source string, opts storage.ReadOptions) (*DataFrame, error)`
**Status**: Good with minor improvements needed

**Strengths**:
- Context as first parameter ✅
- Clear error handling ✅
- Descriptive name ✅

**Issues**:
1. Error messages lack context:
   ```go
   // Current
   return nil, fmt.Errorf("failed to read from storage: %w", err)

   // Better
   return nil, fmt.Errorf("failed to read DataFrame from source %q: %w", source, err)
   ```

2. No validation of empty source:
   ```go
   // Should add
   if source == "" {
       return nil, fmt.Errorf("source cannot be empty")
   }
   ```

**Recommendations**:
- ✅ Add source validation
- ✅ Improve error messages with source context
- 📝 Consider adding timeout support via context

---

### DataFrame Query Methods

#### ✅ `Schema() *arrow.Schema`
**Status**: Perfect
- Returns pointer to read-only data
- No error needed (always succeeds)
- Clear, concise name

---

#### ✅ `NumRows() int64`
#### ✅ `NumCols() int64`
**Status**: Good

**Discussion**:
- Currently returns `int64` (Arrow native type)
- Consider: Should return `int` for Go idioms?
- Decision: Keep `int64` for Arrow consistency and large dataset support

---

#### ⚠️ `Column(name string) (*Series, error)`
**Status**: Good with optimization opportunity

**Current**:
```go
func (df *DataFrame) Column(name string) (*Series, error) {
    // Linear search through schema
    for i, field := range schema.Fields() {
        if field.Name == name {
            fieldIndex = i
            break
        }
    }
    if fieldIndex == -1 {
        return nil, fmt.Errorf("column not found: %s", name)
    }
    // ...
}
```

**Issues**:
1. Linear search - O(n) complexity
2. Error message could be more helpful

**Recommendations**:
- 📝 Add column name cache (map) for O(1) lookups
- ✅ Improve error message with suggestions for typos:
   ```go
   return nil, fmt.Errorf("column %q not found; available columns: %v", name, df.ColumnNames())
   ```

---

#### ✅ `ColumnAt(index int) (*Series, error)`
**Status**: Good

**Strengths**:
- Bounds checking ✅
- Clear error message ✅

**Minor Enhancement**:
```go
// Current
return nil, fmt.Errorf("column index %d out of bounds (0-%d)", index, numCols-1)

// Consider
return nil, fmt.Errorf("column index %d out of bounds [0, %d)", index, numCols)
```

---

#### ⚠️ `HasColumn(name string) bool`
**Status**: Good but inconsistent

**Issue**: Uses linear search (same as Column)

**Recommendation**: Share column lookup logic
```go
func (df *DataFrame) getColumnIndex(name string) int {
    // Could be optimized with map cache
}

func (df *DataFrame) HasColumn(name string) bool {
    return df.getColumnIndex(name) != -1
}

func (df *DataFrame) Column(name string) (*Series, error) {
    idx := df.getColumnIndex(name)
    // ...
}
```

---

### DataFrame Transformation Methods

#### ✅ `Select(columnNames []string) (*DataFrame, error)`
**Status**: Good

**Strengths**:
- Variadic would be more Go-like: `Select(names ...string)` ✅
- Current design is fine for explicitness

**Current**:
```go
func (df *DataFrame) Select(columnNames []string) (*DataFrame, error)
```

**Alternative** (more Go-like):
```go
func (df *DataFrame) Select(columnNames ...string) (*DataFrame, error)
```

**Decision**: Current design is acceptable, but variadic would be more idiomatic

---

#### ⚠️ `Filter(mask arrow.Array) (*DataFrame, error)`
**Status**: Good but could be more user-friendly

**Current**:
```go
func (df *DataFrame) Filter(mask arrow.Array) (*DataFrame, error)
```

**Issues**:
1. Requires users to work with Arrow arrays directly
2. Type checking happens at runtime

**Enhancement Ideas** (for future):
```go
// Keep current for performance
func (df *DataFrame) Filter(mask arrow.Array) (*DataFrame, error)

// Add convenience method
func (df *DataFrame) FilterExpr(expr expr.Expr) (*DataFrame, error) {
    mask, err := expr.Evaluate(df)
    if err != nil {
        return nil, fmt.Errorf("filter expression evaluation failed: %w", err)
    }
    return df.Filter(mask)
}
```

---

#### ✅ `WithColumn(name string, column arrow.Array) (*DataFrame, error)`
**Status**: Good

**Strengths**:
- Clear naming
- Handles both add and replace
- Good error messages

**Minor Enhancement**:
```go
// Current
return nil, fmt.Errorf("column length %d does not match DataFrame rows %d", column.Len(), df.NumRows())

// More helpful
return nil, fmt.Errorf("column %q length mismatch: got %d, expected %d rows", name, column.Len(), df.NumRows())
```

---

### DataFrame Join Methods

#### ✅ `InnerJoin(other *DataFrame, leftKey, rightKey string) (*DataFrame, error)`
#### ✅ `LeftJoin(other *DataFrame, leftKey, rightKey string) (*DataFrame, error)`
**Status**: Excellent

**Strengths**:
- Clear, SQL-like naming ✅
- Descriptive parameters ✅
- Good error messages ✅

**Nil Check**:
```go
if other == nil {
    return nil, fmt.Errorf("other DataFrame cannot be nil")
}
```
✅ Good defensive programming

---

#### ✅ `Join(other *DataFrame, leftKey, rightKey string, joinType JoinType) (*DataFrame, error)`
**Status**: Good

**Enum Pattern**:
```go
type JoinType int

const (
    InnerJoin JoinType = iota
    LeftJoin
)
```

**Recommendation**: Add String() method for debugging
```go
func (jt JoinType) String() string {
    switch jt {
    case InnerJoin:
        return "InnerJoin"
    case LeftJoin:
        return "LeftJoin"
    default:
        return fmt.Sprintf("JoinType(%d)", jt)
    }
}
```

---

### DataFrame Aggregation Methods

#### ✅ `GroupBy(columnNames ...string) *GroupedDataFrame`
**Status**: Good
- Variadic parameters ✅
- Returns dedicated type ✅

#### ⚠️ Agg methods need context
**Issue**: No context for cancellation in long-running aggregations

**Recommendation** (for future):
```go
func (g *GroupedDataFrame) AggContext(ctx context.Context, aggregations ...Aggregation) (*DataFrame, error)
```

---

### DataFrame Storage Methods

#### ✅ `WriteToStorage(ctx context.Context, backend storage.Backend, destination string, opts storage.WriteOptions) error`
**Status**: Excellent

**Strengths**:
- Context support ✅
- Clear naming ✅
- Options pattern ✅

**Minor Enhancement**:
```go
// Add validation
if destination == "" {
    return fmt.Errorf("destination cannot be empty")
}
```

---

### DataFrame Utility Methods

#### ✅ `Sort(columnName string, ascending bool) (*DataFrame, error)`
**Status**: Good but consider builder pattern

**Current**:
```go
df.Sort("age", false) // descending
```

**Alternative** (more readable):
```go
df.Sort("age", core.Descending)

// Or builder pattern
df.SortBy("age").Descending()
```

**Decision**: Current is acceptable, keep as-is

---

#### ✅ `SortMultiple(keys []SortKey) (*DataFrame, error)`
**Status**: Excellent

**Strengths**:
- Clear struct for multi-column sort ✅
- Good ergonomics ✅

```go
type SortKey struct {
    Column    string
    Ascending bool
}
```

---

#### ✅ `Equal(other *DataFrame) bool`
**Status**: Good

**Enhancement**: Document what "equal" means
```go
// Equal returns true if this DataFrame and other have:
// - Same number of rows and columns
// - Same schema (column names and types)
// - Same data values in same order
func (df *DataFrame) Equal(other *DataFrame) bool
```

---

#### ✅ `String() string`
**Status**: Good
Implements `fmt.Stringer` interface ✅

---

#### ✅ `Release()`
**Status**: Critical - Well Implemented

**Strengths**:
- Proper cleanup ✅
- Nil-safe ✅

```go
func (df *DataFrame) Release() {
    if df.record != nil {
        df.record.Release()
        df.record = nil
    }
}
```

**Recommendation**: Add note about usage in godoc
```go
// Release decrements the reference count of the underlying Arrow Record.
// The DataFrame should not be used after calling Release().
// It is safe to call Release multiple times.
```

---

## Series API Audit

### Constructor Functions

#### ✅ `NewSeries(array arrow.Array, field arrow.Field) *Series`
**Status**: Good
Clear construction pattern ✅

#### ⚠️ `NewSeriesFromData(data interface{}, name string) (*Series, error)`
**Status**: Needs improvement

**Issues**:
1. `interface{}` is not type-safe
2. Error handling for unsupported types unclear

**Recommendations**:
- 📝 Add generic version in Go 1.18+:
  ```go
  func NewSeriesFromSlice[T any](data []T, name string) (*Series, error)
  ```
- 📝 Document supported types clearly
- ✅ Add better error messages for unsupported types

---

### Getter Methods

#### ⚠️ `GetInt64(index int) (int64, error)`
#### ⚠️ `GetFloat64(index int) (float64, error)`
#### ⚠️ `GetString(index int) string`
#### ⚠️ `GetBool(index int) (bool, error)`

**Inconsistency**: GetString doesn't return error

**Current**:
```go
func (s *Series) GetString(index int) string {
    if s.array.IsNull(index) {
        return "" // Silent null handling
    }
    // ...
}
```

**Recommendation**: Consistent error handling
```go
func (s *Series) GetString(index int) (string, error) {
    if index < 0 || index >= s.array.Len() {
        return "", fmt.Errorf("index out of bounds")
    }
    if s.array.IsNull(index) {
        return "", fmt.Errorf("null value at index %d", index)
    }
    // ...
}
```

---

### Series Transformation Methods

#### ✅ `Slice(start, end int) *Series`
#### ✅ `Head(n int) *Series`
#### ✅ `Tail(n int) *Series`

**Status**: Good
Clear, Pandas-like API ✅

---

## Summary of Recommendations

### High Priority (Phase 1)

1. **Error Message Enhancement** ⚠️
   - Add context to all error messages
   - Include relevant values in errors
   - Suggest fixes where possible

2. **Consistency Fixes** ⚠️
   - Make GetString return error like other getters
   - Standardize error messages format
   - Consistent bounds checking

3. **Input Validation** ⚠️
   - Add empty string checks
   - Add nil pointer checks
   - Validate option combinations

4. **Documentation** 📝
   - Add godoc for all exported functions
   - Document thread-safety
   - Document memory ownership

### Medium Priority (Phase 2)

5. **Performance Optimizations** 📝
   - Column name lookup cache
   - Batch operations
   - Lazy evaluation

6. **API Enhancements** 📝
   - Variadic Select
   - Context support in aggregations
   - Builder patterns for complex operations

### Low Priority (Future)

7. **Generics Support** 📝
   - Type-safe data construction
   - Generic aggregations
   - Type-safe getters

8. **Convenience Methods** 📝
   - FilterExpr
   - More intuitive sort options
   - Chain-able transformations

---

## Action Items

### Immediate (This Week)
- [ ] Fix GetString to return error
- [ ] Add context to all error messages
- [ ] Add input validation for empty/nil
- [ ] Add JoinType.String() method
- [ ] Document thread-safety of all methods

### Short Term (Next 2 Weeks)
- [ ] Complete godoc for pkg/core
- [ ] Add column lookup optimization
- [ ] Create API usage examples
- [ ] Add benchmark tests for API patterns

### Long Term (Phase 2)
- [ ] Consider generics migration
- [ ] Add builder patterns
- [ ] Implement context cancellation

---

## Conclusion

The current API design is **solid and follows most Go idioms**. The main areas for improvement are:

1. **Error messages** - Add more context
2. **Consistency** - Fix GetString signature
3. **Validation** - Add defensive checks
4. **Documentation** - Complete godoc

Overall Assessment: **B+ → A- with recommended changes**

The API is production-ready with minor improvements needed for polish.
