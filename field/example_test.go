package field_test

import (
	"fmt"

	"github.com/bartventer/docstore-gen/field"
)

// ================================== expressions =============================

func ExampleField_Eq() {
	expr := field.NewField("user", "password").Eq("123")
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password = 123
}

func ExampleField_Gt() {
	expr := field.NewField("user", "password").Gt(123)
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password > 123
}

func ExampleField_Gte() {
	expr := field.NewField("user", "password").Gte(123)
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password >= 123
}

func ExampleField_Lt() {
	expr := field.NewField("user", "password").Lt(123)
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password < 123
}

func ExampleField_Lte() {
	expr := field.NewField("user", "password").Lte(123)
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password <= 123
}

func ExampleField_In() {
	expr := field.NewField("user", "password").In(123, 456, 789)
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password in [123 456 789]
}

func ExampleField_NotIn() {
	expr := field.NewField("user", "password").NotIn(123, 456, 789)
	fieldPath, op, value := expr.Build()
	fmt.Println(fieldPath, op, value)
	// Output:
	// password not-in [123 456 789]
}

// ================================== ordering ================================

func ExampleString_Asc() {
	expr := field.NewString("table", "column").Asc()
	fieldPath, direction := expr.BuildOrderBy()
	fmt.Println(fieldPath, direction)
	// Output:
	// column asc
}

func ExampleString_Desc() {
	expr := field.NewString("table", "column").Desc()
	fieldPath, direction := expr.BuildOrderBy()
	fmt.Println(fieldPath, direction)
	// Output:
	// column desc
}
