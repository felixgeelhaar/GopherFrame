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

func TestWriteReadAvro_Roundtrip(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "score", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"Alice", "Bob", "Carol"}, nil)
	na := nb.NewArray()
	defer na.Release()
	nb.Release()

	sb := array.NewFloat64Builder(pool)
	sb.AppendValues([]float64{95.5, 87.3, 92.1}, nil)
	sa := sb.NewArray()
	defer sa.Release()
	sb.Release()

	rec := array.NewRecord(schema, []arrow.Array{na, sa}, 3)
	defer rec.Release()

	df := NewDataFrame(rec)

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.avro")

	// Write
	err := WriteAvro(df, path)
	require.NoError(t, err)

	// Verify file exists and has content
	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Greater(t, info.Size(), int64(0))

	// Read back
	df2, err := ReadAvro(path)
	require.NoError(t, err)
	assert.Equal(t, int64(3), df2.NumRows())
	assert.True(t, df2.HasColumn("name"))
	assert.True(t, df2.HasColumn("score"))
}

func TestReadAvro_InvalidFile(t *testing.T) {
	_, err := ReadAvro("/nonexistent/test.avro")
	assert.Error(t, err)
}

func TestReadAvro_NotAvroFile(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bad.avro")
	err := os.WriteFile(path, []byte("not an avro file"), 0600)
	require.NoError(t, err)

	_, err = ReadAvro(path)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid magic")
}

func TestWriteAvro_InvalidPath(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
	}, nil)
	xb := array.NewFloat64Builder(pool)
	xb.Append(1.0)
	xa := xb.NewArray()
	defer xa.Release()
	xb.Release()

	rec := array.NewRecord(schema, []arrow.Array{xa}, 1)
	defer rec.Release()

	df := NewDataFrame(rec)
	err := WriteAvro(df, "/nonexistent/dir/test.avro")
	assert.Error(t, err)
}

func TestWriteReadAvro_IntColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		{Name: "label", Type: arrow.BinaryTypes.String},
	}, nil)

	ib := array.NewInt64Builder(pool)
	ib.AppendValues([]int64{1, 2, 3}, nil)
	ia := ib.NewArray()
	defer ia.Release()
	ib.Release()

	lb := array.NewStringBuilder(pool)
	lb.AppendValues([]string{"a", "b", "c"}, nil)
	la := lb.NewArray()
	defer la.Release()
	lb.Release()

	rec := array.NewRecord(schema, []arrow.Array{ia, la}, 3)
	defer rec.Release()

	df := NewDataFrame(rec)

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "int_test.avro")

	err := WriteAvro(df, path)
	require.NoError(t, err)

	df2, err := ReadAvro(path)
	require.NoError(t, err)
	assert.Equal(t, int64(3), df2.NumRows())
}

func TestWriteReadAvro_BoolColumn(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "active", Type: arrow.FixedWidthTypes.Boolean},
	}, nil)

	bb := array.NewBooleanBuilder(pool)
	bb.AppendValues([]bool{true, false, true}, nil)
	ba := bb.NewArray()
	defer ba.Release()
	bb.Release()

	rec := array.NewRecord(schema, []arrow.Array{ba}, 3)
	defer rec.Release()

	df := NewDataFrame(rec)

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bool_test.avro")

	err := WriteAvro(df, path)
	require.NoError(t, err)

	df2, err := ReadAvro(path)
	require.NoError(t, err)
	assert.Equal(t, int64(3), df2.NumRows())
}
