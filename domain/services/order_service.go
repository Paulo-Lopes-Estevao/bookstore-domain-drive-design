package services

import (
	"bookstore/domain/entities/order"
	"errors"
)

func PlaceOrder(customerId string, items ...*order.OrderItem) (*order.Order, error) {
	if len(items) == 0 {
		errors.New("Order must have at least one item")
	}
	order, err := order.NewOrder(customerId, items...)
	if err != nil {
		return nil, err
	}
	return order, nil
}
