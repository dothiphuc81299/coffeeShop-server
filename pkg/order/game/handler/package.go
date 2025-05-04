package handler

import (
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

type PackageAdminHandler struct {
	PackageAdminService      model.PackageAdminService
	PackageGroupAdminService model.PackageGroupAdminService
}

// Create ...
func (d *PackageAdminHandler) Create(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		body      = c.Get("body").(model.PackageBody)
	)

	err := d.PackageAdminService.Create(customCtx.GetRequestCtx(), body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}
	return customCtx.Response200(result, "")
}

func (d *PackageAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		raw       = c.Get("package").(model.PackageRaw)
		body      = c.Get("body").(model.PackageBody)
	)

	err := d.PackageAdminService.Update(customCtx.GetRequestCtx(), raw, body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}

	return customCtx.Response200(result, "")
}

func (d *PackageAdminHandler) ListAll(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Keyword: c.QueryParam("keyword"),
			Active:  c.QueryParam("active"),
			Limit:   customCtx.GetLimitQuery(),
			Page:    customCtx.GetPageQuery(),
		}
	)

	data, total := d.PackageAdminService.ListAll(customCtx.GetRequestCtx(), query)

	result := model.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return customCtx.Response200(result, "")
}

func (d *PackageAdminHandler) GetDetail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		raw       = c.Get("package").(model.PackageRaw)
	)

	data := d.PackageAdminService.GetDetail(customCtx.GetRequestCtx(), raw)
	return customCtx.Response200(echo.Map{
		"result": data,
	}, "")
}

func (d *PackageAdminHandler) PackageGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("packageID")
		if id == "" {
			return next(c)
		}
		packageID := util.Getprimitive.ObjectIDFromHex(id)
		if packageID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		packag, err := d.PackageAdminService.FindByID(customCtx.GetRequestCtx(), packageID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("package", packag)
		return next(c)
	}
}

func (p *PackageAdminHandler) GetPackageGroupByPackageID(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		raw       = c.Get("package").(model.PackageRaw)
	)
	fmt.Println("raw", raw.ID)
	data := p.PackageGroupAdminService.GetPackageGroupByPackageID(customCtx.GetRequestCtx(), raw.ID)
	return customCtx.Response200(echo.Map{
		"result": data,
	}, "")
}
