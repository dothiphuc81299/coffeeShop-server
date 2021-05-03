package service

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type DrinkAnalyticService struct {
	DrinkAnalytic model.DrinkAnalyticDAO
	Order         model.OrderDAO
}

func NewDrinkAnalyticService(d *model.CommonDAO) model.DrinkAnalyticService {
	return &DrinkAnalyticService{
		DrinkAnalytic: d.DrinkAnalytic,
		Order:         d.Order,
	}
}

func (d *DrinkAnalyticService) ListAll(ctx context.Context, q model.CommonQuery) []model.DrinkAnalyticResponse {
	var (
		cond = bson.M{}

		res = make([]model.DrinkAnalyticResponse, 0)
	)

	q.AssignStartAtAndEndAtForDrink(&cond)
	res, _ = d.DrinkAnalytic.AggregateDrink(ctx, cond)
	return res
}
