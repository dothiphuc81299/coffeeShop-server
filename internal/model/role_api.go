package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// RoleBody ...
type RoleBody struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// Validate ...
func (r *RoleBody) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required.Error(locale.CommonKeyNameIsRequired)),
		validation.Field(&r.Permissions, validation.Required.Error(locale.CommonKeyPermissionIsRequired)),
	)
}

// RoleAdminResponse ...
type RoleAdminResponse struct {
	ID          AppID     `json:"_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	Permissions []string  `json:"permissions"`
}
