package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/user/handler"
	"github.com/labstack/echo/v4"
)

// InitUserAdmin ...
func InitUserAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.UserAdminHandler{
		UserAdminService: cs.User,
	}

	r := e.Group("/users")

	// Change status
	r.PATCH("/:userID/status", h.ChangeStatus, middleware.CheckPermission(config.ModelFieldUser, config.PermissionEdit, d), h.UserGetByID)
	r.GET("/list", h.List, middleware.CheckPermission(config.ModelFieldUser, config.PermissionView, d))
	r.GET("/:userID", h.GetDetailUser, h.UserGetByID)
}
