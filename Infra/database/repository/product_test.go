package repository_test

import (
	"bookstore/Infra/database"
	"bookstore/Infra/database/postgresql"
	"bookstore/Infra/database/repository"
	"bookstore/domain/aggregate"
	"bookstore/domain/factory"
	"log"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
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
	product.Price = 5400
	product.CategoryID = category.ID.String()

	produFactory, errProd := factory.CreatedProduct(product.CategoryID, product.Name, product.Price)
	assert.Nil(t, errProd)

	errProduct := productRepo.Create(produFactory)
	if errProduct != nil {
		t.Errorf("Failed to repository create product: %v", errProduct)
	}

}
