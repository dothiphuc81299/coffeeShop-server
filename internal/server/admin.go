package server

import (
	drinkroute "github.com/dothiphuc81299/coffeeShop-server/drink/route"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"

	categoryroute "github.com/dothiphuc81299/coffeeShop-server/category/route"
	middleware "github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/labstack/echo/v4"
)

// StartAdmin ...
func StartAdmin(service *model.AdminService, d *model.CommonDAO) *echo.Echo {
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
	startAdminHandler(server, service, d)
	return server
}

func startAdminHandler(e *echo.Echo, service *model.AdminService, d *model.CommonDAO) {
	// drink

	drinkroute.InitDrinkAdmin(e, service, d)
	categoryroute.InitCategoryAdmin(e, service, d)
}
