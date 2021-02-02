package locale

import (
	"log"

	"github.com/magiconair/properties"
)

// Keys ...
const (
	CategoryKeyNameIsRequired = "nameIsRequired"
)

type (
	categoryLang struct {
		NameIsRequired string `properties:"nameIsRequired"`
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
	}
	return response
}
