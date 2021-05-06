package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const drinkAnalyticCol = "drinkAnalytics"

// DrinkAnalyticDAO ....
type DrinkAnalyticDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewDrinkAnalyticDAO ...
func NewDrinkAnalyticDAO(db *mongo.Database) model.DrinkAnalyticDAO {
	return &DrinkAnalyticDAO{
		DB:  db,
		Col: db.Collection(drinkAnalyticCol),
	}
}

// InsertOne ...
func (d *DrinkAnalyticDAO) InsertOne(ctx context.Context, u model.DrinkAnalyticRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *DrinkAnalyticDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DrinkAnalyticRaw, err error) {
	cursor, err := w.Col.Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

// CountByCondition ...
func (w *DrinkAnalyticDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByCondition ...
func (w *DrinkAnalyticDAO) UpdateByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, cond, payload)
	return err
}

// FindOneByCondition ...
func (w *DrinkAnalyticDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.DrinkAnalyticRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}

func (w *DrinkAnalyticDAO) InsertMany(ctx context.Context, docs []interface{}) error {
	_, err := w.Col.InsertMany(ctx, docs)
	return err
}

func (w *DrinkAnalyticDAO) AggregateDrink(ctx context.Context, cond interface{}) ([]model.DrinkAnalyticTempResponse, error) {
	res := make([]model.DrinkAnalyticTempResponse, 0)
	// match := bson.M{
	// 	"$match": cond,
	// }
	group := bson.M{
		"$group": bson.M{
			"_id": "$name",
			"total": bson.M{
				"$sum": "$totalDrink",
			},
		},
	}
	cursor, err := w.Col.Aggregate(ctx, []bson.M{group})
	if err != nil {
		fmt.Println("Error : ", err)
	}

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &res)
	log.Println("err", res)
	return res, err

}
