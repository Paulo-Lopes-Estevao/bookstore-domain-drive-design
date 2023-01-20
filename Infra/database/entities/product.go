package entities

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	CategoryID  string
	Name        string
	Price       float64
	Description string
	Base
}

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Name string
	Base
}
