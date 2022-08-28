package main

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jackc/pgx"
	"go-sql-no-sql-comparison/gql"
	"go-sql-no-sql-comparison/nosql/mongodb"
	"go-sql-no-sql-comparison/sql/pgxdriver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
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

	uri := "mongodb://mongouser:mongopassword@localhost:27017/"

	client, errMongo := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if errMongo != nil {
		panic(errMongo)
		return
	}
	mongodb.MongoClient = client

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodb.Context = context

	mongodb.GetMenus()

	var schema, errGraphql = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "rootQuery",
			Fields: gql.InitQueries(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "rootMutation",
			Fields: gql.InitMutations(),
		}),
	})
	if errGraphql != nil {
		fmt.Println(errGraphql)
		return
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/gql", h)
	errGraphql = http.ListenAndServe(":8000", nil)
	if errGraphql != nil {
		return
	}

	defer disconnectMongo()
}

func disconnectMongo() {
	if err := mongodb.MongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
