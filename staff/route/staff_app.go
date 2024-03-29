package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/staff/handler"
	"github.com/dothiphuc81299/coffeeShop-server/staff/validation"
	"github.com/labstack/echo/v4"
)

func InitStaffApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.StaffAppHandler{
		StaffService: cs.Staff,
	}
	r := e.Group("/staff")

	// nhan vien cap nhat tai khoan
	r.PUT("/update", h.Update, middleware.CheckStaff(d), validation.StaffUpdateBodyByItValidation)

	// change password do nhan vien
	r.PUT("/me/password", h.UpdatePassword, middleware.CheckStaff(d), validation.StaffChangePasswordBodyValidation)

	// get detail staff
	r.GET("/me", h.GetDetailStaff, middleware.CheckStaff(d))
}
