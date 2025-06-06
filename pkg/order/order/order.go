package order

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
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
	Create(ctx context.Context, body OrderBody) (OrderResponse, error)
	Search(ctx context.Context, query *SearchOrdersQuery) ([]OrderResponse, int64)
	GetDetail(ctx context.Context, id primitive.ObjectID) (OrderResponse, error)
	RejectOrder(ctx context.Context, cmd *UpdateOrderStatusCommand) error
	ApproveOrder(ctx context.Context, cmd *UpdateOrderStatusCommand) error
	GetStatistic(ctx context.Context, query query.CommonQuery) (StatisticResponse, error)
}
