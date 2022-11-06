package product

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID    uuid.UUID
	Name  string
	Price float64
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	}

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
