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

func TestReplace(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hello world")
	builder.Append("foo bar foo")
	builder.Append("no match here")
	builder.Append("")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Replace "world" with "earth"
	expr := Col("text").Replace(Lit("world"), Lit("earth"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "hello earth", strResult.Value(0))
	assert.Equal(t, "foo bar foo", strResult.Value(1)) // No match, unchanged
	assert.Equal(t, "no match here", strResult.Value(2))
	assert.Equal(t, "", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestReplaceMultiple(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("aaa bbb aaa")
	builder.Append("xxxyyyxxx")
	builder.Append("ababab")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 4)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Replace "aaa" with "zzz" - multiple occurrences
	expr := Col("text").Replace(Lit("aaa"), Lit("zzz"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "zzz bbb zzz", strResult.Value(0)) // Two replacements
	assert.Equal(t, "xxxyyyxxx", strResult.Value(1))   // No match
	assert.Equal(t, "ababab", strResult.Value(2))      // No match
	assert.True(t, strResult.IsNull(3))
}

func TestPadLeft(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hi")
	builder.Append("x")
	builder.Append("abc")
	builder.Append("")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Pad to length 5 with "0"
	expr := Col("text").PadLeft(Lit(5), Lit("0"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "000hi", strResult.Value(0))
	assert.Equal(t, "0000x", strResult.Value(1))
	assert.Equal(t, "00abc", strResult.Value(2))
	assert.Equal(t, "00000", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestPadRight(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hi")
	builder.Append("x")
	builder.Append("abc")
	builder.Append("")
	builder.AppendNull()

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Pad to length 5 with "*"
	expr := Col("text").PadRight(Lit(5), Lit("*"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "hi***", strResult.Value(0))
	assert.Equal(t, "x****", strResult.Value(1))
	assert.Equal(t, "abc**", strResult.Value(2))
	assert.Equal(t, "*****", strResult.Value(3))
	assert.True(t, strResult.IsNull(4))
}

func TestPadLeftAlreadyLong(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("already long enough")
	builder.Append("exact")
	builder.Append("hi")

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Pad to length 5 with "0"
	expr := Col("text").PadLeft(Lit(5), Lit("0"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "already long enough", strResult.Value(0)) // Already longer, unchanged
	assert.Equal(t, "exact", strResult.Value(1))               // Exactly 5 chars, unchanged
	assert.Equal(t, "000hi", strResult.Value(2))               // Shorter, padded
}

func TestPadRightAlreadyLong(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("already long enough")
	builder.Append("exact")

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Pad to length 5 with "-"
	expr := Col("text").PadRight(Lit(5), Lit("-"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "already long enough", strResult.Value(0)) // Already longer, unchanged
	assert.Equal(t, "exact", strResult.Value(1))               // Exactly 5 chars, unchanged
}

func TestReplace_TypeErrors(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create numeric array instead of string
	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	builder.Append(123)
	builder.Append(456)

	numArray := builder.NewArray()
	defer numArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "number", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{numArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("number").Replace(Lit("a"), Lit("b"))
	_, err := expr.Evaluate(df)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requires string type")
}

func TestPadLeft_TypeErrors(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewInt64Builder(pool)
	defer builder.Release()

	builder.Append(123)

	numArray := builder.NewArray()
	defer numArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "number", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{numArray}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	expr := Col("number").PadLeft(Lit(5), Lit("0"))
	_, err := expr.Evaluate(df)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requires string type")
}

func TestReplace_EmptyPattern(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hello")
	builder.Append("world")

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Replace empty string with "X" - inserts between every character
	expr := Col("text").Replace(Lit(""), Lit("X"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	// strings.ReplaceAll("hello", "", "X") inserts X between every char
	assert.Equal(t, "XhXeXlXlXoX", strResult.Value(0))
	assert.Equal(t, "XwXoXrXlXdX", strResult.Value(1))
}

func TestPadLeft_MultiCharPad(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("hi")
	builder.Append("a")

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Pad with multi-char pad string "ab"
	expr := Col("text").PadLeft(Lit(6), Lit("ab"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "ababhi", strResult.Value(0)) // Need 4 chars of padding from "ab" repeated
	assert.Equal(t, "ababaa", strResult.Value(1)) // Need 5 chars of padding from "ab" repeated
}

func TestReplace_ChainedWithOtherOps(t *testing.T) {
	pool := memory.NewGoAllocator()

	builder := array.NewStringBuilder(pool)
	defer builder.Release()

	builder.Append("Hello World")
	builder.Append("Foo Bar")

	strArray := builder.NewArray()
	defer strArray.Release()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	record := array.NewRecord(schema, []arrow.Array{strArray}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Chain: Replace then Upper
	expr := Col("text").Replace(Lit(" "), Lit("_")).Upper()
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strResult := result.(*array.String)
	assert.Equal(t, "HELLO_WORLD", strResult.Value(0))
	assert.Equal(t, "FOO_BAR", strResult.Value(1))
}

func TestTernaryExpr_String(t *testing.T) {
	expr := Col("text").Replace(Lit("a"), Lit("b"))
	assert.Contains(t, expr.String(), "replace")

	expr2 := Col("text").PadLeft(Lit(5), Lit("0"))
	assert.Contains(t, expr2.String(), "pad_left")

	expr3 := Col("text").PadRight(Lit(5), Lit("0"))
	assert.Contains(t, expr3.String(), "pad_right")
}
