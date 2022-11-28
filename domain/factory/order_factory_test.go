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
	orderItemAggregate := &aggregate.OrderItemAggregate{}

	orderAggregate.CostumerID = "123"

	orderAggregate.Address.Province = "Luanda"
	orderAggregate.Address.County = "Luanda"
	orderAggregate.Address.Street = "Morro Bento ||"
	orderAggregate.Address.Number = 1234
	orderAggregate.Address.Country = "Angola"

	orderItemAggregate.Name = "Harry Potter"
	orderItemAggregate.Description = "Description"
	orderItemAggregate.Quantity = 1
	orderItemAggregate.Price = 7000
	orderItemAggregate.ProductID = uuid.New().String()

	factoryOrder := factory.NewAggregateFactory(orderAggregate)

	addOrderItem, err := factoryOrder.CreateOrderItem(orderItemAggregate)
	assert.Nil(t, err)

	orderAggregate.Item = addOrderItem

	newOder, err := factoryOrder.CreateOrder(&orderAggregate)

	assert.Nil(t, err)

	assert.Equal(t, newOder.CostumerID, "123")
}
