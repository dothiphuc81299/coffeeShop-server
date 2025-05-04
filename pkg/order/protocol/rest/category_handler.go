package rest

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
)

func (s *Server) NewCategoryHandler(e *echo.Echo) {
	g := e.Group("/api/category")
	g.POST("/", s.createCategory, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionView, token.Staff))
	g.PUT("/detail/:categoryID", s.updateCategory, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionEdit, token.Staff))

	g.GET("/", s.searchCategories)
	g.GET("/detail/:categoryID", s.getCategoryByID)

	g.DELETE("/detail/:categoryID", s.DeleteCategory, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionDelete, token.Staff))
}

func (s *Server) createCategory(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		cmd       category.CategoryBody
	)

	if err := c.Bind(&cmd); err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	err := s.Dependences.CategorySrv.Create(customCtx.GetRequestCtx(), cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200("", "")
}

func (s *Server) updateCategory(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		cmd       category.CategoryBody
	)

	if err := c.Bind(&cmd); err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	id := c.Param("categoryID")
	if id == "" {
		return customCtx.Response400(nil, "categoryID is required")
	}
	categoryID := util.GetObjectIDFromHex(id)

	err := s.Dependences.CategorySrv.Update(customCtx.GetRequestCtx(), categoryID, cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200("", "")
}

func (s *Server) searchCategories(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		cmd       = query.CommonQuery{
			Keyword: c.QueryParam("keyword"),
			Limit:   customCtx.GetLimitQuery(),
			Page:    customCtx.GetPageQuery(),
		}
	)

	data, total := s.Dependences.CategorySrv.ListAll(context.Background(), cmd)
	result := model.ResponseAdminListData{
		Data:  data,
		Total: total,
	}
	return customCtx.Response200(result, "")
}

func (s *Server) getCategoryByID(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
	)

	id := c.Param("categoryID")
	if id == "" {
		return customCtx.Response400(nil, "categoryID is required")
	}
	categoryID := util.GetObjectIDFromHex(id)

	data, err := s.Dependences.CategorySrv.GetDetail(customCtx.GetRequestCtx(), categoryID)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(echo.Map{
		"category": data,
	}, "")

}

func (s *Server) DeleteCategory(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
	)

	id := c.Param("categoryID")
	if id == "" {
		return customCtx.Response400(nil, "categoryID is required")
	}
	categoryID := util.GetObjectIDFromHex(id)
	err := s.Dependences.CategorySrv.DeleteCategory(customCtx.GetRequestCtx(), categoryID)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}
	return customCtx.Response200(nil, "")
}
