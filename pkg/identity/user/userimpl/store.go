package userimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usercol = "users"

type UserDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

func NewUserDAO(db *mongo.Database) user.UserDAO {
	return &UserDAO{
		DB:  db,
		Col: db.Collection(usercol),
	}
}

func (d *UserDAO) InsertOne(ctx context.Context, u user.UserRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

func (w *UserDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []user.UserRaw, err error) {
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

func (w *UserDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

func (w *UserDAO) UpdateByID(ctx context.Context, id user.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

func (w *UserDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc user.UserRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}

func (w *UserDAO) UpdateByCondition(ctx context.Context, cond, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, cond, payload)
	return err
}
