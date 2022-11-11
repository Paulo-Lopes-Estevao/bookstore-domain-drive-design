package factory

import (
	"bookstore/domain/aggregate"
)

func CreatedProduct(categoryID, name string, price float64) (*aggregate.Product, error) {
	product, err := aggregate.NewProduct(categoryID, name, price)
	if err != nil {
		return nil, err
	}
	return product, nil
}
