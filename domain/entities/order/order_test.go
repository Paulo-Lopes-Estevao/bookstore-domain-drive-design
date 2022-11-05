package order_test

import (
	"bookstore/domain/entities/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOderwhenOrderIdIsEmptyWithCustomerId(t *testing.T) {
	_, err := order.NewOrder("123", nil)
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
	orderItem := order.NewOrderItem("Harry Potter", "History", 400, 2)
	orderItem2 := order.NewOrderItem("Harry Potter", "History", 300, 2)

	order2, err := order.NewOrder("123", orderItem, orderItem2)
	assert.Nil(t, err)

	assert.Equal(t, order2.Total(), 1400)

}

func TestOderwhenOrderItemQtyIsLessOrEqualZero0(t *testing.T) {
	orderItem := order.NewOrderItem("Harry Potter", "History", 400, 0)

	_, err := order.NewOrder("123", orderItem)

	if assert.Error(t, err) {
		assert.Equal(t, "OrderItem Quantity cannot be zero", err.Error())
	}

}
