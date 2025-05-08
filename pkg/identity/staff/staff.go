package staff

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (Staff, error)
	InsertOne(ctx context.Context, u Staff) error
	FindByID(ctx context.Context, id primitive.ObjectID, opts ...*options.FindOneOptions) (Staff, error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]Staff, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	UpdateBycondition(ctx context.Context, cond interface{}, payload interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type Service interface {
	Create(ctx context.Context, body CreateStaffCommand) error
	ListStaff(ctx context.Context, q *query.CommonQuery) ([]Staff, int64)
	LoginStaff(ctx context.Context, LoginStaff LoginStaffCommand) (*StaffResponse, error)
	GetStaffByID(ctx context.Context, id primitive.ObjectID) (Staff, error)
	UpdateRole(ctx context.Context, cmd *UpdateStaffRoleCommand) error
	Update(ctx context.Context, cmd *UpdateStaffCommand) error
	ChangePassword(ctx context.Context, body *PasswordBody) error
}
