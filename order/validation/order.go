package validation

import (
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

func OrderBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.OrderBody
		c.Bind(&body)
		if err := body.Validate(); err != nil {
			fmt.Println("1")
			return util.ValidationError(c, err)
		}

		c.Set("orderBody", body)
		return next(c)

	}
}
