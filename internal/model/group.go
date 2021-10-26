package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GroupDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (QuizGroupRaw, error)
	InsertOne(ctx context.Context, u QuizGroupRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]QuizGroupRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// GroupAdminService ....
type GroupAdminService interface {
	Create(ctx context.Context, body QuizGroupBody) error
	ListAll(ctx context.Context, q CommonQuery) ([]QuizGroupCommon, int64)
	Update(ctx context.Context, c QuizGroupRaw, body QuizGroupBody) error
	FindByID(ctx context.Context, id AppID) (Group QuizGroupRaw, err error)
	//GetDetail(ctx context.Context, cate QuizGroupRaw) GroupAdminResponse
	ChangeStatus(ctx context.Context, c QuizGroupRaw) (bool, error)
}

type QuizGroupRaw struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Name          string               `bson:"name"`
	Quizzes       []primitive.ObjectID `bson:"quizzes"`
	Active        bool                 `bson:"active"`
	CreatedAt     time.Time            `bson:"createdAt"`
	UpdatedAt     time.Time            `bson:"updatedAt"`
	TotalQuestion float64              `bson:"totalQuestion"`
}
