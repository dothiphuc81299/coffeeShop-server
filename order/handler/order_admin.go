package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderAdminHandler ...
type OrderAdminHandler struct {
	OrderAdminService model.OrderAdminService
}

func (h *OrderAdminHandler) GetListByStatus(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		query = model.CommonQuery{
			Status: c.QueryParam("status"),
		}
	)

	data, total := h.OrderAdminService.GetListByStatus(cc.GetRequestCtx(), query)

	return cc.Response200(echo.Map{
		"order": data,
		"total": total,
	}, "")
}

func (h *OrderAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = c.Get("order").(model.OrderRaw)
	)

	status := c.Get("statusBody").(model.StatusBody)

	data, err := h.OrderAdminService.ChangeStatus(cc.GetRequestCtx(), order, status)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"status": data,
	}, "")
}

// GetByID ...
func (h *OrderAdminHandler) GetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = util.EchoGetCustomCtx(c)
		)
		orderIDString := c.Param("orderID")
		orderID, err := primitive.ObjectIDFromHex(orderIDString)

		if orderID.IsZero() || err != nil {

			return cc.Response404(nil, locale.CommonKeyNotFound)
		}

		order, err := h.OrderAdminService.FindByID(cc.GetRequestCtx(), orderID)

		if order.ID.IsZero() || err != nil {

			return cc.Response400(nil, locale.CommonKeyNotFound)
		}

		c.Set("order", order)
		return next(c)
	}
}
