package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserAppService struct {
	UserDAO model.UserDAO
	CodeDAO model.CodedRegisterDAO
}

// NewUserAppService ...
func NewUserAppService(d *model.CommonDAO) model.UserAppService {
	return &UserAppService{
		UserDAO: d.User,
		CodeDAO: d.CodeDAO,
	}
}

func (u *UserAppService) UserSignUp(ctx context.Context, body model.UserSignUpBody) (err error) {
	payload := body.NewUserRaw()

	// find db
	count := u.UserDAO.CountByCondition(ctx, bson.M{"username": payload.Username})
	if count > 0 {
		return errors.New(locale.CommonyKeyUserNameIsExisted)
	}

	countEmail := u.UserDAO.CountByCondition(ctx, bson.M{"email": payload.Email, "active": true})
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

func (u *UserAppService) SendEmail(ctx context.Context, mail model.UserSendEmailBody) error {
	code, _ := GenerateOTP(6)
	//err := redisapp.SetKeyValue(code, mail.Email, 24*time.Hour)

	argsCode := model.CodedRegisterRaw{
		Id:    model.NewAppID(),
		Email: mail.Email,
		Code:  code,
	}
	err := u.CodeDAO.InsertOne(ctx, argsCode)

	if err != nil {
		return err
	}
	fmt.Println(err)
	mailw := mail.Email
	args := &model.UserVerifyEmail{
		Email: mailw,
		Code:  code,
	}
	u.SendVerifyMemberEmail(args)

	return nil

}

const otpChars = "1234567890"

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func (u *UserAppService) VerifyEmail(ctx context.Context, args model.VerifyEmailBody) error {
	// result := redisapp.GetValueByKey(args.Code)
	// var res string
	// if result == "" {
	// 	return fmt.Errorf("code not found")
	// }
	// if err := json.Unmarshal([]byte(result), &res); err != nil {
	// 	return err
	// }

	// if res != args.Email {
	// 	fmt.Println("ok")
	// 	return fmt.Errorf("Email Khong hop le")
	// }

	result, err := u.CodeDAO.FindOneByCondition(ctx, bson.M{"code": args.Code})
	if err != nil {
		return err
	}

	if result.Email != args.Code {
		return fmt.Errorf(" Khong hop le")
	}

	// update active user
	err = u.UserDAO.UpdateByCondition(ctx, bson.M{"email": args.Email}, bson.M{"$set": bson.M{"active": true}})
	if err != nil {
		return err
	}

	// err = redisapp.DelKey(args.Code)
	err = u.CodeDAO.DeleteOne(ctx, args.Email)
	if err != nil {
		return err
	}
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

	err := u.UserDAO.UpdateByCondition(ctx, user.ID, bson.M{"$set": payload})
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
