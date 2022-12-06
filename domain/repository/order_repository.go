package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IOrderRepository interface {
	repository.Repository[aggregate.OrderAggregate]
	CreateOrderItem(orderItemFactory *aggregate.OrderItemAggregate) error
}
