package expr

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/stretchr/testify/assert"
)

func TestErrorFunctions(t *testing.T) {
	t.Run("errTypeMismatch", func(t *testing.T) {
		err := errTypeMismatch("int64", "string value")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "type mismatch")
		assert.Contains(t, err.Error(), "int64")
		assert.Contains(t, err.Error(), "string")
	})

	t.Run("errUnsupportedType", func(t *testing.T) {
		// Use an actual arrow DataType
		err := errUnsupportedType(arrow.BinaryTypes.String)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported literal type")
		assert.Contains(t, err.Error(), "utf8")
	})

	t.Run("errLengthMismatch", func(t *testing.T) {
		err := errLengthMismatch(10, 20)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "length mismatch")
		assert.Contains(t, err.Error(), "10")
		assert.Contains(t, err.Error(), "20")
	})

	t.Run("errUnsupportedTypes", func(t *testing.T) {
		// Use actual arrow DataTypes
		err := errUnsupportedTypes("add", arrow.PrimitiveTypes.Int64, arrow.BinaryTypes.String)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported types")
		assert.Contains(t, err.Error(), "add")
		assert.Contains(t, err.Error(), "int64")
		assert.Contains(t, err.Error(), "utf8")
	})

	t.Run("errColumnNotFound", func(t *testing.T) {
		err := errColumnNotFound("missing_column")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "column not found")
		assert.Contains(t, err.Error(), "missing_column")
	})

	t.Run("errDivisionByZero", func(t *testing.T) {
		err := errDivisionByZero(5)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "division by zero")
		assert.Contains(t, err.Error(), "5")
	})
}
