package repository_test

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/factory"
	"bookstore/infra/database"
	"bookstore/infra/database/postgresql"
	"bookstore/infra/database/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	Env()
}

func TestCreateCategory(t *testing.T) {

	dsn := postgresql.PostgreSQLConnect()
	db, err := database.DB("postgres", dsn)
	assert.Nil(t, err)

	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	category, categoryErr := aggregate.NewCategory("Crônica")
	assert.Nil(t, categoryErr)

	errCategory := productRepo.CreateCategory(category.Name)
	assert.Nil(t, errCategory)

	product := aggregate.Product{}
	product.Name = "Harry Potter"
	product.Price = 5400.0
	product.CategoryID = category.ID.String()

	produFactory, errProd := factory.CreatedProduct(product.CategoryID, product.Name, product.Price, product.Description)
	assert.Nil(t, errProd)

	errProduct := productRepo.Create(produFactory)
	if errProduct != nil {
		t.Errorf("Failed to repository create product: %v", errProduct)
	}

}
