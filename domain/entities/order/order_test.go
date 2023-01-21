package order_test

import (
	"bookstore/domain/entities/order"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOderwhenOrderIdIsEmptyWithCustomerId(t *testing.T) {
	_, err := order.NewOrder("", &order.OrderItem{})
	if assert.Error(t, err) {
		assert.Equal(t, "Costumer ID cannot be empty", err.Error())
	}
}

func TestOderwhenOrderItemsIsEmpty(t *testing.T) {
	_, err := order.NewOrder("123", &order.OrderItem{})
	if assert.Error(t, err) {
		assert.Equal(t, "OrderItem ID cannot be empty", err.Error())
	}
}

func TestOderwhenOrderCalculateTotal(t *testing.T) {
	orderItem, _ := order.NewOrderItem("Harry Potter", "History", uuid.New(), 400, 2, uuid.New())
	orderItem2, _ := order.NewOrderItem("Harry Potter", "History", uuid.New(), 300, 2, uuid.New())

	order2, err := order.NewOrder("123", orderItem, orderItem2)
	assert.Nil(t, err)

	assert.Equal(t, order2.Total(), 1400)

}

func TestOderwhenOrderItemQtyIsLessOrEqualZero0(t *testing.T) {
	_, err := order.NewOrderItem("Harry Potter", "History", uuid.New(), 400, 0, uuid.New())

	if assert.Error(t, err) {
		assert.Equal(t, "OrderItem Quantity cannot be zero", err.Error())
	}

}
