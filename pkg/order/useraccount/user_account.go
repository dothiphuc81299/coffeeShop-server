package useraccount

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Create(ctx context.Context, body *CreateUserAccountCommand) (err error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) (UserAccountRaw, error)
}

type Store interface {
	InsertOne(ctx context.Context, u UserAccountRaw) error
	FindOneByCondition(ctx context.Context, cond interface{}) (u UserAccountRaw, err error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []UserAccountRaw, err error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) error
}
