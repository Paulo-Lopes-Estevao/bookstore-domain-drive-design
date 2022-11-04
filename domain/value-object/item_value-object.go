package valueobject

import "github.com/asaskevich/govalidator"

type ItemValueObject struct {
	Name        string  `valid:"required"`
	Price       float64 `valid:"required"`
	Quantity    int     `valid:"required"`
	Description string  `valid:"-"`
}

func NewItemValueObject(name, description string, price float64, quantity int) (*ItemValueObject, error) {
	t := ItemValueObject{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	err := t.IsValid()
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (itemValueObject *ItemValueObject) IsValid() error {
	_, err := govalidator.ValidateStruct(itemValueObject)
	if err != nil {
		return err
	}
	return nil
}
