package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PackageGroupDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (PackageGroupRaw, error)
	InsertOne(ctx context.Context, u PackageGroupRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]PackageGroupRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// PackageGroupAdminService ....
type PackageGroupAdminService interface {
	Create(ctx context.Context, body PackageGroupBody) (PackageGroupAdminResponse, error)
	ListAll(ctx context.Context, q CommonQuery) ([]PackageGroupAdminResponse, int64)
	Update(ctx context.Context, c PackageGroupRaw, body PackageGroupBody) (PackageGroupAdminResponse, error)
	FindByID(ctx context.Context, id AppID) (PackageGroup PackageGroupRaw, err error)
	GetDetail(ctx context.Context, cate PackageGroupRaw) PackageGroupAdminResponse
}

type PackageGroupRaw struct {
	ID         primitive.ObjectID `bson:"_id"`
	PackageID  primitive.ObjectID `bson:"packageId"`
	GroupID    primitive.ObjectID `bson:"groupId"`
	NumberQuiz float64            `bson:"numberQuiz"`
}
