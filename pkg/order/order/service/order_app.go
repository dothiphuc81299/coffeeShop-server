 package service

// import (
// 	"context"
// 	"errors"
// 	"sync"
// 	"time"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type OrderAppService struct {
// 	OrderDAO         model.OrderDAO
// 	DrinkDAO         model.DrinkDAO
// 	UserDAO          model.UserDAO
// 	DrinkAnalyticDAO model.DrinkAnalyticDAO
// }

// const POINT = 3000

// func NewOrderAppService(d *model.CommonDAO) model.OrderAppService {
// 	return &OrderAppService{
// 		OrderDAO:         d.Order,
// 		DrinkDAO:         d.Drink,
// 		UserDAO:          d.User,
// 		DrinkAnalyticDAO: d.DrinkAnalytic,
// 	}
// }

// func (o *OrderAppService) Create(ctx context.Context, user model.UserRaw, order model.OrderBody) (doc model.OrderResponse, err error) {
// 	// convert order payload
// 	drinks := make([]model.DrinkInfo, 0)
// 	for _, value := range order.Drink {
// 		drinkID, _ := primitive.ObjectIDFromHex(value.Name)
// 		drinkRaw, err := o.DrinkDAO.FindOneByCondition(ctx, bson.M{"_id": drinkID})
// 		if err != nil {
// 			return doc, err
// 		}
// 		drink := model.DrinkInfo{
// 			ID:       drinkID,
// 			Name:     drinkRaw.Name,
// 			Price:    drinkRaw.Price,
// 			Quantity: value.Quantity,
// 		}
// 		doc.TotalPrice += drink.Price * float64(drink.Quantity)
// 		drinks = append(drinks, drink)
// 	}
// 	var currentPointUpdate float64
// 	// if !order.IsPoint {
// 	// 	// calculate currentPoint
// 	// 	if (doc.TotalPrice >= 30000) && (doc.TotalPrice) <= 50000 {
// 	// 		currentPointUpdate = user.CurrentPoint + 1
// 	// 	} else if (doc.TotalPrice > 50000) && (doc.TotalPrice) <= 100000 {
// 	// 		currentPointUpdate = user.CurrentPoint + 2
// 	// 	} else if doc.TotalPrice > 100000 {
// 	// 		currentPointUpdate = user.CurrentPoint + 3
// 	// 	}
// 	// }
// 	if order.Point > 0 {
// 		if user.CurrentPoint < order.Point {
// 			return doc, errors.New(locale.CurrentPointIsNotEnough)
// 		}
// 		missPoint := order.Point * POINT

// 		if missPoint > doc.TotalPrice {
// 			point := missPoint - doc.TotalPrice
// 			point = point / POINT
// 			///	currentPointUpdate = point
// 			currentPointUpdate = user.CurrentPoint - order.Point + point
// 			order.Point = order.Point - point
// 			doc.TotalPrice = 0
// 		} else {
// 			doc.TotalPrice -= missPoint
// 			currentPointUpdate = user.CurrentPoint - order.Point
// 		}
// 	}

// 	orderPayload := order.NewOrderRaw(user.ID, drinks, doc.TotalPrice)

// 	err = o.OrderDAO.InsertOne(ctx, orderPayload)
// 	if err != nil {
// 		return doc, errors.New(locale.OrderKeyCanNotCreateOrder)
// 	}

// 	if err = o.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"currentPoint": currentPointUpdate}}); err != nil {
// 		return doc, errors.New(locale.UpdatePointFailed)
// 	}

// 	userInfo := user.GetUserInfo()
// 	res := orderPayload.GetResponse(userInfo, drinks, orderPayload.Status)
// 	return res, nil

// }

// func (o *OrderAppService) FindByID(ctx context.Context, id model.primitive.ObjectID) (model.OrderRaw, error) {
// 	return o.OrderDAO.FindOneByCondition(ctx, bson.M{"_id": id})
// }

// func (o *OrderAppService) GetDetail(ctx context.Context, order model.OrderRaw) (doc model.OrderResponse) {
// 	user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})

// 	if user.ID.IsZero() {
// 		return
// 	}

// 	userInfo := user.GetUserInfo()
// 	res := order.GetResponse(userInfo, order.Drink, order.Status)
// 	return res
// }

// func (o *OrderAppService) Search(ctx context.Context, query model.CommonQuery, user model.UserRaw) ([]model.OrderResponse, int64) {
// 	var (
// 		cond = bson.M{
// 			"user": user.ID,
// 		}
// 		total int64
// 		wg    sync.WaitGroup
// 		res   = make([]model.OrderResponse, 0)
// 	)
// 	// assign
// 	query.AssignStatus(&cond)
// 	total = o.OrderDAO.CountByCondition(ctx, cond)
// 	orders, _ := o.OrderDAO.FindByCondition(ctx, cond, query.GetFindOptsUsingPage())

// 	if len(orders) > 0 {
// 		wg.Add(len(orders))
// 		res = make([]model.OrderResponse, len(orders))
// 		for index, order := range orders {
// 			go func(od model.OrderRaw, i int) {
// 				defer wg.Done()
// 				user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": od.User})

// 				userInfo := user.GetUserInfo()

// 				temp := od.GetResponse(userInfo, od.Drink, od.Status)
// 				res[i] = temp

// 			}(order, index)
// 		}
// 		wg.Wait()

// 	}
// 	return res, total
// }

// func (o *OrderAppService) RejectOrder(ctx context.Context, user model.UserRaw, order model.OrderRaw) error {
// 	now := time.Now()
// 	createdAt := order.CreatedAt
// 	then := createdAt.Add(time.Duration(+2) * time.Minute)

// 	if now.After(then) || order.Status != "pending" {
// 		return errors.New(locale.OrderCanNotCancel)
// 	}

// 	payload := bson.M{
// 		"updatedAt": time.Now(),
// 		"status":    "cancel",
// 		"updatedBy": user.ID,
// 	}

// 	err := o.OrderDAO.UpdateByID(ctx, order.ID, bson.M{"$set": payload})
// 	if err != nil {
// 		return errors.New(locale.CommonKeyErrorWhenHandle)
// 	}

// 	// check lai diem
// 	currentPointUpdate := user.CurrentPoint + order.Point
// 	if order.IsPoint && order.Point > 0 {
// 		if err = o.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"currentPoint": currentPointUpdate}}); err != nil {
// 			return errors.New(locale.UpdatePointFailed)
// 		}
// 	}

// 	return nil

// }
