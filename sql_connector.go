package gopherframe

import (
	"database/sql"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// ReadSQL reads data from a SQL database query into a DataFrame.
// Requires a *sql.DB connection and a SQL query string.
// Supports standard database/sql drivers (PostgreSQL, MySQL, SQLite, etc.)
func ReadSQL(db *sql.DB, query string) (*DataFrame, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection cannot be nil")
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer func() { _ = rows.Close() }()

	// Get column info
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, fmt.Errorf("failed to get column types: %w", err)
	}

	// Collect all rows
	var allRows [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		allRows = append(allRows, values)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	pool := memory.NewGoAllocator()
	numRows := len(allRows)

	// Build Arrow columns
	var fields []arrow.Field
	var arrowColumns []arrow.Array

	for i, col := range columns {
		arrowType := sqlTypeToArrow(columnTypes[i])
		fields = append(fields, arrow.Field{Name: col, Type: arrowType})

		switch arrowType.ID() {
		case arrow.FLOAT64:
			b := array.NewFloat64Builder(pool)
			for _, row := range allRows {
				if row[i] == nil {
					b.AppendNull()
				} else {
					switch v := row[i].(type) {
					case float64:
						b.Append(v)
					case int64:
						b.Append(float64(v))
					case []byte:
						b.AppendNull()
					default:
						b.AppendNull()
					}
				}
			}
			arrowColumns = append(arrowColumns, b.NewArray())
			b.Release()

		case arrow.INT64:
			b := array.NewInt64Builder(pool)
			for _, row := range allRows {
				if row[i] == nil {
					b.AppendNull()
				} else {
					switch v := row[i].(type) {
					case int64:
						b.Append(v)
					case float64:
						b.Append(int64(v))
					case []byte:
						b.AppendNull()
					default:
						b.AppendNull()
					}
				}
			}
			arrowColumns = append(arrowColumns, b.NewArray())
			b.Release()

		default: // STRING
			b := array.NewStringBuilder(pool)
			for _, row := range allRows {
				if row[i] == nil {
					b.AppendNull()
				} else {
					switch v := row[i].(type) {
					case string:
						b.Append(v)
					case []byte:
						b.Append(string(v))
					default:
						b.Append(fmt.Sprintf("%v", v))
					}
				}
			}
			arrowColumns = append(arrowColumns, b.NewArray())
			b.Release()
		}
	}

	schema := arrow.NewSchema(fields, nil)
	record := array.NewRecord(schema, arrowColumns, int64(numRows))
	return NewDataFrame(record), nil
}

// WriteSQL writes a DataFrame to a SQL database table.
// Creates the table if it doesn't exist (using CREATE TABLE IF NOT EXISTS).
func WriteSQL(df *DataFrame, db *sql.DB, tableName string) error {
	if df.err != nil {
		return df.err
	}
	if db == nil {
		return fmt.Errorf("database connection cannot be nil")
	}

	record := df.coreDF.Record()
	schema := record.Schema()
	numRows := int(record.NumRows())

	// Create table
	var colDefs []string
	for _, field := range schema.Fields() {
		sqlType := arrowTypeToSQL(field.Type)
		colDefs = append(colDefs, fmt.Sprintf("%s %s", field.Name, sqlType))
	}
	createSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName,
		joinStrings(colDefs, ", "))

	if _, err := db.Exec(createSQL); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Insert rows
	placeholders := make([]string, len(schema.Fields()))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	insertSQL := fmt.Sprintf("INSERT INTO %s VALUES (%s)", tableName,
		joinStrings(placeholders, ", "))

	for i := 0; i < numRows; i++ {
		values := make([]interface{}, len(schema.Fields()))
		for j := range schema.Fields() {
			col := record.Column(j)
			if col.IsNull(i) {
				values[j] = nil
			} else {
				values[j] = getTypedValue(col, i)
			}
		}
		if _, err := db.Exec(insertSQL, values...); err != nil {
			return fmt.Errorf("failed to insert row %d: %w", i, err)
		}
	}

	return nil
}

// sqlTypeToArrow maps SQL column types to Arrow types.
func sqlTypeToArrow(ct *sql.ColumnType) arrow.DataType {
	typeName := ct.DatabaseTypeName()
	switch typeName {
	case "INTEGER", "INT", "BIGINT", "SMALLINT", "TINYINT":
		return arrow.PrimitiveTypes.Int64
	case "REAL", "FLOAT", "DOUBLE", "NUMERIC", "DECIMAL":
		return arrow.PrimitiveTypes.Float64
	case "BOOLEAN", "BOOL":
		return arrow.FixedWidthTypes.Boolean
	default:
		return arrow.BinaryTypes.String
	}
}

// arrowTypeToSQL maps Arrow types to SQL types.
func arrowTypeToSQL(dt arrow.DataType) string {
	switch dt.ID() {
	case arrow.INT64:
		return "BIGINT"
	case arrow.FLOAT64:
		return "DOUBLE PRECISION"
	case arrow.BOOL:
		return "BOOLEAN"
	default:
		return "TEXT"
	}
}

// joinStrings joins strings with a separator (avoids importing strings in this file).
func joinStrings(s []string, sep string) string {
	result := ""
	for i, v := range s {
		if i > 0 {
			result += sep
		}
		result += v
	}
	return result
}
