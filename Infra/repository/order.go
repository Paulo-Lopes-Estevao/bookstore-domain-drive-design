package repository

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/entities/order"

	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(entity *aggregate.Order) error {
	return r.db.Create(entity.Order).Error
}

func (r *OrderRepository) Update(entity *aggregate.Order) error {
	return r.db.Model(entity).Update(entity.Order).Error
}

func (r *OrderRepository) Find(id string) (*aggregate.Order, error) {
	var entity aggregate.Order
	if err := r.db.Where("id =?", id).Find(&entity.Order).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *OrderRepository) FindAll(filter *aggregate.Order) ([]aggregate.Order, error) {
	var entities []aggregate.Order
	if err := r.db.Where(filter).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *OrderRepository) Delete(id string) error {
	var entities aggregate.Order
	return r.db.Where("id =?", id).Delete(entities.Order).Error
}

func (r *OrderRepository) CreateOrderItem(productID, name, description, productId string, price float64, quantity int) error {
	return r.db.Create(order.OrderItem{
		ProductID:   productID,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}).Error
}
