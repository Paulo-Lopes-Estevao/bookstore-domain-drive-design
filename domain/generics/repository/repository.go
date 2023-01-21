package repository

type IRepository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	FindByID(id string) (*T, error)
	FindAll() ([]*T, error)
	Delete(id string) error
}
