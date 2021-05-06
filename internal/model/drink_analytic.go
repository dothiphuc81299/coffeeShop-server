package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DrinkAnalyticDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (DrinkAnalyticRaw, error)
	InsertOne(ctx context.Context, u DrinkAnalyticRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]DrinkAnalyticRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByCondition(ctx context.Context, cond interface{}, payload interface{}) error
	InsertMany(ctx context.Context, cond []interface{}) error
	AggregateDrink(ctx context.Context, cond interface{}) ([]DrinkAnalyticTempResponse, error)
}

type DrinkAnalyticService interface {
	ListAll(ctx context.Context, q CommonQuery) []DrinkAnalyticResponse
	//FindByID(ctx context.Context, id AppID) (drink DrinkAnalyticResponse, err error)
}

type DrinkAnalyticRaw struct {
	ID         primitive.ObjectID `bson:"_id"`
	TotalDrink float64            `bson:"totalDrink"`
	Name       primitive.ObjectID `bson:"name"`
	Category   primitive.ObjectID `bson:"category"`
	UpdateAt   time.Time          `bson:"updatedAt"`
	CreatedAt  time.Time          `bson:"createdAt"`
}

type DrinkAnalyticResponse struct {
	//ID       primitive.ObjectID `json:"_id"`
	Total    float64            `json:"total"`
	Drink    DrinkAnalyticInfo  `json:"drink"`
	Category CategoryInfo       `json:"category"`
	//UpdateAt time.Time          `json:"updatedAt"`
}

type DrinkAnalyticTempResponse struct {
	ID    primitive.ObjectID `json:"_id"`
	Total float64            `json:"total"`
}

type DrinkAnalyticInfo struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func (d *DrinkRaw) GetResponse() DrinkAnalyticInfo {
	return DrinkAnalyticInfo{
		ID:   d.ID,
		Name: d.Name,
	}
}

func (d *CategoryRaw) GetResponse() CategoryInfo {
	return CategoryInfo{
		ID:   d.ID,
		Name: d.Name,
	}
}
