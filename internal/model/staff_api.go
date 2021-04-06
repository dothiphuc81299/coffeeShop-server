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
	ID          AppID      `json:"_id"`
	Username    string     `json:"username"`
	Token       string     `json:"token"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
	Avatar      *FilePhoto `json:"avatar"`
	Permissions []string   `json:"permissions"`
}

// StaffGetResponseAdmin ...
type StaffGetResponseAdmin struct {
	ID          AppID      `json:"_id"`
	Username    string     `json:"username"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
	Avatar      *FilePhoto `json:"avatar"`
	Permissions []string   `json:"permissions"`
	CreatedAt   time.Time  `json:"createdAt"`
	Active      bool       `json:"active"`
	Role        AppID      `json:"role"`
	IsRoot      bool       `json:"isRoot"`
}

// StaffBody ...
type StaffBody struct {
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Avatar      *FilePhoto `json:"avatar"`
	Role        string     `json:"role"`
	Permissions []string   `json:"permissions"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
}

// Validate ...
func (stf StaffBody) Validate() error {
	return validation.ValidateStruct(&stf,
		validation.Field(&stf.Username, validation.Required.Error(locale.CommonKeyUsernameIsRequired)),
		validation.Field(&stf.Phone, validation.Required.Error(locale.CommonKeyPhoneIsRequired)),
		validation.Field(&stf.Password, validation.Required.Error(locale.CommonKeyPasswordRequired)),
		validation.Field(&stf.Address, validation.Required.Error(locale.CommonKeyContactAddressIsRequired)),
		validation.Field(&stf.Role,
			validation.Required.Error(locale.CommonKeyPermissionIsRequired),
			is.MongoID.Error(locale.CommonKeyIDMongoInvalid)),
	)
}

// StaffNewBSON ...
func (stf *StaffBody) StaffNewBSON() StaffRaw {
	roleID, _ := primitive.ObjectIDFromHex(stf.Role)
	now := time.Now()
	var permissions = make([]string, 0)
	if len(stf.Permissions) > 0 {
		permissions = stf.Permissions
	}
	avatar := FileDefaultPhoto()
	if stf.Avatar != nil {
		avatar = stf.Avatar.GetResponseData()
	}
	return StaffRaw{
		ID:          primitive.NewObjectID(),
		Password:    stf.Password,
		Username:    stf.Username,
		Phone:       stf.Phone,
		Address:     stf.Address,
		Role:        roleID,
		Avatar:      avatar,
		CreatedAt:   now,
		UpdatedAt:   now,
		Permissions: permissions,
	}
}
