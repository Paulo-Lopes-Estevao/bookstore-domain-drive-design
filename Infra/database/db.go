package database

import (
	"bookstore/infra/database/entities"
	"fmt"

	"github.com/jinzhu/gorm"
)

// drive: mysql, sqlite3, postgres
func DB(drive string, dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(drive, dsn)

	if err != nil {

		return nil, err

	}

	db.AutoMigrate(&entities.Customer{}, &entities.OrderItem{}, &entities.Order{}, &entities.Category{}, &entities.Product{})

	fmt.Println("Connection Opened to Database")

	return db, nil

}
