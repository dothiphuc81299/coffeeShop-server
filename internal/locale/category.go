package locale

import (
	"log"

	"github.com/magiconair/properties"
)

// Keys ...
const (
	CategoryKeyNameIsRequired = "nameIsRequired"
	CategoryKeyNameExisted    = "NameExisted"

	EventKeyNameIsRequired = "EventNameIsRequired"
	EventKeyDescIsRequired = "EventDescIsRequired"
	EventKeyCanNotUpdate   = "EventCanNotUpdate"

	ShiftKeyNameIsRequired = "ShiftKeyNameIsRequired"
	ShiftKeyNameInvalid    = "ShiftKeyNameInvalid"
	ShiftKeyDateIsRequired = "ShiftKeyDateIsRequired"
	ShiftKeyDateInvalid    = "ShiftKeyDateInvalid"
)

type (
	categoryLang struct {
		NameIsRequired string `properties:"nameIsRequired"`
		NameExisted    string `properties:"nameExisted"`

		EventNameIsRequired string `properties:"eventNameIsRequired"`
		EventDescIsRequired string `properties:"eventDescIsRequired"`
		EventCanNotUpdate   string `properties:"eventCanNotUpdate"`

		ShiftKeyNameIsRequired string `prooperties:"shiftNameIsRequired"`
		ShiftKeyNameInvalid    string `properties:"shiftNameInvalid"`
		ShiftKeyDateIsRequired string `properties:"shiftDateIsRequired"`
		ShiftKeyDateInvalid    string `properties:"shiftDateInvalid"`
	}
)

var (
	categoryVi categoryLang
)

func init() {
	// Load properties
	p2 := properties.MustLoadFile(getLocalePath()+"/properties/category.properties", properties.UTF8)
	if err := p2.Decode(&categoryVi); err != nil {
		log.Fatal(err)
	}
}

func categoryLoadLocales() (response []Locale) {
	// 200-299
	response = []Locale{
		{
			Key: CategoryKeyNameIsRequired,
			Message: &Message{
				Vi: categoryVi.NameIsRequired,
			},
			Code: 200,
		},
		{
			Key: CategoryKeyNameExisted,
			Message: &Message{
				Vi: categoryVi.NameExisted,
			},
			Code: 201,
		},

		{
			Key: EventKeyNameIsRequired,
			Message: &Message{
				Vi: categoryVi.EventNameIsRequired,
			},
			Code: 202,
		},

		{
			Key: EventKeyDescIsRequired,
			Message: &Message{
				Vi: categoryVi.EventDescIsRequired,
			},
			Code: 203,
		},

		{
			Key: EventKeyCanNotUpdate,
			Message: &Message{
				Vi: categoryVi.EventCanNotUpdate,
			},
			Code: 204,
		},

		{
			Key: ShiftKeyNameIsRequired,
			Message: &Message{
				Vi: categoryVi.ShiftKeyNameIsRequired,
			},
			Code: 205,
		},

		{
			Key: ShiftKeyNameInvalid,
			Message: &Message{
				Vi: categoryVi.ShiftKeyNameInvalid,
			},
			Code: 206,
		},

		{
			Key: ShiftKeyDateIsRequired,
			Message: &Message{
				Vi: categoryVi.ShiftKeyDateIsRequired,
			},
			Code: 207,
		},

		{
			Key: ShiftKeyDateInvalid,
			Message: &Message{
				Vi: categoryVi.ShiftKeyDateInvalid,
			},
			Code: 208,
		},
	}
	return response
}
