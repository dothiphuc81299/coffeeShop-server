package server

import (
	"net/http"

	drinkroute "github.com/dothiphuc81299/coffeeShop-server/drink/route"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StartAdmin ...
func StartAdmin(service *model.AdminService, d *model.CommonDAO) *echo.Echo {
	server := echo.New()
	//CORS
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderAuthorization},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		MaxAge:           600,
		AllowCredentials: false,
	}))

	// Middleware
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	server.Use(middleware.Recover())
	startAdminHandler(server, service, d)
	return server
}

func startAdminHandler(e *echo.Echo, service *model.AdminService, d *model.CommonDAO) {
	// drink

	drinkroute.InitDrinkAdmin(e, service, d)
}
