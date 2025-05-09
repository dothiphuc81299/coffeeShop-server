package order

import (
	"errors"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type OrderRaw struct {
// 	ID         primitive.ObjectID `bson:"_id" json:"id"`
// 	User       primitive.ObjectID `bson:"user" json:"user"`
// 	Drink      []DrinkInfo        `bson:"drink" json:"drink"`
// 	Status     string             `bson:"status" json:"status"`
// 	TotalPrice float64            `bson:"totalPrice" json:"totalPrice"`
// 	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
// 	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
// 	UpdatedBy  primitive.ObjectID `bson:"updatedBy,omitempty" json:"updatedBy,omitempty"`
// 	IsPoint    bool               `bson:"is_point" json:"isPoint"`
// 	Point      float64            `bson:"point" json:"point"`
// }

var (
	ErrCannotCreateOrder       = errors.New("cannot create order")
	ErrUpdatePointFailed       = errors.New("update point failed")
	ErrOrderNotFound           = errors.New("order not found")
	ErrOrderCanNotCancel       = errors.New("order can not cancel")
	ErrUserNotFound            = errors.New("user not found")
	ErrOrderStatusCanNotUpdate = errors.New("order status can not update")
	ErrCommonKeyInvalidID      = errors.New("common key invalid id")
)

type OrderRaw struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID            primitive.ObjectID `bson:"user_id" json:"user_id"`                         // Reference tới user
	ShippingAddressID primitive.ObjectID `bson:"shipping_address_id" json:"shipping_address_id"` // Tham chiếu tới địa chỉ giao hàng đã chọn
	Items             []*OrderItemRaw    `bson:"items" json:"items"`
	Status            string             `bson:"status" json:"status"` // Consider: dùng enum/const
	Total             float64            `bson:"total" json:"total"`
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
	UsePoint          bool               `bson:"use_point" json:"use_point"`
	Point             float64            `bson:"point" json:"point"`
}

type OrderItemRaw struct {
	DrinkID  primitive.ObjectID `bson:"drink_id" json:"drink_id"`
	Quantity int                `bson:"quantity" json:"quantity"`
	Price    float64            `bson:"price" json:"price"`
	Total    float64            `bson:"total" json:"total"`
	Name     string             `bson:"name" json:"name"`
}

type OrderBody struct {
	UserID            primitive.ObjectID
	Items             []*OrderItemRaw                               `json:"items"`
	UsePoint          bool                                          `json:"use_point"`
	Point             float64                                       `json:"point"`
	ShippingAddressID string                                        `json:"shipping_address_id"`
	Shipping          *shippingaddress.CreateShippingAddressCommand `json:"shipping_address"`
}

type DrinkUserBody struct {
	Name     string `json:"name"` // ObjectID
	Quantity int    `json:"quantity"`
}

type OrderResponse struct {
	ID         primitive.ObjectID                      `json:"_id"`
	User       UserInfo                                `json:"user"`
	Drink      []*OrderItemRaw                         `json:"drink"`
	Shipping   *shippingaddress.UserShippingAddressRaw `json:"shipping_address"`
	Status     string                                  `json:"status"` //pending,success
	TotalPrice float64                                 `json:"totalPrice"`
	CreatedAt  time.Time                               `json:"createdAt"`
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

type StatisticResponse struct {
	Statistic     []StatisticByDrink `json:"statistic"`
	TotalQuantity float64            `json:"totalQuanity"`
	TotalSale     float64            `json:"totalSale"`
}

type StatisticByDrink struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Name          string             `bson:"name" json:"name"`
	TotalSale     float64            `bson:"totalSale" json:"totalSale"`
	TotalQuantity float64            `bson:"totalQuantity" json:"totalQuantity"`
}

type SearchOrdersQuery struct {
	OrderID string `schema:"order_id"`
	UserID  string `schema:"user_id"`
	Page    int64  `schema:"page"`
	Limit   int64  `schema:"limit"`
	Status  string `schema:"status"`
}

type UpdateOrderStatusCommand struct {
	ID     primitive.ObjectID
	Status string `json:"status"`
}

func (o *OrderRaw) GetResponse(userInfo UserInfo, items []*OrderItemRaw, status string, shippingAddr *shippingaddress.UserShippingAddressRaw) OrderResponse {
	return OrderResponse{
		ID:         o.ID,
		User:       userInfo,
		Drink:      items,
		Status:     status,
		Shipping:   shippingAddr,
		TotalPrice: o.Total,
		CreatedAt:  o.CreatedAt,
	}
}

// func (o OrderBody) Validate() error {
// 	err := validation.Validate(o.Drink, validation.Required)
// 	if err != nil {
// 		return err
// 	}

// 	return validation.ValidateStruct(&o,
// 		validation.Field(&o.Point, validation.Required.When(o.IsPoint)),
// 	)
// }

// func (d DrinkUserBody) Validate() error {
// 	return validation.ValidateStruct(&d,
// 		validation.Field(&d.Name, validation.Required,
// 			is.MongoID),
// 		validation.Field(&d.Quantity, validation.Required),
// 	)
// }

// func (s StatusBody) Validate() error {
// 	return validation.ValidateStruct(&s,
// 		validation.Field(&s.Status,
// 			validation.In("cancel", "success"),
// 			validation.Required))
// }

// func (o *OrderRaw) GetResponse(u UserInfo, d []OrderItemRaw, status string) OrderResponse {
// 	return OrderResponse{
// 		ID:         o.ID,
// 		User:       u,
// 		Drink:      d,
// 		Status:     status,
// 		TotalPrice: o.Total,
// 		CreatedAt:  o.CreatedAt,
// 	}
// }

func (o OrderBody) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.UserID, validation.Required),
		validation.Field(&o.Items, validation.Required, validation.Each(validation.NotNil)),
		validation.Field(&o.Point, validation.Min(0.0)),
		validation.Field(&o.ShippingAddressID, validation.When(o.Shipping == nil,
			validation.Required, is.Hexadecimal)),
		validation.Field(&o.Shipping, validation.When(o.ShippingAddressID == "",
			validation.Required, validation.By(func(value interface{}) error {
				if v, ok := value.(*shippingaddress.CreateShippingAddressCommand); ok && v != nil {
					return v.Validate()
				}
				return nil
			}))),
	)
}

// func (o OrderBody) NewOrderRaw(userID primitive.ObjectID, drink []DrinkInfo, totalPrice float64) OrderRaw {
// 	return OrderRaw{
// 		ID:         primitive.NewObjectID(),
// 		User:       userID,
// 		Status:     "pending",
// 		Drink:      drink,
// 		TotalPrice: totalPrice,
// 		CreatedAt:  time.Now(),
// 		UpdatedAt:  time.Now(),
// 		IsPoint:    o.IsPoint,
// 		Point:      o.Point,
// 	}
// }

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
