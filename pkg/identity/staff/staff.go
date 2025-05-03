package staff

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (Staff, error)
	InsertOne(ctx context.Context, u Staff) error
	FindByID(ctx context.Context, id primitive.ObjectID) (Staff, error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]Staff, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	UpdateBycondition(ctx context.Context, cond interface{}, payload interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

type Service interface {
	Create(ctx context.Context, body *CreateStaffCommand) error
	ListStaff(ctx context.Context, q query.CommonQuery) (*SearchStaffResult, int64)
	FindByID(ctx context.Context, id primitive.ObjectID) (Staff, error)
	DeleteStaff(ctx context.Context, raw Staff) error
	GetToken(ctx context.Context, staffID primitive.ObjectID) (string, error)
	GetDetailStaff(ctx context.Context, staff Staff) *Staff
	LoginStaff(ctx context.Context, LoginStaff LoginStaffCommand) (LoginStaffResult, error)
	GetStaffByID(ctx context.Context, id primitive.ObjectID) Staff

	Update(ctx context.Context, body UpdateStaffCommand, raw Staff) error
	ChangePassword(ctx context.Context, staff Staff, body PasswordBody) error
}
