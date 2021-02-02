package service

import (
	"context"
	"errors"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// CategoryAdminService ...
type CategoryAdminService struct {
	CategoryDAO model.CategoryDAO
}

// NewCategoryAdminService ...
func NewCategoryAdminService(d *model.CommonDAO) model.CategoryAdminService {
	return &CategoryAdminService{
		CategoryDAO: d.Category,
	}
}

// Create ...
func (d *CategoryAdminService) Create(ctx context.Context, body model.CategoryBody) (model.AppID, error) {
	if d.checkNameExisted(ctx, body.Name) {
		return model.AppID{}, errors.New(locale.CategoryKeyNameIsRequired)
	}
	doc := body.NewCategoryRaw()
	err := d.CategoryDAO.InsertOne(ctx, doc)
	return doc.ID, err
}

func (d *CategoryAdminService) checkNameExisted(ctx context.Context, name string) bool {
	total := d.CategoryDAO.CountByCondition(ctx, bson.M{"name": name})
	return total > 0
}

// ListAll ...
func (d *CategoryAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.CategoryAdminResponse, int64) {
	panic("implement it")
}

// Update ....
func (d *CategoryAdminService) Update(ctx context.Context, category model.CategoryRaw, body model.CategoryBody) error {
	doc := body.NewCategoryRaw()

	// assgin
	category.Name = doc.Name
	category.SearchString = doc.SearchString
	category.UpdatedAt = doc.UpdatedAt

	err := d.CategoryDAO.UpdateByID(ctx, category.ID, category)
	return err
}

// FindByID ...
func (d *CategoryAdminService) FindByID(ctx context.Context, id model.AppID) (Category model.CategoryRaw, err error) {
	panic("implement it ")
}
