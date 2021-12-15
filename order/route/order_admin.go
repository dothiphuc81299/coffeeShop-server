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

	r.GET("", h.GetListByStatus)
	r.PUT("/:orderID/status", h.ChangeStatus, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID, validation.StatusBodyValidation)
	r.GET("/:orderID", h.GetDetail, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionView, d), h.GetByID)
	r.GET("/statistic", h.GetStatistic, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionView, d))

	r.PUT("/:orderID/success", h.UpdateOrderSuccess, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID)
	r.PUT("/:orderID/cancel", h.CancelOrder, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID)

}
