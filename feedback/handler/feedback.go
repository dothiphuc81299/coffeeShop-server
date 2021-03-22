package handler

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// FeedbackAdminHandler ...
type FeedbackAdminHandler struct {
	FeedbackAdminService model.FeedbackAdminService
}

// Create ...
func (d *FeedbackAdminHandler) Create(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		FeedbackBody = c.Get("feedbackBody").(model.FeedbackBody)
		userID       = customCtx.GetCurrentUserID()
	)

	data, err := d.FeedbackAdminService.Create(customCtx.GetRequestCtx(), userID, FeedbackBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

func (d *FeedbackAdminHandler) GetList(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Keyword: c.QueryParam("keyword"),
			Active:  c.QueryParam("active"),
		}
	)

	data, total := d.FeedbackAdminService.ListAll(context.Background(), query)

	result := model.ResponseAdminListData{
		Data:  data,
		Total: total,
	}
	return customCtx.Response200(result, "")
}

func (d *FeedbackAdminHandler) Update(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		FeedbackBody = c.Get("FeedbackBody").(model.FeedbackBody)
		Feedback     = c.Get("Feedback").(model.FeedbackRaw)
		userID       = customCtx.GetCurrentUserID()
	)

	data, err := d.FeedbackAdminService.Update(customCtx.GetRequestCtx(), userID, Feedback, FeedbackBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

// FeedbackGetByID ...
func (d *FeedbackAdminHandler) FeedbackGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("FeedbackID")
		if id == "" {
			return next(c)
		}
		FeedbackID := util.GetAppIDFromHex(id)
		if FeedbackID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		Feedback, err := d.FeedbackAdminService.FindByID(customCtx.GetRequestCtx(), FeedbackID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("Feedback", Feedback)
		return next(c)
	}
}
