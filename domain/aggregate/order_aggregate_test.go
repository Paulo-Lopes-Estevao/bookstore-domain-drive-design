package aggregate_test

import (
	"bookstore/domain/aggregate"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCustomerOrder(t *testing.T) {
	Order := &aggregate.Order{}
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

	ordered, err := aggregate.NewOrder(Order)
	assert.Nil(t, err)
	assert.Equal(t, ordered.CostumerID, "123")

}
