package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/category/handler"
	"github.com/dothiphuc81299/coffeeShop-server/category/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitCategoryAdmin ...
func InitCategoryAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.CategoryAdminHandler{
		CategoryAdminService: cs.Category,
	}

	g := e.Group("/category")
	g.POST("", h.Create, validation.CategoryBodyValidation)

	g.PUT("/:categoryID", h.Update, h.CategoryGetByID, validation.CategoryBodyValidation)

	g.GET("", h.ListAll)
}
