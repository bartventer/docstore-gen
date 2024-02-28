package field

// Bool bool type field
type Bool struct{ expr }

var _ simpleExpression[bool] = new(Bool)

// Eq checks if the field is equal to the provided bool value
func (field Bool) Eq(value bool) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}
