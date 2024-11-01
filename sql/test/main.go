package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	pg "github.com/infraregistry/schema/pg"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=127.0.0.1 port=15432 user=postgres password=postgres dbname=test sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := pg.New(conn)

	// list all authors
	authors, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateUser(ctx, pg.CreateUserParams{
		Name: "Brian Kernighan",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetUser(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
