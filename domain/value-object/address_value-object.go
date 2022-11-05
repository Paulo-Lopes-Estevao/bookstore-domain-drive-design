package valueobject

import "github.com/asaskevich/govalidator"

type Address struct {
	Province string
	County   string
	Street   string
	Number   int
	Country  string
}

func NewAddress(province string, county string, street string, number int, country string) (*Address, error) {
	address := &Address{
		Province: province,
		County:   county,
		Street:   street,
		Number:   number,
		Country:  country,
	}

	if err := address.IsValid(); err != nil {
		return nil, err
	}
	return address, nil
}

func (address *Address) IsValid() error {
	_, err := govalidator.ValidateStruct(address)
	if err != nil {
		return err
	}
	return nil
}
