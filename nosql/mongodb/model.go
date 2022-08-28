package mongodb

type Menu struct {
	Id          string
	Name        string
	Description string
	Price       int32
	Category    *Category
}

type Category struct {
	Id   string
	Name string
}
