package service

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserAppService struct {
	UserDAO model.UserDAO
}

// NewUserAppService ...
func NewUserAppService(d *model.CommonDAO) model.UserAppService {
	return &UserAppService{
		UserDAO: d.User,
	}
}

func (u *UserAppService) UserSignUp(ctx context.Context, body model.UserSignUpBody) (string, error) {
	payload := body.NewUserRaw()

	err := u.UserDAO.InsertOne(ctx, payload)
	if err != nil {
		return "tao tai khoan that bai", err
	}

	return "tao tai khoan thanh cong", nil
}

func (u *UserAppService) UserLoginIn(ctx context.Context, body model.UserLoginBody) (doc model.UserLoginResponse, err error) {
	cond := bson.M{
		"username": body.Username,
		"password": body.Password,
	}

	user, err := u.UserDAO.FindOneByCondition(ctx, cond)
	if err != nil {
		return doc, err
	}

	token := user.GenerateToken()
	doc = user.GetUserLoginInResponse(token)
	return doc, nil
}

func (u *UserAppService) UserUpdateAccount(ctx context.Context, user model.UserRaw, body model.UserSignUpBody) (string, error) {
	payload := body.NewUserRaw()

	// assign
	user.Password = payload.Password
	user.Phone = payload.Phone
	user.Username = payload.Username
	user.Address = payload.Address
	user.Avatar = payload.Avatar

	err := u.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": user})
	if err != nil {
		return "khong the cap nhat tai khoan", err
	}

	return "thanh cong", nil
}

func (u *UserAppService) GetDetailUser(ctx context.Context, user model.UserRaw) model.UserLoginResponse {
	doc, _ := u.UserDAO.FindOneByCondition(ctx, bson.M{"_id": user.ID})
	token := user.GenerateToken()
	res := doc.GetUserLoginInResponse(token)
	return res
}
