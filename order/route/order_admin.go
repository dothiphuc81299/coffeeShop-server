package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/order/handler"
	"github.com/dothiphuc81299/coffeeShop-server/order/validation"
	"github.com/labstack/echo/v4"
)

// InitOrderAdmin ...
func InitOrderAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.OrderAdminHandler{
		OrderAdminService: cs.Order,
	}

	r := e.Group("/orders")

	r.GET("", h.GetListByStatus, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionView, d))

	r.PUT("/:orderID/status", h.ChangeStatus, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID, validation.StatusBodyValidation)

}
