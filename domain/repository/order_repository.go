package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IOrderRepository interface {
	repository.Repository[aggregate.OrderAggregate]
	CreateOrderItem(productID, name, description, productId string, price float64, quantity int) error
}
