package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EventDAO ...
type EventDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (EventRaw, error)
	InsertOne(ctx context.Context, u EventRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]EventRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error
	DeleteByID(ctx context.Context, id primitive.ObjectID) error
}

// EventAdminService ....
type EventAdminService interface {
	Create(ctx context.Context, body EventBody) (EventAdminResponse, error)
	ListAll(ctx context.Context, q CommonQuery) ([]EventAdminResponse, int64)
	Update(ctx context.Context, c EventRaw, body EventBody) (EventAdminResponse, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (Event EventRaw, err error)
	ChangeStatus(ctx context.Context, c EventRaw) error
	GetDetail(ctx context.Context, c EventRaw) EventAdminResponse
	DeleteEvent(ctx context.Context, c EventRaw) error
	SendEmail(ctx context.Context, c EventRaw) error
}

// EventRaw ....
type EventRaw struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Desc      string             `bson:"desc"`
	Active    bool               `bson:"active"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
