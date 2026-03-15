package gopherframe

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// FuzzReadCSV fuzzes the CSV reader with arbitrary input data.
func FuzzReadCSV(f *testing.F) {
	// Seed corpus with valid CSV
	f.Add([]byte("name,age\nAlice,30\nBob,25\n"))
	f.Add([]byte("a,b,c\n1,2,3\n"))
	f.Add([]byte("col\n"))
	f.Add([]byte(""))

	f.Fuzz(func(t *testing.T, data []byte) {
		tmpDir := t.TempDir()
		path := filepath.Join(tmpDir, "fuzz.csv")
		if err := os.WriteFile(path, data, 0600); err != nil {
			return
		}

		// Should not panic
		df, err := ReadCSV(path)
		if err != nil {
			return
		}
		if df != nil {
			_ = df.NumRows()
			_ = df.NumCols()
		}
	})
}

// FuzzReadJSON fuzzes the JSON reader with arbitrary input data.
func FuzzReadJSON(f *testing.F) {
	// Seed corpus
	f.Add([]byte(`[{"a":1,"b":"hello"}]`))
	f.Add([]byte(`[{"x":true},{"x":false}]`))
	f.Add([]byte(`[]`))
	f.Add([]byte(`[{}]`))

	f.Fuzz(func(t *testing.T, data []byte) {
		tmpDir := t.TempDir()
		path := filepath.Join(tmpDir, "fuzz.json")
		if err := os.WriteFile(path, data, 0600); err != nil {
			return
		}

		df, err := ReadJSON(path)
		if err != nil {
			return
		}
		if df != nil {
			_ = df.NumRows()
			_ = df.NumCols()
		}
	})
}

// FuzzReadNDJSON fuzzes the NDJSON reader with arbitrary input data.
func FuzzReadNDJSON(f *testing.F) {
	f.Add([]byte("{\"a\":1}\n{\"a\":2}\n"))
	f.Add([]byte("{\"x\":\"hello\"}\n"))
	f.Add([]byte(""))

	f.Fuzz(func(t *testing.T, data []byte) {
		tmpDir := t.TempDir()
		path := filepath.Join(tmpDir, "fuzz.ndjson")
		if err := os.WriteFile(path, data, 0600); err != nil {
			return
		}

		df, err := ReadNDJSON(path)
		if err != nil {
			return
		}
		if df != nil {
			_ = df.NumRows()
			_ = df.NumCols()
		}
	})
}

// FuzzJSONRoundtrip fuzzes JSON read/write roundtrip.
func FuzzJSONRoundtrip(f *testing.F) {
	f.Add([]byte(`[{"name":"Alice","score":95.5}]`))

	f.Fuzz(func(t *testing.T, data []byte) {
		// Try to parse as valid JSON array of objects
		var records []map[string]interface{}
		if err := json.Unmarshal(data, &records); err != nil {
			return
		}
		if len(records) == 0 {
			return
		}

		tmpDir := t.TempDir()
		inPath := filepath.Join(tmpDir, "in.json")
		outPath := filepath.Join(tmpDir, "out.json")

		if err := os.WriteFile(inPath, data, 0600); err != nil {
			return
		}

		df, err := ReadJSON(inPath)
		if err != nil {
			return
		}

		// Write back
		_ = WriteJSON(df, outPath)
	})
}
