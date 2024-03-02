package field

import (
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
)

// Columner column interface
type Columner interface {
	// FieldPath returns the field path
	FieldPath() docstore.FieldPath
	// ColumnName returns the column name (same as the docstore tag name)
	ColumnName() string
}

// Mod modifier interface
type Mod interface {
	// BuildMod build a field path and value for the modifier
	BuildMod() (fieldPath docstore.FieldPath, value interface{})
}

// coreModifier unsetter coreModifier interface
type coreModifier interface {
	// Unset unset (delete) value from document
	Unset() Mod
}

// simpleModifier simple modifier interface for simple type
type simpleModifier[T any] interface {
	coreModifier
	// Set set value
	Set(value T) Mod
}

// Set set field value
type Set struct {
	Column Column // Column the column to set
	Value  any    // Value the value to set
}

// BuildMod build a field path and value for the modifier
func (set Set) BuildMod() (fieldPath docstore.FieldPath, value any) {
	return set.Column.FieldPath(), set.Value
}

// Unset unset field value
type Unset struct {
	Column Column // Column the column to unset
}

// BuildMod build a field path and value for the modifier
func (unset Unset) BuildMod() (fieldPath docstore.FieldPath, value interface{}) {
	return unset.Column.FieldPath(), nil
}

// numeric numeric type
type numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// numericModifier numeric modifier interface for numeric type
type numericModifier[T numeric] interface {
	simpleModifier[T]
	// Inc increment value
	Inc(value T) Mod
}

// Inc represents an increment operation on a numeric column.
type Inc[T numeric] struct {
	Column Column // Column represents the column to increment.
	Value  T      // Value represents the value to increment.
}

// BuildMod returns the field path and value for an increment operation.
// It constructs a docstore.FieldPath using the column's field path and
// creates a [driver.IncOp] with the specified increment amount.
//
// [driver.IncOp]: https://pkg.go.dev/gocloud.dev/docstore/driver#IncOp
func (inc Inc[T]) BuildMod() (fieldPath docstore.FieldPath, value interface{}) {
	return inc.Column.FieldPath(), driver.IncOp{Amount: inc.Value}
}

// newInc creates a new instance of the Inc struct with the specified column and value.
// The column parameter represents the column to be incremented, and the value parameter
// represents the value by which the column should be incremented.
// The type parameter T specifies the numeric type of the value parameter.
func newInc[T numeric](column Column, value T) Inc[T] {
	return Inc[T]{Column: column, Value: value}
}

// ConvertMods converts a slice of [Mod] objects to a [docstore.Mods] map.
// It iterates over each Mod in the input slice, calls the BuildMod method
// to get the field path and value, and adds them to the resulting map.
// The resulting map is then returned.
func ConvertMods(mods []Mod) docstore.Mods {
	m := make(docstore.Mods, len(mods))
	for _, mod := range mods {
		fieldPath, value := mod.BuildMod()
		m[fieldPath] = value
	}
	return m
}
