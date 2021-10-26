package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

type UserPackageAppHandler struct {
	UserPackageAppService model.UserPackageGroupAppService
}

func (u *UserPackageAppHandler) Create(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		body      = c.Get("body").(model.UserPackageBody)
	)

	err := u.UserPackageAppService.ChoosePakage(customCtx.GetRequestCtx(), body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}
	return customCtx.Response200(result, "")
}
