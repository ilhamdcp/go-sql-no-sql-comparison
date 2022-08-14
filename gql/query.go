package gql

import (
	"github.com/graphql-go/graphql"
	"go-sql-no-sql-comparison/gql/query"
)

var Queries = graphql.NewObject(graphql.ObjectConfig{
	Name:   "rootQuery",
	Fields: query.MenuQueryFields,
})
