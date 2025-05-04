 package middleware

// import (
// 	"net/http"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// // AppMiddleware ....
// func AppMiddleware() []echo.MiddlewareFunc {
// 	return []echo.MiddlewareFunc{cors, log, middleware.Recover(), jwt()}
// }

// var allowHeaders = []string{
// 	config.HeaderAcceptLanguage,
// 	config.HeaderAuthorization,
// 	config.HeaderContentLength,
// 	config.HeaderContentType,
// 	config.HeaderOrigin,
// }

// var allowMethods = []string{
// 	http.MethodGet,
// 	http.MethodPost,
// 	http.MethodOptions,
// 	http.MethodPut,
// 	http.MethodPatch,
// 	http.MethodDelete,
// }

// var cors = middleware.CORSWithConfig(middleware.CORSConfig{
// 	AllowOrigins:     []string{"*"},
// 	AllowMethods:     allowMethods,
// 	AllowHeaders:     allowHeaders,
// 	AllowCredentials: false,
// 	MaxAge:           600,
// })

// var log = middleware.LoggerWithConfig(middleware.LoggerConfig{
// 	Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
// })

// func jwt() echo.MiddlewareFunc {
// 	return middleware.JWTWithConfig(middleware.JWTConfig{
// 		SigningKey: []byte(config.GetEnv().AuthSecret),
// 		Skipper: func(c echo.Context) bool {
// 			cc := util.EchoGetCustomCtx(c)
// 			auth := cc.QueryParam("auth")
// 			authHeader := cc.GetHeaderKey(config.HeaderAuthorization)
// 			if authHeader == "" && auth != "" {
// 				authHeader = "Bearer" + auth
// 				cc.Request().Header.Set(config.HeaderAuthorization, authHeader)
// 				cc.Request().Header.Set("user", authHeader)
// 			}
// 			// return true to skip middleware
// 			return authHeader == "" || authHeader == "Bearer" || authHeader == "Bearer "

// 		},
// 	})
// }
