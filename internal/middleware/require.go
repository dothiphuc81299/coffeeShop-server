package middleware

// import (
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
// 	"github.com/labstack/echo/v4"
// )

// // RequireLogin ...
// func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		cc := util.EchoGetCustomCtx(c)

// 		userID := cc.GetCurrentUserID()
// 		if userID.IsZero() {
// 			return cc.Response401(nil, "")
// 		}
// 		return next(c)
// 	}
// }
