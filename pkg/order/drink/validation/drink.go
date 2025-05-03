package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// DrinkBodyValidation ...
func DrinkBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.DrinkBody
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}
		c.Set("drinkBody", body)
		return next(c)

	}
}
