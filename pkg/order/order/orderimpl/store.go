package orderimpl

import (
	"context"
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/order"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ordercol = "orders"

type store struct {
	Col    *mongo.Collection
	Client *mongo.Client
}

func NewStore(db *mongodb.Database) *store {
	return &store{
		Col:    db.GetCollection(ordercol),
		Client: db.Client,
	}
}

func (s *store) InsertOne(ctx context.Context, u order.OrderRaw) error {
	_, err := s.Col.InsertOne(ctx, u)
	return err
}

func (s *store) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []order.OrderRaw, err error) {
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

func (s *store) FindOneByCondition(ctx context.Context, cond interface{}) (doc order.OrderRaw, err error) {
	err = s.Col.FindOne(ctx, cond).Decode(&doc)
	return
}

func (s *store) AggregateOrder(ctx context.Context, cond interface{}) ([]*order.StatisticByDrink, error) {
	var (
		results = make([]*order.StatisticByDrink, 0)
	)
	match := bson.M{
		"$match": cond,
	}

	project := bson.M{
		"$project": bson.M{
			"_id":           1,
			"totalQuantity": 1,
		},
	}

	group := bson.M{
		"$group": bson.M{
			"_id": "$drink._id",
			"totalQuantity": bson.M{
				"$sum": "$drink.quantity",
			},
		},
	}

	cursor, err := s.Col.Aggregate(ctx, []bson.M{match, project, group})
	if err != nil {
		fmt.Println("Error : ", err)
		return results, nil
	}

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &results)
	return results, err
}
