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

// UserSignUpBody ...
type UserSignUpBody struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Phone    string            `json:"phone"`
	Avatar   *FilePhotoRequest `json:"avatar,omitempty"`
	Address  string            `json:"address"`
}

// UserUpdateBody ...
type UserUpdateBody struct {
	Username string            `json:"username"`
	Phone    string            `json:"phone"`
	Avatar   *FilePhotoRequest `json:"avatar,omitempty"`
	Address  string            `json:"address"`
}

// UserAdminResponse ...
type UserAdminResponse struct {
	ID        AppID      `json:"_id"`
	UserName  string     `json:"username"`
	Phone     string     `json:"phone"`
	Active    bool       `json:"active"`
	Avatar    *FilePhoto `json:"avatar"`
	CreatedAt time.Time  `json:"createdAt"`
	Address   string     `json:"address"`
}

type UserLoginResponse struct {
	ID        AppID      `json:"_id"`
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	Active    bool       `json:"active"`
	Avatar    *FilePhoto `json:"avatar"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Address   string     `json:"address"`
	Token     string     `json:"token"`
}

// Validate ...
func (alg UserLoginBody) Validate() error {
	return validation.ValidateStruct(&alg,
		validation.Field(&alg.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&alg.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
	)
}

// Validate ...
func (u UserSignUpBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&u.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&u.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
		validation.Field(&u.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
	)
}

// Validate ...
func (u UserUpdateBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&u.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&u.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
	)
}

// NewUserRaw																																													 ...
func (u *UserSignUpBody) NewUserRaw() UserRaw {
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

func (u *UserRaw) GetUserLoginInResponse(token string) UserLoginResponse {
	return UserLoginResponse{
		ID:        u.ID,
		Username:  u.Username,
		Phone:     u.Phone,
		Address:   u.Address,
		Avatar:    u.Avatar.GetResponseData(),
		CreatedAt: u.CreatedAt,
		Token:     token,
		Active:    u.Active,
	}
}
