package resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"go-sql-no-sql-comparison/nosql/mongodb"
	"go-sql-no-sql-comparison/sql/pgxdriver"
)

func GetAllMenus(p graphql.ResolveParams) (interface{}, error) {
	platform, platformOk := p.Args["platform"].(string)
	determinedPlatform := determinePlatform(platform, platformOk)
	if determinedPlatform == "mongodb" {
		id, ok := p.Args["id"].(string)
		if !ok {
			id = ""
		}
		menus, err := mongodb.GetMenus(id, fmt.Sprintf("%v", p.Args["keyword"]))
		if err != nil {
			return graphql.List{}, err
		}
		return menus, err
	} else {
		id, ok := p.Args["id"].(int)
		if !ok {
			id = 0
		}
		menus, err := pgxdriver.GetMenus(id, fmt.Sprintf("%v", p.Args["keyword"]))
		if err != nil {
			return graphql.List{}, err
		}
		return menus, err
	}
}

func determinePlatform(platform string, ok bool) interface{} {
	result := "postgresql"
	if ok {
		result = platform
	}
	return result
}

func UpdateMenu(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	if id == 0 {
		return nil, errors.New("Invalid ID")
	}
	name, nameOk := p.Args["name"].(string)
	if !nameOk {
		name = ""
	}

	description, descriptionOk := p.Args["description"].(string)
	if !descriptionOk {
		description = ""
	}

	price, priceOk := p.Args["price"].(float64)
	if !priceOk {
		price = -1
	}
	menu, err := pgxdriver.UpdateMenu(id, name, description, float32(price))
	if err != nil {
		return graphql.List{}, err
	}
	return menu, err
}
