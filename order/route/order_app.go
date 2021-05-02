package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/order/handler"
	"github.com/dothiphuc81299/coffeeShop-server/order/validation"
	"github.com/labstack/echo/v4"
)

// InitOrderApp ...
func InitOrderApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.OrderAppHandler{
		OrderAppService: cs.Order,
	}

	r := e.Group("/orders")

	r.POST("", h.Create, middleware.CheckUser(d), validation.OrderBodyValidation)

	// xem chi tieet don hang

	r.GET("/:orderID", h.GetDetail, middleware.CheckUser(d), h.GetByID)

	r.GET("/me", h.GetList, middleware.CheckUser(d))
}
