package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const feedbackCol = "feedbacks"

// FeedbackDAO ....
type FeedbackDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewFeedbackDAO ...
func NewFeedbackDAO(db *mongo.Database) model.FeedbackDAO {
	return &FeedbackDAO{
		DB:  db,
		Col: db.Collection(feedbackCol),
	}
}

// InsertOne ...
func (d *FeedbackDAO) InsertOne(ctx context.Context, u model.FeedbackRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *FeedbackDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.FeedbackRaw, err error) {
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
func (w *FeedbackDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *FeedbackDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *FeedbackDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.FeedbackRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
