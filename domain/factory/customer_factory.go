package factory

import "bookstore/domain/aggregate"

type CustomerFactory struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

func NewCustomer(objCustomer CustomerFactory) (*aggregate.Customer, error) {
	customer, err := aggregate.NewCustomer(objCustomer.Name, objCustomer.Phone, objCustomer.Email, objCustomer.Password)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
