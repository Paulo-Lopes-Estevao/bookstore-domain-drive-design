package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID          uuid.UUID
	ProductID   uuid.UUID
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewOrderItem(name, description string, productId uuid.UUID, price float64, quantity int) (*OrderItem, error) {
	orderItem := &OrderItem{
		ID:          uuid.New(),
		ProductID:   productId,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	orderItem.CreatedAt = time.Now()
	if err := orderItem.Validate(); err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (orderItem *OrderItem) orderItemTotal() int {
	return int(orderItem.Price) * orderItem.Quantity
}

func (orderItem *OrderItem) Validate() error {
	if orderItem.ID == uuid.Nil {
		return errors.New("OrderItem ID cannot be empty")
	}
	if orderItem.ProductID == uuid.Nil {
		return errors.New("Product ID cannot be empty")
	}
	if orderItem.Name == "" {
		return errors.New("OrderItem Name cannot be empty")
	}
	if orderItem.Description == "" {
		return errors.New("OrderItem Description cannot be empty")
	}
	if orderItem.Price == 0 {
		return errors.New("OrderItem Price cannot be zero")
	}
	if orderItem.Quantity == 0 {
		return errors.New("OrderItem Quantity cannot be zero")
	}
	return nil
}
