package field

// Int int type field
type Int struct{ exprOrderable }

var _ expression[int] = new(Int)

// Eq checks if the field is equal to the provided int value
func (field Int) Eq(value int) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided int value
func (field Int) Gt(value int) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided int value
func (field Int) Gte(value int) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided int value
func (field Int) Lt(value int) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided int value
func (field Int) Lte(value int) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided int values
func (field Int) In(values ...int) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided int values
func (field Int) NotIn(values ...int) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[int] = new(Int)

// Set set value
func (field Int) Set(value int) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided int value
func (field Int) Inc(value int) Mod {
	return mod{m: newInc(field.col, value)}
}

// toIntSlice converts a slice of int to a slice of interface{}
func (field Int) toSlice(values []int) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Int8 int8 type field
type Int8 Int

var _ expression[int8] = new(Int8)

// Eq checks if the field is equal to the provided int8 value
func (field Int8) Eq(value int8) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided int8 value
func (field Int8) Gt(value int8) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided int8 value
func (field Int8) Gte(value int8) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided int8 value
func (field Int8) Lt(value int8) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided int8 value
func (field Int8) Lte(value int8) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided int8 values
func (field Int8) In(values ...int8) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided int8 values
func (field Int8) NotIn(values ...int8) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[int8] = new(Int8)

// Set set value
func (field Int8) Set(value int8) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided int8 value
func (field Int8) Inc(value int8) Mod {
	return mod{m: newInc(field.col, value)}
}

// toInt8Slice converts a slice of int8 to a slice of interface{}
func (field Int8) toSlice(values []int8) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Int16 int16 type field
type Int16 Int

var _ expression[int16] = new(Int16)

// Eq checks if the field is equal to the provided int16 value
func (field Int16) Eq(value int16) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided int16 value
func (field Int16) Gt(value int16) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided int16 value
func (field Int16) Gte(value int16) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided int16 value
func (field Int16) Lt(value int16) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided int16 value
func (field Int16) Lte(value int16) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided int16 values
func (field Int16) In(values ...int16) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided int16 values
func (field Int16) NotIn(values ...int16) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[int16] = new(Int16)

// Set set value
func (field Int16) Set(value int16) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided int16 value
func (field Int16) Inc(value int16) Mod {
	return mod{m: newInc(field.col, value)}
}

// toInt16Slice converts a slice of int16 to a slice of interface{}
func (field Int16) toSlice(values []int16) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Int32 int32 type field
type Int32 Int

var _ expression[int32] = new(Int32)

// Eq checks if the field is equal to the provided int32 value
func (field Int32) Eq(value int32) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided int32 value
func (field Int32) Gt(value int32) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided int32 value
func (field Int32) Gte(value int32) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided int32 value
func (field Int32) Lt(value int32) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided int32 value
func (field Int32) Lte(value int32) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided int32 values
func (field Int32) In(values ...int32) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided int32 values
func (field Int32) NotIn(values ...int32) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[int32] = new(Int32)

// Set set value
func (field Int32) Set(value int32) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided int32 value
func (field Int32) Inc(value int32) Mod {
	return mod{m: newInc(field.col, value)}
}

// toInt32Slice converts a slice of int32 to a slice of interface{}
func (field Int32) toSlice(values []int32) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Int64 int64 type field
type Int64 Int

var _ expression[int64] = new(Int64)

// Eq checks if the field is equal to the provided int64 value
func (field Int64) Eq(value int64) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided int64 value
func (field Int64) Gt(value int64) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided int64 value
func (field Int64) Gte(value int64) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided int64 value
func (field Int64) Lt(value int64) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided int64 value
func (field Int64) Lte(value int64) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided int64 values
func (field Int64) In(values ...int64) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided int64 values
func (field Int64) NotIn(values ...int64) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[int64] = new(Int64)

// Set set value
func (field Int64) Set(value int64) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided int64 value
func (field Int64) Inc(value int64) Mod {
	return mod{m: newInc(field.col, value)}
}

// toInt64Slice converts a slice of int64 to a slice of interface{}
func (field Int64) toSlice(values []int64) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Uint uint type field
type Uint Int

var _ expression[uint] = new(Uint)

// Eq checks if the field is equal to the provided uint value
func (field Uint) Eq(value uint) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided uint value
func (field Uint) Gt(value uint) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided uint value
func (field Uint) Gte(value uint) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided uint value
func (field Uint) Lt(value uint) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided uint value
func (field Uint) Lte(value uint) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided uint values
func (field Uint) In(values ...uint) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided uint values
func (field Uint) NotIn(values ...uint) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[uint] = new(Uint)

// Set set value
func (field Uint) Set(value uint) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided uint value
func (field Uint) Inc(value uint) Mod {
	return mod{m: newInc(field.col, value)}
}

// toUintSlice converts a slice of uint to a slice of uinterface{}
func (field Uint) toSlice(values []uint) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Uint8 uint8 type field
type Uint8 Int

var _ expression[uint8] = new(Uint8)

// Eq checks if the field is equal to the provided uint8 value
func (field Uint8) Eq(value uint8) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided uint8 value
func (field Uint8) Gt(value uint8) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided uint8 value
func (field Uint8) Gte(value uint8) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided uint8 value
func (field Uint8) Lt(value uint8) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided uint8 value
func (field Uint8) Lte(value uint8) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided uint8 values
func (field Uint8) In(values ...uint8) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided uint8 values
func (field Uint8) NotIn(values ...uint8) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[uint8] = new(Uint8)

// Set set value
func (field Uint8) Set(value uint8) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided uint8 value
func (field Uint8) Inc(value uint8) Mod {
	return mod{m: newInc(field.col, value)}
}

// toUint8Slice converts a slice of uint8 to a slice of uinterface{}
func (field Uint8) toSlice(values []uint8) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Uint16 uint16 type field
type Uint16 Int

var _ expression[uint16] = new(Uint16)

// Eq checks if the field is equal to the provided uint16 value
func (field Uint16) Eq(value uint16) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided uint16 value
func (field Uint16) Gt(value uint16) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided uint16 value
func (field Uint16) Gte(value uint16) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided uint16 value
func (field Uint16) Lt(value uint16) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided uint16 value
func (field Uint16) Lte(value uint16) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided uint16 values
func (field Uint16) In(values ...uint16) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided uint16 values
func (field Uint16) NotIn(values ...uint16) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[uint16] = new(Uint16)

// Set set value
func (field Uint16) Set(value uint16) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided uint16 value
func (field Uint16) Inc(value uint16) Mod {
	return mod{m: newInc(field.col, value)}
}

// toUint16Slice converts a slice of uint16 to a slice of uinterface{}
func (field Uint16) toSlice(values []uint16) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Uint32 uint32 type field
type Uint32 Int

var _ expression[uint32] = new(Uint32)

// Eq checks if the field is equal to the provided uint32 value
func (field Uint32) Eq(value uint32) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided uint32 value
func (field Uint32) Gt(value uint32) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided uint32 value
func (field Uint32) Gte(value uint32) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided uint32 value
func (field Uint32) Lt(value uint32) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided uint32 value
func (field Uint32) Lte(value uint32) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided uint32 values
func (field Uint32) In(values ...uint32) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided uint32 values
func (field Uint32) NotIn(values ...uint32) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[uint32] = new(Uint32)

// Set set value
func (field Uint32) Set(value uint32) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided uint32 value
func (field Uint32) Inc(value uint32) Mod {
	return mod{m: newInc(field.col, value)}
}

// toUint32Slice converts a slice of uint32 to a slice of uinterface{}
func (field Uint32) toSlice(values []uint32) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}

// Uint64 uint64 type field
type Uint64 Int

var _ expression[uint64] = new(Uint64)

// Eq checks if the field is equal to the provided uint64 value
func (field Uint64) Eq(value uint64) Expr {
	return expr{e: Eq{Column: field.col, Value: value}}
}

// Gt checks if the field is greater than the provided uint64 value
func (field Uint64) Gt(value uint64) Expr {
	return expr{e: Gt{Column: field.col, Value: value}}
}

// Gte checks if the field is greater than or equal to the provided uint64 value
func (field Uint64) Gte(value uint64) Expr {
	return expr{e: Gte{Column: field.col, Value: value}}
}

// Lt checks if the field is less than the provided uint64 value
func (field Uint64) Lt(value uint64) Expr {
	return expr{e: Lt{Column: field.col, Value: value}}
}

// Lte checks if the field is less than or equal to the provided uint64 value
func (field Uint64) Lte(value uint64) Expr {
	return expr{e: Lte{Column: field.col, Value: value}}
}

// In checks if the field is in the provided uint64 values
func (field Uint64) In(values ...uint64) Expr {
	return expr{e: In{Column: field.col, Values: field.toSlice(values)}}
}

// NotIn checks if the field is not in the provided uint64 values
func (field Uint64) NotIn(values ...uint64) Expr {
	return expr{e: NotIn{Column: field.col, Values: field.toSlice(values)}}
}

var _ numericModifier[uint64] = new(Uint64)

// Set set value
func (field Uint64) Set(value uint64) Mod {
	return mod{m: Set{Column: field.col, Value: value}}
}

// Inc increments the field by the provided uint64 value
func (field Uint64) Inc(value uint64) Mod {
	return mod{m: newInc(field.col, value)}
}

// toUint64Slice converts a slice of uint64 to a slice of uinterface{}
func (field Uint64) toSlice(values []uint64) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
