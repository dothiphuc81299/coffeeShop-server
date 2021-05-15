package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/event/handler"
	"github.com/dothiphuc81299/coffeeShop-server/event/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitEventAdmin ...
func InitEventAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.EventAdminHandler{
		EventAdminService: cs.Event,
	}

	g := e.Group("/event")
	g.POST("", h.Create, middleware.CheckPermissionRoot(d), validation.EventBodyValidation)

	g.PUT("/:eventID", h.Update, middleware.CheckPermissionRoot(d), h.EventGetByID, validation.EventBodyValidation)

	g.GET("", h.ListAll)
	g.GET("/:eventID", h.GetDetail, h.EventGetByID)

	g.PATCH("/:eventID/status", h.ChangeStatus, middleware.CheckPermissionRoot(d), h.EventGetByID)
}
