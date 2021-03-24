package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (OrderRaw, error)
	InsertOne(ctx context.Context, u OrderRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]OrderRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// OrderService ...
type OrderAdminService interface {
	Create(ctx context.Context, userID primitive.ObjectID, body OrderBody) (OrderResponse, error)
}
	
type OrderRaw struct {
	ID         primitive.ObjectID   `bson:"_id"`
	User       primitive.ObjectID   `bson:"user"`
	Drink      []primitive.ObjectID `bson:"drink"`
	Status     bool                 `bson:"status"`
	TotalPrice float64              `bson:"totalPrice"`
	CreatedAt  time.Time            `bson:"createdAt"`
	UpdatedAt  time.Time            `bson:"updatedAt"`
}
