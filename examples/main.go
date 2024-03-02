package main

import (
	"context"
	"fmt"
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
		IsAdmin:     true,
		Expenditure: 100.00,
	}
	if err := q.WithCollection(coll).Create(ctx, &user); err != nil {
		panic(err)
	}
	fmt.Printf("[created user] via `Create` method:\n\t=>%+v\n", user)

	// Set the user's preferences theme to "light", increment the expenditure by 20.50, and unset is_admin
	user2 := models.User{ID: user.ID}
	if err := q.WithCollection(coll).Actions().Update(&user,
		q.Preferences.Set([]byte(`{"theme": "light"}`)),
		q.Expenditure.Inc(20.50),
		q.IsAdmin.Unset(),
	).Get(&user2).Do(ctx); err != nil {
		panic(err)
	}
	fmt.Printf("[updated user] changes via `Update` method:\n\t=>%+v\n", map[string]interface{}{
		"preferences": string(user2.Preferences),
		"expenditure": user2.Expenditure,
		"is_admin":    user2.IsAdmin,
	})

}
