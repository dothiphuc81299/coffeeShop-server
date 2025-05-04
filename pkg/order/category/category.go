package category

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (CategoryRaw, error)
	InsertOne(ctx context.Context, u CategoryRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]CategoryRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type Service interface {
	Create(ctx context.Context, body CategoryBody) ( error)
	ListAll(ctx context.Context, q query.CommonQuery) ([]CategoryRaw, int64)
	Update(ctx context.Context,id primitive.ObjectID, body CategoryBody) ( error)
	GetDetail(ctx context.Context, id primitive.ObjectID) (CategoryRaw,error)
	DeleteCategory(ctx context.Context, id primitive.ObjectID) error
}
