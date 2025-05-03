package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type QuestionAdminHandler struct {
	QuestionAdminService model.QuestionAdminService
}

// Create ...
func (d *QuestionAdminHandler) Create(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		QuestionBody = c.Get("questionBody").(model.QuestionBody)
	)

	err := d.QuestionAdminService.Create(customCtx.GetRequestCtx(), QuestionBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}
	return customCtx.Response200(result, "")
}

func (d *QuestionAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		Question  = c.Get("question").(model.QuestionRaw)
		body      = c.Get("body").(model.QuestionBodyUpdate)
	)

	err := d.QuestionAdminService.Update(customCtx.GetRequestCtx(), Question, body)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: "ok",
	}

	return customCtx.Response200(result, "")
}

func (d *QuestionAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		question  = c.Get("question").(model.QuestionRaw)
	)

	data, err := d.QuestionAdminService.ChangeStatus(customCtx.GetRequestCtx(), question)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

func (d *QuestionAdminHandler) ListAll(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Keyword: c.QueryParam("keyword"),
			Active:  c.QueryParam("active"),
			Limit:   customCtx.GetLimitQuery(),
			Page:    customCtx.GetPageQuery(),
			Sort: bson.D{
				bson.E{
					"order", 1},
			},
		}
	)
	data, total := d.QuestionAdminService.ListAll(customCtx.GetRequestCtx(), query)

	result := model.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return customCtx.Response200(result, "")
}

func (d *QuestionAdminHandler) QuestionGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("questionID")
		if id == "" {
			return next(c)
		}
		questionID := util.GetAppIDFromHex(id)
		if questionID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		question, err := d.QuestionAdminService.FindByID(customCtx.GetRequestCtx(), questionID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("question", question)
		return next(c)
	}
}

func (d *QuestionAdminHandler) GetDetail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		question = c.Get("question").(model.QuestionRaw)
	)

	data := d.QuestionAdminService.GetDetail(customCtx.GetRequestCtx(), question)
	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}
