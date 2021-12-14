package service

import (
	"context"
	"errors"
	"sync"

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
func (d *CategoryAdminService) Create(ctx context.Context, body model.CategoryBody) (doc model.CategoryAdminResponse, err error) {
	if d.checkNameExisted(ctx, body.Name) {
		return doc, errors.New(locale.CategoryKeyNameExisted)
	}
	payload := body.NewCategoryRaw()
	err = d.CategoryDAO.InsertOne(ctx, payload)
	res := payload.CategoryGetAdminResponse()
	return res, err
}

func (d *CategoryAdminService) checkNameExisted(ctx context.Context, name string) bool {
	total := d.CategoryDAO.CountByCondition(ctx, bson.M{"name": name})
	return total > 0
}

// ListAll ...
func (d *CategoryAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.CategoryAdminResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		total int64
		res   = make([]model.CategoryAdminResponse, 0)
	)

	q.AssignKeyword(&cond)
	wg.Add(2)
	go func() {
		defer wg.Done()
		total = d.CategoryDAO.CountByCondition(ctx, cond)
	}()

	go func() {
		defer wg.Done()
		categories, _ := d.CategoryDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPageOne())
		for _, value := range categories {
			temp := value.CategoryGetAdminResponse()
			res = append(res, temp)
		}
	}()

	wg.Wait()
	return res, total
}

func (d *CategoryAdminService) GetDetail(ctx context.Context, cate model.CategoryRaw) model.CategoryAdminResponse {
	return cate.CategoryGetAdminResponse()
}

func (d *CategoryAdminService) DeleteCategory(ctx context.Context, cate model.CategoryRaw) error {
	err := d.CategoryDAO.DeleteByID(ctx, cate.ID)
	return err
}

// Update ....
func (d *CategoryAdminService) Update(ctx context.Context, category model.CategoryRaw, body model.CategoryBody) (doc model.CategoryAdminResponse, err error) {
	payload := body.NewCategoryRaw()

	// assgin
	category.Name = payload.Name
	category.SearchString = payload.SearchString
	category.UpdatedAt = payload.UpdatedAt

	err = d.CategoryDAO.UpdateByID(ctx, category.ID, bson.M{"$set": category})
	if err != nil {
		return doc, errors.New(locale.CategoryKeyCanNotUpdate)
	}

	cat, _ := d.CategoryDAO.FindOneByCondition(ctx, bson.M{"_id": category.ID})
	res := cat.CategoryGetAdminResponse()
	return res, err
}

// FindByID ...
func (d *CategoryAdminService) FindByID(ctx context.Context, id model.AppID) (Category model.CategoryRaw, err error) {
	return d.CategoryDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}
