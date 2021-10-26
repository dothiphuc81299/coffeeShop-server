package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/game/handler"
	validation "github.com/dothiphuc81299/coffeeShop-server/game/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitPackageAdmin
func InitPackageAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.PackageAdminHandler{
		PackageAdminService:      cs.Package,
		PackageGroupAdminService: cs.PackageGroup,
	}

	g := e.Group("/game/packages")
	g.POST("", h.Create, validation.PackageBodyValidation)

	g.PUT("/:packageID", h.Update, h.PackageGetByID, validation.PackageBodyValidation)
	g.GET("", h.ListAll)
	g.GET("/:packageID", h.GetDetail, h.PackageGetByID)
	//g.GET("/:packageID", h.GetPackageGroupByPackageID,)
	g.GET("-groups/:packageID", h.GetPackageGroupByPackageID, h.PackageGetByID)
}
