package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RoleAdminHandler ...
type RoleAdminHandler struct {
	RoleAdminService model.RoleService
}

// Create ...
func (h *RoleAdminHandler) Create(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		body = c.Get("body").(model.RoleBody)
	)

	data, err := h.RoleAdminService.Create(cc.GetRequestCtx(), body)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

// List ...
func (h *RoleAdminHandler) List(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
		q   = model.CommonQuery{
			Sort: bson.D{
				bson.E{Key: "createdAt", Value: -1},
			},
			Limit: cc.GetLimitQuery(),
			Page:  cc.GetPageQuery(),
		}
	)

	data, total := h.RoleAdminService.List(ctx, q)
	result := model.ResponseAppListData{
		Data:  data,
		Total: total,
	}
	return cc.Response200(result, "")
}

// Update ...
func (h *RoleAdminHandler) Update(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		body = c.Get("body").(model.RoleBody)
		role = c.Get("role").(model.RoleRaw)
	)
	data, err := h.RoleAdminService.Update(cc.GetRequestCtx(), body, role)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(model.ResponseAdminData{
		Data: data,
	}, "")
}

func (h *RoleAdminHandler) GetDetail(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		role = c.Get("role").(model.RoleRaw)
	)
	data := h.RoleAdminService.GetDetail(cc.GetRequestCtx(), role)
	return cc.Response200(model.ResponseAdminData{
		Data: data,
	}, "")
}

// GetByID ...
func (h *RoleAdminHandler) GetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = util.EchoGetCustomCtx(c)
		)
		staffRoleIDString := c.Param("roleID")
		staffRoleID, err := primitive.ObjectIDFromHex(staffRoleIDString)
		if staffRoleID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}

		role, err := h.RoleAdminService.FindByID(cc.GetRequestCtx(), staffRoleID)
		if role.ID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}
		c.Set("role", role)
		return next(c)
	}
}
