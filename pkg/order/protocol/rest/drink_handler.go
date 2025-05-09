package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	reqDrinkCreate = role.ResourceDrink + "_" + role.PermissionCreate
	reqDrinkUpdate = role.ResourceDrink + "_" + role.PermissionUpdate
	reqDrinkDelete = role.ResourceDrink + "_" + role.PermissionDelete
	reqDrinkView   = role.ResourceDrink + "_" + role.PermissionView
)

func (s *Server) NewDrinkHandler(e *echo.Echo) {
	user := e.Group("/api/drink")
	admin := e.Group("/api/admin/drink")

	admin.POST("/", s.createDrink, middleware.AuthMiddleware(token.Staff, reqDrinkCreate))
	admin.GET("/", s.searchDrinks, middleware.AuthMiddleware(token.Staff, reqDrinkView))
	admin.PUT("/detail/:drinkID", s.updateDrink, middleware.AuthMiddleware(token.Staff, reqDrinkUpdate))
	admin.PATCH("/detail/:drinkID/status", s.ChangeDrinkStatus, middleware.AuthMiddleware(token.Staff, reqDrinkUpdate))
	admin.GET("/detail/:drinkID", s.getDrinkByID, middleware.AuthMiddleware(token.Staff, reqDrinkView))
	admin.DELETE("/detail/:drinkID", s.DeleteDrink, middleware.AuthMiddleware(token.Staff, reqDrinkDelete))

	user.GET("/", s.searchDrinks)
	user.GET("/detail/:drinkID", s.getDrinkByID)
}

func (s *Server) createDrink(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		cmd       drink.DrinkBody
	)

	err := c.Bind(&cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	err = s.Dependences.DrinkSrv.Create(customCtx.GetRequestCtx(), cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200("", "")
}

func (s *Server) searchDrinks(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		q         = query.CommonQuery{
			Keyword:  c.QueryParam("keyword"),
			Active:   c.QueryParam("active"),
			Category: c.QueryParam("category"),
			Limit:    customCtx.GetLimitQuery(),
			Page:     customCtx.GetPageQuery(),
			Sort: bson.D{
				{"createdAt", -1},
			},
		}
	)

	data, total := s.Dependences.DrinkSrv.ListAll(customCtx.GetRequestCtx(), q)
	result := model.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: q.Limit,
	}
	return customCtx.Response200(result, "")
}

func (s *Server) updateDrink(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		cmd       drink.DrinkBody
	)

	err := c.Bind(&cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	params := c.Param("drinkID")
	if params == "" {
		return customCtx.Response400(nil, "drinkID is required")
	}

	drinkID := util.GetObjectIDFromHex(params)
	if err := cmd.Validate(); err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	err = s.Dependences.DrinkSrv.Update(customCtx.GetRequestCtx(), drinkID, cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200("", "")
}

func (s *Server) ChangeDrinkStatus(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
	)

	param := c.Param("drinkID")
	if param == "" {
		return customCtx.Response400(nil, "drinkID is required")
	}
	drinkID := util.GetObjectIDFromHex(param)

	data, err := s.Dependences.DrinkSrv.ChangeStatus(customCtx.GetRequestCtx(), drinkID)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(data, "")
}

func (s *Server) DeleteDrink(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
	)

	param := c.Param("drinkID")
	if param == "" {
		return customCtx.Response400(nil, "drinkID is required")
	}
	drinkID := util.GetObjectIDFromHex(param)

	err := s.Dependences.DrinkSrv.DeleteDrink(customCtx.GetRequestCtx(), drinkID)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}
	return customCtx.Response200(nil, "")
}

func (s *Server) getDrinkByID(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
	)

	param := c.Param("drinkID")
	if param == "" {
		return customCtx.Response400(nil, "drinkID is required")
	}
	drinkID := util.GetObjectIDFromHex(param)

	data, err := s.Dependences.DrinkSrv.FindByID(customCtx.GetRequestCtx(), drinkID)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(echo.Map{
		"drink": data,
	}, "")
}
