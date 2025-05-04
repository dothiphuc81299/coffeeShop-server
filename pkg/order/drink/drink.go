package drink

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (DrinkRaw, error)
	InsertOne(ctx context.Context, u DrinkRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]DrinkRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
	DeleteByCategoryID(ctx context.Context, categoryID primitive.ObjectID) error
}

type Service interface {
	Create(ctx context.Context, body DrinkBody) error
	ListAll(ctx context.Context, q query.CommonQuery) ([]DrinkAdminResponse, int64)
	Update(ctx context.Context, id primitive.ObjectID, body DrinkBody) error
	ChangeStatus(ctx context.Context, Drink DrinkRaw) (bool, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (Drink DrinkRaw, err error)
	GetDetail(ctx context.Context, drink DrinkRaw) DrinkAdminResponse
	DeleteDrink(ctx context.Context, drink DrinkRaw) error
}
