package dependency

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Dependency(db *gorm.DB, echo *echo.Echo) *echo.Echo {
	CustomerDependency(db, echo)
	return echo
}
