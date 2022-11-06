package contract

import (
	"bookstore/Infra/repository"
	repoDomain "bookstore/domain/repository"

	"github.com/jinzhu/gorm"
)

type IOrderContractRepository interface {
	OrderContractRepository() repoDomain.IOrderRepository
}

type OrderContractRepository struct {
	db *gorm.DB
}

func NewOrderContractRepository(db *gorm.DB) *OrderContractRepository {
	return &OrderContractRepository{
		db: db,
	}
}

func (r *OrderContractRepository) OrderContractRepository() repoDomain.IOrderRepository {
	return repository.NewOrderRepository(r.db)
}
