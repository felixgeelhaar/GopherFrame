package storage

import "errors"

// Common storage backend errors.
var (
	// ErrBackendNotFound is returned when a requested backend is not registered.
	ErrBackendNotFound = errors.New("storage backend not found")

	// ErrSourceNotFound is returned when a requested data source doesn't exist.
	ErrSourceNotFound = errors.New("data source not found")

	// ErrInvalidSource is returned when a source identifier is malformed.
	ErrInvalidSource = errors.New("invalid source identifier")

	// ErrPermissionDenied is returned when access to a source is denied.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrSchemaConflict is returned when schemas don't match expectations.
	ErrSchemaConflict = errors.New("schema conflict")

	// ErrUnsupportedOperation is returned when an operation isn't supported.
	ErrUnsupportedOperation = errors.New("unsupported operation")

	// ErrConnectionFailed is returned when connection to backend fails.
	ErrConnectionFailed = errors.New("connection failed")

	// ErrTimeout is returned when an operation times out.
	ErrTimeout = errors.New("operation timeout")

	// ErrCorruptData is returned when data integrity issues are detected.
	ErrCorruptData = errors.New("corrupt data detected")

	// ErrQuotaExceeded is returned when storage quota is exceeded.
	ErrQuotaExceeded = errors.New("storage quota exceeded")

	// ErrReadOnly is returned when trying to write to a read-only backend.
	ErrReadOnly = errors.New("backend is read-only")
)
