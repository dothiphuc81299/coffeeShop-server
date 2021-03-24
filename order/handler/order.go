package handler

import (
	"log"

	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
)

// OrderAdminHandler ...
type OrderAdminHandler struct {
	OrderAdminService model.OrderAdminService
}

// Update ...
func (h *OrderAdminHandler) Create(c echo.Context) error {
	var (
		cc     = util.EchoGetCustomCtx(c)
		body   = c.Get("orderBody").(model.OrderBody)
		userID = cc.GetCurrentUserID()
	//	user   = c.Get("user").(model.UserRaw)
	)

	log.Println("bosy", body)
	log.Println("userID2", userID)
	data, err := h.OrderAdminService.Create(cc.GetRequestCtx(), userID, body)
	log.Println("done")

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}
