package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/category/handler"
	"github.com/dothiphuc81299/coffeeShop-server/category/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitCategoryAdmin ...
func InitCategoryAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.CategoryAdminHandler{
		CategoryAdminService: cs.Category,
	}

	g := e.Group("/category")
	g.POST("", h.Create, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionView, d), validation.CategoryBodyValidation)
	g.PUT("/:categoryID", h.Update, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionEdit, d), h.CategoryGetByID, validation.CategoryBodyValidation)

	g.GET("", h.ListAll)
	g.GET("/:categoryID", h.GetDetail, h.CategoryGetByID)

	g.DELETE("/:categoryID", h.DeleteCategory, h.CategoryGetByID, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionDelete, d))
}
