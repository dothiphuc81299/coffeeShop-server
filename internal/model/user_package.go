package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserPackageGroupDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (UserPackageGroupRaw, error)
	InsertOne(ctx context.Context, u UserPackageGroupRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]UserPackageGroupRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}
