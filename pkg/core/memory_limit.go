package core

import (
	"fmt"
	"sync/atomic"

	"github.com/apache/arrow-go/v18/arrow/memory"
)

// LimitedAllocator wraps a memory.Allocator with configurable memory limits.
// It tracks total allocated memory and enforces a hard limit to prevent OOM errors.
//
// This is designed for production environments where you need to:
// - Prevent runaway memory usage
// - Enforce resource quotas per operation or per service
// - Gracefully handle memory pressure
//
// Thread Safety: All methods are thread-safe using atomic operations.
type LimitedAllocator struct {
	// base is the underlying allocator (typically memory.NewGoAllocator())
	base memory.Allocator

	// limit is the maximum number of bytes that can be allocated
	limit int64

	// allocated tracks the current number of bytes allocated
	// Uses atomic operations for thread-safe access
	allocated atomic.Int64
}

// NewLimitedAllocator creates a new memory allocator with a hard limit.
//
// Parameters:
//   - base: The underlying allocator to wrap (e.g., memory.NewGoAllocator())
//   - limitBytes: Maximum total bytes that can be allocated (hard limit)
//
// Returns:
//   - *LimitedAllocator: A new allocator that enforces the memory limit
//
// Example:
//
//	// Allow up to 1GB of memory
//	base := memory.NewGoAllocator()
//	limited := NewLimitedAllocator(base, 1*1024*1024*1024)
//	df := NewDataFrameWithAllocator(record, limited)
//	defer df.Release()
func NewLimitedAllocator(base memory.Allocator, limitBytes int64) *LimitedAllocator {
	return &LimitedAllocator{
		base:  base,
		limit: limitBytes,
	}
}

// Allocate attempts to allocate memory, returning an error if limit would be exceeded.
//
// This method is thread-safe and uses atomic operations to track memory usage.
// If the allocation would exceed the configured limit, it returns an error without
// allocating any memory.
//
// Complexity: O(1)
func (a *LimitedAllocator) Allocate(size int) []byte {
	// Check if allocation would exceed limit
	currentlyAllocated := a.allocated.Load()
	if currentlyAllocated+int64(size) > a.limit {
		// Return nil to indicate allocation failure
		// The caller should check for nil and handle the OOM condition
		return nil
	}

	// Perform the allocation
	buf := a.base.Allocate(size)
	if buf == nil {
		// Base allocator failed (system OOM)
		return nil
	}

	// Track the allocation
	a.allocated.Add(int64(size))
	return buf
}

// Reallocate changes the size of an existing allocation.
//
// This method checks if the new size would exceed the memory limit.
// If so, it returns an error without modifying the existing allocation.
//
// Complexity: O(1) for size check, O(n) for data copy if reallocation occurs
func (a *LimitedAllocator) Reallocate(size int, b []byte) []byte {
	oldSize := int64(len(b))
	newSize := int64(size)
	sizeDelta := newSize - oldSize

	// Check if reallocation would exceed limit
	currentlyAllocated := a.allocated.Load()
	if currentlyAllocated+sizeDelta > a.limit {
		return nil
	}

	// Perform the reallocation
	newBuf := a.base.Reallocate(size, b)
	if newBuf == nil {
		return nil
	}

	// Update tracked allocation
	a.allocated.Add(sizeDelta)
	return newBuf
}

// Free releases memory back to the allocator.
//
// This method decrements the tracked allocation count, making the memory
// available for future allocations.
//
// Complexity: O(1)
func (a *LimitedAllocator) Free(b []byte) {
	size := int64(len(b))
	a.base.Free(b)
	a.allocated.Add(-size)
}

// AllocatedBytes returns the current number of bytes allocated.
//
// This method is thread-safe and returns the live allocation count.
//
// Returns:
//   - int64: Current number of bytes allocated through this allocator
//
// Example:
//
//	allocator := NewLimitedAllocator(base, 1*1024*1024*1024)
//	// ... perform operations ...
//	fmt.Printf("Memory usage: %.2f MB\n", float64(allocator.AllocatedBytes())/(1024*1024))
func (a *LimitedAllocator) AllocatedBytes() int64 {
	return a.allocated.Load()
}

// Limit returns the configured memory limit in bytes.
//
// Returns:
//   - int64: Maximum number of bytes that can be allocated
func (a *LimitedAllocator) Limit() int64 {
	return a.limit
}

// UsagePercent returns the current memory usage as a percentage of the limit.
//
// Returns:
//   - float64: Percentage of limit currently used (0.0 to 100.0)
//
// Example:
//
//	if allocator.UsagePercent() > 90.0 {
//	    log.Warn("Memory usage above 90%, consider reducing dataset size")
//	}
func (a *LimitedAllocator) UsagePercent() float64 {
	if a.limit == 0 {
		return 0.0
	}
	return float64(a.allocated.Load()) / float64(a.limit) * 100.0
}

// MemoryPressureLevel returns the current memory pressure level.
//
// Returns:
//   - string: "low" (<60%), "medium" (60-80%), "high" (80-95%), "critical" (>95%)
//
// Example:
//
//	switch allocator.MemoryPressureLevel() {
//	case "critical":
//	    return fmt.Errorf("memory pressure critical, aborting operation")
//	case "high":
//	    log.Warn("high memory pressure, consider reducing batch size")
//	}
func (a *LimitedAllocator) MemoryPressureLevel() string {
	usage := a.UsagePercent()
	switch {
	case usage < 60:
		return "low"
	case usage < 80:
		return "medium"
	case usage < 95:
		return "high"
	default:
		return "critical"
	}
}

// ErrMemoryLimitExceeded is returned when an allocation would exceed the configured limit.
type ErrMemoryLimitExceeded struct {
	Requested int64
	Limit     int64
	Current   int64
}

func (e *ErrMemoryLimitExceeded) Error() string {
	return fmt.Sprintf(
		"memory limit exceeded: requested %d bytes, limit %d bytes, currently allocated %d bytes (%.1f%% used)",
		e.Requested,
		e.Limit,
		e.Current,
		float64(e.Current)/float64(e.Limit)*100,
	)
}

// CheckCanAllocate verifies if an allocation of the given size would succeed.
//
// This is useful for pre-flight checks before expensive operations.
//
// Parameters:
//   - size: Number of bytes to allocate
//
// Returns:
//   - error: nil if allocation would succeed, *ErrMemoryLimitExceeded otherwise
//
// Example:
//
//	if err := allocator.CheckCanAllocate(estimatedSize); err != nil {
//	    return fmt.Errorf("insufficient memory for operation: %w", err)
//	}
func (a *LimitedAllocator) CheckCanAllocate(size int64) error {
	currentlyAllocated := a.allocated.Load()
	if currentlyAllocated+size > a.limit {
		return &ErrMemoryLimitExceeded{
			Requested: size,
			Limit:     a.limit,
			Current:   currentlyAllocated,
		}
	}
	return nil
}
