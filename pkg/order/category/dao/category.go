package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const categorycol = "categories"

// CategoryDAO ....
type CategoryDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewCategoryDAO ...
func NewCategoryDAO(db *mongo.Database) model.CategoryDAO {
	return &CategoryDAO{
		DB:  db,
		Col: db.Collection(categorycol),
	}
}

// InsertOne ...
func (d *CategoryDAO) InsertOne(ctx context.Context, u model.CategoryRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *CategoryDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.CategoryRaw, err error) {
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
func (w *CategoryDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *CategoryDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *CategoryDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.CategoryRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}

func (w *CategoryDAO) DeleteByID(ctx context.Context, id model.AppID) error {
	_, err := w.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
