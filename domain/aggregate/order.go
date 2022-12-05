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
		Item       OrderItemAggregate
		CreatedAt  time.Time
		UpdatedAt  time.Time
		Address    valueobject.Address
	}

	OrderItemAggregate struct {
		ID          uuid.UUID
		ProductID   uuid.UUID
		Name        string
		Description string
		Price       float64
		Quantity    int
	}
)

func NewOrderItem(orderItemAggregate *OrderItemAggregate) (*OrderItemAggregate, error) {
	item, errItem := order.NewOrderItem(orderItemAggregate.Name, orderItemAggregate.Description, orderItemAggregate.ProductID, orderItemAggregate.Price, orderItemAggregate.Quantity)
	if errItem != nil {
		return nil, errItem
	}
	return &OrderItemAggregate{
		ID:          item.ID,
		ProductID:   item.ProductID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
	}, nil
}

func NewOrder(orderAggregate *OrderAggregate, orderItemAggregate *OrderItemAggregate) (*OrderAggregate, error) {

	err := orderAggregate.AddAddress()
	if err != nil {
		return nil, err
	}

	item := orderItemAggregate.OrderedItem()

	order, err := order.NewOrder(orderAggregate.CostumerID, item)
	if err != nil {
		return nil, err
	}

	return orderAggregate.OrderInformation(order, item), nil
}

func (o OrderAggregate) GetCustomerID() string {
	return o.CostumerID
}

func (orderItemAggregate *OrderItemAggregate) OrderedItem() *order.OrderItem {
	item := &order.OrderItem{
		ID:          orderItemAggregate.ID,
		ProductID:   orderItemAggregate.ProductID,
		Name:        orderItemAggregate.Name,
		Description: orderItemAggregate.Description,
		Price:       orderItemAggregate.Price,
		Quantity:    orderItemAggregate.Quantity,
	}
	return item
}

func (orderAggregate *OrderAggregate) OrderInformation(order *order.Order, item *order.OrderItem) *OrderAggregate {
	return &OrderAggregate{
		ID:         order.ID,
		CostumerID: order.CostumerID,
		Item: OrderItemAggregate{
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
