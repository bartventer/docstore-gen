package field

// Float64 float64 type field
type Float64 struct{ exprOrderable }

var _ expression[float64] = new(Float64)

// Eq checks if the field is equal to the provided float64 value
func (field Float64) Eq(value float64) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided float64 value
func (field Float64) Gt(value float64) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided float64 value
func (field Float64) Gte(value float64) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided float64 value
func (field Float64) Lt(value float64) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided float64 value
func (field Float64) Lte(value float64) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided float64 values
func (field Float64) In(values ...float64) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided float64 values
func (field Float64) NotIn(values ...float64) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

// toSlice converts a slice of float64 to a slice of interface{}
func (field Float64) toSlice(values []float64) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Float32 float32 type field
type Float32 Float64

var _ expression[float32] = new(Float32)

// Eq checks if the field is equal to the provided float32 value
func (field Float32) Eq(value float32) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided float32 value
func (field Float32) Gt(value float32) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided float32 value
func (field Float32) Gte(value float32) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided float32 value
func (field Float32) Lt(value float32) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided float32 value
func (field Float32) Lte(value float32) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided float32 values
func (field Float32) In(values ...float32) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided float32 values
func (field Float32) NotIn(values ...float32) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

// toSlice converts a slice of float32 to a slice of interface{}
func (field Float32) toSlice(values []float32) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
