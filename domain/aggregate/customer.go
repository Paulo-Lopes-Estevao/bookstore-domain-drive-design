package aggregate

import (
	"bookstore/domain/entities/item"
	"bookstore/domain/entities/person"
	"bookstore/domain/entities/person/event"
)

type Customer struct {
	person   *person.Person
	products []*item.Item
}

func NewCustomer(name string, phone string, email string, password string) (*Customer, error) {

	person, err := person.NewPerson(name, phone, email, password)
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		person:   person,
		products: make([]*item.Item, 0),
	}

	customerEvent := event.CustomerCreatedEvent{}
	event.NewCustumerCreatedEvent(email)
	customerEvent.DispatchSendEmailWhenCustomerIsCreatedToConfirmAccount()

	return customer, nil
}
