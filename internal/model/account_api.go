package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccountBody ...
type AccountBody struct {
	User        string   `json:"user"`
	Type        string   `json:"type"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// Validate ...
func (a *AccountBody) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.User,
			validation.Required.Error(locale.AuthKeyUserIsRequired),
			is.MongoID.Error(locale.AuthKeyUserInvalid)),
		validation.Field(&a.Permissions,
			validation.Required.Error(locale.CommonKeyPermissionIsRequired)),
	)
}

// NewRaw ...
func (a *AccountBody) NewRaw() AccountRaw {
	now := time.Now()
	userID, _ := primitive.ObjectIDFromHex(a.User)
	roleID, _ := primitive.ObjectIDFromHex(a.Role)
	return AccountRaw{
		ID:          NewAppID(),
		User:        userID,
		Permissions: a.Permissions,
		Role:        roleID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// AccountResponse ...
type AccountResponse struct {
	ID          AppID     `json:"_id"`
	Active      bool      `json:"active"`
	User        AppID     `json:"user"`
	IsRoot      bool      `json:"isRoot"`
	Permissions []string  `json:"permissions"`
	Role        AppID     `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
}

// GetResponse ...
func (r *AccountRaw) GetResponse() AccountResponse {
	return AccountResponse{
		ID:          r.ID,
		Active:      r.Active,
		User:        r.User,
		Permissions: r.Permissions,
		Role:        r.Role,
		CreatedAt:   r.CreatedAt,
	}
}
