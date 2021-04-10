package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const shiftcol = "shifts"

// ShiftDAO ....
type ShiftDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewShiftDAO ...
func NewShiftDAO(db *mongo.Database) model.ShiftDAO {
	return &ShiftDAO{
		DB:  db,
		Col: db.Collection(shiftcol),
	}
}

// InsertOne ...
func (d *ShiftDAO) InsertOne(ctx context.Context, u model.ShiftRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *ShiftDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.ShiftRaw, err error) {
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
func (w *ShiftDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *ShiftDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *ShiftDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.ShiftRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
