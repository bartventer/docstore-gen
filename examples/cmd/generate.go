package main

import (
	docstoregen "github.com/bartventer/docstore-gen"
	"github.com/bartventer/docstore-gen/examples/cmd/models"
)

func main() {
	g := docstoregen.NewGenerator(docstoregen.Config{
		OutPath: "../out/query",
		OutFile: "gen.go",
	})

	g.ApplyInterface(&models.User{})

	g.Execute()
}
