package repository

import (
	"bookstore/domain/entities/product"
	"bookstore/domain/generics/repository"
)

type IProductRepository interface {
	repository.Repository[product.Product]
}
