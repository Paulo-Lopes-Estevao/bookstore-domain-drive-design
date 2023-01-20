package product

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID
	CategoryID string
	Name       string
	Price      float64
	Description string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewProduct(id uuid.UUID, categoryID, name string, price float64, description string) (*Product, error) {
	product := &Product{
		ID:         id,
		CategoryID: categoryID,
		Name:       name,
		Price:      price,
		Description: description,
	}
	product.CreatedAt = time.Now()
	if err := product.Validate(); err != nil {
		return nil, err

	}

	return product, nil

}

func (p *Product) Validate() error {

	if p.ID == uuid.Nil {
		return errors.New("Product ID cannot be empty")
	}
	if p.Name == "" {
		return errors.New("Product Name cannot be empty")
	}
	if p.Price == 0 {
		return errors.New("Product Price cannot be zero")
	}
	return nil
}

func (p *Product) GetID() uuid.UUID {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
