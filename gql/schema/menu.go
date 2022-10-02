package schema

import "github.com/graphql-go/graphql"

var MenuType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Menu",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.Int,
		},
		"idString": &graphql.Field{
			Name: "idString",
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Name:        "name",
			Type:        graphql.String,
			Description: "Name of the menu",
		},
		"description": &graphql.Field{
			Name:        "description",
			Type:        graphql.String,
			Description: "Description of the menu (optional)",
		},
		"price": &graphql.Field{
			Name:        "price",
			Type:        graphql.Float,
			Description: "Price of the menu in Rupiah",
		},
		"category": &graphql.Field{
			Name:        "price",
			Type:        MenuCategoryType,
			Description: "Price of the menu in Rupiah",
		},
	},
})

var MenuCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "MenuCategory",
	Description: "Category object of the menu",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.Int,
		},
		"idString": &graphql.Field{
			Name: "idString",
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
		},
	},
})

var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Order",
	Description: "Order object that is created when user confirms the order",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.Int,
		},
		"customerName": &graphql.Field{
			Name:        "customerName",
			Type:        graphql.String,
			Description: "Name of the menu",
		},
		"courierName": &graphql.Field{
			Name:        "courierName",
			Type:        graphql.String,
			Description: "Name of thee courier, when customer orders through apps",
		},
		"totalPrice": &graphql.Field{
			Name:        "totalPrice",
			Type:        graphql.Float,
			Description: "Total price of the order",
		},
		"createdAt": &graphql.Field{
			Name:        "createdAt",
			Type:        graphql.DateTime,
			Description: "Timestamp of when the order was created",
		},
	},
})

var PlatformEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name: "platform",
	Values: graphql.EnumValueConfigMap{
		"POSTGRESQL": &graphql.EnumValueConfig{
			Value: "postgresql",
		},
		"MONGODB": &graphql.EnumValueConfig{
			Value: "mongodb",
		},
	},
})
