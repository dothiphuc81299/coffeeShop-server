package category

import (
	"context"

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
	Create(ctx context.Context, cmd *CategoryBody) error
	ListAll(ctx context.Context, q *SearchCategoryQuery) ([]CategoryRaw, int64,error)
	Update(ctx context.Context, id primitive.ObjectID, body CategoryBody) error
	GetDetail(ctx context.Context, id primitive.ObjectID) (CategoryRaw, error)
	DeleteCategory(ctx context.Context, id primitive.ObjectID) error
}
