package product

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	err := category.Validate()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) Validate() error {
	if c.ID == uuid.Nil {
		return errors.New("id must not be empty")
	}
	if c.Name == "" {
		return errors.New("name must not be empty")
	}
	return nil
}
