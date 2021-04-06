package service

import (
	"context"
	"errors"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// ShiftAdminService ...
type ShiftAdminService struct {
	ShiftDAO model.ShiftDAO
}

// NewShiftAdminService ...
func NewShiftAdminService(d *model.CommonDAO) model.ShiftAdminService {
	return &ShiftAdminService{
		ShiftDAO: d.Shift,
	}
}

// Create ...
func (d *ShiftAdminService) Create(ctx context.Context, body model.ShiftBody) (doc model.ShiftAdminResponse, err error) {

	payload := body.NewShiftRaw()
	err = d.ShiftDAO.InsertOne(ctx, payload)
	res := payload.ShiftGetAdminResponse()
	return res, err
}

// ListAll ...
func (d *ShiftAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.ShiftAdminResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		total int64
		res   = make([]model.ShiftAdminResponse, 0)
	)

	q.AssignKeyword(&cond)
	wg.Add(2)
	go func() {
		defer wg.Done()
		total = d.ShiftDAO.CountByCondition(ctx, cond)
	}()

	go func() {
		defer wg.Done()
		categories, _ := d.ShiftDAO.FindByCondition(ctx, cond)
		for _, value := range categories {
			temp := value.ShiftGetAdminResponse()
			res = append(res, temp)
		}
	}()

	wg.Wait()
	return res, total
}

// Update ....
func (d *ShiftAdminService) Update(ctx context.Context, Shift model.ShiftRaw, body model.ShiftBody) (doc model.ShiftAdminResponse, err error) {
	payload := body.NewShiftRaw()

	// assgin
	Shift.Name = payload.Name
	Shift.UpdatedAt = payload.UpdatedAt

	err = d.ShiftDAO.UpdateByID(ctx, Shift.ID, bson.M{"$set": Shift})
	if err != nil {
		return doc, errors.New(locale.ShiftKeyCanNotUpdate)
	}

	cat, _ := d.ShiftDAO.FindOneByCondition(ctx, bson.M{"_id": Shift.ID})
	res := cat.ShiftGetAdminResponse()
	return res, err
}

// FindByID ...
func (d *ShiftAdminService) FindByID(ctx context.Context, id model.AppID) (Shift model.ShiftRaw, err error) {
	return d.ShiftDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}
