package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/game/handler"
	validation "github.com/dothiphuc81299/coffeeShop-server/game/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitGroupAdmin ...
func InitGroupAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.GroupAdminHandler{
		GroupAdminService: cs.Group,
	}

	g := e.Group("/game/groups")
	g.POST("", h.Create, validation.GroupBodyValidation)

	g.PUT("/:groupID", h.Update, h.GroupGetByID, validation.GroupBodyValidation)

	// change status
	g.PATCH("/:groupID", h.ChangeStatus, h.GroupGetByID)
	g.GET("", h.ListAll)
}
