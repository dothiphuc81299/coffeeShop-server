package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// UserSignUpBodyValidation ...
func UserSignUpBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.UserSignUpBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

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

// UserUpdateBodyValidation ...
func UserUpdateBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.UserUpdateBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

func UserChangePasswordValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.UserChangePasswordBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

func SendEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.UserSendEmailBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

func VerifyEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.VerifyEmailBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}
