package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// DrinkAdminHandler ...
type DrinkAdminHandler struct {
	DrinkAdminService model.DrinkAdminService
}

// Create ...
func (d *DrinkAdminHandler) Create(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		drinkBody = c.Get("drinkBody").(model.DrinkBody)
	)

	data, err := d.DrinkAdminService.Create(customCtx.GetRequestCtx(), drinkBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}
