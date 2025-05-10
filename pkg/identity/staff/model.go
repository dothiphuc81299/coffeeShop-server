package staff

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrStaffExisted                  = errors.New("Staff existed")
	ErrStaffNotFound                 = errors.New("Staff not found")
	ErrPasswordInvalid               = errors.New("Password invalid")
	ErrUserNameOrPasswordIsIncorrect = errors.New("User name or password is incorrect")
	ErrStaffIsDeleted                = errors.New("Staff is deleted")
	ErrCanNotUpdateRole              = errors.New("Can not update role")
)

type Staff struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Address     string             `bson:"address" json:"address"`
	Phone       string             `bson:"phone" json:"phone"`
	Active      bool               `bson:"active" json:"active"`
	Role        primitive.ObjectID `bson:"role,omitempty" json:"role,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
	IsRoot      bool               `bson:"isRoot" json:"isRoot"`
	Permissions []string           `bson:"permissions" json:"permissions"`
}

type CreateStaffCommand struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UpdateStaffCommand struct {
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type LoginStaffCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordBody struct {
	Password         string `json:"password"`
	NewPassword      string `json:"newPassword"`
	NewPasswordAgain string `json:"newPasswordAgain"`
}

type LoginStaffResult struct {
}

type StaffResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	Username    string             `json:"username"`
	Address     string             `json:"address"`
	Phone       string             `json:"phone"`
	Permissions []string           `json:"permissions"`
	Token       string             `json:"token"`
}

type UpdateStaffRoleCommand struct {
	ID   primitive.ObjectID
	Role string `json:"role"`
}

func (s UpdateStaffRoleCommand) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required),
		validation.Field(&s.Role, validation.Required))
}
func (stf UpdateStaffCommand) Validate() error {
	return validation.ValidateStruct(&stf,
		validation.Field(&stf.Phone, validation.Required),
		validation.Field(&stf.Address, validation.Required),
	)
}

func (stf CreateStaffCommand) Validate() error {
	return validation.ValidateStruct(&stf,
		validation.Field(&stf.Username, validation.Required),
		validation.Field(&stf.Phone),
		validation.Field(&stf.Address),
		validation.Field(&stf.Password, validation.Required),
		validation.Field(&stf.Role,
			is.MongoID.Error("MongoID invalid"), validation.Required),
	)
}

func (alg LoginStaffCommand) Validate() error {
	return validation.ValidateStruct(&alg,
		validation.Field(&alg.Username, validation.Required),
		validation.Field(&alg.Password, validation.Required),
	)
}

func (a PasswordBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Password, validation.Required),
		validation.Field(&a.NewPassword, validation.Required),
		validation.Field(&a.NewPasswordAgain, validation.Required),
	)
}
