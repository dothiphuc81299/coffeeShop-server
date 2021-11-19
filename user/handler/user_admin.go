package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// UserAdminHandler ...
type UserAdminHandler struct {
	UserAdminService model.UserAdminService
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

	data, total := h.UserAdminService.GetList(cc.GetRequestCtx(), query)
	result := model.ResponseAppListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return cc.Response200(result, "")
}

// ChangeStatus ...
func (h *UserAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		user = c.Get("user").(model.UserRaw)
	)
	active, err := h.UserAdminService.ConfirmAccountActive(cc.GetRequestCtx(), user)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"active": active,
	}, "")
}

func (h *UserAdminHandler) GetDetailUser(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		user = c.Get("user").(model.UserRaw)
	)

	result := h.UserAdminService.GetDetailUser(cc.GetRequestCtx(), user)
	return cc.Response200(echo.Map{
		"data": result,
	}, "")

}

// UserGetByID ...
func (d *UserAdminHandler) UserGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("userID")
		if id == "" {
			return next(c)
		}
		userID := util.GetAppIDFromHex(id)
		if userID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		user, err := d.UserAdminService.FindByID(customCtx.GetRequestCtx(), userID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("user", user)
		return next(c)
	}
}
