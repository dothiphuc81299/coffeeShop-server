package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const accountcol = "accounts"

// AccountDAO ....
type AccountDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewAccountDAO ...
func NewAccountDAO(db *mongo.Database) model.AccountDAO {
	return &AccountDAO{
		DB:  db,
		Col: db.Collection(accountcol),
	}
}

// InsertOne ...
func (d *AccountDAO) InsertOne(ctx context.Context, u model.AccountRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *AccountDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.AccountRaw, err error) {
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
func (w *AccountDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *AccountDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *AccountDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.AccountRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
