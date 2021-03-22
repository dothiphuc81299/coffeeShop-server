package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/account/handler"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/role/validation"
	"github.com/labstack/echo/v4"
)

// InitAccountAdmin ...
func InitAccountAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.AccountAdminHandler{
		AccountAdminService: cs.Account,
	}

	r := e.Group("/accounts")
	// Update
	r.PUT("/:accountID", h.Update, h.GetByID, validation.RoleBodyValidation)
	
	// Change status
	r.PATCH("/:accountID/status", h.ChangeStatus,
		h.GetByID)
	// Generate token
	r.GET("/:accountID/token", h.GenerateToken, h.GetByID)
}
