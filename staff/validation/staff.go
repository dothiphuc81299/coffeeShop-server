package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// StaffBodyValidation ...
func StaffBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.StaffBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

// StaffUpdateRoleBodyValidation ...
func StaffUpdateRoleBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.StaffUpdateRoleBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

// StaffLoginBodyValidation ...
func StaffLoginBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.StaffLoginBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

func StaffChangePasswordBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.PasswordBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

func StaffUpdateBodyByItValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.StaffUpdateBodyByIt
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("bodyUpdate", body)
		return next(c)
	}
}
