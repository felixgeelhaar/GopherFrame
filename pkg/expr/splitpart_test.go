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

func TestSplitPart_Evaluate(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "path", Type: arrow.BinaryTypes.String},
	}, nil)

	pb := array.NewStringBuilder(pool)
	pb.AppendValues([]string{"a/b/c", "x/y", "single"}, nil)
	pArr := pb.NewArray()
	defer pArr.Release()
	pb.Release()

	record := array.NewRecord(schema, []arrow.Array{pArr}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Split by "/" and get index 0
	e := NewColumnExpr("path").SplitPart(NewLiteralExpr("/"), NewLiteralExpr(int64(0)))
	result, err := e.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strArr := result.(*array.String)
	assert.Equal(t, "a", strArr.Value(0))
	assert.Equal(t, "x", strArr.Value(1))
	assert.Equal(t, "single", strArr.Value(2))
}

func TestSplitPart_MiddleIndex(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "csv", Type: arrow.BinaryTypes.String},
	}, nil)

	cb := array.NewStringBuilder(pool)
	cb.AppendValues([]string{"a,b,c", "x,y,z"}, nil)
	cArr := cb.NewArray()
	defer cArr.Release()
	cb.Release()

	record := array.NewRecord(schema, []arrow.Array{cArr}, 2)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	e := NewColumnExpr("csv").SplitPart(NewLiteralExpr(","), NewLiteralExpr(int64(1)))
	result, err := e.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strArr := result.(*array.String)
	assert.Equal(t, "b", strArr.Value(0))
	assert.Equal(t, "y", strArr.Value(1))
}

func TestSplitPart_OutOfBounds(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "text", Type: arrow.BinaryTypes.String},
	}, nil)

	tb := array.NewStringBuilder(pool)
	tb.AppendValues([]string{"a,b"}, nil)
	tArr := tb.NewArray()
	defer tArr.Release()
	tb.Release()

	record := array.NewRecord(schema, []arrow.Array{tArr}, 1)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	e := NewColumnExpr("text").SplitPart(NewLiteralExpr(","), NewLiteralExpr(int64(5)))
	result, err := e.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	strArr := result.(*array.String)
	assert.Equal(t, "", strArr.Value(0))
}

func TestSplitPart_WithNull(t *testing.T) {
	pool := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "data", Type: arrow.BinaryTypes.String},
	}, nil)

	db := array.NewStringBuilder(pool)
	db.Append("a-b")
	db.AppendNull()
	db.Append("c-d")
	dArr := db.NewArray()
	defer dArr.Release()
	db.Release()

	record := array.NewRecord(schema, []arrow.Array{dArr}, 3)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	e := NewColumnExpr("data").SplitPart(NewLiteralExpr("-"), NewLiteralExpr(int64(0)))
	result, err := e.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	assert.True(t, result.IsNull(1))
	strArr := result.(*array.String)
	assert.Equal(t, "a", strArr.Value(0))
	assert.Equal(t, "c", strArr.Value(2))
}

func TestSplitPart_String(t *testing.T) {
	e := NewColumnExpr("path").SplitPart(NewLiteralExpr("/"), NewLiteralExpr(int64(0)))
	assert.Contains(t, e.String(), "split_part")
}
