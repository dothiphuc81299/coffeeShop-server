package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShiftAnalyticRaw struct {
	ID         primitive.ObjectID `bson:"_id"`
	Staff      AppID              `bson:"staff"`
	TotalShift float64            `bson:"totalShift"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

type ShiftAnalyticDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (ShiftAnalyticRaw, error)
	InsertOne(ctx context.Context, u ShiftAnalyticRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]ShiftAnalyticRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

