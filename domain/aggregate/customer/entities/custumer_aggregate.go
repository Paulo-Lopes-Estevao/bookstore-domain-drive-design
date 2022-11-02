package entities

import (
	"bookstore/domain/entities/item"
	"bookstore/domain/entities/person"

	"golang.org/x/crypto/bcrypt"
)

type Custumer struct {
	person   *person.Person
	products []*item.Item
}

func NewCustumer(name string, phone string, email string, password string) (*Custumer, error) {
	person := &person.Person{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
	}
	custumer := &Custumer{
		person:   person,
		products: make([]*item.Item, 0),
	}
	err := custumer.passwordEncrypt()
	if err != nil {
		return nil, err
	}
	return custumer, nil
}

func (custumer *Custumer) passwordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(custumer.person.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	custumer.person.Password = string(password)

	err = custumer.IsValid()
	if err != nil {
		return err
	}

	return nil

}

func (custumer *Custumer) VerifyIsExistEmail() bool {
	return custumer.person.VerifyEmail == true
}

func (custumer *Custumer) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(custumer.person.Password), []byte(password))
	return err == nil
}
