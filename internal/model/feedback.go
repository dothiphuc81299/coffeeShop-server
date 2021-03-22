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

// FeedbackAdminService ...
type FeedbackAdminService interface {
	Create(ctx context.Context, userID primitive.ObjectID, body FeedbackBody) (FeedbackResponse, error)
	ListAll(ctx context.Context, q CommonQuery) ([]FeedbackResponse, int64)
	Update(ctx context.Context, userID primitive.ObjectID, Feedback FeedbackRaw, body FeedbackBody) (FeedbackResponse, error)
	FindByID(ctx context.Context, id AppID) (Feedback FeedbackRaw, err error)
}

type FeedbackRaw struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"createdAt"`
	Rating    int                `bson:"rating"`
	User      primitive.ObjectID `bson:"user"`
	Order     primitive.ObjectID `bson:"order"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
