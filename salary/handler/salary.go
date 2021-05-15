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

func (h *SalaryAdminHandler) GetDetail(c echo.Context) error {
	var (
		cc     = util.EchoGetCustomCtx(c)
		salary = c.Get("salaryBody").(model.SalaryBody)
		staff  = c.Get("staff").(model.StaffRaw)
	)
	
	data := h.SalaryAdminService.GetDetail(cc.GetRequestCtx(), salary, staff)

	return cc.Response200(echo.Map{
		"Salary": data,
	}, "")
}
