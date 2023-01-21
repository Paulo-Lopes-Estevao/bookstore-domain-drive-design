package services

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/repository"
)

// OrderItemServices is a service that handles order item related operations
type OrderItemServices struct {
	IOrderItemRepository repository.IOrderItemRepository
}

// NewOrderItemServices creates a new instance of OrderItemServices with the given order item repository implementation as parameter
func NewOrderItemServices(orderItemRepository repository.IOrderItemRepository) *OrderItemServices {
	return &OrderItemServices{IOrderItemRepository: orderItemRepository}
}

// CreateOrderItem creates a new order item 
func (s *OrderItemServices) CreateOrderItem(orderItem *aggregate.OrderItem) error {
	err := s.IOrderItemRepository.Create(orderItem)
	if err != nil {
		return err
	}
	return nil
}

// FindOrderItemByID finds an order item by its ID
func (s *OrderItemServices) FindOrderItemByID(id string) (*aggregate.OrderItem, error) {
	orderItem, err := s.IOrderItemRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

// FindAllOrderItems finds all order items
func (s *OrderItemServices) FindAllOrderItems() ([]*aggregate.OrderItem, error) {
	orderItems, err := s.IOrderItemRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

// DeleteOrderItem deletes an order item by its ID
func (s *OrderItemServices) DeleteOrderItem(id string) error {
	err := s.IOrderItemRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
