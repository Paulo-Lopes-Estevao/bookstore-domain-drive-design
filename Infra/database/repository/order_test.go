package repository_test

import (
	"bookstore/Infra/database"
	"bookstore/Infra/database/postgresql"
	"bookstore/Infra/database/repository"
	"bookstore/domain/aggregate"
	"bookstore/domain/entities/order"
	"bookstore/domain/factory"
	valueobject "bookstore/domain/value-object"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	Env()
}

func TestCreateOrder(t *testing.T) {
	dsn := postgresql.PostgreSQLConnect()
	db, err := database.DB("postgres", dsn)
	assert.Nil(t, err)

	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	item := &order.OrderItem{
		ID:          uuid.New(),
		ProductID:   "123",
		Name:        "Harry Potter",
		Description: "History",
		Price:       100,
		Quantity:    2,
	}
	item2 := &order.OrderItem{
		ID:          uuid.New(),
		ProductID:   "123",
		Name:        "Harry Potter",
		Description: "History",
		Price:       200,
		Quantity:    2,
	}
	value := []*order.OrderItem{
		item, item2,
	}

	address := &valueobject.Address{
		Province: "Luanda",
		County:   "Luanda",
		Street:   "Morro Bento ||",
		Number:   123,
		Country:  "Angola",
	}

	order := factory.OrderFactory{
		Order: aggregate.Order{
			CustomerID: "123",
			Order: order.Order{
				Item: value,
			},
			Address: address,
		},
	}

	orderAggre, orderErr := factory.CreateOrder(order)

	err := orderRepo.Create(orderAggre)
}
