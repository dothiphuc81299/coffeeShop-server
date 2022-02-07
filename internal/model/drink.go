package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// DrinkDAO ...
type DrinkDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (DrinkRaw, error)
	InsertOne(ctx context.Context, u DrinkRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]DrinkRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
	DeleteByID(ctx context.Context, id AppID) error
	DeleteByCategoryID(ctx context.Context, categoryID AppID) error
}

// DrinkAdminService ...
type DrinkAdminService interface {
	Create(ctx context.Context, body DrinkBody) (DrinkAdminResponse, error)
	ListAll(ctx context.Context, q CommonQuery) ([]DrinkAdminResponse, int64)
	Update(ctx context.Context, Drink DrinkRaw, body DrinkBody) (DrinkAdminResponse, error)
	ChangeStatus(ctx context.Context, Drink DrinkRaw) (bool, error)
	FindByID(ctx context.Context, id AppID) (Drink DrinkRaw, err error)
	GetDetail(ctx context.Context, drink DrinkRaw) DrinkAdminResponse
	DeleteDrink(ctx context.Context, drink DrinkRaw) error
}

// DrinkRaw ...
type DrinkRaw struct {
	ID           AppID   `bson:"_id"`
	Name         string  `bson:"name"`
	Category     AppID   `bson:"category"`
	Price        float64 `bson:"price"`
	SearchString string  `bson:"searchString"`
	//	Photo        *FilePhoto `bson:"photo,omitempty"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Active    bool      `bson:"active"`
	Image     string    `bson:"image"`
}
