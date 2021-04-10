package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/shift/handler"
	"github.com/dothiphuc81299/coffeeShop-server/shift/validation"
	"github.com/labstack/echo/v4"
)

// InitShiftAdmin ...
func InitShiftAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.ShiftAdminHandler{
		ShiftAdminService: cs.Shift,
	}

	g := e.Group("/shift")
	g.POST("", h.Create, middleware.CheckPermission(config.ModelFieldShift, config.PermissionEdit, d), validation.ShiftBodyValidation)

	g.PUT("/:shiftID", h.Update, middleware.CheckPermission(config.ModelFieldShift, config.PermissionEdit, d), h.ShiftGetByID, validation.ShiftBodyValidation)
	
	g.GET("", h.ListAll, middleware.CheckPermission(config.ModelFieldShift, config.PermissionView, d))

	g.PATCH("/:shiftID/status", h.AcceptShiftByAdmin, h.ShiftGetByID, middleware.CheckPermissionRoot(d))
}
