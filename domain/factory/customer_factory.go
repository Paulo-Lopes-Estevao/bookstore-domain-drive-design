package factory

import (
	entities "bookstore/domain/aggregate/customer"
)

type CustomerFactory struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

func NewCustomer(objCustomer CustomerFactory) (*entities.Customer, error) {
	customer, err := entities.NewCustomerAggregate(objCustomer.Name, objCustomer.Phone, objCustomer.Email, objCustomer.Password)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
