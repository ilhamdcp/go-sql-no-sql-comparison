package query

import (
	"github.com/graphql-go/graphql"
	"go-sql-no-sql-comparison/gql/schema"
	"go-sql-no-sql-comparison/resolver"
)

var MenuQueryFields = graphql.Fields{
	"menus": &graphql.Field{
		Name:        "menus",
		Description: "Get all menus",
		Type:        graphql.NewList(schema.MenuType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID of the menu in the database",
			},
			"keyword": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Search keyword, can be name or description of the menu",
			},
			"category": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Category of the menu",
			},
		},
		Resolve: resolver.GetAllMenus,
	},
}

//var MenuMutationFields = gql.Fields{
//	"update": &gql.Field{
//		Name: "update menu",
//		Description: "Update selected menu",
//		Type: schema.MenuType,
//		Args: gql.FieldConfi
//	}
//}
