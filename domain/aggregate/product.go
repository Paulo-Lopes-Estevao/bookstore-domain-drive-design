package aggregate

import (
	"fmt"

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

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New(),
		Name: name,
	}
}

func NewProduct(categoryID, name string, price float64) (*Product, error) {
	product := &Product{
		ID:         uuid.New(),
		CategoryID: categoryID,
		Name:       name,
		Price:      price,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.CategoryID == "" {
		return fmt.Errorf("categoryID is empty")
	}
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.Price == 0 {
		return fmt.Errorf("price is required")
	}
	return nil
}
