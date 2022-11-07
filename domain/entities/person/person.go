package person

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	Name             string
	Phone            string `valid:"required"`
	Email            string `valid:"required"`
	Password         string `valid:"required"`
	VerifyEmail      bool   `valid:"-"`
	VerificationCode string `valid:"-"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewPerson(name string, phone string, email string, password string) (*Person, error) {
	person := &Person{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: password,
	}
	person.CreatedAt = time.Now()
	err := person.passwordEncrypt()
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (person *Person) passwordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(person.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	person.Password = string(password)

	return nil

}

func (person *Person) VerifyIsExistEmail() bool {
	return person.VerifyEmail == true
}

func (person *Person) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))
	return err == nil
}
