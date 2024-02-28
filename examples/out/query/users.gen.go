// Code generated by github.com/bartventer/docstore-gen. DO NOT EDIT.
// Code generated by github.com/bartventer/docstore-gen. DO NOT EDIT.
// Code generated by github.com/bartventer/docstore-gen. DO NOT EDIT.

package query

import "github.com/bartventer/docstore-gen/field"

// newUser create new User query struct
func newUser() user {
	_user := user{}

	tableName := _user.TableName()

	_user.ID = field.NewString(tableName, "id")
	_user.Name = field.NewString(tableName, "name")
	_user.Age = field.NewUint(tableName, "age")
	_user.DateJoiend = field.NewTime(tableName, "date_joined")
	_user.IsAdmin = field.NewBool(tableName, "is_admin")
	_user.Preferences = field.NewBytes(tableName, "preferences")
	_user.Expenditure = field.NewFloat64(tableName, "expenditure")

	return _user
}

type user struct {
	ID          field.String
	Name        field.String
	Age         field.Uint
	DateJoiend  field.Time
	IsAdmin     field.Bool
	Preferences field.Bytes
	Expenditure field.Float64
}

// TableName get table name
func (q user) TableName() string {
	return "users"
}
