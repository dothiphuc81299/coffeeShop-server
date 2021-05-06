package service

import (
	"context"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type DrinkAnalyticService struct {
	DrinkAnalytic model.DrinkAnalyticDAO
	Order         model.OrderDAO
	Drink         model.DrinkDAO
	Category      model.CategoryDAO
}

func NewDrinkAnalyticService(d *model.CommonDAO) model.DrinkAnalyticService {
	return &DrinkAnalyticService{
		DrinkAnalytic: d.DrinkAnalytic,
		Order:         d.Order,
		Drink:         d.Drink,
		Category:      d.Category,
	}
}

func (d *DrinkAnalyticService) ListAll(ctx context.Context, q model.CommonQuery) []model.DrinkAnalyticResponse {
	var (
		cond = bson.M{}
		wg   sync.WaitGroup
		res  = make([]model.DrinkAnalyticResponse, 0)
	)

//	q.AssignStartAtAndEndAtForDrink(&cond)
	temps, _ := d.DrinkAnalytic.AggregateDrink(ctx, cond)
	if len(temps) > 0 {
		wg.Add(len(temps))
		for index, drink := range temps {
			res = make([]model.DrinkAnalyticResponse, len(temps))
			go func(dr model.DrinkAnalyticTempResponse, i int) {
				defer wg.Done()

				temp := model.DrinkAnalyticResponse{}
				drinkRaw, _ := d.Drink.FindOneByCondition(ctx, dr.ID)
				temp.Drink = drinkRaw.GetResponse()
				categoryRaw, _ := d.Category.FindOneByCondition(ctx, drinkRaw.Category)
				temp.Category = categoryRaw.GetResponse()
				temp.Total = dr.Total
				res[i] = temp
			}(drink, index)
		}
		wg.Wait()
	}
	return res
}
