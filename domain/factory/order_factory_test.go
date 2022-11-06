package factory_test

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/entities/order"
	"bookstore/domain/factory"
	valueobject "bookstore/domain/value-object"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	item := &order.OrderItem{
		ID:          uuid.New(),
		Name:        "Harry Potter",
		Description: "History",
		Price:       100,
		Quantity:    2,
	}
	item2 := &order.OrderItem{
		ID:          uuid.New(),
		Name:        "Harry Potter",
		Description: "History",
		Price:       200,
		Quantity:    2,
	}
	value := []*order.OrderItem{
		item, item2,
	}

	address := &valueobject.Address{
		Province: "Luanda",
		County:   "Luanda",
		Street:   "Morro Bento ||",
		Number:   123,
		Country:  "Angola",
	}

	order := factory.OrderFactory{
		Order: aggregate.Order{
			CustomerID: "123",
			OrderItem:  value,
			Address:    address,
		},
	}
	newOder, err := factory.CreateOrder(order)

	assert.Nil(t, err)

	assert.Equal(t, newOder.Order.OrderItem, value)
}
