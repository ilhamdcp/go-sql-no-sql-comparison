package mutation

import (
	"github.com/graphql-go/graphql"
	"go-sql-no-sql-comparison/gql/schema"
	"go-sql-no-sql-comparison/resolver"
)

var MenuMutationFields = graphql.Fields{
	"updateMenu": &graphql.Field{
		Name:        "update menu",
		Description: "Update selected menu",
		Type:        schema.MenuType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "ID of the menu",
			},
			"name": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "New menu name",
			},
			"description": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "New menu description",
			},
			"price": &graphql.ArgumentConfig{
				Type:        graphql.Float,
				Description: "New menu price",
			},
		},
		Resolve: resolver.UpdateMenu,
	},
}
