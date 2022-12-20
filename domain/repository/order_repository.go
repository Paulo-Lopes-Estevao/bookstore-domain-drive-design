package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IOrderRepository interface {
	repository.IRepository[aggregate.Order]
	CreateOrderItem(orderItemFactory *aggregate.OrderItem) error
}
