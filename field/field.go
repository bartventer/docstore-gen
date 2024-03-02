// Package field provides a way to create field expressions for use with the docstore package.
package field

import (
	"strings"

	"gocloud.dev/docstore"
)

// ================================ expr ================================

var _ Expr = new(expr)
var _ Mod = new(expr)
var _ coreModifier = new(expr)

type expr struct {
	col Column
	e   Expression
	m   Mod
}

// AddFieldPath adds a field path to the column.
//
// Example:
//
//		field := field.Field{Column: field.Column{Name: "name"}}
//	 // "name"
//		field = field.AddFieldPath("user")
//	 // "user.name"
func (e expr) AddFieldPath(path docstore.FieldPath) expr {
	e.col.path = append(e.col.path, string(path))
	return e
}

// Build returns the field path, operator, and value for the expression.
// If the expression is nil, it returns zero values for all three return types.
func (e expr) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	if e.e == nil {
		return
	}
	return e.e.Build()
}

// ColumnName returns the column name (same as the docstore tag name)
func (e expr) ColumnName() string {
	return e.col.Name
}

// FieldPath returns the [docstore.FieldPath] for the column.
//
// [docstore.FieldPath]: https://pkg.go.dev/gocloud.dev/docstore#FieldPath
func (e expr) FieldPath() docstore.FieldPath {
	return docstore.FieldPath(strings.Join(e.col.path, "."))
}

// BuildMod returns the [docstore.FieldPath] and value for the modifier.
//
// [docstore.FieldPath]: https://pkg.go.dev/gocloud.dev/docstore#FieldPath
func (e expr) BuildMod() (fieldPath docstore.FieldPath, value interface{}) {
	if e.m == nil {
		return
	}
	return e.m.BuildMod()
}

// Unset returns a modifier to unset (delete) the value from the document.
func (e expr) Unset() Mod {
	return mod{m: Unset{Column: e.col}}
}

// ================================ exprOrderable ================================

// exprOrderable represents a field that can be ordered
type exprOrderable struct {
	expr
	o OrderByExpression
}

var _ orderable = new(exprOrderable)

// Asc returns the field in ascending order
func (e exprOrderable) Asc() OrderByExpression {
	return exprOrderable{o: Asc{Column: e.col}}
}

// Desc returns the field in descending order
func (e exprOrderable) Desc() OrderByExpression {
	return exprOrderable{o: Desc{Column: e.col}}
}

// BuildOrderBy returns the field and direction for the order by expression
func (e exprOrderable) BuildOrderBy() (field string, direction string) {
	if e.o == nil {
		return
	}
	return e.o.BuildOrderBy()
}

// ================================ mod ================================

// mod represents a modifier struct
type mod struct {
	m Mod
}

var _ Mod = new(mod)

// BuildMod returns the [docstore.FieldPath] and value for the modifier.
//
// [docstore.FieldPath]: https://pkg.go.dev/gocloud.dev/docstore#FieldPath
func (m mod) BuildMod() (fieldPath docstore.FieldPath, value interface{}) {
	if m.m == nil {
		return
	}
	return m.m.BuildMod()
}

// ================================ Field ================================

// Field represents a standard field struct
type Field struct{ exprOrderable }

var _ expression[interface{}] = new(Field)

// Eq checks if the field is equal to the provided value
func (field Field) Eq(value interface{}) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided value
func (field Field) Gt(value interface{}) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided value
func (field Field) Gte(value interface{}) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided value
func (field Field) Lt(value interface{}) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided value
func (field Field) Lte(value interface{}) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided values
func (field Field) In(values ...interface{}) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values...)}}
}

// NotIn checks if the field is not in the provided values
func (field Field) NotIn(values ...interface{}) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values...)}}
}

var _ simpleModifier[interface{}] = new(Field)

// Set set value
func (field Field) Set(value interface{}) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// toSlice converts a variadic interface{} argument to a slice of interface{}
func (field Field) toSlice(values ...interface{}) []interface{} {
	slice := make([]interface{}, len(values))
	copy(slice, values)
	return slice
}
