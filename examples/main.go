package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/bartventer/docstore-gen/examples/models"
	"github.com/bartventer/docstore-gen/examples/out/query"
	"gocloud.dev/docstore"

	// Enable in-memory driver
	_ "gocloud.dev/docstore/memdocstore"
)

func main() {
	// start your project here

	// initialize docstore query
	query.Initialize()
	// use q data object
	q := query.User

	// use new collection
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	coll, err := docstore.OpenCollection(ctx, fmt.Sprintf("mem://%s/%s", q.TableName(), q.ID.ColumnName())) // "mem://users/id"
	if err != nil {
		panic(err)
	}
	defer coll.Close()

	// Create a user
	user := models.User{
		ID:          "1",
		Name:        "John Doe",
		Age:         30,
		DateJoiend:  time.Now(),
		IsAdmin:     false,
		Preferences: []byte(`{"theme": "dark"}`),
		Expenditure: 100.50,
	}
	if err := q.WithCollection(coll).Create(ctx, &user); err != nil {
		panic(err)
	}
	fmt.Printf("created user: %+v\n", user)

	// Find the user
	var res []models.User
	iter := q.WithCollection(coll).Query().Where(q.ID.Eq("1")).Get(ctx)
	defer iter.Stop()
	for {
		var u models.User
		err := iter.Next(ctx, &u)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			res = append(res, u)
		}
	}
	if len(res) == 0 {
		panic("user not found")
	}
	fmt.Printf("found user: %+v\n", res[0])
}
