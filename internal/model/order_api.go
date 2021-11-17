package model

import (
	"log"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderBody struct {
	Drink   []DrinkUserBody `json:"drink"`
	IsPoint bool            `json:"is_point"`
	Point   float64         `json:"point"`
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
	Status     string             `json:"status"` //pending,success
	TotalPrice float64            `json:"totalPrice"`
	CreatedAt  time.Time          `json:"createdAt"`
}

type DrinkInfo struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Price    float64            `json:"price" bson:"price"`
	Quantity int                `json:"quantity" bson:"quantity"`
}

type UserInfo struct {
	ID       primitive.ObjectID `json:"_id"`
	UserName string             `json:"username"`
	Address  string             `json:"address"`
}

type StatusBody struct {
	Status string `json:"status"`
}

func (o OrderBody) Validate() error {
	err := validation.Validate(o.Drink, validation.Required.Error("Don hang dang bi trong"))
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&o,
		validation.Field(&o.Point, validation.Required.When(o.IsPoint).Error("PointISRequired")),
	)
}

func (d DrinkUserBody) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required.Error(locale.OrderKeyNameDrinkIsRequired),
			is.MongoID.Error(locale.OrderKeyNameDrinkInvalid)),
		validation.Field(&d.Quantity, validation.Required.Error(locale.OrderKeyQuantityIsRequired)),
	)
}

func (s StatusBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Status,
			validation.In("delivery", "success").Error("trang thai khong hop le[delivery,success"),
			validation.Required.Error("trang thai don hang duoc yeu cau")))
}

func (o *OrderRaw) GetResponse(u UserInfo, d []DrinkInfo, status string) OrderResponse {
	return OrderResponse{
		ID:         o.ID,
		User:       u,
		Drink:      d,
		Status:     status,
		TotalPrice: o.TotalPrice,
		CreatedAt:  o.CreatedAt,
	}
}

// NewOrderRaw ...
func (o OrderBody) NewOrderRaw(userID primitive.ObjectID, drink []DrinkInfo, totalPrice float64) OrderRaw {
	return OrderRaw{
		ID:         primitive.NewObjectID(),
		User:       userID,
		Status:     "pending",
		Drink:      drink,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		IsPoint:    o.IsPoint,
		Point:      o.Point,
	}
}

// func (d DrinkUserBody) NewDrinkRawInfo() DrinkInfo {
// 	return DrinkInfo{
// 		ID:       a.ID,
// 		Name:     a.Name,
// 		Quantity: d.Quantity,
// 		Price:    a.Price,
// 	}
// }

func (d *DrinkUserBody) NewDrinkInfo(e DrinkInfo) DrinkInfo {
	id, _ := primitive.ObjectIDFromHex(d.Name)
	return DrinkInfo{
		ID:       id,
		Quantity: d.Quantity,
		Name:     e.Name,
		Price:    e.Price,
	}

}

func (d *DrinkRaw) GetResponseDrink() DrinkInfo {
	log.Println("drinkPrice", d.Price)
	return DrinkInfo{
		Name:  d.Name,
		Price: d.Price,
	}
}

func (u UserRaw) GetUserInfo() UserInfo {
	return UserInfo{
		ID:       u.ID,
		UserName: u.Username,
		Address:  u.Address,
	}
}
