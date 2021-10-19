package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userPackageGroupCol = "user-package-group"

// UserPackageGroupDAO ....
type UserPackageGroupDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewUserPackageGroupDAO ...
func NewUserPackageGroupDAO(db *mongo.Database) model.UserPackageGroupDAO {
	return &UserPackageGroupDAO{
		DB:  db,
		Col: db.Collection(userPackageGroupCol),
	}
}

// InsertOne ...
func (d *UserPackageGroupDAO) InsertOne(ctx context.Context, u model.UserPackageGroupRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *UserPackageGroupDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.UserPackageGroupRaw, err error) {
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
func (w *UserPackageGroupDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *UserPackageGroupDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *UserPackageGroupDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.UserPackageGroupRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
