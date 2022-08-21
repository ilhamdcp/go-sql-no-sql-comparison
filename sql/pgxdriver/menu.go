package pgxdriver

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
)

var (
	DbConn *pgx.Conn
)

func insertOne(db *pgx.Conn, menu *Menu) (*Menu, error) {
	_, err := db.Exec("INSERT INTO menus(name, description, price) VALUES($1, $2, $3)", menu.Name, menu.Description, menu.Price)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert menu")
	}
	return menu, nil
}

func GetMenus(id int, keyword string) ([]Menu, error) {
	rows, err := DbConn.Query("select m.id as id, "+
		"m.name as name, "+
		"m.description as description, "+
		"m.price as price, "+
		"mc.id as menuCategoryId, "+
		"mc.name as menuCategoryName "+
		"from menus m join menu_categories mc on m.category_id = mc.id "+
		"where ($1 <= 0 or m.id = $1) and "+
		"(coalesce($2, null) = '<nil>' or m.name ilike '%' || $2 || '%' or m.description ilike '%' || $2 || '%')",
		id, keyword)
	if err != nil {
		return nil, err
	}
	var menus []Menu
	for rows.Next() {
		var id int
		var name string
		var description string
		var price float32
		var menuCategoryId int
		var menuCategoryName string
		err := rows.Scan(&id, &name, &description, &price, &menuCategoryId, &menuCategoryName)
		if err != nil {
			return nil, err
		}
		menus = append(menus, Menu{Id: id, Name: name, Description: description, Price: price, Category: &Category{Id: menuCategoryId, Name: menuCategoryName}})
	}
	return menus, nil
}

func UpdateMenu(id int, name string, description string, price float32) (*Menu, error) {
	menus, err := GetMenus(id, "")
	if err != nil {
		return nil, err
	}
	if menus == nil || len(menus) == 0 {
		return nil, errors.New("menu not found")
	}
	updatedName := menus[0].Name
	updatedDescription := menus[0].Description
	updatedPrice := menus[0].Price
	if name != "" {
		updatedName = name
	}
	if description != "" {
		updatedDescription = description
	}
	if price > 0 {
		updatedPrice = price
	}
	query, err := DbConn.Exec("update menus set name = $1, description = $2, price = $3 where id = $4", updatedName, updatedDescription, updatedPrice, id)
	if err != nil {
		return nil, err
	}
	if query.RowsAffected() > 0 {
		updatedMenu, _ := GetMenus(id, "")
		return &updatedMenu[0], nil
	}
	return nil, nil
}
