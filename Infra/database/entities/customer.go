package entities

import "github.com/google/uuid"

type Customer struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key"`
	Person
	Base
}

type Person struct {
	Name             string
	Phone            string
	Email            string
	Password         string
	VerifyEmail      bool
	VerificationCode string
}
