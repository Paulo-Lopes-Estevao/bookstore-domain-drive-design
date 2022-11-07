package contract

import (
	"bookstore/Infra/database/repository"
	repoDomain "bookstore/domain/repository"

	"github.com/jinzhu/gorm"
)

type ICustomerContractRepository interface {
	CustomerContractRepository() repoDomain.ICustomerRepository
}

type CustomerContract struct {
	db *gorm.DB
}

func NewCustomerContractRepository(db *gorm.DB) ICustomerContractRepository {
	return &CustomerContract{db: db}
}

func (c *CustomerContract) CustomerContractRepository() repoDomain.ICustomerRepository {
	return repository.NewCustomerRepository(c.db)
}
