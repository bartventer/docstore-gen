package model

import (
	"strings"

	"github.com/bartventer/docstore-gen/field"
)

// Field user input structures
type Field struct {
	Name        string
	Type        string
	ColumnName  string
	Tag         field.Tag
	DocstoreTag field.DocstoreTag
}

// GenType get the gen field type
func (m *Field) GenType() string {
	typ := strings.TrimLeft(m.Type, "*")
	switch typ {
	case "string", "bytes":
		return strings.Title(typ)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return strings.Title(typ)
	case "float64", "float32":
		return strings.Title(typ)
	case "bool":
		return strings.Title(typ)
	case "time.Time":
		return "Time"
	case "json.RawMessage", "[]byte":
		return "Bytes"
	case "serializer":
		return "Serializer"
	default:
		return "Field"
	}
}

// SourceCode source code
type SourceCode int

const (
	// Struct ...
	Struct SourceCode = iota
	// Table ...
	Table
	// Object ...
	Object
)
