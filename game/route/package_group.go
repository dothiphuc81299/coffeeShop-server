package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/game/handler"
	validation "github.com/dothiphuc81299/coffeeShop-server/game/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitPackageGroupAdmin
func InitPackageGroupAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.PackageGroupAdminHandler{
		PackageGroupAdminService: cs.PackageGroup,
	}

	g := e.Group("/game/package-groups")
	g.POST("", h.Create, validation.PackageGroupBodyValidation)
	g.PUT("/:packageGroupID", h.Update, h.PackageGroupGetByID, validation.PackageGroupBodyValidation)
}
