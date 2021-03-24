package route

import (
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

	r.POST("", h.Create, validation.OrderBodyValidation)
}
