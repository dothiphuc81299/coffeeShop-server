package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FeedbackDAO ...
type FeedbackDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (FeedbackRaw, error)
	InsertOne(ctx context.Context, u FeedbackRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]FeedbackRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// FeedbackAppService ...
type FeedbackAppService interface {
	Create(ctx context.Context, body FeedbackBody, user UserRaw) (FeedbackResponse, error)
	ListAll(ctx context.Context) ([]FeedbackResponse, int64)
	Update(ctx context.Context, body FeedbackBody, userID UserRaw, feedback FeedbackRaw) (FeedbackResponse, error)
	FindByID(ctx context.Context, id AppID) (FeedbackRaw, error)
	GetDetail(ctx context.Context, order FeedbackRaw) FeedbackResponse
}

type FeedbackAdminService interface {
	ChangeStatus(ctx context.Context, raw FeedbackRaw) (bool, error)
	FindByID(ctx context.Context, id AppID) (FeedbackRaw, error)
}
type FeedbackRaw struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"createdAt"`
	Rating    int                `bson:"rating"`
	User      primitive.ObjectID `bson:"user"`
	Order     primitive.ObjectID `bson:"order"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Active    bool               `bson:"active"`
	Drink     primitive.ObjectID `bson:"drink"`
}
