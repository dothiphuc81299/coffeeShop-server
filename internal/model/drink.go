package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DrinkDAO ...
type DrinkDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (DrinkRaw, error)
	InsertOne(ctx context.Context, u DrinkRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]DrinkRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// DrinkAdminService ...
type DrinkAdminService interface {
	Create(ctx context.Context, body DrinkBody) (primitive.ObjectID, error)
	ListAll(ctx context.Context, q CommonQuery) ([]DrinkAdminResponse, int64)
	Update(ctx context.Context, Drink DrinkRaw, body DrinkBody) error
	ChangeStatus(ctx context.Context, Drink DrinkRaw) error
	FindByID(ctx context.Context, id AppID) (Drink DrinkRaw, err error)
}

// DrinkRaw ...
type DrinkRaw struct {
	ID           AppID   `bson:"_id"`
	Name         string  `bson:"_id"`
	CategoryID   AppID   `bson:"categoryID"`
	Price        float64 `bson:"price"`
	SearchString string  `bson:"searchString"`
	FeedBack     AppID   `bson:"feedback"`
	Quantity     int     `bson:"quantity"`
	//	Photo        *FilePhoto `bson:"photo"`      // TODO
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}


