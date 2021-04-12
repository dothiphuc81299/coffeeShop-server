package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// FeedbackAdminHandler ...
type FeedbackAdminHandler struct {
	FeedbackAdminService model.FeedbackAdminService
}

func (f *FeedbackAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		Feedback  = c.Get("Feedback").(model.FeedbackRaw)
	)

	data, err := f.FeedbackAdminService.ChangeStatus(customCtx.GetRequestCtx(), Feedback)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(echo.Map{
		"status": data,
	}, "")
}

// FeedbackGetByID ...
func (d *FeedbackAdminHandler) FeedbackGetByID(next echo.HandlerFunc) echo.HandlerFunc {
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
		Feedback, err := d.FeedbackAdminService.FindByID(customCtx.GetRequestCtx(), FeedbackID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("Feedback", Feedback)
		return next(c)
	}
}
