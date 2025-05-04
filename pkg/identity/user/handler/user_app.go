 package handler

// import (
// 	"fmt"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
// 	"github.com/labstack/echo/v4"
// )

// type UserAppHandler struct {
// 	UserAppService model.UserAppService
// }

// func (u *UserAppHandler) CreateUser(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.CreateUserCommand)
// 	)

// 	email, err := u.UserAppService.CreateUser(cc.GetRequestCtx(), body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(echo.Map{
// 		"data": email,
// 	}, "")
// }

// func (u *UserAppHandler) SendEmail(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.SendUserEmailCommand)
// 	)

// 	err := u.UserAppService.SendEmail(cc.GetRequestCtx(), body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(nil, "")
// }

// func (u *UserAppHandler) VerifyEmail(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.VerifyEmailCommand)
// 	)

// 	err := u.UserAppService.VerifyEmail(cc.GetRequestCtx(), body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(nil, "")
// }

// func (u *UserAppHandler) LoginUser(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.CreateLoginUserCommand)
// 	)

// 	data, err := u.UserAppService.LoginUser(cc.GetRequestCtx(), body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(echo.Map{
// 		"data": data,
// 	}, "")
// }

// func (u *UserAppHandler) UpdateUser(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.UpdateUserCommand)
// 		user = c.Get("user").(model.UserRaw)
// 	)

// 	err := u.UserAppService.UpdateUser(cc.GetRequestCtx(), user, body)
// 	fmt.Println(err)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(nil, "")
// }

// func (u *UserAppHandler) GetDetailUser(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		user = c.Get("user").(model.UserRaw)
// 	)

// 	data := u.UserAppService.GetDetailUser(cc.GetRequestCtx(), user)
// 	return cc.Response200(echo.Map{
// 		"data": data,
// 	}, "")
// }

// func (u *UserAppHandler) ChangePassword(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		user = c.Get("user").(model.UserRaw)
// 		body = c.Get("body").(model.ChangePasswordUserCommand)
// 	)

// 	err := u.UserAppService.ChangePassword(cc.GetRequestCtx(), user, body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(echo.Map{}, "")
// }
