package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
			Limit:  cc.GetLimitQuery(),
			Page:   cc.GetPageQuery(),
			Sort: bson.D{
				{"createdAt",-1},
			},
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
		staff = c.Get("staff").(model.StaffRaw)
	)

	status := c.Get("statusBody").(model.StatusBody)

	data, err := h.OrderAdminService.ChangeStatus(cc.GetRequestCtx(), order, status, staff)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"status": data,
	}, "")
}

func (h *OrderAdminHandler) UpdateOrderSuccess(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = c.Get("order").(model.OrderRaw)
		staff = c.Get("staff").(model.StaffRaw)
	)

	err := h.OrderAdminService.UpdateOrderSuccess(cc.GetRequestCtx(), order, staff)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (h *OrderAdminHandler) CancelOrder(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = c.Get("order").(model.OrderRaw)
		staff = c.Get("staff").(model.StaffRaw)
	)

	err := h.OrderAdminService.CancelOrder(cc.GetRequestCtx(), order, staff)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (h *OrderAdminHandler) GetDetail(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = c.Get("order").(model.OrderRaw)
	)

	data := h.OrderAdminService.GetDetail(cc.GetRequestCtx(), order)

	return cc.Response200(echo.Map{
		"order": data,
	}, "")
}

func (h *OrderAdminHandler) GetStatistic(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		query = model.CommonQuery{
			Sort: bson.D{
				bson.E{
					"order", 1},
			},
			StartAt: util.TimeParseISODate(cc.QueryParam("startAt")),
			EndAt:   util.TimeParseISODate(cc.QueryParam("endAt")),
		}
	)

	result, err := h.OrderAdminService.GetStatistic(cc.GetRequestCtx(), query)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"result": result,
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
