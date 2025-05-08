package route

// import (
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
// 	"github.com/dothiphuc81299/coffeeShop-server/staff/handler"
// 	"github.com/dothiphuc81299/coffeeShop-server/staff/validation"
// 	"github.com/labstack/echo/v4"
// )

// // InitStaffAdmin ...
// func InitStaffAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
// 	h := &handler.StaffAdminHandler{
// 		StaffService: cs.Staff,
// 	}

// 	// only root
// 	r := e.Group("/staff")
// 	r.POST("", h.Create, middleware.AuthMiddleware(token.Root,""), validation.CreateStaffCommandValidation)
// 	r.GET("", h.ListStaff, middleware.CheckPermission(config.ModelFieldStaff, config.PermissionView, d))
// 	r.GET("/:staffID", h.GetStaffByID, middleware.AuthMiddleware(token.Root,""))
// 	r.GET("/token", h.GetToken)
// 	// r.PATCH("/:staffID/status", h.ChangeStatus, middleware.AuthMiddleware(token.Root,""), h.StaffGetByID)
// 	r.PUT("/:staffID", h.UpdateRole, middleware.AuthMiddleware(token.Root,""), h.StaffGetByID, validation.UpdateStaffRoleCommandValidation)
// 	r.DELETE("/:staffID", h.DeleteStaff, middleware.AuthMiddleware(token.Root,""), h.StaffGetByID)
// 	r.POST("/log-in", h.LoginStaff, validation.LoginStaffCommandValidation)

// }
