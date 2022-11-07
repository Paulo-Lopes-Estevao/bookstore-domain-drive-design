package repository

import (
	"bookstore/domain/aggregate"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(entity *aggregate.Product) error {
	return r.db.Create(entity).Error
}

func (r *ProductRepository) Update(entity *aggregate.Product) error {
	return r.db.Model(entity).Update(entity).Error
}

func (r *ProductRepository) Delete(id string) error {
	return r.db.Where("id =?", id).Delete(aggregate.Product{}).Error
}

func (r *ProductRepository) Find(id string) (*aggregate.Product, error) {
	var entity aggregate.Product
	if err := r.db.Where("id =?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *ProductRepository) FindAll() ([]aggregate.Product, error) {
	var entities []aggregate.Product
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *ProductRepository) CreateCategory(name string) error {
	return r.db.Model(&aggregate.Category{}).Create(&aggregate.Category{Name: name}).Error
}
