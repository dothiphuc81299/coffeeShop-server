package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDAO ...
type CodedRegisterDAO interface {
	InsertOne(ctx context.Context, u CodedRegisterRaw) error
	DeleteOne(ctx context.Context, u string) error
	FindOneByCondition(ctx context.Context, cond interface{}) (CodedRegisterRaw, error)
}

type CodedRegisterRaw struct {
	Id    primitive.ObjectID `bson:"_id"`
	Email string             `bson:"email"`
	Code  string             `bson:"code"`
}
