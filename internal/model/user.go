package model

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserDAO ...
type UserDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (UserRaw, error)
	InsertOne(ctx context.Context, u UserRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]UserRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// UserAppService ...
type UserAppService interface {
	UserSignUp(ctx context.Context, body UserSignUpBody) (string, error)
	UserLoginIn(ctx context.Context, body UserLoginBody) (UserLoginResponse, error)
	UserUpdateAccount(ctx context.Context, user UserRaw, body UserUpdateBody) (string, error)
	GetDetailUser(ctx context.Context, user UserRaw) UserLoginResponse
	ChangePassword(ctx context.Context,user UserRaw,body UserChangePasswordBody) error
}

type UserAdminService interface {
	GetList(ctx context.Context, query CommonQuery) ([]UserAdminResponse, int64)
	ConfirmAccountActive(ctx context.Context, user UserRaw) (bool, error)
	FindByID(ctx context.Context, id AppID) (UserRaw, error)
}

// UserRaw ...
type UserRaw struct {
	ID           AppID     `bson:"_id"`
	Username     string    `bson:"username"`
	Password     string    `bson:"password"`
	Phone        string    `bson:"phone"`
	Active       bool      `bson:"active"`
	Avatar       string    `bson:"avatar"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
	Address      string    `bson:"address"`
	SearchString string    `bson:"searchString"`
}

// GetAdminResponse ...
func (u *UserRaw) GetAdminResponse() UserAdminResponse {
	return UserAdminResponse{
		ID:        u.ID,
		UserName:  u.Username,
		Phone:     u.Phone,
		Active:    u.Active,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt,
		Address:   u.Address,
	}
}

// GenerateToken generate token for authentication
func (u *UserRaw) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      u.ID,
		"username": u.Username,
		"password": u.Password,
		"exp":      time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte(config.GetEnv().AuthSecret))
	return tokenString
}
