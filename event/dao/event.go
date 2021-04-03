package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const eventcol = "events"

// EventDAO ....
type EventDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewEventDAO ...
func NewEventDAO(db *mongo.Database) model.EventDAO {
	return &EventDAO{
		DB:  db,
		Col: db.Collection(eventcol),
	}
}

// InsertOne ...
func (d *EventDAO) InsertOne(ctx context.Context, u model.EventRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *EventDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.EventRaw, err error) {
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
func (w *EventDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *EventDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *EventDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.EventRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
