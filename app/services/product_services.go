package services

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/repository"
)

// type ProductServices struct { }
type ProductServices struct {
	IProductRepository repository.IProductRepository
}

// create a new product service instance with the given product repository instance as a dependency injection
func NewProductServices(productRepository repository.IProductRepository) *ProductServices {
	return &ProductServices{productRepository}
}

// create a new product category
func (s *ProductServices) CreateCategory(name string) error {
	return s.IProductRepository.CreateCategory(name)
}

// create a new product
func (s *ProductServices) CreateProduct(entity *aggregate.Product) error {
	return s.IProductRepository.Create(entity)
}

// get all products
func (s *ProductServices) GetAllProducts() ([]*aggregate.Product, error) {
	return s.IProductRepository.FindAll()
}

// get a product by id
func (s *ProductServices) GetProductByID(id string) (*aggregate.Product, error) {
	return s.IProductRepository.FindByID(id)
}

// update a product
func (s *ProductServices) UpdateProduct(entity *aggregate.Product) error {
	return s.IProductRepository.Update(entity)
}

// delete a product
func (s *ProductServices) DeleteProduct(id string) error {
	return s.IProductRepository.Delete(id)
}

// get all product categories
func (s *ProductServices) GetAllCategories() ([]*aggregate.Category, error) {
	return s.IProductRepository.FindAllCategory()
}

// get a product category by id
func (s *ProductServices) GetCategoryByID(id string) (*aggregate.Category, error) {
	return s.IProductRepository.FindCategoryByID(id)
}

// get a product category by name
func (s *ProductServices) GetCategoryByName(name string) (*aggregate.Category, error) {
	return s.IProductRepository.FindCategoryByName(name)
}

// get product with category name or category id
func (s *ProductServices) GetProductWithCategory(categoryID string) ([]*aggregate.Product, error) {
	return s.IProductRepository.FindProductByCategory(categoryID)
}
