package repository

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Find(id string) (*T, error)
	FindAll(filter *T) ([]T, error)
	Delete(id string) error
}
