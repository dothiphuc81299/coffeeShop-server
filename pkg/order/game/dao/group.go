package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const groupCol = "groups"

// GroupDAO ....
type GroupDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewGroupDAO ...
func NewGroupDAO(db *mongo.Database) model.GroupDAO {
	return &GroupDAO{
		DB:  db,
		Col: db.Collection(groupCol),
	}
}

// InsertOne ...
func (d *GroupDAO) InsertOne(ctx context.Context, u model.QuizGroupRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *GroupDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.QuizGroupRaw, err error) {
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
func (w *GroupDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *GroupDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *GroupDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.QuizGroupRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
