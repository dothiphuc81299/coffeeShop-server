package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

func OrderBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.OrderBody
		c.Bind(&body)
		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}
		// log.Println("bodt2", body)
		c.Set("orderBody", body)

		//	log.Println("next", next(c))
		return next(c)

	}
}
