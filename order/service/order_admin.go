package service

import (
	"context"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

const delivery = "delivery"
const success = "success"

type OrderAdminService struct {
	OrderDAO model.OrderDAO
	DrinkDAO model.DrinkDAO
	UserDAO  model.UserDAO
}

func NewOrderAdminService(d *model.CommonDAO) model.OrderAdminService {
	return &OrderAdminService{
		OrderDAO: d.Order,
		DrinkDAO: d.Drink,
		UserDAO:  d.User,
	}
}

func (o *OrderAdminService) GetListByStatus(ctx context.Context, query model.CommonQuery) ([]model.OrderResponse, int64) {
	var (
		cond  = bson.M{}
		total int64
		wg    sync.WaitGroup
		res   = make([]model.OrderResponse, 0)
	)

	// assign
	query.AssignStatus(&cond)
	total = o.OrderDAO.CountByCondition(ctx, cond)
	orders, _ := o.OrderDAO.FindByCondition(ctx, cond, query.GetFindOptsUsingPage())

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

func (o *OrderAdminService) ChangeStatus(ctx context.Context, order model.OrderRaw, status model.StatusBody, staff model.StaffRaw) (res string, err error) {
	payload := bson.M{
		"updatedAt": time.Now(),
		"status":    status.Status,
		"shipper":   staff.ID,
	}

	err = o.OrderDAO.UpdateByID(ctx, order.ID, bson.M{"$set": payload})
	if err != nil {
		return res, err
	}

	return status.Status, err
}

func (o *OrderAdminService) FindByID(ctx context.Context, id model.AppID) (model.OrderRaw, error) {
	return o.OrderDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (o *OrderAdminService) GetDetail(ctx context.Context, order model.OrderRaw) (doc model.OrderResponse) {
	user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})
	if user.ID.IsZero() {
		return
	}
	userInfo := user.GetUserInfo()

	res := order.GetResponse(userInfo, order.Drink, order.Status)
	return res
}
