package service

import (
	"context"
	"fmt"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

type SalaryAdminService struct {
	ShiftDAO  model.ShiftDAO
	StaffDAO  model.StaffDAO
	StaffRole model.StaffRoleDAO
}

// NewSalaryAdminService ...
func NewSalaryAdminService(d *model.CommonDAO) model.SalaryAdminService {
	return &SalaryAdminService{
		ShiftDAO:  d.Shift,
		StaffDAO:  d.Staff,
		StaffRole: d.StaffRole,
	}
}

func (s *SalaryAdminService) GetMonth(cond bson.M, month string) {
	var (
		now      = time.Now()
		from, to time.Time
	)
	y, _, _ := now.Date()

	switch month {
	case config.SalaryMonthTypeJanuary:
		from = time.Date(y, 1, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 1+1, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeFebruary:
		from = time.Date(y, 2, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 3, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeMarch:
		from = time.Date(y, 3, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 4, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeApril:
		from = time.Date(y, 4, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 5, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeMay:
		from = time.Date(y, 5, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 6, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeJune:
		from = time.Date(y, 6, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 7, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeJuly:
		from = time.Date(y, 7, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 8, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeAugust:
		from = time.Date(y, 8, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 9, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeSeptember:
		from = time.Date(y, 9, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 10, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeOctober:
		from = time.Date(y, 10, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 11, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeNovember:
		from = time.Date(y, 10, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y, 11, 1, 0, 0, 0, 0, now.Location())
	case config.SalaryMonthTypeDecember:
		from = time.Date(y, 12, 1, 0, 0, 0, 0, now.Location())
		to = time.Date(y+1, 1, 1, 0, 0, 0, 0, now.Location())
	}

	to = util.TimeStartOfDayInHCM(to)
	from = util.TimeStartOfDayInHCM(from)
	cond["date"] = bson.M{
		"$gte": from,
		"$lt":  to,
	}
	fmt.Println("Cond : ", cond)
}

func (s *SalaryAdminService) GetDetail(ctx context.Context, salary model.SalaryBody, staff model.StaffRaw) (res model.SalaryResponse) {

	staffRes := model.StaffInfo{
		ID:       staff.ID,
		Username: staff.Username,
		Address:  staff.Address,
		Phone:    staff.Phone,
	}

	cond := bson.M{
		"isCheck": true,
		"staff":   staff.ID,
	}

	s.GetMonth(cond, salary.Month)

	res.Allowance = 200000

	totalShift := s.ShiftDAO.CountByCondition(ctx, cond)
	if totalShift < 10 {
		res.Allowance = 100000
	}
	res.Coefficient = 50000
	res.TotalShift = float64(totalShift)
	res.TotalSalary = float64(totalShift)*res.Coefficient + res.Allowance
	res.Month = salary.Month
	res.Staff = staffRes

	return res
}
