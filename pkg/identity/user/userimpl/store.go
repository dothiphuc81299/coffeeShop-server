package userimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usercol = "users"

type store struct {
	Col *mongo.Collection
}

func NewStore(db mongodb.DBConnector) *store {
	return &store{
		Col: db.GetCollection(usercol),
	}
}

func (s *store) InsertOne(ctx context.Context, u user.UserRaw) error {
	_, err := s.Col.InsertOne(ctx, u)
	return err
}

func (s *store) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []user.UserRaw, err error) {
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

func (s *store) FindOneByCondition(ctx context.Context, cond interface{}) (doc user.UserRaw, err error) {
	err = s.Col.FindOne(ctx, cond).Decode(&doc)
	return
}

func (s *store) UpdateByCondition(ctx context.Context, cond, payload interface{}) error {
	_, err := s.Col.UpdateOne(ctx, cond, payload)
	return err
}
