package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/staffrole/handler"
	"github.com/dothiphuc81299/coffeeShop-server/staffrole/validation"
	"github.com/labstack/echo/v4"
)

// InitStaffRoleAdmin ...
func InitStaffRoleAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.StaffRoleAdminHandler{
		StaffRoleAdminService: cs.StaffRole,
	}

	g := e.Group("/staffRole")
	g.POST("", h.Create, middleware.CheckPermissionRoot(d), validation.StaffRoleBodyValidation)

	g.PUT("/:roleID", h.Update, middleware.CheckPermissionRoot(d), h.StaffRoleGetByID, validation.StaffRoleBodyValidation)

	// Get list roles
	g.GET("", h.ListRoleStaff, middleware.CheckPermissionRoot(d))

	// get List permission
	g.GET("/permissions", h.GetListPermission, middleware.CheckPermissionRoot(d))
}
