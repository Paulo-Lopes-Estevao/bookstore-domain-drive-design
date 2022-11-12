package aggregate

import (
	"bookstore/domain/entities/person"

	"github.com/google/uuid"
)

type Customer struct {
	ID uuid.UUID
	person.Person
}

func NewCustomer(name string, phone string, email string, password string) (*Customer, error) {

	person, err := person.NewPerson(name, phone, email, password)
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		ID:     uuid.New(),
		Person: *person,
	}

	return customer, nil
}
