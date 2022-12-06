package aggregate_test

import (
	"bookstore/domain/aggregate"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCustomerOrder(t *testing.T) {
	orderAggregate := &aggregate.OrderAggregate{}
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

	orderItem, err := aggregate.NewOrderItem(OrderItem)
	assert.Nil(t, err)

	ordered, err := aggregate.NewOrder(orderAggregate, orderItem)
	assert.Nil(t, err)
	assert.Equal(t, ordered.CostumerID, "123")

}
