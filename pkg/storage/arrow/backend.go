// Package arrow provides an Apache Arrow storage backend implementation.
// This is the default backend that handles Arrow IPC format files.
package arrow

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/ipc"
	"github.com/felixgeelhaar/GopherFrame/pkg/storage"
)

// Backend implements storage.Backend for Apache Arrow IPC format.
// This is the reference implementation and provides optimal performance
// for Arrow-native operations.
type Backend struct {
	// Future: connection pooling, caching, etc.
}

// NewBackend creates a new Arrow storage backend.
func NewBackend() storage.Backend {
	return &Backend{}
}

// Read implements storage.Backend.Read for Arrow IPC files.
func (b *Backend) Read(ctx context.Context, source string, opts storage.ReadOptions) (storage.RecordReader, error) {
	// Validate source path
	if source == "" {
		return nil, storage.ErrInvalidSource
	}

	// Open the file
	file, err := os.Open(source)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, storage.ErrSourceNotFound
		}
		return nil, fmt.Errorf("failed to open Arrow file: %w", err)
	}

	// Create Arrow IPC file reader
	reader, err := ipc.NewFileReader(file)
	if err != nil {
		_ = file.Close()
		return nil, fmt.Errorf("failed to create Arrow file reader: %w", err)
	}

	// Create our record reader wrapper
	return &recordReader{
		reader: reader,
		file:   file,
		opts:   opts,
		ctx:    ctx,
	}, nil
}

// Write implements storage.Backend.Write for Arrow IPC files.
func (b *Backend) Write(ctx context.Context, destination string, records storage.RecordReader, opts storage.WriteOptions) error {
	// Validate destination
	if destination == "" {
		return storage.ErrInvalidSource
	}

	// Create destination directory if needed
	dir := filepath.Dir(destination)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Check if file exists and handle overwrite
	if !opts.Overwrite {
		if _, err := os.Stat(destination); err == nil {
			return fmt.Errorf("destination exists and overwrite not specified: %s", destination)
		}
	}

	// Create the file
	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create Arrow file: %w", err)
	}
	defer func() { _ = file.Close() }()

	// Create Arrow IPC file writer
	writer, err := ipc.NewFileWriter(file, ipc.WithSchema(records.Schema()))
	if err != nil {
		return fmt.Errorf("failed to create Arrow file writer: %w", err)
	}
	defer func() { _ = writer.Close() }()

	// Write all records
	for records.Next() {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		record := records.Record()
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	// Check for reader errors
	if err := records.Err(); err != nil {
		return fmt.Errorf("error reading records: %w", err)
	}

	return nil
}

// Scan implements storage.Backend.Scan for Arrow files.
func (b *Backend) Scan(ctx context.Context, pattern string) ([]storage.SourceInfo, error) {
	// Use filepath.Glob to find matching files
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to scan pattern %s: %w", pattern, err)
	}

	var sources []storage.SourceInfo
	for _, match := range matches {
		// Only include files with Arrow extensions
		ext := strings.ToLower(filepath.Ext(match))
		if ext != ".arrow" && ext != ".ipc" {
			continue
		}

		// Get file info
		info, err := os.Stat(match)
		if err != nil {
			continue // Skip files we can't stat
		}

		// Try to read schema without loading data
		schema, err := b.Schema(ctx, match)
		if err != nil {
			schema = nil // Schema read failed, but file exists
		}

		source := storage.SourceInfo{
			Name:     filepath.Base(match),
			Path:     match,
			Size:     info.Size(),
			Schema:   schema,
			Modified: info.ModTime().Unix(),
			Metadata: map[string]string{
				"type": "arrow",
				"ext":  ext,
			},
		}

		sources = append(sources, source)
	}

	return sources, nil
}

// Schema implements storage.Backend.Schema for Arrow files.
func (b *Backend) Schema(_ context.Context, source string) (*arrow.Schema, error) {
	// Open the file
	file, err := os.Open(source)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, storage.ErrSourceNotFound
		}
		return nil, fmt.Errorf("failed to open Arrow file: %w", err)
	}
	defer func() { _ = file.Close() }()

	// Create Arrow IPC file reader
	reader, err := ipc.NewFileReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create Arrow file reader: %w", err)
	}
	defer func() { _ = reader.Close() }()

	// Return the schema
	return reader.Schema(), nil
}

// Close implements storage.Backend.Close.
func (b *Backend) Close() error {
	// Arrow backend doesn't maintain persistent connections
	return nil
}

// recordReader wraps an Arrow IPC file reader to implement storage.RecordReader.
type recordReader struct {
	reader *ipc.FileReader
	file   *os.File
	opts   storage.ReadOptions
	ctx    context.Context

	currentIndex int
	err          error
}

// Next implements storage.RecordReader.Next.
func (r *recordReader) Next() bool {
	if r.err != nil {
		return false
	}

	// Check context cancellation
	if r.ctx.Err() != nil {
		r.err = r.ctx.Err()
		return false
	}

	// Check if we have more records
	if r.currentIndex >= r.reader.NumRecords() {
		return false
	}

	// Check limit
	if r.opts.Limit > 0 && int64(r.currentIndex) >= r.opts.Limit {
		return false
	}

	return true
}

// Record implements storage.RecordReader.Record.
func (r *recordReader) Record() arrow.Record {
	if r.currentIndex >= r.reader.NumRecords() {
		return nil
	}

	record, err := r.reader.RecordAt(r.currentIndex)
	if err != nil {
		r.err = err
		return nil
	}

	r.currentIndex++
	return record
}

// Schema implements storage.RecordReader.Schema.
func (r *recordReader) Schema() *arrow.Schema {
	return r.reader.Schema()
}

// Err implements storage.RecordReader.Err.
func (r *recordReader) Err() error {
	return r.err
}

// Close implements storage.RecordReader.Close.
func (r *recordReader) Close() error {
	if r.reader != nil {
		_ = r.reader.Close()
	}
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}
