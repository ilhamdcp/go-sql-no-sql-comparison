package resolver

import (
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
