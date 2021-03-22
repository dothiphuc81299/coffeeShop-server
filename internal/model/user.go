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

// UserAdminService ...
type UserAdminService interface {
	Create(ctx context.Context, body UserBody) (UserAdminResponse, error)
	List(ctx context.Context, q CommonQuery) ([]UserAdminResponse, int64)
	FindByID(ctx context.Context, id AppID) (UserRaw, error)
	Update(ctx context.Context, body UserBody, raw UserRaw) (UserAdminResponse, error)
	ChangeStatus(ctx context.Context, raw UserRaw) (bool, error)
	GetDetail(ctx context.Context, raw UserRaw) UserAdminResponse
}

// UserRaw ...
type UserRaw struct {
	ID        AppID      `bson:"_id"`
	Username  string     `bson:"username"`
	Password  string     `bson:"password"`
	Phone     string     `bson:"phone"`
	Active    bool       `bson:"active"`
	Avatar    *FilePhoto `bson:"avatar"`
	IsRoot    bool       `bson:"isRoot"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	Address   string     `bson:"address"`
}

// GetAdminResponse ...
func (u *UserRaw) GetAdminResponse() UserAdminResponse {
	return UserAdminResponse{
		ID:        u.ID,
		UserName:  u.Username,
		Password:  u.Password,
		Phone:     u.Phone,
		Active:    u.Active,
		Avatar:    u.Avatar.GetResponseData(),
		IsRoot:    u.IsRoot,
		CreatedAt: u.CreatedAt,
		Address:   u.Address,
	}
}

// GenerateToken generate token for authentication
func (u *UserRaw) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      u.ID,
		"username": u.Username,
		"phone":    u.Phone,
		"exp":      time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte(config.GetEnv().AuthSecret))
	return tokenString
}
