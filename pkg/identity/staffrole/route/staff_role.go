 package route

// import (
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
// 	"github.com/dothiphuc81299/coffeeShop-server/staffrole/handler"
// 	"github.com/dothiphuc81299/coffeeShop-server/staffrole/validation"
// 	"github.com/labstack/echo/v4"
// )

// // InitStaffRoleAdmin ...
// func InitStaffRoleAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
// 	h := &handler.StaffRoleAdminHandler{
// 		StaffRoleAdminService: cs.StaffRole,
// 	}

// 	g := e.Group("/staffRole")
// 	g.POST("", h.Create, middleware.CheckPermissionRoot(token.Root), validation.CreateStaffRoleCommandValidation)

// 	g.PUT("/:roleID", h.Update, middleware.CheckPermissionRoot(token.Root), h.StaffRoleGetByID, validation.CreateStaffRoleCommandValidation)

// 	// Get list roles
// 	g.GET("", h.ListRoleStaff, middleware.CheckPermission(config.ModelFieldRole, config.PermissionView, d))

// 	// get List permission
// 	g.GET("/permissions", h.SearchPermission)

// 	// delete role
// 	g.DELETE("/:roleID", h.Delete, middleware.CheckPermissionRoot(token.Root), h.StaffRoleGetByID)
// }
