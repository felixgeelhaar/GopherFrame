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

// TestStringOperations_Contains tests Contains string operation
func TestStringOperations_Contains(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Create test schema and data
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "text", Type: arrow.BinaryTypes.String},
			{Name: "pattern", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	// Test data: various contains scenarios
	textBuilder := array.NewStringBuilder(pool)
	textBuilder.AppendValues([]string{"hello world", "test string", "golang", "empty", "null test"}, []bool{true, true, true, true, false})
	textArray := textBuilder.NewArray()
	defer textArray.Release()

	patternBuilder := array.NewStringBuilder(pool)
	patternBuilder.AppendValues([]string{"world", "str", "go", "", "test"}, []bool{true, true, true, true, true})
	patternArray := patternBuilder.NewArray()
	defer patternArray.Release()

	record := array.NewRecord(schema, []arrow.Array{textArray, patternArray}, 5)
	defer record.Release()

	df := core.NewDataFrame(record)
	defer df.Release()

	// Test Contains operation
	expr := Col("text").Contains(Lit("world"))
	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	// Verify result is boolean array
	boolResult, ok := result.(*array.Boolean)
	assert.True(t, ok, "Result should be boolean array")

	// Verify specific results
	assert.True(t, boolResult.Value(0), "Should find 'world' in 'hello world'")
}

// TestStringOperations_ContainsWithColumn tests Contains with column expression
func TestStringOperations_ContainsWithColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "text", Type: arrow.BinaryTypes.String},
			{Name: "pattern", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	textBuilder := array.NewStringBuilder(pool)
	textBuilder.AppendValues([]string{"hello world", "test string", "golang"}, nil)
	textArray := textBuilder.NewArray()
	defer textArray.Release()

	patternBuilder := array.NewStringBuilder(pool)
	patternBuilder.AppendValues([]string{"world", "str", "rust"}, nil)
	patternArray := patternBuilder.NewArray()
	defer patternArray.Release()

	record := array.NewRecord(schema, []arrow.Array{textArray, patternArray}, 3)
	defer record.Release()

	// Test Contains with column reference
	expr := Col("text").Contains(Col("pattern"))
	df := core.NewDataFrame(record)
	defer df.Release()

	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0), "'hello world' contains 'world'")
	assert.True(t, boolResult.Value(1), "'test string' contains 'str'")
	assert.False(t, boolResult.Value(2), "'golang' does not contain 'rust'")
}

// TestStringOperations_ContainsWithNulls tests Contains with null values
func TestStringOperations_ContainsWithNulls(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "text", Type: arrow.BinaryTypes.String},
			{Name: "pattern", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	textBuilder := array.NewStringBuilder(pool)
	textBuilder.AppendValues([]string{"hello", "world", "test"}, []bool{true, false, true})
	textArray := textBuilder.NewArray()
	defer textArray.Release()

	patternBuilder := array.NewStringBuilder(pool)
	patternBuilder.AppendValues([]string{"he", "or", "no"}, []bool{true, true, false})
	patternArray := patternBuilder.NewArray()
	defer patternArray.Release()

	record := array.NewRecord(schema, []arrow.Array{textArray, patternArray}, 3)
	defer record.Release()

	expr := Col("text").Contains(Col("pattern"))
	df := core.NewDataFrame(record)
	defer df.Release()

	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0), "'hello' contains 'he'")
	assert.True(t, boolResult.IsNull(1), "Null text should produce null result")
	assert.True(t, boolResult.IsNull(2), "Null pattern should produce null result")
}

// TestStringOperations_StartsWith tests StartsWith string operation
func TestStringOperations_StartsWith(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "text", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	textBuilder := array.NewStringBuilder(pool)
	textBuilder.AppendValues([]string{"hello world", "test string", "golang", "hello"}, nil)
	textArray := textBuilder.NewArray()
	defer textArray.Release()

	record := array.NewRecord(schema, []arrow.Array{textArray}, 4)
	defer record.Release()

	// Test StartsWith operation
	expr := Col("text").StartsWith(Lit("hello"))
	df := core.NewDataFrame(record)
	defer df.Release()

	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0), "'hello world' starts with 'hello'")
	assert.False(t, boolResult.Value(1), "'test string' does not start with 'hello'")
	assert.False(t, boolResult.Value(2), "'golang' does not start with 'hello'")
	assert.True(t, boolResult.Value(3), "'hello' starts with 'hello'")
}

// TestStringOperations_StartsWithColumn tests StartsWith with column expression
func TestStringOperations_StartsWithColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "text", Type: arrow.BinaryTypes.String},
			{Name: "prefix", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	textBuilder := array.NewStringBuilder(pool)
	textBuilder.AppendValues([]string{"hello world", "test string", "golang"}, nil)
	textArray := textBuilder.NewArray()
	defer textArray.Release()

	prefixBuilder := array.NewStringBuilder(pool)
	prefixBuilder.AppendValues([]string{"hello", "test", "rust"}, nil)
	prefixArray := prefixBuilder.NewArray()
	defer prefixArray.Release()

	record := array.NewRecord(schema, []arrow.Array{textArray, prefixArray}, 3)
	defer record.Release()

	expr := Col("text").StartsWith(Col("prefix"))
	df := core.NewDataFrame(record)
	defer df.Release()

	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0))
	assert.True(t, boolResult.Value(1))
	assert.False(t, boolResult.Value(2))
}

// TestStringOperations_EndsWith tests EndsWith string operation
func TestStringOperations_EndsWith(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "filename", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	filenameBuilder := array.NewStringBuilder(pool)
	filenameBuilder.AppendValues([]string{"test.go", "main.py", "data.json", "readme.md"}, nil)
	filenameArray := filenameBuilder.NewArray()
	defer filenameArray.Release()

	record := array.NewRecord(schema, []arrow.Array{filenameArray}, 4)
	defer record.Release()

	// Test EndsWith operation
	expr := Col("filename").EndsWith(Lit(".go"))
	df := core.NewDataFrame(record)
	defer df.Release()

	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0), "'test.go' ends with '.go'")
	assert.False(t, boolResult.Value(1), "'main.py' does not end with '.go'")
	assert.False(t, boolResult.Value(2), "'data.json' does not end with '.go'")
	assert.False(t, boolResult.Value(3), "'readme.md' does not end with '.go'")
}

// TestStringOperations_EndsWithColumn tests EndsWith with column expression
func TestStringOperations_EndsWithColumn(t *testing.T) {
	pool := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "filename", Type: arrow.BinaryTypes.String},
			{Name: "extension", Type: arrow.BinaryTypes.String},
		},
		nil,
	)

	filenameBuilder := array.NewStringBuilder(pool)
	filenameBuilder.AppendValues([]string{"test.go", "main.py", "data.json"}, nil)
	filenameArray := filenameBuilder.NewArray()
	defer filenameArray.Release()

	extensionBuilder := array.NewStringBuilder(pool)
	extensionBuilder.AppendValues([]string{".go", ".py", ".xml"}, nil)
	extensionArray := extensionBuilder.NewArray()
	defer extensionArray.Release()

	record := array.NewRecord(schema, []arrow.Array{filenameArray, extensionArray}, 3)
	defer record.Release()

	expr := Col("filename").EndsWith(Col("extension"))
	df := core.NewDataFrame(record)
	defer df.Release()

	result, err := expr.Evaluate(df)
	require.NoError(t, err)
	defer result.Release()

	boolResult := result.(*array.Boolean)
	assert.True(t, boolResult.Value(0))
	assert.True(t, boolResult.Value(1))
	assert.False(t, boolResult.Value(2))
}

// TestStringOperations_Errors tests error cases for string operations
func TestStringOperations_Errors(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Test type mismatch error for Contains
	t.Run("Contains with non-string type", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "number", Type: arrow.PrimitiveTypes.Int64},
				{Name: "text", Type: arrow.BinaryTypes.String},
			},
			nil,
		)

		numberBuilder := array.NewInt64Builder(pool)
		numberBuilder.AppendValues([]int64{1, 2, 3}, nil)
		numberArray := numberBuilder.NewArray()
		defer numberArray.Release()

		textBuilder := array.NewStringBuilder(pool)
		textBuilder.AppendValues([]string{"a", "b", "c"}, nil)
		textArray := textBuilder.NewArray()
		defer textArray.Release()

		record := array.NewRecord(schema, []arrow.Array{numberArray, textArray}, 3)
		defer record.Release()

		df := core.NewDataFrame(record)
		defer df.Release()

		expr := Col("number").Contains(Lit("test"))
		_, err := expr.Evaluate(df)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "requires string operands")
	})

	// Test StartsWith with non-string
	t.Run("StartsWith with non-string type", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "number", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		numberBuilder := array.NewFloat64Builder(pool)
		numberBuilder.AppendValues([]float64{1.5, 2.5}, nil)
		numberArray := numberBuilder.NewArray()
		defer numberArray.Release()

		record := array.NewRecord(schema, []arrow.Array{numberArray}, 2)
		defer record.Release()

		df := core.NewDataFrame(record)
		defer df.Release()

		expr := Col("number").StartsWith(Lit("1"))
		_, err := expr.Evaluate(df)
		assert.Error(t, err)
	})

	// Test EndsWith with non-string
	t.Run("EndsWith with non-string type", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "bool_col", Type: arrow.FixedWidthTypes.Boolean},
			},
			nil,
		)

		boolBuilder := array.NewBooleanBuilder(pool)
		boolBuilder.AppendValues([]bool{true, false}, nil)
		boolArray := boolBuilder.NewArray()
		defer boolArray.Release()

		record := array.NewRecord(schema, []arrow.Array{boolArray}, 2)
		defer record.Release()

		df := core.NewDataFrame(record)
		defer df.Release()

		expr := Col("bool_col").EndsWith(Lit("true"))
		_, err := expr.Evaluate(df)
		assert.Error(t, err)
	})
}

// TestHelperFunctions tests asInt64Array and asStringArray helper functions
func TestHelperFunctions(t *testing.T) {
	pool := memory.NewGoAllocator()

	t.Run("asInt64Array with valid int64", func(t *testing.T) {
		builder := array.NewInt64Builder(pool)
		builder.AppendValues([]int64{1, 2, 3}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		result, ok := asInt64Array(arr)
		assert.True(t, ok, "Should successfully cast to Int64 array")
		assert.NotNil(t, result)
		assert.Equal(t, int64(1), result.Value(0))
	})

	t.Run("asInt64Array with non-int64", func(t *testing.T) {
		builder := array.NewFloat64Builder(pool)
		builder.AppendValues([]float64{1.5, 2.5}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		result, ok := asInt64Array(arr)
		assert.False(t, ok, "Should fail to cast float64 to Int64 array")
		assert.Nil(t, result)
	})

	t.Run("asStringArray with valid string", func(t *testing.T) {
		builder := array.NewStringBuilder(pool)
		builder.AppendValues([]string{"hello", "world"}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		result, ok := asStringArray(arr)
		assert.True(t, ok, "Should successfully cast to String array")
		assert.NotNil(t, result)
		assert.Equal(t, "hello", result.Value(0))
	})

	t.Run("asStringArray with non-string", func(t *testing.T) {
		builder := array.NewInt64Builder(pool)
		builder.AppendValues([]int64{1, 2}, nil)
		arr := builder.NewArray()
		defer arr.Release()

		result, ok := asStringArray(arr)
		assert.False(t, ok, "Should fail to cast int64 to String array")
		assert.Nil(t, result)
	})
}

// TestInferDataType_AdditionalTypes tests inferDataType with more type cases
func TestInferDataType_AdditionalTypes(t *testing.T) {
	testCases := []struct {
		name     string
		value    interface{}
		expected arrow.DataType
	}{
		{"int8", int8(10), arrow.PrimitiveTypes.Int8},
		{"int16", int16(1000), arrow.PrimitiveTypes.Int16},
		{"int", int(12345), arrow.PrimitiveTypes.Int64}, // int promotes to int64
		{"uint8", uint8(255), arrow.PrimitiveTypes.Uint8},
		{"uint16", uint16(65535), arrow.PrimitiveTypes.Uint16},
		{"uint", uint(12345), arrow.PrimitiveTypes.Uint32}, // uint promotes to uint32
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := inferDataType(tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// TestBinaryExpr_EdgeCases tests edge cases for binary expressions
func TestBinaryExpr_EdgeCases(t *testing.T) {
	pool := memory.NewGoAllocator()

	// Test division by zero
	t.Run("Division by zero", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "numerator", Type: arrow.PrimitiveTypes.Float64},
				{Name: "denominator", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

		numBuilder := array.NewFloat64Builder(pool)
		numBuilder.AppendValues([]float64{10.0, 20.0}, nil)
		numArray := numBuilder.NewArray()
		defer numArray.Release()

		denomBuilder := array.NewFloat64Builder(pool)
		denomBuilder.AppendValues([]float64{0.0, 2.0}, nil)
		denomArray := denomBuilder.NewArray()
		defer denomArray.Release()

		record := array.NewRecord(schema, []arrow.Array{numArray, denomArray}, 2)
		defer record.Release()

		df := core.NewDataFrame(record)
		defer df.Release()

		expr := Col("numerator").Div(Col("denominator"))
		_, err := expr.Evaluate(df)
		// Division by zero produces an error in this implementation
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "division by zero")
	})

	// Test with null values in arithmetic
	t.Run("Arithmetic with nulls", func(t *testing.T) {
		schema := arrow.NewSchema(
			[]arrow.Field{
				{Name: "a", Type: arrow.PrimitiveTypes.Int64},
				{Name: "b", Type: arrow.PrimitiveTypes.Int64},
			},
			nil,
		)

		aBuilder := array.NewInt64Builder(pool)
		aBuilder.AppendValues([]int64{10, 20, 30}, []bool{true, false, true})
		aArray := aBuilder.NewArray()
		defer aArray.Release()

		bBuilder := array.NewInt64Builder(pool)
		bBuilder.AppendValues([]int64{1, 2, 3}, []bool{true, true, false})
		bArray := bBuilder.NewArray()
		defer bArray.Release()

		record := array.NewRecord(schema, []arrow.Array{aArray, bArray}, 3)
		defer record.Release()

		expr := Col("a").Add(Col("b"))
		df := core.NewDataFrame(record)
		defer df.Release()

		result, err := expr.Evaluate(df)
		require.NoError(t, err)
		defer result.Release()

		int64Result := result.(*array.Int64)
		assert.Equal(t, int64(11), int64Result.Value(0))
		assert.True(t, int64Result.IsNull(1), "Null in operand should produce null result")
		assert.True(t, int64Result.IsNull(2), "Null in operand should produce null result")
	})
}

// TestLiteralExpr_StringMethods tests string operation methods on LiteralExpr
func TestLiteralExpr_StringMethods(t *testing.T) {
	lit := Lit("test")

	// Test Contains method
	containsExpr := lit.Contains(Lit("es"))
	assert.NotNil(t, containsExpr)
	_, ok := containsExpr.(*BinaryExpr)
	assert.True(t, ok, "Contains should return a BinaryExpr")

	// Test StartsWith method
	startsWithExpr := lit.StartsWith(Lit("te"))
	assert.NotNil(t, startsWithExpr)
	_, ok = startsWithExpr.(*BinaryExpr)
	assert.True(t, ok, "StartsWith should return a BinaryExpr")

	// Test EndsWith method
	endsWithExpr := lit.EndsWith(Lit("st"))
	assert.NotNil(t, endsWithExpr)
	_, ok = endsWithExpr.(*BinaryExpr)
	assert.True(t, ok, "EndsWith should return a BinaryExpr")
}
