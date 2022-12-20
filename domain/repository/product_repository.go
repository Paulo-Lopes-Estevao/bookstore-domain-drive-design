package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IProductRepository interface {
	repository.IRepository[aggregate.Product]
	CreateCategory(name string) error
}
