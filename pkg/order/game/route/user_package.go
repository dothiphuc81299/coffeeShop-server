package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/game/handler"
	validation "github.com/dothiphuc81299/coffeeShop-server/game/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

func InitUserPackageApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.UserPackageAppHandler{
		UserPackageAppService: cs.UserPackage,
	}

	g := e.Group(("game/user-packages"))

	g.POST("", h.Create, validation.UserPackageGroupBodyValidation)
	//g.PUT("/:packageGroupID", h.Update, h.PackageGroupGetByID, validation.PackageGroupBodyValidation)

}
