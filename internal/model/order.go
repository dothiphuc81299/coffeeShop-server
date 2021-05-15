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
type OrderAppService interface {
	Create(ctx context.Context, userID UserRaw, body OrderBody) (OrderResponse, error)
	FindByID(ctx context.Context, id AppID) (OrderRaw, error)
	GetList(ctx context.Context, query CommonQuery, user UserRaw) ([]OrderResponse, int64)
	GetDetail(ctx context.Context, order OrderRaw) OrderResponse
}

type OrderAdminService interface {
	ChangeStatus(ctx context.Context, order OrderRaw, status StatusBody, staff StaffRaw) (string, error)
	GetListByStatus(ctx context.Context, query CommonQuery) ([]OrderResponse, int64)
	FindByID(ctx context.Context, id AppID) (OrderRaw, error)
	GetDetail(ctx context.Context, order OrderRaw) OrderResponse
}

type OrderRaw struct {
	ID         primitive.ObjectID `bson:"_id"`
	User       primitive.ObjectID `bson:"user"`
	Drink      []DrinkInfo        `bson:"drink"`
	Status     string             `bson:"status"`
	TotalPrice float64            `bson:"totalPrice"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
	Shipper    primitive.ObjectID `bson:"shipper,omitempty"`
}
