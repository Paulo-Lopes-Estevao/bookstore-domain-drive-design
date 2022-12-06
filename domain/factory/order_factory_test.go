package factory_test

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/factory"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {

	Order := aggregate.Order{}
	OrderItem := &aggregate.OrderItem{}

	Order.CostumerID = "123"

	Order.Address.Province = "Luanda"
	Order.Address.County = "Luanda"
	Order.Address.Street = "Morro Bento ||"
	Order.Address.Number = 1234
	Order.Address.Country = "Angola"

	OrderItem.Name = "Harry Potter"
	OrderItem.Description = "Description"
	OrderItem.Quantity = 1
	OrderItem.Price = 7000
	OrderItem.ProductID = uuid.New()

	factoryOrder := factory.NewAggregateFactory(Order, OrderItem)

	addOrderItem, err := factoryOrder.CreateOrderItem(OrderItem)
	assert.Nil(t, err)

	Order.Item = addOrderItem

	newOder, err := factoryOrder.CreateOrder(&Order)

	assert.Nil(t, err)

	assert.Equal(t, newOder.CostumerID, "123")
}
