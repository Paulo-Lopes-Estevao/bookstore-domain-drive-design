package entities

type Product struct {
	Base
	CategoryID string
	Name       string
	Price      float64
}

type Category struct {
	Base
	Name string
}
