package pgxdriver

import (
	"github.com/jackc/pgx/pgtype"
)

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       float32
}

type Order struct {
	Id              int64              `json:"id"`
	CustomerName    pgtype.Text        `json:"customer_name"`
	CourierName     pgtype.Text        `json:"courier_name"`
	PaymentTypeId   pgtype.Int4        `json:"payment_type_id"`
	OrderPlatformId pgtype.Int4        `json:"order_platform_id"`
	TotalPrice      pgtype.Float4      `json:"total_price"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
}

type OrderItem struct {
	Item     Menu
	Quantity int
}
