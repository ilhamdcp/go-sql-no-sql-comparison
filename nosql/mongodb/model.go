package mongodb

type Menu struct {
	IdString    string
	Name        string
	Description string
	Price       int32
	Category    *Category
}

type Category struct {
	IdString string
	Name     string
}
