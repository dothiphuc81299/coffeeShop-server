package handler

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// CategoryAdminHandler ...
type CategoryAdminHandler struct {
	CategoryAdminService model.CategoryAdminService
}

// Create ...
func (d *CategoryAdminHandler) Create(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		CategoryBody = c.Get("categoryBody").(model.CategoryBody)
	)

	data, err := d.CategoryAdminService.Create(customCtx.GetRequestCtx(), CategoryBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

// Update ...
func (d *CategoryAdminHandler) Update(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		CategoryBody = c.Get("categoryBody").(model.CategoryBody)
		category     = c.Get("category").(model.CategoryRaw)
	)

	err := d.CategoryAdminService.Update(customCtx.GetRequestCtx(), category, CategoryBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(nil, "")
}

// ListAll ...
func (d *CategoryAdminHandler) ListAll(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Keyword: c.QueryParam("keyword"),
		}
	)

	data, total := d.CategoryAdminService.ListAll(context.Background(), query)

	result := model.ResponseAdminListData{
		Data:  data,
		Total: total,
	}
	return customCtx.Response200(result, "")
}

// CategoryGetByID ...
func (d *CategoryAdminHandler) CategoryGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("categoryID")
		if id == "" {
			return next(c)
		}
		categoryID := util.GetAppIDFromHex(id)
		if categoryID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		category, err := d.CategoryAdminService.FindByID(customCtx.GetRequestCtx(), categoryID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("category", category)
		return next(c)
	}
}
