package services_test

import (
	"bookstore/domain/entities/order"
	"bookstore/domain/services"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTotalOrderPrice(t *testing.T) {
	orderItem, _ := order.NewOrderItem("Harry Potter", "History", uuid.New(), 400, 3, uuid.New())

	order, err := order.NewOrder("123", orderItem)
	assert.NoError(t, err)

	total := services.TotalOrderPrice(order)
	assert.Equal(t, 1200, total)

}
