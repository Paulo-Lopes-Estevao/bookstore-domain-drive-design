package interfaces

type IBaseService[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	GetById(id string) (*T, error)
	GetAll() ([]*T, error)
	Remove(id string) error
}
