package service

import (
	"context"
	"errors"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// DrinkAdminService ...
type DrinkAdminService struct {
	DrinkDAO model.DrinkDAO
}

// NewDrinkAdminService ...
func NewDrinkAdminService(d *model.CommonDAO) model.DrinkAdminService {
	return &DrinkAdminService{
		DrinkDAO: d.Drink,
	}
}

// Create ...
func (d *DrinkAdminService) Create(ctx context.Context, body model.DrinkBody) (model.AppID, error) {
	if d.checkNameExisted(ctx, body.Name) {
		return model.AppID{}, errors.New(locale.DrinkKeyNameIsRequired)
	}
	doc := body.NewDrinkRaw()
	err := d.DrinkDAO.InsertOne(ctx, doc)
	return doc.ID, err
}

func (d *DrinkAdminService) checkNameExisted(ctx context.Context, name string) bool {
	total := d.DrinkDAO.CountByCondition(ctx, bson.M{"name": name})
	return total > 0
}

// ListAll ...
func (d *DrinkAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.DrinkAdminResponse, int64) {
	panic("implement it")
}


func (d *DrinkAdminService) Update(ctx context.Context)