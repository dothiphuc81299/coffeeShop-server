package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/salary/handler"
	"github.com/labstack/echo/v4"
)

// InitSalaryRoot ...
func InitSalaryRoot(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.SalaryRootHandler{
		SalaryRootService: cs.Salary,
	}

	g := e.Group("/salary")
	g.GET("", h.GetList, middleware.CheckPermissionRoot(d))

}
