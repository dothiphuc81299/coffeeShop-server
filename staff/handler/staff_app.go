package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

type StaffAppHandler struct {
	StaffService model.StaffAppService
}

// Update ...
func (h *StaffAppHandler) Update(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body  = c.Get("body").(model.StaffBody)
		staff = c.Get("staff").(model.StaffRaw)
	)

	data, err := h.StaffService.Update(cc.GetRequestCtx(), body, staff)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"staff": data,
	}, "")
}


func (h *StaffAppHandler) UpdatePassword(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body  = c.Get("body").(model.PasswordBody)
		staff = c.Get("staff").(model.StaffRaw)
	)

	err := h.StaffService.ChangePassword(cc.GetRequestCtx(), staff, body)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}
