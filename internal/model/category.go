package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// CategoryDAO ...
type CategoryDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (CategoryRaw, error)
	InsertOne(ctx context.Context, u CategoryRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]CategoryRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// CategoryAdminService ....
type CategoryAdminService interface {
	Create(ctx context.Context, body CategoryBody) (AppID, error)
	ListAll(ctx context.Context, q CommonQuery) ([]CategoryAdminResponse, int64)
	Update(ctx context.Context, c CategoryRaw, body CategoryBody) error
	//ChangeStatus(ctx context.Context, Category CategoryRaw) error
	FindByID(ctx context.Context, id AppID) (Category CategoryRaw, err error)
}

// CategoryRaw ....
type CategoryRaw struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
