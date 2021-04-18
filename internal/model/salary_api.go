package model

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SalaryResponse struct {
	Staff       StaffInfo `json:"staff"`
	Month       string    `json:"month"`
	TotalShift  float64   `json:"totalShift"`
	Coefficient float64   `json:"coefficient"`
	TotalSalary float64   `json:"totalSalary"`
	Allowance   float64   `json:"allowance"`
}

type SalaryBody struct {
	Month string `json:"month"`
}

func (s SalaryBody) Validate() error {
	typeMonth := []interface{}{
		config.SalaryMonthTypeJanuary,
		config.SalaryMonthTypeFebruary,
		config.SalaryMonthTypeMarch,
		config.SalaryMonthTypeApril,
		config.SalaryMonthTypeMay,
		config.SalaryMonthTypeJune,
		config.SalaryMonthTypeJuly,
		config.SalaryMonthTypeAugust,
		config.SalaryMonthTypeSeptember,
		config.SalaryMonthTypeOctober,
		config.SalaryMonthTypeNovember,
		config.SalaryMonthTypeDecember,
	}
	return validation.ValidateStruct(&s,
		validation.Field(&s.Month, validation.In(typeMonth...)),
	)
}
