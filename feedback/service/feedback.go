package service

import (
	"context"
	"errors"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FeedbackAdminService ...
type FeedbackAdminService struct {
	FeedbackDAO model.FeedbackDAO
}

// NewFeedbackAdminService ...
func NewFeedbackAdminService(d *model.CommonDAO) model.FeedbackAdminService {
	return &FeedbackAdminService{
		FeedbackDAO: d.Feedback,
	}
}

// Create ...
func (d *FeedbackAdminService) Create(ctx context.Context, userID primitive.ObjectID, body model.FeedbackBody) (doc model.FeedbackResponse, err error) {
	payload := body.NewFeedbackBSON(userID)
	err = d.FeedbackDAO.InsertOne(ctx, payload)
	if err != nil {
		return doc, errors.New(locale.FeedbackKeyCanNotCreate)
	}

	return payload.GetResponse(), nil
}

// ListAll ...
func (d *FeedbackAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.FeedbackResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  bson.M
		res   = make([]model.FeedbackResponse, 0)
		total int64
	)

	q.AssignKeyword(&cond)
	q.AssignActive(&cond)
	wg.Add(2)
	go func() {
		defer wg.Done()
		total = d.FeedbackDAO.CountByCondition(ctx, cond)
	}()

	go func() {
		defer wg.Done()
		feedbacks, _ := d.FeedbackDAO.FindByCondition(ctx, cond)
		for _, value := range feedbacks {
			temp := value.GetResponse()
			res = append(res, temp)
		}
	}()

	wg.Wait()
	return res, total
}

// Update ....
func (d *FeedbackAdminService) Update(ctx context.Context, userID primitive.ObjectID, feedback model.FeedbackRaw, body model.FeedbackBody) (res model.FeedbackResponse, err error) {
	doc := body.NewFeedbackBSON(userID)

	// assign
	feedback.Name = doc.Name
	feedback.Rating = doc.Rating
	feedback.Order = doc.Order
	feedback.UpdatedAt = doc.UpdatedAt

	err = d.FeedbackDAO.UpdateByID(ctx, feedback.ID, bson.M{"$set": feedback})
	if err != nil {
		return res, err
	}

	docBson, _ := d.FeedbackDAO.FindOneByCondition(ctx, bson.M{"_id": feedback.ID})
	return docBson.GetResponse(), nil
}

// FindByID ...
func (d *FeedbackAdminService) FindByID(ctx context.Context, id model.AppID) (model.FeedbackRaw, error) {
	return d.FeedbackDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

