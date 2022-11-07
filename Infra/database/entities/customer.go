package entities

type Customer struct {
	Base
	Person
}

type Person struct {
	Name             string
	Phone            string
	Email            string
	Password         string
	VerifyEmail      bool
	VerificationCode string
}
