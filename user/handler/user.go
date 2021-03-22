package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserAdminHandler ...
type UserAdminHandler struct {
	UserAdminService model.UserAdminService
}

// Create ...
func (h *UserAdminHandler) Create(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		body = c.Get("body").(model.UserBody)
	)
	data, err := h.UserAdminService.Create(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

// List ...
func (h *UserAdminHandler) List(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		query = model.CommonQuery{
			Limit:   cc.GetLimitQuery(),
			Page:    cc.GetPageQuery(),
			Active:  c.QueryParam("active"),
			Keyword: c.QueryParam("keyword"),
			Sort:    bson.D{{"createdAt", -1}},
		}
	)

	data, total := h.UserAdminService.List(cc.GetRequestCtx(), query)
	result := model.ResponseAppListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return cc.Response200(result, "")
}

// Update ...
func (h *UserAdminHandler) Update(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		body = c.Get("body").(model.UserBody)
		user = c.Get("user").(model.UserRaw)
	)
	data, err := h.UserAdminService.Update(cc.GetRequestCtx(), body, user)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

// ChangeStatus ...
func (h *UserAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		user = c.Get("user").(model.UserRaw)
	)
	active, err := h.UserAdminService.ChangeStatus(cc.GetRequestCtx(), user)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"active": active,
	}, "")
}

// Detail ...
func (h *UserAdminHandler) Detail(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		user = c.Get("user").(model.UserRaw)
	)
	data := h.UserAdminService.GetDetail(cc.GetRequestCtx(), user)
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

// GetByID ...
func (h *UserAdminHandler) GetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = util.EchoGetCustomCtx(c)
		)
		userIDString := c.Param("userID")
		userID, err := primitive.ObjectIDFromHex(userIDString)
		if userID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}
		user, err := h.UserAdminService.FindByID(cc.GetRequestCtx(), userID)
		if user.ID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}
		c.Set("user", user)
		return next(c)
	}
}
