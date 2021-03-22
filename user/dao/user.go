package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usercol = "users"

// UserDAO ....
type UserDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewUserDAO ...
func NewUserDAO(db *mongo.Database) model.UserDAO {
	return &UserDAO{
		DB:  db,
		Col: db.Collection(usercol),
	}
}

// InsertOne ...
func (d *UserDAO) InsertOne(ctx context.Context, u model.UserRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *UserDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.UserRaw, err error) {
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
func (w *UserDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *UserDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *UserDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.UserRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
