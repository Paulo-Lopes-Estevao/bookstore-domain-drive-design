package repository

import (
	"bookstore/domain/aggregate"
	"errors"

	"github.com/google/uuid"
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

func (r *ProductRepository) FindAll() ([]*aggregate.Product, error) {
	var entities []*aggregate.Product
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *ProductRepository) FindByID(id string) (*aggregate.Product, error) {
	var entity *aggregate.Product
	if err := r.db.Where("id =?", id).Find(&entity).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &aggregate.Product{
		ID:          entity.ID,
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
	}, nil
}

func (r *ProductRepository) FindLastProduct() (*aggregate.Product, error) {
	var entity *aggregate.Product
	if err := r.db.Last(&entity).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &aggregate.Product{
		ID:    entity.ID,
		Name:  entity.Name,
		Price: entity.Price,
	}, nil
}

func (r *ProductRepository) CreateCategory(name string) error {
	return r.db.Model(&aggregate.Category{}).Create(&aggregate.Category{ID: uuid.New(), Name: name}).Error
}

func (r *ProductRepository) FindAllCategory() ([]*aggregate.Category, error) {
	var entities []*aggregate.Category
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *ProductRepository) FindCategoryByID(id string) (*aggregate.Category, error) {
	var entity *aggregate.Category
	if err := r.db.Where("id =?", id).Find(&entity).Error; err != nil {
		return nil, errors.New("category not found")
	}
	return &aggregate.Category{
		ID:   entity.ID,
		Name: entity.Name,
	}, nil
}

func (r *ProductRepository) FindCategoryByName(name string) (*aggregate.Category, error) {
	var entity *aggregate.Category
	if err := r.db.Where("name =?", name).Find(&entity).Error; err != nil {
		return nil, errors.New("category not found")
	}
	return &aggregate.Category{
		ID:   entity.ID,
		Name: entity.Name,
	}, nil
}

func (r *ProductRepository) FindProductByCategory(categoryID string) ([]*aggregate.Product, error) {
	var entities []*aggregate.Product
	if err := r.db.Where("category_id =?", categoryID).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}