package router

import "github.com/labstack/echo/v4"


type RouterConfig interface {
	Router() *echo.Echo
}
