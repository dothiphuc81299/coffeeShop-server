package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (QuestionRaw, error)
	InsertOne(ctx context.Context, u QuestionRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]QuestionRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// QuestionAdminService ....
type QuestionAdminService interface {
	Create(ctx context.Context, body QuestionBody) error
	ListAll(ctx context.Context, q CommonQuery) ([]QuestionCommon, int64)
	Update(ctx context.Context, c QuestionRaw, body QuestionBodyUpdate) error
	FindByID(ctx context.Context, id AppID) (Question QuestionRaw, err error)
	GetDetail(ctx context.Context, q QuestionRaw) QuestionCommon
	ChangeStatus(ctx context.Context, q QuestionRaw) (bool, error)
}

type QuestionRaw struct {
	ID        primitive.ObjectID    `bson:"_id"`
	Order     int                   `bson:"order"`
	Question  string                `bson:"question"`
	Answers   []QuestionAnswersBSON `bson:"answers"`
	Active    bool                  `bson:"active"`
	CreatedAt time.Time             `bson:"createdAt"`
	UpdatedAt time.Time             `bson:"updatedAt"`
}

// QuestionAnswersBSON  ...
type QuestionAnswersBSON struct {
	ID      primitive.ObjectID `bson:"_id"`
	Answer  string             `bson:"answer"`
	Correct bool               `bson:"correct"`
	Order   int                `bson:"order"`
}
