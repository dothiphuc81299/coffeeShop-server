package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// SalaryAppHandler ...
type SalaryAppHandler struct {
	SalaryAppService model.SalaryAppService
}

func (h *SalaryAppHandler) GetDetail(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)
		// salary = c.Get("salaryBody").(model.SalaryBody)
		staff = c.Get("staff").(model.StaffRaw)
		query = model.CommonQuery{
			Month: cc.QueryParam("month"),
		}
	)

	data := h.SalaryAppService.GetDetail(cc.GetRequestCtx(), query, staff)

	return cc.Response200(echo.Map{
		"Salary": data,
	}, "")
}
