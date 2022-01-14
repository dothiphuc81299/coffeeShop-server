package service

import (
	"context"
	"errors"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
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

func (u *UserAppService) UserSignUp(ctx context.Context, body model.UserSignUpBody) (err error) {
	payload := body.NewUserRaw()

	// find db
	count := u.UserDAO.CountByCondition(ctx, bson.M{"username": payload.Username})
	if count > 0 {
		return errors.New(locale.CommonyKeyUserNameIsExisted)
	}

	countEmail := u.UserDAO.CountByCondition(ctx, bson.M{"email": payload.Email})
	if countEmail > 0 {
		return errors.New(locale.CommonyKeyEmailIsExisted)
	}

	err = u.UserDAO.InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	// save 
	return nil
}

func (u *UserAppService) UserLoginIn(ctx context.Context, body model.UserLoginBody) (doc model.UserLoginResponse, err error) {
	cond := bson.M{
		"username": body.Username,
		"password": body.Password,
	}

	user, err := u.UserDAO.FindOneByCondition(ctx, cond)
	if err != nil {
		return doc, errors.New(locale.UserNameOrPasswordIsIncorrect)
	}

	token := user.GenerateToken()
	doc = user.GetUserLoginInResponse(token)
	return doc, nil
}

func (u *UserAppService) UserUpdateAccount(ctx context.Context, user model.UserRaw, body model.UserUpdateBody) error {
	payload := bson.M{
		"phone":   body.Phone,
		"address": body.Address,
		// "avatar":   body.Avatar,
	}

	err := u.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": payload})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserAppService) GetDetailUser(ctx context.Context, user model.UserRaw) model.UserLoginResponse {
	doc, _ := u.UserDAO.FindOneByCondition(ctx, bson.M{"_id": user.ID})
	token := user.GenerateToken()
	res := doc.GetUserLoginInResponse(token)
	return res
}

func (u *UserAppService) ChangePassword(ctx context.Context, user model.UserRaw, body model.UserChangePasswordBody) (err error) {
	res, _ := u.UserDAO.FindOneByCondition(ctx, bson.M{"_id": user.ID})
	if res.ID.IsZero() {
		return errors.New(locale.UserIsNotExisted)
	}

	if body.Password != res.Password || body.NewPassword != body.NewPasswordAgain || body.NewPassword == body.Password {
		return errors.New(locale.PasswordIsIncorrect)
	}

	err = u.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"password": body.NewPassword}})
	if err != nil {
		return
	}
	return
}
