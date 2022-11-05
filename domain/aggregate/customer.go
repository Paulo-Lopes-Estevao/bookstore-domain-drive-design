package aggregate

import (
	"bookstore/domain/entities/order"
	"bookstore/domain/entities/person"
	"bookstore/domain/entities/person/event"
	valueobject "bookstore/domain/value-object"

	"github.com/google/uuid"
)

type Customer struct {
	ID       uuid.UUID
	person   *person.Person
	products []*order.OrderItem
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
		ID:       uuid.New(),
		person:   person,
		products: make([]*order.OrderItem, 0),
	}

	customerEvent := event.CustomerCreatedEvent{}
	event.NewCustumerCreatedEvent(email)
	customerEvent.DispatchSendEmailWhenCustomerIsCreatedToConfirmAccount()

	return customer, nil
}
