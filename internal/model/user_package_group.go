package model

// import (
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type UserPackageGroupAppService interface {
// 	ChoosePakage(ctx context.Context, body UserPackageBody) error
// }

// type UserPackageGroupDAO interface {
// 	FindOneByCondition(ctx context.Context, cond interface{}) (UserPackageGroupRaw, error)
// 	InsertOne(ctx context.Context, u UserPackageGroupRaw) error
// 	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]UserPackageGroupRaw, error)
// 	CountByCondition(ctx context.Context, cond interface{}) int64
// 	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
// }

// type UserPackageGroupRaw struct {
// 	ID             primitive.ObjectID `bson:"_id"`
// 	PackageID      primitive.ObjectID `bson:"packageId"`
// 	UserID         primitive.ObjectID `bson:"userId"`
// 	SubmissionTime time.Duration      `bson:"submissionTime"`
// 	TotalPoint     float64            `bson:"totalPoint"` // tong so diem
// 	IsPass         bool               `bson:"isPass"`
// 	TotalQuiz      float64            `bson:"totalQuiz"` // tong so cau hoi
// 	GroupID        primitive.ObjectID `bson:"groupID"`
// 	CreatedAt      time.Time          `bson:"createdAt"`
// 	UpdatedAt      time.Time          `bson:"updatedAt"`
// }
