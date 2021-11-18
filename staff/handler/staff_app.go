package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffAppHandler struct {
	StaffService model.StaffAppService
}

// Update ...
func (h *StaffAppHandler) Update(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body  = c.Get("bodyUpdate").(model.StaffUpdateBodyByIt)
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

func (h *StaffAppHandler) GetDetailStaff(c echo.Context) error {

	cc := util.EchoGetCustomCtx(c)
	var (
		staff = c.Get("staff").(model.StaffRaw)
	)
	data := h.StaffService.GetDetailStaff(cc.GetRequestCtx(), staff)

	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (h *StaffAppHandler) StaffGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = util.EchoGetCustomCtx(c)
		)
		staffIDString := c.Param("staffID")
		staffID, err := primitive.ObjectIDFromHex(staffIDString)
		if staffID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}

		staff, err := h.StaffService.FindByID(cc.GetRequestCtx(), staffID)
		if staff.ID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}
		c.Set("staff", staff)
		return next(c)
	}
}

