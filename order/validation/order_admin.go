package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

func StatusBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.StatusBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {

			return util.ValidationError(c, err)
		}

		c.Set("statusBody", body)

		return next(c)

	}
}
