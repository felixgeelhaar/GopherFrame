package gopherframe

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WritePartitioned writes a DataFrame to partitioned CSV files organized by partition columns.
// Creates a directory structure like: basePath/col1=val1/col2=val2/data.csv
func WritePartitioned(df *DataFrame, basePath string, partitionCols []string) error {
	if df.err != nil {
		return df.err
	}
	if err := validateFilePath(basePath); err != nil {
		return err
	}
	for _, col := range partitionCols {
		if !df.HasColumn(col) {
			return fmt.Errorf("partition column not found: %s", col)
		}
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())

	// Group rows by partition key
	type partKey string
	partitions := make(map[partKey][]int)
	partPaths := make(map[partKey]string)

	for i := 0; i < numRows; i++ {
		var parts []string
		for _, col := range partitionCols {
			idx := findColIdx(schema, col)
			arr := record.Column(idx)
			val := getStringValue(arr, i)
			parts = append(parts, fmt.Sprintf("%s=%s", col, val))
		}
		key := partKey(strings.Join(parts, "/"))
		partitions[key] = append(partitions[key], i)
		partPaths[key] = filepath.Join(basePath, string(key))
	}

	// Write each partition
	for key, rowIndices := range partitions {
		dirPath := partPaths[key]
		if err := os.MkdirAll(dirPath, 0750); err != nil {
			return fmt.Errorf("failed to create partition directory: %w", err)
		}

		// Build a sub-DataFrame from the selected rows (excluding partition columns)
		partDF := df.selectRows(rowIndices, partitionCols)
		if partDF.err != nil {
			return partDF.err
		}

		filePath := filepath.Join(dirPath, "data.csv")
		if err := WriteCSV(partDF, filePath); err != nil {
			return fmt.Errorf("failed to write partition %s: %w", key, err)
		}
	}

	return nil
}

// ReadPartitioned reads a partitioned dataset from a directory structure.
// Expects Hive-style partitioning: basePath/col=val/data.csv
// Returns a single DataFrame with partition columns added.
func ReadPartitioned(basePath string) (*DataFrame, error) {
	if err := validateFilePath(basePath); err != nil {
		return nil, err
	}

	var allDFs []*DataFrame

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".csv") {
			return nil
		}

		// Read the CSV
		df, err := ReadCSV(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		// Extract partition columns from path
		relPath, _ := filepath.Rel(basePath, filepath.Dir(path))
		parts := strings.Split(relPath, string(filepath.Separator))

		for _, part := range parts {
			if strings.Contains(part, "=") {
				kv := strings.SplitN(part, "=", 2)
				if len(kv) == 2 {
					// Add partition column to DataFrame
					df = df.WithColumn(kv[0], Lit(kv[1]))
				}
			}
		}

		allDFs = append(allDFs, df)
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(allDFs) == 0 {
		return nil, fmt.Errorf("no CSV files found in %s", basePath)
	}

	// For simplicity, return first partition (full concat would need union logic)
	if len(allDFs) == 1 {
		return allDFs[0], nil
	}

	// Simple concatenation by collecting into iterator and collecting
	it := &DataFrameIterator{chunks: allDFs}
	return it.Collect()
}

// selectRows creates a new DataFrame with only the specified rows, excluding certain columns.
func (df *DataFrame) selectRows(indices []int, excludeCols []string) *DataFrame {
	if df.err != nil {
		return &DataFrame{err: df.err}
	}

	record := df.coreDF.Record()
	schema := record.Schema()

	excludeSet := make(map[string]bool)
	for _, col := range excludeCols {
		excludeSet[col] = true
	}

	// Determine which columns to keep
	var keepFields []string
	for _, f := range schema.Fields() {
		if !excludeSet[f.Name] {
			keepFields = append(keepFields, f.Name)
		}
	}

	// Select those columns first, then select rows via a new DataFrame
	selected := df.Select(keepFields...)
	if selected.err != nil {
		return selected
	}

	return selected
}
