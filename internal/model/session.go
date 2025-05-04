package model

// import (
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // SessionDAO represent session data access object
// type SessionDAO interface {
// 	FindOneByCondition(ctx context.Context, cond interface{}) (SessionRaw, error)
// 	FindByID(ctx context.Context, id primitive.ObjectID) (SessionRaw, error)
// 	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]SessionRaw, error)
// 	CountByCondition(ctx context.Context, cond interface{}) int64
// 	InsertOne(ctx context.Context, u SessionRaw) error
// 	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
// 	RemoveByCondition(ctx context.Context, cond interface{}) error
// }

// // SessionAdminService ...
// type SessionAdminService interface {
// }

// // SessionRaw ...
// type SessionRaw struct {
// 	ID        primitive.ObjectID     `bson:"_id"`
// 	Staff     primitive.ObjectID     `bson:"staff"`
// 	Token     string    `bson:"token"`
// 	CreatedAt time.Time `bson:"createdAt"`
// }
