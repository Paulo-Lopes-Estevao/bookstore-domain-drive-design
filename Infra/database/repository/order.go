package repository

import (
	"bookstore/Infra/database/entities"
	"bookstore/domain/aggregate"

	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(entity *aggregate.OrderAggregate) error {
	return r.db.Create(entity).Error
}

func (r *OrderRepository) Update(entity *aggregate.OrderAggregate) error {
	return r.db.Model(entity).Update(entity).Error
}

func (r *OrderRepository) Find(id string) (*aggregate.OrderAggregate, error) {
	var entity aggregate.OrderAggregate
	if err := r.db.Where("id =?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *OrderRepository) FindAll() ([]aggregate.OrderAggregate, error) {
	var entities []aggregate.OrderAggregate
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *OrderRepository) Delete(id string) error {
	var entities aggregate.OrderAggregate
	return r.db.Where("id =?", id).Delete(entities).Error
}

func (r *OrderRepository) CreateOrderItem(orderItem *aggregate.OrderItemAggregate) error {
	var entity entities.OrderItem
	/* {
		ID:          orderItem.ID,
		ProductID:   orderItem.ProductID,
		Name:        orderItem.Name,
		Description: orderItem.Description,
		Price:       orderItem.Price,
		Quantity:    orderItem.Quantity,
	} */
	return r.db.Create(entity).Error
}
