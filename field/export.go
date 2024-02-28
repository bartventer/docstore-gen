package field

import (
	"strings"

	"gocloud.dev/docstore"
)

// Column represents a column in a table
type Column struct {
	Table string
	Name  string
	path  []string
}

// FieldPath returns the field path
func (c Column) FieldPath() docstore.FieldPath {
	return docstore.FieldPath(strings.Join(c.path, "."))
}

// Option field option
type Option func(Column) Column

// toColumn convert to column
func toColumn(table, column string, opts ...Option) Column {
	col := Column{Table: table, Name: column, path: []string{column}}
	for _, opt := range opts {
		col = opt(col)
	}
	return col
}

// newExprOrderable creates a new exprOrderable with the specified table, column, and options.
// It returns the created exprOrderable.
func newExprOrderable(table, column string, opts ...Option) exprOrderable {
	return exprOrderable{
		expr{col: toColumn(table, column, opts...)},
		nil,
	}
}

// NewField creates a new Field with the specified table, column, and options.
// It returns a Field struct that represents the field expression.
func NewField(table, column string, opts ...Option) Field {
	return Field{newExprOrderable(table, column, opts...)}
}

// NewString creates a new String field
func NewString(table, column string, opts ...Option) String {
	return String{newExprOrderable(table, column, opts...)}
}

// NewBytes creates a new Bytes field
func NewBytes(table, column string, opts ...Option) Bytes {
	return Bytes{newExprOrderable(table, column, opts...)}
}

// NewInt creates a new Int field
func NewInt(table, column string, opts ...Option) Int {
	return Int{newExprOrderable(table, column, opts...)}
}

// NewInt8 creates a new Int8 field
func NewInt8(table, column string, opts ...Option) Int8 {
	return Int8{newExprOrderable(table, column, opts...)}
}

// NewInt16 creates a new Int16 field
func NewInt16(table, column string, opts ...Option) Int16 {
	return Int16{newExprOrderable(table, column, opts...)}
}

// NewInt32 creates a new Int32 field
func NewInt32(table, column string, opts ...Option) Int32 {
	return Int32{newExprOrderable(table, column, opts...)}
}

// NewInt64 creates a new Int64 field
func NewInt64(table, column string, opts ...Option) Int64 {
	return Int64{newExprOrderable(table, column, opts...)}
}

// NewUint creates a new Uint field
func NewUint(table, column string, opts ...Option) Uint {
	return Uint{newExprOrderable(table, column, opts...)}
}

// NewUint8 creates a new Uint8 field
func NewUint8(table, column string, opts ...Option) Uint8 {
	return Uint8{newExprOrderable(table, column, opts...)}
}

// NewUint16 creates a new Uint16 field
func NewUint16(table, column string, opts ...Option) Uint16 {
	return Uint16{newExprOrderable(table, column, opts...)}
}

// NewUint32 creates a new Uint32 field
func NewUint32(table, column string, opts ...Option) Uint32 {
	return Uint32{newExprOrderable(table, column, opts...)}
}

// NewUint64 creates a new Uint64 field
func NewUint64(table, column string, opts ...Option) Uint64 {
	return Uint64{newExprOrderable(table, column, opts...)}
}

// NewFloat32 creates a new Float32 field
func NewFloat32(table, column string, opts ...Option) Float32 {
	return Float32{newExprOrderable(table, column, opts...)}
}

// NewFloat64 creates a new Float64 field
func NewFloat64(table, column string, opts ...Option) Float64 {
	return Float64{newExprOrderable(table, column, opts...)}
}

// NewTime creates a new Time field
func NewTime(table, column string, opts ...Option) Time {
	return Time{newExprOrderable(table, column, opts...)}
}

// NewBool creates a new Bool field
func NewBool(table, column string, opts ...Option) Bool {
	return Bool{expr{col: toColumn(table, column, opts...)}}
}
