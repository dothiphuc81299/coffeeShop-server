package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const rolecol = "roles"

// RoleDAO ....
type RoleDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewRoleDAO ...
func NewRoleDAO(db *mongo.Database) model.RoleDAO {
	return &RoleDAO{
		DB:  db,
		Col: db.Collection(rolecol),
	}
}

// InsertOne ...
func (d *RoleDAO) InsertOne(ctx context.Context, u model.RoleRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *RoleDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.RoleRaw, err error) {
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
func (w *RoleDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *RoleDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *RoleDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.RoleRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
