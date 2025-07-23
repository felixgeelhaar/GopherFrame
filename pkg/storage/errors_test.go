package storage

import (
	"testing"
)

func TestErrors(t *testing.T) {
	// Test that error variables are not nil and have expected messages
	if ErrInvalidSource == nil {
		t.Error("ErrInvalidSource should not be nil")
	}

	if ErrSourceNotFound == nil {
		t.Error("ErrSourceNotFound should not be nil")
	}

	if ErrSchemaConflict == nil {
		t.Error("ErrSchemaConflict should not be nil")
	}

	// Test error messages
	expectedMessages := map[error]string{
		ErrInvalidSource:  "invalid source identifier",
		ErrSourceNotFound: "data source not found",
		ErrSchemaConflict: "schema conflict",
	}

	for err, expectedMsg := range expectedMessages {
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message %q, got %q", expectedMsg, err.Error())
		}
	}
}
