package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderBody struct {
	Drink []DrinkUserBody `json:"drink"`
}

func (o OrderBody) Validate() error {
	return validation.Validate(o.Drink)
}

func (d DrinkUserBody) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required.Error(locale.OrderKeyNameDrinkIsRequired),
			is.MongoID.Error(locale.OrderKeyNameDrinkInvalid)),
		validation.Field(&d.Quantity, validation.Required.Error(locale.OrderKeyQuantityIsRequired)),
	)
}

type DrinkUserBody struct {
	Name     string `json:"name"` // ObjectID
	Quantity int    `json:"quantity"`
}

// OrderResponse ..
type OrderResponse struct {
	ID         primitive.ObjectID `json:"_id"`
	User       UserInfo           `json:"user"`
	Drink      []DrinkInfo        `json:"drink"`
	Status     string             `json:"status"`
	TotalPrice float64            `json:"totalPrice"`
	CreatedAt  time.Time          `json:"createdAt"`
}

type DrinkInfo struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Price    float64            `json:"price"`
	Quantity int                `json:"quantity"`
}

type UserInfo struct {
	ID       primitive.ObjectID `json:"_id"`
	UserName string             `json:"username"`
	Address  string             `json:"address"`
}

func (o OrderRaw) GetResponse(u UserInfo, d []DrinkInfo) OrderResponse {
	return OrderResponse{
		ID:    o.ID,
		User:  u,
		Drink: d,
		//	Status: o.Status,
		//		TotalPrice: o.TotalPrice,
		CreatedAt: o.CreatedAt,
	}
}

// NewOrderRaw ...
func (o OrderBody) NewOrderRaw(userID primitive.ObjectID, drink []primitive.ObjectID, totalPrice float64) OrderRaw {
	return OrderRaw{
		ID:   primitive.NewObjectID(),
		User: userID,
		//	Status: false,
		Drink:      drink,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (d DrinkUserBody) NewDrinkRawInfo(a DrinkRaw) DrinkInfo {
	return DrinkInfo{
		ID:       a.ID,
		Name:     a.Name,
		Quantity: d.Quantity,
		Price:    a.Price,
	}
}

func (u UserRaw) GetUserInfo() UserInfo {
	return UserInfo{
		ID:       u.ID,
		UserName: u.Username,
		Address:  u.Address,
	}
}
