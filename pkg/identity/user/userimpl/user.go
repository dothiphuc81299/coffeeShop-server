package userimpl

import (
	"context"
	"sync"

	orderapi "github.com/dothiphuc81299/coffeeShop-server/pkg/apis/order"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/code"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/password"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	store       *store
	codeStore   code.Store
	orderClient orderapi.OrderClient
}

func NewService(store *store, codeStore code.Store, orderClient orderapi.OrderClient) user.Service {
	return &service{
		store:       store,
		codeStore:   codeStore,
		orderClient: orderClient,
	}
}

func (s *service) CreateUser(ctx context.Context, body user.CreateUserCommand) (email string, err error) {
	payload := body.NewUserRaw()

	countEmail := s.store.CountByCondition(ctx, bson.M{"email": payload.Email, "active": true})
	if countEmail > 0 {
		return "", user.ErrEmailExisted
	}

	err = s.store.InsertOne(ctx, payload)
	if err != nil {
		return "", err
	}

	_, err = s.orderClient.CreateUserAccount(ctx, &orderapi.CreateUserAccountCommand{
		UserId:    payload.ID.Hex(),
		LoginName: payload.Username,
		Active:    true,
	})
	if err != nil {
		return "", err
	}

	// s.SendEmail(ctx, user.SendUserEmailCommand{
	// 	Email: body.Email,
	// })

	return payload.Email, nil
}

func (s *service) LoginUser(ctx context.Context, body user.CreateLoginUserCommand) (doc user.CreateLoginUserResult, err error) {
	cond := bson.M{
		"username": body.Username,
		"active":   true,
	}

	result, err := s.store.FindOneByCondition(ctx, cond)
	if err != nil {
		return doc, err
	}

	if !password.CheckPassword(body.Password, result.Password) {
		return doc, user.ErrPasswordInvalid
	}

	jwt, err := token.GenerateJWT(result.ID, result.Username, token.User)
	if err != nil {
		return doc, err
	}
	doc = result.GetLoginUserResponse(jwt)
	return doc, nil
}

func (s *service) UpdateUser(ctx context.Context, cmd *user.UpdateUserCommand) error {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return user.ErrAccountIsInvalid
	}

	payload := bson.M{
		"phone":   cmd.Phone,
		"address": cmd.Address,
	}

	err := s.store.UpdateByCondition(ctx, bson.M{"_id": account.ID}, bson.M{"$set": payload})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetDetailUser(ctx context.Context, id primitive.ObjectID) user.CreateLoginUserResult {
	doc, _ := s.store.FindOneByCondition(ctx, bson.M{"_id": id})
	jwt, err := token.GenerateJWT(doc.ID, doc.Username, token.User)
	if err != nil {
		return user.CreateLoginUserResult{}
	}
	result := doc.GetLoginUserResponse(jwt)
	return result
}

func (s *service) ChangePassword(ctx context.Context, cmd *user.ChangePasswordUserCommand) (err error) {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return user.ErrAccountIsInvalid
	}

	res, err := s.store.FindOneByCondition(ctx, bson.M{"_id": account.ID})
	if err != nil {
		return err
	}

	if res.ID.IsZero() {
		return user.ErrUserNotFound
	}

	if !password.CheckPassword(cmd.Password, res.Password) {
		return user.ErrPasswordInvalid
	}

	if cmd.NewPassword != cmd.NewPasswordAgain || cmd.NewPassword == cmd.Password {
		return user.ErrPasswordInvalid
	}

	hashPassword, err := password.HashPassword(cmd.NewPassword)
	if err != nil {
		return err
	}

	err = s.store.UpdateByID(ctx, account.ID, bson.M{"$set": bson.M{"password": hashPassword}})
	if err != nil {
		return
	}
	return
}

func (s *service) Search(ctx context.Context, q *query.CommonQuery) ([]user.UserRaw, int64) {
	var (
		wg    sync.WaitGroup
		res   = make([]user.UserRaw, 0)
		total int64
		cond  = bson.M{}
	)

	q.AssignActive(&cond)
	q.AssignKeyword(&cond)

	wg.Add(2)
	go func() {
		defer wg.Done()
		docs, _ := s.store.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
		res = docs
	}()
	go func() {
		defer wg.Done()
		total = s.store.CountByCondition(ctx, cond)
	}()
	wg.Wait()
	return res, total
}

func (s *service) FindByID(ctx context.Context, id primitive.ObjectID) (user.UserRaw, error) {
	return s.store.FindOneByCondition(ctx, bson.M{"_id": id})
}
