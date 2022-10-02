package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient *mongo.Client
	Context     context.Context
)

func GetMenus(id string, keyword string) ([]*Menu, error) {
	query := bson.D{}
	if id != "" {
		idQuery := bson.E{
			Key: "$match", Value: bson.D{
				{
					Key: "$or", Value: bson.D{
						{
							Key: "name", Value: bson.E{Key: "$regex", Value: fmt.Sprintf("/%v/", keyword)},
						},
						{
							Key: "description", Value: bson.E{Key: "$regex", Value: fmt.Sprintf("/%v/", keyword)},
						},
					},
				},
			},
		}
		query = append(query, idQuery)
	}
	//if keyword != "" {
	//	keywordQuery := bson.E{Key: "keyword", Value: keyword}
	//	query = append(query, keywordQuery)
	//}
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

		menu.IdString = current.Lookup("_id").ObjectID().Hex()
		menu.Name = current.Lookup("name").StringValue()
		menu.Price = current.Lookup("price").AsInt32()
		categories := current.Lookup("category").Array()
		if elements, _ := categories.Elements(); len(elements) > 0 {
			menu.Category = &Category{}
			menu.Category.IdString = categories.Index(0).Value().Document().Lookup("_id").ObjectID().Hex()
			menu.Category.Name = categories.Index(0).Value().Document().Lookup("name").StringValue()
		}

		menus = append(menus, &menu)
	}
	return menus, nil
}
