package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

func (custumer *Customer) IsValid() error {
	_, err := govalidator.ValidateStruct(custumer)
	if err != nil {
		return err
	}
	return nil
}
