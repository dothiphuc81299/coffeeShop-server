package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const packageGroupCol = "package-group"

// PackageGroupDAO ....
type PackageGroupDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewPackageGroupDAO ...
func NewPackageGroupDAO(db *mongo.Database) model.PackageGroupDAO {
	return &PackageGroupDAO{
		DB:  db,
		Col: db.Collection(packageGroupCol),
	}
}

// InsertOne ...
func (d *PackageGroupDAO) InsertOne(ctx context.Context, u model.PackageGroupRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *PackageGroupDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.PackageGroupRaw, err error) {
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
func (w *PackageGroupDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *PackageGroupDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *PackageGroupDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.PackageGroupRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
