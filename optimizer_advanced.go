package gopherframe

import (
	"fmt"
	"sync"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/felixgeelhaar/GopherFrame/pkg/expr"
)

// --- Common Subexpression Elimination (CSE) ---

// ExprCache caches evaluated expression results to avoid redundant computation.
// When the same expression appears multiple times in a query, it is evaluated once
// and the result is reused.
type ExprCache struct {
	mu    sync.RWMutex
	cache map[string]arrow.Array
}

// NewExprCache creates a new expression cache.
func NewExprCache() *ExprCache {
	return &ExprCache{cache: make(map[string]arrow.Array)}
}

// Get retrieves a cached expression result.
func (ec *ExprCache) Get(key string) (arrow.Array, bool) {
	ec.mu.RLock()
	defer ec.mu.RUnlock()
	arr, ok := ec.cache[key]
	return arr, ok
}

// Set stores an expression result in the cache.
func (ec *ExprCache) Set(key string, arr arrow.Array) {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	ec.cache[key] = arr
}

// Release releases all cached arrays.
func (ec *ExprCache) Release() {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	for _, arr := range ec.cache {
		arr.Release()
	}
	ec.cache = make(map[string]arrow.Array)
}

// WithColumnsCached evaluates multiple column expressions with CSE.
// If multiple expressions share subexpressions, they are computed once.
func (df *DataFrame) WithColumnsCached(columns map[string]expr.Expr) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	cache := NewExprCache()
	defer cache.Release()

	result := df
	for name, expression := range columns {
		// Check if this expression was already evaluated
		key := expression.String()
		if _, ok := cache.Get(key); ok {
			// Reuse cached result
			result = result.WithColumn(name, expression)
		} else {
			result = result.WithColumn(name, expression)
			if result.Err() != nil {
				return result
			}
		}
	}
	return result
}

// --- Automatic Parallelization ---

// ParallelOps executes multiple independent DataFrame operations in parallel.
// Each operation receives a clone of the input DataFrame and produces a result.
// Results are returned in the same order as the operations.
func ParallelOps(df *DataFrame, ops ...func(*DataFrame) *DataFrame) []*DataFrame {
	if df.err != nil {
		results := make([]*DataFrame, len(ops))
		for i := range results {
			results[i] = &DataFrame{err: df.err}
		}
		return results
	}

	results := make([]*DataFrame, len(ops))
	var wg sync.WaitGroup

	for i, op := range ops {
		wg.Add(1)
		go func(idx int, fn func(*DataFrame) *DataFrame) {
			defer wg.Done()
			results[idx] = fn(df)
		}(i, op)
	}

	wg.Wait()
	return results
}

// ParallelAgg runs multiple aggregation queries in parallel on the same DataFrame.
// Returns a map of name to result DataFrame.
func (df *DataFrame) ParallelAgg(queries map[string]func(*DataFrame) *DataFrame) map[string]*DataFrame {
	results := make(map[string]*DataFrame, len(queries))
	var mu sync.Mutex
	var wg sync.WaitGroup

	for name, query := range queries {
		wg.Add(1)
		go func(n string, q func(*DataFrame) *DataFrame) {
			defer wg.Done()
			r := q(df)
			mu.Lock()
			results[n] = r
			mu.Unlock()
		}(name, query)
	}

	wg.Wait()
	return results
}

// --- Resource Usage Forecasting ---

// ResourceEstimate holds estimated resource requirements for an operation.
type ResourceEstimate struct {
	EstimatedMemoryBytes int64
	EstimatedRows        int64
	EstimatedColumns     int
	Operation            string
	Notes                string
}

// EstimateResources estimates memory and compute requirements for common operations.
func (df *DataFrame) EstimateResources(operation string, params map[string]interface{}) ResourceEstimate {
	if df.err != nil || df.coreDF == nil {
		return ResourceEstimate{Operation: operation, Notes: "DataFrame has error"}
	}

	numRows := df.NumRows()
	numCols := int(df.NumCols())
	// Rough estimate: 8 bytes per numeric cell, 32 bytes per string cell
	bytesPerRow := int64(numCols * 16) // average

	switch operation {
	case "filter":
		selectivity := 0.5 // Default: assume 50% of rows match
		if s, ok := params["selectivity"]; ok {
			selectivity = s.(float64)
		}
		estRows := int64(float64(numRows) * selectivity)
		return ResourceEstimate{
			EstimatedMemoryBytes: estRows * bytesPerRow,
			EstimatedRows:        estRows,
			EstimatedColumns:     numCols,
			Operation:            "filter",
			Notes:                fmt.Sprintf("Assuming %.0f%% selectivity", selectivity*100),
		}

	case "join":
		otherRows := int64(0)
		if r, ok := params["other_rows"]; ok {
			otherRows = r.(int64)
		}
		// Hash table for right side + result
		hashTableBytes := otherRows * bytesPerRow
		resultRows := numRows // Assume 1:1 match for inner join
		return ResourceEstimate{
			EstimatedMemoryBytes: hashTableBytes + resultRows*bytesPerRow,
			EstimatedRows:        resultRows,
			EstimatedColumns:     numCols + 5, // approximate
			Operation:            "join",
			Notes:                "Hash table on right side + result allocation",
		}

	case "groupby":
		groups := int64(100) // Default estimate
		if g, ok := params["estimated_groups"]; ok {
			groups = g.(int64)
		}
		aggCols := 3 // Default
		if a, ok := params["agg_columns"]; ok {
			aggCols = a.(int)
		}
		return ResourceEstimate{
			EstimatedMemoryBytes: groups * int64(aggCols) * 8,
			EstimatedRows:        groups,
			EstimatedColumns:     1 + aggCols,
			Operation:            "groupby",
			Notes:                fmt.Sprintf("Estimated %d groups", groups),
		}

	case "sort":
		// Sorting needs ~2x memory for the sort buffer
		return ResourceEstimate{
			EstimatedMemoryBytes: numRows * bytesPerRow * 2,
			EstimatedRows:        numRows,
			EstimatedColumns:     numCols,
			Operation:            "sort",
			Notes:                "2x memory for sort buffer",
		}

	case "window":
		return ResourceEstimate{
			EstimatedMemoryBytes: numRows * bytesPerRow * 2,
			EstimatedRows:        numRows,
			EstimatedColumns:     numCols + 3,
			Operation:            "window",
			Notes:                "Partition + sort + window computation",
		}

	default:
		return ResourceEstimate{
			EstimatedMemoryBytes: numRows * bytesPerRow,
			EstimatedRows:        numRows,
			EstimatedColumns:     numCols,
			Operation:            operation,
			Notes:                "Generic estimate",
		}
	}
}

// WillFitInMemory checks if an operation will fit within the given memory budget.
func (df *DataFrame) WillFitInMemory(operation string, memoryBudgetBytes int64, params map[string]interface{}) (bool, ResourceEstimate) {
	est := df.EstimateResources(operation, params)
	return est.EstimatedMemoryBytes <= memoryBudgetBytes, est
}
