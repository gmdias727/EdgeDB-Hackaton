package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/edgedb/edgedb-go"
)

type User struct {
	ID   edgedb.UUID `edgedb:"id"`
	Name string      `edgedb:"name"`
	DOB  time.Time   `edgedb:"dob"`
}

func main() {
	fmt.Println("program started")
	opts := edgedb.Options{Concurrency: 4}
	ctx := context.Background()
	db, err := edgedb.CreateClientDSN(ctx, "edgedb://edgedb@localhost/test", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create a user object type.
	err = db.Execute(ctx, `
        CREATE TYPE User {
            CREATE REQUIRED PROPERTY name -> str;
            CREATE PROPERTY dob -> datetime;
        }
    `)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new user.
	var inserted struct{ id edgedb.UUID }
	err = db.QuerySingle(ctx, `
        INSERT User {
            name := <str>$0,
            dob := <datetime>$1
        }
    `, &inserted, "Bob", time.Date(1984, 3, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		log.Fatal(err)
	}

	// Select users.
	var users []User
	args := map[string]interface{}{"name": "Bob"}
	query := "SELECT User {name, dob} FILTER .name = <str>$name"
	err = db.Query(ctx, query, &users, args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
