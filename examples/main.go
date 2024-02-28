package main

import (
	"fmt"
	"time"

	"github.com/bartventer/docstore-gen/examples/out/query"
)

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.Initialize()

	user := query.User
	fmt.Println(user.DateJoiend.Lt(time.Now().AddDate(0, 0, -1)))
}
