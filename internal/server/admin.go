package server

import (
	drinkroute "github.com/dothiphuc81299/coffeeShop-server/drink/route"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"

	categoryroute "github.com/dothiphuc81299/coffeeShop-server/category/route"
	middleware "github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	orderroute "github.com/dothiphuc81299/coffeeShop-server/order/route"
	userroute "github.com/dothiphuc81299/coffeeShop-server/user/route"
	"github.com/labstack/echo/v4"

	eventroute "github.com/dothiphuc81299/coffeeShop-server/event/route"

	gameroute "github.com/dothiphuc81299/coffeeShop-server/game/route"
	staffroute "github.com/dothiphuc81299/coffeeShop-server/staff/route"
	staffrole "github.com/dothiphuc81299/coffeeShop-server/staffrole/route"
)

// StartAdmin ...
func StartAdmin(service *model.AdminService, app *model.AppService, d *model.CommonDAO) *echo.Echo {
	server := echo.New()
	// //CORS
	// server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderAuthorization},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
	// 	MaxAge:           600,
	// 	AllowCredentials: false,
	// }))

	// // Middleware
	// server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	// }))
	// server.Use(middleware.Recover())
	server.Use(middleware.AppMiddleware()...)
	startAdminHandler(server, service, app, d)
	return server
}

func startAdminHandler(e *echo.Echo, service *model.AdminService, app *model.AppService, d *model.CommonDAO) {

	drinkroute.InitDrinkAdmin(e, service, d)
	categoryroute.InitCategoryAdmin(e, service, d)
	userroute.InitUserAdmin(e, service, d)

	eventroute.InitEventAdmin(e, service, d)

	staffroute.InitStaffAdmin(e, service, d)
	staffroute.InitStaffApp(e, app, d)
	staffrole.InitStaffRoleAdmin(e, service, d)

	userroute.InitUserApp(e, app, d)

	orderroute.InitOrderApp(e, app, d)
	orderroute.InitOrderAdmin(e, service, d)
	gameroute.InitQuestionAdmin(e, service, d)
	gameroute.InitGroupAdmin(e, service, d)
	gameroute.InitPackageAdmin(e, service, d)
	gameroute.InitPackageGroupAdmin(e, service, d)

}
