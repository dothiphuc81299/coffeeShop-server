package service

import (
	"context"
	"errors"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// EventAdminService ...
type EventAdminService struct {
	EventDAO model.EventDAO
}

// NewEventAdminService ...
func NewEventAdminService(d *model.CommonDAO) model.EventAdminService {
	return &EventAdminService{
		EventDAO: d.Event,
	}
}

// Create ...
func (d *EventAdminService) Create(ctx context.Context, body model.EventBody) (doc model.EventAdminResponse, err error) {
	payload := body.NewEventRaw()
	err = d.EventDAO.InsertOne(ctx, payload)
	res := payload.EventGetAdminResponse()
	return res, err
}

func (d *EventAdminService) GetDetail(ctx context.Context, event model.EventRaw) model.EventAdminResponse {
	return event.EventGetAdminResponse()
}

// ListAll ...
func (d *EventAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.EventAdminResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		total int64
		res   = make([]model.EventAdminResponse, 0)
	)

	q.AssignActive(&cond)

	wg.Add(2)
	go func() {
		defer wg.Done()
		total = d.EventDAO.CountByCondition(ctx, cond)
	}()

	go func() {
		defer wg.Done()
		events, _ := d.EventDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
		for _, value := range events {
			temp := value.EventGetAdminResponse()
			res = append(res, temp)
		}
	}()

	wg.Wait()
	return res, total
}

// Update ....
func (d *EventAdminService) Update(ctx context.Context, Event model.EventRaw, body model.EventBody) (doc model.EventAdminResponse, err error) {
	payload := body.NewEventRaw()

	// assgin
	Event.Name = payload.Name
	Event.Desc = payload.Desc
	Event.UpdatedAt = payload.UpdatedAt

	err = d.EventDAO.UpdateByID(ctx, Event.ID, bson.M{"$set": Event})
	if err != nil {
		return doc, errors.New(locale.EventKeyCanNotUpdate)
	}

	event, _ := d.EventDAO.FindOneByCondition(ctx, bson.M{"_id": Event.ID})
	res := event.EventGetAdminResponse()
	return res, err
}

// FindByID ...
func (d *EventAdminService) FindByID(ctx context.Context, id model.AppID) (event model.EventRaw, err error) {
	return d.EventDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (d *EventAdminService) ChangeStatus(ctx context.Context, event model.EventRaw) (status bool, err error) {
	active := !event.Active

	payload := bson.M{
		"$set": bson.M{
			"active": active,
		},
	}

	err = d.EventDAO.UpdateByID(ctx, event.ID, payload)
	if err != nil {
		return status, errors.New(locale.EventKeyCanNotUpdate)
	}
	return active, nil

}
