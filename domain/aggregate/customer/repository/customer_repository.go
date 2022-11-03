package repository

import (
	"bookstore/domain/aggregate/customer/entities"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Get(uuid.UUID) (entities.Customer, error)
	Add(entities.Customer) error
	Update(entities.Customer) error
}
