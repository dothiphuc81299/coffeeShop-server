package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) NewDrinkHandler(e *echo.Echo) {
	g := e.Group("/api/drink")

	g.POST("/", s.createDrink, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionDelete, token.Staff))
	g.GET("/", s.searchDrinks)
	g.PUT("/:drinkID", s.updateDrink, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionEdit, token.Staff))
	g.PATCH("/:drinkID/status", s.ChangeDrinkStatus, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionEdit, token.Staff))
	g.GET("/:drinkID", s.getDrinkByID)
	g.DELETE("/:drinkID", s.DeleteDrink, middleware.CheckPermission(config.ModelFieldDrink, config.PermissionDelete, token.Staff))
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
