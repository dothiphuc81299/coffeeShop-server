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

type SalaryAppService struct {
	ShiftDAO  model.ShiftDAO
	StaffDAO  model.StaffDAO
	StaffRole model.StaffRoleDAO
	OrderDAO  model.OrderDAO
}

// NewSalaryAppService ...
func NewSalaryAppService(d *model.CommonDAO) model.SalaryAppService {
	return &SalaryAppService{
		ShiftDAO:  d.Shift,
		StaffDAO:  d.Staff,
		StaffRole: d.StaffRole,
		OrderDAO:  d.Order,
	}
}

func (s *SalaryAppService) getMonth(cond bson.M, date string, month string) {
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
	cond[date] = bson.M{
		"$gte": from,
		"$lt":  to,
	}
	fmt.Println("Cond : ", cond)
}

func (s *SalaryAppService) GetDetail(ctx context.Context, query model.CommonQuery, staff model.StaffRaw) (res model.SalaryResponse) {
	staffRes := model.StaffInfo{
		ID:       staff.ID,
		Username: staff.Username,
		Address:  staff.Address,
		Phone:    staff.Phone,
	}

	if staff.Username == "shipper" {
		cond := bson.M{
			"shipper": staff.ID,
		}
		s.getMonth(cond, "updatedAt", query.Month)
		total := s.OrderDAO.CountByCondition(ctx, cond)
		if total > 10 {
			res.Allowance = 10000
			res.Coefficient = 50000
			res.TotalShift = float64(total)
			res.TotalSalary = float64(total)*res.Coefficient + res.Allowance
		} else {
			res.Allowance = 0
			res.Coefficient = 50000
			res.TotalShift = float64(total)
			res.TotalSalary = float64(total)*res.Coefficient + res.Allowance
		}

		res.Month = query.Month
		res.Staff = staffRes

	} else {
		cond := bson.M{
			"isCheck": true,
			"staff":   staff.ID,
		}

		s.getMonth(cond, "date", query.Month)

		res.Allowance = 200000

		totalShift := s.ShiftDAO.CountByCondition(ctx, cond)
		if totalShift > 10 {
			res.Allowance = 100000
			res.Coefficient = 50000
			res.TotalShift = float64(totalShift)
			res.TotalSalary = float64(totalShift)*res.Coefficient + res.Allowance

		} else {
			res.Allowance = 0
			res.Coefficient = 50000
			res.TotalShift = float64(totalShift)
			res.TotalSalary = float64(totalShift)*res.Coefficient + res.Allowance
		}
		res.Month = query.Month
		res.Staff = staffRes

	}
	return
}
