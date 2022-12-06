package factory

import (
	"bookstore/domain/aggregate"
)

type (
	// AggregateFactory is the factory interface for Aggregate
	AggregateFactory interface {
		CreateOrder(orderFactory *aggregate.OrderAggregate) (*aggregate.OrderAggregate, error)
		CreateOrderItem(orderItemFactory *aggregate.OrderItemAggregate) (*aggregate.OrderItemAggregate, error)
	}

	// AggregateFactoryImpl is the factory implementation for Aggregate
	AggregateFactoryImpl struct {
		aggregateOrder     aggregate.OrderAggregate
		aggregateOrderItem *aggregate.OrderItemAggregate

	}
)

// NewAggregateFactory is the factory constructor for Aggregate
func NewAggregateFactory(aggregateOrder aggregate.OrderAggregate, aggregateOrderItem *aggregate.OrderItemAggregate) AggregateFactory {
	return &AggregateFactoryImpl{
		aggregateOrder:     aggregateOrder,
		aggregateOrderItem: aggregateOrderItem,
	}

}

func NewAggregateOrderItemFactory(aggregateOrderItem *aggregate.OrderItemAggregate) AggregateFactory {
	return &AggregateFactoryImpl{
		aggregateOrderItem: aggregateOrderItem,
	}

}

// CreateOrder is the factory method for CreateOrder
func (factory *AggregateFactoryImpl) CreateOrder(orderFactory *aggregate.OrderAggregate) (*aggregate.OrderAggregate, error) {
	order, err := aggregate.NewOrder(orderFactory, orderFactory.Item)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateOrderItem is the factory method for CreateOrderItem
func (factory *AggregateFactoryImpl) CreateOrderItem(orderItemFactory *aggregate.OrderItemAggregate) (*aggregate.OrderItemAggregate, error) {
	orderItem, err := aggregate.NewOrderItem(orderItemFactory)
	if err != nil {
		return nil, err
	}
	return &aggregate.OrderItemAggregate{
		ID:          orderItem.ID,
		ProductID:   orderItem.ProductID,
		Name:        orderItem.Name,
		Description: orderItem.Description,
		Quantity:    orderItem.Quantity,
		Price:       orderItem.Price,
	}, nil
}
