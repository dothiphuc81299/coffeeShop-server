package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// CategoryBodyValidation ...
func CategoryBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.CategoryBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}
		c.Set("categoryBody", body)
		return next(c)

	}
}
