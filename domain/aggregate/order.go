package aggregate

import (
	"bookstore/domain/entities/order"
	valueobject "bookstore/domain/value-object"
	"time"

	"github.com/google/uuid"
)

type (
	Order struct {
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
		OrderID     uuid.UUID
		Name        string
		Description string
		Price       float64
		Quantity    int
	}
)

func NewOrderItem(orderItem *OrderItem) (*OrderItem, error) {
	item, errItem := order.NewOrderItem(
		orderItem.Name,
		orderItem.Description,
		orderItem.ProductID,
		orderItem.Price,
		orderItem.Quantity,
		orderItem.OrderID,
	)
	if errItem != nil {
		return nil, errItem
	}
	return &OrderItem{
		ID:          item.ID,
		ProductID:   item.ProductID,
		OrderID:     item.OrderID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    item.Quantity,
	}, nil
}

func NewOrder(orderParam *Order) (*Order, error) {

	err := orderParam.AddAddress()
	if err != nil {
		return nil, err
	}

	orderNew, err := order.NewOrder(orderParam.CostumerID, nil)
	if err != nil {
		return nil, err
	}
	return &Order{
		ID:         orderNew.ID,
		CostumerID: orderNew.CostumerID,
		
	}, nil
}

func (o Order) GetCustomerID() string {
	return o.CostumerID
}

func (OrderItem *OrderItem) OrderedItem() *order.OrderItem {
	item := &order.OrderItem{
		ID:          OrderItem.ID,
		ProductID:   OrderItem.ProductID,
		OrderID:     OrderItem.OrderID,
		Name:        OrderItem.Name,
		Description: OrderItem.Description,
		Price:       OrderItem.Price,
		Quantity:    OrderItem.Quantity,
	}
	return item
}

func (orderAddress *Order) AddAddress() error {
	_, err := valueobject.NewAddress(orderAddress.Address.Province, orderAddress.Address.County, orderAddress.Address.Street, orderAddress.Address.Number, orderAddress.Address.Country)
	if err != nil {
		return err
	}
	return nil
}
