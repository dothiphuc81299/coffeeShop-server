package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UserLoginBody ...
type UserLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserSignUpBody ...
type UserSignUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar,omitempty"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type UserSendEmailBody struct {
	Email string `json:"email"`
}

type VerifyEmailBody struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// UserUpdateBody ...
type UserUpdateBody struct {
	Phone string `json:"phone"`
	// Avatar   *FilePhotoRequest `json:"avatar,omitempty"`
	Address string `json:"address"`
}

// UserAdminResponse ...
type UserAdminResponse struct {
	ID           AppID     `json:"_id"`
	UserName     string    `json:"username"`
	Phone        string    `json:"phone"`
	Active       bool      `json:"active"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"createdAt"`
	Address      string    `json:"address"`
	CurrentPoint float64   `json:"currentPoint"`
	Email        string    `json:"email"`
}

type UserLoginResponse struct {
	ID           AppID     `json:"_id"`
	Username     string    `json:"username"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Active       bool      `json:"active"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Address      string    `json:"address"`
	Token        string    `json:"token"`
	Password     string    `json:"password"`
	CurrentPoint float64   `json:"currentPoint"`
}

type UserChangePasswordBody struct {
	Password         string `json:"password"`
	NewPassword      string `json:"newPassword"`
	NewPasswordAgain string `json:"newPasswordAgain"`
}

type UserVerifyEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Message struct {
	To            []string
	From          string
	Template      string
	Subject       string
	Body          string
	Data          map[string]interface{}
	EmbeddedFiles []string
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
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&u.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&u.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
		validation.Field(&u.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
		validation.Field(&u.Email, validation.Required.Error(locale.CommonKeyEmailIsRequired), is.Email.Error(locale.CommonKeyEmailInvalid)))

	if err != nil {
		return err
	}

	return nil
}

// Validate ...
func (u UserUpdateBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&u.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
	)
}

func (a UserChangePasswordBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
		validation.Field(&a.NewPassword, validation.Required.Error(locale.CommonKeyPasswordRequired)),
		validation.Field(&a.NewPasswordAgain, validation.Required.Error(locale.CommonKeyPasswordRequired)),
	)
}

func (u UserSendEmailBody) Validate() error {
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required.Error(locale.CommonKeyEmailIsRequired), is.Email.Error(locale.CommonKeyEmailInvalid)))

	if err != nil {
		return err
	}
	return nil
}

func (u VerifyEmailBody) Validate() error {
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required.Error(locale.CommonKeyEmailIsRequired), is.Email.Error(locale.CommonKeyEmailInvalid)),
		validation.Field(&u.Code, validation.Required.Error(locale.CodeIsRequired)))
	if err != nil {
		return err
	}
	return nil
}

// NewUserRaw																																													 ...
func (u *UserSignUpBody) NewUserRaw() UserRaw {
	now := time.Now()
	return UserRaw{
		ID:           NewAppID(),
		Username:     u.Username,
		Password:     u.Password,
		Email:        u.Email,
		Active:       false,
		Phone:        u.Phone,
		Avatar:       avt,
		CreatedAt:    now,
		UpdatedAt:    now,
		Address:      u.Address,
		CurrentPoint: 0,
		SearchString: format.NonAccentVietnamese(u.Username),
	}
}

func (u *UserRaw) GetUserLoginInResponse(token string) UserLoginResponse {
	return UserLoginResponse{
		ID:           u.ID,
		Username:     u.Username,
		Phone:        u.Phone,
		Email:        u.Email,
		Address:      u.Address,
		Avatar:       u.Avatar,
		CreatedAt:    u.CreatedAt,
		Token:        token,
		Active:       u.Active,
		Password:     u.Password,
		CurrentPoint: u.CurrentPoint,
	}
}
