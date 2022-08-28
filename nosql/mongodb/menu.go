package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient *mongo.Client
	Context     context.Context
)

func GetMenus() ([]*Menu, error) {
	query := bson.D{}
	categoryLookUp := bson.E{Key: "$lookup", Value: bson.D{
		{"from", "menu_categories"},
		{"localField", "category"},
		{"foreignField", "_id"},
		{"as", "category"},
	}}
	query = append(query, categoryLookUp)
	result, err := MongoClient.Database("mongo_comparison").Collection("menus").Aggregate(Context, mongo.Pipeline{query})
	if err != nil {
		return nil, err
	}

	menus := []*Menu{}

	for result.Next(Context) {
		menu := Menu{}
		current := result.Current

		menu.Id = current.Lookup("_id").ObjectID().Hex()
		menu.Name = current.Lookup("name").String()
		menu.Price = current.Lookup("price").AsInt32()
		categories := current.Lookup("category").Array()
		if elements, _ := categories.Elements(); len(elements) > 0 {
			menu.Category = &Category{}
			menu.Category.Id = categories.Index(0).Value().Document().Lookup("_id").ObjectID().Hex()
			menu.Category.Name = categories.Index(0).Value().Document().Lookup("name").String()
		}

		menus = append(menus, &menu)
	}
	return menus, nil
}
