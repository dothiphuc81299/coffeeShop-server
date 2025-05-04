package model

// import (
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type PackageGroupDAO interface {
// 	FindOneByCondition(ctx context.Context, cond interface{}) (PackageGroupRaw, error)
// 	InsertOne(ctx context.Context, u PackageGroupRaw) error
// 	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]PackageGroupRaw, error)
// 	CountByCondition(ctx context.Context, cond interface{}) int64
// 	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
// }

// // PackageGroupAdminService ....
// type PackageGroupAdminService interface {
// 	Create(ctx context.Context, body PackageGroupBody) error
// 	//	ListAll(ctx context.Context, q CommonQuery) ([]PackageGroupAdminResponse, int64)
// 	Update(ctx context.Context, c PackageGroupRaw, body PackageGroupBody) error
// 	FindByID(ctx context.Context, id primitive.ObjectID) (PackageGroup PackageGroupRaw, err error)
// 	GetPackageGroupByPackageID(ctx context.Context, packageID primitive.ObjectID) []PackageGroupAdminResponse
// }

// type PackageGroupRaw struct {
// 	ID         primitive.ObjectID `bson:"_id"`
// 	PackageID  primitive.ObjectID `bson:"packageId"`
// 	GroupID    primitive.ObjectID `bson:"groupId"`
// 	NumberQuiz float64            `bson:"numberQuiz"`
// 	CreatedAt  time.Time          `bson:"createdAt"`
// }
