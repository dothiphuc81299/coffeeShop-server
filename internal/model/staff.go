package model

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StaffDAO ...
type StaffDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (StaffRaw, error)
	InsertOne(ctx context.Context, u StaffRaw) error
	FindByID(ctx context.Context, id AppID) (StaffRaw, error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]StaffRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// StaffAdminService ...
type StaffAdminService interface {
	Create(ctx context.Context, body StaffBody) (StaffGetResponseAdmin, error)
	ListStaff(ctx context.Context, q CommonQuery) ([]StaffGetResponseAdmin, int64)
	FindByID(ctx context.Context, id AppID) (StaffRaw, error)
	Update(ctx context.Context, body StaffBody, raw StaffRaw) (StaffGetResponseAdmin, error)
	ChangeStatus(ctx context.Context, raw StaffRaw) (bool, error)
	GetToken(ctx context.Context, staffID AppID) (string, error)
}

// StaffRaw ...
type StaffRaw struct {
	ID          AppID      `bson:"_id"`
	Username    string     `bson:"username"`
	Password    string     `bson:"password"`
	Address     string     `bson:"address"`
	Phone       string     `bson:"phone"`
	Active      bool       `bson:"active"`
	Role        AppID      `bson:"role,omitempty"`
	Avatar      *FilePhoto `bson:"avatar,omitempty"`
	CreatedAt   time.Time  `bson:"createdAt"`
	UpdatedAt   time.Time  `bson:"updatedAt"`
	IsRoot      bool       `bson:"isRoot"`
	Permissions []string   `bson:"permissions"`
}

// GetAdminResponse ...
func (u *StaffRaw) GetStaffResponseAdmin() StaffGetResponseAdmin {
	return StaffGetResponseAdmin{
		ID:        u.ID,
		Username:  u.Username,
		Phone:     u.Phone,
		Active:    u.Active,
		Avatar:    u.Avatar.GetResponseData(),
		IsRoot:    u.IsRoot,
		CreatedAt: u.CreatedAt,
		Address:   u.Address,
	}
}

// GenerateToken generate token for authentication
func (u *StaffRaw) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      u.ID,
		"username": u.Username,
		"password": u.Password,
		"exp":      time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte(config.GetEnv().AuthSecret))
	return tokenString
}
