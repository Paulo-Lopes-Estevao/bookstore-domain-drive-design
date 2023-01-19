package services

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/repository"
)

type OrderService struct {
	repository repository.IOrderRepository
}

type IOrders interface {
}

// create NewOrderService with repository as parameter and return pointer to OrderService struct for dependency injection
func NewOrderService(repository repository.IOrderRepository) *OrderService {
	return &OrderService{repository: repository}
}

// create function to get all orders
func (s *OrderService) GetOrders() ([]*aggregate.Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// create function to create order
func (s *OrderService) CreateOrder(entity *aggregate.Order) error {
	if err := s.repository.Create(entity); err != nil {
		return err
	}

	return nil
}

// create function to update order
func (s *OrderService) UpdateOrder(entity *aggregate.Order) error {
	if err := s.repository.Update(entity); err != nil {
		return err
	}

	return nil
}

// create function to delete order
func (s *OrderService) DeleteOrder(entity *aggregate.Order) error {
	if err := s.repository.Delete(entity.ID.String()); err != nil {
		return err
	}

	return nil
}
