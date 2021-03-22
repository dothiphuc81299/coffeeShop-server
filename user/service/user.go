package service

import (
	"context"
	"errors"

	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// UserAdminService ...
type UserAdminService struct {
	UserDAO model.UserDAO
}

// Create ...
func (us *UserAdminService) Create(ctx context.Context, body model.UserBody) (doc model.UserAdminResponse, err error) {
	// Check exist phone
	cond := bson.M{
		"phone": body.Phone,
	}
	userCount := us.UserDAO.CountByCondition(ctx, cond)
	if userCount > 0 {
		return doc, errors.New(locale.CommonKeyPhoneExisted)
	}

	// New
	payload := body.NewRaw()
	err = us.UserDAO.InsertOne(ctx, payload)
	if err != nil {
		return doc, errors.New(locale.CommonKeyCanNotCreateUser)
	}
	res := payload.GetAdminResponse()
	return res, err
}

// List ...
func (us *UserAdminService) List(ctx context.Context, q model.CommonQuery) ([]model.UserAdminResponse, int64) {
	var (
		wg     sync.WaitGroup
		result = make([]model.UserAdminResponse, 0)
		total  int64
		cond   = bson.M{}
	)
	// Assign cond
	q.AssignActive(&cond)
	q.AssignKeyword(&cond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		total = us.UserDAO.CountByCondition(ctx, cond)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs, _ := us.UserDAO.FindByCondition(ctx, cond)
		for _, u := range docs {
			user := u.GetAdminResponse()
			result = append(result, user)
		}
	}()

	wg.Wait()
	return result, total
}

// FindByID ...
func (us *UserAdminService) FindByID(ctx context.Context, id model.AppID) (model.UserRaw, error) {
	return us.UserDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

// Update ...
func (us *UserAdminService) Update(ctx context.Context, body model.UserBody, raw model.UserRaw) (model.UserAdminResponse, error) {

	raw.Phone = body.Phone
	raw.Avatar = body.Avatar.ConvertToFilePhoto()
	raw.UpdatedAt = time.Now()
	raw.Username = body.Username
	raw.Password = body.Password
	raw.Address = body.Address

	err := us.UserDAO.UpdateByID(ctx, raw.ID, bson.M{"$set": raw})

	doc, _ := us.UserDAO.FindOneByCondition(ctx, bson.M{"_id": raw.ID})
	return doc.GetAdminResponse(), err
}

// ChangeStatus ...
func (us *UserAdminService) ChangeStatus(ctx context.Context, raw model.UserRaw) (active bool, err error) {
	raw.Active = !raw.Active
	raw.UpdatedAt = time.Now()
	err = us.UserDAO.UpdateByID(ctx, raw.ID, bson.M{"$set": raw})
	if err != nil {
		return active, errors.New(locale.UserKeyCanNotChangeStatus)
	}
	return raw.Active, err
}

// GetDetail ...
func (us *UserAdminService) GetDetail(ctx context.Context, raw model.UserRaw) model.UserAdminResponse {
	return raw.GetAdminResponse()
}

// NewUserAdminService ...
func NewUserAdminService(d *model.CommonDAO) model.UserAdminService {
	return &UserAdminService{
		UserDAO: d.User,
	}
}
