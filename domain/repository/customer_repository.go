package repository

import (
	"bookstore/domain/aggregate"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
