package shippingaddress

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserShippingAddressRaw struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	FullName  string             `bson:"full_name" json:"full_name"`
	Phone     string             `bson:"phone" json:"phone"`
	Address   string             `bson:"address" json:"address"`
	Province  string             `bson:"province" json:"province"`
	City      string             `bson:"city" json:"city"`
	Ward      string             `bson:"ward" json:"ward"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateShippingAddressCommand struct {
	UserID   primitive.ObjectID `json:"user_id"`
	FullName string             `json:"full_name"`
	Phone    string             `json:"phone"`
	Address  string             `json:"address"`
	Province string             `json:"province"`
	City     string             `json:"city"`
	Ward     string             `json:"ward"`
}

func (c CreateShippingAddressCommand) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FullName, validation.Required),
		validation.Field(&c.Phone, validation.Required, is.Digit, validation.Length(9, 15)),
		validation.Field(&c.Address, validation.Required),
	)
}
