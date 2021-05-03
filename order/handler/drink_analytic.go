package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// DrinkAnalyticHandler ...
type DrinkAnalyticHandler struct {
	DrinkAnalyticService model.DrinkAnalyticService
}

func (h *DrinkAnalyticHandler) GetList(c echo.Context) error {
	var (
		cc      = util.EchoGetCustomCtx(c)
		startAt = util.TimeParseISODate(cc.QueryParam("startAt"))
		endAt   = util.TimeParseISODate(cc.QueryParam("endAt"))
		query   = model.CommonQuery{
			StartAt: startAt,
			EndAt:   endAt,
		}
	)
	data := h.DrinkAnalyticService.ListAll(cc.GetRequestCtx(), query)

	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}
