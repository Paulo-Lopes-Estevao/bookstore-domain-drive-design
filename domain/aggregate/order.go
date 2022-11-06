package aggregate

import (
	"bookstore/domain/entities/order"
	valueobject "bookstore/domain/value-object"
)

type Order struct {
	CustomerID string
	OrderItem  []*order.OrderItem
	Address    *valueobject.Address
}

func NewOrder(orderAggregate Order) (*Order, error) {

	err := orderAggregate.AddAddress()
	if err != nil {
		return nil, err
	}

	err = orderAggregate.AddOrderItem()
	if err != nil {
		return nil, err
	}

	order, err := order.NewOrder(orderAggregate.CustomerID, orderAggregate.OrderItem...)
	if err != nil {
		return nil, err
	}

	return &Order{
		CustomerID: order.CostumerID,
		OrderItem:  orderAggregate.OrderItem,
		Address:    orderAggregate.Address,
	}, nil
}

func (o Order) GetAddress() *valueobject.Address {
	return o.Address
}

func (o Order) GetCustomerID() string {
	return o.CustomerID
}

func (o Order) GetOrderItem() []*order.OrderItem {
	return o.OrderItem
}

func (orderItem *Order) AddOrderItem() error {
	for _, v := range orderItem.OrderItem {
		_, err := order.NewOrderItem(v.Name, v.Description, v.Price, v.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (orderAddress *Order) AddAddress() error {
	_, err := valueobject.NewAddress(orderAddress.Address.Province, orderAddress.Address.County, orderAddress.Address.Street, orderAddress.Address.Number, orderAddress.Address.Country)
	if err != nil {
		return err
	}
	return nil
}
