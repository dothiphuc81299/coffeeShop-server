package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/staff/handler"
	"github.com/dothiphuc81299/coffeeShop-server/staff/validation"
	"github.com/labstack/echo/v4"
)

// InitStaffAdmin ...
func InitStaffAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.StaffAdminHandler{
		StaffService: cs.Staff,
	}

	r := e.Group("/staff")
	r.GET("/token", h.GetToken)

	r.GET("", h.ListStaff,
		middleware.RequireLogin,
		middleware.CheckPermissionRoot(d))

	r.POST("", h.Create, middleware.RequireLogin,
		middleware.CheckPermissionRoot(d), validation.StaffBodyValidation)

	r.PUT("/:staffID", h.Update,
		middleware.CheckPermission(config.ModelFieldCategory, config.PermissionView, d),
		h.StaffGetByID, validation.StaffBodyValidation)

	r.PATCH("/:staffID/status", h.ChangeStatus, middleware.RequireLogin,
		middleware.CheckPermissionRoot(d),
		h.StaffGetByID)

	// get detail user
	r.GET("/me", h.GetDetailStaff, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionView, d))

	r.GET("/:staffID/me", h.GetDetailStaffByAdmin)

	r.POST("/log-in", h.StaffLogin, validation.StaffLoginBodyValidation)
}
