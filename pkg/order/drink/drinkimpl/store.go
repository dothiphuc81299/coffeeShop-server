package drinkimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const drinkCol = "drinks"

type store struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

func NewStore(db *mongo.Database) *store {
	return &store{
		DB:  db,
		Col: db.Collection(drinkCol),
	}
}

func (s *store) InsertOne(ctx context.Context, u drink.DrinkRaw) error {
	_, err := s.Col.InsertOne(ctx, u)
	return err
}

func (s *store) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []drink.DrinkRaw, err error) {
	cursor, err := s.Col.Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

func (s *store) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := s.Col.CountDocuments(ctx, cond)
	return total
}

func (s *store) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	_, err := s.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

func (s *store) FindOneByCondition(ctx context.Context, cond interface{}) (doc drink.DrinkRaw, err error) {
	err = s.Col.FindOne(ctx, cond).Decode(&doc)
	return
}

func (s *store) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (s *store) DeleteByCategoryID(ctx context.Context, categoryID primitive.ObjectID) error {
	_, err := s.Col.DeleteOne(ctx, bson.M{"category": categoryID})
	return err
}
