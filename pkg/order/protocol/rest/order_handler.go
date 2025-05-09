package rest

import "github.com/labstack/echo/v4"

func (s *Server) NewOrderHandler(e *echo.Echo) {
	r := e.Group("/orders")

	r.GET("", h.SearchByStatus)
	r.PUT("/:orderID/status", h.ChangeStatus, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID, validation.StatusBodyValidation)
	r.GET("/:orderID", h.GetDetail, h.GetByID)
	r.GET("/statistic", h.GetStatistic, middleware.CheckPermission(config.ModelFieldStatistic, config.PermissionAdmin, d))

	r.PUT("/:orderID/success", h.UpdateOrderSuccess, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID)
	r.PUT("/:orderID/cancel", h.CancelOrder, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID)
}


package handler


package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/order/handler"
	"github.com/dothiphuc81299/coffeeShop-server/order/validation"
	"github.com/labstack/echo/v4"
)

// InitOrderAdmin ...
func InitOrderAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.OrderAdminHandler{
		OrderAdminService: cs.Order,
	}

	r := e.Group("/orders")

	r.GET("", h.SearchByStatus)
	r.PUT("/:orderID/status", h.ChangeStatus, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID, validation.StatusBodyValidation)
	r.GET("/:orderID", h.GetDetail, h.GetByID)
	r.GET("/statistic", h.GetStatistic, middleware.CheckPermission(config.ModelFieldStatistic, config.PermissionAdmin, d))

	r.PUT("/:orderID/success", h.UpdateOrderSuccess, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID)
	r.PUT("/:orderID/cancel", h.CancelOrder, middleware.CheckPermission(config.ModelFieldOrder, config.PermissionEdit, d), h.GetByID)

}

package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/order/handler"
	"github.com/dothiphuc81299/coffeeShop-server/order/validation"
	"github.com/labstack/echo/v4"
)

// InitOrderApp ...
func InitOrderApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.OrderAppHandler{
		OrderAppService: cs.Order,
	}

	r := e.Group("/orders")

	r.POST("", h.Create, middleware.CheckUser(d), validation.OrderBodyValidation)

	// xem chi tieet don hang
	r.GET("/:orderID/me", h.GetDetail, middleware.CheckUser(d), h.GetByID)

	r.GET("/me", h.Search, middleware.CheckUser(d))

	r.PUT("/:orderID/reject", h.RejectOrder, middleware.CheckUser(d), h.GetByID)
}

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderAdminHandler ...
type OrderAdminHandler struct {
	OrderAdminService model.OrderAdminService
}

func (h *OrderAdminHandler) SearchByStatus(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)
		q  = query.CommonQuery{
			Status:   c.QueryParam("status"),
			Limit:    cc.GetLimitQuery(),
			Page:     cc.GetPageQuery(),
			Username: cc.QueryParam("username"),
			Sort: bson.D{
				{"createdAt", -1},
			},
		}
	)

	data, total := h.OrderAdminService.SearchByStatus(cc.GetRequestCtx(), query)

	return cc.Response200(echo.Map{
		"order": data,
		"total": total,
	}, "")
}

func (h *OrderAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = c.Get("order").(model.OrderRaw)
		staff = c.Get("staff").(model.Staff)
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
		staff = c.Get("staff").(model.Staff)
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
		staff = c.Get("staff").(model.Staff)
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


package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

func (h *OrderAppHandler) Search(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		user  = c.Get("user").(model.UserRaw)
		query = model.CommonQuery{
			Status: c.QueryParam("status"),
			Limit:  cc.GetLimitQuery(),
			Page:   cc.GetPageQuery(),
			Sort: bson.D{
				{"createdAt", -1},
			},
		}
	)

	data, total := h.OrderAppService.Search(cc.GetRequestCtx(), query, user)

	return cc.Response200(echo.Map{
		"order": data,
		"total": total,
	}, "")
}

func (h *OrderAppHandler) RejectOrder(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		order = cc.Get("order").(model.OrderRaw)
		user  = c.Get("user").(model.UserRaw)
	)

	err := h.OrderAppService.RejectOrder(cc.GetRequestCtx(), user, order)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}
