package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CategoryBody ...
type CategoryBody struct {
	Name string `json:"name"`
}

// CategoryAdminResponse ...
type CategoryAdminResponse struct {
	ID        AppID     `json:"_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Validate ...
func (c CategoryBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error(locale.CategoryKeyNameIsRequired)),
	)
}

// NewCategoryRaw ...
func (c CategoryBody) NewCategoryRaw() CategoryRaw {
	now := time.Now()
	return CategoryRaw{
		ID:           NewAppID(),
		Name:         c.Name,
		SearchString: format.NonAccentVietnamese(c.Name),
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func CategoryGetAdminResponse(c CategoryRaw) CategoryAdminResponse {
	return CategoryAdminResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
	}
}
