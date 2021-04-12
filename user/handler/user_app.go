package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

type UserAppHandler struct {
	UserAppService model.UserAppService
}

func (u *UserAppHandler) UserSignUp(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(model.UserSignUpBody)
	)

	data, err := u.UserAppService.UserSignUp(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (u *UserAppHandler) UserLoginIn(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(model.UserLoginBody)
	)

	data, err := u.UserAppService.UserLoginIn(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (u *UserAppHandler) UserUpdateAccount(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(model.UserSignUpBody)
		user = c.Get("user").(model.UserRaw)
	)

	data, err := u.UserAppService.UserUpdateAccount(cc.GetRequestCtx(), user, body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (u *UserAppHandler) GetDetailUser(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		user = c.Get("user").(model.UserRaw)
	)

	data := u.UserAppService.GetDetailUser(cc.GetRequestCtx(), user)
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}