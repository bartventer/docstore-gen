package field

// String string type field
type String struct{ exprOrderable }

var _ expression[string] = new(String)

// Eq checks if the field is equal to the provided string value
func (field String) Eq(value string) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided string value
func (field String) Gt(value string) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided string value
func (field String) Gte(value string) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided string value
func (field String) Lt(value string) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided string value
func (field String) Lte(value string) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided string values
func (field String) In(values ...string) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided string values
func (field String) NotIn(values ...string) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

// toStringSlice converts a slice of string to a slice of interface{}
func (field String) toSlice(values []string) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Bytes []byte type field
type Bytes String

var _ expression[[]byte] = new(Bytes)

// Eq checks if the field is equal to the provided bytes value
func (field Bytes) Eq(value []byte) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided bytes value
func (field Bytes) Gt(value []byte) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided bytes value
func (field Bytes) Gte(value []byte) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided bytes value
func (field Bytes) Lt(value []byte) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided bytes value
func (field Bytes) Lte(value []byte) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided bytes values
func (field Bytes) In(values ...[]byte) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided bytes values
func (field Bytes) NotIn(values ...[]byte) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

// toSlice converts a slice of bytes to a slice of interface{}
func (field Bytes) toSlice(values [][]byte) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
