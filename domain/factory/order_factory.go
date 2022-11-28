package factory

import (
	"bookstore/domain/aggregate"
)

type (
	// AggregateFactory is the factory interface for Aggregate
	AggregateFactory interface {
		CreateOrder(orderFactory *aggregate.OrderAggregate) (*aggregate.OrderAggregate, error)
		CreateOrderItem(orderItemFactory *aggregate.OrderItemAggregate) (aggregate.OrderItemAggregate, error)
	}

	// AggregateFactoryImpl is the factory implementation for Aggregate
	AggregateFactoryImpl struct {
		//aggregateRepository aggregate.Repository
		aggregateOrder     aggregate.OrderAggregate
		aggregateOrderItem aggregate.OrderItemAggregate
		//aggregateService    aggregate.Service
		//logger              Logger
	}
)

// NewAggregateFactory is the factory constructor for Aggregate
func NewAggregateFactory(aggregateOrder aggregate.OrderAggregate) AggregateFactory {
	return &AggregateFactoryImpl{
		aggregateOrder: aggregateOrder,
	}

}

// CreateOrder is the factory method for CreateOrder
func (factory *AggregateFactoryImpl) CreateOrder(orderFactory *aggregate.OrderAggregate) (*aggregate.OrderAggregate, error) {
	order, err := aggregate.NewOrder(orderFactory, &orderFactory.Item)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateOrderItem is the factory method for CreateOrderItem
func (factory *AggregateFactoryImpl) CreateOrderItem(orderItemFactory *aggregate.OrderItemAggregate) (aggregate.OrderItemAggregate, error) {
	orderItem, err := aggregate.NewOrderItem(orderItemFactory)
	if err != nil {
		return aggregate.OrderItemAggregate{}, err
	}
	return aggregate.OrderItemAggregate{
		ID:          orderItem.ID,
		ProductID:   orderItem.ProductID,
		Name:        orderItem.Name,
		Description: orderItem.Description,
		Quantity:    orderItem.Quantity,
		Price:       orderItem.Price,
	}, nil
}
