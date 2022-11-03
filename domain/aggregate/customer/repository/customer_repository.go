package repository

import (
	"bookstore/domain/aggregate/customer"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Get(uuid.UUID) (customer.Customer, error)
	Add(customer.Customer) error
	Update(customer.Customer) error
}
