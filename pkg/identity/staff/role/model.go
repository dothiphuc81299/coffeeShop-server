package role

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrCanNotUpdate = errors.New("Can not update")
)

type StaffRoleRaw struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
	Permissions []string           `bson:"permissions" json:"permissions"`
}

type CreateStaffRoleCommand struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type UpdateStaffRoleCommand struct {
	ID          primitive.ObjectID
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (cmd CreateStaffRoleCommand) Validate() error {
	return validation.ValidateStruct(&cmd,
		validation.Field(&cmd.Name, validation.Required),
		validation.Field(&cmd.Permissions, validation.Required),
	)
}

func (cmd UpdateStaffRoleCommand) Validate() error {
	return validation.ValidateStruct(&cmd,
		validation.Field(&cmd.ID, validation.Required),
		validation.Field(&cmd.Name, validation.Required),
		validation.Field(&cmd.Permissions, validation.Required),
	)
}