package order

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (OrderRaw, error)
	InsertOne(ctx context.Context, u OrderRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]OrderRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id, payload interface{}) error
	AggregateOrder(ctx context.Context, cond interface{}) ([]*StatisticByDrink, error)
}

type Service interface {
	Create(ctx context.Context, userID UserRaw, body OrderBody) (OrderResponse, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (OrderRaw, error)
	Search(ctx context.Context, query CommonQuery, user UserRaw) ([]OrderResponse, int64)
	GetDetail(ctx context.Context, order OrderRaw) OrderResponse
	RejectOrder(ctx context.Context, user UserRaw, order OrderRaw) error
}

type OrderAdminService interface {
	ChangeStatus(ctx context.Context, order OrderRaw, status StatusBody, staff Staff) (string, error)
	UpdateOrderSuccess(ctx context.Context, order OrderRaw, staff Staff) error
	CancelOrder(ctx context.Context, order OrderRaw, staff Staff) error
	SearchByStatus(ctx context.Context, query CommonQuery) ([]OrderResponse, int64)
	FindByID(ctx context.Context, id primitive.ObjectID) (OrderRaw, error)
	GetDetail(ctx context.Context, order OrderRaw) OrderResponse
	GetStatistic(ctx context.Context, query CommonQuery) (StatisticResponse, error)
}
