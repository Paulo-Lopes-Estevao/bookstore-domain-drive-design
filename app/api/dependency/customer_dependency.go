package dependency

import (
	"bookstore/Infra/database/repository"
	"bookstore/app/api/controller"
	"bookstore/app/api/router"
	"bookstore/app/services"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func CustomerDependency(db *gorm.DB, echo *echo.Echo) *echo.Echo {
	customerRepository := repository.NewCustomerRepository(db)
	customerService := services.NewCustomerServices(customerRepository)
	customerController := controller.NewCustomerController(customerService)
	route := router.NewRouter(echo, customerController)
	return route.Router()
}
