package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/user/handler"
	"github.com/dothiphuc81299/coffeeShop-server/user/validation"
	"github.com/labstack/echo/v4"
)

// InitUserApp ...
func InitUserApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.UserAppHandler{
		UserAppService: cs.User,
	}

	r := e.Group("/users")

	// sign up
	r.POST("/sign-up", h.UserSignUp, validation.UserSignUpBodyValidation)

	// log in

	r.POST("/log-in", h.UserLoginIn, validation.UserLoginBodyValidation)

	// update user
	r.PUT("/update", h.UserUpdateAccount, middleware.CheckUser(d), validation.UserSignUpBodyValidation)

	// get detail user
	r.GET("/me", h.GetDetailUser, middleware.CheckUser(d))
}
