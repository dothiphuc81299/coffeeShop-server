package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderAppService struct {
	OrderDAO         model.OrderDAO
	DrinkDAO         model.DrinkDAO
	UserDAO          model.UserDAO
	DrinkAnalyticDAO model.DrinkAnalyticDAO
}

func NewOrderAppService(d *model.CommonDAO) model.OrderAppService {
	return &OrderAppService{
		OrderDAO:         d.Order,
		DrinkDAO:         d.Drink,
		UserDAO:          d.User,
		DrinkAnalyticDAO: d.DrinkAnalytic,
	}
}

func (o *OrderAppService) Create(ctx context.Context, user model.UserRaw, order model.OrderBody) (doc model.OrderResponse, err error) {
	// convert order payload
	drinks := make([]model.DrinkInfo, 0)
	drinkAnalytics := make([]model.DrinkAnalyticRaw, 0)
	for _, value := range order.Drink {
		drinkAnalytic := model.DrinkAnalyticRaw{}
		drinkID, _ := primitive.ObjectIDFromHex(value.Name)
		drinkRaw, err := o.DrinkDAO.FindOneByCondition(ctx, bson.M{"_id": drinkID})
		if err != nil {
			return doc, err
		}
		drink := model.DrinkInfo{
			ID:       drinkID,
			Name:     drinkRaw.Name,
			Price:    drinkRaw.Price,
			Quantity: value.Quantity,
		}
		doc.TotalPrice += drink.Price * float64(drink.Quantity)

		// append
		drinkAnalytic.ID = primitive.NewObjectID()
		drinkAnalytic.Category = drinkRaw.Category
		drinkAnalytic.Name = drinkRaw.ID
		drinkAnalytic.TotalDrink = float64(value.Quantity)
		drinkAnalytic.UpdateAt = time.Now()
		drinkAnalytic.CreatedAt = time.Now()

		drinks = append(drinks, drink)
		drinkAnalytics = append(drinkAnalytics, drinkAnalytic)

	}

	orderPayload := order.NewOrderRaw(user.ID, drinks, doc.TotalPrice)

	err = o.OrderDAO.InsertOne(ctx, orderPayload)
	if err != nil {
		return doc, errors.New(locale.OrderKeyCanNotCreateOrder)
	}

	// for _, drink := range drinks {
	// 	a := model.DrinkAnalyticRaw{}
	// 	a.Name = drink.ID
	// 	a.TotalDrink = float64(drink.Quantity)
	// 	a.UpdateAt = time.Now()
	// 	a.CreatedAt = time.Now()
	// 	drinkAnalytic = append(drinkAnalytic, a)
	// }
	// log.Println("doodod", drinkAnalytic)
	var docs []interface{}
	for _, item := range drinkAnalytics {
		docs = append(docs, item)
	}
	if len(docs) > 0 {
		if err := o.DrinkAnalyticDAO.InsertMany(ctx, docs); err != nil {
			fmt.Println("Insert analytic err: ", err)
		}
	}
	userInfo := user.GetUserInfo()

	res := orderPayload.GetResponse(userInfo, drinks, orderPayload.Status)
	return res, nil

}

func (o *OrderAppService) FindByID(ctx context.Context, id model.AppID) (model.OrderRaw, error) {
	return o.OrderDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (o *OrderAppService) GetDetail(ctx context.Context, order model.OrderRaw) (doc model.OrderResponse) {
	user, err := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})
	if err != nil {
		return
	}
	userInfo := user.GetUserInfo()

	res := order.GetResponse(userInfo, order.Drink, order.Status)
	return res
}

func (o *OrderAppService) GetList(ctx context.Context, user model.UserRaw) ([]model.OrderResponse, int64) {
	var (
		cond = bson.M{
			"user": user.ID,
		}
		total int64
		wg    sync.WaitGroup
		res   = make([]model.OrderResponse, 0)
	)

	total = o.OrderDAO.CountByCondition(ctx, cond)
	orders, _ := o.OrderDAO.FindByCondition(ctx, cond)

	if len(orders) > 0 {
		wg.Add(len(orders))
		res = make([]model.OrderResponse, len(orders))
		for index, order := range orders {
			go func(od model.OrderRaw, i int) {
				defer wg.Done()
				user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": od.User})

				userInfo := user.GetUserInfo()

				temp := od.GetResponse(userInfo, od.Drink, od.Status)
				res[i] = temp

			}(order, index)
		}
		wg.Wait()

	}
	return res, total
}
