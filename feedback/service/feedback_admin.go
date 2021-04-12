package service

import (
	"context"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (d *FeedbackAdminService) ChangeStatus(ctx context.Context, feedback model.FeedbackRaw) (status bool, err error) {
	payload := bson.M{
		"active":    !feedback.Active,
		"updatedAt": time.Now(),
	}

	err = d.FeedbackDAO.UpdateByID(ctx, feedback.ID, bson.M{"$set": payload})
	if err != nil {
		return
	}

	return !feedback.Active, nil

}

// FindByID ...
func (d *FeedbackAdminService) FindByID(ctx context.Context, id model.AppID) (model.FeedbackRaw, error) {
	return d.FeedbackDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}
