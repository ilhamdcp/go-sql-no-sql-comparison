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
	rows, err := DbConn.Query("select * from menus where "+
		"($1 <= 0 or id = $1 )and "+
		"(coalesce($2, null) = '<nil>' or name ilike '%' || $2 || '%' or description ilike '%' || $2 || '%')",
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
		err := rows.Scan(&id, &name, &description, &price)
		if err != nil {
			return nil, err
		}
		menus = append(menus, Menu{Id: id, Name: name, Description: description, Price: price})
	}
	return menus, nil
}

func UpdateMenu(id int, name string, description string, price string) (*Menu, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}
	menus, err := GetMenus(id, "")
	if err != nil {
		return nil, err
	}
	if menus == nil || len(menus) == 0 {
		return nil, errors.New("menu not found")
	}

	return nil, nil
}

func GetOrderById(db *pgx.Conn, orderId int) (*Order, error) {
	row := db.QueryRow("SELECT * FROM orders where id = $1", orderId)
	//var (
	//	id            pgtype.Int8
	//	customerName  pgtype.Varchar
	//	courierName   string
	//	paymentType   constants.PaymentType
	//	orderPlatform constants.OrderPlatform
	//	totalPrice    float32
	//	createdAt     time.Time
	//)
	var order Order
	err := row.Scan(&order.Id, &order.CustomerName, &order.CourierName, &order.PaymentTypeId, &order.OrderPlatformId,
		&order.TotalPrice, &order.CreatedAt)
	if err != nil && err == pgx.ErrNoRows {
		return nil, err
	}

	return &order, nil
}
