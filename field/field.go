package field

import (
	"strings"

	"gocloud.dev/docstore"
)

var _ Expr = new(expr)

type expr struct {
	col Column
	e   Expression
}

// AddFieldPath implements Expr.
func (e expr) AddFieldPath(path docstore.FieldPath) expr {
	e.col.path = append(e.col.path, string(path))
	return e
}

// Build implements Expr.
func (e expr) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	if e.e == nil {
		return
	}
	return e.e.Build()
}

// ColumnName implements Expr.
func (e expr) ColumnName() string {
	return e.col.Name
}

// FieldPath implements Expr.
func (e expr) FieldPath() docstore.FieldPath {
	return docstore.FieldPath(strings.Join(e.col.path, "."))
}

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

// BuildOrderBy implements OrderableExpr.
func (e exprOrderable) BuildOrderBy() (field string, direction string) {
	if e.o == nil {
		return
	}
	return e.o.BuildOrderBy()
}

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

// toSlice converts a variadic interface{} argument to a slice of interface{}
func (field Field) toSlice(values ...interface{}) []interface{} {
	slice := make([]interface{}, len(values))
	copy(slice, values)
	return slice
}
