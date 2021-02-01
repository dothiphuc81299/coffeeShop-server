package locale

import (
	"log"

	"github.com/magiconair/properties"
)

// Keys ..
const (
	DrinkKeyNameIsRequired       = "nameIsRequired"
	DrinkKeyCategoryIDIsRequired = "categoryIDIsRequired"
	DrinkKeyCategoryInvalid      = "categoryInvalid"
	DrinkKeyPriceIsRequired      = "priceIsRequired"
)

type (
	drinkLang struct {
		NameIsRequired       string `properties:"nameIsRequired"`
		CategoryIDIsRequired string `properties:"categoryIDIsRequired"`
		CategoryInvalid      string `properties:"categoryInvalid"`
		PriceIsRequired      string `properties:"priceIsRequired"`
	}
)

var (
	drinkVi drinkLang
)

func init() {
	// Load properties
	p2 := properties.MustLoadFile(getLocalePath()+"/properties/drink.properties", properties.UTF8)
	if err := p2.Decode(&drinkVi); err != nil {
		log.Fatal(err)
	}
}

func drinkLoadLocales() (response []Locale) {
	// 100-199
	response = []Locale{
		{
			Key: DrinkKeyNameIsRequired,
			Message: &Message{
				Vi: drinkVi.NameIsRequired,
			},
			Code: 101,
		},
		{
			Key: DrinkKeyCategoryIDIsRequired,
			Message: &Message{
				Vi: drinkVi.CategoryIDIsRequired,
			},
			Code: 102,
		},
		{
			Key: DrinkKeyCategoryInvalid,
			Message: &Message{
				Vi: drinkVi.CategoryInvalid,
			},
			Code: 103,
		},
		{
			Key: DrinkKeyPriceIsRequired,
			Message: &Message{
				Vi: drinkVi.PriceIsRequired,
			},
			Code: 104,
		},
	}
	return response
}
