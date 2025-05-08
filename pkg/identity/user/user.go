package user

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (UserRaw, error)
	InsertOne(ctx context.Context, u UserRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]UserRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	UpdateByCondition(ctx context.Context, cond, payload interface{}) error
}

type Service interface {
	CreateUser(ctx context.Context, body CreateUserCommand) (string, error)
	LoginUser(ctx context.Context, body CreateLoginUserCommand) (CreateLoginUserResult, error)
	UpdateUser(ctx context.Context, cmd *UpdateUserCommand) error
	GetDetailUser(ctx context.Context, id primitive.ObjectID) CreateLoginUserResult
	ChangePassword(ctx context.Context, cmd *ChangePasswordUserCommand) error
	SendEmail(ctx context.Context, mail SendUserEmailCommand) error
	VerifyEmail(ctx context.Context, mail VerifyEmailCommand) error
	Search(ctx context.Context, query *query.CommonQuery) ([]UserRaw, int64)
	FindByID(ctx context.Context, id primitive.ObjectID) (UserRaw, error)
}
