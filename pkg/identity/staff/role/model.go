package role

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Role string `json:"role"`
}

func (s UpdateStaffRoleCommand) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Role, validation.Required))
}

func (sr *CreateStaffRoleCommand) Validate() error {
	return validation.ValidateStruct(sr,
		validation.Field(&sr.Name, validation.Required),
		validation.Field(&sr.Permissions, validation.Required),
	)
}
