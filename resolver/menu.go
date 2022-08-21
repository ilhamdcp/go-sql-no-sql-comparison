package resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"go-sql-no-sql-comparison/sql/pgxdriver"
)

func GetAllMenus(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		id = 0
	}
	menus, err := pgxdriver.GetMenus(id, fmt.Sprintf("%v", p.Args["keyword"]))
	if err != nil {
		return graphql.List{}, err
	}
	fmt.Println(menus)
	return menus, err
}

func UpdateMenu(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	if id == 0 {
		return nil, errors.New("Invalid ID")
	}
	menu, err := pgxdriver.UpdateMenu(id, p.Args["name"].(string), p.Args["description"].(string), p.Args["price"].(float32))
	if err != nil {
		return graphql.List{}, err
	}
	return menu, err
}
