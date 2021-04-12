package service

import (
	"context"
	"errors"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// FeedbackAppService ...
type FeedbackAppService struct {
	FeedbackDAO model.FeedbackDAO
	UserDAO     model.UserDAO
}

// NewFeedbackAppService ...
func NewFeedbackAppService(d *model.CommonDAO) model.FeedbackAppService {
	return &FeedbackAppService{
		FeedbackDAO: d.Feedback,
		UserDAO:     d.User,
	}
}

// Create ...
func (d *FeedbackAppService) Create(ctx context.Context, body model.FeedbackBody, user model.UserRaw) (doc model.FeedbackResponse, err error) {
	payload := body.NewFeedbackBSON(user.ID)
	err = d.FeedbackDAO.InsertOne(ctx, payload)
	if err != nil {
		return doc, errors.New(locale.FeedbackKeyCanNotCreate)
	}

	userInfo := model.UserInfo{
		ID:       user.ID,
		UserName: user.Username,
		Address:  user.Address,
	}

	return payload.GetResponse(userInfo), nil
}

// ListAll ...
func (d *FeedbackAppService) ListAll(ctx context.Context) ([]model.FeedbackResponse, int64) {
	var (
		wg   sync.WaitGroup
		cond = bson.M{
			"active": true,
		}
		res   = make([]model.FeedbackResponse, 0)
		total int64
	)

	total = d.FeedbackDAO.CountByCondition(ctx, cond)
	feedbacks, _ := d.FeedbackDAO.FindByCondition(ctx, cond)
	if len(feedbacks) > 0 {
		wg.Add(len(feedbacks))
		for index, feedback := range feedbacks {
			go func(f model.FeedbackRaw, i int) {
				defer wg.Done()
				user, _ := d.UserDAO.FindOneByCondition(ctx, bson.M{"_id": f.User})
				userInfo := model.UserInfo{
					ID:       user.ID,
					UserName: user.Username,
					Address:  user.Address,
				}

				temp := f.GetResponse(userInfo)
				res = append(res, temp)
			}(feedback, index)
		}
		wg.Wait()
	}

	return res, total
}

// Update ....
func (d *FeedbackAppService) Update(ctx context.Context, body model.FeedbackBody, user model.UserRaw, feedback model.FeedbackRaw) (res model.FeedbackResponse, err error) {
	doc := body.NewFeedbackBSON(user.ID)

	// assign
	feedback.Name = doc.Name
	feedback.Rating = doc.Rating
	feedback.Order = doc.Order
	feedback.UpdatedAt = doc.UpdatedAt

	err = d.FeedbackDAO.UpdateByID(ctx, feedback.ID, bson.M{"$set": feedback})
	if err != nil {
		return res, err
	}

	userInfo := model.UserInfo{
		ID:       user.ID,
		UserName: user.Username,
		Address:  user.Address,
	}

	return feedback.GetResponse(userInfo), nil
}

// FindByID ...
func (d *FeedbackAppService) FindByID(ctx context.Context, id model.AppID) (model.FeedbackRaw, error) {
	return d.FeedbackDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (d *FeedbackAppService) GetDetail(ctx context.Context, feedback model.FeedbackRaw) model.FeedbackResponse {
	user, _ := d.UserDAO.FindOneByCondition(ctx, bson.M{"_id": feedback.User})
	userInfo := model.UserInfo{
		ID:       user.ID,
		UserName: user.Username,
		Address:  user.Address,
	}
	return feedback.GetResponse(userInfo)
}
