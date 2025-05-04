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
	g := e.Group("/category")
	g.POST("", s.createCategory, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionView, token.Staff))
	g.PUT("/:categoryID", s.updateCategory, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionEdit, d), h.CategoryGetByID, validation.CategoryBodyValidation)

	g.GET("", h.ListAll)
	g.GET("/:categoryID", h.GetDetail, h.CategoryGetByID)

	g.DELETE("/:categoryID", h.DeleteCategory, h.CategoryGetByID, middleware.CheckPermission(config.ModelFieldCategory, config.PermissionDelete, d))
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
	categoryID := util.GetAppIDFromHex(id)

	err := s.Dependences.CategorySrv.Update(customCtx.GetRequestCtx(), categoryID, cmd)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200("", "")
}

// ListAll ...
func (s *Server) ListAll(c echo.Context) error {
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

func (s *Server) GetDetail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		category  = c.Get("category").(model.CategoryRaw)
	)

	data := s.Dependences.CategorySrvGetDetail(customCtx.GetRequestCtx(), category)
	return customCtx.Response200(echo.Map{
		"category": data,
	}, "")

}

func (s *Server) DeleteCategory(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		category  = c.Get("category").(model.CategoryRaw)
	)

	err := s.Dependences.CategorySrvDeleteCategory(customCtx.GetRequestCtx(), category)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}
	return customCtx.Response200(nil, "")
}

// CategoryGetByID ...
func (s *Server) CategoryGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("categoryID")
		if id == "" {
			return next(c)
		}
		categoryID := util.GetAppIDFromHex(id)
		if categoryID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		category, err := s.Dependences.CategorySrvFindByID(customCtx.GetRequestCtx(), categoryID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("category", category)
		return next(c)
	}
}
