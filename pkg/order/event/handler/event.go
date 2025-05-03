package handler

import (
	"context"
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// EventAdminHandler ...
type EventAdminHandler struct {
	EventAdminService model.EventAdminService
}

// Create ...
func (d *EventAdminHandler) Create(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		EventBody = c.Get("EventBody").(model.EventBody)
	)

	data, err := d.EventAdminService.Create(customCtx.GetRequestCtx(), EventBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}
	return customCtx.Response200(result, "")
}

// Update ...
func (d *EventAdminHandler) Update(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		EventBody = c.Get("EventBody").(model.EventBody)
		Event     = c.Get("Event").(model.EventRaw)
	)
	fmt.Println(1)
	data, err := d.EventAdminService.Update(customCtx.GetRequestCtx(), Event, EventBody)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	result := model.ResponseAdminData{
		Data: data,
	}

	return customCtx.Response200(result, "")
}

// ListAll ...
func (d *EventAdminHandler) ListAll(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		query     = model.CommonQuery{
			Active: c.QueryParam("active"),
			Limit:  customCtx.GetLimitQuery(),
			Page:   customCtx.GetPageQuery(),
		}
	)

	data, total := d.EventAdminService.ListAll(context.Background(), query)

	result := model.ResponseAdminListData{
		Data:         data,
		Total:        total,
		LimitPerPage: query.Limit,
	}
	return customCtx.Response200(result, "")
}

// ChangeStatus ...
func (d *EventAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		Event     = c.Get("Event").(model.EventRaw)
	)

	err := d.EventAdminService.ChangeStatus(customCtx.GetRequestCtx(), Event)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(nil, "")
}

func (d *EventAdminHandler) SendEmail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		Event     = c.Get("Event").(model.EventRaw)
	)

	err := d.EventAdminService.SendEmail(customCtx.GetRequestCtx(), Event)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(nil, "")
}

func (d *EventAdminHandler) GetDetail(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		event     = customCtx.Get("Event").(model.EventRaw)
	)

	data := d.EventAdminService.GetDetail(customCtx.GetRequestCtx(), event)
	return customCtx.Response200(echo.Map{
		"event": data,
	}, "")
}

func (d *EventAdminHandler) DeleteEvent(c echo.Context) error {
	var (
		customCtx = util.EchoGetCustomCtx(c)
		event     = customCtx.Get("Event").(model.EventRaw)
	)

	err := d.EventAdminService.DeleteEvent(customCtx.GetRequestCtx(), event)
	if err != nil {
		return customCtx.Response400(nil, err.Error())
	}

	return customCtx.Response200(nil, "")
}

// EventGetByID ...
func (d *EventAdminHandler) EventGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("eventID")
		if id == "" {
			return next(c)
		}
		EventID := util.GetAppIDFromHex(id)
		if EventID.IsZero() {
			return customCtx.Response400(nil, locale.CommonKeyBadRequest)
		}
		Event, err := d.EventAdminService.FindByID(customCtx.GetRequestCtx(), EventID)
		if err != nil {
			return customCtx.Response404(nil, locale.CommonKeyNotFound)
		}

		c.Set("Event", Event)
		return next(c)
	}
}
