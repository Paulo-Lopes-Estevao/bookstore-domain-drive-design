package repository

import (
	"bookstore/domain/aggregate"

	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(entity *aggregate.Order) error {
	return r.db.Create(entity).Error
}

func (r *OrderRepository) Update(entity *aggregate.Order) error {
	return r.db.Save(entity).Error
}

func (r *OrderRepository) FindByID(id string) (*aggregate.Order, error) {
	var entity aggregate.Order
	if err := r.db.Where("id =?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *OrderRepository) FindAll() ([]*aggregate.Order, error) {
	var order []*aggregate.Order
	if err := r.db.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Delete(id string) error {
	var entities aggregate.Order
	return r.db.Where("id =?", id).Delete(entities).Error
}