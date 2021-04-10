package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ShiftDAO ...
type ShiftDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (ShiftRaw, error)
	InsertOne(ctx context.Context, u ShiftRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]ShiftRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// ShiftAdminService ....
type ShiftAdminService interface {
	Create(ctx context.Context, body ShiftBody, staff StaffRaw) (ShiftResponse, error)
	ListAll(ctx context.Context, q CommonQuery) ([]ShiftResponse, int64)
	Update(ctx context.Context, c ShiftRaw, body ShiftBody, staff StaffRaw) (ShiftResponse, error)
	FindByID(ctx context.Context, id AppID) (Shift ShiftRaw, err error)
	AcceptShiftByAdmin(ctx context.Context, raw ShiftRaw) (bool, error)
}

// ShiftRaw ....
type ShiftRaw struct {
	ID        AppID              `bson:"_id"`
	Name      string             `bson:"name"`
	IsCheck   bool               `bson:"isCheck"`
	Date      time.Time          `bson:"date"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Staff     primitive.ObjectID `bson:"staff"`
}
