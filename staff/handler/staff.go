package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffAdminHandler struct {
	StaffService model.StaffAdminService
}

// ListStaff ...
func (h *StaffAdminHandler) ListStaff(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
	)

	q := model.CommonQuery{
		Page:  cc.GetPageQuery(),
		Limit: cc.GetLimitQuery(),
		Sort: bson.D{
			bson.E{Key: "createdAt", Value: -1},
		},
		Active:  c.QueryParam("active"),
		Keyword: c.QueryParam("keyword"),
	}

	staffs, total := h.StaffService.ListStaff(ctx, q)

	return cc.Response200(echo.Map{
		"staffs": staffs,
		"total":  total,
		"limit":  q.Limit,
	}, "")
}

// Create ...
func (h *StaffAdminHandler) Create(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(model.StaffBody)
	)

	data, err := h.StaffService.Create(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"staff": data,
	}, "")
}

// Update ...
func (h *StaffAdminHandler) Update(c echo.Context) error {
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

// ChangeStatus ...
func (h *StaffAdminHandler) ChangeStatus(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		staff = c.Get("staff").(model.StaffRaw)
	)

	data, err := h.StaffService.ChangeStatus(cc.GetRequestCtx(), staff)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"staff": data,
	}, "")
}

// GetToken ...
func (h *StaffAdminHandler) GetToken(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	ctx := cc.GetRequestCtx()

	staffID, _ := primitive.ObjectIDFromHex(c.QueryParam("staffId"))

	data, err := h.StaffService.GetToken(ctx, staffID)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"token": data,
	}, "")
}

func (h *StaffAdminHandler) StaffGetByID(next echo.HandlerFunc) echo.HandlerFunc {
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
