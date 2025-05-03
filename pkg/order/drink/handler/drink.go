package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

func (d *DrinkAdminHandler) Search(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Keyword:  c.QueryParam("keyword"),
			Active:   c.QueryParam("active"),
			Category: c.QueryParam("category"),
			Limit:    customCtx.GetLimitQuery(),
			Page:     customCtx.GetPageQuery(),
			Sort: bson.D{
				{"createdAt", -1},
			},
		}
	)

	data, total := d.DrinkAdminService.ListAll(customCtx.GetRequestCtx(), query)

	result := model.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return customCtx.Response200(result, "")
}

func (d *DrinkAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		drinkBody = c.Get("drinkBody").(model.DrinkBody)
		drink     = c.Get("drink").(model.DrinkRaw)
	)

	data, err := d.DrinkAdminService.Update(customCtx.GetRequestCtx(), drink, drinkBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

func (d *DrinkAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		drink     = c.Get("drink").(model.DrinkRaw)
	)

	data, err := d.DrinkAdminService.ChangeStatus(customCtx.GetRequestCtx(), drink)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

func (d *DrinkAdminHandler) GetDetail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)

		drink = c.Get("drink").(model.DrinkRaw)
	)

	data := d.DrinkAdminService.GetDetail(customCtx.GetRequestCtx(), drink)

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

func (d *DrinkAdminHandler) DeleteDrink(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		drink     = c.Get("drink").(model.DrinkRaw)
	)

	err := d.DrinkAdminService.DeleteDrink(customCtx.GetRequestCtx(), drink)

	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}
	return customCtx.Response200(nil, "")
}

// DrinkGetByID ...
func (d *DrinkAdminHandler) DrinkGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("drinkID")
		if id == "" {
			return next(c)
		}
		drinkID := util.GetAppIDFromHex(id)
		if drinkID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		drink, err := d.DrinkAdminService.FindByID(customCtx.GetRequestCtx(), drinkID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("drink", drink)
		return next(c)
	}
}
