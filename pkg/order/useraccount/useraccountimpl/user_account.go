package useraccountimpl

import (
	"context"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	store *store
}

func NewService(store *store) useraccount.Service {
	return &service{
		store: store,
	}
}

func (s *service) Create(ctx context.Context, body *useraccount.CreateUserAccountCommand) (err error) {
	entity := useraccount.UserAccountRaw{
		ID:           primitive.NewObjectID(),
		UserID:       body.UserID,
		LoginName:    body.LoginName,
		Active:       body.Active,
		CurrentPoint: 0,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	return s.store.InsertOne(ctx, entity)
}

func (s *service) GetByUserID(ctx context.Context, userID primitive.ObjectID) (useraccount.UserAccountRaw, error) {
	cond := bson.M{
		"user_id": userID,
	}

	return s.store.FindOneByCondition(ctx, cond)
}
