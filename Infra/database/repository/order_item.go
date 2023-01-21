package repository

import (
	"bookstore/domain/aggregate"

	"github.com/jinzhu/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Create(entity *aggregate.OrderItem) error {
	err := r.db.Create(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderItemRepository) Update(entity *aggregate.OrderItem) error {
	return r.db.Save(entity).Error
}

func (r *OrderItemRepository) FindByID(id string) (*aggregate.OrderItem, error) {
	var entity aggregate.OrderItem
	if err := r.db.Where("id =?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *OrderItemRepository) FindAll() ([]*aggregate.OrderItem, error) {
	var order []*aggregate.OrderItem
	if err := r.db.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderItemRepository) Delete(id string) error {
	var entities aggregate.OrderItem
	return r.db.Where("id =?", id).Delete(entities).Error
}
