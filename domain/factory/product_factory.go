package factory

import (
	"bookstore/domain/entities/product"
)

func CreatedProduct(categoryID, name string, price float64) (*product.Product, error) {
	product, err := product.NewProduct(categoryID, name, price)
	if err != nil {
		return nil, err
	}
	return product, nil
}
