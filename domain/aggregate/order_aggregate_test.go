package aggregate_test

import (
	"bookstore/domain/aggregate"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCustomerOrder(t *testing.T) {
	orderAggregate := &aggregate.OrderAggregate{}
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
	orderItemAggregate.ProductID = uuid.New()

	orderItem, err := aggregate.NewOrderItem(orderItemAggregate)
	assert.Nil(t, err)

	ordered, err := aggregate.NewOrder(orderAggregate, orderItem)
	assert.Nil(t, err)
	assert.Equal(t, ordered.CostumerID, "123")

}
