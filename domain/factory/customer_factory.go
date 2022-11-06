package factory

import "bookstore/domain/aggregate"

type CustomerFactory struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

func NewCustomer(Customer CustomerFactory) (*aggregate.Customer, error) {
	customer, err := aggregate.NewCustomer(Customer.Name, Customer.Phone, Customer.Email, Customer.Password)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
