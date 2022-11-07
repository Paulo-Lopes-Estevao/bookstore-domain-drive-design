package repository

import (
	"bookstore/domain/aggregate"

	"github.com/jinzhu/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(entity *aggregate.Customer) error {
	return r.db.Create(entity).Error
}

func (r *CustomerRepository) Update(entity *aggregate.Customer) error {
	return r.db.Save(entity).Error
}

func (r *CustomerRepository) Find(id string) (*aggregate.Customer, error) {
	var customer aggregate.Customer
	if err := r.db.Where("id =?", id).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) FindAll() ([]aggregate.Customer, error) {
	var customers []aggregate.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) Delete(id string) error {
	return r.db.Where("id =?", id).Delete(&aggregate.Customer{}).Error
}
