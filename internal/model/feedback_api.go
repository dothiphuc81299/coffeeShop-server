package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedbackBody struct {
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	Order  string `json:"order"`
	Drink  string `json:"drink"`
}

type FeedbackResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Order     primitive.ObjectID `json:"order"`
	Rating    int                `json:"rating"`
	User      UserInfo           `json:"user"`
	CreatedAt TimeResponse       `json:"createdAt"`
	Active    bool               `json:"active"`
	Drink     DrinkInfo          `json:"drink"`
}

func (f FeedbackBody) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required.Error(locale.FeedbackKeyNameIsRequired)),
		validation.Field(&f.Rating, validation.Required.Error(locale.FeedbackKeyRatingIsRequired)),
		validation.Field(&f.Order, validation.Required.Error(locale.FeedbackKeyOrderIsRequired),
			is.MongoID.Error(locale.FeedbackKeyOrderInvalid)),
		validation.Field(&f.Drink, validation.Required.Error("drink is required"),
			is.MongoID.Error("drink invalid")),
	)
}

func (f *FeedbackBody) NewFeedbackBSON(userID primitive.ObjectID) FeedbackRaw {
	orderID, _ := primitive.ObjectIDFromHex(f.Order)
	return FeedbackRaw{
		ID:        primitive.NewObjectID(),
		Name:      f.Name,
		Rating:    f.Rating,
		Order:     orderID,
		User:      userID,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (f FeedbackRaw) GetResponse(user UserInfo) FeedbackResponse {
	return FeedbackResponse{
		ID:        f.ID,
		Name:      f.Name,
		Rating:    f.Rating,
		Order:     f.Order,
		User:      user,
		Active:    f.Active,
		CreatedAt: TimeResponse{Time: f.CreatedAt},
	}
}
