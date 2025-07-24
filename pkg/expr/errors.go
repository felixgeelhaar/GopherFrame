package expr

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
)

// Common expression errors
func errTypeMismatch(expected string, actual interface{}) error {
	return fmt.Errorf("type mismatch: expected %s, got %T", expected, actual)
}

func errUnsupportedType(dataType arrow.DataType) error {
	return fmt.Errorf("unsupported literal type: %s", dataType)
}

func errLengthMismatch(left, right int) error {
	return fmt.Errorf("array length mismatch: %d vs %d", left, right)
}

func errUnsupportedTypes(operation string, left, right arrow.DataType) error {
	return fmt.Errorf("unsupported types for %s: %s and %s", operation, left, right)
}

func errColumnNotFound(name string) error {
	return fmt.Errorf("column not found: %s", name)
}

func errDivisionByZero(index int) error {
	return fmt.Errorf("division by zero at index %d", index)
}
