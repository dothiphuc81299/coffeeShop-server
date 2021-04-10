package service

import (
	"context"

	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// UserAdminService ...
type UserAdminService struct {
	UserDAO model.UserDAO
}

// List ...
func (us *UserAdminService) GetList(ctx context.Context, q model.CommonQuery) ([]model.UserAdminResponse, int64) {
	var (
		wg     sync.WaitGroup
		result = make([]model.UserAdminResponse, 0)
		total  int64
		cond   = bson.M{}
	)
	// Assign cond
	q.AssignActive(&cond)
	q.AssignKeyword(&cond)
	total = us.UserDAO.CountByCondition(ctx, cond)
	docs, _ := us.UserDAO.FindByCondition(ctx, cond)
	if len(docs) > 0 {
		wg.Add(len(docs))
		result = make([]model.UserAdminResponse, len(docs))
		for index, user := range docs {
			go func(u model.UserRaw, i int) {
				defer wg.Done()
				temp := u.GetAdminResponse()
				result[i] = temp
			}(user, index)
		}
		wg.Wait()
	}

	return result, total
}

// ChangeStatus ...
func (us *UserAdminService) ConfirmAccountActive(ctx context.Context, user model.UserRaw) (active bool, err error) {

	payload := bson.M{
		"active": !user.Active,
	}

	err = us.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": payload})
	if err != nil {
		return user.Active, err
	}
	return !user.Active, nil
}

func (us *UserAdminService) FindByID(ctx context.Context, id model.AppID) (model.UserRaw, error) {
	return us.UserDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

// NewUserAdminService ...
func NewUserAdminService(d *model.CommonDAO) model.UserAdminService {
	return &UserAdminService{
		UserDAO: d.User,
	}
}
