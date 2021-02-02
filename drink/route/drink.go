package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/drink/handler"
	"github.com/dothiphuc81299/coffeeShop-server/drink/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitDrinkAdmin ...
func InitDrinkAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.DrinkAdminHandler{
		DrinkAdminService: cs.Drink,
	}

	g := e.Group("/drink")
	g.POST("", h.Create, validation.DrinkBodyValidation)
}
