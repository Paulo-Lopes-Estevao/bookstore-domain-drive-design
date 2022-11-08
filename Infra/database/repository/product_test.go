package repository_test

import (
	"bookstore/Infra/database"
	"bookstore/Infra/database/postgresql"
	"bookstore/Infra/database/repository"
	"bookstore/domain/aggregate"
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
	product := aggregate.Product{}
	product.Name = "Harry Potter"
	product.Price = 400

	errCategory := productRepo.CreateCategory("Crônica")
	assert.Nil(t, errCategory)

	errProduct := productRepo.Create(&product)
	if errProduct != nil {
		t.Errorf("Failed to repository create product: %v", errProduct)
	}

}
