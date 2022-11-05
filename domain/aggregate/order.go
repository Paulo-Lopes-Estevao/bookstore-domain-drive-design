package aggregate

import (
	"bookstore/domain/entities/order"
	"bookstore/domain/services"
)

type Order struct {
	customer Customer
	items    []*order.OrderItem
}

func NewOrder(customer Customer, items ...*order.OrderItem) (*Order, error) {
	order, err := services.PlaceOrder(customer.ID.String(), items...)
	if err != nil {
		return nil, err
	}
	return &Order{
		customer: customer,
		items:    order.Item,
	}, nil
}

func (o Order) GetCustomer() Customer {
	return o.customer
}

func (o Order) GetItems() []*order.OrderItem {
	return o.items
}
