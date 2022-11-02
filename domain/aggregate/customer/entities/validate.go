package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

func (custumer *Custumer) IsValid() error {
	_, err := govalidator.ValidateStruct(custumer)
	if err != nil {
		return err
	}
	return nil
}
