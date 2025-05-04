package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EventBody ...
type EventBody struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// EventAdminResponse ...
type EventAdminResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Desc      string             `json:"desc"`
	CreatedAt time.Time          `json:"createdAt"`
	Active    bool               `json:"active"`
}

// Validate ...
func (c EventBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error(locale.EventKeyNameIsRequired)),
		validation.Field(&c.Desc, validation.Required.Error(locale.EventKeyDescIsRequired)),
	)
}

// NewEventRaw ...
func (c EventBody) NewEventRaw() EventRaw {
	now := time.Now()
	return EventRaw{
		ID:        Newprimitive.ObjectID(),
		Name:      c.Name,
		Desc:      c.Desc,
		CreatedAt: now,
		UpdatedAt: now,
		Active:    false,
	}
}

func (c EventRaw) EventGetAdminResponse() EventAdminResponse {
	return EventAdminResponse{
		ID:        c.ID,
		Name:      c.Name,
		Desc:      c.Desc,
		CreatedAt: c.CreatedAt,
		Active:    c.Active,
	}
}
