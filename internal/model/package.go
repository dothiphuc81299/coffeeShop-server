package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PackageDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (PackageRaw, error)
	InsertOne(ctx context.Context, u PackageRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]PackageRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// PackageAdminService ....
type PackageAdminService interface {
	Create(ctx context.Context, body PackageBody) (PackageAdminResponse, error)
	ListAll(ctx context.Context, q CommonQuery) ([]PackageAdminResponse, int64)
	Update(ctx context.Context, c PackageRaw, body PackageBody) (PackageAdminResponse, error)
	FindByID(ctx context.Context, id AppID) (Package PackageRaw, err error)
	GetDetail(ctx context.Context, cate PackageRaw) PackageAdminResponse
}

type PackageRaw struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Level      Level              `bson:"level"` // easy , medium, hard
	Reward     float64            `bson:"reward"`
	NumberQuiz float64            `bson:"numberQuiz"`
	MinusPoint float64            `bson:"minusPoint"`
}
