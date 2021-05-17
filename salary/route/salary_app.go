package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/salary/handler"
	"github.com/labstack/echo/v4"
)

// InitSalaryApp ...
func InitSalaryApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.SalaryAppHandler{
		SalaryAppService: cs.Salary,
	}

	g := e.Group("/salary")
	g.GET("/me", h.GetDetail, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionView, d))
}
