package core

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func TestLimitedAllocator_Basic(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(1024) // 1KB limit
	allocator := NewLimitedAllocator(base, limit)

	// Test basic allocation
	buf := allocator.Allocate(512)
	if buf == nil {
		t.Fatal("expected successful allocation of 512 bytes")
	}
	defer allocator.Free(buf)

	if allocator.AllocatedBytes() != 512 {
		t.Errorf("expected 512 bytes allocated, got %d", allocator.AllocatedBytes())
	}

	if allocator.UsagePercent() != 50.0 {
		t.Errorf("expected 50%% usage, got %.2f%%", allocator.UsagePercent())
	}
}

func TestLimitedAllocator_ExceedLimit(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(1024) // 1KB limit
	allocator := NewLimitedAllocator(base, limit)

	// Allocate up to limit
	buf1 := allocator.Allocate(512)
	if buf1 == nil {
		t.Fatal("expected successful allocation of 512 bytes")
	}
	defer allocator.Free(buf1)

	buf2 := allocator.Allocate(400)
	if buf2 == nil {
		t.Fatal("expected successful allocation of 400 bytes")
	}
	defer allocator.Free(buf2)

	// This should fail - would exceed limit (512 + 400 + 200 = 1112 > 1024)
	buf3 := allocator.Allocate(200)
	if buf3 != nil {
		allocator.Free(buf3)
		t.Fatal("expected allocation to fail when exceeding limit")
	}

	// Verify allocated amount hasn't changed
	if allocator.AllocatedBytes() != 912 {
		t.Errorf("expected 912 bytes allocated, got %d", allocator.AllocatedBytes())
	}
}

func TestLimitedAllocator_FreeMemory(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(1024)
	allocator := NewLimitedAllocator(base, limit)

	// Allocate and free
	buf := allocator.Allocate(512)
	if buf == nil {
		t.Fatal("expected successful allocation")
	}

	if allocator.AllocatedBytes() != 512 {
		t.Errorf("expected 512 bytes allocated, got %d", allocator.AllocatedBytes())
	}

	allocator.Free(buf)

	if allocator.AllocatedBytes() != 0 {
		t.Errorf("expected 0 bytes allocated after free, got %d", allocator.AllocatedBytes())
	}

	// Should be able to allocate again
	buf2 := allocator.Allocate(512)
	if buf2 == nil {
		t.Fatal("expected successful allocation after freeing memory")
	}
	defer allocator.Free(buf2)
}

func TestLimitedAllocator_Reallocate(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(2048)
	allocator := NewLimitedAllocator(base, limit)

	// Initial allocation
	buf := allocator.Allocate(512)
	if buf == nil {
		t.Fatal("expected successful allocation")
	}

	// Reallocate to larger size
	newBuf := allocator.Reallocate(1024, buf)
	if newBuf == nil {
		t.Fatal("expected successful reallocation")
	}
	defer allocator.Free(newBuf)

	if allocator.AllocatedBytes() != 1024 {
		t.Errorf("expected 1024 bytes allocated after reallocation, got %d", allocator.AllocatedBytes())
	}

	// Try to reallocate beyond limit
	tooBig := allocator.Reallocate(4096, newBuf)
	if tooBig != nil {
		allocator.Free(tooBig)
		t.Fatal("expected reallocation to fail when exceeding limit")
	}
}

func TestLimitedAllocator_MemoryPressureLevels(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(1000)

	tests := []struct {
		allocate int
		expected string
	}{
		{500, "low"},      // 50%
		{700, "medium"},   // 70%
		{850, "high"},     // 85%
		{970, "critical"}, // 97%
	}

	for _, tt := range tests {
		// Create fresh allocator for each test case
		allocator := NewLimitedAllocator(base, limit)

		buf := allocator.Allocate(tt.allocate)
		if buf == nil {
			t.Fatalf("expected successful allocation of %d bytes", tt.allocate)
		}
		defer allocator.Free(buf)

		level := allocator.MemoryPressureLevel()
		if level != tt.expected {
			t.Errorf("allocated %d bytes (%.0f%%): expected pressure level %q, got %q",
				tt.allocate, allocator.UsagePercent(), tt.expected, level)
		}
	}
}

func TestLimitedAllocator_CheckCanAllocate(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(1024)
	allocator := NewLimitedAllocator(base, limit)

	// Should succeed
	err := allocator.CheckCanAllocate(512)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	// Allocate some memory
	buf := allocator.Allocate(800)
	if buf == nil {
		t.Fatal("expected successful allocation")
	}
	defer allocator.Free(buf)

	// Should fail - would exceed limit
	err = allocator.CheckCanAllocate(300)
	if err == nil {
		t.Error("expected error when checking allocation that would exceed limit")
	}

	// Verify error type
	if _, ok := err.(*ErrMemoryLimitExceeded); !ok {
		t.Errorf("expected *ErrMemoryLimitExceeded, got %T", err)
	}
}

func TestLimitedAllocator_WithDataFrame(t *testing.T) {
	// Create allocator with 10MB limit
	base := memory.NewGoAllocator()
	limit := int64(10 * 1024 * 1024) // 10MB
	allocator := NewLimitedAllocator(base, limit)

	// Create a schema
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	// Build arrays
	idBuilder := array.NewInt64Builder(allocator)
	defer idBuilder.Release()
	valueBuilder := array.NewFloat64Builder(allocator)
	defer valueBuilder.Release()

	// Add 1000 rows
	for i := 0; i < 1000; i++ {
		idBuilder.Append(int64(i))
		valueBuilder.Append(float64(i) * 1.5)
	}

	idArray := idBuilder.NewArray()
	defer idArray.Release()
	valueArray := valueBuilder.NewArray()
	defer valueArray.Release()

	// Create record
	record := array.NewRecord(
		schema,
		[]arrow.Array{idArray, valueArray},
		1000,
	)
	defer record.Release()

	// Create DataFrame with limited allocator
	df := NewDataFrameWithAllocator(record, allocator)
	defer df.Release()

	// Verify DataFrame works
	if df.NumRows() != 1000 {
		t.Errorf("expected 1000 rows, got %d", df.NumRows())
	}

	// Check memory usage
	t.Logf("Memory usage: %d bytes (%.2f%% of %d byte limit)",
		allocator.AllocatedBytes(),
		allocator.UsagePercent(),
		allocator.Limit())

	// Memory pressure should be low for this small dataset
	if allocator.MemoryPressureLevel() == "critical" {
		t.Error("unexpected critical memory pressure for small dataset")
	}
}

func TestLimitedAllocator_OOMScenario(t *testing.T) {
	// Create allocator with very small limit to trigger OOM
	base := memory.NewGoAllocator()
	limit := int64(1024) // Only 1KB
	allocator := NewLimitedAllocator(base, limit)

	// Pre-flight check: verify we cannot allocate a large block
	largeSize := int64(10000) // 10KB
	err := allocator.CheckCanAllocate(largeSize)
	if err == nil {
		t.Error("expected error when checking allocation that exceeds limit")
	}

	// Verify error details
	if oomErr, ok := err.(*ErrMemoryLimitExceeded); ok {
		if oomErr.Requested != largeSize {
			t.Errorf("expected requested=%d, got %d", largeSize, oomErr.Requested)
		}
		if oomErr.Limit != limit {
			t.Errorf("expected limit=%d, got %d", limit, oomErr.Limit)
		}
		t.Logf("OOM error correctly reported: %v", oomErr)
	} else {
		t.Errorf("expected *ErrMemoryLimitExceeded, got %T", err)
	}

	// Allocate up to the limit successfully
	buf := allocator.Allocate(900)
	if buf == nil {
		t.Fatal("expected successful allocation within limit")
	}
	defer allocator.Free(buf)

	// Verify memory pressure is high
	level := allocator.MemoryPressureLevel()
	if level != "high" && level != "critical" {
		t.Errorf("expected high/critical memory pressure at 90%% usage, got %s", level)
	}

	// Try to allocate beyond limit - should fail gracefully
	buf2 := allocator.Allocate(200)
	if buf2 != nil {
		allocator.Free(buf2)
		t.Error("expected allocation to fail when exceeding limit")
	}

	t.Logf("Memory limit enforced successfully: %d/%d bytes used (%.1f%%)",
		allocator.AllocatedBytes(), allocator.Limit(), allocator.UsagePercent())
}

func TestLimitedAllocator_ConcurrentSafety(t *testing.T) {
	base := memory.NewGoAllocator()
	limit := int64(100 * 1024) // 100KB
	allocator := NewLimitedAllocator(base, limit)

	// Launch multiple goroutines allocating and freeing memory
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				buf := allocator.Allocate(128)
				if buf != nil {
					allocator.Free(buf)
				}
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// All memory should be freed
	if allocator.AllocatedBytes() != 0 {
		t.Errorf("expected 0 bytes allocated after concurrent operations, got %d",
			allocator.AllocatedBytes())
	}
}
