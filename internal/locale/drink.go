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
	DrinkKeyCanNotCreate         = "drinkCanNotCreate"
	DrinkKeyNameExisted          = "DrinkKeyNameExisted"
	OrderKeyNameDrinkIsRequired  = "OrderKeyNameDrinkIsRequired"
	OrderKeyNameDrinkInvalid     = "OrderKeyNameDrinkInvalid"
	OrderKeyQuantityIsRequired   = "OrderKeyQuantityIsRequired"
)

type (
	drinkLang struct {
		NameIsRequired           string `properties:"nameIsRequired"`
		CategoryIDIsRequired     string `properties:"categoryIDIsRequired"`
		CategoryInvalid          string `properties:"categoryInvalid"`
		PriceIsRequired          string `properties:"priceIsRequired"`
		DrinkCanNotCreate        string `properties:"drinkCanNotCreate"`
		DrinkNameExisted         string `properties:"drinkNameExisted"`
		OrderNameDrinkIsRequired string `properties:"orderNameDrinkIsRequired"`
		OrderNameDrinkInvalid    string `properties:"orderNameDrinkInvalid"`
		OrderQuantityIsRequired  string `properties:"orderQuantityIsRequired"`
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
		{
			Key: DrinkKeyCanNotCreate,
			Message: &Message{
				Vi: drinkVi.DrinkCanNotCreate,
			},
			Code: 105,
		},
		{
			Key: DrinkKeyNameExisted,
			Message: &Message{
				Vi: drinkVi.DrinkNameExisted,
			},
			Code: 106,
		},
		{
			Key: OrderKeyNameDrinkIsRequired,
			Message: &Message{
				Vi: drinkVi.OrderNameDrinkIsRequired,
			},
			Code: 107,
		},
		{
			Key: OrderKeyNameDrinkInvalid,
			Message: &Message{
				Vi: drinkVi.OrderNameDrinkInvalid,
			},
			Code: 108,
		},
		{
			Key: OrderKeyQuantityIsRequired,
			Message: &Message{
				Vi: drinkVi.OrderQuantityIsRequired,
			},
			Code: 109,
		},
	}
	return response
}
