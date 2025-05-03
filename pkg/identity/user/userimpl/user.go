package userimpl

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/code"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	UserDAO user.UserDAO
	CodeDAO code.CodedRegisterDAO
}

func NewService(dao *service) user.Service {
	return &service{
		UserDAO: dao.UserDAO,
		CodeDAO: dao.CodeDAO,
	}
}

func (s *service) CreateUser(ctx context.Context, body user.CreateUserCommand) (email string, err error) {
	payload := body.NewUserRaw()

	countEmail := s.UserDAO.CountByCondition(ctx, bson.M{"email": payload.Email, "active": true})
	if countEmail > 0 {
		return "", errors.New(locale.CommonyKeyEmailIsExisted)
	}

	err = s.UserDAO.InsertOne(ctx, payload)
	if err != nil {
		return "", err
	}

	s.SendEmail(ctx, user.SendUserEmailCommand{
		Email: body.Email,
	})

	return payload.Email, nil
}

func (s *service) SendEmail(ctx context.Context, mail user.SendUserEmailCommand) error {
	otp, _ := GenerateOTP(6)
	//err := redisapp.SetKeyValue(code, mail.Email, 24*time.Hour)

	argsCode := code.CodedRegisterRaw{
		Id:    primitive.NewObjectID(),
		Email: mail.Email,
		Code:  otp,
	}
	err := s.CodeDAO.InsertOne(ctx, argsCode)

	if err != nil {
		return err
	}
	fmt.Println(err)
	mailw := mail.Email
	args := &user.UserVerifyEmail{
		Email: mailw,
		Code:  otp,
	}
	s.SendVerifyMemberEmail(args)

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

func (s *service) VerifyEmail(ctx context.Context, args user.VerifyEmailCommand) error {
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

	result, err := s.CodeDAO.FindOneByCondition(ctx, bson.M{"email": args.Email})
	if err != nil {
		return err
	}

	if result.Code != args.Code {
		return fmt.Errorf(" Khong hop le")
	}

	err = s.UserDAO.UpdateByCondition(ctx, bson.M{"email": args.Email}, bson.M{"$set": bson.M{"active": true}})
	if err != nil {
		return err
	}

	// err = redisapp.DelKey(args.Code)
	err = s.CodeDAO.DeleteOne(ctx, args.Email)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) LoginUser(ctx context.Context, body user.CreateLoginUserCommand) (doc user.CreateLoginUserResult, err error) {
	cond := bson.M{
		"username": body.Username,
		"password": body.Password,
		"active":   true,
	}

	user, err := s.UserDAO.FindOneByCondition(ctx, cond)
	if err != nil {
		return doc, errors.New(locale.UserNameOrPasswordIsIncorrect)
	}

	token := user.GenerateToken()
	doc = user.GetLoginUserResponse(token)
	return doc, nil
}

func (s *service) UpdateUser(ctx context.Context, entity user.UserRaw, body user.UpdateUserCommand) error {
	payload := bson.M{
		"phone":   body.Phone,
		"address": body.Address,
	}

	err := s.UserDAO.UpdateByCondition(ctx, bson.M{"_id": entity.ID}, bson.M{"$set": payload})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetDetailUser(ctx context.Context, entity user.UserRaw) user.CreateLoginUserResult {
	doc, _ := s.UserDAO.FindOneByCondition(ctx, bson.M{"_id": entity.ID})
	token := entity.GenerateToken()
	res := doc.GetLoginUserResponse(token)
	return res
}

func (s *service) ChangePassword(ctx context.Context, entity user.UserRaw, body user.ChangePasswordUserCommand) (err error) {
	res, _ := s.UserDAO.FindOneByCondition(ctx, bson.M{"_id": entity.ID})
	if res.ID.IsZero() {
		return errors.New(locale.UserIsNotExisted)
	}

	if body.Password != res.Password || body.NewPassword != body.NewPasswordAgain || body.NewPassword == body.Password {
		return errors.New(locale.PasswordIsIncorrect)
	}

	err = s.UserDAO.UpdateByID(ctx, entity.ID, bson.M{"$set": bson.M{"password": body.NewPassword}})
	if err != nil {
		return
	}
	return
}

func (s *service) Search(ctx context.Context, q query.CommonQuery) ([]user.UserRaw, int64) {
	var (
		// wg     sync.WaitGroup
		// result = make([]user.UserRaw, 0)
		total int64
		cond  = bson.M{}
	)

	q.AssignActive(&cond)
	q.AssignKeyword(&cond)
	total = s.UserDAO.CountByCondition(ctx, cond)
	docs, _ := s.UserDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPageOne())
	// if len(docs) > 0 {
	// 	wg.Add(len(docs))
	// 	result = make([]user.UserRaw, len(docs))
	// 	for index, value := range docs {
	// 		go func(u user.UserRaw, i int) {
	// 			defer wg.Done()
	// 			result[i] = u
	// 		}(value, index)
	// 	}
	// 	wg.Wait()
	// }

	return docs, total
}

func (s *service) FindByID(ctx context.Context, id user.AppID) (user.UserRaw, error) {
	return s.UserDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}
