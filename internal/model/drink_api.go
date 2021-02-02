package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DrinkBody ...
type DrinkBody struct {
	Name string `json:"name"`
	//CategoryID string  `json:"categoryID"`
	Price float64 `json:"price"`
	//	Photo      *FilePhoto `json:"photo"`
}

// DrinkAdminResponse ...
type DrinkAdminResponse struct {
	ID   AppID  `json:"_id"`
	Name string `json:"name"`
	//	Category CategoryInfo `json:"category"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Validate ....
func (d DrinkBody) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required.Error(locale.DrinkKeyNameIsRequired)),
		//	validation.Field(&d.CategoryID, validation.Required.Error(locale.DrinkKeyCategoryIDIsRequired),
		//	is.MongoID.Error(locale.DrinkKeyCategoryInvalid)),
		validation.Field(&d.Price, validation.Required.Error(locale.DrinkKeyPriceIsRequired)),
	)
}

// NewDrinkRaw ...
func (d DrinkBody) NewDrinkRaw() DrinkRaw {
	//categoryID, _ := primitive.ObjectIDFromHex(d.CategoryID)
	now := time.Now()
	return DrinkRaw{
		ID:           NewAppID(),
		Name:         d.Name,
		SearchString: format.NonAccentVietnamese(d.Name),
		Price:        d.Price,
		//CategoryID:   categoryID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
