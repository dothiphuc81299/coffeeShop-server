package route

import (
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

	// only root
	r := e.Group("/staff")
	r.POST("", h.Create,  middleware.CheckPermissionRoot(d), validation.StaffBodyValidation)
	r.GET("", h.ListStaff,middleware.CheckPermissionRoot(d))
	r.GET("/:staffID", h.GetStaffByID,middleware.CheckPermissionRoot(d))
	r.GET("/token", h.GetToken)
	r.PATCH("/:staffID/status", h.ChangeStatus, middleware.CheckPermissionRoot(d), h.StaffGetByID)
	r.PUT("/:staffID", h.UpdateRole, middleware.CheckPermissionRoot(d), h.StaffGetByID, validation.StaffUpdateRoleBodyValidation)
	
	r.POST("/log-in", h.StaffLogin, validation.StaffLoginBodyValidation)

}
