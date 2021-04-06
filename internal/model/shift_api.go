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

type ShiftAdminResponse struct {
	ID      AppID        `json:"_id"`
	Name    string       `json:"name"`
	Date    TimeResponse `json:"date"`
	IsCheck bool         `json:"isCheck"`
	User    UserInfo     `json:"user"`
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

func (s *ShiftBody) NewShiftRaw(user AppID) ShiftRaw {
	shift := util.TimeParseISODate(s.Date)
	return ShiftRaw{
		ID:        primitive.NewObjectID(),
		Name:      s.Name,
		Date:      shift,
		IsCheck:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User:      user,
	}
}

func (s *ShiftRaw) GetResponse(user UserInfo) ShiftAdminResponse {
	return ShiftAdminResponse{
		ID:      s.ID,
		Name:    s.Name,
		Date:    TimeResponse{Time: s.Date},
		User:    user,
		IsCheck: s.IsCheck,
	}
}
