package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UserLoginBody ...
type UserLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate ...
func (alg UserLoginBody) Validate() error {
	return validation.ValidateStruct(&alg,
		validation.Field(&alg.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&alg.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
	)
}

// UserBody ...
type UserBody struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Phone    string            `json:"phone"`
	Avatar   *FilePhotoRequest `json:"avatar,omitempty"`
	Address  string            `json:"address"`
}

// Validate ...
func (u UserBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&u.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&u.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
		validation.Field(&u.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
	)
}

// UserAdminResponse ...
type UserAdminResponse struct {
	ID        AppID      `json:"_id"`
	UserName  string     `json:"username"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Active    bool       `json:"active"`
	Avatar    *FilePhoto `json:"avatar"`
	IsRoot    bool       `json:"isRoot"`
	CreatedAt time.Time  `json:"createdAt"`
	Token     string     `json:"token,omitempty"`
	Address   string     `json:"address"`
}

// NewRaw ...
func (u *UserBody) NewRaw() UserRaw {
	now := time.Now()
	return UserRaw{
		ID:           NewAppID(),
		Username:     u.Username,
		Password:     u.Password,
		Active:       false,
		Phone:        u.Phone,
		Avatar:       u.Avatar.ConvertToFilePhoto(),
		CreatedAt:    now,
		UpdatedAt:    now,
		Address:      u.Address,
		SearchString: format.NonAccentVietnamese(u.Username),
	}
}
