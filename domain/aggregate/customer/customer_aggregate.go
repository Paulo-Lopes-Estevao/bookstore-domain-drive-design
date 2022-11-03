package entities

import (
	"bookstore/domain/entities/item"
	"bookstore/domain/entities/person"

	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	person   *person.Person
	products []*item.Item
}

func NewCustomerAggregate(name string, phone string, email string, password string) (*Customer, error) {
	person := &person.Person{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
	}
	customer := &Customer{
		person:   person,
		products: make([]*item.Item, 0),
	}
	err := customer.passwordEncrypt()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (customer *Customer) passwordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(customer.person.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	customer.person.Password = string(password)

	err = customer.IsValid()
	if err != nil {
		return err
	}

	return nil

}

func (customer *Customer) VerifyIsExistEmail() bool {
	return customer.person.VerifyEmail == true
}

func (customer *Customer) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(customer.person.Password), []byte(password))
	return err == nil
}
