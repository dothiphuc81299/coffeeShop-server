package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

type PackageGroupAdminHandler struct {
	PackageGroupAdminService model.PackageGroupAdminService
}

func (p *PackageGroupAdminHandler) Create(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		body      = c.Get("body").(model.PackageGroupBody)
	)

	err := p.PackageGroupAdminService.Create(customCtx.GetRequestCtx(), body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}
	return customCtx.Response200(result, "")
}

func (p *PackageGroupAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		raw       = c.Get("package-group").(model.PackageGroupRaw)
		body      = c.Get("body").(model.PackageGroupBody)
	)

	err := p.PackageGroupAdminService.Update(customCtx.GetRequestCtx(), raw, body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}

	return customCtx.Response200(result, "")
}

func (p *PackageGroupAdminHandler) GetPackageGroupByPackageID(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		raw       = c.Get("package").(model.PackageRaw)
	)

	data := p.PackageGroupAdminService.GetPackageGroupByPackageID(customCtx.GetRequestCtx(), raw.ID)
	return customCtx.Response200(echo.Map{
		"result": data,
	}, "")
}

func (p *PackageGroupAdminHandler) PackageGroupGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("packageGroupID")
		if id == "" {
			return next(c)
		}

		packageID := util.GetAppIDFromHex(id)
		if packageID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		packag, err := p.PackageGroupAdminService.FindByID(customCtx.GetRequestCtx(), packageID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("package-group", packag)
		return next(c)
	}
}
