package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/role/handler"
	"github.com/dothiphuc81299/coffeeShop-server/role/validation"
	"github.com/labstack/echo/v4"
)

// InitRoleAdmin ...
func InitRoleAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.RoleAdminHandler{
		RoleAdminService: cs.Role,
	}

	r := e.Group("/roles")

	// Create
	r.POST("", h.Create, validation.RoleBodyValidation)

	// List
	r.GET("", h.List)

	r.PUT("/:roleID", h.Update,
		h.GetByID, validation.RoleBodyValidation)
}
