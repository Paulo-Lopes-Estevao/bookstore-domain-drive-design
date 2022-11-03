package person

import (
	"github.com/asaskevich/govalidator"
)

func (person *Person) IsValid() error {
	_, err := govalidator.ValidateStruct(person)
	if err != nil {
		return err
	}
	return nil
}
