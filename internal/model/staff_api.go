package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StaffMeResponse ...
type StaffMeResponse struct {
	ID       AppID  `json:"_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	//Avatar      *FilePhoto `json:"avatar"`
	Avatar      string   `json:"avatar"`
	Permissions []string `json:"permissions"`
}

type StaffInfo struct {
	ID       AppID  `json:"_id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

// StaffGetResponseAdmin ...
type StaffGetResponseAdmin struct {
	ID       AppID  `json:"_id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	//Avatar      *FilePhoto `json:"avatar"`
	Avatar      string       `json:"avatar"`
	Permissions []string     `json:"permissions"`
	CreatedAt   TimeResponse `json:"createdAt"`
	Active      bool         `json:"active"`
	Role        AppID        `json:"role"`
	IsRoot      bool         `json:"isRoot"`
}

// StaffResponse ...
type StaffResponse struct {
	ID       AppID  `json:"_id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	//Avatar      *FilePhoto `json:"avatar"`
	Avatar      string   `json:"avatar"`
	Permissions []string `json:"permissions"`
	Token       string   `json:"token"`
}

// StaffBody ...
type StaffBody struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type StaffUpdateRoleBody struct {
	Role string `json:"role"`
}

func (s StaffUpdateRoleBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Role, is.MongoID.Error(locale.CommonKeyIDMongoInvalid)),
	)
}

type StaffUpdateBodyByIt struct {
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type StaffLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordBody struct {
	Password         string `json:"password"`
	NewPassword      string `json:"newPassword"`
	NewPasswordAgain string `json:"newPasswordAgain"`
}

func (stf StaffUpdateBodyByIt) Validate() error {
	return validation.ValidateStruct(&stf,
		validation.Field(&stf.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&stf.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
	)
}

// Validate ...
func (stf StaffBody) Validate() error {
	return validation.ValidateStruct(&stf,
		validation.Field(&stf.Username, validation.Required.Error("username is required")),
		validation.Field(&stf.Phone),
		validation.Field(&stf.Address),
		validation.Field(&stf.Password, validation.Required.Error("Password is required")),
		validation.Field(&stf.Role,
			is.MongoID.Error(locale.CommonKeyIDMongoInvalid), validation.Required.Error("Role is Required")),
	)
}

// Validate ...
func (alg StaffLoginBody) Validate() error {
	return validation.ValidateStruct(&alg,
		validation.Field(&alg.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&alg.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
	)
}

func (a PasswordBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Password, validation.Required.Error("password duoc yeu cau")),
		validation.Field(&a.NewPassword, validation.Required.Error("password duoc yeu cau")),
		validation.Field(&a.NewPasswordAgain, validation.Required.Error("password duoc yeu cau")),
	)
}

const avt = "https://banner2.cleanpng.com/20180402/ojw/kisspng-united-states-avatar-organization-information-user-avatar-5ac20804a62b58.8673620215226654766806.jpg"

// StaffNewBSON ...
func (stf *StaffBody) StaffNewBSON(permissions []string) StaffRaw {
	roleID, _ := primitive.ObjectIDFromHex(stf.Role)
	now := time.Now()

	return StaffRaw{
		ID:          primitive.NewObjectID(),
		Password:    stf.Password,
		Username:    stf.Username,
		Phone:       stf.Phone,
		Address:     stf.Address,
		Active:      true,
		Avatar:      avt,
		Role:        roleID,
		CreatedAt:   now,
		UpdatedAt:   now,
		Permissions: permissions,
	}
}
