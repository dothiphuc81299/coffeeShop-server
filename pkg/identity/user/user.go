package user

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (UserRaw, error)
	InsertOne(ctx context.Context, u UserRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]UserRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
	UpdateByCondition(ctx context.Context, cond, payload interface{}) error
}

type Service interface {
	CreateUser(ctx context.Context, body CreateUserCommand) (string, error)
	LoginUser(ctx context.Context, body CreateLoginUserCommand) (CreateLoginUserResult, error)
	UpdateUser(ctx context.Context, user UserRaw, body UpdateUserCommand) error
	GetDetailUser(ctx context.Context, user UserRaw) CreateLoginUserResult
	ChangePassword(ctx context.Context, user UserRaw, body ChangePasswordUserCommand) error
	SendEmail(ctx context.Context, mail SendUserEmailCommand) error
	VerifyEmail(ctx context.Context, mail VerifyEmailCommand) error
	Search(ctx context.Context, query query.CommonQuery) ([]UserRaw, int64)
	FindByID(ctx context.Context, id AppID) (UserRaw, error)
}

func (u *UserRaw) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      u.ID,
		"username": u.Username,
		"phone":    u.Phone,
		//	"exp":      time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte(config.GetEnv().AuthSecret))
	return tokenString
}
