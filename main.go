package main

import (
	"context"
	"log"

	"github.com/edgedb/edgedb-go"
)

func main() {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var (
		age   int64 = 21
		users []struct {
			ID   edgedb.UUID `edgedb:"id"`
			Name string      `edgedb:"name"`
		}
	)

	query := "SELECT User{name} FILTER .age = <int64>$0"
	err = client.Query(ctx, query, &users, age)

}
