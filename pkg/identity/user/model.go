package user

import (
	"errors"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/format"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/password"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmailExisted     = errors.New("Email existed")
	ErrUserNotFound     = errors.New("User not found")
	ErrPasswordInvalid  = errors.New("Password is incorrect")
	ErrAccountIsInvalid = errors.New("Account is invalid")
)

type UserRaw struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	Phone        string             `bson:"phone" json:"phone"`
	Active       bool               `bson:"active" json:"active"`
	Avatar       string             `bson:"avatar" json:"avatar"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
	Address      string             `bson:"address" json:"address"`
	SearchString string             `bson:"searchString" json:"searchString"`
	CurrentPoint float64            `bson:"currentPoint" json:"currentPoint"`
	Email        string             `bson:"email" json:"email"`
}

type CreateLoginUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar,omitempty"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type SendUserEmailCommand struct {
	Email string `json:"email"`
}

type VerifyEmailCommand struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UpdateUserCommand struct {
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UserAdminResponse struct {
	ID           primitive.ObjectID `json:"_id"`
	UserName     string             `json:"username"`
	Phone        string             `json:"phone"`
	Active       bool               `json:"active"`
	Avatar       string             `json:"avatar"`
	CreatedAt    time.Time          `json:"createdAt"`
	Address      string             `json:"address"`
	CurrentPoint float64            `json:"currentPoint"`
	Email        string             `json:"email"`
}

type CreateLoginUserResult struct {
	ID           primitive.ObjectID `json:"_id"`
	Username     string             `json:"username"`
	Phone        string             `json:"phone"`
	Email        string             `json:"email"`
	Active       bool               `json:"active"`
	Avatar       string             `json:"avatar"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	Address      string             `json:"address"`
	Token        string             `json:"token"`
	Password     string             `json:"password"`
	CurrentPoint float64            `json:"currentPoint"`
}

type ChangePasswordUserCommand struct {
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

func (alg CreateLoginUserCommand) Validate() error {
	return validation.ValidateStruct(&alg,
		validation.Field(&alg.Username, validation.Required),
		validation.Field(&alg.Password, validation.Required),
	)
}

func (u CreateUserCommand) Validate() error {
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Phone, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Address, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email))

	if err != nil {
		return err
	}

	return nil
}

func (u UpdateUserCommand) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Phone, validation.Required),
		validation.Field(&u.Address, validation.Required),
	)
}

func (a ChangePasswordUserCommand) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Password, validation.Required),
		validation.Field(&a.NewPassword, validation.Required),
		validation.Field(&a.NewPasswordAgain, validation.Required),
	)
}

func (u SendUserEmailCommand) Validate() error {
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email))

	if err != nil {
		return err
	}
	return nil
}

func (u VerifyEmailCommand) Validate() error {
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Code, validation.Required))
	if err != nil {
		return err
	}
	return nil
}

func (u *CreateUserCommand) NewUserRaw() UserRaw {
	now := time.Now().UTC()
	hashPassword, _ := password.HashPassword(u.Password)
	return UserRaw{
		ID:           primitive.NewObjectID(),
		Username:     u.Username,
		Password:     hashPassword,
		Email:        u.Email,
		Active:       true,
		Phone:        u.Phone,
		CreatedAt:    now,
		UpdatedAt:    now,
		Address:      u.Address,
		CurrentPoint: 0,
		SearchString: format.NonAccentVietnamese(u.Username),
	}
}

func (u *UserRaw) GetLoginUserResponse(token string) CreateLoginUserResult {
	return CreateLoginUserResult{
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
