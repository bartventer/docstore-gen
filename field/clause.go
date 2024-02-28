package field

import (
	"gocloud.dev/docstore"
)

// Eq equal to for where
type Eq struct {
	Column Column      // Column the column to compare
	Value  interface{} // Value the value to compare
}

// Build builds the where clause
func (eq Eq) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return eq.Column.FieldPath(), "=", eq.Value
}

// Gt greater than for where
type Gt Eq

// Build builds the where clause
func (gt Gt) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return gt.Column.FieldPath(), ">", gt.Value
}

// Gte greater than or equal to for where
type Gte Eq

// Build builds the where clause
func (gte Gte) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return gte.Column.FieldPath(), ">=", gte.Value
}

// Lt less than for where
type Lt Eq

// Build builds the where clause
func (lt Lt) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return lt.Column.FieldPath(), "<", lt.Value
}

// Lte less than or equal to for where
type Lte Eq

// Build builds the where clause
func (lte Lte) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return lte.Column.FieldPath(), "<=", lte.Value
}

// In in for where
type In struct {
	Column Column
	Values []interface{}
}

// Build builds the where clause
func (in In) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return in.Column.FieldPath(), "in", in.Values
}

// NotIn not in for where
type NotIn In

// Build builds the where clause
func (notIn NotIn) Build() (fieldPath docstore.FieldPath, op string, value interface{}) {
	return notIn.Column.FieldPath(), "not-in", notIn.Values
}

// Asc asc for order by
type Asc struct {
	Column Column
}

// Build builds the order by clause

// BuildOrderBy builds the order by clause
func (asc Asc) BuildOrderBy() (field string, direction string) {
	return string(asc.Column.FieldPath()), docstore.Ascending
}

// Desc desc for order by
type Desc Asc

// BuildOrderBy builds the order by clause
func (desc Desc) BuildOrderBy() (field string, direction string) {
	return string(desc.Column.FieldPath()), docstore.Descending
}
