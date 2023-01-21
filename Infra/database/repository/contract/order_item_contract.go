package contract

import (
	repoDomain "bookstore/domain/repository"
	"bookstore/infra/database/repository"

	"github.com/jinzhu/gorm"
)

type IOrderItemContractRepository interface {
	OrderItemContractRepository() repoDomain.IOrderItemRepository
}

type OrderItemContractRepository struct {
	db *gorm.DB
}

func NewOrderItemContractRepository(db *gorm.DB) *OrderItemContractRepository {
	return &OrderItemContractRepository{
		db: db,
	}
}

func (r *OrderItemContractRepository) OrderItemContractRepository() repoDomain.IOrderItemRepository {
	return repository.NewOrderItemRepository(r.db)
}
