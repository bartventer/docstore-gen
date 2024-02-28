// Package docstoregen provides a code generator for the [go-cloud.dev/docstore] package.
//
// The generator creates a query code file for the go-cloud.dev/docstore package.
//
// The generator's basic configuration is defined in the Config struct.
//
// The generator's main function is Execute, which generates the query code file.
//
// Example:
//
//	package main
//
//	import (
//		"github.com/bartventer/docstore-gen"
//	)
//
//	type User struct {
//		ID   string `json:"id" docstore:"id"`
//		Name string `json:"name" docstore:"name"`
//	}
//
//	func (User) TableName() string { return "user" }
//
//	type Copmany struct {
//		ID   string `json:"id" docstore:"id"`
//		Name string `json:"name" docstore:"name"`
//	}
//
//	func (Copmany) TableName() string { return "company" }
//
//	func main() {
//		g := docstoregen.NewGenerator(docstoregen.Config{
//			OutPath: "query",
//		})
//		g.ApplyInterface(&User{}, &Copmany{})
//		g.Execute()
//	}
//
// [go-cloud.dev/docstore]: https://gocloud.dev/howto/docstore/
package docstoregen
