package field

import "time"

// Time time type field
type Time struct{ exprOrderable }

var _ expression[time.Time] = new(Time)

// Eq checks if the field is equal to the provided time value
func (field Time) Eq(value time.Time) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided time value
func (field Time) Gt(value time.Time) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided time value
func (field Time) Gte(value time.Time) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided time value
func (field Time) Lt(value time.Time) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided time value
func (field Time) Lte(value time.Time) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided time values
func (field Time) In(values ...time.Time) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided time values
func (field Time) NotIn(values ...time.Time) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ simpleModifier[time.Time] = new(Time)

// Set set value
func (field Time) Set(value time.Time) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// toSlice converts a slice of time.Time to a slice of interface{}
func (field Time) toSlice(values []time.Time) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
