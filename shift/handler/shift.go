package handler

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ShiftAdminHandler ...
type ShiftAdminHandler struct {
	ShiftAdminService model.ShiftAdminService
}

// Create ...
func (d *ShiftAdminHandler) Create(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		shiftBody = c.Get("shiftBody").(model.ShiftBody)
		staff     = c.Get("staff").(model.StaffRaw)
	)

	data, err := d.ShiftAdminService.Create(customCtx.GetRequestCtx(), shiftBody, staff)

	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

// Update ...
func (d *ShiftAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		shiftBody = c.Get("shiftBody").(model.ShiftBody)
		shift     = c.Get("shift").(model.ShiftRaw)
		staff     = c.Get("staff").(model.StaffRaw)
	)

	data, err := d.ShiftAdminService.Update(customCtx.GetRequestCtx(), shift, shiftBody, staff)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}

	return customCtx.Response200(result, "")
}

// ListAll ...
// list theo user,
// list theo thoi diem
// list theo trang thai
func (d *ShiftAdminHandler) ListAll(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)

		staff, _ = primitive.ObjectIDFromHex(customCtx.QueryParam("staff"))
		startAt  = util.TimeParseISODate(customCtx.QueryParam("startAt"))
		endAt    = util.TimeParseISODate(customCtx.QueryParam("endAt"))

		query = model.CommonQuery{
			IsCheck: c.QueryParam("isCheck"),
			Staff:   staff,
			StartAt: startAt,
			EndAt:   endAt,
			Keyword: c.QueryParam("keyword"),
		}
	)

	data, total := d.ShiftAdminService.ListAll(context.Background(), query)

	result := model.ResponseAdminListData{
		Data:  data,
		Total: total,
	}
	return customCtx.Response200(result, "")
}

func (d *ShiftAdminHandler) AcceptShiftByAdmin(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		shift     = c.Get("shift").(model.ShiftRaw)
	)

	data, err := d.ShiftAdminService.AcceptShiftByAdmin(customCtx.GetRequestCtx(), shift)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}

	return customCtx.Response200(result, "")
}

// ShiftGetByID ...
func (d *ShiftAdminHandler) ShiftGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("shiftID")
		if id == "" {
			return next(c)
		}
		ShiftID := util.GetAppIDFromHex(id)
		if ShiftID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		Shift, err := d.ShiftAdminService.FindByID(customCtx.GetRequestCtx(), ShiftID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("shift", Shift)
		return next(c)
	}
}
