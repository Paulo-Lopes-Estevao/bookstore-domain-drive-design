package order

import (
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID
	CostumerID string
	Item       []*OrderItem
}

func NewOrder(customerId string, items ...*OrderItem) (*Order, error) {
	order := &Order{
		ID:         uuid.New(),
		CostumerID: customerId,
		Item:       items,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

func (order *Order) Total() int {
	total := 0
	lengthOrderItemTotal := len(order.Item)
	for i := 0; i < lengthOrderItemTotal; i++ {
		price := order.Item[i].orderItemTotal()
		total += price
	}
	return total
}

func (order *Order) Validate() error {
	if len(order.ID) == 0 {
		return errors.New("Order ID cannot be empty")
	}
	if len(order.CostumerID) == 0 {
		return errors.New("Costumer ID cannot be empty")
	}
	if len(order.Item) == 0 {
		return errors.New("Items cannot be empty")
	}
	for _, item := range order.Item {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
