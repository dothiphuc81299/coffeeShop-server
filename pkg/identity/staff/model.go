package staff

import (
	"errors"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrStaffExisted = errors.New("Staff existed")
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

type SearchStaffResult struct {
	Staffs  []*Staff `json:"staffs"`
	Total   int64    `json:"total"`
	Page    int64    `json:"page"`
	PerPage int64    `json:"perPage"`
}

type LoginStaffResult struct {
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
			is.MongoID.Error(locale.CommonKeyIDMongoInvalid), validation.Required),
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
