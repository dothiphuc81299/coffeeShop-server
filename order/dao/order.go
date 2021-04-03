package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ordercol = "orders"

// OrderDAO ....
type OrderDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewOrderDAO ...
func NewOrderDAO(db *mongo.Database) model.OrderDAO {
	return &OrderDAO{
		DB:  db,
		Col: db.Collection(ordercol),
	}
}

// InsertOne ...
func (d *OrderDAO) InsertOne(ctx context.Context, u model.OrderRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *OrderDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.OrderRaw, err error) {
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
func (w *OrderDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *OrderDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *OrderDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.OrderRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
