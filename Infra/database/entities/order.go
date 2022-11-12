package entities

import "github.com/google/uuid"

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	CustomerID string
	OrderItem  []OrderItem
	Address
	Base
}

type OrderItem struct {
	Base
	ProductID   string
	Name        string
	Description string
	Price       float64
	Quantity    int
	OrderID     uuid.UUID
}

type Address struct {
	Province string
	County   string
	Street   string
	Number   int
	Country  string
}
