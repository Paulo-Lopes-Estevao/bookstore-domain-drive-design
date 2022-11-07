package product

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
