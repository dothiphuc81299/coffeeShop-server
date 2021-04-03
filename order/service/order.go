package service

import (
	"context"
	"errors"
	"log"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderAdminService struct {
	OrderDAO model.OrderDAO
	DrinkDAO model.DrinkDAO
	UserDAO  model.UserDAO
}

func NewOrderService(d *model.CommonDAO) model.OrderAdminService {
	return &OrderAdminService{
		OrderDAO: d.Order,
		DrinkDAO: d.Drink,
		UserDAO:  d.User,
	}
}

func (o *OrderAdminService) Create(ctx context.Context, userID primitive.ObjectID, order model.OrderBody) (doc model.OrderResponse, err error) {
	// convert order payload
	drinksID := make([]primitive.ObjectID, 0)
	for _, value := range order.Drink {
		drinkID, _ := primitive.ObjectIDFromHex(value.Name)
		drinksID = append(drinksID, drinkID)
	}
	log.Println("userID", userID)
	log.Println("log1")
	// totalPrice
	cond := bson.M{
		"_id": bson.M{
			"$in": drinksID,
		},
	}
	log.Println("log2")
	drinksRaw, err := o.DrinkDAO.FindByCondition(ctx, cond)
	log.Println(err)
	if err != nil {
		return doc, errors.New(locale.OrderKeyDrinkCanNotFind)
	}

	drinksInfo := make([]model.DrinkInfo, 0)
	for _, value := range drinksRaw {
		for _, a := range order.Drink {
			doc.TotalPrice += value.Price * float64(a.Quantity)
			drinkInfo := a.NewDrinkRawInfo(value)
			drinksInfo = append(drinksInfo, drinkInfo)
		}
	}

	orderPayload := order.NewOrderRaw(userID, drinksID, doc.TotalPrice)
	err = o.OrderDAO.InsertOne(ctx, orderPayload)
	if err != nil {
		return doc, errors.New(locale.OrderKeyCanNotCreateOrder)
	}

	// response
	// 1. Get info
	userInfoRaw, err := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": userID})
	log.Println("err", userID)
	if err != nil {
		return doc, errors.New(locale.OrderKeyCanNotFindUserByUserID)
	}
	userInfo := userInfoRaw.GetUserInfo()

	res := orderPayload.GetResponse(userInfo, drinksInfo)
	return res, nil

}
