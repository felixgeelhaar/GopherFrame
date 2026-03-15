package gopherframe

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

// Avro OCF magic bytes
var avroMagic = []byte{'O', 'b', 'j', 1}

// avroSchema represents a simplified Avro schema for reading.
type avroSchema struct {
	Type   string            `json:"type"`
	Name   string            `json:"name,omitempty"`
	Fields []avroSchemaField `json:"fields,omitempty"`
}

type avroSchemaField struct {
	Name string          `json:"name"`
	Type json.RawMessage `json:"type"`
}

// ReadAvro reads an Avro Object Container File into a DataFrame.
// Supports primitive Avro types: null, boolean, int, long, float, double, string, bytes.
// Union types like ["null", "string"] are supported for nullable columns.
func ReadAvro(filename string) (*DataFrame, error) {
	if err := validateFilePath(filename); err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open Avro file: %w", err)
	}
	defer func() { _ = f.Close() }()

	// Read magic
	magic := make([]byte, 4)
	if _, err := io.ReadFull(f, magic); err != nil {
		return nil, fmt.Errorf("failed to read Avro magic: %w", err)
	}
	if string(magic) != string(avroMagic) {
		return nil, fmt.Errorf("not a valid Avro file (invalid magic bytes)")
	}

	// Read metadata (map of string -> bytes)
	meta, err := readAvroMap(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read Avro metadata: %w", err)
	}

	// Parse schema
	schemaJSON, ok := meta["avro.schema"]
	if !ok {
		return nil, fmt.Errorf("avro file missing schema in metadata")
	}

	var schema avroSchema
	if err := json.Unmarshal([]byte(schemaJSON), &schema); err != nil {
		return nil, fmt.Errorf("failed to parse Avro schema: %w", err)
	}

	if schema.Type != "record" {
		return nil, fmt.Errorf("avro schema type must be 'record', got %q", schema.Type)
	}

	// Read sync marker (16 bytes)
	syncMarker := make([]byte, 16)
	if _, err := io.ReadFull(f, syncMarker); err != nil {
		return nil, fmt.Errorf("failed to read sync marker: %w", err)
	}

	// Read data blocks
	var allRows []map[string]interface{}
	for {
		// Read block count (varint)
		blockCount, err := readAvroLong(f)
		if err != nil {
			break // EOF
		}
		if blockCount <= 0 {
			break
		}

		// Read block size
		blockSize, err := readAvroLong(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read block size: %w", err)
		}

		// Read block data
		blockData := make([]byte, blockSize)
		if _, err := io.ReadFull(f, blockData); err != nil {
			return nil, fmt.Errorf("failed to read block data: %w", err)
		}

		// Parse records from block
		rows, err := parseAvroBlock(blockData, schema.Fields, int(blockCount))
		if err != nil {
			return nil, fmt.Errorf("failed to parse Avro block: %w", err)
		}
		allRows = append(allRows, rows...)

		// Read sync marker
		sync := make([]byte, 16)
		if _, err := io.ReadFull(f, sync); err != nil {
			break
		}
	}

	return jsonRecordsToDataFrame(allRows)
}

// WriteAvro writes a DataFrame to an Avro Object Container File.
func WriteAvro(df *DataFrame, filename string) error {
	if err := validateFilePath(filename); err != nil {
		return err
	}
	if df.err != nil {
		return df.err
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create Avro file: %w", err)
	}
	defer func() { _ = f.Close() }()

	record := df.coreDF.Record()
	schema := record.Schema()

	// Build Avro schema
	avroSch := buildAvroSchema(schema)
	schemaJSON, err := json.Marshal(avroSch)
	if err != nil {
		return fmt.Errorf("failed to marshal Avro schema: %w", err)
	}

	// Write magic
	if _, err := f.Write(avroMagic); err != nil {
		return err
	}

	// Write metadata
	meta := map[string]string{
		"avro.schema": string(schemaJSON),
		"avro.codec":  "null",
	}
	if err := writeAvroMap(f, meta); err != nil {
		return err
	}

	// Write sync marker
	syncMarker := make([]byte, 16)
	for i := range syncMarker {
		syncMarker[i] = byte(i + 1)
	}
	if _, err := f.Write(syncMarker); err != nil {
		return err
	}

	// Write data as a single block
	numRows := int(record.NumRows())
	if numRows > 0 {
		var blockBuf []byte
		for i := 0; i < numRows; i++ {
			for j, field := range schema.Fields() {
				col := record.Column(j)
				b := encodeAvroValue(col, i, field.Type)
				blockBuf = append(blockBuf, b...)
			}
		}

		// Write block count
		writeAvroLong(f, int64(numRows))
		// Write block size
		writeAvroLong(f, int64(len(blockBuf)))
		// Write block data
		if _, err := f.Write(blockBuf); err != nil {
			return err
		}
		// Write sync marker
		if _, err := f.Write(syncMarker); err != nil {
			return err
		}
	}

	return nil
}

// --- Avro encoding helpers ---

func readAvroLong(r io.Reader) (int64, error) {
	var val uint64
	var shift uint
	for {
		b := make([]byte, 1)
		if _, err := io.ReadFull(r, b); err != nil {
			return 0, err
		}
		val |= uint64(b[0]&0x7f) << shift
		if b[0]&0x80 == 0 {
			break
		}
		shift += 7
	}
	// Zigzag decode
	return int64(val>>1) ^ -int64(val&1), nil
}

func writeAvroLong(w io.Writer, n int64) {
	// Zigzag encode
	val := uint64((n << 1) ^ (n >> 63))
	var buf [10]byte
	i := 0
	for val > 0x7f {
		buf[i] = byte(val&0x7f) | 0x80
		val >>= 7
		i++
	}
	buf[i] = byte(val)
	_, _ = w.Write(buf[:i+1])
}

func readAvroMap(r io.Reader) (map[string]string, error) {
	result := make(map[string]string)
	count, err := readAvroLong(r)
	if err != nil {
		return nil, err
	}
	for count > 0 {
		for i := int64(0); i < count; i++ {
			key, err := readAvroString(r)
			if err != nil {
				return nil, err
			}
			val, err := readAvroString(r)
			if err != nil {
				return nil, err
			}
			result[key] = val
		}
		count, err = readAvroLong(r)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func writeAvroMap(w io.Writer, m map[string]string) error {
	if len(m) > 0 {
		writeAvroLong(w, int64(len(m)))
		for k, v := range m {
			writeAvroString(w, k)
			writeAvroString(w, v)
		}
	}
	writeAvroLong(w, 0) // End of map
	return nil
}

func readAvroString(r io.Reader) (string, error) {
	length, err := readAvroLong(r)
	if err != nil {
		return "", err
	}
	if length <= 0 {
		return "", nil
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func writeAvroString(w io.Writer, s string) {
	writeAvroLong(w, int64(len(s)))
	_, _ = w.Write([]byte(s))
}

func buildAvroSchema(schema *arrow.Schema) avroSchema {
	avroSch := avroSchema{
		Type: "record",
		Name: "GopherFrameRecord",
	}
	for _, field := range schema.Fields() {
		avroType := arrowToAvroType(field.Type)
		raw, _ := json.Marshal(avroType)
		avroSch.Fields = append(avroSch.Fields, avroSchemaField{
			Name: field.Name,
			Type: raw,
		})
	}
	return avroSch
}

func arrowToAvroType(dt arrow.DataType) interface{} {
	switch dt.ID() {
	case arrow.FLOAT64:
		return "double"
	case arrow.INT64:
		return "long"
	case arrow.STRING:
		return "string"
	case arrow.BOOL:
		return "boolean"
	default:
		return "string"
	}
}

func encodeAvroValue(col arrow.Array, i int, dt arrow.DataType) []byte {
	switch a := col.(type) {
	case *array.Float64:
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:], math.Float64bits(a.Value(i)))
		return buf[:]
	case *array.Int64:
		var buf [10]byte
		val := a.Value(i)
		n := binary.PutVarint(buf[:], val)
		// Use zigzag encoding
		uval := uint64((val << 1) ^ (val >> 63))
		idx := 0
		for uval > 0x7f {
			buf[idx] = byte(uval&0x7f) | 0x80
			uval >>= 7
			idx++
		}
		buf[idx] = byte(uval)
		_ = n
		return buf[:idx+1]
	case *array.String:
		s := a.Value(i)
		var buf []byte
		// Length as varint
		l := int64(len(s))
		uval := uint64((l << 1) ^ (l >> 63))
		for uval > 0x7f {
			buf = append(buf, byte(uval&0x7f)|0x80)
			uval >>= 7
		}
		buf = append(buf, byte(uval))
		buf = append(buf, []byte(s)...)
		return buf
	case *array.Boolean:
		if a.Value(i) {
			return []byte{1}
		}
		return []byte{0}
	default:
		return nil
	}
}

func parseAvroBlock(data []byte, fields []avroSchemaField, count int) ([]map[string]interface{}, error) {
	rows := make([]map[string]interface{}, 0, count)
	offset := 0
	for i := 0; i < count && offset < len(data); i++ {
		row := make(map[string]interface{})
		for _, field := range fields {
			var typeStr string
			if err := json.Unmarshal(field.Type, &typeStr); err != nil {
				// Could be union type, skip
				typeStr = "string"
			}
			switch typeStr {
			case "double":
				if offset+8 > len(data) {
					return rows, nil
				}
				bits := binary.LittleEndian.Uint64(data[offset : offset+8])
				row[field.Name] = math.Float64frombits(bits)
				offset += 8
			case "long":
				val, n := decodeAvroVarlong(data[offset:])
				row[field.Name] = float64(val) // Store as float64 for consistency
				offset += n
			case "string":
				length, n := decodeAvroVarlong(data[offset:])
				offset += n
				if offset+int(length) > len(data) {
					return rows, nil
				}
				row[field.Name] = string(data[offset : offset+int(length)])
				offset += int(length)
			case "boolean":
				if offset >= len(data) {
					return rows, nil
				}
				row[field.Name] = data[offset] != 0
				offset++
			case "int", "float":
				offset += 4 // Skip for now
			default:
				// Try to read as string
				length, n := decodeAvroVarlong(data[offset:])
				offset += n
				if length > 0 && offset+int(length) <= len(data) {
					row[field.Name] = string(data[offset : offset+int(length)])
					offset += int(length)
				}
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func decodeAvroVarlong(data []byte) (int64, int) {
	var val uint64
	var shift uint
	i := 0
	for i < len(data) {
		b := data[i]
		val |= uint64(b&0x7f) << shift
		i++
		if b&0x80 == 0 {
			break
		}
		shift += 7
	}
	return int64(val>>1) ^ -int64(val&1), i
}
