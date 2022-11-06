package factory

import (
	"bookstore/domain/aggregate"
)

type OrderFactory struct {
	Order aggregate.Order
}

func CreateOrder(orderFactory OrderFactory) (*OrderFactory, error) {
	order, err := aggregate.NewOrder(orderFactory.Order)
	if err != nil {
		return nil, err
	}

	return &OrderFactory{
		Order: *order,
	}, nil
}
