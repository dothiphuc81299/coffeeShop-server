package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const packageCol = "packages"

// PackageDAO ....
type PackageDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewPackageDAO ...
func NewPackageDAO(db *mongo.Database) model.PackageDAO {
	return &PackageDAO{
		DB:  db,
		Col: db.Collection(packageCol),
	}
}

// InsertOne ...
func (d *PackageDAO) InsertOne(ctx context.Context, u model.PackageRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *PackageDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.PackageRaw, err error) {
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
func (w *PackageDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *PackageDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *PackageDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.PackageRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
