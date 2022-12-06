package factory

import (
	"bookstore/domain/aggregate"
)

type (
	// AggregateFactory is the factory interface for Aggregate
	AggregateFactory interface {
		CreateOrder(orderFactory *aggregate.Order) (*aggregate.Order, error)
		CreateOrderItem(orderItemFactory *aggregate.OrderItem) (*aggregate.OrderItem, error)
	}

	// AggregateFactoryImpl is the factory implementation for Aggregate
	AggregateFactoryImpl struct {
		aggregateOrder     aggregate.Order
		aggregateOrderItem *aggregate.OrderItem
	}
)

// NewAggregateFactory is the factory constructor for Aggregate
func NewAggregateFactory(aggregateOrder aggregate.Order, aggregateOrderItem *aggregate.OrderItem) AggregateFactory {
	return &AggregateFactoryImpl{
		aggregateOrder:     aggregateOrder,
		aggregateOrderItem: aggregateOrderItem,
	}

}

func NewAggregateOrderItemFactory(aggregateOrderItem *aggregate.OrderItem) AggregateFactory {
	return &AggregateFactoryImpl{
		aggregateOrderItem: aggregateOrderItem,
	}

}

// CreateOrder is the factory method for CreateOrder
func (factory *AggregateFactoryImpl) CreateOrder(orderFactory *aggregate.Order) (*aggregate.Order, error) {
	order, err := aggregate.NewOrder(orderFactory, orderFactory.Item)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateOrderItem is the factory method for CreateOrderItem
func (factory *AggregateFactoryImpl) CreateOrderItem(orderItemFactory *aggregate.OrderItem) (*aggregate.OrderItem, error) {
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
