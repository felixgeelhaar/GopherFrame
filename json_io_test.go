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

func TestReadWriteJSON_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "name", Type: arrow.BinaryTypes.String},
		{Name: "score", Type: arrow.PrimitiveTypes.Float64},
	}, nil)

	nb := array.NewStringBuilder(pool)
	nb.AppendValues([]string{"Alice", "Bob"}, nil)
	nArr := nb.NewArray()
	defer nArr.Release()
	nb.Release()

	sb := array.NewFloat64Builder(pool)
	sb.AppendValues([]float64{95.5, 87.3}, nil)
	sArr := sb.NewArray()
	defer sArr.Release()
	sb.Release()

	record := array.NewRecord(schema, []arrow.Array{nArr, sArr}, 2)
	defer record.Release()

	df := NewDataFrame(record)

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.json")

	err := WriteJSON(df, path)
	require.NoError(t, err)

	// Verify file exists
	_, err = os.Stat(path)
	require.NoError(t, err)

	// Read back
	df2, err := ReadJSON(path)
	require.NoError(t, err)
	assert.Equal(t, int64(2), df2.NumRows())
	assert.True(t, df2.HasColumn("name"))
	assert.True(t, df2.HasColumn("score"))
}

func TestReadWriteNDJSON_Basic(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "id", Type: arrow.PrimitiveTypes.Float64},
		{Name: "city", Type: arrow.BinaryTypes.String},
	}, nil)

	ib := array.NewFloat64Builder(pool)
	ib.AppendValues([]float64{1, 2, 3}, nil)
	iArr := ib.NewArray()
	defer iArr.Release()
	ib.Release()

	cb := array.NewStringBuilder(pool)
	cb.AppendValues([]string{"NYC", "LA", "CHI"}, nil)
	cArr := cb.NewArray()
	defer cArr.Release()
	cb.Release()

	record := array.NewRecord(schema, []arrow.Array{iArr, cArr}, 3)
	defer record.Release()

	df := NewDataFrame(record)

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.ndjson")

	err := WriteNDJSON(df, path)
	require.NoError(t, err)

	df2, err := ReadNDJSON(path)
	require.NoError(t, err)
	assert.Equal(t, int64(3), df2.NumRows())
	assert.True(t, df2.HasColumn("id"))
	assert.True(t, df2.HasColumn("city"))
}

func TestReadJSON_InvalidPath(t *testing.T) {
	_, err := ReadJSON("/nonexistent/path/test.json")
	assert.Error(t, err)
}

func TestWriteJSON_InvalidPath(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "a", Type: arrow.BinaryTypes.String},
	}, nil)
	ab := array.NewStringBuilder(pool)
	ab.Append("x")
	aArr := ab.NewArray()
	defer aArr.Release()
	ab.Release()

	record := array.NewRecord(schema, []arrow.Array{aArr}, 1)
	defer record.Release()

	df := NewDataFrame(record)
	err := WriteJSON(df, "/nonexistent/dir/test.json")
	assert.Error(t, err)
}

func TestReadJSON_MalformedJSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bad.json")
	err := os.WriteFile(path, []byte(`{not valid json`), 0600)
	require.NoError(t, err)

	_, err = ReadJSON(path)
	assert.Error(t, err)
}

func TestReadNDJSON_MalformedLine(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bad.ndjson")
	err := os.WriteFile(path, []byte("{\"a\":1}\n{bad json\n"), 0600)
	require.NoError(t, err)

	_, err = ReadNDJSON(path)
	assert.Error(t, err)
}

func TestReadJSON_BooleanValues(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "bool.json")
	data := `[{"name": "Alice", "active": true}, {"name": "Bob", "active": false}]`
	err := os.WriteFile(path, []byte(data), 0600)
	require.NoError(t, err)

	df, err := ReadJSON(path)
	require.NoError(t, err)
	assert.Equal(t, int64(2), df.NumRows())
	assert.True(t, df.HasColumn("active"))
}

func TestReadJSON_EmptyArray(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "empty.json")
	err := os.WriteFile(path, []byte(`[]`), 0600)
	require.NoError(t, err)

	df, err := ReadJSON(path)
	require.NoError(t, err)
	assert.Equal(t, int64(0), df.NumRows())
}
