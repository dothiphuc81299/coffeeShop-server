package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/salary/handler"
	"github.com/labstack/echo/v4"
)

// InitSalaryAdmin ...
func InitSalaryAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.SalaryAdminHandler{
		SalaryAdminService: cs.Salary,
	}

	g := e.Group("/salary")
	g.GET("", h.GetList, middleware.CheckPermissionRoot(d))
	//g.GET("/staffID", h.GetDetail, middleware.CheckPermissionRoot(d), h.GetStaffByID)
}
