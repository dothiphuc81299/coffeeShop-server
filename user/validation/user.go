package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// UserBodyValidation ...
func UserBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.UserBody
		c.Bind(&body)
		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}
		c.Set("body", body)
		return next(c)
	}
}

// UserLoginBodyValidation ...
func UserLoginBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.UserLoginBody
		c.Bind(&body)
		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}
		c.Set("body", body)
		return next(c)
	}
}
