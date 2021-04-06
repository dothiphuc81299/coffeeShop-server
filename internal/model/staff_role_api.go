package model

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// StaffRoleAdminResponse ...
type StaffRoleAdminResponse struct {
	ID          AppID        `json:"_id"`
	Name        string       `json:"name"`
	CreatedAt   TimeResponse `json:"createdAt"`
	Permissions []string     `json:"permissions"`
}

// StaffRoleBody ...
type StaffRoleBody struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// Validate ...
func (sr *StaffRoleBody) Validate() error {
	return validation.ValidateStruct(sr,
		validation.Field(&sr.Name, validation.Required.Error(locale.CommonKeyNameIsRequired)),
		validation.Field(&sr.Permissions, validation.Required.Error(locale.CommonKeyPermissionIsRequired)),
	)
}
