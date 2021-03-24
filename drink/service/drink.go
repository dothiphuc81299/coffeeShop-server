package service

import (
	"context"
	"errors"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// DrinkAdminService ...
type DrinkAdminService struct {
	DrinkDAO    model.DrinkDAO
	CategoryDAO model.CategoryDAO
}

// NewDrinkAdminService ...
func NewDrinkAdminService(d *model.CommonDAO) model.DrinkAdminService {
	return &DrinkAdminService{
		DrinkDAO:    d.Drink,
		CategoryDAO: d.Category,
	}
}

// Create ...
func (d *DrinkAdminService) Create(ctx context.Context, body model.DrinkBody) (doc model.DrinkAdminResponse, err error) {
	if d.checkNameExisted(ctx, body.Name) {
		return doc, errors.New(locale.DrinkKeyNameExisted)
	}
	payload := body.NewDrinkRaw()
	err = d.DrinkDAO.InsertOne(ctx, payload)
	if err != nil {
		return doc, errors.New(locale.DrinkKeyCanNotCreate)
	}
	cat, _ := d.CategoryDAO.FindOneByCondition(ctx, bson.M{"_id": payload.Category})
	catTemp := model.CategoryGetInfo(cat)
	temp := payload.DrinkGetAdminResponse(catTemp)
	return temp, err
}

func (d *DrinkAdminService) checkNameExisted(ctx context.Context, name string) bool {
	total := d.DrinkDAO.CountByCondition(ctx, bson.M{"name": name})
	return total > 0
}

// ListAll ...
func (d *DrinkAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.DrinkAdminResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		res   = make([]model.DrinkAdminResponse, 0)
		total int64
	)

	q.AssignKeyword(&cond)
	q.AssignActive(&cond)
	wg.Add(2)
	go func() {
		defer wg.Done()
		total = d.DrinkDAO.CountByCondition(ctx, cond)
	}()

	go func() {
		defer wg.Done()
		drinks, _ := d.DrinkDAO.FindByCondition(ctx, cond)
		for _, value := range drinks {
			cat, _ := d.CategoryDAO.FindOneByCondition(ctx, bson.M{"_id": value.Category})
			catTemp := model.CategoryGetInfo(cat)
			temp := value.DrinkGetAdminResponse(catTemp)
			res = append(res, temp)
		}
	}()

	wg.Wait()
	return res, total
}

// Update ....
func (d *DrinkAdminService) Update(ctx context.Context, drink model.DrinkRaw, body model.DrinkBody) (res model.DrinkAdminResponse, err error) {
	doc := body.NewDrinkRaw()

	// assign
	drink.Name = doc.Name
	drink.SearchString = doc.SearchString
	drink.Category = doc.Category
	drink.Price = doc.Price

	err = d.DrinkDAO.UpdateByID(ctx, drink.ID, bson.M{"$set": drink})
	if err != nil {
		return res, err
	}

	drinkRaw, _ := d.DrinkDAO.FindOneByCondition(ctx, bson.M{"_id": drink.ID})
	cat, _ := d.CategoryDAO.FindOneByCondition(ctx, bson.M{"_id": drinkRaw.Category})
	catTemp := model.CategoryGetInfo(cat)
	temp := drinkRaw.DrinkGetAdminResponse(catTemp)
	return temp, nil
}

// FindByID ...
func (d *DrinkAdminService) FindByID(ctx context.Context, id model.AppID) (model.DrinkRaw, error) {
	return d.DrinkDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (d *DrinkAdminService) ChangeStatus(ctx context.Context, drink model.DrinkRaw) (status bool, err error) {
	active := !drink.Active
	payload := bson.M{
		"$set": bson.M{
			"active": active,
		},
	}
	err = d.DrinkDAO.UpdateByID(ctx, drink.ID, payload)
	if err != nil {
		return
	}
	return active, nil

}

func (d *DrinkAdminService) GetDetail(ctx context.Context, drink model.DrinkRaw) model.DrinkAdminResponse {
	cat, _ := d.CategoryDAO.FindOneByCondition(ctx, bson.M{"_id": drink.Category})
	catTemp := model.CategoryGetInfo(cat)
	temp := drink.DrinkGetAdminResponse(catTemp)
	return temp
}
