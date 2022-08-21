package gql

import (
	"github.com/graphql-go/graphql"
	"go-sql-no-sql-comparison/gql/mutation"
	"go-sql-no-sql-comparison/gql/query"
)

func InitMutations() graphql.Fields {
	result := map[string]*graphql.Field{}
	queries := []graphql.Fields{mutation.MenuMutationFields}
	for i := range queries {
		for name, field := range queries[i] {
			result[name] = field
		}
	}
	return result
}

func InitQueries() graphql.Fields {
	result := map[string]*graphql.Field{}
	queries := []graphql.Fields{query.MenuQueryFields}
	for i := range queries {
		for name, field := range queries[i] {
			result[name] = field
		}
	}
	return result
}
