package controller

import (
	"bookstore/app/services"
	"bookstore/domain/aggregate"

	"github.com/labstack/echo/v4"
)


type ICustomerController interface{
	CreateCustomer(ctx echo.Context) error
	UpdateCustomer(ctx echo.Context) error
	GetCustomerById(ctx echo.Context) error
	GetAllCustomer(ctx echo.Context) error
	DeleteCustomer(ctx echo.Context) error
}

type customerController struct {
	customerService services.ICustomerService
}

func NewCustomerController(customerService services.ICustomerService) ICustomerController {
	return &customerController{
		customerService: customerService,
	}
}

func (c *customerController) CreateCustomer(ctx echo.Context) error {
	var customer aggregate.Customer
	if err := ctx.Bind(&customer); err != nil {
		ctx.JSON(400, err)
	}
	if err := c.customerService.Create(&customer); err != nil {
		ctx.JSON(400, err)
	}
	return ctx.JSON(201, customer)
}

func (c *customerController) UpdateCustomer(ctx echo.Context) error {
	panic("implement me")
}

func (c *customerController) GetCustomerById(ctx echo.Context) error {
	panic("implement me")
}

func (c *customerController) GetAllCustomer(ctx echo.Context) error {
	panic("implement me")
}

func (c *customerController) DeleteCustomer(ctx echo.Context) error {
	panic("implement me")
}