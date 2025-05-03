package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/drink/handler"
	"github.com/dothiphuc81299/coffeeShop-server/drink/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitDrinkAdmin ...
func InitDrinkAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.DrinkAdminHandler{
		DrinkAdminService: cs.Drink,
	}

	g := e.Group("/drink")
	g.POST("", h.Create, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionDelete, d), validation.DrinkBodyValidation)
	g.GET("", h.Search)
	g.PUT("/:drinkID", h.Update, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionEdit, d), validation.DrinkBodyValidation, h.DrinkGetByID)
	g.PATCH("/:drinkID/status", h.ChangeStatus, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionEdit, d), h.DrinkGetByID)
	g.GET("/:drinkID", h.GetDetail, h.DrinkGetByID)
	g.DELETE("/:drinkID", h.DeleteDrink, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionDelete, d), h.DrinkGetByID)
}
