package valueobject

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type PersonValueObject struct {
	Name     string `valid:"required"`
	Phone    string `valid:"required"`
	Email    string `valid:"required,email"`
	Password string `valid:"required,minstringlength(4)"`
}

func NewPersonValueObject(name, phone, email, password string) (*PersonValueObject, error) {
	p := PersonValueObject{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
	}

	err := p.IsValid()
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (personValueObject *PersonValueObject) IsValid() error {
	_, err := govalidator.ValidateStruct(personValueObject)
	if err != nil {
		return err
	}
	return nil
}
