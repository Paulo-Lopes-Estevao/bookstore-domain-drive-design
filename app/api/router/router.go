package router

import (
	"bookstore/app/api/controller"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo                *echo.Echo
	ICustomerController controller.ICustomerController
}

func NewRouter(echo *echo.Echo, customerController controller.ICustomerController) *Router {
	return &Router{
		Echo:                echo,
		ICustomerController: customerController,
	}
}

func (r *Router) Router() *echo.Echo {
	v1 := r.Echo.Group("/api/v1")
	{
		customer := v1.Group("/customer")
		{
			customer.POST("", r.ICustomerController.CreateCustomer)
		}
	}

	return r.Echo
}
