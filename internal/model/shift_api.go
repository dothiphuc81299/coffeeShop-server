package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShiftBody struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type ShiftResponse struct {
	ID      AppID        `json:"_id"`
	Name    string       `json:"name"`
	Date    TimeResponse `json:"date"`
	IsCheck bool         `json:"isCheck"`
	Staff   StaffInfo    `json:"staff"`
}

type ShiftAdminResponse struct {
	ID      AppID        `json:"_id"`
	Name    string       `json:"name"`
	Date    TimeResponse `json:"date"`
	IsCheck bool         `json:"isCheck"`
}

func (s ShiftBody) Validate() error {

	var shiftType = []interface{}{
		config.ShiftTypeOne,
		config.ShiftTypeTwo,
		config.ShiftTypeThree,
	}

	return validation.ValidateStruct(&s,
		validation.Field(&s.Name,
			validation.Required.Error(locale.ShiftKeyNameIsRequired),
			validation.In(shiftType...).Error(locale.ShiftKeyNameInvalid)),

		validation.Field(&s.Date,
			validation.Required.Error(locale.ShiftKeyDateIsRequired),
			validation.Date("2006-01-02T15:00:00.000Z").Error(locale.ShiftKeyDateInvalid)),
	)

}

func (s *ShiftBody) NewShiftRaw(staff AppID) ShiftRaw {
	shift := util.TimeParseISODate(s.Date)
	return ShiftRaw{
		ID:        primitive.NewObjectID(),
		Name:      s.Name,
		Date:      shift,
		IsCheck:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Staff:     staff,
	}
}

func (s *ShiftRaw) GetResponse(Staff StaffInfo) ShiftResponse {
	return ShiftResponse{
		ID:      s.ID,
		Name:    s.Name,
		Date:    TimeResponse{Time: s.Date},
		Staff:   Staff,
		IsCheck: s.IsCheck,
	}
}

func (s *StaffRaw) GetStaffInfo() StaffInfo {
	return StaffInfo{
		ID:       s.ID,
		Username: s.Username,
		Address:  s.Address,
		Phone:    s.Phone,
	}
}
