package dao

import (
	"context"
	"log"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const drinkCol = "drinks"

// DrinkDAO ....
type DrinkDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewDrinkDAO ...
func NewDrinkDAO(db *mongo.Database) model.DrinkDAO {
	return &DrinkDAO{
		DB:  db,
		Col: db.Collection(drinkCol),
	}
}

// InsertOne ...
func (d *DrinkDAO) InsertOne(ctx context.Context, u model.DrinkRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	log.Println("err", err)
	return err
}

// FindByCondition ...
func (w *DrinkDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DrinkRaw, err error) {
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
func (w *DrinkDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *DrinkDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *DrinkDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.DrinkRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
