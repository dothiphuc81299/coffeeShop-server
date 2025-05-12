package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/order"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	reqOrderView   = role.ResourceOrder + "_" + role.PermissionView
	reqOrderUpdate = role.ResourceOrder + "_" + role.PermissionUpdate
)

func (s *Server) NewOrderHandler(e *echo.Echo) {
	admin := e.Group("/api/admin/orders")
	user := e.Group("/api/orders")

	admin.GET("/", s.searchOrders, middleware.AuthMiddleware(token.Staff, reqOrderView))
	admin.GET("/detail/:orderID", s.getOrderByID, middleware.AuthMiddleware(token.Staff, reqOrderView))
	admin.GET("/statistic", s.GetStatistic, middleware.AuthMiddleware(token.Staff, reqOrderView))
	admin.PUT("/detail/:orderID/success", s.approveOrder, middleware.AuthMiddleware(token.Staff, reqOrderUpdate))
	admin.PUT("/detail/:orderID/cancel", s.CancelOrder, middleware.AuthMiddleware(token.Staff, reqOrderUpdate))

	user.POST("/", s.CreateOrder, middleware.AuthMiddleware(token.User, ""))
	user.GET("/detail/:orderID/me", s.getOrderByID, middleware.AuthMiddleware(token.User, ""))
	user.GET("/", s.searchOrders, middleware.AuthMiddleware(token.User, ""))
	user.PUT("/detail/:orderID/reject", s.RejectOrder, middleware.AuthMiddleware(token.User, ""))
}

func (s *Server) searchOrders(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)

		query order.SearchOrdersQuery
	)
	if err := c.Bind(&query); err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, total := s.Dependences.OrderSrv.Search(cc.GetRequestCtx(), &query)
	return cc.Response200(echo.Map{
		"order": data,
		"total": total,
	}, "")
}

func (s *Server) approveOrder(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		cmd order.UpdateOrderStatusCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	params := c.Param("orderID")
	if params == "" {
		return cc.Response400(nil, "orderID is required")
	}

	orderID := util.GetObjectIDFromHex(params)
	cmd.ID = orderID

	err := s.Dependences.OrderSrv.ApproveOrder(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) CancelOrder(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		cmd order.UpdateOrderStatusCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	params := c.Param("orderID")
	if params == "" {
		return cc.Response400(nil, "orderID is required")
	}

	orderID := util.GetObjectIDFromHex(params)
	cmd.ID = orderID

	err := s.Dependences.OrderSrv.RejectOrder(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) getOrderByID(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)
	)

	params := c.Param("orderID")
	if params == "" {
		return cc.Response400(nil, "orderID is required")
	}

	orderID := util.GetObjectIDFromHex(params)

	data, err := s.Dependences.OrderSrv.GetDetail(cc.GetRequestCtx(), orderID)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	if data.ID.IsZero() || err != nil {
		return cc.Response400(nil, "order not found")
	}

	return cc.Response200(echo.Map{
		"order": data,
	}, "")
}

func (s *Server) GetStatistic(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		cmd = query.CommonQuery{
			Sort: bson.D{
				bson.E{
					"order", 1},
			},
			// StartAt: util.TimeParseISODate(cc.QueryParam("startAt")),
			// EndAt:   util.TimeParseISODate(cc.QueryParam("endAt")),
		}
	)

	result, err := s.Dependences.OrderSrv.GetStatistic(cc.GetRequestCtx(), cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"result": result,
	}, "")

}

func (s *Server) CreateOrder(c echo.Context) error {

	var (
		cc = util.EchoGetCustomCtx(c)
	)

	var body order.OrderBody
	if err := c.Bind(&body); err != nil {
		return cc.Response400(nil, err.Error())
	}
	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, err := s.Dependences.OrderSrv.Create(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"order": data,
	}, "")
}

func (s *Server) RejectOrder(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)
	)

	orderID := c.Param("orderID")
	if orderID == "" {
		return cc.Response400(nil, "orderID is required")
	}

	orderIDObj := util.GetObjectIDFromHex(orderID)
	if orderIDObj.IsZero() {
		return cc.Response400(nil, "orderID is invalid")
	}

	var cmd order.UpdateOrderStatusCommand
	cmd.Status = "cancel"
	cmd.ID = orderIDObj
	err := s.Dependences.OrderSrv.RejectOrder(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}
