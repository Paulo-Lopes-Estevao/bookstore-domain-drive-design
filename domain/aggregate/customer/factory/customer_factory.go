package factory

import (
	"bookstore/domain/aggregate/customer/entities"
	"bookstore/domain/aggregate/customer/event"
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

	customerEvent := event.CustomerCreatedEvent{}
	event.NewCustumerCreatedEvent(customer)
	customerEvent.SendEmailWhenCustomerIsCreatedDispatch()

	return customer, nil
}
