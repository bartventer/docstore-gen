package field

import (
	"gocloud.dev/docstore"
)

// Expression expression interface
type Expression interface {
	// Build build expression
	Build() (fieldPath docstore.FieldPath, op string, value interface{})
}

// Expr a query expression about field
type Expr interface {
	ColumnName() string
	FieldPath() docstore.FieldPath
	Expression
}

// OrderByExpression order by expression interface
type OrderByExpression interface {
	// BuildOrderBy build order by expression
	BuildOrderBy() (field string, direction string)
}
type orderable interface {
	Asc() OrderByExpression
	Desc() OrderByExpression
}

// simpleExpression simple expression interface
type simpleExpression[T any] interface {
	Eq(value T) Expr
}

// expression expression interface
type expression[T any] interface {
	simpleExpression[T]
	Gt(value T) Expr
	Gte(value T) Expr
	In(values ...T) Expr
	Lt(value T) Expr
	Lte(value T) Expr
	NotIn(values ...T) Expr
}
