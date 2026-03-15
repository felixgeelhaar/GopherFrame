package gopherframe

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// --- Backpressure Streaming ---

// StreamingReader provides backpressure-aware chunked reading via Go channels.
// The reader produces DataFrames into a channel, and the consumer controls pace
// by reading from the channel. If the consumer is slow, the producer blocks.
type StreamingReader struct {
	chunkCh  chan *DataFrame
	errCh    chan error
	cancelFn context.CancelFunc
}

// ReadCSVStreaming creates a backpressure-aware streaming CSV reader.
// bufferSize controls how many chunks can be buffered ahead of the consumer.
// The channel blocks the producer when the buffer is full.
func ReadCSVStreaming(filename string, chunkSize, bufferSize int) (*StreamingReader, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}
	if chunkSize <= 0 || bufferSize <= 0 {
		return nil, fmt.Errorf("chunk size and buffer size must be positive")
	}

	ctx, cancel := context.WithCancel(context.Background())
	sr := &StreamingReader{
		chunkCh:  make(chan *DataFrame, bufferSize),
		errCh:    make(chan error, 1),
		cancelFn: cancel,
	}

	go func() {
		defer close(sr.chunkCh)
		it, err := ReadCSVChunked(filename, chunkSize)
		if err != nil {
			sr.errCh <- err
			return
		}
		for it.HasNext() {
			select {
			case <-ctx.Done():
				return
			case sr.chunkCh <- it.Next():
			}
		}
	}()

	return sr, nil
}

// Chunks returns the channel of DataFrames. Range over this to consume chunks.
func (sr *StreamingReader) Chunks() <-chan *DataFrame {
	return sr.chunkCh
}

// Err returns any error that occurred during reading.
func (sr *StreamingReader) Err() error {
	select {
	case err := <-sr.errCh:
		return err
	default:
		return nil
	}
}

// Cancel stops the streaming reader.
func (sr *StreamingReader) Cancel() {
	sr.cancelFn()
}

// --- Partition Pruning ---

// ReadPartitionedWithPruning reads a partitioned dataset, skipping partitions
// that don't match the given predicates. This avoids reading unnecessary files.
//
// predicates is a map of partition column name to allowed values.
// Only partitions where all predicates match are read.
func ReadPartitionedWithPruning(basePath string, predicates map[string][]string) (*DataFrame, error) {
	if err := validateFilePath(basePath); err != nil {
		return nil, err
	}

	predicateSet := make(map[string]map[string]bool)
	for col, vals := range predicates {
		predicateSet[col] = make(map[string]bool)
		for _, v := range vals {
			predicateSet[col][v] = true
		}
	}

	var allDFs []*DataFrame

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".csv") {
			return nil
		}

		// Extract partition columns from path
		relPath, _ := filepath.Rel(basePath, filepath.Dir(path))
		parts := strings.Split(relPath, string(filepath.Separator))

		// Check if this partition matches all predicates
		for _, part := range parts {
			if strings.Contains(part, "=") {
				kv := strings.SplitN(part, "=", 2)
				if len(kv) == 2 {
					col, val := kv[0], kv[1]
					if allowed, ok := predicateSet[col]; ok {
						if !allowed[val] {
							return nil // Prune: skip this partition
						}
					}
				}
			}
		}

		// Read the matching partition
		df, err := ReadCSV(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		// Add partition columns
		for _, part := range parts {
			if strings.Contains(part, "=") {
				kv := strings.SplitN(part, "=", 2)
				if len(kv) == 2 {
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
		return nil, fmt.Errorf("no matching partitions found in %s", basePath)
	}
	if len(allDFs) == 1 {
		return allDFs[0], nil
	}

	it := &DataFrameIterator{chunks: allDFs}
	return it.Collect()
}

// --- Multi-file Parallel Reads ---

// ReadCSVParallel reads multiple CSV files concurrently and returns a combined DataFrame.
// maxWorkers controls the number of concurrent file readers.
func ReadCSVParallel(filenames []string, maxWorkers int) (*DataFrame, error) {
	if len(filenames) == 0 {
		return nil, fmt.Errorf("no files to read")
	}
	if maxWorkers <= 0 {
		maxWorkers = 4
	}

	type result struct {
		df  *DataFrame
		err error
		idx int
	}

	results := make([]result, len(filenames))
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i, filename := range filenames {
		wg.Add(1)
		go func(idx int, fn string) {
			defer wg.Done()
			sem <- struct{}{}        // Acquire semaphore
			defer func() { <-sem }() // Release semaphore

			df, err := ReadCSV(fn)
			results[idx] = result{df: df, err: err, idx: idx}
		}(i, filename)
	}
	wg.Wait()

	// Check for errors and collect DataFrames
	var dfs []*DataFrame
	for _, r := range results {
		if r.err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", filenames[r.idx], r.err)
		}
		dfs = append(dfs, r.df)
	}

	if len(dfs) == 1 {
		return dfs[0], nil
	}

	it := &DataFrameIterator{chunks: dfs}
	return it.Collect()
}

// ReadJSONParallel reads multiple JSON files concurrently.
func ReadJSONParallel(filenames []string, maxWorkers int) (*DataFrame, error) {
	if len(filenames) == 0 {
		return nil, fmt.Errorf("no files to read")
	}
	if maxWorkers <= 0 {
		maxWorkers = 4
	}

	type result struct {
		df  *DataFrame
		err error
		idx int
	}

	results := make([]result, len(filenames))
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i, filename := range filenames {
		wg.Add(1)
		go func(idx int, fn string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			df, err := ReadJSON(fn)
			results[idx] = result{df: df, err: err, idx: idx}
		}(i, filename)
	}
	wg.Wait()

	var dfs []*DataFrame
	for _, r := range results {
		if r.err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", filenames[r.idx], r.err)
		}
		dfs = append(dfs, r.df)
	}

	if len(dfs) == 1 {
		return dfs[0], nil
	}

	it := &DataFrameIterator{chunks: dfs}
	return it.Collect()
}
