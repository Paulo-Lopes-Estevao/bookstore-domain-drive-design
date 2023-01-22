package services

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/factory"
	"bookstore/domain/repository"
)

// OrderItemServices is a service that handles order item related operations
type OrderItemServices struct {
	IOrderItemRepository       repository.IOrderItemRepository
	IOrderItemAggregateFactory factory.IOrderAggregateFactory
}

// NewOrderItemServices creates a new instance of OrderItemServices with the given order item repository implementation as parameter
func NewOrderItemServices(orderItemRepository repository.IOrderItemRepository, orderItemAggregateFactory factory.IOrderAggregateFactory) *OrderItemServices {
	return &OrderItemServices{
		IOrderItemRepository:       orderItemRepository,
		IOrderItemAggregateFactory: orderItemAggregateFactory,
	}
}

// CreateOrderItem creates a new order item
func (s *OrderItemServices) CreateOrderItem(orderItem *aggregate.OrderItem) error {
	orderItemFactory, err := s.IOrderItemAggregateFactory.CreateOrderItem(orderItem)
	if err != nil {
		return err
	}

	errOrderRepository := s.IOrderItemRepository.Create(orderItemFactory)
	if errOrderRepository != nil {
		return errOrderRepository
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
