package item

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
