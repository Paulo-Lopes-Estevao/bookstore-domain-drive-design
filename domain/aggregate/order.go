package aggregate

import (
	"bookstore/domain/entities/order"
	valueobject "bookstore/domain/value-object"
)

type Order struct {
	CustomerID string
	Address    *valueobject.Address
	Item       []*OrderItem
}

type OrderItem struct {
	ProductID   string
	Name        string
	Description string
	Price       float64
	Quantity    int
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

	order, err := order.NewOrder(orderAggregate.CustomerID, OrderItem...)
	if err != nil {
		return nil, err
	}

	return &Order{
		CustomerID: order.CostumerID,
		Order:      *order,
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
	return o.Order.Item
}

func (orderItem *Order) AddOrderItem() (*order.OrderItem, error) {
	for _, v := range orderItem.Item {
		item, err := order.NewOrderItem(v.Name, v.Description, v.ProductID, v.Price, v.Quantity)
		if err != nil {
			return nil, err
		}
		
	}
}

func (orderAddress *Order) AddAddress() error {
	_, err := valueobject.NewAddress(orderAddress.Address.Province, orderAddress.Address.County, orderAddress.Address.Street, orderAddress.Address.Number, orderAddress.Address.Country)
	if err != nil {
		return err
	}
	return nil
}
