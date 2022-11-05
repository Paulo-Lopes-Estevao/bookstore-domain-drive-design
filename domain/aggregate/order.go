package aggregate

import (
	"bookstore/domain/entities/order"
	valueobject "bookstore/domain/value-object"
)

type Order struct {
	customerID Customer
	items      []*order.OrderItem
	Address    valueobject.Address
}

func NewOrder(customer Customer, address valueobject.Address, items ...*order.OrderItem) (*Order, error) {
	order := &Order{
		customerID: customer,
		items:      items,
		Address:    address,
	}
	if err := address.IsValid(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o Order) GetCustomer() Customer {
	return o.customerID
}

func (o Order) GetItems() []*order.OrderItem {
	return o.items
}
