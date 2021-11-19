package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
)

// OrderAppHandler ...
type OrderAppHandler struct {
	OrderAppService model.OrderAppService
}

// Update ...
func (h *OrderAppHandler) Create(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		body = c.Get("orderBody").(model.OrderBody)
		user = c.Get("user").(model.UserRaw)
	)

	data, err := h.OrderAppService.Create(cc.GetRequestCtx(), user, body)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"order": data,
	}, "")
}

func (h *OrderAppHandler) GetDetail(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = cc.Get("order").(model.OrderRaw)
	)

	data := h.OrderAppService.GetDetail(cc.GetRequestCtx(), order)

	return cc.Response200(echo.Map{
		"order": data,
	}, "")
}

// GetByID ...
func (h *OrderAppHandler) GetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = util.EchoGetCustomCtx(c)
		)
		orderIDString := c.Param("orderID")
		orderID, err := primitive.ObjectIDFromHex(orderIDString)
		if orderID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}

		order, err := h.OrderAppService.FindByID(cc.GetRequestCtx(), orderID)
		if order.ID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}
		c.Set("order", order)
		return next(c)
	}
}

func (h *OrderAppHandler) GetList(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		user  = c.Get("user").(model.UserRaw)
		query = model.CommonQuery{
			Status: c.QueryParam("status"),
			Limit:  cc.GetLimitQuery(),
			Page:   cc.GetPageQuery(),
		}
	)

	data, total := h.OrderAppService.GetList(cc.GetRequestCtx(), query, user)

	return cc.Response200(echo.Map{
		"order": data,
		"total": total,
	}, "")
}
