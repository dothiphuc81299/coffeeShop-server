package handler

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// FeedbackAppHandler ...
type FeedbackAppHandler struct {
	FeedbackAppService model.FeedbackAppService
}

// Create ...
func (d *FeedbackAppHandler) Create(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		FeedbackBody = c.Get("feedbackBody").(model.FeedbackBody)
		user         = c.Get("user").(model.UserRaw)
	)

	data, err := d.FeedbackAppService.Create(customCtx.GetRequestCtx(), FeedbackBody, user)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(echo.Map{
		"data": data,
	}, "")
}

func (d *FeedbackAppHandler) GetList(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
	)

	data, total := d.FeedbackAppService.ListAll(context.Background())

	result := model.ResponseAppListData{
		Data:  data,
		Total: total,
	}
	return customCtx.Response200(result, "")
}

func (d *FeedbackAppHandler) Update(c echo.Context) error {
	var (
		customCtx    = util.EchoGetCustomCtx(c)
		FeedbackBody = c.Get("feedbackBody").(model.FeedbackBody)
		Feedback     = c.Get("Feedback").(model.FeedbackRaw)
		user         = c.Get("user").(model.UserRaw)
	)

	data, err := d.FeedbackAppService.Update(customCtx.GetRequestCtx(), FeedbackBody, user, Feedback)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(echo.Map{
		"feedback": data,
	}, "")
}

func (d *FeedbackAppHandler) GetDetail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		Feedback  = c.Get("Feedback").(model.FeedbackRaw)
	)

	data := d.FeedbackAppService.GetDetail(customCtx.GetRequestCtx(), Feedback)

	return customCtx.Response200(echo.Map{
		"feedback": data,
	}, "")

}

// FeedbackGetByID ...
func (d *FeedbackAppHandler) FeedbackGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("feedbackID")
		if id == "" {
			return next(c)
		}
		FeedbackID := util.GetAppIDFromHex(id)
		if FeedbackID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		Feedback, err := d.FeedbackAppService.FindByID(customCtx.GetRequestCtx(), FeedbackID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("Feedback", Feedback)
		return next(c)
	}
}
