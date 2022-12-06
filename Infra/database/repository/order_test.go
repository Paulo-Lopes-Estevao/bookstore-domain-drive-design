package repository_test

import (
	"bookstore/Infra/database"
	"bookstore/Infra/database/postgresql"
	"bookstore/Infra/database/repository"
	"bookstore/domain/aggregate"
	"bookstore/domain/factory"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	Env()
}

func TestAddOrderItem(t *testing.T) {
	dsn := postgresql.PostgreSQLConnect()
	db, err := database.DB("postgres", dsn)
	assert.Nil(t, err)

	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	productRepo := repository.NewProductRepository(db)

	product, err := productRepo.FindLastProduct()
	assert.Nil(t, err)

	OrderItem := &aggregate.OrderItem{}

	factoryOrder := factory.NewAggregateOrderItemFactory(OrderItem)

	OrderItem.Name = "Harry Potter"
	OrderItem.Description = "Description"
	OrderItem.Quantity = 1
	OrderItem.Price = 7000
	OrderItem.ProductID = product.ID

	addOrderItem, err := factoryOrder.CreateOrderItem(OrderItem)
	assert.Nil(t, err)

	err = orderRepo.CreateOrderItem(addOrderItem)
	assert.Nil(t, err)
}

func TestCreateOrder(t *testing.T) {
	dsn := postgresql.PostgreSQLConnect()
	db, err := database.DB("postgres", dsn)
	assert.Nil(t, err)

	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	orderAggregate := aggregate.OrderAggregate{}
	OrderItem := aggregate.OrderItem{}

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

	factoryOrder := factory.NewAggregateFactory(orderAggregate, &OrderItem)

	addOrderItem, err := factoryOrder.CreateOrderItem(&OrderItem)
	assert.Nil(t, err)

	orderAggregate.Item = addOrderItem

	newOder, err := factoryOrder.CreateOrder(&orderAggregate)
	assert.Nil(t, err)

	err = orderRepo.CreateOrderItem(addOrderItem)
	assert.Nil(t, err)

	err = orderRepo.Create(newOder)
	assert.Nil(t, err)
}
