package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (AccountRaw, error)
	InsertOne(ctx context.Context, u AccountRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]AccountRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// AccountService ...
type AccountAdminService interface {
	FindByID(ctx context.Context, id AppID) (AccountRaw, error)
	Update(ctx context.Context, body AccountBody, raw AccountRaw) (AccountResponse, error)
	ChangeStatus(ctx context.Context, account AccountRaw) (bool, error)
	GenerateToken(ctx context.Context, account AccountRaw, id AppID) (string, error)
}

type AccountRaw struct {
	ID          primitive.ObjectID `bson:"_id"`
	Active      bool               `bson:"active"`
	User        primitive.ObjectID `bson:"user"`
	Permissions []string           `bson:"permissions"`
	Role        primitive.ObjectID `bson:"role"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}
