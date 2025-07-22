package persistence

import "errors"

// Common persistence backend errors.
var (
	ErrBackendNotFound      = errors.New("persistence backend not found")
	ErrSourceNotFound       = errors.New("data source not found")
	ErrInvalidSource        = errors.New("invalid source identifier")
	ErrPermissionDenied     = errors.New("permission denied")
	ErrSchemaConflict       = errors.New("schema conflict")
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrConnectionFailed     = errors.New("connection failed")
	ErrTimeout              = errors.New("operation timeout")
	ErrCorruptData          = errors.New("corrupt data detected")
	ErrQuotaExceeded        = errors.New("storage quota exceeded")
	ErrReadOnly             = errors.New("backend is read-only")
)
