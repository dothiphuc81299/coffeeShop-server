 package model

// import (
// 	"context"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // StaffDAO ...
// type StaffDAO interface {
// 	FindOneByCondition(ctx context.Context, cond interface{}) (Staff, error)
// 	InsertOne(ctx context.Context, u Staff) error
// 	FindByID(ctx context.Context, id AppID) (Staff, error)
// 	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]Staff, error)
// 	CountByCondition(ctx context.Context, cond interface{}) int64
// 	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
// 	UpdateBycondition(ctx context.Context, cond interface{}, payload interface{}) error
// 	DeleteByID(ctx context.Context, id AppID) error
// }

// // StaffAdminService ...
// type StaffAdminService interface {
// 	Create(ctx context.Context, body CreateStaffCommand) (StaffGetResponseAdmin, error)
// 	ListStaff(ctx context.Context, q CommonQuery) ([]StaffGetResponseAdmin, int64)
// 	FindByID(ctx context.Context, id AppID) (Staff, error)
// 	UpdateRole(ctx context.Context, body UpdateStaffRoleCommand, raw Staff) error
// 	DeleteStaff(ctx context.Context, raw Staff) error
// 	GetToken(ctx context.Context, staffID AppID) (string, error)
// 	GetDetailStaff(ctx context.Context, staff Staff) StaffMeResponse
// 	LoginStaff(ctx context.Context, LoginStaff LoginStaffCommand) (StaffResponse, error)
// 	GetStaffByID(ctx context.Context, id AppID) StaffGetResponseAdmin
// }

// type StaffAppService interface {
// 	Update(ctx context.Context, body UpdateStaffCommand, raw Staff) (StaffGetResponseAdmin, error)
// 	ChangePassword(ctx context.Context, staff Staff, body PasswordBody) error
// 	FindByID(ctx context.Context, id AppID) (Staff, error)
// 	GetDetailStaff(ctx context.Context, staff Staff) StaffMeResponse
// }

// // Staff ...
// type Staff struct {
// 	ID       AppID  `bson:"_id"`
// 	Username string `bson:"username"`
// 	Password string `bson:"password"`
// 	Address  string `bson:"address"`
// 	Phone    string `bson:"phone"`
// 	Active   bool   `bson:"active"`
// 	Role     AppID  `bson:"role,omitempty"`
// 	//Avatar      *FilePhoto `bson:"avatar,omitempty"`
// 	Avatar      string    `bson:"avatar,omitempty"`
// 	CreatedAt   time.Time `bson:"createdAt"`
// 	UpdatedAt   time.Time `bson:"updatedAt"`
// 	IsRoot      bool      `bson:"isRoot"`
// 	Permissions []string  `bson:"permissions"`
// }

// // GetAdminResponse ...
// func (u *Staff) GetStaffResponseAdmin() StaffGetResponseAdmin {
// 	return StaffGetResponseAdmin{
// 		ID:       u.ID,
// 		Username: u.Username,
// 		Phone:    u.Phone,
// 		Active:   u.Active,
// 		Role:     u.Role,
// 		//Avatar:      u.Avatar.GetResponseData(),
// 		Avatar: u.Avatar,
// 		IsRoot: u.IsRoot,
// 		CreatedAt: TimeResponse{
// 			Time: u.CreatedAt,
// 		},
// 		Address:     u.Address,
// 		Permissions: u.Permissions,
// 	}
// }

// // GenerateToken generate token for authentication
// func (u *Staff) GenerateToken() string {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"_id":      u.ID,
// 		"username": u.Username,
// 		"phone":    u.Phone,
// 		//	"exp":      time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
// 	})
// 	tokenString, _ := token.SignedString([]byte(config.GetEnv().AuthSecret))
// 	return tokenString
// }

// func (u *Staff) GetStaffResponse(token string) StaffResponse {
// 	return StaffResponse{
// 		ID:          u.ID,
// 		Username:    u.Username,
// 		Phone:       u.Phone,
// 		Address:     u.Address,
// 		Token:       token,
// 		Permissions: u.Permissions,
// 		//Avatar:      u.Avatar.GetResponseData(),
// 		Avatar: u.Avatar,
// 	}
// }
