package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
)

func (s *Server) NewShippingHandler(e *echo.Echo) {
	user := e.Group("/api/shipping")

	user.GET("/", s.searchShippings, middleware.AuthMiddleware(token.User, ""))
	user.GET("/detail/:shippingID", s.getShippingByID, middleware.AuthMiddleware(token.User, ""))
	user.POST("/", s.CreateShipping, middleware.AuthMiddleware(token.User, ""))
	user.PUT("/detail/:shippingID", s.updateShipping, middleware.AuthMiddleware(token.User, ""))
	user.DELETE("/detail/:shippingID", s.DeleteShipping, middleware.AuthMiddleware(token.User, ""))
}

func (s *Server) searchShippings(c echo.Context) error {
	var (
		cc    = util.EchoGetCustomCtx(c)
		query shippingaddress.SearchShippingAddressQuery
	)

	if err := c.Bind(&query); err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, total, err := s.Dependences.ShippingSrv.Search(cc.GetRequestCtx(), &query)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"order": data,
		"total": total,
	}, "")

}

func (s *Server) getShippingByID(c echo.Context) error {
	var (
		cc         = util.EchoGetCustomCtx(c)
		shippingID = c.Param("shippingID")
	)

	if shippingID == "" {
		return cc.Response400(nil, "shippingID is required")
	}

	shippingIDObj := util.GetObjectIDFromHex(shippingID)
	if shippingIDObj.IsZero() {
		return cc.Response400(nil, "shippingID is invalid")
	}

	data, err := s.Dependences.ShippingSrv.GetDetail(cc.GetRequestCtx(), shippingIDObj)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"shipping": data,
	}, "")
}

func (s *Server) CreateShipping(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		cmd shippingaddress.CreateShippingAddressCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.ShippingSrv.Create(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) updateShipping(c echo.Context) error {
	var (
		cc         = util.EchoGetCustomCtx(c)
		shippingID = c.Param("shippingID")
		cmd        shippingaddress.UpdateShippingAddressCommand
	)

	if shippingID == "" {
		return cc.Response400(nil, "shippingID is required")
	}

	shippingIDObj := util.GetObjectIDFromHex(shippingID)
	if shippingIDObj.IsZero() {
		return cc.Response400(nil, "shippingID is invalid")
	}
	cmd.ID = shippingIDObj

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.ShippingSrv.Update(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) DeleteShipping(c echo.Context) error {
	var (
		cc         = util.EchoGetCustomCtx(c)
		shippingID = c.Param("shippingID")
	)

	if shippingID == "" {
		return cc.Response400(nil, "shippingID is required")
	}

	shippingIDObj := util.GetObjectIDFromHex(shippingID)
	if shippingIDObj.IsZero() {
		return cc.Response400(nil, "shippingID is invalid")
	}
	err := s.Dependences.ShippingSrv.Delete(cc.GetRequestCtx(), shippingIDObj)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}
