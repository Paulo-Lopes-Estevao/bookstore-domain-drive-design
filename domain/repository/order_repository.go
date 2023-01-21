package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IOrderRepository interface {
	repository.IRepository[aggregate.Order]
}

type IOrderItemRepository interface {
	repository.IRepository[aggregate.OrderItem]
}