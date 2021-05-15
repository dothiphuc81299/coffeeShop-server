package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DrinkBody ...
type DrinkBody struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	//	Photo    *FilePhoto `json:"photo"`
	Image string `json:"image"`
}

// DrinkAdminResponse ...
type DrinkAdminResponse struct {
	ID       AppID        `json:"_id"`
	Name     string       `json:"name"`
	Category CategoryInfo `json:"category"`
	Price    float64      `json:"price"`
	//	Photo    *FilePhoto   `json:"photo,omitempty"`
	Image  string `json:"image"`
	Active bool   `json:"active"`
}

type CategoryInfo struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}

// Validate ....
func (d DrinkBody) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required.Error(locale.DrinkKeyNameIsRequired)),
		validation.Field(&d.Category, validation.Required.Error(locale.DrinkKeyCategoryIDIsRequired),
			is.MongoID.Error(locale.DrinkKeyCategoryInvalid)),
		validation.Field(&d.Price, validation.Required.Error(locale.DrinkKeyPriceIsRequired)),
	)
}

// NewDrinkRaw ...
func (d DrinkBody) NewDrinkRaw() DrinkRaw {
	categoryID, _ := primitive.ObjectIDFromHex(d.Category)
	now := time.Now()
	return DrinkRaw{
		ID:           NewAppID(),
		Name:         d.Name,
		SearchString: format.NonAccentVietnamese(d.Name),
		Price:        d.Price,
		Category:     categoryID,
		CreatedAt:    now,
		UpdatedAt:    now,
		Active:       true,
		Image:        d.Image,
	}
}

func (b DrinkRaw) DrinkGetAdminResponse(c CategoryInfo) DrinkAdminResponse {
	return DrinkAdminResponse{
		ID:       b.ID,
		Name:     b.Name,
		Category: c,
		Price:    b.Price,
		//Photo:    b.Photo,
		Image:  b.Image,
		Active: b.Active,
	}
}

func CategoryGetInfo(r CategoryRaw) CategoryInfo {
	return CategoryInfo{
		ID:   r.ID,
		Name: r.Name,
	}
}
