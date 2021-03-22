package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/user/handler"
	"github.com/dothiphuc81299/coffeeShop-server/user/validation"
	"github.com/labstack/echo/v4"
)

// InitUserAdmin ...
func InitUserAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.UserAdminHandler{
		UserAdminService: cs.User,
	}

	r := e.Group("/users")

	// Create
	r.POST("", h.Create, validation.UserBodyValidation)

	// List
	r.GET("", h.List)

	// Detail
	r.GET("/:userID", h.Detail, h.GetByID)
	// Update
	r.PUT("/:userID", h.Update,
		h.GetByID, validation.UserBodyValidation)

	// Change status
	r.PATCH("/:userID/status", h.ChangeStatus, h.GetByID)
}
