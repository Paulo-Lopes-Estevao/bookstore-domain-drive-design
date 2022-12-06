package factory_test

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/factory"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {

	orderAggregate := aggregate.OrderAggregate{}
	OrderItem := &aggregate.OrderItem{}

	orderAggregate.CostumerID = "123"

	orderAggregate.Address.Province = "Luanda"
	orderAggregate.Address.County = "Luanda"
	orderAggregate.Address.Street = "Morro Bento ||"
	orderAggregate.Address.Number = 1234
	orderAggregate.Address.Country = "Angola"

	OrderItem.Name = "Harry Potter"
	OrderItem.Description = "Description"
	OrderItem.Quantity = 1
	OrderItem.Price = 7000
	OrderItem.ProductID = uuid.New()

	factoryOrder := factory.NewAggregateFactory(orderAggregate, OrderItem)

	addOrderItem, err := factoryOrder.CreateOrderItem(OrderItem)
	assert.Nil(t, err)

	orderAggregate.Item = addOrderItem

	newOder, err := factoryOrder.CreateOrder(&orderAggregate)

	assert.Nil(t, err)

	assert.Equal(t, newOder.CostumerID, "123")
}
