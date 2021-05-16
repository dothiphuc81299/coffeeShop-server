package service

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
)

type SalaryAdminService struct {
	ShiftDAO  model.ShiftDAO
	StaffDAO  model.StaffDAO
	StaffRole model.StaffRoleDAO
	OrderDAO  model.OrderDAO
}

func NewSalaryAdminService(d *model.CommonDAO) model.SalaryAdminService {
	return &SalaryAdminService{
		ShiftDAO:  d.Shift,
		StaffDAO:  d.Staff,
		StaffRole: d.StaffRole,
		OrderDAO:  d.Order,
	}
}

func (s *SalaryAdminService) GetList(ctx context.Context) []model.SalaryResponse {
	panic("")
}
