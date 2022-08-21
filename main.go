package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jackc/pgx"
	"go-sql-no-sql-comparison/gql"
	"go-sql-no-sql-comparison/sql/pgxdriver"
	"net/http"
)

func main() {
	conn, _ := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})

	pgxdriver.DbConn = conn

	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "rootQuery",
			Fields: gql.InitQueries(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "rootMutation",
			Fields: gql.InitMutations(),
		}),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/gql", h)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
