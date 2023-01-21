package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/generics/repository"
)

type IProductRepository interface {
	repository.IRepository[aggregate.Product]
	CreateCategory(name string) error
	FindAllCategory() ([]*aggregate.Category, error)
	FindCategoryByID(id string) (*aggregate.Category, error)
	FindCategoryByName(name string) (*aggregate.Category, error)
	FindProductByCategory(categoryID string) ([]*aggregate.Product, error)
}