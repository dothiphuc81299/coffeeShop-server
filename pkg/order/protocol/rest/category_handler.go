package rest

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
)

var (
	reqCategoryCreate = role.ResourceCategory + "_" + role.PermissionCreate
	reqCategoryUpdate = role.ResourceCategory + "_" + role.PermissionUpdate
	reqCategoryDelete = role.ResourceCategory + "_" + role.PermissionDelete
	reqCategoryView   = role.ResourceCategory + "_" + role.PermissionView
)

func (s *Server) NewCategoryHandler(e *echo.Echo) {
	admin := e.Group("/api/admin/category")
	user := e.Group("/api/category")

	user.GET("/", s.searchCategories)
	user.GET("/detail/:categoryID", s.getCategoryByID)

	admin.POST("/", s.createCategory, middleware.AuthMiddleware(token.Staff, reqCategoryCreate))
	admin.PUT("/detail/:categoryID", s.updateCategory, middleware.AuthMiddleware(token.Staff, reqCategoryUpdate))
	admin.DELETE("/detail/:categoryID", s.DeleteCategory, middleware.AuthMiddleware(token.Staff, reqCategoryDelete))
	admin.GET("/", s.searchCategories, middleware.AuthMiddleware(token.Staff, reqCategoryView))
	admin.GET("/detail/:categoryID", s.getCategoryByID, middleware.AuthMiddleware(token.Staff, reqCategoryView))

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

	err := s.Dependences.CategorySrv.Create(customCtx.GetRequestCtx(), &cmd)
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
		query     category.SearchCategoryQuery
	)

	err := c.Bind(&query)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	data, total, err := s.Dependences.CategorySrv.ListAll(context.Background(), &query)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := category.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
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
