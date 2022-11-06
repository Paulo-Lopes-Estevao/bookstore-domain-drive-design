package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IOrderRepository interface {
	repository.Repository[aggregate.Order]
	CreateOrderItem(productID, name, description, productId string, price float64, quantity int) error
}
