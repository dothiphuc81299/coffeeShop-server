package shippingaddressimpl

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	store     *store
	userStore useraccount.Store
}

func NewService(store *store, userStore useraccount.Store) shippingaddress.Service {
	return &service{
		store:     store,
		userStore: userStore,
	}
}

func (s *service) Create(ctx context.Context, cmd *shippingaddress.CreateShippingAddressCommand) error {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return shippingaddress.ErrAccountIsInvalid
	}

	userAcc, err := s.userStore.FindOneByCondition(ctx, bson.M{"user_id": account.ID})
	if err != nil {
		return err
	}

	cmd.UserID = userAcc.UserID

	return s.store.InsertOne(ctx, shippingaddress.UserShippingAddressRaw{
		ID:        primitive.NewObjectID(),
		UserID:    cmd.UserID,
		FullName:  cmd.FullName,
		Phone:     cmd.Phone,
		Address:   cmd.Address,
		Province:  cmd.Province,
		City:      cmd.City,
		Ward:      cmd.Ward,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
}

func (s *service) Update(ctx context.Context, cmd *shippingaddress.UpdateShippingAddressCommand) error {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return shippingaddress.ErrAccountIsInvalid
	}

	result, err := s.store.FindOneByCondition(ctx, bson.M{"_id": cmd.ID})
	if err != nil {
		return err
	}

	if result.ID.IsZero() {
		return shippingaddress.ErrShippingAddressNotFound
	}
	err = s.store.UpdateOne(ctx, bson.M{"_id": cmd.ID}, bson.M{"$set": bson.M{
		"full_name":  cmd.FullName,
		"phone":      cmd.Phone,
		"address":    cmd.Address,
		"province":   cmd.Province,
		"city":       cmd.City,
		"ward":       cmd.Ward,
		"updated_at": time.Now().UTC(),
	}})

	return err
}

func (s *service) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.store.DeleteOne(ctx, id)
}

func (s *service) GetDetail(ctx context.Context, id primitive.ObjectID) (shippingaddress.UserShippingAddressRaw, error) {
	return s.store.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (s *service) Search(ctx context.Context, query *shippingaddress.SearchShippingAddressQuery) ([]shippingaddress.UserShippingAddressRaw, int64, error) {

	var (
		wg      sync.WaitGroup
		mutex   sync.Mutex
		cond    = bson.M{}
		total   int64
		res     = make([]shippingaddress.UserShippingAddressRaw, 0)
		errList []error
	)

	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return nil, 0, shippingaddress.ErrAccountIsInvalid
	}

	cond["user_id"] = account.ID

	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	skip := (query.Page - 1) * query.Limit

	findOpts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(query.Limit))

	wg.Add(2)

	go func() {
		defer wg.Done()
		count, _ := s.store.CountByCondition(ctx, cond)
		mutex.Lock()
		total = count
		mutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		result, err := s.store.FindByCondition(ctx, cond, findOpts)
		if err != nil {
			mutex.Lock()
			errList = append(errList, err)
			mutex.Unlock()
			return
		}
		mutex.Lock()
		res = result
		mutex.Unlock()
	}()

	wg.Wait()

	if len(errList) > 0 {
		return nil, 0, fmt.Errorf("error(s) occurred: %v", errList)
	}

	return res, total, nil
}
