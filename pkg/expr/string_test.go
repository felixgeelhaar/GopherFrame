package expr

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/felixgeelhaar/GopherFrame/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpr_Upper(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hello")
	builder.Append("World")
	builder.Append("ALREADY UPPER")
	builder.Append("MiXeD CaSe")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Upper()
	expr := Col("text").Upper()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	strResult := result.(*array.String)
	assert.Equal(t, "HELLO", strResult.Value(0))
	assert.Equal(t, "WORLD", strResult.Value(1))
	assert.Equal(t, "ALREADY UPPER", strResult.Value(2))
	assert.Equal(t, "MIXED CASE", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestExpr_Lower(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("HELLO")
	builder.Append("World")
	builder.Append("already lower")
	builder.Append("MiXeD CaSe")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Lower()
	expr := Col("text").Lower()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	strResult := result.(*array.String)
	assert.Equal(t, "hello", strResult.Value(0))
	assert.Equal(t, "world", strResult.Value(1))
	assert.Equal(t, "already lower", strResult.Value(2))
	assert.Equal(t, "mixed case", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestExpr_Trim(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings with various whitespace
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("  hello  ")
	builder.Append("\tworld\t")
	builder.Append("\n  test  \n")
	builder.Append("no whitespace")
	builder.Append("   ")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 6)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Trim()
	expr := Col("text").Trim()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	strResult := result.(*array.String)
	assert.Equal(t, "hello", strResult.Value(0))
	assert.Equal(t, "world", strResult.Value(1))
	assert.Equal(t, "test", strResult.Value(2))
	assert.Equal(t, "no whitespace", strResult.Value(3))
	assert.Equal(t, "", strResult.Value(4)) // All whitespace becomes empty
	assert.True(t, strResult.IsNull(5))
}

func TestExpr_TrimLeft(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("  hello")
	builder.Append("\tworld")
	builder.Append("  test  ")
	builder.Append("no leading")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test TrimLeft()
	expr := Col("text").TrimLeft()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	strResult := result.(*array.String)
	assert.Equal(t, "hello", strResult.Value(0))
	assert.Equal(t, "world", strResult.Value(1))
	assert.Equal(t, "test  ", strResult.Value(2)) // Trailing whitespace preserved
	assert.Equal(t, "no leading", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestExpr_TrimRight(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hello  ")
	builder.Append("world\t")
	builder.Append("  test  ")
	builder.Append("no trailing")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test TrimRight()
	expr := Col("text").TrimRight()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	strResult := result.(*array.String)
	assert.Equal(t, "hello", strResult.Value(0))
	assert.Equal(t, "world", strResult.Value(1))
	assert.Equal(t, "  test", strResult.Value(2)) // Leading whitespace preserved
	assert.Equal(t, "no trailing", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestExpr_Length(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hello")
	builder.Append("")
	builder.Append("longer string")
	builder.Append("a")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Length()
	expr := Col("text").Length()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results - Length returns Int64
	int64Result := result.(*array.Int64)
	assert.Equal(t, int64(5), int64Result.Value(0))
	assert.Equal(t, int64(0), int64Result.Value(1)) // Empty string has length 0
	assert.Equal(t, int64(13), int64Result.Value(2))
	assert.Equal(t, int64(1), int64Result.Value(3))
	assert.True(t, int64Result.IsNull(4)) // Null propagates
}

func TestExpr_Match(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	strBuilder := array.NewStringBuilder(pool)
	defer strBuilder.Release()

	strBuilder.Append("hello@example.com")
	strBuilder.Append("not-an-email")
	strBuilder.Append("another@test.org")
	strBuilder.Append("invalid@")
	strBuilder.AppendNull()

	strArray := strBuilder.NewArray()
	defer strArray.Release()

	// Create pattern strings (email regex)
	patternBuilder := array.NewStringBuilder(pool)
	defer patternBuilder.Release()

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	for i := 0; i < 5; i++ {
		patternBuilder.Append(emailRegex)
	}

	patternArray := patternBuilder.NewArray()
	defer patternArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "email", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "pattern", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray, patternArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Match()
	expr := Col("email").Match(Col("pattern"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results - Match returns Boolean
	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0))  // Valid email
	assert.False(t, boolResult.Value(1)) // Not an email
	assert.True(t, boolResult.Value(2))  // Valid email
	assert.False(t, boolResult.Value(3)) // Invalid email
	assert.True(t, boolResult.IsNull(4)) // Null propagates
}

func TestExpr_Match_WithLiteral(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("test123")
	builder.Append("abc")
	builder.Append("456")
	builder.Append("test456")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Match() with literal pattern - match strings containing digits
	expr := Col("text").Match(Lit(`\d+`))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0))  // Contains digits
	assert.False(t, boolResult.Value(1)) // No digits
	assert.True(t, boolResult.Value(2))  // All digits
	assert.True(t, boolResult.Value(3))  // Contains digits
	assert.True(t, boolResult.IsNull(4)) // Null propagates
}

func TestExpr_Match_InvalidRegex(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test string
	strBuilder := array.NewStringBuilder(pool)
	defer strBuilder.Release()
	strBuilder.Append("test")
	strArray := strBuilder.NewArray()
	defer strArray.Release()

	// Create invalid regex pattern
	patternBuilder := array.NewStringBuilder(pool)
	defer patternBuilder.Release()
	patternBuilder.Append("[invalid(regex") // Unclosed bracket
	patternArray := patternBuilder.NewArray()
	defer patternArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
		{Name: "pattern", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray, patternArray}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Match() with invalid regex - should return error
	expr := Col("text").Match(Col("pattern"))
	_, err := expr.Evaluate(df)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid regex pattern")
}

func TestExpr_String_ChainedOperations(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("  HELLO WORLD  ")
	builder.Append("  goodbye  ")
	builder.Append("MiXeD")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 4)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test chained operations: Trim() then Lower()
	expr := Col("text").Trim().Lower()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify results
	strResult := result.(*array.String)
	assert.Equal(t, "hello world", strResult.Value(0))
	assert.Equal(t, "goodbye", strResult.Value(1))
	assert.Equal(t, "mixed", strResult.Value(2))
	assert.True(t, strResult.IsNull(3))
}

func TestExpr_String_EmptyStrings(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings including empty strings
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("")
	builder.Append("   ")
	builder.Append("a")

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test various operations on empty strings
	t.Run("Upper on empty string", func(t *testing.T) {
		expr := Col("text").Upper()
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		strResult := result.(*array.String)
		assert.Equal(t, "", strResult.Value(0))
	})

	t.Run("Length on empty string", func(t *testing.T) {
		expr := Col("text").Length()
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		int64Result := result.(*array.Int64)
		assert.Equal(t, int64(0), int64Result.Value(0))
	})

	t.Run("Trim on whitespace-only string", func(t *testing.T) {
		expr := Col("text").Trim()
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		strResult := result.(*array.String)
		assert.Equal(t, "", strResult.Value(1)) // Whitespace-only becomes empty
	})
}

func TestExpr_String_TypeErrors(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create numeric array instead of string
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	builder.Append(123)
	builder.Append(456)

	numArray := builder.NewArray()
	defer numArray.Release()

	// Create DataFrame with numeric column
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "number", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{numArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	t.Run("Upper on non-string", func(t *testing.T) {
		expr := Col("number").Upper()
		_, err := expr.Evaluate(df)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requires string type")
	})

	t.Run("Lower on non-string", func(t *testing.T) {
		expr := Col("number").Lower()
		_, err := expr.Evaluate(df)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requires string type")
	})

	t.Run("Trim on non-string", func(t *testing.T) {
		expr := Col("number").Trim()
		_, err := expr.Evaluate(df)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requires string type")
	})

	t.Run("Length on non-string", func(t *testing.T) {
		expr := Col("number").Length()
		_, err := expr.Evaluate(df)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requires string type")
	})

	t.Run("Match on non-string", func(t *testing.T) {
		expr := Col("number").Match(Lit("test"))
		_, err := expr.Evaluate(df)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requires string operands")
	})
}

func TestExpr_String_UnicodeSupport(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test strings with unicode characters
	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("héllo wörld")
	builder.Append("ПРИВЕТ МИР")
	builder.Append("こんにちは")
	builder.Append("مرحبا")

	strArray := builder.NewArray()
	defer strArray.Release()

	// Create DataFrame
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 4)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	t.Run("Upper with unicode", func(t *testing.T) {
		expr := Col("text").Upper()
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		strResult := result.(*array.String)
		assert.Equal(t, "HÉLLO WÖRLD", strResult.Value(0))
		assert.Equal(t, "ПРИВЕТ МИР", strResult.Value(1))
	})

	t.Run("Lower with unicode", func(t *testing.T) {
		expr := Col("text").Lower()
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		strResult := result.(*array.String)
		assert.Equal(t, "héllo wörld", strResult.Value(0))
		assert.Equal(t, "привет мир", strResult.Value(1))
	})

	t.Run("Length with unicode", func(t *testing.T) {
		expr := Col("text").Length()
		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		int64Result := result.(*array.Int64)
		// Length counts bytes, not characters in Go strings.len()
		// For proper character counting would need utf8.RuneCountInString
		assert.Greater(t, int64Result.Value(0), int64(0))
		assert.Greater(t, int64Result.Value(2), int64(0)) // Japanese text
		assert.Greater(t, int64Result.Value(3), int64(0)) // Arabic text
	})
}
