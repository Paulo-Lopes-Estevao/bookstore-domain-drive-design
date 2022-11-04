package aggregate

import (
	"bookstore/domain/entities/item"
	"bookstore/domain/entities/person"
	"bookstore/domain/entities/person/event"
	valueobject "bookstore/domain/value-object"
)

type Customer struct {
	person   *person.Person
	products []*item.Item
}

func NewCustomer(name string, phone string, email string, password string) (*Customer, error) {

	personValueObject, err := valueobject.NewPersonValueObject(name, phone, email, password)
	if err != nil {
		return nil, err
	}

	person, err := person.NewPerson(personValueObject.Name, personValueObject.Phone, personValueObject.Email, personValueObject.Password)
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
