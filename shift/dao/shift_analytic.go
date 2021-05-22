package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const shiftAnalyticcol = "shiftAnalytics"

// ShiftAnalyticDAO ....
type ShiftAnalyticDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewShiftAnalyticDAO ...
func NewShiftAnalyticDAO(db *mongo.Database) model.ShiftAnalyticDAO {
	return &ShiftAnalyticDAO{
		DB:  db,
		Col: db.Collection(shiftAnalyticcol),
	}
}

// InsertOne ...
func (d *ShiftAnalyticDAO) InsertOne(ctx context.Context, u model.ShiftAnalyticRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *ShiftAnalyticDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.ShiftAnalyticRaw, err error) {
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
func (w *ShiftAnalyticDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *ShiftAnalyticDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *ShiftAnalyticDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.ShiftAnalyticRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
