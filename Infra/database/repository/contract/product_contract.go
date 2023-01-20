package contract

import (
	repoDomain "bookstore/domain/repository"
	"bookstore/infra/database/repository"

	"github.com/jinzhu/gorm"
)

type IProductContractRepository interface {
	ProductContractRepository() repoDomain.IProductRepository
}

type ProductContractRepository struct {
	db *gorm.DB
}

func NewProductContractRepository(db *gorm.DB) *ProductContractRepository {
	return &ProductContractRepository{
		db: db,
	}
}

func (r *ProductContractRepository) ProductContractRepository() repoDomain.IProductRepository {
	return repository.NewProductRepository(r.db)
}
