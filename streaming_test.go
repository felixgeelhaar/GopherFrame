package gopherframe

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadCSVChunked_Basic(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.csv")

	data := "name,score\nAlice,95.5\nBob,87.3\nCarol,92.1\nDave,78.9\nEve,88.0\n"
	err := os.WriteFile(path, []byte(data), 0600)
	require.NoError(t, err)

	it, err := ReadCSVChunked(path, 2)
	require.NoError(t, err)
	assert.Equal(t, 3, it.Len()) // 5 rows / 2 = 3 chunks (2+2+1)

	// First chunk
	chunk := it.Next()
	assert.Equal(t, int64(2), chunk.NumRows())
	assert.True(t, chunk.HasColumn("name"))

	// Second chunk
	chunk = it.Next()
	assert.Equal(t, int64(2), chunk.NumRows())

	// Third chunk (1 remaining row)
	chunk = it.Next()
	assert.Equal(t, int64(1), chunk.NumRows())

	// No more
	assert.False(t, it.HasNext())
}

func TestReadCSVChunked_Collect(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.csv")

	data := "id,value\n1,10.0\n2,20.0\n3,30.0\n"
	err := os.WriteFile(path, []byte(data), 0600)
	require.NoError(t, err)

	it, err := ReadCSVChunked(path, 2)
	require.NoError(t, err)

	df, err := it.Collect()
	require.NoError(t, err)
	assert.Equal(t, int64(3), df.NumRows())
	assert.True(t, df.HasColumn("id"))
	assert.True(t, df.HasColumn("value"))
}

func TestReadCSVChunked_ForEachChunk(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.csv")

	data := "x\n1\n2\n3\n4\n"
	err := os.WriteFile(path, []byte(data), 0600)
	require.NoError(t, err)

	it, err := ReadCSVChunked(path, 3)
	require.NoError(t, err)

	totalRows := int64(0)
	err = it.ForEachChunk(func(chunk *DataFrame) error {
		totalRows += chunk.NumRows()
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, int64(4), totalRows)
}

func TestReadCSVChunked_InvalidFile(t *testing.T) {
	_, err := ReadCSVChunked("/nonexistent.csv", 10)
	assert.Error(t, err)
}

func TestReadCSVChunked_InvalidChunkSize(t *testing.T) {
	_, err := ReadCSVChunked("test.csv", 0)
	assert.Error(t, err)
}
