package field

// Bool bool type field
type Bool struct{ expr }

var _ simpleExpression[bool] = new(Bool)
var _ simpleModifier[bool] = new(Bool)

// Eq checks if the field is equal to the provided bool value
func (field Bool) Eq(value bool) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Set set value
func (field Bool) Set(value bool) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}
