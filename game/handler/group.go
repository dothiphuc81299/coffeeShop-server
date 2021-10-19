package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

type GroupAdminHandler struct {
	GroupAdminService model.GroupAdminService
}

// Create ...
func (d *GroupAdminHandler) Create(c echo.Context) error {
	var (
		customCtx     = util.EchoGetCustomCtx(c)
		QuizGroupBody = c.Get("body").(model.QuizGroupBody)
	)

	err := d.GroupAdminService.Create(customCtx.GetRequestCtx(), QuizGroupBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}
	return customCtx.Response200(result, "")
}

func (d *GroupAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		QuizGroup = c.Get("quizGroup").(model.QuizGroupRaw)
		body      = c.Get("body").(model.QuizGroupBody)
	)

	err := d.GroupAdminService.Update(customCtx.GetRequestCtx(), QuizGroup, body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}

	return customCtx.Response200(result, "")
}

func (d *GroupAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		QuizGroup = c.Get("quizGroup").(model.QuizGroupRaw)
	)

	data, err := d.GroupAdminService.ChangeStatus(customCtx.GetRequestCtx(), QuizGroup)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

func (d *GroupAdminHandler) ListAll(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Keyword: c.QueryParam("keyword"),
			Active:  c.QueryParam("active"),
			Limit:   customCtx.GetLimitQuery(),
			Page:    customCtx.GetPageQuery(),
		}
	)

	data, total := d.GroupAdminService.ListAll(customCtx.GetRequestCtx(), query)

	result := model.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return customCtx.Response200(result, "")
}

func (d *GroupAdminHandler) GroupGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("quizGroupID")
		if id == "" {
			return next(c)
		}
		quizGroupID := util.GetAppIDFromHex(id)
		if quizGroupID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		quizGroup, err := d.GroupAdminService.FindByID(customCtx.GetRequestCtx(), quizGroupID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("quizGroup", quizGroup)
		return next(c)
	}
}
