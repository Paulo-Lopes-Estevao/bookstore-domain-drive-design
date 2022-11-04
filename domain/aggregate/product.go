package aggregate

import (
	"bookstore/domain/entities/item"
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	item     *item.Item
	price    float64
	quantity int
}

var (
	ErrMissingValues = errors.New("missing values")
)

func NewProduct(name, description string, price float64, quantity int) (*Product, error) {
	if name == "" || description == "" {
		return nil, ErrMissingValues
	}

	products := &Product{
		item: &item.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: quantity,
	}

	return products, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *item.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
