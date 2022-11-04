package aggregate

import (
	"bookstore/domain/entities/item"
	valueobject "bookstore/domain/value-object"
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

	itemValueObject, err := valueobject.NewItemValueObject(name, description, price, quantity)
	if err != nil {
		return nil, err
	}

	products := &Product{
		item: &item.Item{
			ID:          uuid.New(),
			Name:        itemValueObject.Name,
			Description: itemValueObject.Description,
		},
		price:    itemValueObject.Price,
		quantity: itemValueObject.Quantity,
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
