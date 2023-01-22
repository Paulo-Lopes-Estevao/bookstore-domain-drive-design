package factory

import (
	"bookstore/domain/aggregate"
)

type (
	// AggregateFactory is the factory interface for Aggregate
	IOrderAggregateFactory interface {
		CreateOrder(orderFactory *aggregate.Order) (*aggregate.Order, error)
		CreateOrderItem(orderItemFactory *aggregate.OrderItem) (*aggregate.OrderItem, error)
	}

	// AggregateFactoryImpl is the factory implementation for Aggregate
	OrderAggregateFactoryImpl struct {
		aggregateOrder     aggregate.Order
		aggregateOrderItem *aggregate.OrderItem
	}
)

// NewAggregateFactory is the factory constructor for Aggregate
func NewAggregateFactory(aggregateOrder aggregate.Order, aggregateOrderItem *aggregate.OrderItem) IOrderAggregateFactory {
	return &OrderAggregateFactoryImpl{
		aggregateOrder:     aggregateOrder,
		aggregateOrderItem: aggregateOrderItem,
	}

}

func NewAggregateOrderItemFactory(aggregateOrderItem *aggregate.OrderItem) IOrderAggregateFactory {
	return &OrderAggregateFactoryImpl{
		aggregateOrderItem: aggregateOrderItem,
	}

}

// CreateOrder is the factory method for CreateOrder
func (factory *OrderAggregateFactoryImpl) CreateOrder(orderFactory *aggregate.Order) (*aggregate.Order, error) {
	order, err := aggregate.NewOrder(orderFactory)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateOrderItem is the factory method for CreateOrderItem
func (factory *OrderAggregateFactoryImpl) CreateOrderItem(orderItemFactory *aggregate.OrderItem) (*aggregate.OrderItem, error) {
	orderItem, err := aggregate.NewOrderItem(orderItemFactory)
	if err != nil {
		return nil, err
	}
	return &aggregate.OrderItem{
		ID:          orderItem.ID,
		ProductID:   orderItem.ProductID,
		Name:        orderItem.Name,
		Description: orderItem.Description,
		Quantity:    orderItem.Quantity,
		Price:       orderItem.Price,
	}, nil
}
