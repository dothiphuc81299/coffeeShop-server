package role

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (StaffRoleRaw, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (StaffRoleRaw, error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []StaffRoleRaw, err error)
	InsertOne(ctx context.Context, u *StaffRoleRaw) error
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type Service interface {
	FindByID(ctx context.Context, id primitive.ObjectID) (StaffRoleRaw, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	Update(ctx context.Context, id primitive.ObjectID, body CreateStaffRoleCommand) error
	ListStaffRole(ctx context.Context, q query.CommonQuery) ([]StaffRoleRaw, int64)
	Create(ctx context.Context, body CreateStaffRoleCommand) error
}
