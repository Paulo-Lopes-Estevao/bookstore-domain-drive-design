package aggregate

import (
	"bookstore/domain/entities/order"
	valueobject "bookstore/domain/value-object"
	"time"

	"github.com/google/uuid"
)

type (
	OrderAggregate struct {
		ID         uuid.UUID
		CostumerID string
		Item       *OrderItem
		CreatedAt  time.Time
		UpdatedAt  time.Time
		Address    valueobject.Address
	}

	OrderItem struct {
		ID          uuid.UUID
		ProductID   uuid.UUID
		Name        string
		Description string
		Price       float64
		Quantity    int
	}
)

func NewOrderItem(orderItem *OrderItem) (*OrderItem, error) {
	item, errItem := order.NewOrderItem(orderItem.Name, orderItem.Description, orderItem.ProductID, orderItem.Price, orderItem.Quantity)
	if errItem != nil {
		return nil, errItem
	}
	return &OrderItem{
		ID:          item.ID,
		ProductID:   item.ProductID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
	}, nil
}

func NewOrder(orderAggregate *OrderAggregate, OrderItem *OrderItem) (*OrderAggregate, error) {

	err := orderAggregate.AddAddress()
	if err != nil {
		return nil, err
	}

	item := OrderItem.OrderedItem()

	order, err := order.NewOrder(orderAggregate.CostumerID, item)
	if err != nil {
		return nil, err
	}

	return orderAggregate.OrderInformation(order, item), nil
}

func (o OrderAggregate) GetCustomerID() string {
	return o.CostumerID
}

func (OrderItem *OrderItem) OrderedItem() *order.OrderItem {
	item := &order.OrderItem{
		ID:          OrderItem.ID,
		ProductID:   OrderItem.ProductID,
		Name:        OrderItem.Name,
		Description: OrderItem.Description,
		Price:       OrderItem.Price,
		Quantity:    OrderItem.Quantity,
	}
	return item
}

func (orderAggregate *OrderAggregate) OrderInformation(order *order.Order, item *order.OrderItem) *OrderAggregate {
	return &OrderAggregate{
		ID:         order.ID,
		CostumerID: order.CostumerID,
		Item: &OrderItem{
			ID:          item.ID,
			ProductID:   item.ProductID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		},
		Address: orderAggregate.Address,
	}
}

func (orderAddress *OrderAggregate) AddAddress() error {
	_, err := valueobject.NewAddress(orderAddress.Address.Province, orderAddress.Address.County, orderAddress.Address.Street, orderAddress.Address.Number, orderAddress.Address.Country)
	if err != nil {
		return err
	}
	return nil
}
