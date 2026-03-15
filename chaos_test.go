package gopherframe

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestChaos_CorruptCSV tests handling of corrupt CSV files.
func TestChaos_CorruptCSV(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name    string
		content string
	}{
		{"empty_file", ""},
		{"header_only", "a,b,c\n"},
		{"mismatched_columns", "a,b\n1,2,3\n4\n"},
		{"binary_garbage", "\x00\x01\x02\x03\x04\x05"},
		{"huge_header", "col" + string(make([]byte, 10000))},
		{"null_bytes", "a,b\n\x00,\x00\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tmpDir, tt.name+".csv")
			err := os.WriteFile(path, []byte(tt.content), 0600)
			require.NoError(t, err)

			// Should not panic
			df, err := ReadCSV(path)
			_ = df
			_ = err
		})
	}
}

// TestChaos_CorruptJSON tests handling of corrupt JSON files.
func TestChaos_CorruptJSON(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name    string
		content string
	}{
		{"empty", ""},
		{"null", "null"},
		{"number", "42"},
		{"string", `"hello"`},
		{"nested_deep", `[{"a":{"b":{"c":{"d":1}}}}]`},
		{"huge_value", `[{"x":"` + string(make([]byte, 100000)) + `"}]`},
		{"unicode", `[{"name":"日本語テスト"}]`},
		{"special_chars", `[{"val":"line1\nline2\ttab"}]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tmpDir, tt.name+".json")
			err := os.WriteFile(path, []byte(tt.content), 0600)
			require.NoError(t, err)

			df, err := ReadJSON(path)
			_ = df
			_ = err
		})
	}
}

// TestChaos_NilDataFrame tests that operations on nil/errored DataFrames don't panic.
func TestChaos_NilDataFrame(t *testing.T) {
	df := &DataFrame{err: assert.AnError}

	// All of these should return errors, not panic
	assert.Error(t, df.Err())
	assert.Equal(t, int64(0), df.NumRows())
	assert.Equal(t, int64(0), df.NumCols())
	assert.Nil(t, df.ColumnNames())
	assert.False(t, df.HasColumn("x"))
	assert.False(t, df.IsComplete())
	assert.Nil(t, df.NullCount())

	// Operations should propagate error
	r := df.Filter(Col("x").Gt(Lit(0.0)))
	assert.Error(t, r.Err())

	r = df.Select("x")
	assert.Error(t, r.Err())

	r = df.Sort("x", true)
	assert.Error(t, r.Err())

	r = df.Pivot([]string{"a"}, "b", "c")
	assert.Error(t, r.Err())

	r = df.Unpivot([]string{"a"}, []string{"b"}, "var", "val")
	assert.Error(t, r.Err())

	r = df.MergeJoin(nil, "a", "b")
	assert.Error(t, r.Err())
}

// TestChaos_ConcurrentAccess tests concurrent read access to a DataFrame.
func TestChaos_ConcurrentAccess(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	for i := 0; i < 1000; i++ {
		xb.Append(float64(i))
	}
	xa := xb.NewArray()
	defer xa.Release()
	xb.Release()

	rec := array.NewRecord(schema, []arrow.Array{xa}, 1000)
	defer rec.Release()
	df := NewDataFrame(rec)

	// Launch 10 concurrent readers
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			_ = df.NumRows()
			_ = df.ColumnNames()
			_ = df.HasColumn("x")
			_ = df.NullCount()
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestChaos_EmptyDataFrame tests operations on empty DataFrames.
func TestChaos_EmptyDataFrame(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xa := xb.NewArray() // 0 rows
	defer xa.Release()
	xb.Release()

	rec := array.NewRecord(schema, []arrow.Array{xa}, 0)
	defer rec.Release()
	df := NewDataFrame(rec)

	assert.Equal(t, int64(0), df.NumRows())
	assert.True(t, df.IsComplete())

	// GroupBy on empty should work
	result := df.GroupBy("x").Agg(Sum("x").As("total"))
	// May error but should not panic
	_ = result

	// Describe empty
	stats, err := df.Describe()
	require.NoError(t, err)
	assert.Len(t, stats, 1)
	assert.Equal(t, int64(0), stats[0].Count)
}

// TestChaos_PathTraversal tests path traversal protection.
func TestChaos_PathTraversal(t *testing.T) {
	_, err := ReadCSV("../../etc/passwd")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "traversal")

	_, err = ReadJSON("../../../secret.json")
	assert.Error(t, err)

	_, err = ReadNDJSON("../../data/../../../etc/shadow")
	assert.Error(t, err)
}

// TestChaos_VeryLargeColumnName tests handling of extremely long column names.
func TestChaos_VeryLargeColumnName(t *testing.T) {
	pool := memory.NewGoAllocator()
	longName := string(make([]byte, 10000))
	for i := range longName {
		longName = longName[:i] + "x" + longName[i+1:]
	}

	schema := arrow.NewSchema([]arrow.Field{
		{Name: longName, Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.Append(1.0)
	xa := xb.NewArray()
	defer xa.Release()
	xb.Release()

	rec := array.NewRecord(schema, []arrow.Array{xa}, 1)
	defer rec.Release()
	df := NewDataFrame(rec)

	assert.True(t, df.HasColumn(longName))
}
