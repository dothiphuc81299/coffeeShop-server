package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// SalaryAdminHandler ...
type SalaryAdminHandler struct {
	SalaryAdminService model.SalaryAdminService
}

func (s *SalaryAdminHandler) GetList(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)
	)

	data := s.SalaryAdminService.GetList(cc.GetRequestCtx())

	return cc.Response200(echo.Map{
		"salaries": data,
	}, "")
}
