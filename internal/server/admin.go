package server

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo"
)

// StartAdmin ...
func StartAdmin(service *model.AdminService, d *model.CommonDAO) *echo.Echo {
	e := echo.New()
	startAdminHandler(e, service, d)
	return e

}

func startAdminHandler(e *echo.Echo, service *model.AdminService, d *model.CommonDAO) {
	// drink

	drinkroute.InitDrinkAdmin(e, service, d)
}
