package repository_test

import (
	"bookstore/Infra/database"
	"bookstore/Infra/database/postgresql"
	"bookstore/Infra/database/repository"
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

func TestCreateCustomer(t *testing.T) {
	dsn := postgresql.PostgreSQLConnect()
	db, err := database.DB("postgres", dsn)
	assert.Nil(t, err)

	defer db.Close()

	customerRepo := repository.NewCustomerRepository(db)
	customerFactory := factory.CustomerFactory{}

	customerFactory.Name = "Paulo Estevão"
	customerFactory.Email = "pl1745240@gmail.com"
	customerFactory.Password = "8hxs27swjaoq"
	customerFactory.Phone = "993836542"

	customer, errFact := factory.NewCustomer(customerFactory)
	if errFact != nil {
		t.Errorf("Failed to create customer: %v", errFact)
	}

	errRepo := customerRepo.Create(customer)
	if errRepo != nil {
		t.Errorf("Failed to repository create customer: %v", errRepo)
	}
}
