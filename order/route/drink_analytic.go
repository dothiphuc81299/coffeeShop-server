package route

import (
	// "github.com/dothiphuc81299/coffeeShop-server/internal/config"
	// "github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/order/handler"
	"github.com/labstack/echo/v4"
)

// InitDrinkAnalyticAdmin ...
func InitDrinkAnalyticAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.DrinkAnalyticHandler{
		DrinkAnalyticService: cs.DrinkAnalytic,
	}

	r := e.Group("/drink-analytics")

	r.GET("", h.GetList)
// middleware.CheckPermission(config.ModelFieldDrink, config.PermissionView, d)
}
