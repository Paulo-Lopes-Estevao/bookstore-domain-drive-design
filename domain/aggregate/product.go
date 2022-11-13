package aggregate

import (
	"bookstore/domain/entities/product"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID
	CategoryID string
	Name       string
	Price      float64
}

type Category struct {
	ID   uuid.UUID
	Name string
}

func NewCategory(name string) (*Category, error) {
	category, err := product.NewCategory(name)
	if err != nil {
		return nil, err
	}
	return &Category{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func NewProduct(categoryID, name string, price float64) (*Product, error) {
	product, err := product.NewProduct(uuid.New(), categoryID, name, price)
	if err != nil {
		return nil, err
	}
	aggreProduct := &Product{
		ID:         product.ID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Price:      product.Price,
	}
	return aggreProduct, nil
}
